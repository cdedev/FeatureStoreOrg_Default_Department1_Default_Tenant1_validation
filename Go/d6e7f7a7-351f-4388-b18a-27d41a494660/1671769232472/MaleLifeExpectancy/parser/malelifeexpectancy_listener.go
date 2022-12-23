// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671769232472/MaleLifeExpectancy.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MaleLifeExpectancy

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MaleLifeExpectancyListener is a complete listener for a parse tree produced by MaleLifeExpectancyParser.
type MaleLifeExpectancyListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMalelifeexpectancy is called when entering the malelifeexpectancy production.
	EnterMalelifeexpectancy(c *MalelifeexpectancyContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMalelifeexpectancy is called when exiting the malelifeexpectancy production.
	ExitMalelifeexpectancy(c *MalelifeexpectancyContext)
}
