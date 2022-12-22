// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671690518250/Relationship.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Relationship

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RelationshipListener is a complete listener for a parse tree produced by RelationshipParser.
type RelationshipListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRelationship is called when entering the relationship production.
	EnterRelationship(c *RelationshipContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRelationship is called when exiting the relationship production.
	ExitRelationship(c *RelationshipContext)
}
