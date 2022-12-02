// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669983532769/EngineSize.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // EngineSize

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseEngineSizeListener is a complete listener for a parse tree produced by EngineSizeParser.
type BaseEngineSizeListener struct{}

var _ EngineSizeListener = &BaseEngineSizeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEngineSizeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEngineSizeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEngineSizeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEngineSizeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseEngineSizeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseEngineSizeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterEnginesize is called when production enginesize is entered.
func (s *BaseEngineSizeListener) EnterEnginesize(ctx *EnginesizeContext) {}

// ExitEnginesize is called when production enginesize is exited.
func (s *BaseEngineSizeListener) ExitEnginesize(ctx *EnginesizeContext) {}
