// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671866796744/CloneSize.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CloneSize

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCloneSizeListener is a complete listener for a parse tree produced by CloneSizeParser.
type BaseCloneSizeListener struct{}

var _ CloneSizeListener = &BaseCloneSizeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCloneSizeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCloneSizeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCloneSizeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCloneSizeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCloneSizeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCloneSizeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterClonesize is called when production clonesize is entered.
func (s *BaseCloneSizeListener) EnterClonesize(ctx *ClonesizeContext) {}

// ExitClonesize is called when production clonesize is exited.
func (s *BaseCloneSizeListener) ExitClonesize(ctx *ClonesizeContext) {}
