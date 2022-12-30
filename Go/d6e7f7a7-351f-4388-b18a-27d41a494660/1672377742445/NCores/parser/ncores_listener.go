// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377742445/NCores.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // NCores

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NCoresListener is a complete listener for a parse tree produced by NCoresParser.
type NCoresListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterNcores is called when entering the ncores production.
	EnterNcores(c *NcoresContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitNcores is called when exiting the ncores production.
	ExitNcores(c *NcoresContext)
}
