// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669659215020/Grade.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Grade

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GradeListener is a complete listener for a parse tree produced by GradeParser.
type GradeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterGrade is called when entering the grade production.
	EnterGrade(c *GradeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitGrade is called when exiting the grade production.
	ExitGrade(c *GradeContext)
}
