// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377534929/FrontCamera.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FrontCamera

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFrontCameraListener is a complete listener for a parse tree produced by FrontCameraParser.
type BaseFrontCameraListener struct{}

var _ FrontCameraListener = &BaseFrontCameraListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFrontCameraListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFrontCameraListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFrontCameraListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFrontCameraListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFrontCameraListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFrontCameraListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFrontcamera is called when production frontcamera is entered.
func (s *BaseFrontCameraListener) EnterFrontcamera(ctx *FrontcameraContext) {}

// ExitFrontcamera is called when production frontcamera is exited.
func (s *BaseFrontCameraListener) ExitFrontcamera(ctx *FrontcameraContext) {}
