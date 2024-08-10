
// Generated from tite.g4 by ANTLR 4.13.2

#pragma once


#include "antlr4-runtime.h"
#include "titeParser.h"


/**
 * This interface defines an abstract listener for a parse tree produced by titeParser.
 */
class  titeListener : public antlr4::tree::ParseTreeListener {
public:

  virtual void enterProgram(titeParser::ProgramContext *ctx) = 0;
  virtual void exitProgram(titeParser::ProgramContext *ctx) = 0;

  virtual void enterSequence(titeParser::SequenceContext *ctx) = 0;
  virtual void exitSequence(titeParser::SequenceContext *ctx) = 0;

  virtual void enterDeclaration(titeParser::DeclarationContext *ctx) = 0;
  virtual void exitDeclaration(titeParser::DeclarationContext *ctx) = 0;

  virtual void enterIdentifiers(titeParser::IdentifiersContext *ctx) = 0;
  virtual void exitIdentifiers(titeParser::IdentifiersContext *ctx) = 0;

  virtual void enterExpression(titeParser::ExpressionContext *ctx) = 0;
  virtual void exitExpression(titeParser::ExpressionContext *ctx) = 0;

  virtual void enterType(titeParser::TypeContext *ctx) = 0;
  virtual void exitType(titeParser::TypeContext *ctx) = 0;

  virtual void enterFunction(titeParser::FunctionContext *ctx) = 0;
  virtual void exitFunction(titeParser::FunctionContext *ctx) = 0;

  virtual void enterCondition(titeParser::ConditionContext *ctx) = 0;
  virtual void exitCondition(titeParser::ConditionContext *ctx) = 0;

  virtual void enterOtherwise(titeParser::OtherwiseContext *ctx) = 0;
  virtual void exitOtherwise(titeParser::OtherwiseContext *ctx) = 0;

  virtual void enterDisjunction(titeParser::DisjunctionContext *ctx) = 0;
  virtual void exitDisjunction(titeParser::DisjunctionContext *ctx) = 0;

  virtual void enterConjunction(titeParser::ConjunctionContext *ctx) = 0;
  virtual void exitConjunction(titeParser::ConjunctionContext *ctx) = 0;

  virtual void enterOr(titeParser::OrContext *ctx) = 0;
  virtual void exitOr(titeParser::OrContext *ctx) = 0;

  virtual void enterXor(titeParser::XorContext *ctx) = 0;
  virtual void exitXor(titeParser::XorContext *ctx) = 0;

  virtual void enterAnd(titeParser::AndContext *ctx) = 0;
  virtual void exitAnd(titeParser::AndContext *ctx) = 0;

  virtual void enterEquality(titeParser::EqualityContext *ctx) = 0;
  virtual void exitEquality(titeParser::EqualityContext *ctx) = 0;

  virtual void enterRelation(titeParser::RelationContext *ctx) = 0;
  virtual void exitRelation(titeParser::RelationContext *ctx) = 0;

  virtual void enterShift(titeParser::ShiftContext *ctx) = 0;
  virtual void exitShift(titeParser::ShiftContext *ctx) = 0;

  virtual void enterAdditive(titeParser::AdditiveContext *ctx) = 0;
  virtual void exitAdditive(titeParser::AdditiveContext *ctx) = 0;

  virtual void enterProduct(titeParser::ProductContext *ctx) = 0;
  virtual void exitProduct(titeParser::ProductContext *ctx) = 0;

  virtual void enterFactor(titeParser::FactorContext *ctx) = 0;
  virtual void exitFactor(titeParser::FactorContext *ctx) = 0;

  virtual void enterPower(titeParser::PowerContext *ctx) = 0;
  virtual void exitPower(titeParser::PowerContext *ctx) = 0;

  virtual void enterPostfix(titeParser::PostfixContext *ctx) = 0;
  virtual void exitPostfix(titeParser::PostfixContext *ctx) = 0;

  virtual void enterPrimary(titeParser::PrimaryContext *ctx) = 0;
  virtual void exitPrimary(titeParser::PrimaryContext *ctx) = 0;

  virtual void enterTag(titeParser::TagContext *ctx) = 0;
  virtual void exitTag(titeParser::TagContext *ctx) = 0;

  virtual void enterAssign(titeParser::AssignContext *ctx) = 0;
  virtual void exitAssign(titeParser::AssignContext *ctx) = 0;

  virtual void enterLiteral(titeParser::LiteralContext *ctx) = 0;
  virtual void exitLiteral(titeParser::LiteralContext *ctx) = 0;

  virtual void enterBlock(titeParser::BlockContext *ctx) = 0;
  virtual void exitBlock(titeParser::BlockContext *ctx) = 0;

  virtual void enterArray(titeParser::ArrayContext *ctx) = 0;
  virtual void exitArray(titeParser::ArrayContext *ctx) = 0;

  virtual void enterInput(titeParser::InputContext *ctx) = 0;
  virtual void exitInput(titeParser::InputContext *ctx) = 0;

  virtual void enterDelim(titeParser::DelimContext *ctx) = 0;
  virtual void exitDelim(titeParser::DelimContext *ctx) = 0;


};

