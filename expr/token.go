package expr

import (
	"fmt"
	"strings"
	"unicode"
)

type token struct {
	typeof tokenType
	lexeme any
}

type scanner struct {
	offset int
	src    []*token
}

func newScanner() *scanner {
	return &scanner{
		offset: -1,
		src:    []*token{},
	}
}

func (s *scanner) peek() *token {
	if (s.offset + 1) >= len(s.src) {
		return nil
	}
	return s.src[s.offset+1]
}

func (s *scanner) next() *token {
	if (s.offset + 1) >= len(s.src) {
		return nil
	}
	s.offset++
	return s.src[s.offset]
}

func (s *scanner) reset() {
	s.offset = -1
	s.src = s.src[:0]
}

// Performs lexical analysis, building the list of tokens from the input string.
func tokenize(input string, sc *scanner, variable any) error {
	var number, functionName string
	var currentToken *token
	var isLastRun, ok bool
	var ch, lookahead rune
	var parens int

	for idx := 0; idx < len(input); idx++ {

		ch = rune(input[idx])
		isLastRun = idx == (len(input) - 1)

		// Skip whitespace
		if unicode.IsSpace(ch) {
			continue
		}

		if isLeftParen(ch) {
			parens++
		}

		if isRightParen(ch) {
			parens--
		}

		if isInvalidChar(ch) {
			return SyntaxError{message: fmt.Sprintf(INVALID_CHAR_FOUND_AT, ch, idx)}
		}

		switch {

		// Variables/Identifiers
		case isPercentSign(ch):
			if isLastRun {
				return SyntaxError{message: fmt.Sprintf(UNEXPECTED_END_OF_EXPR, "variable")}
			}

			lookahead = rune(input[idx+1])
			if !isKeyword(lookahead) {
				return SyntaxError{message: fmt.Sprintf(UNEXPECTED_END_OF_EXPR, "variable")}
			}

			currentToken = &token{typeof: id, lexeme: variable}
			sc.src = append(sc.src, currentToken)
			idx++
			continue

		// Functions
		case isLetter(ch):
			functionName += string(ch)
			if !isLastRun {
				lookahead = rune(input[idx+1])
				if isLetter(lookahead) {
					continue
				}
			}

			if _, ok = funcTable[functionName]; !ok {
				return SyntaxError{message: fmt.Sprintf(EXPECTED_FNC_NAME, functionName)}
			}

			currentToken = &token{typeof: fnc, lexeme: functionName}
			sc.src = append(sc.src, currentToken)
			functionName = ""

		// Numbers
		case isDigit(ch) || isPeriod(ch):
			number += string(ch)
			if !isLastRun {
				lookahead = rune(input[idx+1])
				if isDigit(lookahead) || isPeriod(lookahead) {
					continue
				}
			}

			if isPeriod(rune(number[0])) || strings.Count(number, ".") > 1 {
				return SyntaxError{message: INVALID_NUMBER}
			}

			currentToken = &token{typeof: num, lexeme: number}
			sc.src = append(sc.src, currentToken)
			number = ""
		}

		// Operators
		if currentToken, ok = opTable[ch]; ok {
			sc.src = append(sc.src, currentToken)
		}
	}

	if parens < 0 {
		return SyntaxError{message: UNBAL_PARENS}
	}

	// If we made it here, the expression is valid.
	return nil
}
