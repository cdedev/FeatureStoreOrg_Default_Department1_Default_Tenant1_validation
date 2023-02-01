// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675241758896/Bio.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Bio

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBioListener is a complete listener for a parse tree produced by BioParser.
type BaseBioListener struct{}

var _ BioListener = &BaseBioListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBioListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBioListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBioListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBioListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBioListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBioListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBio is called when production bio is entered.
func (s *BaseBioListener) EnterBio(ctx *BioContext) {}

// ExitBio is called when production bio is exited.
func (s *BaseBioListener) ExitBio(ctx *BioContext) {}
