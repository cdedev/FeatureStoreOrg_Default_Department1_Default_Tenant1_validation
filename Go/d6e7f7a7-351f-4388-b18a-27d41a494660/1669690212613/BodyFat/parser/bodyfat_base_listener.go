// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669690212613/BodyFat.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BodyFat

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBodyFatListener is a complete listener for a parse tree produced by BodyFatParser.
type BaseBodyFatListener struct{}

var _ BodyFatListener = &BaseBodyFatListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBodyFatListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBodyFatListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBodyFatListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBodyFatListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBodyFatListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBodyFatListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBodyfat is called when production bodyfat is entered.
func (s *BaseBodyFatListener) EnterBodyfat(ctx *BodyfatContext) {}

// ExitBodyfat is called when production bodyfat is exited.
func (s *BaseBodyFatListener) ExitBodyfat(ctx *BodyfatContext) {}
