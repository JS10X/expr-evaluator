package main

import (
	"flag"
	"log"
	"strings"

	"github.com/js10x/expr-evaluator/expr"
)

var parser *expr.Parser = expr.NewParser()

func main() {
	var expression string
	var variable float64

	// The expression to evaluate.
	flag.StringVar(&expression, "e", "", "-(7 + 5) * 2")

	// A numeric value that will be inserted into the expression during evaluation anywhere a %P identifier is defined.
	flag.Float64Var(&variable, "v", 1, "-e \"-(%P + 5) * 2 + 2\" -v 7.125")
	flag.Parse()
	if len(strings.TrimSpace(expression)) <= 0 {
		log.Fatalln("An expression must be provided with the 'e' flag")
	}

	evaluated, err := parser.EvalV(expression, variable)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Evaluated -> %v\n", evaluated)
}
