// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669707448675/Address.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Address

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAddressListener is a complete listener for a parse tree produced by AddressParser.
type BaseAddressListener struct{}

var _ AddressListener = &BaseAddressListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAddressListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAddressListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAddressListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAddressListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAddressListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAddressListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAddress is called when production address is entered.
func (s *BaseAddressListener) EnterAddress(ctx *AddressContext) {}

// ExitAddress is called when production address is exited.
func (s *BaseAddressListener) ExitAddress(ctx *AddressContext) {}
