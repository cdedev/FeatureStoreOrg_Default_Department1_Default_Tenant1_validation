// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675068198075/Cpu.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cpu

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CpuListener is a complete listener for a parse tree produced by CpuParser.
type CpuListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCpu is called when entering the cpu production.
	EnterCpu(c *CpuContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCpu is called when exiting the cpu production.
	ExitCpu(c *CpuContext)
}
