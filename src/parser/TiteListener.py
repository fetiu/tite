# Generated from tite.g4 by ANTLR 4.13.2
from antlr4 import *
if "." in __name__:
    from .titeParser import titeParser
else:
    from titeParser import titeParser

# This class defines a complete listener for a parse tree produced by titeParser.
class titeListener(ParseTreeListener):

    # Enter a parse tree produced by titeParser#program.
    def enterProgram(self, ctx:titeParser.ProgramContext):
        pass

    # Exit a parse tree produced by titeParser#program.
    def exitProgram(self, ctx:titeParser.ProgramContext):
        pass


    # Enter a parse tree produced by titeParser#sequence.
    def enterSequence(self, ctx:titeParser.SequenceContext):
        pass

    # Exit a parse tree produced by titeParser#sequence.
    def exitSequence(self, ctx:titeParser.SequenceContext):
        pass


    # Enter a parse tree produced by titeParser#declaration.
    def enterDeclaration(self, ctx:titeParser.DeclarationContext):
        pass

    # Exit a parse tree produced by titeParser#declaration.
    def exitDeclaration(self, ctx:titeParser.DeclarationContext):
        pass


    # Enter a parse tree produced by titeParser#identifiers.
    def enterIdentifiers(self, ctx:titeParser.IdentifiersContext):
        pass

    # Exit a parse tree produced by titeParser#identifiers.
    def exitIdentifiers(self, ctx:titeParser.IdentifiersContext):
        pass


    # Enter a parse tree produced by titeParser#expression.
    def enterExpression(self, ctx:titeParser.ExpressionContext):
        pass

    # Exit a parse tree produced by titeParser#expression.
    def exitExpression(self, ctx:titeParser.ExpressionContext):
        pass


    # Enter a parse tree produced by titeParser#type.
    def enterType(self, ctx:titeParser.TypeContext):
        pass

    # Exit a parse tree produced by titeParser#type.
    def exitType(self, ctx:titeParser.TypeContext):
        pass


    # Enter a parse tree produced by titeParser#function.
    def enterFunction(self, ctx:titeParser.FunctionContext):
        pass

    # Exit a parse tree produced by titeParser#function.
    def exitFunction(self, ctx:titeParser.FunctionContext):
        pass


    # Enter a parse tree produced by titeParser#condition.
    def enterCondition(self, ctx:titeParser.ConditionContext):
        pass

    # Exit a parse tree produced by titeParser#condition.
    def exitCondition(self, ctx:titeParser.ConditionContext):
        pass


    # Enter a parse tree produced by titeParser#otherwise.
    def enterOtherwise(self, ctx:titeParser.OtherwiseContext):
        pass

    # Exit a parse tree produced by titeParser#otherwise.
    def exitOtherwise(self, ctx:titeParser.OtherwiseContext):
        pass


    # Enter a parse tree produced by titeParser#disjunction.
    def enterDisjunction(self, ctx:titeParser.DisjunctionContext):
        pass

    # Exit a parse tree produced by titeParser#disjunction.
    def exitDisjunction(self, ctx:titeParser.DisjunctionContext):
        pass


    # Enter a parse tree produced by titeParser#conjunction.
    def enterConjunction(self, ctx:titeParser.ConjunctionContext):
        pass

    # Exit a parse tree produced by titeParser#conjunction.
    def exitConjunction(self, ctx:titeParser.ConjunctionContext):
        pass


    # Enter a parse tree produced by titeParser#or.
    def enterOr(self, ctx:titeParser.OrContext):
        pass

    # Exit a parse tree produced by titeParser#or.
    def exitOr(self, ctx:titeParser.OrContext):
        pass


    # Enter a parse tree produced by titeParser#xor.
    def enterXor(self, ctx:titeParser.XorContext):
        pass

    # Exit a parse tree produced by titeParser#xor.
    def exitXor(self, ctx:titeParser.XorContext):
        pass


    # Enter a parse tree produced by titeParser#and.
    def enterAnd(self, ctx:titeParser.AndContext):
        pass

    # Exit a parse tree produced by titeParser#and.
    def exitAnd(self, ctx:titeParser.AndContext):
        pass


    # Enter a parse tree produced by titeParser#equality.
    def enterEquality(self, ctx:titeParser.EqualityContext):
        pass

    # Exit a parse tree produced by titeParser#equality.
    def exitEquality(self, ctx:titeParser.EqualityContext):
        pass


    # Enter a parse tree produced by titeParser#relation.
    def enterRelation(self, ctx:titeParser.RelationContext):
        pass

    # Exit a parse tree produced by titeParser#relation.
    def exitRelation(self, ctx:titeParser.RelationContext):
        pass


    # Enter a parse tree produced by titeParser#shift.
    def enterShift(self, ctx:titeParser.ShiftContext):
        pass

    # Exit a parse tree produced by titeParser#shift.
    def exitShift(self, ctx:titeParser.ShiftContext):
        pass


    # Enter a parse tree produced by titeParser#additive.
    def enterAdditive(self, ctx:titeParser.AdditiveContext):
        pass

    # Exit a parse tree produced by titeParser#additive.
    def exitAdditive(self, ctx:titeParser.AdditiveContext):
        pass


    # Enter a parse tree produced by titeParser#product.
    def enterProduct(self, ctx:titeParser.ProductContext):
        pass

    # Exit a parse tree produced by titeParser#product.
    def exitProduct(self, ctx:titeParser.ProductContext):
        pass


    # Enter a parse tree produced by titeParser#factor.
    def enterFactor(self, ctx:titeParser.FactorContext):
        pass

    # Exit a parse tree produced by titeParser#factor.
    def exitFactor(self, ctx:titeParser.FactorContext):
        pass


    # Enter a parse tree produced by titeParser#power.
    def enterPower(self, ctx:titeParser.PowerContext):
        pass

    # Exit a parse tree produced by titeParser#power.
    def exitPower(self, ctx:titeParser.PowerContext):
        pass


    # Enter a parse tree produced by titeParser#postfix.
    def enterPostfix(self, ctx:titeParser.PostfixContext):
        pass

    # Exit a parse tree produced by titeParser#postfix.
    def exitPostfix(self, ctx:titeParser.PostfixContext):
        pass


    # Enter a parse tree produced by titeParser#primary.
    def enterPrimary(self, ctx:titeParser.PrimaryContext):
        pass

    # Exit a parse tree produced by titeParser#primary.
    def exitPrimary(self, ctx:titeParser.PrimaryContext):
        pass


    # Enter a parse tree produced by titeParser#tag.
    def enterTag(self, ctx:titeParser.TagContext):
        pass

    # Exit a parse tree produced by titeParser#tag.
    def exitTag(self, ctx:titeParser.TagContext):
        pass


    # Enter a parse tree produced by titeParser#assign.
    def enterAssign(self, ctx:titeParser.AssignContext):
        pass

    # Exit a parse tree produced by titeParser#assign.
    def exitAssign(self, ctx:titeParser.AssignContext):
        pass


    # Enter a parse tree produced by titeParser#literal.
    def enterLiteral(self, ctx:titeParser.LiteralContext):
        pass

    # Exit a parse tree produced by titeParser#literal.
    def exitLiteral(self, ctx:titeParser.LiteralContext):
        pass


    # Enter a parse tree produced by titeParser#block.
    def enterBlock(self, ctx:titeParser.BlockContext):
        pass

    # Exit a parse tree produced by titeParser#block.
    def exitBlock(self, ctx:titeParser.BlockContext):
        pass


    # Enter a parse tree produced by titeParser#array.
    def enterArray(self, ctx:titeParser.ArrayContext):
        pass

    # Exit a parse tree produced by titeParser#array.
    def exitArray(self, ctx:titeParser.ArrayContext):
        pass


    # Enter a parse tree produced by titeParser#input.
    def enterInput(self, ctx:titeParser.InputContext):
        pass

    # Exit a parse tree produced by titeParser#input.
    def exitInput(self, ctx:titeParser.InputContext):
        pass


    # Enter a parse tree produced by titeParser#delim.
    def enterDelim(self, ctx:titeParser.DelimContext):
        pass

    # Exit a parse tree produced by titeParser#delim.
    def exitDelim(self, ctx:titeParser.DelimContext):
        pass



del titeParser