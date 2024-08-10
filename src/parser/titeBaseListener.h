
// Generated from tite.g4 by ANTLR 4.13.2

#pragma once


#include "antlr4-runtime.h"
#include "titeListener.h"


/**
 * This class provides an empty implementation of titeListener,
 * which can be extended to create a listener which only needs to handle a subset
 * of the available methods.
 */
class  titeBaseListener : public titeListener {
public:

  virtual void enterProgram(titeParser::ProgramContext * /*ctx*/) override { }
  virtual void exitProgram(titeParser::ProgramContext * /*ctx*/) override { }

  virtual void enterSequence(titeParser::SequenceContext * /*ctx*/) override { }
  virtual void exitSequence(titeParser::SequenceContext * /*ctx*/) override { }

  virtual void enterDeclaration(titeParser::DeclarationContext * /*ctx*/) override { }
  virtual void exitDeclaration(titeParser::DeclarationContext * /*ctx*/) override { }

  virtual void enterIdentifiers(titeParser::IdentifiersContext * /*ctx*/) override { }
  virtual void exitIdentifiers(titeParser::IdentifiersContext * /*ctx*/) override { }

  virtual void enterExpression(titeParser::ExpressionContext * /*ctx*/) override { }
  virtual void exitExpression(titeParser::ExpressionContext * /*ctx*/) override { }

  virtual void enterType(titeParser::TypeContext * /*ctx*/) override { }
  virtual void exitType(titeParser::TypeContext * /*ctx*/) override { }

  virtual void enterFunction(titeParser::FunctionContext * /*ctx*/) override { }
  virtual void exitFunction(titeParser::FunctionContext * /*ctx*/) override { }

  virtual void enterCondition(titeParser::ConditionContext * /*ctx*/) override { }
  virtual void exitCondition(titeParser::ConditionContext * /*ctx*/) override { }

  virtual void enterOtherwise(titeParser::OtherwiseContext * /*ctx*/) override { }
  virtual void exitOtherwise(titeParser::OtherwiseContext * /*ctx*/) override { }

  virtual void enterDisjunction(titeParser::DisjunctionContext * /*ctx*/) override { }
  virtual void exitDisjunction(titeParser::DisjunctionContext * /*ctx*/) override { }

  virtual void enterConjunction(titeParser::ConjunctionContext * /*ctx*/) override { }
  virtual void exitConjunction(titeParser::ConjunctionContext * /*ctx*/) override { }

  virtual void enterOr(titeParser::OrContext * /*ctx*/) override { }
  virtual void exitOr(titeParser::OrContext * /*ctx*/) override { }

  virtual void enterXor(titeParser::XorContext * /*ctx*/) override { }
  virtual void exitXor(titeParser::XorContext * /*ctx*/) override { }

  virtual void enterAnd(titeParser::AndContext * /*ctx*/) override { }
  virtual void exitAnd(titeParser::AndContext * /*ctx*/) override { }

  virtual void enterEquality(titeParser::EqualityContext * /*ctx*/) override { }
  virtual void exitEquality(titeParser::EqualityContext * /*ctx*/) override { }

  virtual void enterRelation(titeParser::RelationContext * /*ctx*/) override { }
  virtual void exitRelation(titeParser::RelationContext * /*ctx*/) override { }

  virtual void enterShift(titeParser::ShiftContext * /*ctx*/) override { }
  virtual void exitShift(titeParser::ShiftContext * /*ctx*/) override { }

  virtual void enterAdditive(titeParser::AdditiveContext * /*ctx*/) override { }
  virtual void exitAdditive(titeParser::AdditiveContext * /*ctx*/) override { }

  virtual void enterProduct(titeParser::ProductContext * /*ctx*/) override { }
  virtual void exitProduct(titeParser::ProductContext * /*ctx*/) override { }

  virtual void enterFactor(titeParser::FactorContext * /*ctx*/) override { }
  virtual void exitFactor(titeParser::FactorContext * /*ctx*/) override { }

  virtual void enterPower(titeParser::PowerContext * /*ctx*/) override { }
  virtual void exitPower(titeParser::PowerContext * /*ctx*/) override { }

  virtual void enterPostfix(titeParser::PostfixContext * /*ctx*/) override { }
  virtual void exitPostfix(titeParser::PostfixContext * /*ctx*/) override { }

  virtual void enterPrimary(titeParser::PrimaryContext * /*ctx*/) override { }
  virtual void exitPrimary(titeParser::PrimaryContext * /*ctx*/) override { }

  virtual void enterTag(titeParser::TagContext * /*ctx*/) override { }
  virtual void exitTag(titeParser::TagContext * /*ctx*/) override { }

  virtual void enterAssign(titeParser::AssignContext * /*ctx*/) override { }
  virtual void exitAssign(titeParser::AssignContext * /*ctx*/) override { }

  virtual void enterLiteral(titeParser::LiteralContext * /*ctx*/) override { }
  virtual void exitLiteral(titeParser::LiteralContext * /*ctx*/) override { }

  virtual void enterBlock(titeParser::BlockContext * /*ctx*/) override { }
  virtual void exitBlock(titeParser::BlockContext * /*ctx*/) override { }

  virtual void enterArray(titeParser::ArrayContext * /*ctx*/) override { }
  virtual void exitArray(titeParser::ArrayContext * /*ctx*/) override { }

  virtual void enterInput(titeParser::InputContext * /*ctx*/) override { }
  virtual void exitInput(titeParser::InputContext * /*ctx*/) override { }

  virtual void enterDelim(titeParser::DelimContext * /*ctx*/) override { }
  virtual void exitDelim(titeParser::DelimContext * /*ctx*/) override { }


  virtual void enterEveryRule(antlr4::ParserRuleContext * /*ctx*/) override { }
  virtual void exitEveryRule(antlr4::ParserRuleContext * /*ctx*/) override { }
  virtual void visitTerminal(antlr4::tree::TerminalNode * /*node*/) override { }
  virtual void visitErrorNode(antlr4::tree::ErrorNode * /*node*/) override { }

};

