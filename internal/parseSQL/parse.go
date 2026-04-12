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

func Parse(sqlquery string) Cell{
	resultQuery := make([]string,0)
	spaceSkip := regexp.MustCompile(` +`)
	mode := spaceSkip.Split(sqlquery,-1)
	if slices.Contains(mode, "CREATE") && slices.Contains(mode, "TABLE"){
		resultQuery = append(resultQuery, mode[2])
		takeColums := regexp.MustCompile(`\(([^)]+)\)`)
		allCol := takeColums.FindStringSubmatch(sqlquery)[0]
		resultQuery = append(resultQuery, allCol)
		return Cell{Query: resultQuery, Flag: "t"}
	}
	if slices.Contains(mode, "INSERT") && slices.Contains(mode,"INTO"){

	}

	return Cell{}
}