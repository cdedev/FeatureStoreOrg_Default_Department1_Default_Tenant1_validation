// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671690203209/CollisionType.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CollisionType

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCollisionTypeListener is a complete listener for a parse tree produced by CollisionTypeParser.
type BaseCollisionTypeListener struct{}

var _ CollisionTypeListener = &BaseCollisionTypeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCollisionTypeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCollisionTypeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCollisionTypeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCollisionTypeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCollisionTypeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCollisionTypeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCollisiontype is called when production collisiontype is entered.
func (s *BaseCollisionTypeListener) EnterCollisiontype(ctx *CollisiontypeContext) {}

// ExitCollisiontype is called when production collisiontype is exited.
func (s *BaseCollisionTypeListener) ExitCollisiontype(ctx *CollisiontypeContext) {}
