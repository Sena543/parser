package src

type Node interface {
	TokenLexeme() string
}

type Key interface {
	Node
	keyNode()
}

type Value interface {
	Node
	valueNode()
}

/* type Node struct {
	operator Token
	left     *Node
	right    *Node
} */
