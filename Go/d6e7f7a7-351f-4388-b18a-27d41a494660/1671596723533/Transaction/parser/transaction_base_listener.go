// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671596723533/Transaction.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Transaction

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTransactionListener is a complete listener for a parse tree produced by TransactionParser.
type BaseTransactionListener struct{}

var _ TransactionListener = &BaseTransactionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTransactionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTransactionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTransactionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTransactionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTransactionListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTransactionListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTransaction is called when production transaction is entered.
func (s *BaseTransactionListener) EnterTransaction(ctx *TransactionContext) {}

// ExitTransaction is called when production transaction is exited.
func (s *BaseTransactionListener) ExitTransaction(ctx *TransactionContext) {}
