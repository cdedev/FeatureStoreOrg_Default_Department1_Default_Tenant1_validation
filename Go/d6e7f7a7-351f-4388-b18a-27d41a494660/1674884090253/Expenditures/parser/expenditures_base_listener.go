// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674884090253/Expenditures.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Expenditures

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseExpendituresListener is a complete listener for a parse tree produced by ExpendituresParser.
type BaseExpendituresListener struct{}

var _ ExpendituresListener = &BaseExpendituresListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseExpendituresListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseExpendituresListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseExpendituresListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseExpendituresListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseExpendituresListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseExpendituresListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExpenditures is called when production expenditures is entered.
func (s *BaseExpendituresListener) EnterExpenditures(ctx *ExpendituresContext) {}

// ExitExpenditures is called when production expenditures is exited.
func (s *BaseExpendituresListener) ExitExpenditures(ctx *ExpendituresContext) {}
