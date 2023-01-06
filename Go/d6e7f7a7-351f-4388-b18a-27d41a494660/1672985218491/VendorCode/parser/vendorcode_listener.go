// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672985218491/VendorCode.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // VendorCode

import "github.com/antlr/antlr4/runtime/Go/antlr"

// VendorCodeListener is a complete listener for a parse tree produced by VendorCodeParser.
type VendorCodeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterVendorcode is called when entering the vendorcode production.
	EnterVendorcode(c *VendorcodeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitVendorcode is called when exiting the vendorcode production.
	ExitVendorcode(c *VendorcodeContext)
}
