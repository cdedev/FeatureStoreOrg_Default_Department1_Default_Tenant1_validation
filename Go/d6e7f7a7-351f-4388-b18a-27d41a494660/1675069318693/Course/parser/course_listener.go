// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675069318693/Course.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Course

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CourseListener is a complete listener for a parse tree produced by CourseParser.
type CourseListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCourse is called when entering the course production.
	EnterCourse(c *CourseContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCourse is called when exiting the course production.
	ExitCourse(c *CourseContext)
}
