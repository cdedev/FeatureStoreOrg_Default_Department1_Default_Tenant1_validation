// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672985218491/VendorCode.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // VendorCode

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseVendorCodeListener is a complete listener for a parse tree produced by VendorCodeParser.
type BaseVendorCodeListener struct{}

var _ VendorCodeListener = &BaseVendorCodeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVendorCodeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVendorCodeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVendorCodeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVendorCodeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseVendorCodeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseVendorCodeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterVendorcode is called when production vendorcode is entered.
func (s *BaseVendorCodeListener) EnterVendorcode(ctx *VendorcodeContext) {}

// ExitVendorcode is called when production vendorcode is exited.
func (s *BaseVendorCodeListener) ExitVendorcode(ctx *VendorcodeContext) {}
