// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674454745948/Category.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Category

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCategoryListener is a complete listener for a parse tree produced by CategoryParser.
type BaseCategoryListener struct{}

var _ CategoryListener = &BaseCategoryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCategoryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCategoryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCategoryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCategoryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCategoryListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCategoryListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCategory is called when production category is entered.
func (s *BaseCategoryListener) EnterCategory(ctx *CategoryContext) {}

// ExitCategory is called when production category is exited.
func (s *BaseCategoryListener) ExitCategory(ctx *CategoryContext) {}
