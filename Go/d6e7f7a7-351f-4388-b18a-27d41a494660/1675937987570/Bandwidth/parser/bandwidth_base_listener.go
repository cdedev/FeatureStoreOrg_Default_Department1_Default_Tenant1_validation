// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675937987570/Bandwidth.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Bandwidth

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBandwidthListener is a complete listener for a parse tree produced by BandwidthParser.
type BaseBandwidthListener struct{}

var _ BandwidthListener = &BaseBandwidthListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBandwidthListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBandwidthListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBandwidthListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBandwidthListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBandwidthListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBandwidthListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBandwidth is called when production bandwidth is entered.
func (s *BaseBandwidthListener) EnterBandwidth(ctx *BandwidthContext) {}

// ExitBandwidth is called when production bandwidth is exited.
func (s *BaseBandwidthListener) ExitBandwidth(ctx *BandwidthContext) {}
