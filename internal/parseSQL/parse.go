package parse_sql

import (
	"regexp"
	"slices"
)

// flags:
//	- "t" - table
//	- "r" - rows
type Cell struct{
	Query []string
	Flag string
}

func ParseStack(query string) string{
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

func Parse(sqlquery string) Cell{
	resultQuery := make([]string,0)
	spaceSkip := regexp.MustCompile(` +`)
	mode := spaceSkip.Split(sqlquery,-1)
	if slices.Contains(mode, "CREATE") && slices.Contains(mode, "TABLE"){
		resultQuery = append(resultQuery, mode[2])
		// takeColums := regexp.MustCompile(`\(([^)]+)\)`)
		// takeColums := regexp.MustCompile(`\((.*?)\)`)
		// allCol := takeColums.FindStringSubmatch(sqlquery)[1]
		allCol := ParseStack(sqlquery)
		resultQuery = append(resultQuery, allCol)
		return Cell{Query: resultQuery, Flag: "t"}
	}
	if slices.Contains(mode, "INSERT") && slices.Contains(mode,"INTO"){

	}

	return Cell{}
}