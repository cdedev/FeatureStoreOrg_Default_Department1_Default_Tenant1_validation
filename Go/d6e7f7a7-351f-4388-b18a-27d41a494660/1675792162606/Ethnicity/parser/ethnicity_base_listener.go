// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675792162606/Ethnicity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Ethnicity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseEthnicityListener is a complete listener for a parse tree produced by EthnicityParser.
type BaseEthnicityListener struct{}

var _ EthnicityListener = &BaseEthnicityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEthnicityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEthnicityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEthnicityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEthnicityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseEthnicityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseEthnicityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterEthnicity is called when production ethnicity is entered.
func (s *BaseEthnicityListener) EnterEthnicity(ctx *EthnicityContext) {}

// ExitEthnicity is called when production ethnicity is exited.
func (s *BaseEthnicityListener) ExitEthnicity(ctx *EthnicityContext) {}
