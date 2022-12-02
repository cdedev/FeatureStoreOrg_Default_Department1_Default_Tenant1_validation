// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669983107242/Package.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Package

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePackageListener is a complete listener for a parse tree produced by PackageParser.
type BasePackageListener struct{}

var _ PackageListener = &BasePackageListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePackageListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePackageListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePackageListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePackageListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePackageListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePackageListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPackage is called when production package is entered.
func (s *BasePackageListener) EnterPackage(ctx *PackageContext) {}

// ExitPackage is called when production package is exited.
func (s *BasePackageListener) ExitPackage(ctx *PackageContext) {}
