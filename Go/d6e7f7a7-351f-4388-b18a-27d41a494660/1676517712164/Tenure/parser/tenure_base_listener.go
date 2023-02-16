// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676517712164/Tenure.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Tenure

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTenureListener is a complete listener for a parse tree produced by TenureParser.
type BaseTenureListener struct{}

var _ TenureListener = &BaseTenureListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTenureListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTenureListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTenureListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTenureListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTenureListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTenureListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTenure is called when production tenure is entered.
func (s *BaseTenureListener) EnterTenure(ctx *TenureContext) {}

// ExitTenure is called when production tenure is exited.
func (s *BaseTenureListener) ExitTenure(ctx *TenureContext) {}
