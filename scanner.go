package dsc

import (
	"github.com/viant/toolbox"
)

type scanner struct {
	scanner             Scanner
	columns             []string
	columnTypes         []ColumnType
	columnsInitFlag     chan bool
	columnTypesInitFlag chan bool
}

func (s *scanner) Columns() ([]string, error) {
	if s.columns == nil {
		_, ok := <-s.columnsInitFlag // 并发控制
		if ok {
			columns, err := s.scanner.Columns()
			if err != nil {
				return nil, err
			}
			s.columns = columns
			close(s.columnsInitFlag) // 通知其他进程
		} else { // 被唤醒
			return s.Columns()
		}
	}
	return s.columns, nil
}

func (s *scanner) ColumnTypes() ([]ColumnType, error) {
	if s.columnTypes == nil {
		_, ok := <-s.columnTypesInitFlag
		if ok {
			columnTypes, err := s.scanner.ColumnTypes()
			if err != nil {
				return nil, err
			}
			s.columnTypes = columnTypes
			close(s.columnTypesInitFlag)
		} else {
			return s.ColumnTypes()
		}
	}

	return s.columnTypes, nil
}

func (s *scanner) Scan(destinations ...interface{}) error {
	if len(destinations) == 1 { // 当接受变量是map时
		if toolbox.IsMap(destinations[0]) {
			aMap := toolbox.AsMap(destinations[0]) // 为什么要再转一次
			values, columns, err := ScanRow(s)
			if err != nil {
				return err
			}
			for i, column := range columns {
				aMap[column] = values[i]
			}
			return nil
		}
	}
	err := s.scanner.Scan(destinations...) // 此Scan和当前方法有什么区别？ 调试时直接调用到
	return err
}

func NewScanner(s Scanner) Scanner {
	columnsInitFlag := make(chan bool, 1)
	columnTypesInitFlag := make(chan bool, 1)
	columnsInitFlag <- true
	columnTypesInitFlag <- true

	return &scanner{
		scanner:             s,
		columnsInitFlag:     columnsInitFlag,
		columnTypesInitFlag: columnTypesInitFlag,
	}
}
