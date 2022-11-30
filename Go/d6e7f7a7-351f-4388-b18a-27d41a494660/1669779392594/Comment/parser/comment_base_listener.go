// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669779392594/Comment.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Comment

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCommentListener is a complete listener for a parse tree produced by CommentParser.
type BaseCommentListener struct{}

var _ CommentListener = &BaseCommentListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCommentListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCommentListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCommentListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCommentListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCommentListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCommentListener) ExitExpression(ctx *ExpressionContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseCommentListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseCommentListener) ExitComment(ctx *CommentContext) {}
