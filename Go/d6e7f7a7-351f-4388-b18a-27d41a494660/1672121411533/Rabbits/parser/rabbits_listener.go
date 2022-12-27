// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672121411533/Rabbits.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rabbits

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RabbitsListener is a complete listener for a parse tree produced by RabbitsParser.
type RabbitsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRabbits is called when entering the rabbits production.
	EnterRabbits(c *RabbitsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRabbits is called when exiting the rabbits production.
	ExitRabbits(c *RabbitsContext)
}
