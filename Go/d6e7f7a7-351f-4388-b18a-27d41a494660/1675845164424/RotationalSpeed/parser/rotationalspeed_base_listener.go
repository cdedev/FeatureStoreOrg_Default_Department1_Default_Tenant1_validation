// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675845164424/RotationalSpeed.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RotationalSpeed

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRotationalSpeedListener is a complete listener for a parse tree produced by RotationalSpeedParser.
type BaseRotationalSpeedListener struct{}

var _ RotationalSpeedListener = &BaseRotationalSpeedListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRotationalSpeedListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRotationalSpeedListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRotationalSpeedListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRotationalSpeedListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRotationalSpeedListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRotationalSpeedListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRotationalspeed is called when production rotationalspeed is entered.
func (s *BaseRotationalSpeedListener) EnterRotationalspeed(ctx *RotationalspeedContext) {}

// ExitRotationalspeed is called when production rotationalspeed is exited.
func (s *BaseRotationalSpeedListener) ExitRotationalspeed(ctx *RotationalspeedContext) {}
