package processor

import (
	"testing"
)

func TestLexerHappyPath(t *testing.T) {
	const sqlString = "SELECT * FROM db;"
	var lexedValues = Lexer(sqlString)
	selectClause := lexedValues[Select]
	if !selectClause.Contains("*") {
		t.Errorf("Lexer Happy Path missing SELECT *")
	}
	fromClause := lexedValues[From]
	if !fromClause.Contains("db") {
		t.Errorf("Lexer Happy Path missing FROM db")
	}
}

func TestLexerComplicatedQuery(t *testing.T) {
	const sqlString = "SELECT Table, Owner FROM db WHERE table.name = 'aaron';"
	var lexedValues = Lexer(sqlString)
	selectClause := lexedValues[Select]
	if !selectClause.Contains("Table,") && !selectClause.Contains("Owner") {
		t.Errorf("Lexer Complicated Query missing SELECT clause")
	}
	fromClause := lexedValues[From]
	if !fromClause.Contains("db") {
		t.Errorf("Lexer Complicated Query missing FROM db")
	}

	whereClause := lexedValues[Where]
	if !whereClause.Contains("table.name") && !whereClause.Contains("=") && !whereClause.Contains("'aaron'") {
		t.Errorf("Lexer Complicated Query missing FROM db")
	}
}
