// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675311460151/Experience.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Experience

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ExperienceListener is a complete listener for a parse tree produced by ExperienceParser.
type ExperienceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExperience is called when entering the experience production.
	EnterExperience(c *ExperienceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExperience is called when exiting the experience production.
	ExitExperience(c *ExperienceContext)
}
