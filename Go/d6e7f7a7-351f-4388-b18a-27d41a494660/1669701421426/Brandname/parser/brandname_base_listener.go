// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669701421426/Brandname.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Brandname

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBrandnameListener is a complete listener for a parse tree produced by BrandnameParser.
type BaseBrandnameListener struct{}

var _ BrandnameListener = &BaseBrandnameListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBrandnameListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBrandnameListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBrandnameListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBrandnameListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBrandnameListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBrandnameListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBrandname is called when production brandname is entered.
func (s *BaseBrandnameListener) EnterBrandname(ctx *BrandnameContext) {}

// ExitBrandname is called when production brandname is exited.
func (s *BaseBrandnameListener) ExitBrandname(ctx *BrandnameContext) {}
