package chipres_test

import (
	"github.com/chipres/chipres_sintactico"
	"strings"
	"testing"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s   string
		tok chipres.Token
		lit string
	}{
		// Special tokens (EOF, ILLEGAL, WS)
		{s: ``, tok: chipres.EOF},
		{s: `#`, tok: chipres.ILLEGAL, lit: `#`},
		{s: ` `, tok: chipres.WS, lit: " "},
		{s: "\t", tok: chipres.WS, lit: "\t"},
		{s: "\n", tok: chipres.WS, lit: "\n"},
		{s: "\r", tok: chipres.WS, lit: "\r"},

		// Misc characters
		{s: `*`, tok: chipres.ASTERISK, lit: "*"},
		{s: `:`, tok: chipres.DP, lit: ":"},
		//{s: `(`, tok: chipres.PARDER, lit: "("},
		//{s: `)`, tok: chipres.PARIZQ, lit: ")"},
		//{s: `{`, tok: chipres.LLDER, lit: "{"},
		//{s: `}`, tok: chipres.LLIZQ, lit: "}"},

		// Identifiers
		{s: `foo`, tok: chipres.IDENT, lit: `foo`},
		{s: `Zx12_3U_-`, tok: chipres.IDENT, lit: `Zx12_3U_`},

		// Keywords
		//{s: `FROM`, tok: chipres.FROM, lit: "FROM"},
		//{s: `SELECT`, tok: chipres.SELECT, lit: "SELECT"},
		// Preguntar

		{s: `cicloif`, tok: chipres.cicloif, lit: "cicloif"},
		{s: `celseif`, tok: pasillas.celseif, lit: "celseif"},
		{s: `cicloelse`, tok: pasillas.cicloelse, lit: "cicloelse"},
	}

	for i, tt := range tests {
		s := chipres.NewScanner(strings.NewReader(tt.s))
		tok, lit := s.Scan()
		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)
		} else if tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
		}

	}
}
