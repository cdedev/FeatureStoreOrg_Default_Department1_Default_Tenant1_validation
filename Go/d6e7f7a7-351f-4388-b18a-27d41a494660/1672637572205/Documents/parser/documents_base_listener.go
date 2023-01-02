// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672637572205/Documents.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Documents

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDocumentsListener is a complete listener for a parse tree produced by DocumentsParser.
type BaseDocumentsListener struct{}

var _ DocumentsListener = &BaseDocumentsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDocumentsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDocumentsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDocumentsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDocumentsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDocumentsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDocumentsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDocuments is called when production documents is entered.
func (s *BaseDocumentsListener) EnterDocuments(ctx *DocumentsContext) {}

// ExitDocuments is called when production documents is exited.
func (s *BaseDocumentsListener) ExitDocuments(ctx *DocumentsContext) {}
