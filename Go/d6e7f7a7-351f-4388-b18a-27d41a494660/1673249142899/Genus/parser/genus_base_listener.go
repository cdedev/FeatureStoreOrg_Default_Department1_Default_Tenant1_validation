// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673249142899/Genus.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Genus

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGenusListener is a complete listener for a parse tree produced by GenusParser.
type BaseGenusListener struct{}

var _ GenusListener = &BaseGenusListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGenusListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGenusListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGenusListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGenusListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGenusListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGenusListener) ExitExpression(ctx *ExpressionContext) {}

// EnterGenus is called when production genus is entered.
func (s *BaseGenusListener) EnterGenus(ctx *GenusContext) {}

// ExitGenus is called when production genus is exited.
func (s *BaseGenusListener) ExitGenus(ctx *GenusContext) {}
