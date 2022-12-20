// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671530343405/SubCategory.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SubCategory

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSubCategoryListener is a complete listener for a parse tree produced by SubCategoryParser.
type BaseSubCategoryListener struct{}

var _ SubCategoryListener = &BaseSubCategoryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSubCategoryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSubCategoryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSubCategoryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSubCategoryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSubCategoryListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSubCategoryListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSubcategory is called when production subcategory is entered.
func (s *BaseSubCategoryListener) EnterSubcategory(ctx *SubcategoryContext) {}

// ExitSubcategory is called when production subcategory is exited.
func (s *BaseSubCategoryListener) ExitSubcategory(ctx *SubcategoryContext) {}
