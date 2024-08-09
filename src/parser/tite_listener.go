// Code generated from tite.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // tite

import "github.com/antlr4-go/antlr/v4"

// titeListener is a complete listener for a parse tree produced by titeParser.
type titeListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterSequence is called when entering the sequence production.
	EnterSequence(c *SequenceContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterIdentifiers is called when entering the identifiers production.
	EnterIdentifiers(c *IdentifiersContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterOtherwise is called when entering the otherwise production.
	EnterOtherwise(c *OtherwiseContext)

	// EnterDisjunction is called when entering the disjunction production.
	EnterDisjunction(c *DisjunctionContext)

	// EnterConjunction is called when entering the conjunction production.
	EnterConjunction(c *ConjunctionContext)

	// EnterOr is called when entering the or production.
	EnterOr(c *OrContext)

	// EnterXor is called when entering the xor production.
	EnterXor(c *XorContext)

	// EnterAnd is called when entering the and production.
	EnterAnd(c *AndContext)

	// EnterEquality is called when entering the equality production.
	EnterEquality(c *EqualityContext)

	// EnterRelation is called when entering the relation production.
	EnterRelation(c *RelationContext)

	// EnterShift is called when entering the shift production.
	EnterShift(c *ShiftContext)

	// EnterAdditive is called when entering the additive production.
	EnterAdditive(c *AdditiveContext)

	// EnterProduct is called when entering the product production.
	EnterProduct(c *ProductContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterPower is called when entering the power production.
	EnterPower(c *PowerContext)

	// EnterPostfix is called when entering the postfix production.
	EnterPostfix(c *PostfixContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterTag is called when entering the tag production.
	EnterTag(c *TagContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterInput is called when entering the input production.
	EnterInput(c *InputContext)

	// EnterDelim is called when entering the delim production.
	EnterDelim(c *DelimContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitSequence is called when exiting the sequence production.
	ExitSequence(c *SequenceContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitIdentifiers is called when exiting the identifiers production.
	ExitIdentifiers(c *IdentifiersContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitOtherwise is called when exiting the otherwise production.
	ExitOtherwise(c *OtherwiseContext)

	// ExitDisjunction is called when exiting the disjunction production.
	ExitDisjunction(c *DisjunctionContext)

	// ExitConjunction is called when exiting the conjunction production.
	ExitConjunction(c *ConjunctionContext)

	// ExitOr is called when exiting the or production.
	ExitOr(c *OrContext)

	// ExitXor is called when exiting the xor production.
	ExitXor(c *XorContext)

	// ExitAnd is called when exiting the and production.
	ExitAnd(c *AndContext)

	// ExitEquality is called when exiting the equality production.
	ExitEquality(c *EqualityContext)

	// ExitRelation is called when exiting the relation production.
	ExitRelation(c *RelationContext)

	// ExitShift is called when exiting the shift production.
	ExitShift(c *ShiftContext)

	// ExitAdditive is called when exiting the additive production.
	ExitAdditive(c *AdditiveContext)

	// ExitProduct is called when exiting the product production.
	ExitProduct(c *ProductContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitPower is called when exiting the power production.
	ExitPower(c *PowerContext)

	// ExitPostfix is called when exiting the postfix production.
	ExitPostfix(c *PostfixContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitTag is called when exiting the tag production.
	ExitTag(c *TagContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitInput is called when exiting the input production.
	ExitInput(c *InputContext)

	// ExitDelim is called when exiting the delim production.
	ExitDelim(c *DelimContext)
}
