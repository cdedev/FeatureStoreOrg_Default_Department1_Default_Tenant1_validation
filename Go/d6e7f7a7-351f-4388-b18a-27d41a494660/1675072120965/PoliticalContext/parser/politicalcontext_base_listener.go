// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675072120965/PoliticalContext.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // PoliticalContext

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePoliticalContextListener is a complete listener for a parse tree produced by PoliticalContextParser.
type BasePoliticalContextListener struct{}

var _ PoliticalContextListener = &BasePoliticalContextListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePoliticalContextListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePoliticalContextListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePoliticalContextListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePoliticalContextListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePoliticalContextListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePoliticalContextListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPoliticalcontext is called when production politicalcontext is entered.
func (s *BasePoliticalContextListener) EnterPoliticalcontext(ctx *PoliticalcontextContext) {}

// ExitPoliticalcontext is called when production politicalcontext is exited.
func (s *BasePoliticalContextListener) ExitPoliticalcontext(ctx *PoliticalcontextContext) {}
