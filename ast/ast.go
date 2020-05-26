import "github.com/hnts/gorilla/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statement []Statement
}

type LetStatement struct {
	Token token.Token
	Name *Identifer
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifer struct {
	Token token.Token
	Value string
}

func (i *Identifer) expressionNode {}
func (i *Identifer) TokenLiteral() string { return ls.Token.Literal }

func (p *Program) TokenLiteral() string {
	if len(p.Statement) > 0 {
		return p.Statement[0].TokenLiteral()
	}
	return ""
}