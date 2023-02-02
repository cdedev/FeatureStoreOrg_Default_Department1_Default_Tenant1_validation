// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675309170263/Cost.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cost

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCostListener is a complete listener for a parse tree produced by CostParser.
type BaseCostListener struct{}

var _ CostListener = &BaseCostListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCostListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCostListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCostListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCostListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCostListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCostListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCost is called when production cost is entered.
func (s *BaseCostListener) EnterCost(ctx *CostContext) {}

// ExitCost is called when production cost is exited.
func (s *BaseCostListener) ExitCost(ctx *CostContext) {}
