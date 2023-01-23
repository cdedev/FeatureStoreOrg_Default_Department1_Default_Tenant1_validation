// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674454539826/Assets.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Assets

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAssetsListener is a complete listener for a parse tree produced by AssetsParser.
type BaseAssetsListener struct{}

var _ AssetsListener = &BaseAssetsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAssetsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAssetsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAssetsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAssetsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAssetsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAssetsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAssets is called when production assets is entered.
func (s *BaseAssetsListener) EnterAssets(ctx *AssetsContext) {}

// ExitAssets is called when production assets is exited.
func (s *BaseAssetsListener) ExitAssets(ctx *AssetsContext) {}
