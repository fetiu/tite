// Code generated from tite.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // tite

import "github.com/antlr4-go/antlr/v4"

// BasetiteListener is a complete listener for a parse tree produced by titeParser.
type BasetiteListener struct{}

var _ titeListener = &BasetiteListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasetiteListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasetiteListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasetiteListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasetiteListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasetiteListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasetiteListener) ExitProgram(ctx *ProgramContext) {}

// EnterSequence is called when production sequence is entered.
func (s *BasetiteListener) EnterSequence(ctx *SequenceContext) {}

// ExitSequence is called when production sequence is exited.
func (s *BasetiteListener) ExitSequence(ctx *SequenceContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BasetiteListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BasetiteListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterIdentifiers is called when production identifiers is entered.
func (s *BasetiteListener) EnterIdentifiers(ctx *IdentifiersContext) {}

// ExitIdentifiers is called when production identifiers is exited.
func (s *BasetiteListener) ExitIdentifiers(ctx *IdentifiersContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasetiteListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasetiteListener) ExitExpression(ctx *ExpressionContext) {}

// EnterType is called when production type is entered.
func (s *BasetiteListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BasetiteListener) ExitType(ctx *TypeContext) {}

// EnterFunction is called when production function is entered.
func (s *BasetiteListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BasetiteListener) ExitFunction(ctx *FunctionContext) {}

// EnterCondition is called when production condition is entered.
func (s *BasetiteListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BasetiteListener) ExitCondition(ctx *ConditionContext) {}

// EnterOtherwise is called when production otherwise is entered.
func (s *BasetiteListener) EnterOtherwise(ctx *OtherwiseContext) {}

// ExitOtherwise is called when production otherwise is exited.
func (s *BasetiteListener) ExitOtherwise(ctx *OtherwiseContext) {}

// EnterDisjunction is called when production disjunction is entered.
func (s *BasetiteListener) EnterDisjunction(ctx *DisjunctionContext) {}

// ExitDisjunction is called when production disjunction is exited.
func (s *BasetiteListener) ExitDisjunction(ctx *DisjunctionContext) {}

// EnterConjunction is called when production conjunction is entered.
func (s *BasetiteListener) EnterConjunction(ctx *ConjunctionContext) {}

// ExitConjunction is called when production conjunction is exited.
func (s *BasetiteListener) ExitConjunction(ctx *ConjunctionContext) {}

// EnterOr is called when production or is entered.
func (s *BasetiteListener) EnterOr(ctx *OrContext) {}

// ExitOr is called when production or is exited.
func (s *BasetiteListener) ExitOr(ctx *OrContext) {}

// EnterXor is called when production xor is entered.
func (s *BasetiteListener) EnterXor(ctx *XorContext) {}

// ExitXor is called when production xor is exited.
func (s *BasetiteListener) ExitXor(ctx *XorContext) {}

// EnterAnd is called when production and is entered.
func (s *BasetiteListener) EnterAnd(ctx *AndContext) {}

// ExitAnd is called when production and is exited.
func (s *BasetiteListener) ExitAnd(ctx *AndContext) {}

// EnterEquality is called when production equality is entered.
func (s *BasetiteListener) EnterEquality(ctx *EqualityContext) {}

// ExitEquality is called when production equality is exited.
func (s *BasetiteListener) ExitEquality(ctx *EqualityContext) {}

// EnterRelation is called when production relation is entered.
func (s *BasetiteListener) EnterRelation(ctx *RelationContext) {}

// ExitRelation is called when production relation is exited.
func (s *BasetiteListener) ExitRelation(ctx *RelationContext) {}

// EnterShift is called when production shift is entered.
func (s *BasetiteListener) EnterShift(ctx *ShiftContext) {}

// ExitShift is called when production shift is exited.
func (s *BasetiteListener) ExitShift(ctx *ShiftContext) {}

// EnterAdditive is called when production additive is entered.
func (s *BasetiteListener) EnterAdditive(ctx *AdditiveContext) {}

// ExitAdditive is called when production additive is exited.
func (s *BasetiteListener) ExitAdditive(ctx *AdditiveContext) {}

// EnterProduct is called when production product is entered.
func (s *BasetiteListener) EnterProduct(ctx *ProductContext) {}

// ExitProduct is called when production product is exited.
func (s *BasetiteListener) ExitProduct(ctx *ProductContext) {}

// EnterFactor is called when production factor is entered.
func (s *BasetiteListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BasetiteListener) ExitFactor(ctx *FactorContext) {}

// EnterPower is called when production power is entered.
func (s *BasetiteListener) EnterPower(ctx *PowerContext) {}

// ExitPower is called when production power is exited.
func (s *BasetiteListener) ExitPower(ctx *PowerContext) {}

// EnterPostfix is called when production postfix is entered.
func (s *BasetiteListener) EnterPostfix(ctx *PostfixContext) {}

// ExitPostfix is called when production postfix is exited.
func (s *BasetiteListener) ExitPostfix(ctx *PostfixContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BasetiteListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BasetiteListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterTag is called when production tag is entered.
func (s *BasetiteListener) EnterTag(ctx *TagContext) {}

// ExitTag is called when production tag is exited.
func (s *BasetiteListener) ExitTag(ctx *TagContext) {}

// EnterAssign is called when production assign is entered.
func (s *BasetiteListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BasetiteListener) ExitAssign(ctx *AssignContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BasetiteListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BasetiteListener) ExitLiteral(ctx *LiteralContext) {}

// EnterBlock is called when production block is entered.
func (s *BasetiteListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BasetiteListener) ExitBlock(ctx *BlockContext) {}

// EnterArray is called when production array is entered.
func (s *BasetiteListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BasetiteListener) ExitArray(ctx *ArrayContext) {}

// EnterInput is called when production input is entered.
func (s *BasetiteListener) EnterInput(ctx *InputContext) {}

// ExitInput is called when production input is exited.
func (s *BasetiteListener) ExitInput(ctx *InputContext) {}

// EnterDelim is called when production delim is entered.
func (s *BasetiteListener) EnterDelim(ctx *DelimContext) {}

// ExitDelim is called when production delim is exited.
func (s *BasetiteListener) ExitDelim(ctx *DelimContext) {}
