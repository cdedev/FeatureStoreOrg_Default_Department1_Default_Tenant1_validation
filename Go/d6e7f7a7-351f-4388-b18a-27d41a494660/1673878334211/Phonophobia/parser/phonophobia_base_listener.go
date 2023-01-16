// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673878334211/Phonophobia.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Phonophobia

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePhonophobiaListener is a complete listener for a parse tree produced by PhonophobiaParser.
type BasePhonophobiaListener struct{}

var _ PhonophobiaListener = &BasePhonophobiaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePhonophobiaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePhonophobiaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePhonophobiaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePhonophobiaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePhonophobiaListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePhonophobiaListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPhonophobia is called when production phonophobia is entered.
func (s *BasePhonophobiaListener) EnterPhonophobia(ctx *PhonophobiaContext) {}

// ExitPhonophobia is called when production phonophobia is exited.
func (s *BasePhonophobiaListener) ExitPhonophobia(ctx *PhonophobiaContext) {}
