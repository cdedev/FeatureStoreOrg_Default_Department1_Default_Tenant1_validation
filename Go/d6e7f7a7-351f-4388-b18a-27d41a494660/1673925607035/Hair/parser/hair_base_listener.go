// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673925607035/Hair.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Hair

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseHairListener is a complete listener for a parse tree produced by HairParser.
type BaseHairListener struct{}

var _ HairListener = &BaseHairListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHairListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHairListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHairListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHairListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseHairListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseHairListener) ExitExpression(ctx *ExpressionContext) {}

// EnterHair is called when production hair is entered.
func (s *BaseHairListener) EnterHair(ctx *HairContext) {}

// ExitHair is called when production hair is exited.
func (s *BaseHairListener) ExitHair(ctx *HairContext) {}
