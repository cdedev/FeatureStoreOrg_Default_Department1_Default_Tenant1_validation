// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669692348781/Family.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Family

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFamilyListener is a complete listener for a parse tree produced by FamilyParser.
type BaseFamilyListener struct{}

var _ FamilyListener = &BaseFamilyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFamilyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFamilyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFamilyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFamilyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFamilyListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFamilyListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFamily is called when production family is entered.
func (s *BaseFamilyListener) EnterFamily(ctx *FamilyContext) {}

// ExitFamily is called when production family is exited.
func (s *BaseFamilyListener) ExitFamily(ctx *FamilyContext) {}
