// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674541115725/Married.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Married

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMarriedListener is a complete listener for a parse tree produced by MarriedParser.
type BaseMarriedListener struct{}

var _ MarriedListener = &BaseMarriedListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMarriedListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMarriedListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMarriedListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMarriedListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMarriedListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMarriedListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMarried is called when production married is entered.
func (s *BaseMarriedListener) EnterMarried(ctx *MarriedContext) {}

// ExitMarried is called when production married is exited.
func (s *BaseMarriedListener) ExitMarried(ctx *MarriedContext) {}
