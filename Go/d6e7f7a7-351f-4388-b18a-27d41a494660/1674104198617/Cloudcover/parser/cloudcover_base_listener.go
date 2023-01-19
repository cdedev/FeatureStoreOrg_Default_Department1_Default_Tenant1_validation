// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674104198617/Cloudcover.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cloudcover

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCloudcoverListener is a complete listener for a parse tree produced by CloudcoverParser.
type BaseCloudcoverListener struct{}

var _ CloudcoverListener = &BaseCloudcoverListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCloudcoverListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCloudcoverListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCloudcoverListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCloudcoverListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCloudcoverListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCloudcoverListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCloudcover is called when production cloudcover is entered.
func (s *BaseCloudcoverListener) EnterCloudcover(ctx *CloudcoverContext) {}

// ExitCloudcover is called when production cloudcover is exited.
func (s *BaseCloudcoverListener) ExitCloudcover(ctx *CloudcoverContext) {}
