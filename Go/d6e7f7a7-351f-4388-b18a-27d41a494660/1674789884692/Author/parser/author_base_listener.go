// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674789884692/Author.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Author

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAuthorListener is a complete listener for a parse tree produced by AuthorParser.
type BaseAuthorListener struct{}

var _ AuthorListener = &BaseAuthorListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAuthorListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAuthorListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAuthorListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAuthorListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAuthorListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAuthorListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAuthor is called when production author is entered.
func (s *BaseAuthorListener) EnterAuthor(ctx *AuthorContext) {}

// ExitAuthor is called when production author is exited.
func (s *BaseAuthorListener) ExitAuthor(ctx *AuthorContext) {}
