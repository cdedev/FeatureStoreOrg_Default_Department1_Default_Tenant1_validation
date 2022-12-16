// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671193181705/Cluster.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cluster

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseClusterListener is a complete listener for a parse tree produced by ClusterParser.
type BaseClusterListener struct{}

var _ ClusterListener = &BaseClusterListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseClusterListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseClusterListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseClusterListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseClusterListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseClusterListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseClusterListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCluster is called when production cluster is entered.
func (s *BaseClusterListener) EnterCluster(ctx *ClusterContext) {}

// ExitCluster is called when production cluster is exited.
func (s *BaseClusterListener) ExitCluster(ctx *ClusterContext) {}
