// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671771368345/WeatherWind.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // WeatherWind

import "github.com/antlr/antlr4/runtime/Go/antlr"

// WeatherWindListener is a complete listener for a parse tree produced by WeatherWindParser.
type WeatherWindListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterWeatherwind is called when entering the weatherwind production.
	EnterWeatherwind(c *WeatherwindContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitWeatherwind is called when exiting the weatherwind production.
	ExitWeatherwind(c *WeatherwindContext)
}
