// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672637923711/Product.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Product

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseProductListener is a complete listener for a parse tree produced by ProductParser.
type BaseProductListener struct{}

var _ ProductListener = &BaseProductListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseProductListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseProductListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseProductListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseProductListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseProductListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseProductListener) ExitExpression(ctx *ExpressionContext) {}

// EnterProduct is called when production product is entered.
func (s *BaseProductListener) EnterProduct(ctx *ProductContext) {}

// ExitProduct is called when production product is exited.
func (s *BaseProductListener) ExitProduct(ctx *ProductContext) {}
