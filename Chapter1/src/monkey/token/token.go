package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	// this is how you evaluate if a map contains a key in Go
	// initialize two variables
	// tok will receive either the value of "ident" from the map or an empty string
	// ok will receive a bool that will be true if "ident" is in the map
	// if true (ok is true), then it will execute the body of the if statemtn and tok will be local to that scope
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
