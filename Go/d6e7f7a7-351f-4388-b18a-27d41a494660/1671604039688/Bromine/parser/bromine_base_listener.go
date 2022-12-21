// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671604039688/Bromine.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Bromine

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBromineListener is a complete listener for a parse tree produced by BromineParser.
type BaseBromineListener struct{}

var _ BromineListener = &BaseBromineListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBromineListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBromineListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBromineListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBromineListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBromineListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBromineListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBromine is called when production bromine is entered.
func (s *BaseBromineListener) EnterBromine(ctx *BromineContext) {}

// ExitBromine is called when production bromine is exited.
func (s *BaseBromineListener) ExitBromine(ctx *BromineContext) {}
