// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669694538700/Arrest.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Arrest

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseArrestListener is a complete listener for a parse tree produced by ArrestParser.
type BaseArrestListener struct{}

var _ ArrestListener = &BaseArrestListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseArrestListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseArrestListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseArrestListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseArrestListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseArrestListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseArrestListener) ExitExpression(ctx *ExpressionContext) {}

// EnterArrest is called when production arrest is entered.
func (s *BaseArrestListener) EnterArrest(ctx *ArrestContext) {}

// ExitArrest is called when production arrest is exited.
func (s *BaseArrestListener) ExitArrest(ctx *ArrestContext) {}
