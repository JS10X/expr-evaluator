## Expression Evaluator

An expression evaluator with support for custom functions wrapped around a CLI.

### Examples

Without a variable

```powershell
go build -o ee.exe
./ee.exe -e "BAND(-(7 + 5) / 2, (7 * 5) / 2)"
```

With a variable via `%P` (if the `-v` flag is omitted, `%P` tokens fallback to a default of 1)

```powershell
go build -o ee.exe
./ee.exe -e "BAND(-(%P + 5) / 2, (%P * 5) / 2)" -v 7
```

### Supported Functions

| Function | Description                                                        |
|----------|--------------------------------------------------------------------|
| NEG      | NEG(X): Returns the negation of X                                  |
| ABS      | ABS(X): Returns the absolute value of X                            |
| ACOS     | ACOS(X): Returns the arc cosine of X radians                       |
| ASIN     | ASIN(X): Returns the arc sine of X radians                         |
| ATAN     | ATAN(X): Returns the arc tangent of X radians                      |
| BAND     | BAND(X,Y): Returns the bitwise AND of X and Y                      |
| BANDNOT  | BANDNOT(X,Y): Returns the bitwise AND NOT of X and Y               |
| BNOT     | BNOT(X): Returns the bitwise NOT of X                              |
| BOR      | BOR(X,Y): Returns the bitwise OR of X and Y                        |
| BXOR     | BXOR(X,Y): Returns the bitwise XOR of X and Y                      |
| CEIL     | CEIL(X): Returns the nearest integer greater than or equal to X    |
| COS      | COS(X): Returns the cosine of X radians                            |
| MOD      | MOD(X,Y): Returns the value of X modulo Y                          |
| POW      | POW(X,Y): Returns the X raised to the power of Y                   |
| RND      | RND(X): Returns the integer nearest to X                           |
| SHL      | SHL(X,Y): Returns the value of X shifted left by Y bits            |
| SHR      | SHR(X,Y): Returns the value of X shifted right by Y bits           |
| SIN      | SIN(X): Returns the sine of X radians                              |
| SQR      | SQR(X): Returns the square root of X                               |
| TAN      | TAN(X): Returns the tangent of X radians                           |
| EQ       | EQ(X,Y): Returns 1 if X is equal to Y, otherwise 0                 |
| NE       | NE(X,Y): Returns 1 if X is not equal to Y, otherwise 0             |
| GE       | GE(X,Y): Returns 1 if X is greater than or equal to Y, otherwise 0 |
| GT       | GT(X,Y): Returns 1 if X is greater than Y, otherwise 0             |
| LE       | LE(X,Y): Returns 1 if X is less than or equal to Y, otherwise 0    |
| LT       | LT(X,Y): Returns 1 if X is less than Y, otherwise 0                |
| MIN      | MIN(X,Y): Returns the minimum of X and Y                           |
| MAX      | MAX(X,Y): Returns the maximum of X and Y                           |
| AND      | AND(X,Y): Returns the logical AND of X and Y                       |
| OR       | OR(X,Y): Returns the logical OR of X and Y                         |
| NOT      | NOT(X): Returns the logical NOT of X                               |

### Grammar: LL(1) One token lookahead

Expression as `E`, Term as `T`, Factor as `F`

- Expression: `E` -> `T` { +|-|, `T`}
- Term:       `T` -> `F` { *|/ `F`}
- Factor:     `F` -> `VAR` | `NUM` | (`E`) | -`F` | `FNC`

#### Definitions:

- `VAR`  ::= char{char}
- `NUM`  ::= digit{digit} | digit.digit
- `FNC`  ::= `FNC`(`ARGS`)
- `ARGS` ::= `E` {, `E`}