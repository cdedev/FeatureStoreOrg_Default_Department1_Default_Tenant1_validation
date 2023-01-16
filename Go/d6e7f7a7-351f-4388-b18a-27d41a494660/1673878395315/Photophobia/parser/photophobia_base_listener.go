// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673878395315/Photophobia.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Photophobia

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePhotophobiaListener is a complete listener for a parse tree produced by PhotophobiaParser.
type BasePhotophobiaListener struct{}

var _ PhotophobiaListener = &BasePhotophobiaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePhotophobiaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePhotophobiaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePhotophobiaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePhotophobiaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePhotophobiaListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePhotophobiaListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPhotophobia is called when production photophobia is entered.
func (s *BasePhotophobiaListener) EnterPhotophobia(ctx *PhotophobiaContext) {}

// ExitPhotophobia is called when production photophobia is exited.
func (s *BasePhotophobiaListener) ExitPhotophobia(ctx *PhotophobiaContext) {}
