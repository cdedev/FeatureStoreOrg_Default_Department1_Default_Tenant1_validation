// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671771286494/WeatherTemperature.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // WeatherTemperature

import "github.com/antlr/antlr4/runtime/Go/antlr"

// WeatherTemperatureListener is a complete listener for a parse tree produced by WeatherTemperatureParser.
type WeatherTemperatureListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterWeathertemperature is called when entering the weathertemperature production.
	EnterWeathertemperature(c *WeathertemperatureContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitWeathertemperature is called when exiting the weathertemperature production.
	ExitWeathertemperature(c *WeathertemperatureContext)
}
