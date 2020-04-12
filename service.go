package main

import "github.com/iancoleman/orderedmap"

type Service struct {
	db *DB
}

type StringMap map[string]interface{}

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

func (s *Service) RowsToOrderedMap(rows Rows) []*orderedmap.OrderedMap {
	data := make([]*orderedmap.OrderedMap, len(rows))

	for i, row := range rows {
		dataRow := orderedmap.New()

		for key := range *row {
			dataRow.Set(key, (*row)[key].Data)
		}

		data[i] = dataRow
	}

	return data
}
