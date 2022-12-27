// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672119205375/AverageSpeed.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AverageSpeed

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AverageSpeedListener is a complete listener for a parse tree produced by AverageSpeedParser.
type AverageSpeedListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAveragespeed is called when entering the averagespeed production.
	EnterAveragespeed(c *AveragespeedContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAveragespeed is called when exiting the averagespeed production.
	ExitAveragespeed(c *AveragespeedContext)
}
