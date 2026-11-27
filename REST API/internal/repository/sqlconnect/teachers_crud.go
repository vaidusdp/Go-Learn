package sqlconnect

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"restapi/internal/models"
	"strconv"
	"strings"
)

func queryFilter(r *http.Request, query string, args []interface{}) (string, []interface{}) {
	params := map[string]string{
		"FirstName": "FirstName",
		"LastName":  "LastName",
		"email":     "email",
		"class":     "class",
		"subject":   "subject",
	}

	for param, dbField := range params {
		value := r.URL.Query().Get(param)
		if value != "" {
			query += fmt.Sprintf(" AND %s = @p%d", dbField, len(args)+1)
			args = append(args, value)
		}
	}
	return query, args
}

func addSorting(r *http.Request, query string) string {
	sortParams := r.URL.Query()["sortBy"]
	if len(sortParams) > 0 {
		query += " ORDER BY"
		for i, param := range sortParams {
			parts := strings.Split(param, ":")
			if len(parts) != 2 {
				continue
			}

			field, order := parts[0], parts[1]
			if !isValidSortOrder(order) || !isValidSortFeilds(field) {
				continue
			}

			if i > 0 {
				query += ","
			}
			query += " " + field + " " + order
		}
	}
	return query
}

func isValidSortOrder(order string) bool {
	return order == "asc" || order == "desc"
}

func isValidSortFeilds(feild string) bool {
	validFeilds := map[string]bool{
		"FirstName": true,
		"LastName":  true,
		"email":     true,
		"class":     true,
		"subject":   true,
	}

	return validFeilds[feild]
}

func GetTeachersDbHandler(teachers []models.Teacher, r *http.Request) ([]models.Teacher, error) {
	db, err := ConnectDB()
	if err != nil {
		// http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return nil, err
	}
	defer db.Close()

	query := "SELECT id, FirstName, LastName, email, class, subject FROM teachers WHERE 1=1"
	var args []interface{}

	query, args = queryFilter(r, query, args)

	query = addSorting(r, query)

	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
		// http.Error(w, "Database Query Error", http.StatusInternalServerError)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var teacher models.Teacher
		err = rows.Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)
		if err != nil {
			// http.Error(w, "Error scanning database results", http.StatusInternalServerError)
			return nil, err
		}
		teachers = append(teachers, teacher)
	}
	return teachers, nil
}

func GetOneTeacherDbHandler(id int) (models.Teacher, error) {
	db, err := ConnectDB()
	if err != nil {
		// http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return models.Teacher{}, err
	}
	defer db.Close()

	var teacher models.Teacher
	err = db.QueryRow("SELECT id, FirstName, LastName, email, class, subject FROM teachers WHERE id = @p1", id).Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)
	if err == sql.ErrNoRows {
		// http.Error(w, "Teacher not found", http.StatusNotFound)
		return models.Teacher{}, err
	} else if err != nil {
		// http.Error(w, "Database query error", http.StatusInternalServerError)
		return models.Teacher{}, err
	}
	return teacher, nil
}

func AddTeacherDbHandler(newTeachers []models.Teacher) ([]models.Teacher, error) {
	db, err := ConnectDB()
	if err != nil {
		// http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return nil, err
	}
	defer db.Close()

	query := `
		INSERT INTO teachers (FirstName, LastName, email, class, subject)
		OUTPUT INSERTED.id
		VALUES (@p1, @p2, @p3, @p4, @p5)
	`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println("Prepare error:", err)
		// http.Error(w, "Error in preparing SQL Query", http.StatusBadRequest)
		return nil, err
	}
	defer stmt.Close()

	addedTeachers := make([]models.Teacher, len(newTeachers))
	for i, newTeacher := range newTeachers {
		err := stmt.QueryRow(newTeacher.FirstName, newTeacher.LastName, newTeacher.Email, newTeacher.Class, newTeacher.Subject).Scan(&newTeacher.ID)
		if err != nil {
			// http.Error(w, "Error inserting data into DB", http.StatusBadRequest)
			return nil, err
		}

		addedTeachers[i] = newTeacher
	}
	return addedTeachers, nil
}

