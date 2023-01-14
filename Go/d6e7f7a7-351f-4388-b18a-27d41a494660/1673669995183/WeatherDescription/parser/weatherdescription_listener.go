// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673669995183/WeatherDescription.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // WeatherDescription

import "github.com/antlr/antlr4/runtime/Go/antlr"

// WeatherDescriptionListener is a complete listener for a parse tree produced by WeatherDescriptionParser.
type WeatherDescriptionListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterWeatherdescription is called when entering the weatherdescription production.
	EnterWeatherdescription(c *WeatherdescriptionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitWeatherdescription is called when exiting the weatherdescription production.
	ExitWeatherdescription(c *WeatherdescriptionContext)
}
