// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675740349484/BuyingPower.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BuyingPower

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBuyingPowerListener is a complete listener for a parse tree produced by BuyingPowerParser.
type BaseBuyingPowerListener struct{}

var _ BuyingPowerListener = &BaseBuyingPowerListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBuyingPowerListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBuyingPowerListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBuyingPowerListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBuyingPowerListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBuyingPowerListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBuyingPowerListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBuyingpower is called when production buyingpower is entered.
func (s *BaseBuyingPowerListener) EnterBuyingpower(ctx *BuyingpowerContext) {}

// ExitBuyingpower is called when production buyingpower is exited.
func (s *BaseBuyingPowerListener) ExitBuyingpower(ctx *BuyingpowerContext) {}
