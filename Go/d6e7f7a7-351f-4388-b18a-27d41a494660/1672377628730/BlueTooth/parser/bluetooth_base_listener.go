// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377628730/BlueTooth.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BlueTooth

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBlueToothListener is a complete listener for a parse tree produced by BlueToothParser.
type BaseBlueToothListener struct{}

var _ BlueToothListener = &BaseBlueToothListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBlueToothListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBlueToothListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBlueToothListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBlueToothListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBlueToothListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBlueToothListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBluetooth is called when production bluetooth is entered.
func (s *BaseBlueToothListener) EnterBluetooth(ctx *BluetoothContext) {}

// ExitBluetooth is called when production bluetooth is exited.
func (s *BaseBlueToothListener) ExitBluetooth(ctx *BluetoothContext) {}
