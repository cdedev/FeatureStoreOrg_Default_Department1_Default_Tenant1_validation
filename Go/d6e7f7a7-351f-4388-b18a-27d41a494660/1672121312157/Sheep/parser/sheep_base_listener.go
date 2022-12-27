// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672121312157/Sheep.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sheep

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSheepListener is a complete listener for a parse tree produced by SheepParser.
type BaseSheepListener struct{}

var _ SheepListener = &BaseSheepListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSheepListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSheepListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSheepListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSheepListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSheepListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSheepListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSheep is called when production sheep is entered.
func (s *BaseSheepListener) EnterSheep(ctx *SheepContext) {}

// ExitSheep is called when production sheep is exited.
func (s *BaseSheepListener) ExitSheep(ctx *SheepContext) {}
