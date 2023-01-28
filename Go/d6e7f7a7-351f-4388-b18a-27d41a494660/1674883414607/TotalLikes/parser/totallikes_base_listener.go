// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674883414607/TotalLikes.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // TotalLikes

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTotalLikesListener is a complete listener for a parse tree produced by TotalLikesParser.
type BaseTotalLikesListener struct{}

var _ TotalLikesListener = &BaseTotalLikesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTotalLikesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTotalLikesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTotalLikesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTotalLikesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTotalLikesListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTotalLikesListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTotallikes is called when production totallikes is entered.
func (s *BaseTotalLikesListener) EnterTotallikes(ctx *TotallikesContext) {}

// ExitTotallikes is called when production totallikes is exited.
func (s *BaseTotalLikesListener) ExitTotallikes(ctx *TotallikesContext) {}
