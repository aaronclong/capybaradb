package processor

import "strings"

type Clause byte

const (
	SELECT Clause = iota
	UPDATE
	INSERT
	FROM
	WHERE
	GROUP_BY
	EQUALS
	NOT_A_CLAUSE
)

func EvaluateWord(word string) Clause {
	switch strings.ToUpper(word) {
	case "SELECT":
		return SELECT
	case "FROM":
		return FROM
	case "WHERE":
		return WHERE
	case "GROUP BY":
		return GROUP_BY
	case "=":
		return EQUALS
	default:
		return NOT_A_CLAUSE
	}
}
