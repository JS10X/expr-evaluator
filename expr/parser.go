package expr

import "fmt"

type Parser struct {
	scn *scanner
}

func NewParser() *Parser {
	return &Parser{scn: newScanner()}
}

func (p *Parser) EvalV(input string, variable any) (float64, error) {

	defer p.scn.reset()

	err := tokenize(input, p.scn, variable)
	if err != nil {
		return 0, err
	}

	res, err := parse(p.scn)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (p *Parser) Eval(input string) (float64, error) {

	defer p.scn.reset()

	err := tokenize(input, p.scn, nil)
	if err != nil {
		return 0, err
	}

	res, err := parse(p.scn)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func parse(sc *scanner) (float64, error) {

	ast := parseE(sc)
	if ast == nil {
		return 0, SyntaxError{message: INVALID_EXPR_GENERAL}
	}

	if err, ok := ast.(SyntaxError); ok {
		return 0, err
	}

	evaluated, err := ast.Eval()
	if err != nil {
		return 0, err
	}

	res, ok := evaluated.(float64)
	if !ok {
		return 0, SyntaxError{message: INVALID_EXPR_GENERAL}
	}
	return res, nil
}

// Expression: E -> T { +|-|, T}
func parseE(sc *scanner) treeNode {

	var lookahead *token
	var nA, nB treeNode
	var args []treeNode

	nA = parseT(sc)
	if err, ok := nA.(SyntaxError); ok {
		return err
	}

	for {
		if sc.peek() == nil {
			return nA
		}

		switch sc.peek().typeof {
		case add:
			sc.next() // scan past '+'
			nB = parseT(sc)
			if nA == nil || nB == nil {
				return SyntaxError{message: fmt.Sprintf(UNEXPECTED_TERM_CONNECTED_BY, '+')}
			}

			nA = newAdd(nA, nB)
			if len(args) > 0 {
				lookahead = sc.peek()
				if lookahead == nil || lookahead.typeof == rparen {
					args = append(args, nA)
					nA = newFunctionArgs(args)
				}
			}

		case subtract:
			sc.next() // scan past '-'
			nB = parseT(sc)
			if nA == nil || nB == nil {
				return SyntaxError{message: fmt.Sprintf(UNEXPECTED_TERM_CONNECTED_BY, '-')}
			}

			nA = newSubtract(nA, nB)
			if len(args) > 0 {
				lookahead = sc.peek()
				if lookahead == nil || lookahead.typeof == rparen {
					args = append(args, nA)
					nA = newFunctionArgs(args)
				}
			}

		case comma:
			sc.next() // scan past ','
			if nA == nil {
				return SyntaxError{message: fmt.Sprintf(UNEXPECTED_TERM_CONNECTED_BY, ',')}
			}

			args = append(args, nA)
			nA = parseT(sc)
			if nA == nil {
				return SyntaxError{message: fmt.Sprintf(UNEXPECTED_TERM_CONNECTED_BY, ',')}
			}

			lookahead = sc.peek()
			if lookahead == nil {
				return SyntaxError{message: fmt.Sprintf(UNEXPECTED_TERM_CONNECTED_BY, ',')}
			}

			if lookahead.typeof == rparen {
				args = append(args, nA)
				nA = newFunctionArgs(args)
			}

		default:
			return nA
		}
	}
}

// Term: T -> F { *|/ F}
func parseT(sc *scanner) treeNode {

	var nA, nB treeNode
	nA = parseF(sc)
	if err, ok := nA.(SyntaxError); ok {
		return err
	}

	for {
		if sc.peek() == nil {
			return nA
		}

		switch sc.peek().typeof {
		case multiply:
			sc.next() // scan past '*'
			nB = parseF(sc)
			if nA == nil || nB == nil {
				return SyntaxError{message: fmt.Sprintf(UNEXPECTED_TERM_CONNECTED_BY, '*')}
			}
			nA = newMultiply(nA, nB)

		case divide:
			sc.next() // scan past '/'
			nB = parseF(sc)
			if nA == nil || nB == nil {
				return SyntaxError{message: fmt.Sprintf(UNEXPECTED_TERM_CONNECTED_BY, '/')}
			}
			nA = newDivide(nA, nB)

		default:
			return nA
		}
	}
}

// Factor: F -> VAR | NUM | (E) | -F | FNC
func parseF(sc *scanner) treeNode {

	var next, lookahead *token
	var nA treeNode
	var ok bool
	var fn string

	for {
		if sc.peek() == nil {
			return nA
		}

		switch sc.peek().typeof {

		case id:
			nA = newIdentifer(sc.next())
			return nA

		case num:
			nA = newNumber(sc.next())
			return nA

		case lparen:
			sc.next() // scan past the '('
			nA = parseE(sc)
			if nA == nil {
				return SyntaxError{message: fmt.Sprintf(UNEXPECTED_TERM_AFTER, '(')}
			}

			lookahead = sc.peek()
			if lookahead == nil {
				return SyntaxError{message: fmt.Sprintf(UNEXPECTED_TERM_AFTER, '(')}
			}

			if lookahead.typeof == rparen {
				sc.next() // scan past the ')'
				return nA
			} else {
				return SyntaxError{message: fmt.Sprintf(UNEXPECTED_END_OF_EXPR, ')')}
			}

		case subtract:
			sc.next() // scan past the '-'
			nA = parseF(sc)
			if nA == nil {
				return SyntaxError{message: fmt.Sprintf(UNEXPECTED_TERM_AFTER, '-')}
			}
			nA = newNegate(nA)
			return nA

		case fnc:
			next = sc.next() // scan past the 'function name'
			if fn, ok = next.lexeme.(string); !ok {
				return SyntaxError{message: fmt.Sprintf(UNEXPECTED_END_OF_EXPR, "function")}
			}

			next = sc.next() // scan past the '('
			if next == nil || next.typeof != lparen {
				return SyntaxError{message: INVALID_FNC_DECL}
			}

			nA = parseE(sc)

			switch node := nA.(type) {

			case *function:
				nA = newFunction(fn, []treeNode{node})

			case *functionArgs:
				node.owner = fn
				nA = newFunction(fn, node.args)

			case *number:
				nA = newFunction(fn, []treeNode{node})

			case *identifer:
				nA = newFunction(fn, []treeNode{node})

			case *addition:
				nA = newFunction(fn, []treeNode{node})

			case *subtraction:
				nA = newFunction(fn, []treeNode{node})

			case *multiplication:
				nA = newFunction(fn, []treeNode{node})

			case *division:
				nA = newFunction(fn, []treeNode{node})

			default:
				return SyntaxError{message: fmt.Sprintf(INVALID_FNC_ARGS_FOR, fn)}
			}

			next = sc.next()
			if next == nil || next.typeof != rparen {
				return SyntaxError{message: fmt.Sprintf(INVALID_FNC_DECL_FOR, fn)}
			}

		default:
			return nA
		}
	}
}
