// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673924919318/Aquatic.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Aquatic

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAquaticListener is a complete listener for a parse tree produced by AquaticParser.
type BaseAquaticListener struct{}

var _ AquaticListener = &BaseAquaticListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAquaticListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAquaticListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAquaticListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAquaticListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAquaticListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAquaticListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAquatic is called when production aquatic is entered.
func (s *BaseAquaticListener) EnterAquatic(ctx *AquaticContext) {}

// ExitAquatic is called when production aquatic is exited.
func (s *BaseAquaticListener) ExitAquatic(ctx *AquaticContext) {}
