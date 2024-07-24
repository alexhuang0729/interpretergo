package token

type TokenType string

type Token struct {
	Type    TokenType //Allows us to use many different values as TokenTypes, distinguish between different types of tokens
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // Signifies a token/character we don't know about
	EOF     = "EOF"     // Stands for "end of file" and tells our parser later on that it can

	// Identifiers and literals
	IDENT = "IDENT" // add, foobar, x, y
	INT   = "INT"   // 1, 2, 3, 4, 5, 6, etc

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