func UpdateTeacherDbHandler(id int, updatedTeacher models.Teacher) (models.Teacher, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Unable to connect to database", http.StatusInternalServerError)
		return models.Teacher{}, err
	}
	defer db.Close()

	var existingTeacher models.Teacher
	err = db.QueryRow("SELECT id, FirstName, LastName, email, class, subject FROM teachers WHERE id = @p1", id).Scan(&existingTeacher.ID, &existingTeacher.FirstName, &existingTeacher.LastName, &existingTeacher.Email, &existingTeacher.Class, &existingTeacher.Subject)
	if err == sql.ErrNoRows {
		// http.Error(w, "Teacher Not Found", http.StatusNotFound)
		return models.Teacher{}, err
	} else if err != nil {
		// http.Error(w, "Unable to retrieve data", http.StatusInternalServerError)
		return models.Teacher{}, err
	}

	updatedTeacher.ID = existingTeacher.ID
	_, err = db.Exec("UPDATE teachers SET FirstName = @p1, LastName = @p2, email = @p3, class = @p4, subject = @p5 WHERE id = @p6", updatedTeacher.FirstName, updatedTeacher.LastName, updatedTeacher.Email, updatedTeacher.Class, updatedTeacher.Subject, updatedTeacher.ID)
	if err != nil {
		log.Println(err)
		// http.Error(w, "Error updating teacher", http.StatusInternalServerError)
		return models.Teacher{}, err
	}
	return updatedTeacher, nil
}

func PatchTeacherDbHandler(updates []map[string]interface{}) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Unable to connect to database", http.StatusInternalServerError)
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		// http.Error(w, "Error starting transaction", http.StatusInternalServerError)
	}

	for _, update := range updates {
		idStr, ok := update["id"].(string)
		if !ok {
			tx.Rollback()
			// http.Error(w, "Invalid Teacher Id in update", http.StatusBadRequest)
			return err
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			tx.Rollback()
			// http.Error(w, "Error converting id to int", http.StatusBadRequest)
			return err
		}

		var teacherFromDb models.Teacher
		err = db.QueryRow("SELECT id, FirstName, LastName, email, class, subject FROM teachers WHERE id = @p1", id).Scan(&teacherFromDb.ID, &teacherFromDb.FirstName, &teacherFromDb.LastName, &teacherFromDb.Email, &teacherFromDb.Class, &teacherFromDb.Subject)
		if err != nil {
			tx.Rollback()
			if err != sql.ErrNoRows {
				// http.Error(w, "Teacher not found", http.StatusNotFound)
				return err
			}

			// http.Error(w, "Teacher not found", http.StatusNotFound)
			return err
		}

		teacherVal := reflect.ValueOf(&teacherFromDb).Elem()
		teacherType := teacherVal.Type()

		for k, v := range update {
			if k == "id" {
				continue
			}

			for i := 0; i < teacherVal.NumField(); i++ {
				field := teacherType.Field(i)
				if field.Tag.Get("json") == k+",omitempty" {
					fieldVal := teacherVal.Field(i)
					if fieldVal.CanSet() {
						val := reflect.ValueOf(v)
						if val.Type().ConvertibleTo(fieldVal.Type()) {
							fieldVal.Set(val.Convert(fieldVal.Type()))
						} else {
							tx.Rollback()
							log.Panicf("Cannot convert %v to %v", val.Type(), fieldVal.Type())
						}
					}
				}
			}
		}
		_, err = tx.Exec("UPDATE teachers SET FirstName = @p1, LastName = @p2, email = @p3, class = @p4, subject = @p5 WHERE id = @p6", teacherFromDb.FirstName, teacherFromDb.LastName, teacherFromDb.Email, teacherFromDb.Class, teacherFromDb.Subject, teacherFromDb.ID)

		if err != nil {
			tx.Rollback()
			// http.Error(w, "Error Updating Teacher", http.StatusInternalServerError)
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		// http.Error(w, "Error commiting transaction", http.StatusInternalServerError)
		return err
	}
	return nil
}

