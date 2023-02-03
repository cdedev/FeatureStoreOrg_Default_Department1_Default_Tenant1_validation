// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675404923351/Tilt.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Tilt

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTiltListener is a complete listener for a parse tree produced by TiltParser.
type BaseTiltListener struct{}

var _ TiltListener = &BaseTiltListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTiltListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTiltListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTiltListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTiltListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTiltListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTiltListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTilt is called when production tilt is entered.
func (s *BaseTiltListener) EnterTilt(ctx *TiltContext) {}

// ExitTilt is called when production tilt is exited.
func (s *BaseTiltListener) ExitTilt(ctx *TiltContext) {}
