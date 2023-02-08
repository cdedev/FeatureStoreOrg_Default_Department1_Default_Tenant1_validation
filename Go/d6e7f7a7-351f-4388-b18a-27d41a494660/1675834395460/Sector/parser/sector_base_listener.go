// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675834395460/Sector.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sector

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSectorListener is a complete listener for a parse tree produced by SectorParser.
type BaseSectorListener struct{}

var _ SectorListener = &BaseSectorListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSectorListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSectorListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSectorListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSectorListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSectorListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSectorListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSector is called when production sector is entered.
func (s *BaseSectorListener) EnterSector(ctx *SectorContext) {}

// ExitSector is called when production sector is exited.
func (s *BaseSectorListener) ExitSector(ctx *SectorContext) {}
