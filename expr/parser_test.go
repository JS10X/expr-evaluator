package expr_test

import (
	"testing"

	"github.com/js10x/expr-evaluator/expr"
)

const (
	expected_but_got_for_expr = "expected %v, but got %v for expression %v"
)

func TestEval(t *testing.T) {

	tests := []struct {
		input  string
		expect float64
	}{
		{input: "2 + 3 * 4", expect: 14},
		{input: "2 + 3 / 4", expect: 2.75},
		{input: "3 / 3", expect: 1},
		{input: "-(7 + 5) * 2", expect: -24},
		{input: "ABS(-(7 + 5) - 2)", expect: 14},
		{input: "ABS(-(7 + 5) + 2)", expect: 10},
		{input: "ABS(-(7 + 5) * 2)", expect: 24},
		{input: "ABS(-(7 + 5) / 2)", expect: 6},
		{input: "ACOS(0.125)", expect: 1.445468495626831},
		{input: "ASIN(0.125)", expect: 0.12532783116806537},
		{input: "ATAN(0.125)", expect: 0.12435499454676144},
		{input: "BAND(255,4)", expect: 4},
		{input: "BANDNOT(255,8)", expect: 247},
		{input: "BNOT(255)", expect: -256},
		{input: "BOR(15,2)", expect: 15},
		{input: "BXOR(15,2)", expect: 13},
		{input: "CEIL(15.5)", expect: 16},
		{input: "COS(360)", expect: -0.2836910914865273},
		{input: "MOD(4,2)", expect: 0},
		{input: "POW(2,2)", expect: 4},
		{input: "RND(3.4)", expect: 3},
		{input: "SHL(255,8)", expect: 65280},
		{input: "SHR(255,8)", expect: 0},
		{input: "SIN(360)", expect: 0.9589157234143065},
		{input: "SQR(5)", expect: 2.23606797749979},
		{input: "TAN(360)", expect: -3.380140413960958},
		{input: "EQ(1,8)", expect: 0},
		{input: "EQ(8,8)", expect: 1},
		{input: "NE(1,8)", expect: 1},
		{input: "NE(8,8)", expect: 0},
		{input: "GE(1,8)", expect: 0},
		{input: "GE(8,8)", expect: 1},
		{input: "GT(4,8)", expect: 0},
		{input: "GT(8,4)", expect: 1},
		{input: "LE(8,7)", expect: 0},
		{input: "LE(8,8)", expect: 1},
		{input: "LT(8,4)", expect: 0},
		{input: "LT(4,8)", expect: 1},
		{input: "MIN(8,4)", expect: 4},
		{input: "MIN(4,8)", expect: 4},
		{input: "MAX(8,4)", expect: 8},
		{input: "MAX(4,8)", expect: 8},
		{input: "AND(1,0)", expect: 0},
		{input: "AND(1,1)", expect: 1},
		{input: "OR(1,0)", expect: 1},
		{input: "OR(0,0)", expect: 0},
		{input: "NOT(0)", expect: 1},
		{input: "NOT(1)", expect: 0},
		{input: "BAND(5,-(CEIL(CEIL(CEIL(1.5)))))", expect: 4},
		{input: "BAND(5,-(CEIL(SQR(SHL(5, 2))))) * -(5 + 1) * 1", expect: -6},
		{input: "BAND(-(7 + 5) / 2, (7 * 5) / 2)", expect: 16},
		{input: "BAND(-(7+5)/2, BANDNOT(-(7*5)/2,5))", expect: -22},
		{input: "SHL(BAND(CEIL(BAND(CEIL(249.50), 15)), BAND(CEIL(219.50), 10)), 4)", expect: 128},
	}

	var res float64
	var err error
	parser := expr.NewParser()

	for _, tc := range tests {

		res, err = parser.Eval(tc.input)
		if err != nil {
			t.Fatalf(expected_but_got_for_expr, nil, err.Error(), tc.input)
		}

		if tc.expect != res {
			t.Fatalf(expected_but_got_for_expr, tc.expect, res, tc.input)
		}
	}
}

