package processor

import (
	"strings"

	"github.com/emirpasic/gods/lists/arraylist"
)

func lexer(query string) map[Clause]*arraylist.List {
	var hashMap = make(map[Clause]*arraylist.List)
	var subSlices = strings.Split(query, " ")
	var currentClause = NOT_A_CLAUSE
	for _, element := range subSlices {
		var result = EvaluateWord(element)
		if result != NOT_A_CLAUSE {
			currentClause = result
		} else {
			if clause, ok := hashMap[currentClause]; ok {
				clause.Add(currentClause)
			} else {
				hashMap[currentClause] = arraylist.New()
			}
		}
	}
	return hashMap
}
