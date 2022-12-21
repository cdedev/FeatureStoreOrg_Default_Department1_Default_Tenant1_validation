// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671604404766/Cesium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cesium

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCesiumListener is a complete listener for a parse tree produced by CesiumParser.
type BaseCesiumListener struct{}

var _ CesiumListener = &BaseCesiumListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCesiumListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCesiumListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCesiumListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCesiumListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCesiumListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCesiumListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCesium is called when production cesium is entered.
func (s *BaseCesiumListener) EnterCesium(ctx *CesiumContext) {}

// ExitCesium is called when production cesium is exited.
func (s *BaseCesiumListener) ExitCesium(ctx *CesiumContext) {}
