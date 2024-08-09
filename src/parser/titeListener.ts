// Generated from tite.g4 by ANTLR 4.13.2

import {ParseTreeListener} from "antlr4";


import { ProgramContext } from "./titeParser.js";
import { SequenceContext } from "./titeParser.js";
import { DeclarationContext } from "./titeParser.js";
import { IdentifiersContext } from "./titeParser.js";
import { ExpressionContext } from "./titeParser.js";
import { TypeContext } from "./titeParser.js";
import { FunctionContext } from "./titeParser.js";
import { ConditionContext } from "./titeParser.js";
import { OtherwiseContext } from "./titeParser.js";
import { DisjunctionContext } from "./titeParser.js";
import { ConjunctionContext } from "./titeParser.js";
import { OrContext } from "./titeParser.js";
import { XorContext } from "./titeParser.js";
import { AndContext } from "./titeParser.js";
import { EqualityContext } from "./titeParser.js";
import { RelationContext } from "./titeParser.js";
import { ShiftContext } from "./titeParser.js";
import { AdditiveContext } from "./titeParser.js";
import { ProductContext } from "./titeParser.js";
import { FactorContext } from "./titeParser.js";
import { PowerContext } from "./titeParser.js";
import { PostfixContext } from "./titeParser.js";
import { PrimaryContext } from "./titeParser.js";
import { TagContext } from "./titeParser.js";
import { AssignContext } from "./titeParser.js";
import { LiteralContext } from "./titeParser.js";
import { BlockContext } from "./titeParser.js";
import { ArrayContext } from "./titeParser.js";
import { InputContext } from "./titeParser.js";
import { DelimContext } from "./titeParser.js";


/**
 * This interface defines a complete listener for a parse tree produced by
 * `titeParser`.
 */
