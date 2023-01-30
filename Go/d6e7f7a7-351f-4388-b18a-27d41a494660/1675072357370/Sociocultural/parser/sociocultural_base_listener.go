// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675072357370/Sociocultural.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sociocultural

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSocioculturalListener is a complete listener for a parse tree produced by SocioculturalParser.
type BaseSocioculturalListener struct{}

var _ SocioculturalListener = &BaseSocioculturalListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSocioculturalListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSocioculturalListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSocioculturalListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSocioculturalListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSocioculturalListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSocioculturalListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSociocultural is called when production sociocultural is entered.
func (s *BaseSocioculturalListener) EnterSociocultural(ctx *SocioculturalContext) {}

// ExitSociocultural is called when production sociocultural is exited.
func (s *BaseSocioculturalListener) ExitSociocultural(ctx *SocioculturalContext) {}
