// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674787464947/Barcode.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Barcode

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBarcodeListener is a complete listener for a parse tree produced by BarcodeParser.
type BaseBarcodeListener struct{}

var _ BarcodeListener = &BaseBarcodeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBarcodeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBarcodeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBarcodeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBarcodeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBarcodeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBarcodeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBarcode is called when production barcode is entered.
func (s *BaseBarcodeListener) EnterBarcode(ctx *BarcodeContext) {}

// ExitBarcode is called when production barcode is exited.
func (s *BaseBarcodeListener) ExitBarcode(ctx *BarcodeContext) {}
