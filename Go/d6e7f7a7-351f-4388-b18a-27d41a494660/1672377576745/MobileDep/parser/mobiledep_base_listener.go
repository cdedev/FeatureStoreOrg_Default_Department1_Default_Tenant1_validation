// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377576745/MobileDep.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MobileDep

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMobileDepListener is a complete listener for a parse tree produced by MobileDepParser.
type BaseMobileDepListener struct{}

var _ MobileDepListener = &BaseMobileDepListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMobileDepListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMobileDepListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMobileDepListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMobileDepListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMobileDepListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMobileDepListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMobiledep is called when production mobiledep is entered.
func (s *BaseMobileDepListener) EnterMobiledep(ctx *MobiledepContext) {}

// ExitMobiledep is called when production mobiledep is exited.
func (s *BaseMobileDepListener) ExitMobiledep(ctx *MobiledepContext) {}
