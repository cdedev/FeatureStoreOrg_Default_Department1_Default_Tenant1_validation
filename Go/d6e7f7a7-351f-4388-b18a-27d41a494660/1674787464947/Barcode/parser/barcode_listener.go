// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674787464947/Barcode.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Barcode

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BarcodeListener is a complete listener for a parse tree produced by BarcodeParser.
type BarcodeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBarcode is called when entering the barcode production.
	EnterBarcode(c *BarcodeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBarcode is called when exiting the barcode production.
	ExitBarcode(c *BarcodeContext)
}
