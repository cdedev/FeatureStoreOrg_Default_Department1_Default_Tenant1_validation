// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672282520870/EverMarried.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // EverMarried

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseEverMarriedListener is a complete listener for a parse tree produced by EverMarriedParser.
type BaseEverMarriedListener struct{}

var _ EverMarriedListener = &BaseEverMarriedListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEverMarriedListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEverMarriedListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEverMarriedListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEverMarriedListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseEverMarriedListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseEverMarriedListener) ExitExpression(ctx *ExpressionContext) {}

// EnterEvermarried is called when production evermarried is entered.
func (s *BaseEverMarriedListener) EnterEvermarried(ctx *EvermarriedContext) {}

// ExitEvermarried is called when production evermarried is exited.
func (s *BaseEverMarriedListener) ExitEvermarried(ctx *EvermarriedContext) {}
