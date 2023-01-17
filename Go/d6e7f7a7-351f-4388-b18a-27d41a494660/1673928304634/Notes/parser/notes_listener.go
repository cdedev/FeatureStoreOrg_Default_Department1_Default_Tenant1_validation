// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673928304634/Notes.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Notes

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NotesListener is a complete listener for a parse tree produced by NotesParser.
type NotesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterNotes is called when entering the notes production.
	EnterNotes(c *NotesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitNotes is called when exiting the notes production.
	ExitNotes(c *NotesContext)
}
