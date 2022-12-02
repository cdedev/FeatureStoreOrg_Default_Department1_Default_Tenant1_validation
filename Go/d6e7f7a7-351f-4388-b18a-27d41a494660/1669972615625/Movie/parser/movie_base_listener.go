// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669972615625/Movie.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Movie

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMovieListener is a complete listener for a parse tree produced by MovieParser.
type BaseMovieListener struct{}

var _ MovieListener = &BaseMovieListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMovieListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMovieListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMovieListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMovieListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMovieListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMovieListener) ExitExpression(ctx *ExpressionContext) {}

// EnterField0 is called when production field0 is entered.
func (s *BaseMovieListener) EnterField0(ctx *Field0Context) {}

// ExitField0 is called when production field0 is exited.
func (s *BaseMovieListener) ExitField0(ctx *Field0Context) {}
