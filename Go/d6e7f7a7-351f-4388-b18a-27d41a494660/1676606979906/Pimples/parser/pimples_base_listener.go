// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676606979906/Pimples.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Pimples

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePimplesListener is a complete listener for a parse tree produced by PimplesParser.
type BasePimplesListener struct{}

var _ PimplesListener = &BasePimplesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePimplesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePimplesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePimplesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePimplesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePimplesListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePimplesListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPimples is called when production pimples is entered.
func (s *BasePimplesListener) EnterPimples(ctx *PimplesContext) {}

// ExitPimples is called when production pimples is exited.
func (s *BasePimplesListener) ExitPimples(ctx *PimplesContext) {}
