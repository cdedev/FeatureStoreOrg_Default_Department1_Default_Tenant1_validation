// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672121133708/NonHumanPrimates.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // NonHumanPrimates

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNonHumanPrimatesListener is a complete listener for a parse tree produced by NonHumanPrimatesParser.
type BaseNonHumanPrimatesListener struct{}

var _ NonHumanPrimatesListener = &BaseNonHumanPrimatesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNonHumanPrimatesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNonHumanPrimatesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNonHumanPrimatesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNonHumanPrimatesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseNonHumanPrimatesListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseNonHumanPrimatesListener) ExitExpression(ctx *ExpressionContext) {}

// EnterNonhumanprimates is called when production nonhumanprimates is entered.
func (s *BaseNonHumanPrimatesListener) EnterNonhumanprimates(ctx *NonhumanprimatesContext) {}

// ExitNonhumanprimates is called when production nonhumanprimates is exited.
func (s *BaseNonHumanPrimatesListener) ExitNonhumanprimates(ctx *NonhumanprimatesContext) {}
