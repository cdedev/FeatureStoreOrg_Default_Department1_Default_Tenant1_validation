// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673855347164/Debitscore.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Debitscore

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDebitscoreListener is a complete listener for a parse tree produced by DebitscoreParser.
type BaseDebitscoreListener struct{}

var _ DebitscoreListener = &BaseDebitscoreListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDebitscoreListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDebitscoreListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDebitscoreListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDebitscoreListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDebitscoreListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDebitscoreListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDebitscore is called when production debitscore is entered.
func (s *BaseDebitscoreListener) EnterDebitscore(ctx *DebitscoreContext) {}

// ExitDebitscore is called when production debitscore is exited.
func (s *BaseDebitscoreListener) ExitDebitscore(ctx *DebitscoreContext) {}
