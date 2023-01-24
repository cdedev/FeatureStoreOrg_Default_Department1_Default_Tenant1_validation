// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674580642195/Religiousness.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Religiousness

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseReligiousnessListener is a complete listener for a parse tree produced by ReligiousnessParser.
type BaseReligiousnessListener struct{}

var _ ReligiousnessListener = &BaseReligiousnessListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseReligiousnessListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseReligiousnessListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseReligiousnessListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseReligiousnessListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseReligiousnessListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseReligiousnessListener) ExitExpression(ctx *ExpressionContext) {}

// EnterReligiousness is called when production religiousness is entered.
func (s *BaseReligiousnessListener) EnterReligiousness(ctx *ReligiousnessContext) {}

// ExitReligiousness is called when production religiousness is exited.
func (s *BaseReligiousnessListener) ExitReligiousness(ctx *ReligiousnessContext) {}
