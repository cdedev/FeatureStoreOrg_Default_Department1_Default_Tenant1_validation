// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672973369964/Contractions.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Contractions

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseContractionsListener is a complete listener for a parse tree produced by ContractionsParser.
type BaseContractionsListener struct{}

var _ ContractionsListener = &BaseContractionsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseContractionsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseContractionsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseContractionsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseContractionsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseContractionsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseContractionsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterContractions is called when production contractions is entered.
func (s *BaseContractionsListener) EnterContractions(ctx *ContractionsContext) {}

// ExitContractions is called when production contractions is exited.
func (s *BaseContractionsListener) ExitContractions(ctx *ContractionsContext) {}
