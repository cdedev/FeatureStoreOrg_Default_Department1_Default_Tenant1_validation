// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671514143700/PostalCode.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // PostalCode

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePostalCodeListener is a complete listener for a parse tree produced by PostalCodeParser.
type BasePostalCodeListener struct{}

var _ PostalCodeListener = &BasePostalCodeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePostalCodeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePostalCodeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePostalCodeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePostalCodeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePostalCodeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePostalCodeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPostalcode is called when production postalcode is entered.
func (s *BasePostalCodeListener) EnterPostalcode(ctx *PostalcodeContext) {}

// ExitPostalcode is called when production postalcode is exited.
func (s *BasePostalCodeListener) ExitPostalcode(ctx *PostalcodeContext) {}
