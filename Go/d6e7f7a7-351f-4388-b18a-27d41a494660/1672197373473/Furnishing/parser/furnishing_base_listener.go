// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672197373473/Furnishing.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Furnishing

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFurnishingListener is a complete listener for a parse tree produced by FurnishingParser.
type BaseFurnishingListener struct{}

var _ FurnishingListener = &BaseFurnishingListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFurnishingListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFurnishingListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFurnishingListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFurnishingListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFurnishingListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFurnishingListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFurnishing is called when production furnishing is entered.
func (s *BaseFurnishingListener) EnterFurnishing(ctx *FurnishingContext) {}

// ExitFurnishing is called when production furnishing is exited.
func (s *BaseFurnishingListener) ExitFurnishing(ctx *FurnishingContext) {}
