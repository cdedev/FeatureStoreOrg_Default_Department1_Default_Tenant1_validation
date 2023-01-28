// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674881896098/Grocery.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Grocery

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGroceryListener is a complete listener for a parse tree produced by GroceryParser.
type BaseGroceryListener struct{}

var _ GroceryListener = &BaseGroceryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGroceryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGroceryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGroceryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGroceryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGroceryListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGroceryListener) ExitExpression(ctx *ExpressionContext) {}

// EnterGrocery is called when production grocery is entered.
func (s *BaseGroceryListener) EnterGrocery(ctx *GroceryContext) {}

// ExitGrocery is called when production grocery is exited.
func (s *BaseGroceryListener) ExitGrocery(ctx *GroceryContext) {}
