// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676260183645/IpAddressPresent.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // IpAddressPresent

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseIpAddressPresentListener is a complete listener for a parse tree produced by IpAddressPresentParser.
type BaseIpAddressPresentListener struct{}

var _ IpAddressPresentListener = &BaseIpAddressPresentListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseIpAddressPresentListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseIpAddressPresentListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseIpAddressPresentListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseIpAddressPresentListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseIpAddressPresentListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseIpAddressPresentListener) ExitExpression(ctx *ExpressionContext) {}

// EnterIpaddresspresent is called when production ipaddresspresent is entered.
func (s *BaseIpAddressPresentListener) EnterIpaddresspresent(ctx *IpaddresspresentContext) {}

// ExitIpaddresspresent is called when production ipaddresspresent is exited.
func (s *BaseIpAddressPresentListener) ExitIpaddresspresent(ctx *IpaddresspresentContext) {}
