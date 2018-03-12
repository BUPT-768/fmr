package bnf

import (
	//"fmt"
	"github.com/liuzl/goutil"
	"testing"
)

var tests = []string{
	`<list>  =  "<" <items> ">"               ;
	<items> =  <items> " " <item> {     nf.math.sum($1,$3)} | <item>   ;
	<item>  =  "f    \\uoo\n" | "bar\t" | "baz"|"好吧"         ;
	`,
}

func TestLex(t *testing.T) {
	for _, c := range tests {
		g, err := CFGrammar(c)
		if err != nil {
			t.Error(err)
		}
		_, err = goutil.JsonMarshalIndent(g, "", "  ")
		if err != nil {
			t.Error(err)
		}
		//fmt.Println(string(b))
	}
}