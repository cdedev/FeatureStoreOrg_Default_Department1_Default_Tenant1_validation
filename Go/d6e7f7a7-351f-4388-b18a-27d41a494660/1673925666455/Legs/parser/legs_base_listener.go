// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673925666455/Legs.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Legs

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLegsListener is a complete listener for a parse tree produced by LegsParser.
type BaseLegsListener struct{}

var _ LegsListener = &BaseLegsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLegsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLegsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLegsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLegsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseLegsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseLegsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLegs is called when production legs is entered.
func (s *BaseLegsListener) EnterLegs(ctx *LegsContext) {}

// ExitLegs is called when production legs is exited.
func (s *BaseLegsListener) ExitLegs(ctx *LegsContext) {}
