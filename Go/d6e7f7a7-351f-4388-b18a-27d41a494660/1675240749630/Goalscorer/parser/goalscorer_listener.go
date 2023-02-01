// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675240749630/Goalscorer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Goalscorer

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GoalscorerListener is a complete listener for a parse tree produced by GoalscorerParser.
type GoalscorerListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterGoalscorer is called when entering the goalscorer production.
	EnterGoalscorer(c *GoalscorerContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitGoalscorer is called when exiting the goalscorer production.
	ExitGoalscorer(c *GoalscorerContext)
}
