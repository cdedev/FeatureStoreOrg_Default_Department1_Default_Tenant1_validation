// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669623376814/Latitude.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Latitude

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLatitudeListener is a complete listener for a parse tree produced by LatitudeParser.
type BaseLatitudeListener struct{}

var _ LatitudeListener = &BaseLatitudeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLatitudeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLatitudeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLatitudeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLatitudeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseLatitudeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseLatitudeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLatitude is called when production latitude is entered.
func (s *BaseLatitudeListener) EnterLatitude(ctx *LatitudeContext) {}

// ExitLatitude is called when production latitude is exited.
func (s *BaseLatitudeListener) ExitLatitude(ctx *LatitudeContext) {}
