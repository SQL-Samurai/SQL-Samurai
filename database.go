package database

import (
	"database/sql"
	"fmt"
	"regexp"
	"strconv"

	_ "github.com/lib/pq"
)

func RunQuery(databaseURL string, query string) ([]map[string]interface{}, error) {
	// Убедитесь, что запрос является запросом на чтение (SELECT), игнорируя регистр символов
	if ok, _ := regexp.MatchString(`(?i)^\s*SELECT\s`, query); !ok {
		return nil, fmt.Errorf("Only SELECT queries are allowed")
	}

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Failed to execute query: %v", err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve column names: %v", err)
	}

	var result []map[string]interface{}
	for rows.Next() {
		row := make(map[string]interface{})
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, fmt.Errorf("Failed to scan row values: %v", err)
		}

		for i, col := range columns {
			val := values[i]

			// Обработка типа данных numeric
			if b, ok := val.([]byte); ok {
				strVal := string(b)
				floatVal, err := strconv.ParseFloat(strVal, 64)
				if err != nil {
					row[col] = strVal // если не удалось преобразовать, оставляем строку
				} else {
					row[col] = floatVal
				}
			} else {
				row[col] = val
			}
		}
		result = append(result, row)
	}

	return result, nil
}
