// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676260183645/IpAddressPresent.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // IpAddressPresent

import "github.com/antlr/antlr4/runtime/Go/antlr"

// IpAddressPresentListener is a complete listener for a parse tree produced by IpAddressPresentParser.
type IpAddressPresentListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterIpaddresspresent is called when entering the ipaddresspresent production.
	EnterIpaddresspresent(c *IpaddresspresentContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitIpaddresspresent is called when exiting the ipaddresspresent production.
	ExitIpaddresspresent(c *IpaddresspresentContext)
}
