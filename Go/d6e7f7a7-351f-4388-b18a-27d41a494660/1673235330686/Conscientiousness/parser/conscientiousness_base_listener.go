// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673235330686/Conscientiousness.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Conscientiousness

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseConscientiousnessListener is a complete listener for a parse tree produced by ConscientiousnessParser.
type BaseConscientiousnessListener struct{}

var _ ConscientiousnessListener = &BaseConscientiousnessListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseConscientiousnessListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseConscientiousnessListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseConscientiousnessListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseConscientiousnessListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseConscientiousnessListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseConscientiousnessListener) ExitExpression(ctx *ExpressionContext) {}

// EnterConscientiousness is called when production conscientiousness is entered.
func (s *BaseConscientiousnessListener) EnterConscientiousness(ctx *ConscientiousnessContext) {}

// ExitConscientiousness is called when production conscientiousness is exited.
func (s *BaseConscientiousnessListener) ExitConscientiousness(ctx *ConscientiousnessContext) {}
