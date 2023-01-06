// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672983153378/EmailerForPromotion.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // EmailerForPromotion

import "github.com/antlr/antlr4/runtime/Go/antlr"

// EmailerForPromotionListener is a complete listener for a parse tree produced by EmailerForPromotionParser.
type EmailerForPromotionListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterEmailerforpromotion is called when entering the emailerforpromotion production.
	EnterEmailerforpromotion(c *EmailerforpromotionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitEmailerforpromotion is called when exiting the emailerforpromotion production.
	ExitEmailerforpromotion(c *EmailerforpromotionContext)
}
