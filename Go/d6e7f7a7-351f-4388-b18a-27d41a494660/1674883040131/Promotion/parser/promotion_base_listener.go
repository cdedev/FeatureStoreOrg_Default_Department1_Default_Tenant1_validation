// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674883040131/Promotion.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Promotion

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePromotionListener is a complete listener for a parse tree produced by PromotionParser.
type BasePromotionListener struct{}

var _ PromotionListener = &BasePromotionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePromotionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePromotionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePromotionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePromotionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePromotionListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePromotionListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPromotion is called when production promotion is entered.
func (s *BasePromotionListener) EnterPromotion(ctx *PromotionContext) {}

// ExitPromotion is called when production promotion is exited.
func (s *BasePromotionListener) ExitPromotion(ctx *PromotionContext) {}
