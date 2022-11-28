// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669657491481/Genre.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Genre

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGenreListener is a complete listener for a parse tree produced by GenreParser.
type BaseGenreListener struct{}

var _ GenreListener = &BaseGenreListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGenreListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGenreListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGenreListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGenreListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGenreListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGenreListener) ExitExpression(ctx *ExpressionContext) {}

// EnterGenre is called when production genre is entered.
func (s *BaseGenreListener) EnterGenre(ctx *GenreContext) {}

// ExitGenre is called when production genre is exited.
func (s *BaseGenreListener) ExitGenre(ctx *GenreContext) {}
