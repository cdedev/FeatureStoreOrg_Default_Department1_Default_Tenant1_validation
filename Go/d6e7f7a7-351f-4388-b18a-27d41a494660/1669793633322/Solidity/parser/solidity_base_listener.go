// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669793633322/Solidity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Solidity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSolidityListener is a complete listener for a parse tree produced by SolidityParser.
type BaseSolidityListener struct{}

var _ SolidityListener = &BaseSolidityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSolidityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSolidityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSolidityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSolidityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSolidityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSolidityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSolidity is called when production solidity is entered.
func (s *BaseSolidityListener) EnterSolidity(ctx *SolidityContext) {}

// ExitSolidity is called when production solidity is exited.
func (s *BaseSolidityListener) ExitSolidity(ctx *SolidityContext) {}
