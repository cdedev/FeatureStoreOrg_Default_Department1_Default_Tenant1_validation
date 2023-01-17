// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673925094551/Breathes.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Breathes

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBreathesListener is a complete listener for a parse tree produced by BreathesParser.
type BaseBreathesListener struct{}

var _ BreathesListener = &BaseBreathesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBreathesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBreathesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBreathesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBreathesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBreathesListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBreathesListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBreathes is called when production breathes is entered.
func (s *BaseBreathesListener) EnterBreathes(ctx *BreathesContext) {}

// ExitBreathes is called when production breathes is exited.
func (s *BaseBreathesListener) ExitBreathes(ctx *BreathesContext) {}
