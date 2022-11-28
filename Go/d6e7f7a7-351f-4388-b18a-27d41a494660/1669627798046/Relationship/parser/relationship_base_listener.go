// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669627798046/Relationship.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Relationship

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRelationshipListener is a complete listener for a parse tree produced by RelationshipParser.
type BaseRelationshipListener struct{}

var _ RelationshipListener = &BaseRelationshipListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRelationshipListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRelationshipListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRelationshipListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRelationshipListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRelationshipListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRelationshipListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRelationship is called when production relationship is entered.
func (s *BaseRelationshipListener) EnterRelationship(ctx *RelationshipContext) {}

// ExitRelationship is called when production relationship is exited.
func (s *BaseRelationshipListener) ExitRelationship(ctx *RelationshipContext) {}
