# Generated from Tite.g4 by ANTLR 4.13.1
# encoding: utf-8
from antlr4 import *
from io import StringIO
import sys
if sys.version_info[1] > 5:
	from typing import TextIO
else:
	from typing.io import TextIO

def serializedATN():
    return [
        4,1,14,68,2,0,7,0,2,1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,
        6,2,7,7,7,1,0,4,0,18,8,0,11,0,12,0,19,1,1,3,1,23,8,1,1,1,1,1,1,1,
        1,1,1,1,1,1,1,1,3,1,32,8,1,1,2,1,2,1,3,1,3,1,3,1,3,1,3,3,3,41,8,
        3,1,4,1,4,1,4,1,5,1,5,3,5,48,8,5,1,6,1,6,1,6,1,6,1,6,1,6,1,7,1,7,
        1,7,1,7,1,7,1,7,1,7,5,7,63,8,7,10,7,12,7,66,9,7,1,7,0,0,8,0,2,4,
        6,8,10,12,14,0,1,1,0,1,3,69,0,17,1,0,0,0,2,22,1,0,0,0,4,33,1,0,0,
        0,6,40,1,0,0,0,8,42,1,0,0,0,10,47,1,0,0,0,12,49,1,0,0,0,14,55,1,
        0,0,0,16,18,3,2,1,0,17,16,1,0,0,0,18,19,1,0,0,0,19,17,1,0,0,0,19,
        20,1,0,0,0,20,1,1,0,0,0,21,23,3,4,2,0,22,21,1,0,0,0,22,23,1,0,0,
        0,23,24,1,0,0,0,24,25,5,9,0,0,25,31,5,14,0,0,26,32,3,6,3,0,27,32,
        3,8,4,0,28,29,3,6,3,0,29,30,3,8,4,0,30,32,1,0,0,0,31,26,1,0,0,0,
        31,27,1,0,0,0,31,28,1,0,0,0,32,3,1,0,0,0,33,34,7,0,0,0,34,5,1,0,
        0,0,35,41,5,9,0,0,36,41,5,10,0,0,37,41,5,11,0,0,38,41,5,12,0,0,39,
        41,3,10,5,0,40,35,1,0,0,0,40,36,1,0,0,0,40,37,1,0,0,0,40,38,1,0,
        0,0,40,39,1,0,0,0,41,7,1,0,0,0,42,43,5,4,0,0,43,44,3,10,5,0,44,9,
        1,0,0,0,45,48,5,9,0,0,46,48,3,12,6,0,47,45,1,0,0,0,47,46,1,0,0,0,
        48,11,1,0,0,0,49,50,5,5,0,0,50,51,3,14,7,0,51,52,5,6,0,0,52,53,5,
        7,0,0,53,54,3,10,5,0,54,13,1,0,0,0,55,56,5,9,0,0,56,57,5,14,0,0,
        57,64,3,6,3,0,58,59,5,8,0,0,59,60,5,9,0,0,60,61,5,14,0,0,61,63,3,
        6,3,0,62,58,1,0,0,0,63,66,1,0,0,0,64,62,1,0,0,0,64,65,1,0,0,0,65,
        15,1,0,0,0,66,64,1,0,0,0,6,19,22,31,40,47,64
    ]

