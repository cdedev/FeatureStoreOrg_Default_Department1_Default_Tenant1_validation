// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671186145445/WaterLevel.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // WaterLevel

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseWaterLevelListener is a complete listener for a parse tree produced by WaterLevelParser.
type BaseWaterLevelListener struct{}

var _ WaterLevelListener = &BaseWaterLevelListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWaterLevelListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWaterLevelListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWaterLevelListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWaterLevelListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseWaterLevelListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseWaterLevelListener) ExitExpression(ctx *ExpressionContext) {}

// EnterWaterlevel is called when production waterlevel is entered.
func (s *BaseWaterLevelListener) EnterWaterlevel(ctx *WaterlevelContext) {}

// ExitWaterlevel is called when production waterlevel is exited.
func (s *BaseWaterLevelListener) ExitWaterlevel(ctx *WaterlevelContext) {}
