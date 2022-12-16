// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671203166456/Terrace.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Terrace

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTerraceListener is a complete listener for a parse tree produced by TerraceParser.
type BaseTerraceListener struct{}

var _ TerraceListener = &BaseTerraceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTerraceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTerraceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTerraceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTerraceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTerraceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTerraceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTerrace is called when production terrace is entered.
func (s *BaseTerraceListener) EnterTerrace(ctx *TerraceContext) {}

// ExitTerrace is called when production terrace is exited.
func (s *BaseTerraceListener) ExitTerrace(ctx *TerraceContext) {}
