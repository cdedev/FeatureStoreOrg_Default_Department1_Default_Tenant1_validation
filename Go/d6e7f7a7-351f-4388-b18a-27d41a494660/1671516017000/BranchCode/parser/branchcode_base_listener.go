// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671516017000/BranchCode.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BranchCode

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBranchCodeListener is a complete listener for a parse tree produced by BranchCodeParser.
type BaseBranchCodeListener struct{}

var _ BranchCodeListener = &BaseBranchCodeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBranchCodeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBranchCodeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBranchCodeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBranchCodeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBranchCodeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBranchCodeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBranchcode is called when production branchcode is entered.
func (s *BaseBranchCodeListener) EnterBranchcode(ctx *BranchcodeContext) {}

// ExitBranchcode is called when production branchcode is exited.
func (s *BaseBranchCodeListener) ExitBranchcode(ctx *BranchcodeContext) {}
