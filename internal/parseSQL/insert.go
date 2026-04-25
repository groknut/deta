package parse_sql

import (
	"slices"
	"strings"
)

// Struct for storing default value in column
type CellDefault struct{
	Title []string
	DefaultVal map[string]string
}

// Public function for getting default value
func ParseTableCell(sqlTable string) CellDefault{
	titleColumn := make([]string,0)
	defaultval := make(map[string]string)
	clearColumn := splitColumnTable(sqlTable)
	for _, q := range clearColumn{
		partsQuery := strings.Split(q, " ") 
		idxDEFAULT := slices.Index(partsQuery, "DEFAULT")
		if idxDEFAULT != -1{
			defaultval[partsQuery[0]] = partsQuery[idxDEFAULT+1]
		} else{
			titleColumn = append(titleColumn, partsQuery[0])
		}
	}

	return CellDefault{Title: titleColumn, DefaultVal: defaultval}
}

// Private function to split column of query
func splitColumnTable(query string) []string{
	var result []string
	var current strings.Builder
	parentCnt := 0
	for i := 0; i < len(query); i++{
		ch := query[i]

		switch ch {
		case '(':
			parentCnt++
			current.WriteByte(ch)
		case ')':
			parentCnt--
			current.WriteByte(ch)
		case ',':
			if parentCnt == 0{
				result = append(result, strings.TrimSpace(current.String()))
				current.Reset()
			} else {
				current.WriteByte(ch)
			}
		default:
			current.WriteByte(ch)
		}
	}

	if current.Len() > 0{
		result = append(result, strings.TrimSpace(current.String()))
	}
	return result
}


// Public function prefer of data for model. Also it fills default value
func AddRowsOfModel(insertInto Cell, defValue CellDefault) [][]string{
// 	resRows := make([][]string, 0)
// 	widthRow := len(defValue.Title)
// 	for _, row := range insertInto.InValue{
// 		tempRow := make([]string, widthRow)

// 	}

// 	return resRows
	return nil
}