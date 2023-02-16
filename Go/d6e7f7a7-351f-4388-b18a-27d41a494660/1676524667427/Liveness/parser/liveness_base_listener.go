// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676524667427/Liveness.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Liveness

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLivenessListener is a complete listener for a parse tree produced by LivenessParser.
type BaseLivenessListener struct{}

var _ LivenessListener = &BaseLivenessListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLivenessListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLivenessListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLivenessListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLivenessListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseLivenessListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseLivenessListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLiveness is called when production liveness is entered.
func (s *BaseLivenessListener) EnterLiveness(ctx *LivenessContext) {}

// ExitLiveness is called when production liveness is exited.
func (s *BaseLivenessListener) ExitLiveness(ctx *LivenessContext) {}
