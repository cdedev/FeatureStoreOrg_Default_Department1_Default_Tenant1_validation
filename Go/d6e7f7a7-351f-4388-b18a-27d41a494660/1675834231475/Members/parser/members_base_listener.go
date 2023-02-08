// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675834231475/Members.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Members

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMembersListener is a complete listener for a parse tree produced by MembersParser.
type BaseMembersListener struct{}

var _ MembersListener = &BaseMembersListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMembersListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMembersListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMembersListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMembersListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMembersListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMembersListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMembers is called when production members is entered.
func (s *BaseMembersListener) EnterMembers(ctx *MembersContext) {}

// ExitMembers is called when production members is exited.
func (s *BaseMembersListener) ExitMembers(ctx *MembersContext) {}
