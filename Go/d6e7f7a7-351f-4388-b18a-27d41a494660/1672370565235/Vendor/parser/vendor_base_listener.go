// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672370565235/Vendor.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Vendor

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseVendorListener is a complete listener for a parse tree produced by VendorParser.
type BaseVendorListener struct{}

var _ VendorListener = &BaseVendorListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVendorListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVendorListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVendorListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVendorListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseVendorListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseVendorListener) ExitExpression(ctx *ExpressionContext) {}

// EnterVendor is called when production vendor is entered.
func (s *BaseVendorListener) EnterVendor(ctx *VendorContext) {}

// ExitVendor is called when production vendor is exited.
func (s *BaseVendorListener) ExitVendor(ctx *VendorContext) {}
