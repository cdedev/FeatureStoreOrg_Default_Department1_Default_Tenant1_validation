// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674621397885/Calories.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Calories

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CaloriesListener is a complete listener for a parse tree produced by CaloriesParser.
type CaloriesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCalories is called when entering the calories production.
	EnterCalories(c *CaloriesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCalories is called when exiting the calories production.
	ExitCalories(c *CaloriesContext)
}
