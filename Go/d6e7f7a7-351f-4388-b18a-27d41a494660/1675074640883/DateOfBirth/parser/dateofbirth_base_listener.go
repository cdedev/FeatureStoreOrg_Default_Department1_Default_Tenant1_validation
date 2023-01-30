// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675074640883/DateOfBirth.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DateOfBirth

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDateOfBirthListener is a complete listener for a parse tree produced by DateOfBirthParser.
type BaseDateOfBirthListener struct{}

var _ DateOfBirthListener = &BaseDateOfBirthListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDateOfBirthListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDateOfBirthListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDateOfBirthListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDateOfBirthListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDateOfBirthListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDateOfBirthListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDateofbirth is called when production dateofbirth is entered.
func (s *BaseDateOfBirthListener) EnterDateofbirth(ctx *DateofbirthContext) {}

// ExitDateofbirth is called when production dateofbirth is exited.
func (s *BaseDateOfBirthListener) ExitDateofbirth(ctx *DateofbirthContext) {}
