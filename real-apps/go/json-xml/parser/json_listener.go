// Code generated from JSON.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // JSON

import "github.com/antlr/antlr4/runtime/Go/antlr"

// JSONListener is a complete listener for a parse tree produced by JSONParser.
type JSONListener interface {
	antlr.ParseTreeListener

	// EnterJson is called when entering the json production.
	EnterJson(c *JsonContext)

	// EnterAnObject is called when entering the AnObject production.
	EnterAnObject(c *AnObjectContext)

	// EnterBlankObject is called when entering the BlankObject production.
	EnterBlankObject(c *BlankObjectContext)

	// EnterArrayOfValues is called when entering the ArrayOfValues production.
	EnterArrayOfValues(c *ArrayOfValuesContext)

	// EnterBlankArray is called when entering the BlankArray production.
	EnterBlankArray(c *BlankArrayContext)

	// EnterStringValue is called when entering the StringValue production.
	EnterStringValue(c *StringValueContext)

	// EnterAtom is called when entering the Atom production.
	EnterAtom(c *AtomContext)

	// EnterObjectValue is called when entering the ObjectValue production.
	EnterObjectValue(c *ObjectValueContext)

	// EnterArrayValue is called when entering the ArrayValue production.
	EnterArrayValue(c *ArrayValueContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// ExitJson is called when exiting the json production.
	ExitJson(c *JsonContext)

	// ExitAnObject is called when exiting the AnObject production.
	ExitAnObject(c *AnObjectContext)

	// ExitBlankObject is called when exiting the BlankObject production.
	ExitBlankObject(c *BlankObjectContext)

	// ExitArrayOfValues is called when exiting the ArrayOfValues production.
	ExitArrayOfValues(c *ArrayOfValuesContext)

	// ExitBlankArray is called when exiting the BlankArray production.
	ExitBlankArray(c *BlankArrayContext)

	// ExitStringValue is called when exiting the StringValue production.
	ExitStringValue(c *StringValueContext)

	// ExitAtom is called when exiting the Atom production.
	ExitAtom(c *AtomContext)

	// ExitObjectValue is called when exiting the ObjectValue production.
	ExitObjectValue(c *ObjectValueContext)

	// ExitArrayValue is called when exiting the ArrayValue production.
	ExitArrayValue(c *ArrayValueContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)
}
