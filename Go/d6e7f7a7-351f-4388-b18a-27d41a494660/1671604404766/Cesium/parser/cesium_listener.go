// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671604404766/Cesium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cesium

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CesiumListener is a complete listener for a parse tree produced by CesiumParser.
type CesiumListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCesium is called when entering the cesium production.
	EnterCesium(c *CesiumContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCesium is called when exiting the cesium production.
	ExitCesium(c *CesiumContext)
}
