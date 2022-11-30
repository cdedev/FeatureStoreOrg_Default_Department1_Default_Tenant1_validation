// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669778829094/Contract.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Contract

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseContractListener is a complete listener for a parse tree produced by ContractParser.
type BaseContractListener struct{}

var _ ContractListener = &BaseContractListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseContractListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseContractListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseContractListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseContractListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseContractListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseContractListener) ExitExpression(ctx *ExpressionContext) {}

// EnterContract is called when production contract is entered.
func (s *BaseContractListener) EnterContract(ctx *ContractContext) {}

// ExitContract is called when production contract is exited.
func (s *BaseContractListener) ExitContract(ctx *ContractContext) {}
