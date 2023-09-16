package expr

type SyntaxError struct {
	treeNode
	message string
}

func (s SyntaxError) Error() string {
	return s.message
}

var (
	UNEXPECTED_TERM_AFTER        = "Unexpected term after %v"
	UNEXPECTED_TERM_AT           = "Unexpected term at %v"
	UNEXPECTED_TERM_CONNECTED_BY = "Unexpected term connected by %v"
	EXPR_CANNOT_START_WITH       = "An expression cannot start with %v"
	EXPR_CANNOT_END_WITH         = "An expression cannot end with %v"
	UNEXPECTED_END_OF_EXPR       = "Unexpected end of expression, expected %v"
	INVALID_CHAR_FOUND_AT        = "Invalid character found %v at index %v"
	CONSECUTIVE_COMMAS           = "An expression cannot contain consecutive commas"
	EXPECTED_FNC_NAME            = "Expected valid function name, but got '%v'. Function names are case sensitive."
	INVALID_FNC_ARG_COUNT        = "Expected %v argument(s) for function '%v', but got %v"
	INVALID_FNC_ARGS_FOR         = "Invalid argument(s) for function '%v'"
	INVALID_FNC_DECL             = "Invalid function declaration"
	INVALID_FNC_DECL_FOR         = "Invalid function declaration for '%v'"
	UNBAL_PARENS                 = "Parenthesis missing in expression"
	DIVIDE_BY_ZERO               = "Cannot divide by zero"
	INVALID_IDENTIFIER           = "Invalid identifier in expression"
	INVALID_NUMBER               = "Invalid number in expression"
	INVALID_EXPR_GENERAL         = "Invalid expression"
	VALID_EXPR                   = "Valid expression"
)
