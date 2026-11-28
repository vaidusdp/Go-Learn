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
)

func GetStudentsDbHandler(students []models.Student, r *http.Request) ([]models.Student, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error retrieving data")
	}
	defer db.Close()

	query := "SELECT id, FirstName, LastName, email, class, subject FROM students WHERE 1=1"
	var args []interface{}

	query, args = utils.QueryFilter(r, query, args)

	query = utils.AddSorting(r, query)

	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
		return nil, utils.ErrorHandler(err, "Database Query Error")
	}
	defer rows.Close()

	for rows.Next() {
		var student models.Student
		err = rows.Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email, &student.Class)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error scanning database results")
		}
		students = append(students, student)
	}
	return students, nil
}

func GetOneStudentDbHandler(id int) (models.Student, error) {
	db, err := ConnectDB()
	if err != nil {
		return models.Student{}, utils.ErrorHandler(err, "Error connecting to database")
	}
	defer db.Close()

	var student models.Student
	err = db.QueryRow("SELECT id, FirstName, LastName, email, class, subject FROM students WHERE id = @p1", id).Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email, &student.Class)
	if err == sql.ErrNoRows {
		return models.Student{}, utils.ErrorHandler(err, "Student not found")
	} else if err != nil {
		return models.Student{}, utils.ErrorHandler(err, "Database query error")
	}
	return student, nil
}

func AddStudentDbHandler(newStudents []models.Student) ([]models.Student, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error connecting to database")
	}
	defer db.Close()

	// query := `
	// 	INSERT INTO students (FirstName, LastName, email, class, subject)
	// 	OUTPUT INSERTED.id
	// 	VALUES (@p1, @p2, @p3, @p4, @p5)
	// `
	// stmt, err := db.Prepare(query)
	stmt, err := db.Prepare(utils.GenerateInsertQuery("students", models.Student{}))
	if err != nil {
		log.Println("Prepare error:", err)
		return nil, utils.ErrorHandler(err, "Error in preparing SQL Query")
	}
	defer stmt.Close()

	addedStudents := make([]models.Student, len(newStudents))
	for i, newStudent := range newStudents {
		// res, err := stmt.Exec(newTeacher.FirstName, newTeacher.LastName, newTeacher.Email, newTeacher.Class, newTeacher.Subject)
		values := utils.GetStructValues(newStudent)
		var lastId int64
		err := stmt.QueryRow(values...).Scan(&lastId)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error inserting data into DB")
		}

		newStudent.ID = int(lastId)
		addedStudents[i] = newStudent
	}
	return addedStudents, nil
}

func UpdateStudentDbHandler(id int, updatedStudent models.Student) (models.Student, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		return models.Student{}, utils.ErrorHandler(err, "Unable to connect to database")
	}
	defer db.Close()

	var existingStudent models.Student
	err = db.QueryRow("SELECT id, FirstName, LastName, email, class, subject FROM students WHERE id = @p1", id).Scan(&existingStudent.ID, &existingStudent.FirstName, &existingStudent.LastName, &existingStudent.Email, &existingStudent.Class)
	if err == sql.ErrNoRows {
		return models.Student{}, utils.ErrorHandler(err, "Student Not Found")
	} else if err != nil {
		return models.Student{}, utils.ErrorHandler(err, "Unable to retrieve data")
	}

	updatedStudent.ID = existingStudent.ID
	_, err = db.Exec("UPDATE students SET FirstName = @p1, LastName = @p2, email = @p3, class = @p4 WHERE id = @p5", updatedStudent.FirstName, updatedStudent.LastName, updatedStudent.Email, updatedStudent.Class, updatedStudent.ID)
	if err != nil {
		log.Println(err)
		return models.Student{}, err
	}
	return updatedStudent, nil
}

