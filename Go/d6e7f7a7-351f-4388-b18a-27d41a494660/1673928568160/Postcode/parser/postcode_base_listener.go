// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673928568160/Postcode.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Postcode

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePostcodeListener is a complete listener for a parse tree produced by PostcodeParser.
type BasePostcodeListener struct{}

var _ PostcodeListener = &BasePostcodeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePostcodeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePostcodeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePostcodeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePostcodeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePostcodeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePostcodeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPostcode is called when production postcode is entered.
func (s *BasePostcodeListener) EnterPostcode(ctx *PostcodeContext) {}

// ExitPostcode is called when production postcode is exited.
func (s *BasePostcodeListener) ExitPostcode(ctx *PostcodeContext) {}
