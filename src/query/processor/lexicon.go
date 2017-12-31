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
	EndQuery
)

// evaluateWord converts SQL clause to in memory value
func evaluateWord(word string) Clause {
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
	case ";":
		return EndQuery
	default:
		return NotAClause
	}
}
