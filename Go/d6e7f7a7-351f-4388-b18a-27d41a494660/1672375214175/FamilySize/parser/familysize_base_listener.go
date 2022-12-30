// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672375214175/FamilySize.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FamilySize

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFamilySizeListener is a complete listener for a parse tree produced by FamilySizeParser.
type BaseFamilySizeListener struct{}

var _ FamilySizeListener = &BaseFamilySizeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFamilySizeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFamilySizeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFamilySizeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFamilySizeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFamilySizeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFamilySizeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFamilysize is called when production familysize is entered.
func (s *BaseFamilySizeListener) EnterFamilysize(ctx *FamilysizeContext) {}

// ExitFamilysize is called when production familysize is exited.
func (s *BaseFamilySizeListener) ExitFamilysize(ctx *FamilysizeContext) {}
