// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672376894048/Airbags.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Airbags

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAirbagsListener is a complete listener for a parse tree produced by AirbagsParser.
type BaseAirbagsListener struct{}

var _ AirbagsListener = &BaseAirbagsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAirbagsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAirbagsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAirbagsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAirbagsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAirbagsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAirbagsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAirbags is called when production airbags is entered.
func (s *BaseAirbagsListener) EnterAirbags(ctx *AirbagsContext) {}

// ExitAirbags is called when production airbags is exited.
func (s *BaseAirbagsListener) ExitAirbags(ctx *AirbagsContext) {}
