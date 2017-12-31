package processor

import (
	"strings"

	"github.com/emirpasic/gods/lists/arraylist"
)

func addElement(current Clause, value string, hashMap map[Clause]*arraylist.List) {
	var cleanValue = strings.Replace(value, ";", "", -1)
	if clause, ok := hashMap[current]; ok {
		clause.Add(cleanValue)
	} else {
		hashMap[current] = arraylist.New()
		hashMap[current].Add(cleanValue)
	}
}

func Lexer(query string) map[Clause]*arraylist.List {
	var hashMap = make(map[Clause]*arraylist.List)
	var subSlices = strings.Split(query, " ")
	var currentClause = NotAClause
	for _, element := range subSlices {
		var result = evaluateWord(element)
		if result == EndQuery {
			break
		}
		if result != NotAClause {
			currentClause = result
		} else {
			addElement(currentClause, element, hashMap)
		}
	}
	return hashMap
}
