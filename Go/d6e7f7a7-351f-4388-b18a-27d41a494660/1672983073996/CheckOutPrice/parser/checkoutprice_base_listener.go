// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672983073996/CheckOutPrice.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CheckOutPrice

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCheckOutPriceListener is a complete listener for a parse tree produced by CheckOutPriceParser.
type BaseCheckOutPriceListener struct{}

var _ CheckOutPriceListener = &BaseCheckOutPriceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCheckOutPriceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCheckOutPriceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCheckOutPriceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCheckOutPriceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCheckOutPriceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCheckOutPriceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCheckoutprice is called when production checkoutprice is entered.
func (s *BaseCheckOutPriceListener) EnterCheckoutprice(ctx *CheckoutpriceContext) {}

// ExitCheckoutprice is called when production checkoutprice is exited.
func (s *BaseCheckOutPriceListener) ExitCheckoutprice(ctx *CheckoutpriceContext) {}
