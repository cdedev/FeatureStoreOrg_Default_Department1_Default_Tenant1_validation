// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671698007742/BirthPlace.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BirthPlace

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBirthPlaceListener is a complete listener for a parse tree produced by BirthPlaceParser.
type BaseBirthPlaceListener struct{}

var _ BirthPlaceListener = &BaseBirthPlaceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBirthPlaceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBirthPlaceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBirthPlaceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBirthPlaceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBirthPlaceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBirthPlaceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBirthplace is called when production birthplace is entered.
func (s *BaseBirthPlaceListener) EnterBirthplace(ctx *BirthplaceContext) {}

// ExitBirthplace is called when production birthplace is exited.
func (s *BaseBirthPlaceListener) ExitBirthplace(ctx *BirthplaceContext) {}