func TestEvalV(t *testing.T) {

	tests := []struct {
		input    string
		variable float64
		expect   float64
	}{
		{input: "2 + %P * 4", variable: 3, expect: 14},
		{input: "%P + %P / %P", variable: 4, expect: 5},
		{input: "%P / %P", variable: 4, expect: 1},
		{input: "-(7 + %P) * 2", variable: 5, expect: -24},
		{input: "ABS(-(7 + %P) - 2)", variable: 5, expect: 14},
		{input: "ABS(-(7 + %P) + 2)", variable: 5, expect: 10},
		{input: "ABS(-(7 + %P) * 2)", variable: 5, expect: 24},
		{input: "ABS(-(7 + %P) / 2)", variable: 5, expect: 6},
		{input: "ACOS(%P)", variable: 0.125, expect: 1.445468495626831},
		{input: "ASIN(%P)", variable: 0.125, expect: 0.12532783116806537},
		{input: "ATAN(%P)", variable: 0.125, expect: 0.12435499454676144},
		{input: "BAND(%P,4)", variable: 255, expect: 4},
		{input: "BANDNOT(%P,8)", variable: 255, expect: 247},
		{input: "BNOT(%P)", variable: 255, expect: -256},
		{input: "BOR(15,%P)", variable: 2, expect: 15},
		{input: "BXOR(15,%P)", variable: 2, expect: 13},
		{input: "CEIL(%P)", variable: 15.5, expect: 16},
		{input: "COS(%P)", variable: 360, expect: -0.2836910914865273},
		{input: "MOD(4,%P)", variable: 2, expect: 0},
		{input: "POW(2,%P)", variable: 2, expect: 4},
		{input: "RND(%P)", variable: 3.4, expect: 3},
		{input: "SHL(%P,8)", variable: 255, expect: 65280},
		{input: "SHR(%P,8)", variable: 255, expect: 0},
		{input: "SIN(%P)", variable: 360, expect: 0.9589157234143065},
		{input: "SQR(%P)", variable: 5, expect: 2.23606797749979},
		{input: "TAN(%P)", variable: 360, expect: -3.380140413960958},
		{input: "EQ(1,%P)", variable: 8, expect: 0},
		{input: "EQ(%P,%P)", variable: 8, expect: 1},
		{input: "NE(1,%P)", variable: 8, expect: 1},
		{input: "NE(%P,%P)", variable: 8, expect: 0},
		{input: "GE(1,%P)", variable: 8, expect: 0},
		{input: "GE(%P,%P)", variable: 8, expect: 1},
		{input: "GT(8,%P)", variable: 4, expect: 1},
		{input: "GT(%P-1,%P)", variable: 4, expect: 0},
		{input: "LE(%P,1)", variable: 8, expect: 0},
		{input: "LE(%P,%P)", variable: 8, expect: 1},
		{input: "LT(%P,8)", variable: 4, expect: 1},
		{input: "LT(%P+1,%P)", variable: 4, expect: 0},
		{input: "MIN(8,%P)", variable: 4, expect: 4},
		{input: "MIN(%P-1,%P)", variable: 4, expect: 3},
		{input: "MAX(8,%P)", variable: 4, expect: 8},
		{input: "MAX(%P,%P*3)", variable: 4, expect: 12},
		{input: "AND(%P,%P-1)", variable: 2, expect: 0},
		{input: "AND(%P-1, %P-1)", variable: 2, expect: 1},
		{input: "OR(%P/1,%P)", variable: 1, expect: 1},
		{input: "OR(%P-1,%P-1)", variable: 1, expect: 0},
		{input: "NOT(%P)", variable: 1, expect: 0},
		{input: "NOT(%P)", variable: 0, expect: 1},
		{input: "BAND(5,-(CEIL(CEIL(CEIL(%P)))))", variable: 1.5, expect: 4},
		{input: "BAND(5,-(CEIL(SQR(SHL(%P, 2))))) * -(%P + 1) * 1", variable: 5, expect: -6},
		{input: "BAND(-(%P + 5) / 2, (%P * 5) / 2)", variable: 7, expect: 16},
		{input: "BAND(-(7+%P)/2, BANDNOT(-(7*%P)/2,%P))", variable: 5, expect: -22},
		{input: "SHL(BAND(CEIL(BAND(CEIL(249.50), 15)), BAND(CEIL(219.50), %P)), 4)", variable: 10, expect: 128},
	}

	var res float64
	var err error
	parser := expr.NewParser()

	for _, tc := range tests {

		res, err = parser.EvalV(tc.input, tc.variable)
		if err != nil {
			t.Errorf(expected_but_got_for_expr, nil, err.Error(), tc.input)
		}

		if tc.expect != res {
			t.Errorf(expected_but_got_for_expr, tc.expect, res, tc.input)
		}
	}
}

func TestEvalSyntaxErrors(t *testing.T) {

	tests := []struct {
		input string
	}{
		{input: "ABS(0ABS(0))"},
		{input: "ABS(0,0)"},
		{input: "Abs(0)"},
		{input: "(2/(2*0))*20"},
		{input: ",,"},
		{input: "(,,)"},
		{input: "(),,"},
		{input: ","},
		{input: "()"},
		{input: "((0)"},
		{input: "(,"},
		{input: ").(x"},
		{input: "ABS(0) * 2,,0"},
		{input: "SHL(,)"},
		{input: "SHL(0,,0)"},
		{input: "AND(-+)"},
		{input: "AND(m,n)"},
		{input: "A(0)"},
		{input: "(4)*/28"},
		{input: "<(0)>"},
		{input: "SHL((2+2/))"},
		{input: "(ABS(2)) * (0))"},
	}

	parser := expr.NewParser()
	var err error

	for _, tc := range tests {

		_, err = parser.Eval(tc.input)
		if _, ok := err.(expr.SyntaxError); !ok {
			// Not convertible to syntax error
			t.Errorf(expected_but_got_for_expr, "registered syntax error", err, tc.input)
		}
	}
}

func BenchmarkEvaluate(b *testing.B) {

	parser := expr.NewParser()
	for i := 0; i < b.N; i++ {
		parser.EvalV("BAND(5, 10 * 1000.0)", 10)
	}
}

func FuzzEval(f *testing.F) {

	testcases := []string{
		"BAND(-(7+5)/2, BANDNOT(-(7*5)/2,5))",
		"1985 * 15 / 50.0",
		"BAND(5,CEIL(CEIL(CEIL(1.5))))",
		"CEIL(2)",
		"NOT(1130.0) + ()",
		"(,",
		").(x",
		"AND(-+)",
		"'9(AND(0, 0)    * X * X / 0",
		"0-0-0!",
		"((((059.))))",
		"!",
		"OR(0,0) * -(2 * OR(1,1) / AND(0,1) * ((((0))))",
	}
	for _, tc := range testcases {
		f.Add(tc)
	}

	parser := expr.NewParser()
	f.Fuzz(func(t *testing.T, input string) {
		parser.Eval(input)
	})
}
