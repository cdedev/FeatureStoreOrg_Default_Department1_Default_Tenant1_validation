// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672284871047/SuggestiveThemes.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SuggestiveThemes

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SuggestiveThemesListener is a complete listener for a parse tree produced by SuggestiveThemesParser.
type SuggestiveThemesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSuggestivethemes is called when entering the suggestivethemes production.
	EnterSuggestivethemes(c *SuggestivethemesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSuggestivethemes is called when exiting the suggestivethemes production.
	ExitSuggestivethemes(c *SuggestivethemesContext)
}
