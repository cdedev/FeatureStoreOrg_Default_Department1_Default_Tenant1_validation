// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377628730/BlueTooth.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BlueTooth

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BlueToothListener is a complete listener for a parse tree produced by BlueToothParser.
type BlueToothListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBluetooth is called when entering the bluetooth production.
	EnterBluetooth(c *BluetoothContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBluetooth is called when exiting the bluetooth production.
	ExitBluetooth(c *BluetoothContext)
}
