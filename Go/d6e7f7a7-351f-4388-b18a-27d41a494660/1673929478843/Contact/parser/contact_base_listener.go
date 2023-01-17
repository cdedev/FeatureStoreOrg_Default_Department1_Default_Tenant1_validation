// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673929478843/Contact.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Contact

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseContactListener is a complete listener for a parse tree produced by ContactParser.
type BaseContactListener struct{}

var _ ContactListener = &BaseContactListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseContactListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseContactListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseContactListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseContactListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseContactListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseContactListener) ExitExpression(ctx *ExpressionContext) {}

// EnterContact is called when production contact is entered.
func (s *BaseContactListener) EnterContact(ctx *ContactContext) {}

// ExitContact is called when production contact is exited.
func (s *BaseContactListener) ExitContact(ctx *ContactContext) {}
