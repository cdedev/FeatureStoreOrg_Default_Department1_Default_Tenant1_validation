// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672981954579/InventoryCategory.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // InventoryCategory

import "github.com/antlr/antlr4/runtime/Go/antlr"

// InventoryCategoryListener is a complete listener for a parse tree produced by InventoryCategoryParser.
type InventoryCategoryListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterInventorycategory is called when entering the inventorycategory production.
	EnterInventorycategory(c *InventorycategoryContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitInventorycategory is called when exiting the inventorycategory production.
	ExitInventorycategory(c *InventorycategoryContext)
}
