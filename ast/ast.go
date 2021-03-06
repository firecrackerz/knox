package ast

import (
	"fmt"
	"knox/token"
	"strings"
)

// NodeType is a string.
type NodeType string

// Node is for the AST.
type Node struct {
	Type       NodeType
	Children   []Node // TODO: Should this be a slice of pointers of Nodes?
	TokenStart token.Token
	Symbols    *SymTable // Only blocks get a symbol table.
}

// Predefined AST node types.
const (
	PROGRAM    = "PROGRAM"    // Variable children. One for each funcdecl.
	CLASS      = "CLASS"      // Variable children. One for each funcdecl and vardecl.
	BLOCK      = "BLOCK"      // Variable children. One for each statement.
	EXPRESSION = "EXPRESSION" // One child. Tree of binary ops, unary ops, and primaries.
	BINARYOP   = "BINARYOP"   // Two children.
	UNARYOP    = "UNARYOP"    // One child.
	INDEXOP    = "INDEXOP"    // Two children.
	// TODO: Consider making this a binaryop.
	DOTOP = "DOTOP" // Two children.
	// TODO: Consider making this a binaryop.
	VARDECL = "VARDECL" // Variable children. Name and type for each variable, then expression.
	// TODO: Consider making the third child a VARASSIGN.
	VARTYPE   = "VARTYPE"   // Variable children. Name and optionally a child for each inner type.
	VARASSIGN = "VARASSIGN" // Variable children children. One or more Varref and one expression.
	FUNCDECL  = "FUNCDECL"  // Four children. Name, paramlist for params, returnlist for return, block.
	PARAMLIST = "PARAMLIST" // Variable children.
	// TODO: Consider making the pairs a VARDECL node.
	RETURNLIST     = "RETURNLIST"     //
	IFSTATEMENT    = "IFSTATEMENT"    // Variable children. Condition and block for IF and zero or more for each ELSE IF and zero or one block for ELSE.
	FORSTATEMENT   = "FORSTATEMENT"   // Three children. Vardecl, expr, and  block.
	WHILESTATEMENT = "WHILESTATEMENT" // Two children. Condition and block.
	JUMPSTATEMENT  = "JUMPSTATEMENT"  // Variable children. Zero except for return. Return has zero or more expressions.
	VARREF         = "VARREF"         // Variable children. Name and list of expressions for array indices. TODO: Update this.
	FUNCCALL       = "FUNCCALL"       // Variable children. Name then one expression for each argument.  TODO: Update this.
	// TODO: Consider making this a binaryop. Name and arglist.
	LEFTEXPR = "LEFTEXPR" // One child. Expr.
	NEW      = "NEW"      // One child. Vartype.
	LIST     = "LIST"     // Variable children. Expressions.
	INT      = "INT"      // Leaf.
	FLOAT    = "FLOAT"    // Leaf.
	STRING   = "STRING"   // Leaf.
	BOOL     = "BOOL"     // Leaf.
	NIL      = "NIL"      // Leaf.
	SELF     = "SELF"     // Leaf.
	VOID     = "VOID"     // Leaf.
	IDENT    = "IDENT"    // Leaf.
)

// Print AST.
func Print(node Node) {
	printUtil(node, 0)
	fmt.Println("")
}

func printUtil(node Node, depth int) {
	var prefix = strings.Repeat(">", depth)
	fmt.Printf("%s %s %s\n", prefix, node.Type, node.TokenStart.Literal)

	// Print symbols
	if node.Type == PROGRAM || node.Type == BLOCK {
		fmt.Printf("Symbols (%d): ", len(node.Symbols.Entries))
		for key := range node.Symbols.Entries {
			fmt.Print(key + " ")
		}
		fmt.Println("")
	}
	for _, c := range node.Children {
		printUtil(c, depth+1)
	}
}
