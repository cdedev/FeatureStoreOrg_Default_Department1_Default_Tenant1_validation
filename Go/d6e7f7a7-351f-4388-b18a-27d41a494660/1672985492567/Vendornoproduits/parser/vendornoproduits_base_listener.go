// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672985492567/Vendornoproduits.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Vendornoproduits

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseVendornoproduitsListener is a complete listener for a parse tree produced by VendornoproduitsParser.
type BaseVendornoproduitsListener struct{}

var _ VendornoproduitsListener = &BaseVendornoproduitsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVendornoproduitsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVendornoproduitsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVendornoproduitsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVendornoproduitsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseVendornoproduitsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseVendornoproduitsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterVendornoproduits is called when production vendornoproduits is entered.
func (s *BaseVendornoproduitsListener) EnterVendornoproduits(ctx *VendornoproduitsContext) {}

// ExitVendornoproduits is called when production vendornoproduits is exited.
func (s *BaseVendornoproduitsListener) ExitVendornoproduits(ctx *VendornoproduitsContext) {}