export default class titeListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by `titeParser.program`.
	 * @param ctx the parse tree
	 */
	enterProgram?: (ctx: ProgramContext) => void;
	/**
	 * Exit a parse tree produced by `titeParser.program`.
	 * @param ctx the parse tree
	 */
	exitProgram?: (ctx: ProgramContext) => void;
	/**
	 * Enter a parse tree produced by `titeParser.sequence`.
	 * @param ctx the parse tree
	 */
	enterSequence?: (ctx: SequenceContext) => void;
	/**
	 * Exit a parse tree produced by `titeParser.sequence`.
	 * @param ctx the parse tree
	 */
	exitSequence?: (ctx: SequenceContext) => void;
	/**
	 * Enter a parse tree produced by `titeParser.declaration`.
	 * @param ctx the parse tree
	 */
	enterDeclaration?: (ctx: DeclarationContext) => void;
	/**
	 * Exit a parse tree produced by `titeParser.declaration`.
	 * @param ctx the parse tree
	 */
	exitDeclaration?: (ctx: DeclarationContext) => void;
	/**
	 * Enter a parse tree produced by `titeParser.identifiers`.
	 * @param ctx the parse tree
	 */
	enterIdentifiers?: (ctx: IdentifiersContext) => void;
	/**
	 * Exit a parse tree produced by `titeParser.identifiers`.
	 * @param ctx the parse tree
	 */
	exitIdentifiers?: (ctx: IdentifiersContext) => void;
	/**
	 * Enter a parse tree produced by `titeParser.expression`.
	 * @param ctx the parse tree
	 */
	enterExpression?: (ctx: ExpressionContext) => void;
	/**
	 * Exit a parse tree produced by `titeParser.expression`.
	 * @param ctx the parse tree
	 */
	exitExpression?: (ctx: ExpressionContext) => void;
	/**
	 * Enter a parse tree produced by `titeParser.type`.
	 * @param ctx the parse tree
	 */
	enterType?: (ctx: TypeContext) => void;
	/**
	 * Exit a parse tree produced by `titeParser.type`.
	 * @param ctx the parse tree
	 */
	exitType?: (ctx: TypeContext) => void;
	/**
	 * Enter a parse tree produced by `titeParser.function`.
	 * @param ctx the parse tree
	 */
	enterFunction?: (ctx: FunctionContext) => void;
	/**
	 * Exit a parse tree produced by `titeParser.function`.
	 * @param ctx the parse tree
	 */
	exitFunction?: (ctx: FunctionContext) => void;
	/**
	 * Enter a parse tree produced by `titeParser.condition`.
	 * @param ctx the parse tree
	 */
	enterCondition?: (ctx: ConditionContext) => void;
	/**
	 * Exit a parse tree produced by `titeParser.condition`.
	 * @param ctx the parse tree
	 */
	exitCondition?: (ctx: ConditionContext) => void;
	/**
	 * Enter a parse tree produced by `titeParser.otherwise`.
	 * @param ctx the parse tree
	 */
	enterOtherwise?: (ctx: OtherwiseContext) => void;
	/**
	 * Exit a parse tree produced by `titeParser.otherwise`.
	 * @param ctx the parse tree
	 */
	exitOtherwise?: (ctx: OtherwiseContext) => void;
	/**
	 * Enter a parse tree produced by `titeParser.disjunction`.
	 * @param ctx the parse tree
	 */
	enterDisjunction?: (ctx: DisjunctionContext) => void;
	/**
	 * Exit a parse tree produced by `titeParser.disjunction`.
	 * @param ctx the parse tree
	 */
	exitDisjunction?: (ctx: DisjunctionContext) => void;
	/**
	 * Enter a parse tree produced by `titeParser.conjunction`.
	 * @param ctx the parse tree
	 */
	enterConjunction?: (ctx: ConjunctionContext) => void;
	/**
	 * Exit a parse tree produced by `titeParser.conjunction`.
	 * @param ctx the parse tree
	 */
	exitConjunction?: (ctx: ConjunctionContext) => void;
	/**
	 * Enter a parse tree produced by `titeParser.or`.
	 * @param ctx the parse tree
	 */
	enterOr?: (ctx: OrContext) => void;
	/**
	 * Exit a parse tree produced by `titeParser.or`.
	 * @param ctx the parse tree
	 */
	exitOr?: (ctx: OrContext) => void;
	/**
	 * Enter a parse tree produced by `titeParser.xor`.
	 * @param ctx the parse tree
	 */
	enterXor?: (ctx: XorContext) => void;
	/**
	 * Exit a parse tree produced by `titeParser.xor`.
	 * @param ctx the parse tree
	 */
	exitXor?: (ctx: XorContext) => void;
	/**
	 * Enter a parse tree produced by `titeParser.and`.
	 * @param ctx the parse tree
	 */
	enterAnd?: (ctx: AndContext) => void;
	/**
	 * Exit a parse tree produced by `titeParser.and`.
	 * @param ctx the parse tree
	 */
	exitAnd?: (ctx: AndContext) => void;
	/**
	 * Enter a parse tree produced by `titeParser.equality`.
	 * @param ctx the parse tree
	 */
	enterEquality?: (ctx: EqualityContext) => void;
	/**
	 * Exit a parse tree produced by `titeParser.equality`.
	 * @param ctx the parse tree
	 */
	exitEquality?: (ctx: EqualityContext) => void;
	/**
	 * Enter a parse tree produced by `titeParser.relation`.
	 * @param ctx the parse tree
	 */
	enterRelation?: (ctx: RelationContext) => void;
	/**
	 * Exit a parse tree produced by `titeParser.relation`.
	 * @param ctx the parse tree
	 */
	exitRelation?: (ctx: RelationContext) => void;
	/**
	 * Enter a parse tree produced by `titeParser.shift`.
	 * @param ctx the parse tree
	 */
	enterShift?: (ctx: ShiftContext) => void;
	/**
	 * Exit a parse tree produced by `titeParser.shift`.
	 * @param ctx the parse tree
	 */
	exitShift?: (ctx: ShiftContext) => void;
	/**
	 * Enter a parse tree produced by `titeParser.additive`.
	 * @param ctx the parse tree
	 */
	enterAdditive?: (ctx: AdditiveContext) => void;
	/**
	 * Exit a parse tree produced by `titeParser.additive`.
	 * @param ctx the parse tree
	 */
	exitAdditive?: (ctx: AdditiveContext) => void;
	/**
	 * Enter a parse tree produced by `titeParser.product`.
	 * @param ctx the parse tree
	 */
	enterProduct?: (ctx: ProductContext) => void;
	/**
	 * Exit a parse tree produced by `titeParser.product`.
	 * @param ctx the parse tree
	 */
	exitProduct?: (ctx: ProductContext) => void;
	/**
	 * Enter a parse tree produced by `titeParser.factor`.
	 * @param ctx the parse tree
	 */
	enterFactor?: (ctx: FactorContext) => void;
	/**
	 * Exit a parse tree produced by `titeParser.factor`.
	 * @param ctx the parse tree
	 */
	exitFactor?: (ctx: FactorContext) => void;
	/**
	 * Enter a parse tree produced by `titeParser.power`.
	 * @param ctx the parse tree
	 */
	enterPower?: (ctx: PowerContext) => void;
	/**
	 * Exit a parse tree produced by `titeParser.power`.
	 * @param ctx the parse tree
	 */
	exitPower?: (ctx: PowerContext) => void;
	/**
	 * Enter a parse tree produced by `titeParser.postfix`.
	 * @param ctx the parse tree
	 */
	enterPostfix?: (ctx: PostfixContext) => void;
	/**
	 * Exit a parse tree produced by `titeParser.postfix`.
	 * @param ctx the parse tree
	 */
	exitPostfix?: (ctx: PostfixContext) => void;
	/**
	 * Enter a parse tree produced by `titeParser.primary`.
	 * @param ctx the parse tree
	 */
	enterPrimary?: (ctx: PrimaryContext) => void;
	/**
	 * Exit a parse tree produced by `titeParser.primary`.
	 * @param ctx the parse tree
	 */
	exitPrimary?: (ctx: PrimaryContext) => void;
	/**
	 * Enter a parse tree produced by `titeParser.tag`.
	 * @param ctx the parse tree
	 */
	enterTag?: (ctx: TagContext) => void;
	/**
	 * Exit a parse tree produced by `titeParser.tag`.
	 * @param ctx the parse tree
	 */
	exitTag?: (ctx: TagContext) => void;
	/**
	 * Enter a parse tree produced by `titeParser.assign`.
	 * @param ctx the parse tree
	 */
	enterAssign?: (ctx: AssignContext) => void;
	/**
	 * Exit a parse tree produced by `titeParser.assign`.
	 * @param ctx the parse tree
	 */
	exitAssign?: (ctx: AssignContext) => void;
	/**
	 * Enter a parse tree produced by `titeParser.literal`.
	 * @param ctx the parse tree
	 */
	enterLiteral?: (ctx: LiteralContext) => void;
	/**
	 * Exit a parse tree produced by `titeParser.literal`.
	 * @param ctx the parse tree
	 */
	exitLiteral?: (ctx: LiteralContext) => void;
	/**
	 * Enter a parse tree produced by `titeParser.block`.
	 * @param ctx the parse tree
	 */
	enterBlock?: (ctx: BlockContext) => void;
	/**
	 * Exit a parse tree produced by `titeParser.block`.
	 * @param ctx the parse tree
	 */
	exitBlock?: (ctx: BlockContext) => void;
	/**
	 * Enter a parse tree produced by `titeParser.array`.
	 * @param ctx the parse tree
	 */
	enterArray?: (ctx: ArrayContext) => void;
	/**
	 * Exit a parse tree produced by `titeParser.array`.
	 * @param ctx the parse tree
	 */
	exitArray?: (ctx: ArrayContext) => void;
	/**
	 * Enter a parse tree produced by `titeParser.input`.
	 * @param ctx the parse tree
	 */
	enterInput?: (ctx: InputContext) => void;
	/**
	 * Exit a parse tree produced by `titeParser.input`.
	 * @param ctx the parse tree
	 */
	exitInput?: (ctx: InputContext) => void;
	/**
	 * Enter a parse tree produced by `titeParser.delim`.
	 * @param ctx the parse tree
	 */
	enterDelim?: (ctx: DelimContext) => void;
	/**
	 * Exit a parse tree produced by `titeParser.delim`.
	 * @param ctx the parse tree
	 */
	exitDelim?: (ctx: DelimContext) => void;
}

