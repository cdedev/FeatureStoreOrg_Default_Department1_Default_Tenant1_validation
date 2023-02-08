// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675836088799/Birthrate.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Birthrate

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBirthrateListener is a complete listener for a parse tree produced by BirthrateParser.
type BaseBirthrateListener struct{}

var _ BirthrateListener = &BaseBirthrateListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBirthrateListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBirthrateListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBirthrateListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBirthrateListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBirthrateListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBirthrateListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBirthrate is called when production birthrate is entered.
func (s *BaseBirthrateListener) EnterBirthrate(ctx *BirthrateContext) {}

// ExitBirthrate is called when production birthrate is exited.
func (s *BaseBirthrateListener) ExitBirthrate(ctx *BirthrateContext) {}
