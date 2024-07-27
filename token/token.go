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
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

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
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"

	// Operators
	EQ     = "=="
	NOT_EQ = "!="
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

/*Checks the keywords table to see whether the given identifier is in fact a keyword.
If so, returns keyword's TokenType constant
If not, returns token.IDENT (TokenType for all user-defined identifiers)
*/