func PatchOneTeacherDbHandler(id int, updates map[string]interface{}) (models.Teacher, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Unable to connect to database", http.StatusInternalServerError)
		return models.Teacher{}, err
	}
	defer db.Close()

	var existingTeacher models.Teacher
	err = db.QueryRow("SELECT id, FirstName, LastName, email, class, subject FROM teachers WHERE id = @p1", id).Scan(&existingTeacher.ID, &existingTeacher.FirstName, &existingTeacher.LastName, &existingTeacher.Email, &existingTeacher.Class, &existingTeacher.Subject)
	if err == sql.ErrNoRows {
		// http.Error(w, "Teacher Not Found", http.StatusNotFound)
		return models.Teacher{}, nil
	} else if err != nil {
		// http.Error(w, "Unable to retrieve data", http.StatusInternalServerError)
		return models.Teacher{}, nil
	}

	teacherVal := reflect.ValueOf(&existingTeacher).Elem()
	teacherType := teacherVal.Type()

	for k, v := range updates {
		for i := 0; i < teacherVal.NumField(); i++ {
			feild := teacherType.Field(i)

			if feild.Tag.Get("json") == k+",omitempty" {
				if teacherVal.Field(i).CanSet() {
					teacherVal.Field(i).Set(reflect.ValueOf(v).Convert(teacherVal.Field(i).Type()))
				}
			}
		}
	}

	_, err = db.Exec("UPDATE teachers SET FirstName = @p1, LastName = @p2, email = @p3, class = @p4, subject = @p5 WHERE id = @p6", existingTeacher.FirstName, existingTeacher.LastName, existingTeacher.Email, existingTeacher.Class, existingTeacher.Subject, existingTeacher.ID)
	if err != nil {
		log.Println(err)
		// http.Error(w, "Error updating teacher", http.StatusInternalServerError)
		return models.Teacher{}, nil
	}
	return existingTeacher, nil
}

func DeleteOneTeacherDbHandler(id int) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Unable to connect to database", http.StatusInternalServerError)
		return err
	}
	defer db.Close()

	res, err := db.Exec("DELETE FROM teachers WHERE id = @p1", id)
	if err != nil {
		// http.Error(w, "Error Deleting Teacher", http.StatusInternalServerError)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		// http.Error(w, "Error Retrieving Rows Teacher", http.StatusInternalServerError)
		return err
	}

	if rowsAffected == 0 {
		// http.Error(w, "Teacher not found", http.StatusNotFound)
		return err
	}
	return nil
}

func DeleteTeachersDbHandler(ids []int) ([]int, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Unable to connect to database", http.StatusInternalServerError)
		return nil, err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Panicln(err)
		// http.Error(w, "Error starting transaction", http.StatusInternalServerError)
		return nil, err
	}

	stmt, err := tx.Prepare("DELETE FROM teachers WHERE id = @p1")
	if err != nil {
		log.Panicln(err)
		tx.Rollback()
		// http.Error(w, "Error preparing delete statement", http.StatusInternalServerError)
		return nil, err
	}

	var deletedIds []int

	for _, id := range ids {
		res, err := stmt.Exec(id)
		if err != nil {
			tx.Rollback()
			log.Panicln(err)
			// http.Error(w, "Error deleting teacher", http.StatusInternalServerError)
			return nil, err
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			tx.Rollback()
			log.Panicln(err)
			// http.Error(w, "Error retrieving delete result", http.StatusInternalServerError)
			return nil, err
		}

		if rowsAffected > 0 {
			deletedIds = append(deletedIds, id)
		}

		if rowsAffected < 1 {
			tx.Rollback()
			// http.Error(w, fmt.Sprintf("Id %d Does Not Exist", id), http.StatusInternalServerError)
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Error commiting transaction", http.StatusInternalServerError)
		return nil, err
	}

	if len(deletedIds) < 1 {
		// http.Error(w, "Error retrieving delete result", http.StatusBadGateway)
		return nil, err
	}
	return deletedIds, nil
}