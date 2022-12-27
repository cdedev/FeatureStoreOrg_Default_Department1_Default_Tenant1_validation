// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672116531886/PaternityScore.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // PaternityScore

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePaternityScoreListener is a complete listener for a parse tree produced by PaternityScoreParser.
type BasePaternityScoreListener struct{}

var _ PaternityScoreListener = &BasePaternityScoreListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePaternityScoreListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePaternityScoreListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePaternityScoreListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePaternityScoreListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePaternityScoreListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePaternityScoreListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPaternityscore is called when production paternityscore is entered.
func (s *BasePaternityScoreListener) EnterPaternityscore(ctx *PaternityscoreContext) {}

// ExitPaternityscore is called when production paternityscore is exited.
func (s *BasePaternityScoreListener) ExitPaternityscore(ctx *PaternityscoreContext) {}
