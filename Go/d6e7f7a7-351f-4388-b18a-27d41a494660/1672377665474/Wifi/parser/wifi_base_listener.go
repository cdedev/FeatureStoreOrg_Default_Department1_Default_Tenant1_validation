// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377665474/Wifi.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Wifi

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseWifiListener is a complete listener for a parse tree produced by WifiParser.
type BaseWifiListener struct{}

var _ WifiListener = &BaseWifiListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWifiListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWifiListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWifiListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWifiListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseWifiListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseWifiListener) ExitExpression(ctx *ExpressionContext) {}

// EnterWifi is called when production wifi is entered.
func (s *BaseWifiListener) EnterWifi(ctx *WifiContext) {}

// ExitWifi is called when production wifi is exited.
func (s *BaseWifiListener) ExitWifi(ctx *WifiContext) {}
