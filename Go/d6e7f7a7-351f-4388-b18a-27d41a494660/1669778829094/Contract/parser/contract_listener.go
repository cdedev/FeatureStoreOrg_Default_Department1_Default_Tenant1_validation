// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669778829094/Contract.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Contract

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ContractListener is a complete listener for a parse tree produced by ContractParser.
type ContractListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterContract is called when entering the contract production.
	EnterContract(c *ContractContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitContract is called when exiting the contract production.
	ExitContract(c *ContractContext)
}
