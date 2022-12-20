// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671515680829/Patients.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Patients

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PatientsListener is a complete listener for a parse tree produced by PatientsParser.
type PatientsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPatients is called when entering the patients production.
	EnterPatients(c *PatientsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPatients is called when exiting the patients production.
	ExitPatients(c *PatientsContext)
}
