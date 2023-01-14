// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673675405429/Reviews.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Reviews

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseReviewsListener is a complete listener for a parse tree produced by ReviewsParser.
type BaseReviewsListener struct{}

var _ ReviewsListener = &BaseReviewsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseReviewsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseReviewsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseReviewsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseReviewsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseReviewsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseReviewsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterReviews is called when production reviews is entered.
func (s *BaseReviewsListener) EnterReviews(ctx *ReviewsContext) {}

// ExitReviews is called when production reviews is exited.
func (s *BaseReviewsListener) ExitReviews(ctx *ReviewsContext) {}
