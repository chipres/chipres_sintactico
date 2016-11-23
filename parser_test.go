package chipres_test

import (
	"github.com/chipres/chipres_sintactico"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)  

func TestParser_ParseStatement(t *testing.T) {
	f, _ := ioutil.ReadFile("C://proyects//src//github.com//chipres//chipres_sintactico//Gramatica.txt")
	z := string(f[:])


	
	var tests = []struct {
		s    string
		stmt *chipres.Statement
		err  string
	}{
		
		{
			s:    z,
			stmt: &chipres.Statement{
	
			},
		},

		
	}

	for i, tt := range tests {
		stmt, err := chipres.NewParser(strings.NewReader(tt.s)).Parse()
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
			t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt, stmt)
		}
	}
}


func errstring(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
