// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672981954579/InventoryCategory.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // InventoryCategory

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseInventoryCategoryListener is a complete listener for a parse tree produced by InventoryCategoryParser.
type BaseInventoryCategoryListener struct{}

var _ InventoryCategoryListener = &BaseInventoryCategoryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseInventoryCategoryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseInventoryCategoryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseInventoryCategoryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseInventoryCategoryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseInventoryCategoryListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseInventoryCategoryListener) ExitExpression(ctx *ExpressionContext) {}

// EnterInventorycategory is called when production inventorycategory is entered.
func (s *BaseInventoryCategoryListener) EnterInventorycategory(ctx *InventorycategoryContext) {}

// ExitInventorycategory is called when production inventorycategory is exited.
func (s *BaseInventoryCategoryListener) ExitInventorycategory(ctx *InventorycategoryContext) {}
