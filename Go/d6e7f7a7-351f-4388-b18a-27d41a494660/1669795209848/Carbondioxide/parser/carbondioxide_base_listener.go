// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669795209848/Carbondioxide.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Carbondioxide

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCarbondioxideListener is a complete listener for a parse tree produced by CarbondioxideParser.
type BaseCarbondioxideListener struct{}

var _ CarbondioxideListener = &BaseCarbondioxideListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCarbondioxideListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCarbondioxideListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCarbondioxideListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCarbondioxideListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCarbondioxideListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCarbondioxideListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCarbondioxide is called when production carbondioxide is entered.
func (s *BaseCarbondioxideListener) EnterCarbondioxide(ctx *CarbondioxideContext) {}

// ExitCarbondioxide is called when production carbondioxide is exited.
func (s *BaseCarbondioxideListener) ExitCarbondioxide(ctx *CarbondioxideContext) {}
