// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675242106853/Vendor.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Vendor

import "github.com/antlr/antlr4/runtime/Go/antlr"

// VendorListener is a complete listener for a parse tree produced by VendorParser.
type VendorListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterVendor is called when entering the vendor production.
	EnterVendor(c *VendorContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitVendor is called when exiting the vendor production.
	ExitVendor(c *VendorContext)
}
