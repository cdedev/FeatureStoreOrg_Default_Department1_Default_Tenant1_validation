// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377798987/ThreeG.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ThreeG

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseThreeGListener is a complete listener for a parse tree produced by ThreeGParser.
type BaseThreeGListener struct{}

var _ ThreeGListener = &BaseThreeGListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseThreeGListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseThreeGListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseThreeGListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseThreeGListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseThreeGListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseThreeGListener) ExitExpression(ctx *ExpressionContext) {}

// EnterThreeg is called when production threeg is entered.
func (s *BaseThreeGListener) EnterThreeg(ctx *ThreegContext) {}

// ExitThreeg is called when production threeg is exited.
func (s *BaseThreeGListener) ExitThreeg(ctx *ThreegContext) {}
