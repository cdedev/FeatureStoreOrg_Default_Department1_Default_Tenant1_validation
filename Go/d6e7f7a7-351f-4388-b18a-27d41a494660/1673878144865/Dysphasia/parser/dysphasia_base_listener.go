// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673878144865/Dysphasia.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Dysphasia

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDysphasiaListener is a complete listener for a parse tree produced by DysphasiaParser.
type BaseDysphasiaListener struct{}

var _ DysphasiaListener = &BaseDysphasiaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDysphasiaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDysphasiaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDysphasiaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDysphasiaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDysphasiaListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDysphasiaListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDysphasia is called when production dysphasia is entered.
func (s *BaseDysphasiaListener) EnterDysphasia(ctx *DysphasiaContext) {}

// ExitDysphasia is called when production dysphasia is exited.
func (s *BaseDysphasiaListener) ExitDysphasia(ctx *DysphasiaContext) {}
