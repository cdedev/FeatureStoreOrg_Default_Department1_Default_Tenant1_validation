// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671193515509/ReviewsCount.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ReviewsCount

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseReviewsCountListener is a complete listener for a parse tree produced by ReviewsCountParser.
type BaseReviewsCountListener struct{}

var _ ReviewsCountListener = &BaseReviewsCountListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseReviewsCountListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseReviewsCountListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseReviewsCountListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseReviewsCountListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseReviewsCountListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseReviewsCountListener) ExitExpression(ctx *ExpressionContext) {}

// EnterReviewscount is called when production reviewscount is entered.
func (s *BaseReviewsCountListener) EnterReviewscount(ctx *ReviewscountContext) {}

// ExitReviewscount is called when production reviewscount is exited.
func (s *BaseReviewsCountListener) ExitReviewscount(ctx *ReviewscountContext) {}
