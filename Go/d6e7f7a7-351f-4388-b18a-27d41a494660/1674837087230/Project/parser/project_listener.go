// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674837087230/Project.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Project

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ProjectListener is a complete listener for a parse tree produced by ProjectParser.
type ProjectListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterProject is called when entering the project production.
	EnterProject(c *ProjectContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitProject is called when exiting the project production.
	ExitProject(c *ProjectContext)
}
