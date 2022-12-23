// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671766113878/House.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // House

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseHouseListener is a complete listener for a parse tree produced by HouseParser.
type BaseHouseListener struct{}

var _ HouseListener = &BaseHouseListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHouseListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHouseListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHouseListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHouseListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseHouseListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseHouseListener) ExitExpression(ctx *ExpressionContext) {}

// EnterHouse is called when production house is entered.
func (s *BaseHouseListener) EnterHouse(ctx *HouseContext) {}

// ExitHouse is called when production house is exited.
func (s *BaseHouseListener) ExitHouse(ctx *HouseContext) {}
