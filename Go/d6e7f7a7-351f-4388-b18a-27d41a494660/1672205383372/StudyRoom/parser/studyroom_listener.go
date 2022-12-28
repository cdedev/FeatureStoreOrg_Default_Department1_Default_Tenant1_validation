// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672205383372/StudyRoom.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // StudyRoom

import "github.com/antlr/antlr4/runtime/Go/antlr"

// StudyRoomListener is a complete listener for a parse tree produced by StudyRoomParser.
type StudyRoomListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterStudyroom is called when entering the studyroom production.
	EnterStudyroom(c *StudyroomContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitStudyroom is called when exiting the studyroom production.
	ExitStudyroom(c *StudyroomContext)
}
