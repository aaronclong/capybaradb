package processor

import "strings"

type Clause byte

const (
	Select Clause = iota
	Update
	Insert
	From
	Where
	GroupBy
	Equals
	NotAClause
)

func EvaluateWord(word string) Clause {
	switch strings.ToUpper(word) {
	case "SELECT":
		return Select
	case "FROM":
		return From
	case "WHERE":
		return Where
	case "GROUP BY":
		return GroupBy
	case "=":
		return Equals
	default:
		return NotAClause
	}
}