class TiteParser ( Parser ):

    grammarFileName = "Tite.g4"

    atn = ATNDeserializer().deserialize(serializedATN())

    decisionsToDFA = [ DFA(ds, i) for i, ds in enumerate(atn.decisionToState) ]

    sharedContextCache = PredictionContextCache()

    literalNames = [ "<INVALID>", "'$'", "'.'", "'#'", "'='", "'('", "')'", 
                     "'=>'", "','" ]

    symbolicNames = [ "<INVALID>", "<INVALID>", "<INVALID>", "<INVALID>", 
                      "<INVALID>", "<INVALID>", "<INVALID>", "<INVALID>", 
                      "<INVALID>", "IDENTIFIER", "INT", "FLOAT", "STRING", 
                      "WS", "COLON" ]

    RULE_program = 0
    RULE_declaration = 1
    RULE_tag = 2
    RULE_type = 3
    RULE_initializer = 4
    RULE_expression = 5
    RULE_functionLiteral = 6
    RULE_parameterList = 7

    ruleNames =  [ "program", "declaration", "tag", "type", "initializer", 
                   "expression", "functionLiteral", "parameterList" ]

    EOF = Token.EOF
    T__0=1
    T__1=2
    T__2=3
    T__3=4
    T__4=5
    T__5=6
    T__6=7
    T__7=8
    IDENTIFIER=9
    INT=10
    FLOAT=11
    STRING=12
    WS=13
    COLON=14

    def __init__(self, input:TokenStream, output:TextIO = sys.stdout):
        super().__init__(input, output)
        self.checkVersion("4.13.1")
        self._interp = ParserATNSimulator(self, self.atn, self.decisionsToDFA, self.sharedContextCache)
        self._predicates = None




    class ProgramContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def declaration(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(TiteParser.DeclarationContext)
            else:
                return self.getTypedRuleContext(TiteParser.DeclarationContext,i)


        def getRuleIndex(self):
            return TiteParser.RULE_program

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterProgram" ):
                listener.enterProgram(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitProgram" ):
                listener.exitProgram(self)




    def program(self):

        localctx = TiteParser.ProgramContext(self, self._ctx, self.state)
        self.enterRule(localctx, 0, self.RULE_program)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 17 
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            while True:
                self.state = 16
                self.declaration()
                self.state = 19 
                self._errHandler.sync(self)
                _la = self._input.LA(1)
                if not ((((_la) & ~0x3f) == 0 and ((1 << _la) & 526) != 0)):
                    break

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class DeclarationContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def IDENTIFIER(self):
            return self.getToken(TiteParser.IDENTIFIER, 0)

        def COLON(self):
            return self.getToken(TiteParser.COLON, 0)

        def type_(self):
            return self.getTypedRuleContext(TiteParser.TypeContext,0)


        def initializer(self):
            return self.getTypedRuleContext(TiteParser.InitializerContext,0)


        def tag(self):
            return self.getTypedRuleContext(TiteParser.TagContext,0)


        def getRuleIndex(self):
            return TiteParser.RULE_declaration

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterDeclaration" ):
                listener.enterDeclaration(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitDeclaration" ):
                listener.exitDeclaration(self)




    def declaration(self):

        localctx = TiteParser.DeclarationContext(self, self._ctx, self.state)
        self.enterRule(localctx, 2, self.RULE_declaration)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 22
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            if (((_la) & ~0x3f) == 0 and ((1 << _la) & 14) != 0):
                self.state = 21
                self.tag()


            self.state = 24
            self.match(TiteParser.IDENTIFIER)
            self.state = 25
            self.match(TiteParser.COLON)
            self.state = 31
            self._errHandler.sync(self)
            la_ = self._interp.adaptivePredict(self._input,2,self._ctx)
            if la_ == 1:
                self.state = 26
                self.type_()
                pass

            elif la_ == 2:
                self.state = 27
                self.initializer()
                pass

            elif la_ == 3:
                self.state = 28
                self.type_()
                self.state = 29
                self.initializer()
                pass


        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class TagContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser


        def getRuleIndex(self):
            return TiteParser.RULE_tag

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterTag" ):
                listener.enterTag(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitTag" ):
                listener.exitTag(self)




    def tag(self):

        localctx = TiteParser.TagContext(self, self._ctx, self.state)
        self.enterRule(localctx, 4, self.RULE_tag)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 33
            _la = self._input.LA(1)
            if not((((_la) & ~0x3f) == 0 and ((1 << _la) & 14) != 0)):
                self._errHandler.recoverInline(self)
            else:
                self._errHandler.reportMatch(self)
                self.consume()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class TypeContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def IDENTIFIER(self):
            return self.getToken(TiteParser.IDENTIFIER, 0)

        def INT(self):
            return self.getToken(TiteParser.INT, 0)

        def FLOAT(self):
            return self.getToken(TiteParser.FLOAT, 0)

        def STRING(self):
            return self.getToken(TiteParser.STRING, 0)

        def expression(self):
            return self.getTypedRuleContext(TiteParser.ExpressionContext,0)


        def getRuleIndex(self):
            return TiteParser.RULE_type

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterType" ):
                listener.enterType(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitType" ):
                listener.exitType(self)




    def type_(self):

        localctx = TiteParser.TypeContext(self, self._ctx, self.state)
        self.enterRule(localctx, 6, self.RULE_type)
        try:
            self.state = 40
            self._errHandler.sync(self)
            la_ = self._interp.adaptivePredict(self._input,3,self._ctx)
            if la_ == 1:
                self.enterOuterAlt(localctx, 1)
                self.state = 35
                self.match(TiteParser.IDENTIFIER)
                pass

            elif la_ == 2:
                self.enterOuterAlt(localctx, 2)
                self.state = 36
                self.match(TiteParser.INT)
                pass

            elif la_ == 3:
                self.enterOuterAlt(localctx, 3)
                self.state = 37
                self.match(TiteParser.FLOAT)
                pass

            elif la_ == 4:
                self.enterOuterAlt(localctx, 4)
                self.state = 38
                self.match(TiteParser.STRING)
                pass

            elif la_ == 5:
                self.enterOuterAlt(localctx, 5)
                self.state = 39
                self.expression()
                pass


        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class InitializerContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def expression(self):
            return self.getTypedRuleContext(TiteParser.ExpressionContext,0)


        def getRuleIndex(self):
            return TiteParser.RULE_initializer

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterInitializer" ):
                listener.enterInitializer(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitInitializer" ):
                listener.exitInitializer(self)




    def initializer(self):

        localctx = TiteParser.InitializerContext(self, self._ctx, self.state)
        self.enterRule(localctx, 8, self.RULE_initializer)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 42
            self.match(TiteParser.T__3)
            self.state = 43
            self.expression()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class ExpressionContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def IDENTIFIER(self):
            return self.getToken(TiteParser.IDENTIFIER, 0)

        def functionLiteral(self):
            return self.getTypedRuleContext(TiteParser.FunctionLiteralContext,0)


        def getRuleIndex(self):
            return TiteParser.RULE_expression

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterExpression" ):
                listener.enterExpression(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitExpression" ):
                listener.exitExpression(self)




    def expression(self):

        localctx = TiteParser.ExpressionContext(self, self._ctx, self.state)
        self.enterRule(localctx, 10, self.RULE_expression)
        try:
            self.state = 47
            self._errHandler.sync(self)
            token = self._input.LA(1)
            if token in [9]:
                self.enterOuterAlt(localctx, 1)
                self.state = 45
                self.match(TiteParser.IDENTIFIER)
                pass
            elif token in [5]:
                self.enterOuterAlt(localctx, 2)
                self.state = 46
                self.functionLiteral()
                pass
            else:
                raise NoViableAltException(self)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class FunctionLiteralContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def parameterList(self):
            return self.getTypedRuleContext(TiteParser.ParameterListContext,0)


        def expression(self):
            return self.getTypedRuleContext(TiteParser.ExpressionContext,0)


        def getRuleIndex(self):
            return TiteParser.RULE_functionLiteral

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterFunctionLiteral" ):
                listener.enterFunctionLiteral(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitFunctionLiteral" ):
                listener.exitFunctionLiteral(self)




    def functionLiteral(self):

        localctx = TiteParser.FunctionLiteralContext(self, self._ctx, self.state)
        self.enterRule(localctx, 12, self.RULE_functionLiteral)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 49
            self.match(TiteParser.T__4)
            self.state = 50
            self.parameterList()
            self.state = 51
            self.match(TiteParser.T__5)
            self.state = 52
            self.match(TiteParser.T__6)
            self.state = 53
            self.expression()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class ParameterListContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def IDENTIFIER(self, i:int=None):
            if i is None:
                return self.getTokens(TiteParser.IDENTIFIER)
            else:
                return self.getToken(TiteParser.IDENTIFIER, i)

        def COLON(self, i:int=None):
            if i is None:
                return self.getTokens(TiteParser.COLON)
            else:
                return self.getToken(TiteParser.COLON, i)

        def type_(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(TiteParser.TypeContext)
            else:
                return self.getTypedRuleContext(TiteParser.TypeContext,i)


        def getRuleIndex(self):
            return TiteParser.RULE_parameterList

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterParameterList" ):
                listener.enterParameterList(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitParameterList" ):
                listener.exitParameterList(self)




    def parameterList(self):

        localctx = TiteParser.ParameterListContext(self, self._ctx, self.state)
        self.enterRule(localctx, 14, self.RULE_parameterList)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 55
            self.match(TiteParser.IDENTIFIER)
            self.state = 56
            self.match(TiteParser.COLON)
            self.state = 57
            self.type_()
            self.state = 64
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            while _la==8:
                self.state = 58
                self.match(TiteParser.T__7)
                self.state = 59
                self.match(TiteParser.IDENTIFIER)
                self.state = 60
                self.match(TiteParser.COLON)
                self.state = 61
                self.type_()
                self.state = 66
                self._errHandler.sync(self)
                _la = self._input.LA(1)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx





