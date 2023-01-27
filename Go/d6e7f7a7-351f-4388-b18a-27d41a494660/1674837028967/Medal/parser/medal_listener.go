// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674837028967/Medal.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Medal

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MedalListener is a complete listener for a parse tree produced by MedalParser.
type MedalListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMedal is called when entering the medal production.
	EnterMedal(c *MedalContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMedal is called when exiting the medal production.
	ExitMedal(c *MedalContext)
}