func PatchStudentDbHandler(updates []map[string]interface{}) error {
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
			return utils.ErrorHandler(err, "Invalid Student Id in update")
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			tx.Rollback()
			return utils.ErrorHandler(err, "Error converting id to int")
		}

		var studentFromDb models.Student
		err = db.QueryRow("SELECT id, FirstName, LastName, email, class FROM students WHERE id = @p1", id).Scan(&studentFromDb.ID, &studentFromDb.FirstName, &studentFromDb.LastName, &studentFromDb.Email, &studentFromDb.Class)
		if err != nil {
			tx.Rollback()
			if err != sql.ErrNoRows {
				return utils.ErrorHandler(err, "Student not found")
			}

			return utils.ErrorHandler(err, "Student not found")
		}

		studentVal := reflect.ValueOf(&studentFromDb).Elem()
		studentType := studentVal.Type()

		for k, v := range update {
			if k == "id" {
				continue
			}

			for i := 0; i < studentVal.NumField(); i++ {
				field := studentType.Field(i)
				if field.Tag.Get("json") == k+",omitempty" {
					fieldVal := studentVal.Field(i)
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
		_, err = tx.Exec("UPDATE students SET FirstName = @p1, LastName = @p2, email = @p3, class = @p4 WHERE id = @p5", studentFromDb.FirstName, studentFromDb.LastName, studentFromDb.Email, studentFromDb.Class, studentFromDb.ID)

		if err != nil {
			tx.Rollback()
			return utils.ErrorHandler(err, "Error Updating Student")
		}
	}

	err = tx.Commit()
	if err != nil {
		return utils.ErrorHandler(err, "Error commiting transaction")
	}
	return nil
}

func PatchOneStudentDbHandler(id int, updates map[string]interface{}) (models.Student, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		return models.Student{}, utils.ErrorHandler(err, "Unable to connect to database")
	}
	defer db.Close()

	var existingStudent models.Student
	err = db.QueryRow("SELECT id, FirstName, LastName, email, class FROM students WHERE id = @p1", id).Scan(&existingStudent.ID, &existingStudent.FirstName, &existingStudent.LastName, &existingStudent.Email, &existingStudent.Class)
	if err == sql.ErrNoRows {
		return models.Student{}, utils.ErrorHandler(err, "Student Not Found")
	} else if err != nil {
		return models.Student{}, utils.ErrorHandler(err, "Unable to retrieve data")
	}

	studentVal := reflect.ValueOf(&existingStudent).Elem()
	studentType := studentVal.Type()

	for k, v := range updates {
		for i := 0; i < studentVal.NumField(); i++ {
			feild := studentType.Field(i)

			if feild.Tag.Get("json") == k+",omitempty" {
				if studentVal.Field(i).CanSet() {
					studentVal.Field(i).Set(reflect.ValueOf(v).Convert(studentVal.Field(i).Type()))
				}
			}
		}
	}

	_, err = db.Exec("UPDATE students SET FirstName = @p1, LastName = @p2, email = @p3, class = @p4 WHERE id = @p5", existingStudent.FirstName, existingStudent.LastName, existingStudent.Email, existingStudent.Class, existingStudent.ID)
	if err != nil {
		log.Println(err)
		return models.Student{}, utils.ErrorHandler(err, "Error updating student")
	}
	return existingStudent, nil
}

func DeleteOneStudentDbHandler(id int) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		return utils.ErrorHandler(err, "Unable to connect to database")
	}
	defer db.Close()

	res, err := db.Exec("DELETE FROM students WHERE id = @p1", id)
	if err != nil {
		return utils.ErrorHandler(err, "Error Deleting Student")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return utils.ErrorHandler(err, "Error Retrieving Rows Student")
	}

	if rowsAffected == 0 {
		return utils.ErrorHandler(err, "Student not found")
	}
	return nil
}

func DeleteStudentsDbHandler(ids []int) ([]int, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Unable to connect to database")
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error starting transaction")
	}

	stmt, err := tx.Prepare("DELETE FROM students WHERE id = @p1")
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
			return nil, utils.ErrorHandler(err, "Error deleting student")
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
