// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673928304634/Notes.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Notes

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNotesListener is a complete listener for a parse tree produced by NotesParser.
type BaseNotesListener struct{}

var _ NotesListener = &BaseNotesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNotesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNotesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNotesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNotesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseNotesListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseNotesListener) ExitExpression(ctx *ExpressionContext) {}

// EnterNotes is called when production notes is entered.
func (s *BaseNotesListener) EnterNotes(ctx *NotesContext) {}

// ExitNotes is called when production notes is exited.
func (s *BaseNotesListener) ExitNotes(ctx *NotesContext) {}
