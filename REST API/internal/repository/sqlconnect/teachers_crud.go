package sqlconnect

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"restapi/internal/models"
	"restapi/pkg/utils"
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
		return nil, utils.ErrorHandler(err, "Error retrieving data")
	}
	defer db.Close()

	query := "SELECT id, FirstName, LastName, email, class, subject FROM teachers WHERE 1=1"
	var args []interface{}

	query, args = queryFilter(r, query, args)

	query = addSorting(r, query)

	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
		return nil, utils.ErrorHandler(err, "Database Query Error")
	}
	defer rows.Close()

	for rows.Next() {
		var teacher models.Teacher
		err = rows.Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error scanning database results")
		}
		teachers = append(teachers, teacher)
	}
	return teachers, nil
}

func GetOneTeacherDbHandler(id int) (models.Teacher, error) {
	db, err := ConnectDB()
	if err != nil {
		return models.Teacher{}, utils.ErrorHandler(err, "Error connecting to database")
	}
	defer db.Close()

	var teacher models.Teacher
	err = db.QueryRow("SELECT id, FirstName, LastName, email, class, subject FROM teachers WHERE id = @p1", id).Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)
	if err == sql.ErrNoRows {
		return models.Teacher{}, utils.ErrorHandler(err, "Teacher not found")
	} else if err != nil {
		return models.Teacher{}, utils.ErrorHandler(err, "Database query error")
	}
	return teacher, nil
}

func AddTeacherDbHandler(newTeachers []models.Teacher) ([]models.Teacher, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error connecting to database")
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
		return nil, utils.ErrorHandler(err, "Error in preparing SQL Query")
	}
	defer stmt.Close()

	addedTeachers := make([]models.Teacher, len(newTeachers))
	for i, newTeacher := range newTeachers {
		err := stmt.QueryRow(newTeacher.FirstName, newTeacher.LastName, newTeacher.Email, newTeacher.Class, newTeacher.Subject).Scan(&newTeacher.ID)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error inserting data into DB")
		}

		addedTeachers[i] = newTeacher
	}
	return addedTeachers, nil
}

func UpdateTeacherDbHandler(id int, updatedTeacher models.Teacher) (models.Teacher, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		return models.Teacher{}, utils.ErrorHandler(err, "Unable to connect to database")
	}
	defer db.Close()

	var existingTeacher models.Teacher
	err = db.QueryRow("SELECT id, FirstName, LastName, email, class, subject FROM teachers WHERE id = @p1", id).Scan(&existingTeacher.ID, &existingTeacher.FirstName, &existingTeacher.LastName, &existingTeacher.Email, &existingTeacher.Class, &existingTeacher.Subject)
	if err == sql.ErrNoRows {
		return models.Teacher{}, utils.ErrorHandler(err, "Teacher Not Found")
	} else if err != nil {
		return models.Teacher{}, utils.ErrorHandler(err, "Unable to retrieve data")
	}

	updatedTeacher.ID = existingTeacher.ID
	_, err = db.Exec("UPDATE teachers SET FirstName = @p1, LastName = @p2, email = @p3, class = @p4, subject = @p5 WHERE id = @p6", updatedTeacher.FirstName, updatedTeacher.LastName, updatedTeacher.Email, updatedTeacher.Class, updatedTeacher.Subject, updatedTeacher.ID)
	if err != nil {
		log.Println(err)
		return models.Teacher{}, err
	}
	return updatedTeacher, nil
}

func PatchTeacherDbHandler(updates []map[string]interface{}) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		return utils.ErrorHandler(err, "Unable to connect to database")
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return utils.ErrorHandler(err, "Error starting transaction")
	}

	for _, update := range updates {
		idStr, ok := update["id"].(string)
		if !ok {
			tx.Rollback()
			return utils.ErrorHandler(err, "Invalid Teacher Id in update")
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			tx.Rollback()
			return utils.ErrorHandler(err, "Error converting id to int")
		}

		var teacherFromDb models.Teacher
		err = db.QueryRow("SELECT id, FirstName, LastName, email, class, subject FROM teachers WHERE id = @p1", id).Scan(&teacherFromDb.ID, &teacherFromDb.FirstName, &teacherFromDb.LastName, &teacherFromDb.Email, &teacherFromDb.Class, &teacherFromDb.Subject)
		if err != nil {
			tx.Rollback()
			if err != sql.ErrNoRows {
				return utils.ErrorHandler(err, "Teacher not found")
			}

			return utils.ErrorHandler(err, "Teacher not found")
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
			return utils.ErrorHandler(err, "Error Updating Teacher")
		}
	}

	err = tx.Commit()
	if err != nil {
		return utils.ErrorHandler(err, "Error commiting transaction")
	}
	return nil
}

func PatchOneTeacherDbHandler(id int, updates map[string]interface{}) (models.Teacher, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		return models.Teacher{}, utils.ErrorHandler(err, "Unable to connect to database")
	}
	defer db.Close()

	var existingTeacher models.Teacher
	err = db.QueryRow("SELECT id, FirstName, LastName, email, class, subject FROM teachers WHERE id = @p1", id).Scan(&existingTeacher.ID, &existingTeacher.FirstName, &existingTeacher.LastName, &existingTeacher.Email, &existingTeacher.Class, &existingTeacher.Subject)
	if err == sql.ErrNoRows {
		return models.Teacher{}, utils.ErrorHandler(err, "Teacher Not Found")
	} else if err != nil {
		return models.Teacher{}, utils.ErrorHandler(err, "Unable to retrieve data")
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
		return models.Teacher{}, utils.ErrorHandler(err, "Error updating teacher")
	}
	return existingTeacher, nil
}

func DeleteOneTeacherDbHandler(id int) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		return utils.ErrorHandler(err, "Unable to connect to database")
	}
	defer db.Close()

	res, err := db.Exec("DELETE FROM teachers WHERE id = @p1", id)
	if err != nil {
		return utils.ErrorHandler(err, "Error Deleting Teacher")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return utils.ErrorHandler(err, "Error Retrieving Rows Teacher")
	}

	if rowsAffected == 0 {
		return utils.ErrorHandler(err, "Teacher not found")
	}
	return nil
}

func DeleteTeachersDbHandler(ids []int) ([]int, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Unable to connect to database")
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error starting transaction")
	}

	stmt, err := tx.Prepare("DELETE FROM teachers WHERE id = @p1")
	if err != nil {
		tx.Rollback()
		return nil, utils.ErrorHandler(err, "Error preparing delete statement")
	}

	var deletedIds []int

	for _, id := range ids {
		res, err := stmt.Exec(id)
		if err != nil {
			tx.Rollback()
			log.Panicln(err)
			return nil, utils.ErrorHandler(err, "Error deleting teacher")
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			tx.Rollback()
			log.Panicln(err)
			return nil, utils.ErrorHandler(err, "Error retrieving delete result")
		}

		if rowsAffected > 0 {
			deletedIds = append(deletedIds, id)
		}

		if rowsAffected < 1 {
			tx.Rollback()
			return nil, utils.ErrorHandler(err, fmt.Sprintf("Id %d Does Not Exist", id))
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return nil, utils.ErrorHandler(err, "Error commiting transaction")
	}

	if len(deletedIds) < 1 {
		return nil, utils.ErrorHandler(err, "Error retrieving delete result")
	}
	return deletedIds, nil
}
