package dbdriver

import (
	"github.com/jinzhu/gorm"
	"time"
)

func QueryTopMapSlice(db *gorm.DB, sql string) (result []map[string]interface{}, err error) {
	rows, err := db.Raw(sql).Rows()
	if err != nil {
		return
	}
	cols, err := rows.Columns()
	if err != nil {
		return
	}

	for rows.Next() {
		// Create a slice of interface{}'s to represent each column,
		// and a second slice to contain pointers to each item in the columns slice.
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		// Scan the result into the column pointers...
		if err = rows.Scan(columnPointers...); err != nil {
			return
		}

		// Create our map, and retrieve the value for each column from the pointers slice,
		// storing it in the map with the name of the column as the key.
		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			switch (*val).(type) {
			case time.Time:
				m[colName] = (*val).(time.Time)
			default:
				m[colName] = string((*val).([]byte))
			}
		}

		// Outputs: map[columnName:value columnName2:value2 columnName3:value3 ...]
		result = append(result, m)
	}
	return
}
