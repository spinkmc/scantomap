package scantomap

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ScanToMap(db *sql.DB, query string, args ...any) (rows []map[string]any, err error) {

	var db_rows *sql.Rows

	if len(args) > 0 {
		db_rows, err = db.Query(query, args...)
	} else {
		db_rows, err = db.Query(query)
	}
	if err != nil {
		return nil, err
	}
	columns, _ := db_rows.Columns()
	for db_rows.Next() {
		scans := make([]any, len(columns))
		row := make(map[string]any)

		for i := range scans {
			scans[i] = &scans[i]
		}
		db_rows.Scan(scans...)
		for i, v := range scans {
			var value = ""
			if v != nil {
				value = fmt.Sprintf("%s", v)
			}
			row[columns[i]] = value
		}
		rows = append(rows, row)
	}
	return
}
