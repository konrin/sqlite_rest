package main

type Service struct {
	db *DB
}

func NewService(db *DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) ExecRaw(sql string, args ...interface{}) (lastInsertId, rowsAffected int64, err error) {
	result, _err := s.db.Exec(sql, args...)
	if _err != nil {
		err = _err
		return
	}

	lastInsertId, err = result.LastInsertId()
	rowsAffected, err = result.RowsAffected()

	return
}

func (s *Service) QueryRaw(sql string, args ...interface{}) (Rows, error) {
	rows, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (s *Service) RowsToMap(rows Rows) []map[string]interface{} {
	data := make([]map[string]interface{}, len(rows))

	for i, row := range rows {
		dataRow := make(map[string]interface{})

		for key := range *row {
			dataRow[key] = (*row)[key].Data
		}

		data[i] = dataRow
	}

	return data
}
