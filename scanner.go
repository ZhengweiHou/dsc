package dsc

import "github.com/viant/toolbox"

type scanner struct {
	scanner Scanner
}

func (s *scanner) Columns() ([]string, error) {
	return s.scanner.Columns()
}

func (s *scanner) ColumnTypes() ([]ColumnType, error) {
	return s.scanner.ColumnTypes()
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
	return &scanner{s}
}
