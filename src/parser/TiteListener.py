# Generated from Tite.g4 by ANTLR 4.13.1
from antlr4 import *
if "." in __name__:
    from .TiteParser import TiteParser
else:
    from TiteParser import TiteParser

# This class defines a complete listener for a parse tree produced by TiteParser.
class TiteListener(ParseTreeListener):

    # Enter a parse tree produced by TiteParser#program.
    def enterProgram(self, ctx:TiteParser.ProgramContext):
        pass

    # Exit a parse tree produced by TiteParser#program.
    def exitProgram(self, ctx:TiteParser.ProgramContext):
        pass


    # Enter a parse tree produced by TiteParser#declaration.
    def enterDeclaration(self, ctx:TiteParser.DeclarationContext):
        pass

    # Exit a parse tree produced by TiteParser#declaration.
    def exitDeclaration(self, ctx:TiteParser.DeclarationContext):
        pass


    # Enter a parse tree produced by TiteParser#tag.
    def enterTag(self, ctx:TiteParser.TagContext):
        pass

    # Exit a parse tree produced by TiteParser#tag.
    def exitTag(self, ctx:TiteParser.TagContext):
        pass


    # Enter a parse tree produced by TiteParser#type.
    def enterType(self, ctx:TiteParser.TypeContext):
        pass

    # Exit a parse tree produced by TiteParser#type.
    def exitType(self, ctx:TiteParser.TypeContext):
        pass


    # Enter a parse tree produced by TiteParser#initializer.
    def enterInitializer(self, ctx:TiteParser.InitializerContext):
        pass

    # Exit a parse tree produced by TiteParser#initializer.
    def exitInitializer(self, ctx:TiteParser.InitializerContext):
        pass


    # Enter a parse tree produced by TiteParser#expression.
    def enterExpression(self, ctx:TiteParser.ExpressionContext):
        pass

    # Exit a parse tree produced by TiteParser#expression.
    def exitExpression(self, ctx:TiteParser.ExpressionContext):
        pass


    # Enter a parse tree produced by TiteParser#functionLiteral.
    def enterFunctionLiteral(self, ctx:TiteParser.FunctionLiteralContext):
        pass

    # Exit a parse tree produced by TiteParser#functionLiteral.
    def exitFunctionLiteral(self, ctx:TiteParser.FunctionLiteralContext):
        pass


    # Enter a parse tree produced by TiteParser#parameterList.
    def enterParameterList(self, ctx:TiteParser.ParameterListContext):
        pass

    # Exit a parse tree produced by TiteParser#parameterList.
    def exitParameterList(self, ctx:TiteParser.ParameterListContext):
        pass



del TiteParser