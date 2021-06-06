package token

const (
	Illegal   = "ILLEGAL"
	EOF       = "EOF"
	Ident     = "IDENT"
	Int       = "INT"
	Assign    = "="
	Plus      = "+"
	Minus     = "-"
	Bang      = "!"
	Asterisk  = "*"
	Slash     = "/"
	Lt        = "<"
	Gt        = ">"
	Comma     = ","
	Semicolon = ";"
	LParen    = "("
	RParen    = ")"
	LBrace    = "{"
	RBrace    = "}"
	LBracket  = "["
	RBracket  = "]"
	Function  = "FUNCTION"
	Let       = "LET"
	True      = "TRUE"
	False     = "FALSE"
	If        = "IF"
	Else      = "ELSE"
	Return    = "RETURN"
	Eq        = "=="
	NotEq     = "!="
	String    = "STRING"
	Colon     = ":"
)

type Type string

type Token struct {
	Type    Type
	Literal string
}

var keywords = map[string]Type{
	"fn":     Function,
	"let":    Let,
	"true":   True,
	"false":  False,
	"if":     If,
	"else":   Else,
	"return": Return,
}

func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return Ident
}
