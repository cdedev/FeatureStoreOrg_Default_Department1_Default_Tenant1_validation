// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669793633322/Solidity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Solidity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SolidityListener is a complete listener for a parse tree produced by SolidityParser.
type SolidityListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSolidity is called when entering the solidity production.
	EnterSolidity(c *SolidityContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSolidity is called when exiting the solidity production.
	ExitSolidity(c *SolidityContext)
}
