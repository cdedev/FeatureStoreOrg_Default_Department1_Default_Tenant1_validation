// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671205162446/Branch.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Branch

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBranchListener is a complete listener for a parse tree produced by BranchParser.
type BaseBranchListener struct{}

var _ BranchListener = &BaseBranchListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBranchListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBranchListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBranchListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBranchListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBranchListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBranchListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBranch is called when production branch is entered.
func (s *BaseBranchListener) EnterBranch(ctx *BranchContext) {}

// ExitBranch is called when production branch is exited.
func (s *BaseBranchListener) ExitBranch(ctx *BranchContext) {}
