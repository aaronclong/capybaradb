package query_processor

import (
	"container/list"
	"strings"
	"github.com/emirpasic/gods/lists/arraylist"
)

func lexer(query string) map[byte]list.List {
	var hashMap = make(map[byte]list.List)
	var subSlices = strings.Split(query, " ")
	var currentItem = NOT_A_CLAUSE
	for _, element := range subSlices {
		var result = EvaluateWord(element)
		if result != NOT_A_CLAUSE {

		} else {
			currentItem = result
			if hashMap[:result] == nil {
				hashMap[result] := list.list.New()
			}
		}
	}
	return hashMap
}
