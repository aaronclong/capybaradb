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

// Operation refers to math matical operation
type Operation byte

const (
	Equals Operation = iota
	Multiply
	Divide
	Addition
	Subtract
	WildCard
	Noperation
)

func evaluateForClause(word string) Clause {
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

func evaluateForOperation(statement string) Operation {
	switch statement {
	case "*":
		return WildCard
	case "=":
		return Equals
	case "\\":
		return Divide
	case "+"
		return Addition
	case "-"
		return Subtract
	default:
		return Noperation
	}
}
