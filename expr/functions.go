package expr

import "math"

type fncDescriptor struct {
	args   int
	invoke func(args []treeNode) (any, error)
}

// Functions arguments require validation before invocation.
var funcTable = map[string]*fncDescriptor{

	// NEG(X): Returns the negation of X
	"NEG": {
		args: 1,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) { return -params[0], nil }, args...)
		},
	},

	// ABS(X): Returns the absolute value of X
	"ABS": {
		args: 1,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) { return math.Abs(params[0]), nil }, args...)
		},
	},

	// ACOS(X): Returns the arc cosine of X radians
	"ACOS": {
		args: 1,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) { return math.Acos(params[0]), nil }, args...)
		},
	},

	// ASIN(X): Returns the arc sine of X radians
	"ASIN": {
		args: 1,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) { return math.Asin(params[0]), nil }, args...)
		},
	},

	// ATAN(X): Returns the arc tangent of X radians
	"ATAN": {
		args: 1,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) { return math.Atan(params[0]), nil }, args...)
		},
	},

	// BAND(X,Y): Returns the bitwise AND of X and Y
	"BAND": {
		args: 2,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) { return float64(int(params[0]) & int(params[1])), nil }, args...)
		},
	},

	// BANDNOT(X,Y): Returns the bitwise AND NOT of X and Y
	"BANDNOT": {
		args: 2,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) { return float64(int(params[0]) &^ int(params[1])), nil }, args...)
		},
	},

	// BNOT(X): Returns the bitwise NOT of X
	"BNOT": {
		args: 1,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) { return float64(^int(params[0])), nil }, args...)
		},
	},

	// BOR(X,Y): Returns the bitwise OR of X and Y
	"BOR": {
		args: 2,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) { return float64(int(params[0]) | int(params[1])), nil }, args...)
		},
	},

	// BXOR(X,Y): Returns the bitwise XOR of X and Y
	"BXOR": {
		args: 2,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) { return float64(int(params[0]) ^ int(params[1])), nil }, args...)
		},
	},

	// CEIL(X): Returns the nearest integer greater than or equal to X
	"CEIL": {
		args: 1,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) { return math.Ceil(params[0]), nil }, args...)
		},
	},

	// COS(X): Returns the cosine of X radians
	"COS": {
		args: 1,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) { return math.Cos(params[0]), nil }, args...)
		},
	},

	// MOD(X,Y): Returns the value of X modulo Y
	"MOD": {
		args: 2,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) { return math.Mod(params[0], params[1]), nil }, args...)
		},
	},

	// POW(X,Y): Returns the X raised to the power of Y
	"POW": {
		args: 2,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) { return math.Pow(params[0], params[1]), nil }, args...)
		},
	},

	// RND(X): Returns the integer nearest to X
	"RND": {
		args: 1,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) { return math.RoundToEven(params[0]), nil }, args...)
		},
	},

	// SHL(X,Y): Returns the value of X shifted left by Y bits
	"SHL": {
		args: 2,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) { return float64(int(params[0]) << int(params[1])), nil }, args...)
		},
	},

	// SHR(X,Y): Returns the value of X shifted right by Y bits
	"SHR": {
		args: 2,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) { return float64(int(params[0]) >> int(params[1])), nil }, args...)
		},
	},

	// SIN(X): Returns the sine of X radians
	"SIN": {
		args: 1,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) { return math.Sin(params[0]), nil }, args...)
		},
	},

	// SQR(X): Returns the square root of X
	"SQR": {
		args: 1,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) { return math.Sqrt(params[0]), nil }, args...)
		},
	},

	// TAN(X): Returns the tangent of X radians
	"TAN": {
		args: 1,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) { return math.Tan(params[0]), nil }, args...)
		},
	},

	// EQ(X,Y): Returns 1 if X is equal to Y, otherwise 0
	"EQ": {
		args: 2,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) {
				if params[0] == params[1] {
					return 1.0, nil
				}
				return 0.0, nil
			}, args...)
		},
	},

	// NE(X,Y): Returns 1 if X is not equal to Y, otherwise 0
	"NE": {
		args: 2,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) {
				if params[0] != params[1] {
					return 1.0, nil
				}
				return 0.0, nil
			}, args...)
		},
	},

	// GE(X,Y): Returns 1 if X is greater than or equal to Y, otherwise 0
	"GE": {
		args: 2,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) {
				if params[0] >= params[1] {
					return 1.0, nil
				}
				return 0.0, nil
			}, args...)
		},
	},

	// GT(X,Y): Returns 1 if X is greater than Y, otherwise 0
	"GT": {
		args: 2,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) {
				if params[0] > params[1] {
					return 1.0, nil
				}
				return 0.0, nil
			}, args...)
		},
	},

	// LE(X,Y): Returns 1 if X is less than or equal to Y, otherwise 0
	"LE": {
		args: 2,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) {
				if params[0] <= params[1] {
					return 1.0, nil
				}
				return 0.0, nil
			}, args...)
		},
	},

	// LT(X,Y): Returns 1 if X is less than Y, otherwise 0
	"LT": {
		args: 2,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) {
				if params[0] < params[1] {
					return 1.0, nil
				}
				return 0.0, nil
			}, args...)
		},
	},

	// MIN(X,Y): Returns the minimum of X and Y
	"MIN": {
		args: 2,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) { return math.Min(params[0], params[1]), nil }, args...)
		},
	},

	// MAX(X,Y): Returns the maximum of X and Y
	"MAX": {
		args: 2,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) { return math.Max(params[0], params[1]), nil }, args...)
		},
	},

	// AND(X,Y): Returns the logical AND of X and Y
	"AND": {
		args: 2,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) {
				if params[0] == 1 && params[1] == 1 {
					return 1.0, nil
				}
				return 0.0, nil
			}, args...)
		},
	},

	// OR(X,Y): Returns the logical OR of X and Y
	"OR": {
		args: 2,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) {
				if params[0] == 1 || params[1] == 1 {
					return 1.0, nil
				}
				return 0.0, nil
			}, args...)
		},
	},

	// NOT(X): Returns the logical NOT of X
	"NOT": {
		args: 1,
		invoke: func(args []treeNode) (any, error) {
			return evalT(func(params ...float64) (any, error) {
				if params[0] == 1.0 {
					return 0.0, nil
				}
				return 1.0, nil
			}, args...)
		},
	},
}
