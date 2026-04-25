package parse_sql

import (
	"regexp"
	"slices"
	"strings"
)

// Struct for query
// Flags:
//	- "t" - table
//	- "i" - rows
type Cell struct{
	Query []string
	Flag string
	InValue [][]string // There is for data from request INSERT INTO
}

// Private function to delete brackets in string
func parseStack(query string) string{
	stack := 0
	start := -1

	for i, ch := range query{
		if ch == '('{
			if stack == 0 {
				start = i + 1
			}
			stack++
		} else if ch == ')'{
			stack--
			if stack == 0 && start != -1{
				return query[start:i]
			}
		}
	}
	return ""
}

// Public function to parse string
// This function returns flag and main data from query
// Query for request "CREATE TABLE":
// * first element it's a name table
// * second element it's a data about table
// Query for request "INSERT INTO":
// * first element it's a name table
// * second element colums
func Parse(sqlquery string) Cell{
	resultQuery := make([]string,0)
	spaceSkip := regexp.MustCompile(` +`)
	mode := spaceSkip.Split(sqlquery,-1)
	if slices.Contains(mode, "CREATE") && slices.Contains(mode, "TABLE"){
		resultQuery = append(resultQuery, mode[2])
		allCol := parseStack(sqlquery)
		resultQuery = append(resultQuery, allCol)
		return Cell{Query: resultQuery, Flag: "t"}
	}
	if slices.Contains(mode, "INSERT") && slices.Contains(mode,"INTO"){
		resultQuery = append(resultQuery, mode[2])
		inColStr := parseStack(sqlquery)
		inCol := splitCol(inColStr)
		resultQuery = append(resultQuery, inCol)
		valRe := regexp.MustCompile(`VALUES\s*\(([^)]+)\)`)
		valRow := valRe.FindAllStringSubmatch(sqlquery,-1)
		return Cell{Query: resultQuery, Flag: "i", InValue: valRow}
	}

	return Cell{}
}

// Private function for split "," with n quantity spaces
func splitCol(str string) string{
	re := regexp.MustCompile(`, +`)
	return re.ReplaceAllString(str," ")
}

// Private function for split insert data 
func splitInsertVal(val []string) [][]string{
	result := make([][]string,0)
	for _, v := range val{
		unBracketVal := parseStack(v)
		spaceSplitVal := splitCol(unBracketVal)
		result = append(result, strings.Split(spaceSplitVal," "))
	}
	return  result
}