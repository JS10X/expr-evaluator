package expr

import (
	"unicode"
	"unicode/utf8"
)

type tokenType int

const (
	lparen tokenType = iota
	rparen
	subtract
	multiply
	divide
	add
	comma
	id
	num
	fnc
)

var opTable = map[rune]*token{
	'(': {typeof: lparen, lexeme: '('},
	')': {typeof: rparen, lexeme: ')'},
	'-': {typeof: subtract, lexeme: '-'},
	'*': {typeof: multiply, lexeme: '*'},
	'/': {typeof: divide, lexeme: '/'},
	'+': {typeof: add, lexeme: '+'},
	',': {typeof: comma, lexeme: ','},
}

func isLeftParen(ch rune) bool {
	return (ch - '(') == 0
}

func isRightParen(ch rune) bool {
	return (ch - ')') == 0
}

func isKeyword(ch rune) bool {
	return (ch - 'P') == 0
}

func isPercentSign(ch rune) bool {
	return (ch - '%') == 0
}

func isPeriod(ch rune) bool {
	return (ch - '.') == 0
}

func isLetter(ch rune) bool {
	return 'a' <= lower(ch) && lower(ch) <= 'z' || ch == '_' || ch >= utf8.RuneSelf && unicode.IsLetter(ch)
}

func isDigit(ch rune) bool {
	return isDecimal(ch) || ch >= utf8.RuneSelf && unicode.IsDigit(ch)
}

func lower(ch rune) rune {
	return ('a' - 'A') | ch // returns lower-case ch if ch is ASCII letter
}

func isDecimal(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func isInvalidChar(ch rune) bool {
	switch {
	case
		(ch - '!') == 0,
		(ch - '"') == 0,
		(ch - '#') == 0,
		(ch - '$') == 0,
		(ch - '&') == 0,
		(ch - '\'') == 0,
		(ch - '[') == 0,
		(ch - ']') == 0,
		(ch - ':') == 0,
		(ch - ';') == 0,
		(ch - '<') == 0,
		(ch - '>') == 0,
		(ch - '=') == 0,
		(ch - '^') == 0,
		(ch - '_') == 0,
		(ch - '`') == 0,
		(ch - '{') == 0,
		(ch - '|') == 0,
		(ch - '}') == 0,
		(ch - '~') == 0,
		int(ch) >= 128:
		return true
	}
	return false
}
