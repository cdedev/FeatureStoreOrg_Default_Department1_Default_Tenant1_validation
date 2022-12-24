// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671862557285/StatementOfPurpose.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // StatementOfPurpose

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseStatementOfPurposeListener is a complete listener for a parse tree produced by StatementOfPurposeParser.
type BaseStatementOfPurposeListener struct{}

var _ StatementOfPurposeListener = &BaseStatementOfPurposeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseStatementOfPurposeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseStatementOfPurposeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseStatementOfPurposeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseStatementOfPurposeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseStatementOfPurposeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseStatementOfPurposeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterStatementofpurpose is called when production statementofpurpose is entered.
func (s *BaseStatementOfPurposeListener) EnterStatementofpurpose(ctx *StatementofpurposeContext) {}

// ExitStatementofpurpose is called when production statementofpurpose is exited.
func (s *BaseStatementOfPurposeListener) ExitStatementofpurpose(ctx *StatementofpurposeContext) {}
