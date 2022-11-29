// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669690407274/Protein.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Protein

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseProteinListener is a complete listener for a parse tree produced by ProteinParser.
type BaseProteinListener struct{}

var _ ProteinListener = &BaseProteinListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseProteinListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseProteinListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseProteinListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseProteinListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseProteinListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseProteinListener) ExitExpression(ctx *ExpressionContext) {}

// EnterProtein is called when production protein is entered.
func (s *BaseProteinListener) EnterProtein(ctx *ProteinContext) {}

// ExitProtein is called when production protein is exited.
func (s *BaseProteinListener) ExitProtein(ctx *ProteinContext) {}
