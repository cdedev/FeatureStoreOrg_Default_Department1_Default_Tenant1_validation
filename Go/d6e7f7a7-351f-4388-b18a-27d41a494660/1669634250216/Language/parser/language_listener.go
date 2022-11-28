// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669634250216/Language.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Language

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LanguageListener is a complete listener for a parse tree produced by LanguageParser.
type LanguageListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLanguage is called when entering the language production.
	EnterLanguage(c *LanguageContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLanguage is called when exiting the language production.
	ExitLanguage(c *LanguageContext)
}
