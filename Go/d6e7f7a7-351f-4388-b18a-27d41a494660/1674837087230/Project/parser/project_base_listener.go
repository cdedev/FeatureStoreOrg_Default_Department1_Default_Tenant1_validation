// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674837087230/Project.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Project

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseProjectListener is a complete listener for a parse tree produced by ProjectParser.
type BaseProjectListener struct{}

var _ ProjectListener = &BaseProjectListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseProjectListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseProjectListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseProjectListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseProjectListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseProjectListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseProjectListener) ExitExpression(ctx *ExpressionContext) {}

// EnterProject is called when production project is entered.
func (s *BaseProjectListener) EnterProject(ctx *ProjectContext) {}

// ExitProject is called when production project is exited.
func (s *BaseProjectListener) ExitProject(ctx *ProjectContext) {}
