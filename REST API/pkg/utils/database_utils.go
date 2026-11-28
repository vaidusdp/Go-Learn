package utils

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

func GenerateInsertQuery(tableName string, model interface{}) string {
	modelType := reflect.TypeOf(model)

	var columns, placeholder string
	count := 0

	for i := 0; i < modelType.NumField(); i++ {
		dbTag := modelType.Field(i).Tag.Get("db")
		dbTag = strings.TrimSuffix(dbTag, ",omitempty")

		if dbTag != "" && dbTag != "id" {
			if count > 0 {
				columns += ", "
				placeholder += ", "
			}

			count++
			columns += dbTag
			placeholder += fmt.Sprintf("@p%d", count)
		}
	}

	return fmt.Sprintf("INSERT INTO %s (%s) OUTPUT INSERTED.id VALUES (%s)", tableName, columns, placeholder)
}

func GetStructValues(model interface{}) []interface{} {
	modelValue := reflect.ValueOf(model)
	modelType := modelValue.Type()
	values := []interface{}{}

	for i := 0; i < modelType.NumField(); i++ {
		dbTag := modelType.Field(i).Tag.Get("db")
		dbTag = strings.TrimSuffix(dbTag, ",omitempty")

		if dbTag != "" && dbTag != "id" {
			values = append(values, modelValue.Field(i).Interface())
		}
	}

	return values
}

func QueryFilter(r *http.Request, query string, args []interface{}) (string, []interface{}) {
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

func AddSorting(r *http.Request, query string) string {
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