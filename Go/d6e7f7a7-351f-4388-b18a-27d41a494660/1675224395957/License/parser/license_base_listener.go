// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675224395957/License.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // License

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLicenseListener is a complete listener for a parse tree produced by LicenseParser.
type BaseLicenseListener struct{}

var _ LicenseListener = &BaseLicenseListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLicenseListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLicenseListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLicenseListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLicenseListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseLicenseListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseLicenseListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLicense is called when production license is entered.
func (s *BaseLicenseListener) EnterLicense(ctx *LicenseContext) {}

// ExitLicense is called when production license is exited.
func (s *BaseLicenseListener) ExitLicense(ctx *LicenseContext) {}
