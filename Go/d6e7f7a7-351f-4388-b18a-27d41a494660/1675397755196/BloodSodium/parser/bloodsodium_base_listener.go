// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675397755196/BloodSodium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BloodSodium

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBloodSodiumListener is a complete listener for a parse tree produced by BloodSodiumParser.
type BaseBloodSodiumListener struct{}

var _ BloodSodiumListener = &BaseBloodSodiumListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBloodSodiumListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBloodSodiumListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBloodSodiumListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBloodSodiumListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBloodSodiumListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBloodSodiumListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBloodsodium is called when production bloodsodium is entered.
func (s *BaseBloodSodiumListener) EnterBloodsodium(ctx *BloodsodiumContext) {}

// ExitBloodsodium is called when production bloodsodium is exited.
func (s *BaseBloodSodiumListener) ExitBloodsodium(ctx *BloodsodiumContext) {}
