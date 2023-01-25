// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674621397885/Calories.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Calories

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCaloriesListener is a complete listener for a parse tree produced by CaloriesParser.
type BaseCaloriesListener struct{}

var _ CaloriesListener = &BaseCaloriesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCaloriesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCaloriesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCaloriesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCaloriesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCaloriesListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCaloriesListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCalories is called when production calories is entered.
func (s *BaseCaloriesListener) EnterCalories(ctx *CaloriesContext) {}

// ExitCalories is called when production calories is exited.
func (s *BaseCaloriesListener) ExitCalories(ctx *CaloriesContext) {}
