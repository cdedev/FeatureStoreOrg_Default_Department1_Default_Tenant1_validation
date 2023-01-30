// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675069518379/Review.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Review

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseReviewListener is a complete listener for a parse tree produced by ReviewParser.
type BaseReviewListener struct{}

var _ ReviewListener = &BaseReviewListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseReviewListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseReviewListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseReviewListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseReviewListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseReviewListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseReviewListener) ExitExpression(ctx *ExpressionContext) {}

// EnterReview is called when production review is entered.
func (s *BaseReviewListener) EnterReview(ctx *ReviewContext) {}

// ExitReview is called when production review is exited.
func (s *BaseReviewListener) ExitReview(ctx *ReviewContext) {}
