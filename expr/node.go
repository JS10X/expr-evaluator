package expr

import (
	"fmt"
	"strconv"
)

type treeNode interface {
	Print()
	Eval() (any, error)
}

type addition struct{ left, right treeNode }
type subtraction struct{ left, right treeNode }
type multiplication struct{ left, right treeNode }
type division struct{ left, right treeNode }
type negation struct{ arg treeNode }
type identifer struct{ value any }
type number struct{ value any }

type functionArgs struct {
	owner string
	args  []treeNode
}

type function struct {
	name string
	args []treeNode
}

func newAdd(left, right treeNode) *addition            { return &addition{left, right} }
func newSubtract(left, right treeNode) *subtraction    { return &subtraction{left, right} }
func newMultiply(left, right treeNode) *multiplication { return &multiplication{left, right} }
func newDivide(left, right treeNode) *division         { return &division{left, right} }
func newNegate(arg treeNode) *negation                 { return &negation{arg} }
func newIdentifer(t *token) *identifer                 { return &identifer{t.lexeme} }
func newNumber(t *token) *number                       { return &number{t.lexeme} }
func newFunctionArgs(args []treeNode) *functionArgs    { return &functionArgs{args: args} }

func newFunction(fnc string, args []treeNode) *function {
	return &function{fnc, args}
}

func (o *addition) Eval() (any, error) {
	return evalT(func(params ...float64) (any, error) { return params[0] + params[1], nil }, o.left, o.right)
}

func (o *subtraction) Eval() (any, error) {
	return evalT(func(params ...float64) (any, error) { return params[0] - params[1], nil }, o.left, o.right)
}

func (o *multiplication) Eval() (any, error) {
	return evalT(func(params ...float64) (any, error) { return params[0] * params[1], nil }, o.left, o.right)
}

func (o *division) Eval() (any, error) {
	return evalT(func(params ...float64) (any, error) {
		if params[1] == 0 {
			return nil, SyntaxError{message: DIVIDE_BY_ZERO}
		}
		return params[0] / params[1], nil
	}, o.left, o.right)
}

func (o *negation) Eval() (any, error) {
	return evalT(func(params ...float64) (any, error) { return -params[0], nil }, o.arg)
}

func (o *identifer) Eval() (any, error) {
	return evalN(o.value)
}

func (o *number) Eval() (any, error) {
	return evalN(o.value)
}

func (o *function) Eval() (any, error) {
	fn, ok := funcTable[o.name]
	if !ok {
		return nil, SyntaxError{message: fmt.Sprintf(EXPECTED_FNC_NAME, o.name)}
	}

	if len(o.args) != fn.args {
		return nil, SyntaxError{message: fmt.Sprintf(INVALID_FNC_ARG_COUNT, fn.args, len(o.args))}
	}
	return fn.invoke(o.args)
}

func (o *functionArgs) Eval() (any, error) {
	fn, ok := funcTable[o.owner]
	if !ok {
		return nil, SyntaxError{message: fmt.Sprintf(EXPECTED_FNC_NAME, o.owner)}
	}

	if len(o.args) != fn.args {
		return nil, SyntaxError{message: fmt.Sprintf(INVALID_FNC_ARG_COUNT, fn.args, len(o.args))}
	}
	return fn.invoke(o.args)
}

func (o *addition) Print() {
	fmt.Printf("(")
	o.left.Print()
	fmt.Printf("+")
	o.right.Print()
	fmt.Printf(")")
}

func (o *subtraction) Print() {
	fmt.Printf("(")
	o.left.Print()
	fmt.Printf("-")
	o.right.Print()
	fmt.Printf(")")
}

func (o *multiplication) Print() {
	fmt.Printf("(")
	o.left.Print()
	fmt.Printf("*")
	o.right.Print()
	fmt.Printf(")")
}

func (o *division) Print() {
	fmt.Printf("(")
	o.left.Print()
	fmt.Printf("/")
	o.right.Print()
	fmt.Printf(")")
}

func (o *negation) Print() {
	fmt.Printf("(")
	fmt.Printf("-")
	o.arg.Print()
}

func (o *identifer) Print() {
	fmt.Printf("%v", o.value)
}

func (o *number) Print() {
	fmt.Printf("%v", o.value)
}

func (o *functionArgs) Print() {
	for ix, arg := range o.args {
		arg.Print()
		if ix != (len(o.args) - 1) {
			fmt.Printf(",")
		}
	}
}

func (o *function) Print() {
	fmt.Printf(o.name)
	fmt.Printf("(")
	for ix, arg := range o.args {
		arg.Print()
		if ix != (len(o.args) - 1) {
			fmt.Printf(",")
		}
	}
	fmt.Printf(")")
}

func evalT(fn func(params ...float64) (any, error), nodes ...treeNode) (any, error) {

	var ct any
	var cv float64
	var err error
	var args []float64

	for _, curr := range nodes {

		if err, ok := curr.(SyntaxError); ok {
			return nil, err
		}

		ct, err = curr.Eval()
		if err != nil {
			return nil, err
		}

		cv, err = evalN(ct)
		if err != nil {
			return nil, err
		}
		args = append(args, cv)
	}
	return fn(args...)
}

func evalN(value any) (float64, error) {

	switch v := value.(type) {
	case string:
		num, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, err
		}
		return num, nil

	case int:
		return float64(v), nil

	case float64:
		return v, nil
	}
	return 0, SyntaxError{message: fmt.Sprintf(UNEXPECTED_TERM_AT, value)}
}
