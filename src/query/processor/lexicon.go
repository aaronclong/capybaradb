package processor

import "strings"

// Clause SQL Statement Markers
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

// EvaluateWord converts SQL clause to in memory value
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
