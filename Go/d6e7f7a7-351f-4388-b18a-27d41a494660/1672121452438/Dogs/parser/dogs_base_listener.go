// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672121452438/Dogs.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Dogs

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDogsListener is a complete listener for a parse tree produced by DogsParser.
type BaseDogsListener struct{}

var _ DogsListener = &BaseDogsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDogsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDogsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDogsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDogsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDogsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDogsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDogs is called when production dogs is entered.
func (s *BaseDogsListener) EnterDogs(ctx *DogsContext) {}

// ExitDogs is called when production dogs is exited.
func (s *BaseDogsListener) ExitDogs(ctx *DogsContext) {}
