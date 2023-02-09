// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675920988401/Neighbourhood.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Neighbourhood

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNeighbourhoodListener is a complete listener for a parse tree produced by NeighbourhoodParser.
type BaseNeighbourhoodListener struct{}

var _ NeighbourhoodListener = &BaseNeighbourhoodListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNeighbourhoodListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNeighbourhoodListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNeighbourhoodListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNeighbourhoodListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseNeighbourhoodListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseNeighbourhoodListener) ExitExpression(ctx *ExpressionContext) {}

// EnterNeighbourhood is called when production neighbourhood is entered.
func (s *BaseNeighbourhoodListener) EnterNeighbourhood(ctx *NeighbourhoodContext) {}

// ExitNeighbourhood is called when production neighbourhood is exited.
func (s *BaseNeighbourhoodListener) ExitNeighbourhood(ctx *NeighbourhoodContext) {}
