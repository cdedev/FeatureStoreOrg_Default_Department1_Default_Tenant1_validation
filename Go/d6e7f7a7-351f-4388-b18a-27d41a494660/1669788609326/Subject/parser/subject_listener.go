// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669788609326/Subject.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Subject

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SubjectListener is a complete listener for a parse tree produced by SubjectParser.
type SubjectListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSubject is called when entering the subject production.
	EnterSubject(c *SubjectContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSubject is called when exiting the subject production.
	ExitSubject(c *SubjectContext)
}
