// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676518353306/Geography.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Geography

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGeographyListener is a complete listener for a parse tree produced by GeographyParser.
type BaseGeographyListener struct{}

var _ GeographyListener = &BaseGeographyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGeographyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGeographyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGeographyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGeographyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGeographyListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGeographyListener) ExitExpression(ctx *ExpressionContext) {}

// EnterGeography is called when production geography is entered.
func (s *BaseGeographyListener) EnterGeography(ctx *GeographyContext) {}

// ExitGeography is called when production geography is exited.
func (s *BaseGeographyListener) ExitGeography(ctx *GeographyContext) {}
