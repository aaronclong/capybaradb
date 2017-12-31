package processor

import (
	"testing"
)

func TestLexerHappyPath(t *testing.T) {
	const sqlString = "SELECT * FROM db;"
	var lexedValues = Lexer(sqlString)
	selectClause := lexedValues[Select]
	if selectClause.Contains("*") != true {
		t.Errorf("Lexer Happy Path missing SELECT *")
	}
	fromClause := lexedValues[From]
	if fromClause.Contains("db") != true {
		t.Errorf("Lexer Happy Path missing FROM db")
	}
}
