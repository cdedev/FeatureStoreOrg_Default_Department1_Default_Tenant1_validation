// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675072357370/Sociocultural.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sociocultural

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SocioculturalListener is a complete listener for a parse tree produced by SocioculturalParser.
type SocioculturalListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSociocultural is called when entering the sociocultural production.
	EnterSociocultural(c *SocioculturalContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSociocultural is called when exiting the sociocultural production.
	ExitSociocultural(c *SocioculturalContext)
}
