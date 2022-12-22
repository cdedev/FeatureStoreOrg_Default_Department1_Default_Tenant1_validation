// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671690372200/BodilyInjuries.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BodilyInjuries

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BodilyInjuriesListener is a complete listener for a parse tree produced by BodilyInjuriesParser.
type BodilyInjuriesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBodilyinjuries is called when entering the bodilyinjuries production.
	EnterBodilyinjuries(c *BodilyinjuriesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBodilyinjuries is called when exiting the bodilyinjuries production.
	ExitBodilyinjuries(c *BodilyinjuriesContext)
}
