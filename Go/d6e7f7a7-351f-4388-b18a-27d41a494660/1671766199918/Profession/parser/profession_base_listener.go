// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671766199918/Profession.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Profession

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseProfessionListener is a complete listener for a parse tree produced by ProfessionParser.
type BaseProfessionListener struct{}

var _ ProfessionListener = &BaseProfessionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseProfessionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseProfessionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseProfessionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseProfessionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseProfessionListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseProfessionListener) ExitExpression(ctx *ExpressionContext) {}

// EnterProfession is called when production profession is entered.
func (s *BaseProfessionListener) EnterProfession(ctx *ProfessionContext) {}

// ExitProfession is called when production profession is exited.
func (s *BaseProfessionListener) ExitProfession(ctx *ProfessionContext) {}
