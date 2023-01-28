// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674881674788/Weather.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Weather

import "github.com/antlr/antlr4/runtime/Go/antlr"

// WeatherListener is a complete listener for a parse tree produced by WeatherParser.
type WeatherListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterWeather is called when entering the weather production.
	EnterWeather(c *WeatherContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitWeather is called when exiting the weather production.
	ExitWeather(c *WeatherContext)
}
