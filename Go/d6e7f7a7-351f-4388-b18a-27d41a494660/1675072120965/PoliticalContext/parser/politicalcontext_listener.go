// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675072120965/PoliticalContext.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // PoliticalContext

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PoliticalContextListener is a complete listener for a parse tree produced by PoliticalContextParser.
type PoliticalContextListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPoliticalcontext is called when entering the politicalcontext production.
	EnterPoliticalcontext(c *PoliticalcontextContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPoliticalcontext is called when exiting the politicalcontext production.
	ExitPoliticalcontext(c *PoliticalcontextContext)
}
