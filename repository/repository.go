package repository

import (
	structs "backend_service/structs"
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func GetWords() (res []structs.WordStruct, err error) {
	var responseWord string
	rows, err := DB.Query("SELECT words FROM wordslist ORDER BY random() LIMIT 10")
	if err != nil {
		return res, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&responseWord); err != nil {
			return res, err
		}
		res = append(res, structs.WordStruct{
			Word: responseWord,
		})

	}

	return res, err
}
