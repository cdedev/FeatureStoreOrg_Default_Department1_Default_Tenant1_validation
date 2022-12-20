// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671530343405/SubCategory.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SubCategory

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SubCategoryListener is a complete listener for a parse tree produced by SubCategoryParser.
type SubCategoryListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSubcategory is called when entering the subcategory production.
	EnterSubcategory(c *SubcategoryContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSubcategory is called when exiting the subcategory production.
	ExitSubcategory(c *SubcategoryContext)
}
