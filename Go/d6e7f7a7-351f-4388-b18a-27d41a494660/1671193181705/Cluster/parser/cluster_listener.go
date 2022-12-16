// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671193181705/Cluster.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cluster

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ClusterListener is a complete listener for a parse tree produced by ClusterParser.
type ClusterListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCluster is called when entering the cluster production.
	EnterCluster(c *ClusterContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCluster is called when exiting the cluster production.
	ExitCluster(c *ClusterContext)
}
