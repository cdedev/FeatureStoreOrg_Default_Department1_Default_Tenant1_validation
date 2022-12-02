// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669981752435/Skills.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Skills

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SkillsListener is a complete listener for a parse tree produced by SkillsParser.
type SkillsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSkills is called when entering the skills production.
	EnterSkills(c *SkillsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSkills is called when exiting the skills production.
	ExitSkills(c *SkillsContext)
}
