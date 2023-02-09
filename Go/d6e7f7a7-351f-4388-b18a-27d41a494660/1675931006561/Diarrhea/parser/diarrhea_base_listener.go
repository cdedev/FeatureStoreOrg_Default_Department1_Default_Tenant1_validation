// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675931006561/Diarrhea.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Diarrhea

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDiarrheaListener is a complete listener for a parse tree produced by DiarrheaParser.
type BaseDiarrheaListener struct{}

var _ DiarrheaListener = &BaseDiarrheaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDiarrheaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDiarrheaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDiarrheaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDiarrheaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDiarrheaListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDiarrheaListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDiarrhea is called when production diarrhea is entered.
func (s *BaseDiarrheaListener) EnterDiarrhea(ctx *DiarrheaContext) {}

// ExitDiarrhea is called when production diarrhea is exited.
func (s *BaseDiarrheaListener) ExitDiarrhea(ctx *DiarrheaContext) {}
