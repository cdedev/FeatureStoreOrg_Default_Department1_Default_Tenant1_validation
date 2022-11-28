// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669636205938/Rating.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rating

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRatingListener is a complete listener for a parse tree produced by RatingParser.
type BaseRatingListener struct{}

var _ RatingListener = &BaseRatingListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRatingListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRatingListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRatingListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRatingListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRatingListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRatingListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRating is called when production rating is entered.
func (s *BaseRatingListener) EnterRating(ctx *RatingContext) {}

// ExitRating is called when production rating is exited.
func (s *BaseRatingListener) ExitRating(ctx *RatingContext) {}
