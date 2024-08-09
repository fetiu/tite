// Generated from tite.g4 by ANTLR 4.13.2
// noinspection ES6UnusedImports,JSUnusedGlobalSymbols,JSUnusedLocalSymbols

import {
	ATN,
	ATNDeserializer, DecisionState, DFA, FailedPredicateException,
	RecognitionException, NoViableAltException, BailErrorStrategy,
	Parser, ParserATNSimulator,
	RuleContext, ParserRuleContext, PredictionMode, PredictionContextCache,
	TerminalNode, RuleNode,
	Token, TokenStream,
	Interval, IntervalSet
} from 'antlr4';
import titeListener from "./titeListener.js";
// for running tests with parameters, TODO: discuss strategy for typed parameters in CI
// eslint-disable-next-line no-unused-vars
type int = number;

export default class titeParser extends Parser {
	public static readonly T__0 = 1;
	public static readonly T__1 = 2;
	public static readonly T__2 = 3;
	public static readonly T__3 = 4;
	public static readonly T__4 = 5;
	public static readonly T__5 = 6;
	public static readonly T__6 = 7;
	public static readonly T__7 = 8;
	public static readonly T__8 = 9;
	public static readonly T__9 = 10;
	public static readonly T__10 = 11;
	public static readonly T__11 = 12;
	public static readonly T__12 = 13;
	public static readonly T__13 = 14;
	public static readonly T__14 = 15;
	public static readonly T__15 = 16;
	public static readonly T__16 = 17;
	public static readonly T__17 = 18;
	public static readonly T__18 = 19;
	public static readonly T__19 = 20;
	public static readonly T__20 = 21;
	public static readonly T__21 = 22;
	public static readonly T__22 = 23;
	public static readonly T__23 = 24;
	public static readonly T__24 = 25;
	public static readonly T__25 = 26;
	public static readonly T__26 = 27;
	public static readonly T__27 = 28;
	public static readonly T__28 = 29;
	public static readonly T__29 = 30;
	public static readonly T__30 = 31;
	public static readonly T__31 = 32;
	public static readonly T__32 = 33;
	public static readonly T__33 = 34;
	public static readonly T__34 = 35;
	public static readonly T__35 = 36;
	public static readonly T__36 = 37;
	public static readonly T__37 = 38;
	public static readonly T__38 = 39;
	public static readonly T__39 = 40;
	public static readonly T__40 = 41;
	public static readonly T__41 = 42;
	public static readonly T__42 = 43;
	public static readonly T__43 = 44;
	public static readonly T__44 = 45;
	public static readonly T__45 = 46;
	public static readonly T__46 = 47;
	public static readonly T__47 = 48;
	public static readonly T__48 = 49;
	public static readonly T__49 = 50;
	public static readonly T__50 = 51;
	public static readonly T__51 = 52;
	public static readonly T__52 = 53;
	public static readonly T__53 = 54;
	public static readonly T__54 = 55;
	public static readonly IDENTIFIER = 56;
	public static readonly INT = 57;
	public static readonly FLOAT = 58;
	public static readonly CHAR = 59;
	public static readonly STR = 60;
	public static readonly WS = 61;
	public static readonly LF = 62;
	public static override readonly EOF = Token.EOF;
	public static readonly RULE_program = 0;
	public static readonly RULE_sequence = 1;
	public static readonly RULE_declaration = 2;
	public static readonly RULE_identifiers = 3;
	public static readonly RULE_expression = 4;
	public static readonly RULE_type = 5;
	public static readonly RULE_function = 6;
	public static readonly RULE_condition = 7;
	public static readonly RULE_otherwise = 8;
	public static readonly RULE_disjunction = 9;
	public static readonly RULE_conjunction = 10;
	public static readonly RULE_or = 11;
	public static readonly RULE_xor = 12;
	public static readonly RULE_and = 13;
	public static readonly RULE_equality = 14;
	public static readonly RULE_relation = 15;
	public static readonly RULE_shift = 16;
	public static readonly RULE_additive = 17;
	public static readonly RULE_product = 18;
	public static readonly RULE_factor = 19;
	public static readonly RULE_power = 20;
	public static readonly RULE_postfix = 21;
	public static readonly RULE_primary = 22;
	public static readonly RULE_tag = 23;
	public static readonly RULE_assign = 24;
	public static readonly RULE_literal = 25;
	public static readonly RULE_block = 26;
	public static readonly RULE_array = 27;
	public static readonly RULE_input = 28;
	public static readonly RULE_delim = 29;
	public static readonly literalNames: (string | null)[] = [ null, "':'", 
                                                            "','", "'*'", 
                                                            "'?'", "'||'", 
                                                            "'&&'", "'|'", 
                                                            "'^'", "'&'", 
                                                            "'=='", "'!='", 
                                                            "'==='", "'!=='", 
                                                            "'<'", "'>'", 
                                                            "'<='", "'>='", 
                                                            "'<<'", "'>>'", 
                                                            "'+'", "'-'", 
                                                            "'..'", "'/'", 
                                                            "'%'", "'//'", 
                                                            "'++'", "'--'", 
                                                            "'~'", "'!'", 
                                                            "'**'", "'.'", 
                                                            "'('", "')'", 
                                                            "'$'", "'#'", 
                                                            "'='", "'=>'", 
                                                            "'<-'", "'@'", 
                                                            "'*='", "'/='", 
                                                            "'%='", "'//='", 
                                                            "'**='", "'+='", 
                                                            "'-='", "'<<='", 
                                                            "'>>='", "'&='", 
                                                            "'^='", "'|='", 
                                                            "'{'", "'}'", 
                                                            "'['", "']'" ];
	public static readonly symbolicNames: (string | null)[] = [ null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             null, null, 
                                                             "IDENTIFIER", 
                                                             "INT", "FLOAT", 
                                                             "CHAR", "STR", 
                                                             "WS", "LF" ];
	// tslint:disable:no-trailing-whitespace
	public static readonly ruleNames: string[] = [
		"program", "sequence", "declaration", "identifiers", "expression", "type", 
		"function", "condition", "otherwise", "disjunction", "conjunction", "or", 
		"xor", "and", "equality", "relation", "shift", "additive", "product", 
		"factor", "power", "postfix", "primary", "tag", "assign", "literal", "block", 
		"array", "input", "delim",
	];
	public get grammarFileName(): string { return "tite.g4"; }
	public get literalNames(): (string | null)[] { return titeParser.literalNames; }
	public get symbolicNames(): (string | null)[] { return titeParser.symbolicNames; }
	public get ruleNames(): string[] { return titeParser.ruleNames; }
	public get serializedATN(): number[] { return titeParser._serializedATN; }

	protected createFailedPredicateException(predicate?: string, message?: string): FailedPredicateException {
		return new FailedPredicateException(this, predicate, message);
	}

	constructor(input: TokenStream) {
		super(input);
		this._interp = new ParserATNSimulator(this, titeParser._ATN, titeParser.DecisionsToDFA, new PredictionContextCache());
	}
	// @RuleVersion(0)
	public program(): ProgramContext {
		let localctx: ProgramContext = new ProgramContext(this, this._ctx, this.state);
		this.enterRule(localctx, 0, titeParser.RULE_program);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 60;
			this.sequence();
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public sequence(): SequenceContext {
		let localctx: SequenceContext = new SequenceContext(this, this._ctx, this.state);
		this.enterRule(localctx, 2, titeParser.RULE_sequence);
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 62;
			this.declaration();
			this.state = 68;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 0, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					{
					this.state = 63;
					this.delim();
					this.state = 64;
					this.declaration();
					}
					}
				}
				this.state = 70;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 0, this._ctx);
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public declaration(): DeclarationContext {
		let localctx: DeclarationContext = new DeclarationContext(this, this._ctx, this.state);
		this.enterRule(localctx, 4, titeParser.RULE_declaration);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 74;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 1, this._ctx) ) {
			case 1:
				{
				this.state = 71;
				this.identifiers();
				this.state = 72;
				this.match(titeParser.T__0);
				}
				break;
			}
			this.state = 76;
			this.expression();
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public identifiers(): IdentifiersContext {
		let localctx: IdentifiersContext = new IdentifiersContext(this, this._ctx, this.state);
		this.enterRule(localctx, 6, titeParser.RULE_identifiers);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 78;
			this.primary();
			this.state = 83;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===2) {
				{
				{
				this.state = 79;
				this.match(titeParser.T__1);
				this.state = 80;
				this.primary();
				}
				}
				this.state = 85;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public expression(): ExpressionContext {
		let localctx: ExpressionContext = new ExpressionContext(this, this._ctx, this.state);
		this.enterRule(localctx, 8, titeParser.RULE_expression);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 87;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===3) {
				{
				this.state = 86;
				this.match(titeParser.T__2);
				}
			}

			this.state = 99;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 6, this._ctx) ) {
			case 1:
				{
				this.state = 89;
				this.type_();
				}
				break;
			case 2:
				{
				this.state = 91;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if ((((_la) & ~0x1F) === 0 && ((1 << _la) & 3157262464) !== 0) || ((((_la - 32)) & ~0x1F) === 0 && ((1 << (_la - 32)) & 525336589) !== 0)) {
					{
					this.state = 90;
					this.type_();
					}
				}

				this.state = 93;
				this.assign();
				this.state = 95;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===62) {
					{
					this.state = 94;
					this.match(titeParser.LF);
					}
				}

				this.state = 97;
				this.expression();
				}
				break;
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public type_(): TypeContext {
		let localctx: TypeContext = new TypeContext(this, this._ctx, this.state);
		this.enterRule(localctx, 10, titeParser.RULE_type);
		try {
			this.state = 103;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 7, this._ctx) ) {
			case 1:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 101;
				this.condition();
				}
				break;
			case 2:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 102;
				this.function_();
				}
				break;
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public function_(): FunctionContext {
		let localctx: FunctionContext = new FunctionContext(this, this._ctx, this.state);
		this.enterRule(localctx, 12, titeParser.RULE_function);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 105;
			this.input();
			this.state = 107;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if ((((_la) & ~0x1F) === 0 && ((1 << _la) & 3157262464) !== 0) || ((((_la - 32)) & ~0x1F) === 0 && ((1 << (_la - 32)) & 525336589) !== 0)) {
				{
				this.state = 106;
				this.type_();
				}
			}

			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public condition(): ConditionContext {
		let localctx: ConditionContext = new ConditionContext(this, this._ctx, this.state);
		this.enterRule(localctx, 14, titeParser.RULE_condition);
		let _la: number;
		try {
			this.state = 122;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 12, this._ctx) ) {
			case 1:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 109;
				this.disjunction(0);
				}
				break;
			case 2:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 110;
				this.disjunction(0);
				this.state = 112;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===62) {
					{
					this.state = 111;
					this.match(titeParser.LF);
					}
				}

				this.state = 114;
				this.match(titeParser.T__3);
				this.state = 116;
				this._errHandler.sync(this);
				switch ( this._interp.adaptivePredict(this._input, 10, this._ctx) ) {
				case 1:
					{
					this.state = 115;
					this.match(titeParser.LF);
					}
					break;
				}
				this.state = 120;
				this._errHandler.sync(this);
				switch ( this._interp.adaptivePredict(this._input, 11, this._ctx) ) {
				case 1:
					{
					this.state = 118;
					this.expression();
					}
					break;
				case 2:
					{
					this.state = 119;
					this.otherwise();
					}
					break;
				}
				}
				break;
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public otherwise(): OtherwiseContext {
		let localctx: OtherwiseContext = new OtherwiseContext(this, this._ctx, this.state);
		this.enterRule(localctx, 16, titeParser.RULE_otherwise);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 125;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if ((((_la) & ~0x1F) === 0 && ((1 << _la) & 3157262472) !== 0) || ((((_la - 32)) & ~0x1F) === 0 && ((1 << (_la - 32)) & 526385149) !== 0)) {
				{
				this.state = 124;
				this.expression();
				}
			}

			this.state = 128;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===62) {
				{
				this.state = 127;
				this.match(titeParser.LF);
				}
			}

			this.state = 130;
			this.match(titeParser.T__0);
			this.state = 132;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===62) {
				{
				this.state = 131;
				this.match(titeParser.LF);
				}
			}

			this.state = 134;
			this.condition();
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}

	public disjunction(): DisjunctionContext;
	public disjunction(_p: number): DisjunctionContext;
	// @RuleVersion(0)
	public disjunction(_p?: number): DisjunctionContext {
		if (_p === undefined) {
			_p = 0;
		}

		let _parentctx: ParserRuleContext = this._ctx;
		let _parentState: number = this.state;
		let localctx: DisjunctionContext = new DisjunctionContext(this, this._ctx, _parentState);
		let _prevctx: DisjunctionContext = localctx;
		let _startState: number = 18;
		this.enterRecursionRule(localctx, 18, titeParser.RULE_disjunction, _p);
		let _la: number;
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			{
			this.state = 137;
			this.conjunction(0);
			}
			this._ctx.stop = this._input.LT(-1);
			this.state = 150;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 18, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					if (this._parseListeners != null) {
						this.triggerExitRuleEvent();
					}
					_prevctx = localctx;
					{
					{
					localctx = new DisjunctionContext(this, _parentctx, _parentState);
					this.pushNewRecursionContext(localctx, _startState, titeParser.RULE_disjunction);
					this.state = 139;
					if (!(this.precpred(this._ctx, 1))) {
						throw this.createFailedPredicateException("this.precpred(this._ctx, 1)");
					}
					this.state = 141;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
					if (_la===62) {
						{
						this.state = 140;
						this.match(titeParser.LF);
						}
					}

					this.state = 143;
					this.match(titeParser.T__4);
					this.state = 145;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
					if (_la===62) {
						{
						this.state = 144;
						this.match(titeParser.LF);
						}
					}

					this.state = 147;
					this.conjunction(0);
					}
					}
				}
				this.state = 152;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 18, this._ctx);
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.unrollRecursionContexts(_parentctx);
		}
		return localctx;
	}

	public conjunction(): ConjunctionContext;
	public conjunction(_p: number): ConjunctionContext;
	// @RuleVersion(0)
	public conjunction(_p?: number): ConjunctionContext {
		if (_p === undefined) {
			_p = 0;
		}

		let _parentctx: ParserRuleContext = this._ctx;
		let _parentState: number = this.state;
		let localctx: ConjunctionContext = new ConjunctionContext(this, this._ctx, _parentState);
		let _prevctx: ConjunctionContext = localctx;
		let _startState: number = 20;
		this.enterRecursionRule(localctx, 20, titeParser.RULE_conjunction, _p);
		let _la: number;
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			{
			this.state = 154;
			this.or(0);
			}
			this._ctx.stop = this._input.LT(-1);
			this.state = 167;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 21, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					if (this._parseListeners != null) {
						this.triggerExitRuleEvent();
					}
					_prevctx = localctx;
					{
					{
					localctx = new ConjunctionContext(this, _parentctx, _parentState);
					this.pushNewRecursionContext(localctx, _startState, titeParser.RULE_conjunction);
					this.state = 156;
					if (!(this.precpred(this._ctx, 1))) {
						throw this.createFailedPredicateException("this.precpred(this._ctx, 1)");
					}
					this.state = 158;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
					if (_la===62) {
						{
						this.state = 157;
						this.match(titeParser.LF);
						}
					}

					this.state = 160;
					this.match(titeParser.T__5);
					this.state = 162;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
					if (_la===62) {
						{
						this.state = 161;
						this.match(titeParser.LF);
						}
					}

					this.state = 164;
					this.or(0);
					}
					}
				}
				this.state = 169;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 21, this._ctx);
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.unrollRecursionContexts(_parentctx);
		}
		return localctx;
	}

	public or(): OrContext;
	public or(_p: number): OrContext;
	// @RuleVersion(0)
	public or(_p?: number): OrContext {
		if (_p === undefined) {
			_p = 0;
		}

		let _parentctx: ParserRuleContext = this._ctx;
		let _parentState: number = this.state;
		let localctx: OrContext = new OrContext(this, this._ctx, _parentState);
		let _prevctx: OrContext = localctx;
		let _startState: number = 22;
		this.enterRecursionRule(localctx, 22, titeParser.RULE_or, _p);
		let _la: number;
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			{
			this.state = 171;
			this.xor(0);
			}
			this._ctx.stop = this._input.LT(-1);
			this.state = 184;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 24, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					if (this._parseListeners != null) {
						this.triggerExitRuleEvent();
					}
					_prevctx = localctx;
					{
					{
					localctx = new OrContext(this, _parentctx, _parentState);
					this.pushNewRecursionContext(localctx, _startState, titeParser.RULE_or);
					this.state = 173;
					if (!(this.precpred(this._ctx, 1))) {
						throw this.createFailedPredicateException("this.precpred(this._ctx, 1)");
					}
					this.state = 175;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
					if (_la===62) {
						{
						this.state = 174;
						this.match(titeParser.LF);
						}
					}

					this.state = 177;
					this.match(titeParser.T__6);
					this.state = 179;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
					if (_la===62) {
						{
						this.state = 178;
						this.match(titeParser.LF);
						}
					}

					this.state = 181;
					this.xor(0);
					}
					}
				}
				this.state = 186;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 24, this._ctx);
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.unrollRecursionContexts(_parentctx);
		}
		return localctx;
	}

	public xor(): XorContext;
	public xor(_p: number): XorContext;
	// @RuleVersion(0)
	public xor(_p?: number): XorContext {
		if (_p === undefined) {
			_p = 0;
		}

		let _parentctx: ParserRuleContext = this._ctx;
		let _parentState: number = this.state;
		let localctx: XorContext = new XorContext(this, this._ctx, _parentState);
		let _prevctx: XorContext = localctx;
		let _startState: number = 24;
		this.enterRecursionRule(localctx, 24, titeParser.RULE_xor, _p);
		let _la: number;
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			{
			this.state = 188;
			this.and(0);
			}
			this._ctx.stop = this._input.LT(-1);
			this.state = 201;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 27, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					if (this._parseListeners != null) {
						this.triggerExitRuleEvent();
					}
					_prevctx = localctx;
					{
					{
					localctx = new XorContext(this, _parentctx, _parentState);
					this.pushNewRecursionContext(localctx, _startState, titeParser.RULE_xor);
					this.state = 190;
					if (!(this.precpred(this._ctx, 1))) {
						throw this.createFailedPredicateException("this.precpred(this._ctx, 1)");
					}
					this.state = 192;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
					if (_la===62) {
						{
						this.state = 191;
						this.match(titeParser.LF);
						}
					}

					this.state = 194;
					this.match(titeParser.T__7);
					this.state = 196;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
					if (_la===62) {
						{
						this.state = 195;
						this.match(titeParser.LF);
						}
					}

					this.state = 198;
					this.and(0);
					}
					}
				}
				this.state = 203;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 27, this._ctx);
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.unrollRecursionContexts(_parentctx);
		}
		return localctx;
	}

	public and(): AndContext;
	public and(_p: number): AndContext;
	// @RuleVersion(0)
	public and(_p?: number): AndContext {
		if (_p === undefined) {
			_p = 0;
		}

		let _parentctx: ParserRuleContext = this._ctx;
		let _parentState: number = this.state;
		let localctx: AndContext = new AndContext(this, this._ctx, _parentState);
		let _prevctx: AndContext = localctx;
		let _startState: number = 26;
		this.enterRecursionRule(localctx, 26, titeParser.RULE_and, _p);
		let _la: number;
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			{
			this.state = 205;
			this.equality(0);
			}
			this._ctx.stop = this._input.LT(-1);
			this.state = 218;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 30, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					if (this._parseListeners != null) {
						this.triggerExitRuleEvent();
					}
					_prevctx = localctx;
					{
					{
					localctx = new AndContext(this, _parentctx, _parentState);
					this.pushNewRecursionContext(localctx, _startState, titeParser.RULE_and);
					this.state = 207;
					if (!(this.precpred(this._ctx, 1))) {
						throw this.createFailedPredicateException("this.precpred(this._ctx, 1)");
					}
					this.state = 209;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
					if (_la===62) {
						{
						this.state = 208;
						this.match(titeParser.LF);
						}
					}

					this.state = 211;
					this.match(titeParser.T__8);
					this.state = 213;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
					if (_la===62) {
						{
						this.state = 212;
						this.match(titeParser.LF);
						}
					}

					this.state = 215;
					this.equality(0);
					}
					}
				}
				this.state = 220;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 30, this._ctx);
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.unrollRecursionContexts(_parentctx);
		}
		return localctx;
	}

	public equality(): EqualityContext;
	public equality(_p: number): EqualityContext;
	// @RuleVersion(0)
	public equality(_p?: number): EqualityContext {
		if (_p === undefined) {
			_p = 0;
		}

		let _parentctx: ParserRuleContext = this._ctx;
		let _parentState: number = this.state;
		let localctx: EqualityContext = new EqualityContext(this, this._ctx, _parentState);
		let _prevctx: EqualityContext = localctx;
		let _startState: number = 28;
		this.enterRecursionRule(localctx, 28, titeParser.RULE_equality, _p);
		let _la: number;
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			{
			this.state = 222;
			this.relation(0);
			}
			this._ctx.stop = this._input.LT(-1);
			this.state = 235;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 33, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					if (this._parseListeners != null) {
						this.triggerExitRuleEvent();
					}
					_prevctx = localctx;
					{
					{
					localctx = new EqualityContext(this, _parentctx, _parentState);
					this.pushNewRecursionContext(localctx, _startState, titeParser.RULE_equality);
					this.state = 224;
					if (!(this.precpred(this._ctx, 1))) {
						throw this.createFailedPredicateException("this.precpred(this._ctx, 1)");
					}
					this.state = 226;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
					if (_la===62) {
						{
						this.state = 225;
						this.match(titeParser.LF);
						}
					}

					this.state = 228;
					_la = this._input.LA(1);
					if(!((((_la) & ~0x1F) === 0 && ((1 << _la) & 15360) !== 0))) {
					this._errHandler.recoverInline(this);
					}
					else {
						this._errHandler.reportMatch(this);
					    this.consume();
					}
					this.state = 230;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
					if (_la===62) {
						{
						this.state = 229;
						this.match(titeParser.LF);
						}
					}

					this.state = 232;
					this.relation(0);
					}
					}
				}
				this.state = 237;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 33, this._ctx);
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.unrollRecursionContexts(_parentctx);
		}
		return localctx;
	}

	public relation(): RelationContext;
	public relation(_p: number): RelationContext;
	// @RuleVersion(0)
	public relation(_p?: number): RelationContext {
		if (_p === undefined) {
			_p = 0;
		}

		let _parentctx: ParserRuleContext = this._ctx;
		let _parentState: number = this.state;
		let localctx: RelationContext = new RelationContext(this, this._ctx, _parentState);
		let _prevctx: RelationContext = localctx;
		let _startState: number = 30;
		this.enterRecursionRule(localctx, 30, titeParser.RULE_relation, _p);
		let _la: number;
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			{
			this.state = 239;
			this.shift(0);
			}
			this._ctx.stop = this._input.LT(-1);
			this.state = 252;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 36, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					if (this._parseListeners != null) {
						this.triggerExitRuleEvent();
					}
					_prevctx = localctx;
					{
					{
					localctx = new RelationContext(this, _parentctx, _parentState);
					this.pushNewRecursionContext(localctx, _startState, titeParser.RULE_relation);
					this.state = 241;
					if (!(this.precpred(this._ctx, 1))) {
						throw this.createFailedPredicateException("this.precpred(this._ctx, 1)");
					}
					this.state = 243;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
					if (_la===62) {
						{
						this.state = 242;
						this.match(titeParser.LF);
						}
					}

					this.state = 245;
					_la = this._input.LA(1);
					if(!((((_la) & ~0x1F) === 0 && ((1 << _la) & 245760) !== 0))) {
					this._errHandler.recoverInline(this);
					}
					else {
						this._errHandler.reportMatch(this);
					    this.consume();
					}
					this.state = 247;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
					if (_la===62) {
						{
						this.state = 246;
						this.match(titeParser.LF);
						}
					}

					this.state = 249;
					this.shift(0);
					}
					}
				}
				this.state = 254;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 36, this._ctx);
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.unrollRecursionContexts(_parentctx);
		}
		return localctx;
	}

	public shift(): ShiftContext;
	public shift(_p: number): ShiftContext;
	// @RuleVersion(0)
	public shift(_p?: number): ShiftContext {
		if (_p === undefined) {
			_p = 0;
		}

		let _parentctx: ParserRuleContext = this._ctx;
		let _parentState: number = this.state;
		let localctx: ShiftContext = new ShiftContext(this, this._ctx, _parentState);
		let _prevctx: ShiftContext = localctx;
		let _startState: number = 32;
		this.enterRecursionRule(localctx, 32, titeParser.RULE_shift, _p);
		let _la: number;
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			{
			this.state = 256;
			this.additive(0);
			}
			this._ctx.stop = this._input.LT(-1);
			this.state = 269;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 39, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					if (this._parseListeners != null) {
						this.triggerExitRuleEvent();
					}
					_prevctx = localctx;
					{
					{
					localctx = new ShiftContext(this, _parentctx, _parentState);
					this.pushNewRecursionContext(localctx, _startState, titeParser.RULE_shift);
					this.state = 258;
					if (!(this.precpred(this._ctx, 1))) {
						throw this.createFailedPredicateException("this.precpred(this._ctx, 1)");
					}
					this.state = 260;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
					if (_la===62) {
						{
						this.state = 259;
						this.match(titeParser.LF);
						}
					}

					this.state = 262;
					_la = this._input.LA(1);
					if(!(_la===18 || _la===19)) {
					this._errHandler.recoverInline(this);
					}
					else {
						this._errHandler.reportMatch(this);
					    this.consume();
					}
					this.state = 264;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
					if (_la===62) {
						{
						this.state = 263;
						this.match(titeParser.LF);
						}
					}

					this.state = 266;
					this.additive(0);
					}
					}
				}
				this.state = 271;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 39, this._ctx);
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.unrollRecursionContexts(_parentctx);
		}
		return localctx;
	}

	public additive(): AdditiveContext;
	public additive(_p: number): AdditiveContext;
	// @RuleVersion(0)
	public additive(_p?: number): AdditiveContext {
		if (_p === undefined) {
			_p = 0;
		}

		let _parentctx: ParserRuleContext = this._ctx;
		let _parentState: number = this.state;
		let localctx: AdditiveContext = new AdditiveContext(this, this._ctx, _parentState);
		let _prevctx: AdditiveContext = localctx;
		let _startState: number = 34;
		this.enterRecursionRule(localctx, 34, titeParser.RULE_additive, _p);
		let _la: number;
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			{
			this.state = 273;
			this.product(0);
			}
			this._ctx.stop = this._input.LT(-1);
			this.state = 286;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 42, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					if (this._parseListeners != null) {
						this.triggerExitRuleEvent();
					}
					_prevctx = localctx;
					{
					{
					localctx = new AdditiveContext(this, _parentctx, _parentState);
					this.pushNewRecursionContext(localctx, _startState, titeParser.RULE_additive);
					this.state = 275;
					if (!(this.precpred(this._ctx, 1))) {
						throw this.createFailedPredicateException("this.precpred(this._ctx, 1)");
					}
					this.state = 277;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
					if (_la===62) {
						{
						this.state = 276;
						this.match(titeParser.LF);
						}
					}

					this.state = 279;
					_la = this._input.LA(1);
					if(!((((_la) & ~0x1F) === 0 && ((1 << _la) & 7340032) !== 0))) {
					this._errHandler.recoverInline(this);
					}
					else {
						this._errHandler.reportMatch(this);
					    this.consume();
					}
					this.state = 281;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
					if (_la===62) {
						{
						this.state = 280;
						this.match(titeParser.LF);
						}
					}

					this.state = 283;
					this.product(0);
					}
					}
				}
				this.state = 288;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 42, this._ctx);
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.unrollRecursionContexts(_parentctx);
		}
		return localctx;
	}

	public product(): ProductContext;
	public product(_p: number): ProductContext;
	// @RuleVersion(0)
	public product(_p?: number): ProductContext {
		if (_p === undefined) {
			_p = 0;
		}

		let _parentctx: ParserRuleContext = this._ctx;
		let _parentState: number = this.state;
		let localctx: ProductContext = new ProductContext(this, this._ctx, _parentState);
		let _prevctx: ProductContext = localctx;
		let _startState: number = 36;
		this.enterRecursionRule(localctx, 36, titeParser.RULE_product, _p);
		let _la: number;
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			{
			this.state = 290;
			this.factor();
			}
			this._ctx.stop = this._input.LT(-1);
			this.state = 303;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 45, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					if (this._parseListeners != null) {
						this.triggerExitRuleEvent();
					}
					_prevctx = localctx;
					{
					{
					localctx = new ProductContext(this, _parentctx, _parentState);
					this.pushNewRecursionContext(localctx, _startState, titeParser.RULE_product);
					this.state = 292;
					if (!(this.precpred(this._ctx, 1))) {
						throw this.createFailedPredicateException("this.precpred(this._ctx, 1)");
					}
					this.state = 294;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
					if (_la===62) {
						{
						this.state = 293;
						this.match(titeParser.LF);
						}
					}

					this.state = 296;
					_la = this._input.LA(1);
					if(!((((_la) & ~0x1F) === 0 && ((1 << _la) & 58720264) !== 0))) {
					this._errHandler.recoverInline(this);
					}
					else {
						this._errHandler.reportMatch(this);
					    this.consume();
					}
					this.state = 298;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
					if (_la===62) {
						{
						this.state = 297;
						this.match(titeParser.LF);
						}
					}

					this.state = 300;
					this.factor();
					}
					}
				}
				this.state = 305;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 45, this._ctx);
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.unrollRecursionContexts(_parentctx);
		}
		return localctx;
	}
	// @RuleVersion(0)
	public factor(): FactorContext {
		let localctx: FactorContext = new FactorContext(this, this._ctx, this.state);
		this.enterRule(localctx, 38, titeParser.RULE_factor);
		let _la: number;
		try {
			this.state = 316;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 31:
			case 32:
			case 34:
			case 35:
			case 52:
			case 54:
			case 56:
			case 57:
			case 58:
			case 59:
			case 60:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 306;
				this.power();
				}
				break;
			case 20:
			case 21:
			case 26:
			case 27:
			case 28:
			case 29:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 307;
				_la = this._input.LA(1);
				if(!((((_la) & ~0x1F) === 0 && ((1 << _la) & 1009778688) !== 0))) {
				this._errHandler.recoverInline(this);
				}
				else {
					this._errHandler.reportMatch(this);
				    this.consume();
				}
				this.state = 309;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===62) {
					{
					this.state = 308;
					this.match(titeParser.LF);
					}
				}

				this.state = 311;
				this.factor();
				}
				break;
			case 7:
				this.enterOuterAlt(localctx, 3);
				{
				this.state = 312;
				this.match(titeParser.T__6);
				this.state = 313;
				this.factor();
				this.state = 314;
				this.match(titeParser.T__6);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public power(): PowerContext {
		let localctx: PowerContext = new PowerContext(this, this._ctx, this.state);
		this.enterRule(localctx, 40, titeParser.RULE_power);
		let _la: number;
		try {
			this.state = 329;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 50, this._ctx) ) {
			case 1:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 318;
				this.postfix(0);
				}
				break;
			case 2:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 319;
				this.postfix(0);
				this.state = 321;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===62) {
					{
					this.state = 320;
					this.match(titeParser.LF);
					}
				}

				this.state = 323;
				this.match(titeParser.T__29);
				this.state = 325;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (_la===62) {
					{
					this.state = 324;
					this.match(titeParser.LF);
					}
				}

				this.state = 327;
				this.factor();
				}
				break;
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}

	public postfix(): PostfixContext;
	public postfix(_p: number): PostfixContext;
	// @RuleVersion(0)
	public postfix(_p?: number): PostfixContext {
		if (_p === undefined) {
			_p = 0;
		}

		let _parentctx: ParserRuleContext = this._ctx;
		let _parentState: number = this.state;
		let localctx: PostfixContext = new PostfixContext(this, this._ctx, _parentState);
		let _prevctx: PostfixContext = localctx;
		let _startState: number = 42;
		this.enterRecursionRule(localctx, 42, titeParser.RULE_postfix, _p);
		let _la: number;
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			{
			this.state = 332;
			this.primary();
			}
			this._ctx.stop = this._input.LT(-1);
			this.state = 348;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 53, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					if (this._parseListeners != null) {
						this.triggerExitRuleEvent();
					}
					_prevctx = localctx;
					{
					{
					localctx = new PostfixContext(this, _parentctx, _parentState);
					this.pushNewRecursionContext(localctx, _startState, titeParser.RULE_postfix);
					this.state = 334;
					if (!(this.precpred(this._ctx, 1))) {
						throw this.createFailedPredicateException("this.precpred(this._ctx, 1)");
					}
					this.state = 336;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
					if (_la===62) {
						{
						this.state = 335;
						this.match(titeParser.LF);
						}
					}

					this.state = 344;
					this._errHandler.sync(this);
					switch (this._input.LA(1)) {
					case 26:
						{
						this.state = 338;
						this.match(titeParser.T__25);
						}
						break;
					case 27:
						{
						this.state = 339;
						this.match(titeParser.T__26);
						}
						break;
					case 31:
						{
						this.state = 340;
						this.match(titeParser.T__30);
						this.state = 341;
						this.match(titeParser.IDENTIFIER);
						}
						break;
					case 54:
						{
						this.state = 342;
						this.array();
						}
						break;
					case 32:
						{
						this.state = 343;
						this.input();
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					}
					}
				}
				this.state = 350;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 53, this._ctx);
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.unrollRecursionContexts(_parentctx);
		}
		return localctx;
	}
	// @RuleVersion(0)
	public primary(): PrimaryContext {
		let localctx: PrimaryContext = new PrimaryContext(this, this._ctx, this.state);
		this.enterRule(localctx, 44, titeParser.RULE_primary);
		let _la: number;
		try {
			this.state = 360;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 31:
			case 34:
			case 35:
			case 56:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 352;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				if (((((_la - 31)) & ~0x1F) === 0 && ((1 << (_la - 31)) & 25) !== 0)) {
					{
					this.state = 351;
					this.tag();
					}
				}

				this.state = 354;
				this.match(titeParser.IDENTIFIER);
				}
				break;
			case 52:
			case 54:
			case 57:
			case 58:
			case 59:
			case 60:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 355;
				this.literal();
				}
				break;
			case 32:
				this.enterOuterAlt(localctx, 3);
				{
				this.state = 356;
				this.match(titeParser.T__31);
				this.state = 357;
				this.expression();
				this.state = 358;
				this.match(titeParser.T__32);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public tag(): TagContext {
		let localctx: TagContext = new TagContext(this, this._ctx, this.state);
		this.enterRule(localctx, 46, titeParser.RULE_tag);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 362;
			_la = this._input.LA(1);
			if(!(((((_la - 31)) & ~0x1F) === 0 && ((1 << (_la - 31)) & 25) !== 0))) {
			this._errHandler.recoverInline(this);
			}
			else {
				this._errHandler.reportMatch(this);
			    this.consume();
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public assign(): AssignContext {
		let localctx: AssignContext = new AssignContext(this, this._ctx, this.state);
		this.enterRule(localctx, 48, titeParser.RULE_assign);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 364;
			_la = this._input.LA(1);
			if(!(((((_la - 36)) & ~0x1F) === 0 && ((1 << (_la - 36)) & 65535) !== 0))) {
			this._errHandler.recoverInline(this);
			}
			else {
				this._errHandler.reportMatch(this);
			    this.consume();
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public literal(): LiteralContext {
		let localctx: LiteralContext = new LiteralContext(this, this._ctx, this.state);
		this.enterRule(localctx, 50, titeParser.RULE_literal);
		try {
			this.state = 372;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 57:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 366;
				this.match(titeParser.INT);
				}
				break;
			case 58:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 367;
				this.match(titeParser.FLOAT);
				}
				break;
			case 59:
				this.enterOuterAlt(localctx, 3);
				{
				this.state = 368;
				this.match(titeParser.CHAR);
				}
				break;
			case 60:
				this.enterOuterAlt(localctx, 4);
				{
				this.state = 369;
				this.match(titeParser.STR);
				}
				break;
			case 52:
				this.enterOuterAlt(localctx, 5);
				{
				this.state = 370;
				this.block();
				}
				break;
			case 54:
				this.enterOuterAlt(localctx, 6);
				{
				this.state = 371;
				this.array();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public block(): BlockContext {
		let localctx: BlockContext = new BlockContext(this, this._ctx, this.state);
		this.enterRule(localctx, 52, titeParser.RULE_block);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 374;
			this.match(titeParser.T__51);
			this.state = 376;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 57, this._ctx) ) {
			case 1:
				{
				this.state = 375;
				this.match(titeParser.LF);
				}
				break;
			}
			this.state = 379;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if ((((_la) & ~0x1F) === 0 && ((1 << _la) & 3157262472) !== 0) || ((((_la - 32)) & ~0x1F) === 0 && ((1 << (_la - 32)) & 526385149) !== 0)) {
				{
				this.state = 378;
				this.sequence();
				}
			}

			this.state = 382;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===62) {
				{
				this.state = 381;
				this.match(titeParser.LF);
				}
			}

			this.state = 384;
			this.match(titeParser.T__52);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public array(): ArrayContext {
		let localctx: ArrayContext = new ArrayContext(this, this._ctx, this.state);
		this.enterRule(localctx, 54, titeParser.RULE_array);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 386;
			this.match(titeParser.T__53);
			this.state = 388;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 60, this._ctx) ) {
			case 1:
				{
				this.state = 387;
				this.match(titeParser.LF);
				}
				break;
			}
			this.state = 391;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if ((((_la) & ~0x1F) === 0 && ((1 << _la) & 3157262472) !== 0) || ((((_la - 32)) & ~0x1F) === 0 && ((1 << (_la - 32)) & 526385149) !== 0)) {
				{
				this.state = 390;
				this.sequence();
				}
			}

			this.state = 394;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===62) {
				{
				this.state = 393;
				this.match(titeParser.LF);
				}
			}

			this.state = 396;
			this.match(titeParser.T__54);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public input(): InputContext {
		let localctx: InputContext = new InputContext(this, this._ctx, this.state);
		this.enterRule(localctx, 56, titeParser.RULE_input);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 398;
			this.match(titeParser.T__31);
			this.state = 400;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 63, this._ctx) ) {
			case 1:
				{
				this.state = 399;
				this.match(titeParser.LF);
				}
				break;
			}
			this.state = 403;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if ((((_la) & ~0x1F) === 0 && ((1 << _la) & 3157262472) !== 0) || ((((_la - 32)) & ~0x1F) === 0 && ((1 << (_la - 32)) & 526385149) !== 0)) {
				{
				this.state = 402;
				this.sequence();
				}
			}

			this.state = 406;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===62) {
				{
				this.state = 405;
				this.match(titeParser.LF);
				}
			}

			this.state = 408;
			this.match(titeParser.T__32);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public delim(): DelimContext {
		let localctx: DelimContext = new DelimContext(this, this._ctx, this.state);
		this.enterRule(localctx, 58, titeParser.RULE_delim);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 410;
			_la = this._input.LA(1);
			if(!(_la===2 || _la===62)) {
			this._errHandler.recoverInline(this);
			}
			else {
				this._errHandler.reportMatch(this);
			    this.consume();
			}
			this.state = 412;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===62) {
				{
				this.state = 411;
				this.match(titeParser.LF);
				}
			}

			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}

	public sempred(localctx: RuleContext, ruleIndex: number, predIndex: number): boolean {
		switch (ruleIndex) {
		case 9:
			return this.disjunction_sempred(localctx as DisjunctionContext, predIndex);
		case 10:
			return this.conjunction_sempred(localctx as ConjunctionContext, predIndex);
		case 11:
			return this.or_sempred(localctx as OrContext, predIndex);
		case 12:
			return this.xor_sempred(localctx as XorContext, predIndex);
		case 13:
			return this.and_sempred(localctx as AndContext, predIndex);
		case 14:
			return this.equality_sempred(localctx as EqualityContext, predIndex);
		case 15:
			return this.relation_sempred(localctx as RelationContext, predIndex);
		case 16:
			return this.shift_sempred(localctx as ShiftContext, predIndex);
		case 17:
			return this.additive_sempred(localctx as AdditiveContext, predIndex);
		case 18:
			return this.product_sempred(localctx as ProductContext, predIndex);
		case 21:
			return this.postfix_sempred(localctx as PostfixContext, predIndex);
		}
		return true;
	}
	private disjunction_sempred(localctx: DisjunctionContext, predIndex: number): boolean {
		switch (predIndex) {
		case 0:
			return this.precpred(this._ctx, 1);
		}
		return true;
	}
	private conjunction_sempred(localctx: ConjunctionContext, predIndex: number): boolean {
		switch (predIndex) {
		case 1:
			return this.precpred(this._ctx, 1);
		}
		return true;
	}
	private or_sempred(localctx: OrContext, predIndex: number): boolean {
		switch (predIndex) {
		case 2:
			return this.precpred(this._ctx, 1);
		}
		return true;
	}
	private xor_sempred(localctx: XorContext, predIndex: number): boolean {
		switch (predIndex) {
		case 3:
			return this.precpred(this._ctx, 1);
		}
		return true;
	}
	private and_sempred(localctx: AndContext, predIndex: number): boolean {
		switch (predIndex) {
		case 4:
			return this.precpred(this._ctx, 1);
		}
		return true;
	}
	private equality_sempred(localctx: EqualityContext, predIndex: number): boolean {
		switch (predIndex) {
		case 5:
			return this.precpred(this._ctx, 1);
		}
		return true;
	}
	private relation_sempred(localctx: RelationContext, predIndex: number): boolean {
		switch (predIndex) {
		case 6:
			return this.precpred(this._ctx, 1);
		}
		return true;
	}
	private shift_sempred(localctx: ShiftContext, predIndex: number): boolean {
		switch (predIndex) {
		case 7:
			return this.precpred(this._ctx, 1);
		}
		return true;
	}
	private additive_sempred(localctx: AdditiveContext, predIndex: number): boolean {
		switch (predIndex) {
		case 8:
			return this.precpred(this._ctx, 1);
		}
		return true;
	}
	private product_sempred(localctx: ProductContext, predIndex: number): boolean {
		switch (predIndex) {
		case 9:
			return this.precpred(this._ctx, 1);
		}
		return true;
	}
	private postfix_sempred(localctx: PostfixContext, predIndex: number): boolean {
		switch (predIndex) {
		case 10:
			return this.precpred(this._ctx, 1);
		}
		return true;
	}

	public static readonly _serializedATN: number[] = [4,1,62,415,2,0,7,0,2,
	1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,6,2,7,7,7,2,8,7,8,2,9,7,9,2,
	10,7,10,2,11,7,11,2,12,7,12,2,13,7,13,2,14,7,14,2,15,7,15,2,16,7,16,2,17,
	7,17,2,18,7,18,2,19,7,19,2,20,7,20,2,21,7,21,2,22,7,22,2,23,7,23,2,24,7,
	24,2,25,7,25,2,26,7,26,2,27,7,27,2,28,7,28,2,29,7,29,1,0,1,0,1,1,1,1,1,
	1,1,1,5,1,67,8,1,10,1,12,1,70,9,1,1,2,1,2,1,2,3,2,75,8,2,1,2,1,2,1,3,1,
	3,1,3,5,3,82,8,3,10,3,12,3,85,9,3,1,4,3,4,88,8,4,1,4,1,4,3,4,92,8,4,1,4,
	1,4,3,4,96,8,4,1,4,1,4,3,4,100,8,4,1,5,1,5,3,5,104,8,5,1,6,1,6,3,6,108,
	8,6,1,7,1,7,1,7,3,7,113,8,7,1,7,1,7,3,7,117,8,7,1,7,1,7,3,7,121,8,7,3,7,
	123,8,7,1,8,3,8,126,8,8,1,8,3,8,129,8,8,1,8,1,8,3,8,133,8,8,1,8,1,8,1,9,
	1,9,1,9,1,9,1,9,3,9,142,8,9,1,9,1,9,3,9,146,8,9,1,9,5,9,149,8,9,10,9,12,
	9,152,9,9,1,10,1,10,1,10,1,10,1,10,3,10,159,8,10,1,10,1,10,3,10,163,8,10,
	1,10,5,10,166,8,10,10,10,12,10,169,9,10,1,11,1,11,1,11,1,11,1,11,3,11,176,
	8,11,1,11,1,11,3,11,180,8,11,1,11,5,11,183,8,11,10,11,12,11,186,9,11,1,
	12,1,12,1,12,1,12,1,12,3,12,193,8,12,1,12,1,12,3,12,197,8,12,1,12,5,12,
	200,8,12,10,12,12,12,203,9,12,1,13,1,13,1,13,1,13,1,13,3,13,210,8,13,1,
	13,1,13,3,13,214,8,13,1,13,5,13,217,8,13,10,13,12,13,220,9,13,1,14,1,14,
	1,14,1,14,1,14,3,14,227,8,14,1,14,1,14,3,14,231,8,14,1,14,5,14,234,8,14,
	10,14,12,14,237,9,14,1,15,1,15,1,15,1,15,1,15,3,15,244,8,15,1,15,1,15,3,
	15,248,8,15,1,15,5,15,251,8,15,10,15,12,15,254,9,15,1,16,1,16,1,16,1,16,
	1,16,3,16,261,8,16,1,16,1,16,3,16,265,8,16,1,16,5,16,268,8,16,10,16,12,
	16,271,9,16,1,17,1,17,1,17,1,17,1,17,3,17,278,8,17,1,17,1,17,3,17,282,8,
	17,1,17,5,17,285,8,17,10,17,12,17,288,9,17,1,18,1,18,1,18,1,18,1,18,3,18,
	295,8,18,1,18,1,18,3,18,299,8,18,1,18,5,18,302,8,18,10,18,12,18,305,9,18,
	1,19,1,19,1,19,3,19,310,8,19,1,19,1,19,1,19,1,19,1,19,3,19,317,8,19,1,20,
	1,20,1,20,3,20,322,8,20,1,20,1,20,3,20,326,8,20,1,20,1,20,3,20,330,8,20,
	1,21,1,21,1,21,1,21,1,21,3,21,337,8,21,1,21,1,21,1,21,1,21,1,21,1,21,3,
	21,345,8,21,5,21,347,8,21,10,21,12,21,350,9,21,1,22,3,22,353,8,22,1,22,
	1,22,1,22,1,22,1,22,1,22,3,22,361,8,22,1,23,1,23,1,24,1,24,1,25,1,25,1,
	25,1,25,1,25,1,25,3,25,373,8,25,1,26,1,26,3,26,377,8,26,1,26,3,26,380,8,
	26,1,26,3,26,383,8,26,1,26,1,26,1,27,1,27,3,27,389,8,27,1,27,3,27,392,8,
	27,1,27,3,27,395,8,27,1,27,1,27,1,28,1,28,3,28,401,8,28,1,28,3,28,404,8,
	28,1,28,3,28,407,8,28,1,28,1,28,1,29,1,29,3,29,413,8,29,1,29,0,11,18,20,
	22,24,26,28,30,32,34,36,42,30,0,2,4,6,8,10,12,14,16,18,20,22,24,26,28,30,
	32,34,36,38,40,42,44,46,48,50,52,54,56,58,0,9,1,0,10,13,1,0,14,17,1,0,18,
	19,1,0,20,22,2,0,3,3,23,25,2,0,20,21,26,29,2,0,31,31,34,35,1,0,36,51,2,
	0,2,2,62,62,460,0,60,1,0,0,0,2,62,1,0,0,0,4,74,1,0,0,0,6,78,1,0,0,0,8,87,
	1,0,0,0,10,103,1,0,0,0,12,105,1,0,0,0,14,122,1,0,0,0,16,125,1,0,0,0,18,
	136,1,0,0,0,20,153,1,0,0,0,22,170,1,0,0,0,24,187,1,0,0,0,26,204,1,0,0,0,
	28,221,1,0,0,0,30,238,1,0,0,0,32,255,1,0,0,0,34,272,1,0,0,0,36,289,1,0,
	0,0,38,316,1,0,0,0,40,329,1,0,0,0,42,331,1,0,0,0,44,360,1,0,0,0,46,362,
	1,0,0,0,48,364,1,0,0,0,50,372,1,0,0,0,52,374,1,0,0,0,54,386,1,0,0,0,56,
	398,1,0,0,0,58,410,1,0,0,0,60,61,3,2,1,0,61,1,1,0,0,0,62,68,3,4,2,0,63,
	64,3,58,29,0,64,65,3,4,2,0,65,67,1,0,0,0,66,63,1,0,0,0,67,70,1,0,0,0,68,
	66,1,0,0,0,68,69,1,0,0,0,69,3,1,0,0,0,70,68,1,0,0,0,71,72,3,6,3,0,72,73,
	5,1,0,0,73,75,1,0,0,0,74,71,1,0,0,0,74,75,1,0,0,0,75,76,1,0,0,0,76,77,3,
	8,4,0,77,5,1,0,0,0,78,83,3,44,22,0,79,80,5,2,0,0,80,82,3,44,22,0,81,79,
	1,0,0,0,82,85,1,0,0,0,83,81,1,0,0,0,83,84,1,0,0,0,84,7,1,0,0,0,85,83,1,
	0,0,0,86,88,5,3,0,0,87,86,1,0,0,0,87,88,1,0,0,0,88,99,1,0,0,0,89,100,3,
	10,5,0,90,92,3,10,5,0,91,90,1,0,0,0,91,92,1,0,0,0,92,93,1,0,0,0,93,95,3,
	48,24,0,94,96,5,62,0,0,95,94,1,0,0,0,95,96,1,0,0,0,96,97,1,0,0,0,97,98,
	3,8,4,0,98,100,1,0,0,0,99,89,1,0,0,0,99,91,1,0,0,0,100,9,1,0,0,0,101,104,
	3,14,7,0,102,104,3,12,6,0,103,101,1,0,0,0,103,102,1,0,0,0,104,11,1,0,0,
	0,105,107,3,56,28,0,106,108,3,10,5,0,107,106,1,0,0,0,107,108,1,0,0,0,108,
	13,1,0,0,0,109,123,3,18,9,0,110,112,3,18,9,0,111,113,5,62,0,0,112,111,1,
	0,0,0,112,113,1,0,0,0,113,114,1,0,0,0,114,116,5,4,0,0,115,117,5,62,0,0,
	116,115,1,0,0,0,116,117,1,0,0,0,117,120,1,0,0,0,118,121,3,8,4,0,119,121,
	3,16,8,0,120,118,1,0,0,0,120,119,1,0,0,0,121,123,1,0,0,0,122,109,1,0,0,
	0,122,110,1,0,0,0,123,15,1,0,0,0,124,126,3,8,4,0,125,124,1,0,0,0,125,126,
	1,0,0,0,126,128,1,0,0,0,127,129,5,62,0,0,128,127,1,0,0,0,128,129,1,0,0,
	0,129,130,1,0,0,0,130,132,5,1,0,0,131,133,5,62,0,0,132,131,1,0,0,0,132,
	133,1,0,0,0,133,134,1,0,0,0,134,135,3,14,7,0,135,17,1,0,0,0,136,137,6,9,
	-1,0,137,138,3,20,10,0,138,150,1,0,0,0,139,141,10,1,0,0,140,142,5,62,0,
	0,141,140,1,0,0,0,141,142,1,0,0,0,142,143,1,0,0,0,143,145,5,5,0,0,144,146,
	5,62,0,0,145,144,1,0,0,0,145,146,1,0,0,0,146,147,1,0,0,0,147,149,3,20,10,
	0,148,139,1,0,0,0,149,152,1,0,0,0,150,148,1,0,0,0,150,151,1,0,0,0,151,19,
	1,0,0,0,152,150,1,0,0,0,153,154,6,10,-1,0,154,155,3,22,11,0,155,167,1,0,
	0,0,156,158,10,1,0,0,157,159,5,62,0,0,158,157,1,0,0,0,158,159,1,0,0,0,159,
	160,1,0,0,0,160,162,5,6,0,0,161,163,5,62,0,0,162,161,1,0,0,0,162,163,1,
	0,0,0,163,164,1,0,0,0,164,166,3,22,11,0,165,156,1,0,0,0,166,169,1,0,0,0,
	167,165,1,0,0,0,167,168,1,0,0,0,168,21,1,0,0,0,169,167,1,0,0,0,170,171,
	6,11,-1,0,171,172,3,24,12,0,172,184,1,0,0,0,173,175,10,1,0,0,174,176,5,
	62,0,0,175,174,1,0,0,0,175,176,1,0,0,0,176,177,1,0,0,0,177,179,5,7,0,0,
	178,180,5,62,0,0,179,178,1,0,0,0,179,180,1,0,0,0,180,181,1,0,0,0,181,183,
	3,24,12,0,182,173,1,0,0,0,183,186,1,0,0,0,184,182,1,0,0,0,184,185,1,0,0,
	0,185,23,1,0,0,0,186,184,1,0,0,0,187,188,6,12,-1,0,188,189,3,26,13,0,189,
	201,1,0,0,0,190,192,10,1,0,0,191,193,5,62,0,0,192,191,1,0,0,0,192,193,1,
	0,0,0,193,194,1,0,0,0,194,196,5,8,0,0,195,197,5,62,0,0,196,195,1,0,0,0,
	196,197,1,0,0,0,197,198,1,0,0,0,198,200,3,26,13,0,199,190,1,0,0,0,200,203,
	1,0,0,0,201,199,1,0,0,0,201,202,1,0,0,0,202,25,1,0,0,0,203,201,1,0,0,0,
	204,205,6,13,-1,0,205,206,3,28,14,0,206,218,1,0,0,0,207,209,10,1,0,0,208,
	210,5,62,0,0,209,208,1,0,0,0,209,210,1,0,0,0,210,211,1,0,0,0,211,213,5,
	9,0,0,212,214,5,62,0,0,213,212,1,0,0,0,213,214,1,0,0,0,214,215,1,0,0,0,
	215,217,3,28,14,0,216,207,1,0,0,0,217,220,1,0,0,0,218,216,1,0,0,0,218,219,
	1,0,0,0,219,27,1,0,0,0,220,218,1,0,0,0,221,222,6,14,-1,0,222,223,3,30,15,
	0,223,235,1,0,0,0,224,226,10,1,0,0,225,227,5,62,0,0,226,225,1,0,0,0,226,
	227,1,0,0,0,227,228,1,0,0,0,228,230,7,0,0,0,229,231,5,62,0,0,230,229,1,
	0,0,0,230,231,1,0,0,0,231,232,1,0,0,0,232,234,3,30,15,0,233,224,1,0,0,0,
	234,237,1,0,0,0,235,233,1,0,0,0,235,236,1,0,0,0,236,29,1,0,0,0,237,235,
	1,0,0,0,238,239,6,15,-1,0,239,240,3,32,16,0,240,252,1,0,0,0,241,243,10,
	1,0,0,242,244,5,62,0,0,243,242,1,0,0,0,243,244,1,0,0,0,244,245,1,0,0,0,
	245,247,7,1,0,0,246,248,5,62,0,0,247,246,1,0,0,0,247,248,1,0,0,0,248,249,
	1,0,0,0,249,251,3,32,16,0,250,241,1,0,0,0,251,254,1,0,0,0,252,250,1,0,0,
	0,252,253,1,0,0,0,253,31,1,0,0,0,254,252,1,0,0,0,255,256,6,16,-1,0,256,
	257,3,34,17,0,257,269,1,0,0,0,258,260,10,1,0,0,259,261,5,62,0,0,260,259,
	1,0,0,0,260,261,1,0,0,0,261,262,1,0,0,0,262,264,7,2,0,0,263,265,5,62,0,
	0,264,263,1,0,0,0,264,265,1,0,0,0,265,266,1,0,0,0,266,268,3,34,17,0,267,
	258,1,0,0,0,268,271,1,0,0,0,269,267,1,0,0,0,269,270,1,0,0,0,270,33,1,0,
	0,0,271,269,1,0,0,0,272,273,6,17,-1,0,273,274,3,36,18,0,274,286,1,0,0,0,
	275,277,10,1,0,0,276,278,5,62,0,0,277,276,1,0,0,0,277,278,1,0,0,0,278,279,
	1,0,0,0,279,281,7,3,0,0,280,282,5,62,0,0,281,280,1,0,0,0,281,282,1,0,0,
	0,282,283,1,0,0,0,283,285,3,36,18,0,284,275,1,0,0,0,285,288,1,0,0,0,286,
	284,1,0,0,0,286,287,1,0,0,0,287,35,1,0,0,0,288,286,1,0,0,0,289,290,6,18,
	-1,0,290,291,3,38,19,0,291,303,1,0,0,0,292,294,10,1,0,0,293,295,5,62,0,
	0,294,293,1,0,0,0,294,295,1,0,0,0,295,296,1,0,0,0,296,298,7,4,0,0,297,299,
	5,62,0,0,298,297,1,0,0,0,298,299,1,0,0,0,299,300,1,0,0,0,300,302,3,38,19,
	0,301,292,1,0,0,0,302,305,1,0,0,0,303,301,1,0,0,0,303,304,1,0,0,0,304,37,
	1,0,0,0,305,303,1,0,0,0,306,317,3,40,20,0,307,309,7,5,0,0,308,310,5,62,
	0,0,309,308,1,0,0,0,309,310,1,0,0,0,310,311,1,0,0,0,311,317,3,38,19,0,312,
	313,5,7,0,0,313,314,3,38,19,0,314,315,5,7,0,0,315,317,1,0,0,0,316,306,1,
	0,0,0,316,307,1,0,0,0,316,312,1,0,0,0,317,39,1,0,0,0,318,330,3,42,21,0,
	319,321,3,42,21,0,320,322,5,62,0,0,321,320,1,0,0,0,321,322,1,0,0,0,322,
	323,1,0,0,0,323,325,5,30,0,0,324,326,5,62,0,0,325,324,1,0,0,0,325,326,1,
	0,0,0,326,327,1,0,0,0,327,328,3,38,19,0,328,330,1,0,0,0,329,318,1,0,0,0,
	329,319,1,0,0,0,330,41,1,0,0,0,331,332,6,21,-1,0,332,333,3,44,22,0,333,
	348,1,0,0,0,334,336,10,1,0,0,335,337,5,62,0,0,336,335,1,0,0,0,336,337,1,
	0,0,0,337,344,1,0,0,0,338,345,5,26,0,0,339,345,5,27,0,0,340,341,5,31,0,
	0,341,345,5,56,0,0,342,345,3,54,27,0,343,345,3,56,28,0,344,338,1,0,0,0,
	344,339,1,0,0,0,344,340,1,0,0,0,344,342,1,0,0,0,344,343,1,0,0,0,345,347,
	1,0,0,0,346,334,1,0,0,0,347,350,1,0,0,0,348,346,1,0,0,0,348,349,1,0,0,0,
	349,43,1,0,0,0,350,348,1,0,0,0,351,353,3,46,23,0,352,351,1,0,0,0,352,353,
	1,0,0,0,353,354,1,0,0,0,354,361,5,56,0,0,355,361,3,50,25,0,356,357,5,32,
	0,0,357,358,3,8,4,0,358,359,5,33,0,0,359,361,1,0,0,0,360,352,1,0,0,0,360,
	355,1,0,0,0,360,356,1,0,0,0,361,45,1,0,0,0,362,363,7,6,0,0,363,47,1,0,0,
	0,364,365,7,7,0,0,365,49,1,0,0,0,366,373,5,57,0,0,367,373,5,58,0,0,368,
	373,5,59,0,0,369,373,5,60,0,0,370,373,3,52,26,0,371,373,3,54,27,0,372,366,
	1,0,0,0,372,367,1,0,0,0,372,368,1,0,0,0,372,369,1,0,0,0,372,370,1,0,0,0,
	372,371,1,0,0,0,373,51,1,0,0,0,374,376,5,52,0,0,375,377,5,62,0,0,376,375,
	1,0,0,0,376,377,1,0,0,0,377,379,1,0,0,0,378,380,3,2,1,0,379,378,1,0,0,0,
	379,380,1,0,0,0,380,382,1,0,0,0,381,383,5,62,0,0,382,381,1,0,0,0,382,383,
	1,0,0,0,383,384,1,0,0,0,384,385,5,53,0,0,385,53,1,0,0,0,386,388,5,54,0,
	0,387,389,5,62,0,0,388,387,1,0,0,0,388,389,1,0,0,0,389,391,1,0,0,0,390,
	392,3,2,1,0,391,390,1,0,0,0,391,392,1,0,0,0,392,394,1,0,0,0,393,395,5,62,
	0,0,394,393,1,0,0,0,394,395,1,0,0,0,395,396,1,0,0,0,396,397,5,55,0,0,397,
	55,1,0,0,0,398,400,5,32,0,0,399,401,5,62,0,0,400,399,1,0,0,0,400,401,1,
	0,0,0,401,403,1,0,0,0,402,404,3,2,1,0,403,402,1,0,0,0,403,404,1,0,0,0,404,
	406,1,0,0,0,405,407,5,62,0,0,406,405,1,0,0,0,406,407,1,0,0,0,407,408,1,
	0,0,0,408,409,5,33,0,0,409,57,1,0,0,0,410,412,7,8,0,0,411,413,5,62,0,0,
	412,411,1,0,0,0,412,413,1,0,0,0,413,59,1,0,0,0,67,68,74,83,87,91,95,99,
	103,107,112,116,120,122,125,128,132,141,145,150,158,162,167,175,179,184,
	192,196,201,209,213,218,226,230,235,243,247,252,260,264,269,277,281,286,
	294,298,303,309,316,321,325,329,336,344,348,352,360,372,376,379,382,388,
	391,394,400,403,406,412];

	private static __ATN: ATN;
	public static get _ATN(): ATN {
		if (!titeParser.__ATN) {
			titeParser.__ATN = new ATNDeserializer().deserialize(titeParser._serializedATN);
		}

		return titeParser.__ATN;
	}


	static DecisionsToDFA = titeParser._ATN.decisionToState.map( (ds: DecisionState, index: number) => new DFA(ds, index) );

}

export class ProgramContext extends ParserRuleContext {
	constructor(parser?: titeParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public sequence(): SequenceContext {
		return this.getTypedRuleContext(SequenceContext, 0) as SequenceContext;
	}
    public get ruleIndex(): number {
    	return titeParser.RULE_program;
	}
	public enterRule(listener: titeListener): void {
	    if(listener.enterProgram) {
	 		listener.enterProgram(this);
		}
	}
	public exitRule(listener: titeListener): void {
	    if(listener.exitProgram) {
	 		listener.exitProgram(this);
		}
	}
}


export class SequenceContext extends ParserRuleContext {
	constructor(parser?: titeParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public declaration_list(): DeclarationContext[] {
		return this.getTypedRuleContexts(DeclarationContext) as DeclarationContext[];
	}
	public declaration(i: number): DeclarationContext {
		return this.getTypedRuleContext(DeclarationContext, i) as DeclarationContext;
	}
	public delim_list(): DelimContext[] {
		return this.getTypedRuleContexts(DelimContext) as DelimContext[];
	}
	public delim(i: number): DelimContext {
		return this.getTypedRuleContext(DelimContext, i) as DelimContext;
	}
    public get ruleIndex(): number {
    	return titeParser.RULE_sequence;
	}
	public enterRule(listener: titeListener): void {
	    if(listener.enterSequence) {
	 		listener.enterSequence(this);
		}
	}
	public exitRule(listener: titeListener): void {
	    if(listener.exitSequence) {
	 		listener.exitSequence(this);
		}
	}
}


export class DeclarationContext extends ParserRuleContext {
	constructor(parser?: titeParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public expression(): ExpressionContext {
		return this.getTypedRuleContext(ExpressionContext, 0) as ExpressionContext;
	}
	public identifiers(): IdentifiersContext {
		return this.getTypedRuleContext(IdentifiersContext, 0) as IdentifiersContext;
	}
    public get ruleIndex(): number {
    	return titeParser.RULE_declaration;
	}
	public enterRule(listener: titeListener): void {
	    if(listener.enterDeclaration) {
	 		listener.enterDeclaration(this);
		}
	}
	public exitRule(listener: titeListener): void {
	    if(listener.exitDeclaration) {
	 		listener.exitDeclaration(this);
		}
	}
}


export class IdentifiersContext extends ParserRuleContext {
	constructor(parser?: titeParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public primary_list(): PrimaryContext[] {
		return this.getTypedRuleContexts(PrimaryContext) as PrimaryContext[];
	}
	public primary(i: number): PrimaryContext {
		return this.getTypedRuleContext(PrimaryContext, i) as PrimaryContext;
	}
    public get ruleIndex(): number {
    	return titeParser.RULE_identifiers;
	}
	public enterRule(listener: titeListener): void {
	    if(listener.enterIdentifiers) {
	 		listener.enterIdentifiers(this);
		}
	}
	public exitRule(listener: titeListener): void {
	    if(listener.exitIdentifiers) {
	 		listener.exitIdentifiers(this);
		}
	}
}


export class ExpressionContext extends ParserRuleContext {
	constructor(parser?: titeParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public type_(): TypeContext {
		return this.getTypedRuleContext(TypeContext, 0) as TypeContext;
	}
	public assign(): AssignContext {
		return this.getTypedRuleContext(AssignContext, 0) as AssignContext;
	}
	public expression(): ExpressionContext {
		return this.getTypedRuleContext(ExpressionContext, 0) as ExpressionContext;
	}
	public LF(): TerminalNode {
		return this.getToken(titeParser.LF, 0);
	}
    public get ruleIndex(): number {
    	return titeParser.RULE_expression;
	}
	public enterRule(listener: titeListener): void {
	    if(listener.enterExpression) {
	 		listener.enterExpression(this);
		}
	}
	public exitRule(listener: titeListener): void {
	    if(listener.exitExpression) {
	 		listener.exitExpression(this);
		}
	}
}


export class TypeContext extends ParserRuleContext {
	constructor(parser?: titeParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public condition(): ConditionContext {
		return this.getTypedRuleContext(ConditionContext, 0) as ConditionContext;
	}
	public function_(): FunctionContext {
		return this.getTypedRuleContext(FunctionContext, 0) as FunctionContext;
	}
    public get ruleIndex(): number {
    	return titeParser.RULE_type;
	}
	public enterRule(listener: titeListener): void {
	    if(listener.enterType) {
	 		listener.enterType(this);
		}
	}
	public exitRule(listener: titeListener): void {
	    if(listener.exitType) {
	 		listener.exitType(this);
		}
	}
}


export class FunctionContext extends ParserRuleContext {
	constructor(parser?: titeParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public input(): InputContext {
		return this.getTypedRuleContext(InputContext, 0) as InputContext;
	}
	public type_(): TypeContext {
		return this.getTypedRuleContext(TypeContext, 0) as TypeContext;
	}
    public get ruleIndex(): number {
    	return titeParser.RULE_function;
	}
	public enterRule(listener: titeListener): void {
	    if(listener.enterFunction) {
	 		listener.enterFunction(this);
		}
	}
	public exitRule(listener: titeListener): void {
	    if(listener.exitFunction) {
	 		listener.exitFunction(this);
		}
	}
}


export class ConditionContext extends ParserRuleContext {
	constructor(parser?: titeParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public disjunction(): DisjunctionContext {
		return this.getTypedRuleContext(DisjunctionContext, 0) as DisjunctionContext;
	}
	public expression(): ExpressionContext {
		return this.getTypedRuleContext(ExpressionContext, 0) as ExpressionContext;
	}
	public otherwise(): OtherwiseContext {
		return this.getTypedRuleContext(OtherwiseContext, 0) as OtherwiseContext;
	}
	public LF_list(): TerminalNode[] {
	    	return this.getTokens(titeParser.LF);
	}
	public LF(i: number): TerminalNode {
		return this.getToken(titeParser.LF, i);
	}
    public get ruleIndex(): number {
    	return titeParser.RULE_condition;
	}
	public enterRule(listener: titeListener): void {
	    if(listener.enterCondition) {
	 		listener.enterCondition(this);
		}
	}
	public exitRule(listener: titeListener): void {
	    if(listener.exitCondition) {
	 		listener.exitCondition(this);
		}
	}
}


export class OtherwiseContext extends ParserRuleContext {
	constructor(parser?: titeParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public condition(): ConditionContext {
		return this.getTypedRuleContext(ConditionContext, 0) as ConditionContext;
	}
	public expression(): ExpressionContext {
		return this.getTypedRuleContext(ExpressionContext, 0) as ExpressionContext;
	}
	public LF_list(): TerminalNode[] {
	    	return this.getTokens(titeParser.LF);
	}
	public LF(i: number): TerminalNode {
		return this.getToken(titeParser.LF, i);
	}
    public get ruleIndex(): number {
    	return titeParser.RULE_otherwise;
	}
	public enterRule(listener: titeListener): void {
	    if(listener.enterOtherwise) {
	 		listener.enterOtherwise(this);
		}
	}
	public exitRule(listener: titeListener): void {
	    if(listener.exitOtherwise) {
	 		listener.exitOtherwise(this);
		}
	}
}


export class DisjunctionContext extends ParserRuleContext {
	constructor(parser?: titeParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public conjunction(): ConjunctionContext {
		return this.getTypedRuleContext(ConjunctionContext, 0) as ConjunctionContext;
	}
	public disjunction(): DisjunctionContext {
		return this.getTypedRuleContext(DisjunctionContext, 0) as DisjunctionContext;
	}
	public LF_list(): TerminalNode[] {
	    	return this.getTokens(titeParser.LF);
	}
	public LF(i: number): TerminalNode {
		return this.getToken(titeParser.LF, i);
	}
    public get ruleIndex(): number {
    	return titeParser.RULE_disjunction;
	}
	public enterRule(listener: titeListener): void {
	    if(listener.enterDisjunction) {
	 		listener.enterDisjunction(this);
		}
	}
	public exitRule(listener: titeListener): void {
	    if(listener.exitDisjunction) {
	 		listener.exitDisjunction(this);
		}
	}
}


export class ConjunctionContext extends ParserRuleContext {
	constructor(parser?: titeParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public or(): OrContext {
		return this.getTypedRuleContext(OrContext, 0) as OrContext;
	}
	public conjunction(): ConjunctionContext {
		return this.getTypedRuleContext(ConjunctionContext, 0) as ConjunctionContext;
	}
	public LF_list(): TerminalNode[] {
	    	return this.getTokens(titeParser.LF);
	}
	public LF(i: number): TerminalNode {
		return this.getToken(titeParser.LF, i);
	}
    public get ruleIndex(): number {
    	return titeParser.RULE_conjunction;
	}
	public enterRule(listener: titeListener): void {
	    if(listener.enterConjunction) {
	 		listener.enterConjunction(this);
		}
	}
	public exitRule(listener: titeListener): void {
	    if(listener.exitConjunction) {
	 		listener.exitConjunction(this);
		}
	}
}


export class OrContext extends ParserRuleContext {
	constructor(parser?: titeParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public xor(): XorContext {
		return this.getTypedRuleContext(XorContext, 0) as XorContext;
	}
	public or(): OrContext {
		return this.getTypedRuleContext(OrContext, 0) as OrContext;
	}
	public LF_list(): TerminalNode[] {
	    	return this.getTokens(titeParser.LF);
	}
	public LF(i: number): TerminalNode {
		return this.getToken(titeParser.LF, i);
	}
    public get ruleIndex(): number {
    	return titeParser.RULE_or;
	}
	public enterRule(listener: titeListener): void {
	    if(listener.enterOr) {
	 		listener.enterOr(this);
		}
	}
	public exitRule(listener: titeListener): void {
	    if(listener.exitOr) {
	 		listener.exitOr(this);
		}
	}
}


export class XorContext extends ParserRuleContext {
	constructor(parser?: titeParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public and(): AndContext {
		return this.getTypedRuleContext(AndContext, 0) as AndContext;
	}
	public xor(): XorContext {
		return this.getTypedRuleContext(XorContext, 0) as XorContext;
	}
	public LF_list(): TerminalNode[] {
	    	return this.getTokens(titeParser.LF);
	}
	public LF(i: number): TerminalNode {
		return this.getToken(titeParser.LF, i);
	}
    public get ruleIndex(): number {
    	return titeParser.RULE_xor;
	}
	public enterRule(listener: titeListener): void {
	    if(listener.enterXor) {
	 		listener.enterXor(this);
		}
	}
	public exitRule(listener: titeListener): void {
	    if(listener.exitXor) {
	 		listener.exitXor(this);
		}
	}
}


export class AndContext extends ParserRuleContext {
	constructor(parser?: titeParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public equality(): EqualityContext {
		return this.getTypedRuleContext(EqualityContext, 0) as EqualityContext;
	}
	public and(): AndContext {
		return this.getTypedRuleContext(AndContext, 0) as AndContext;
	}
	public LF_list(): TerminalNode[] {
	    	return this.getTokens(titeParser.LF);
	}
	public LF(i: number): TerminalNode {
		return this.getToken(titeParser.LF, i);
	}
    public get ruleIndex(): number {
    	return titeParser.RULE_and;
	}
	public enterRule(listener: titeListener): void {
	    if(listener.enterAnd) {
	 		listener.enterAnd(this);
		}
	}
	public exitRule(listener: titeListener): void {
	    if(listener.exitAnd) {
	 		listener.exitAnd(this);
		}
	}
}


export class EqualityContext extends ParserRuleContext {
	constructor(parser?: titeParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public relation(): RelationContext {
		return this.getTypedRuleContext(RelationContext, 0) as RelationContext;
	}
	public equality(): EqualityContext {
		return this.getTypedRuleContext(EqualityContext, 0) as EqualityContext;
	}
	public LF_list(): TerminalNode[] {
	    	return this.getTokens(titeParser.LF);
	}
	public LF(i: number): TerminalNode {
		return this.getToken(titeParser.LF, i);
	}
    public get ruleIndex(): number {
    	return titeParser.RULE_equality;
	}
	public enterRule(listener: titeListener): void {
	    if(listener.enterEquality) {
	 		listener.enterEquality(this);
		}
	}
	public exitRule(listener: titeListener): void {
	    if(listener.exitEquality) {
	 		listener.exitEquality(this);
		}
	}
}


export class RelationContext extends ParserRuleContext {
	constructor(parser?: titeParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public shift(): ShiftContext {
		return this.getTypedRuleContext(ShiftContext, 0) as ShiftContext;
	}
	public relation(): RelationContext {
		return this.getTypedRuleContext(RelationContext, 0) as RelationContext;
	}
	public LF_list(): TerminalNode[] {
	    	return this.getTokens(titeParser.LF);
	}
	public LF(i: number): TerminalNode {
		return this.getToken(titeParser.LF, i);
	}
    public get ruleIndex(): number {
    	return titeParser.RULE_relation;
	}
	public enterRule(listener: titeListener): void {
	    if(listener.enterRelation) {
	 		listener.enterRelation(this);
		}
	}
	public exitRule(listener: titeListener): void {
	    if(listener.exitRelation) {
	 		listener.exitRelation(this);
		}
	}
}


export class ShiftContext extends ParserRuleContext {
	constructor(parser?: titeParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public additive(): AdditiveContext {
		return this.getTypedRuleContext(AdditiveContext, 0) as AdditiveContext;
	}
	public shift(): ShiftContext {
		return this.getTypedRuleContext(ShiftContext, 0) as ShiftContext;
	}
	public LF_list(): TerminalNode[] {
	    	return this.getTokens(titeParser.LF);
	}
	public LF(i: number): TerminalNode {
		return this.getToken(titeParser.LF, i);
	}
    public get ruleIndex(): number {
    	return titeParser.RULE_shift;
	}
	public enterRule(listener: titeListener): void {
	    if(listener.enterShift) {
	 		listener.enterShift(this);
		}
	}
	public exitRule(listener: titeListener): void {
	    if(listener.exitShift) {
	 		listener.exitShift(this);
		}
	}
}


export class AdditiveContext extends ParserRuleContext {
	constructor(parser?: titeParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public product(): ProductContext {
		return this.getTypedRuleContext(ProductContext, 0) as ProductContext;
	}
	public additive(): AdditiveContext {
		return this.getTypedRuleContext(AdditiveContext, 0) as AdditiveContext;
	}
	public LF_list(): TerminalNode[] {
	    	return this.getTokens(titeParser.LF);
	}
	public LF(i: number): TerminalNode {
		return this.getToken(titeParser.LF, i);
	}
    public get ruleIndex(): number {
    	return titeParser.RULE_additive;
	}
	public enterRule(listener: titeListener): void {
	    if(listener.enterAdditive) {
	 		listener.enterAdditive(this);
		}
	}
	public exitRule(listener: titeListener): void {
	    if(listener.exitAdditive) {
	 		listener.exitAdditive(this);
		}
	}
}


export class ProductContext extends ParserRuleContext {
	constructor(parser?: titeParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public factor(): FactorContext {
		return this.getTypedRuleContext(FactorContext, 0) as FactorContext;
	}
	public product(): ProductContext {
		return this.getTypedRuleContext(ProductContext, 0) as ProductContext;
	}
	public LF_list(): TerminalNode[] {
	    	return this.getTokens(titeParser.LF);
	}
	public LF(i: number): TerminalNode {
		return this.getToken(titeParser.LF, i);
	}
    public get ruleIndex(): number {
    	return titeParser.RULE_product;
	}
	public enterRule(listener: titeListener): void {
	    if(listener.enterProduct) {
	 		listener.enterProduct(this);
		}
	}
	public exitRule(listener: titeListener): void {
	    if(listener.exitProduct) {
	 		listener.exitProduct(this);
		}
	}
}


export class FactorContext extends ParserRuleContext {
	constructor(parser?: titeParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public power(): PowerContext {
		return this.getTypedRuleContext(PowerContext, 0) as PowerContext;
	}
	public factor(): FactorContext {
		return this.getTypedRuleContext(FactorContext, 0) as FactorContext;
	}
	public LF(): TerminalNode {
		return this.getToken(titeParser.LF, 0);
	}
    public get ruleIndex(): number {
    	return titeParser.RULE_factor;
	}
	public enterRule(listener: titeListener): void {
	    if(listener.enterFactor) {
	 		listener.enterFactor(this);
		}
	}
	public exitRule(listener: titeListener): void {
	    if(listener.exitFactor) {
	 		listener.exitFactor(this);
		}
	}
}


export class PowerContext extends ParserRuleContext {
	constructor(parser?: titeParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public postfix(): PostfixContext {
		return this.getTypedRuleContext(PostfixContext, 0) as PostfixContext;
	}
	public factor(): FactorContext {
		return this.getTypedRuleContext(FactorContext, 0) as FactorContext;
	}
	public LF_list(): TerminalNode[] {
	    	return this.getTokens(titeParser.LF);
	}
	public LF(i: number): TerminalNode {
		return this.getToken(titeParser.LF, i);
	}
    public get ruleIndex(): number {
    	return titeParser.RULE_power;
	}
	public enterRule(listener: titeListener): void {
	    if(listener.enterPower) {
	 		listener.enterPower(this);
		}
	}
	public exitRule(listener: titeListener): void {
	    if(listener.exitPower) {
	 		listener.exitPower(this);
		}
	}
}


export class PostfixContext extends ParserRuleContext {
	constructor(parser?: titeParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public primary(): PrimaryContext {
		return this.getTypedRuleContext(PrimaryContext, 0) as PrimaryContext;
	}
	public postfix(): PostfixContext {
		return this.getTypedRuleContext(PostfixContext, 0) as PostfixContext;
	}
	public IDENTIFIER(): TerminalNode {
		return this.getToken(titeParser.IDENTIFIER, 0);
	}
	public array(): ArrayContext {
		return this.getTypedRuleContext(ArrayContext, 0) as ArrayContext;
	}
	public input(): InputContext {
		return this.getTypedRuleContext(InputContext, 0) as InputContext;
	}
	public LF(): TerminalNode {
		return this.getToken(titeParser.LF, 0);
	}
    public get ruleIndex(): number {
    	return titeParser.RULE_postfix;
	}
	public enterRule(listener: titeListener): void {
	    if(listener.enterPostfix) {
	 		listener.enterPostfix(this);
		}
	}
	public exitRule(listener: titeListener): void {
	    if(listener.exitPostfix) {
	 		listener.exitPostfix(this);
		}
	}
}


export class PrimaryContext extends ParserRuleContext {
	constructor(parser?: titeParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public IDENTIFIER(): TerminalNode {
		return this.getToken(titeParser.IDENTIFIER, 0);
	}
	public tag(): TagContext {
		return this.getTypedRuleContext(TagContext, 0) as TagContext;
	}
	public literal(): LiteralContext {
		return this.getTypedRuleContext(LiteralContext, 0) as LiteralContext;
	}
	public expression(): ExpressionContext {
		return this.getTypedRuleContext(ExpressionContext, 0) as ExpressionContext;
	}
    public get ruleIndex(): number {
    	return titeParser.RULE_primary;
	}
	public enterRule(listener: titeListener): void {
	    if(listener.enterPrimary) {
	 		listener.enterPrimary(this);
		}
	}
	public exitRule(listener: titeListener): void {
	    if(listener.exitPrimary) {
	 		listener.exitPrimary(this);
		}
	}
}


export class TagContext extends ParserRuleContext {
	constructor(parser?: titeParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
    public get ruleIndex(): number {
    	return titeParser.RULE_tag;
	}
	public enterRule(listener: titeListener): void {
	    if(listener.enterTag) {
	 		listener.enterTag(this);
		}
	}
	public exitRule(listener: titeListener): void {
	    if(listener.exitTag) {
	 		listener.exitTag(this);
		}
	}
}


export class AssignContext extends ParserRuleContext {
	constructor(parser?: titeParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
    public get ruleIndex(): number {
    	return titeParser.RULE_assign;
	}
	public enterRule(listener: titeListener): void {
	    if(listener.enterAssign) {
	 		listener.enterAssign(this);
		}
	}
	public exitRule(listener: titeListener): void {
	    if(listener.exitAssign) {
	 		listener.exitAssign(this);
		}
	}
}


export class LiteralContext extends ParserRuleContext {
	constructor(parser?: titeParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public INT(): TerminalNode {
		return this.getToken(titeParser.INT, 0);
	}
	public FLOAT(): TerminalNode {
		return this.getToken(titeParser.FLOAT, 0);
	}
	public CHAR(): TerminalNode {
		return this.getToken(titeParser.CHAR, 0);
	}
	public STR(): TerminalNode {
		return this.getToken(titeParser.STR, 0);
	}
	public block(): BlockContext {
		return this.getTypedRuleContext(BlockContext, 0) as BlockContext;
	}
	public array(): ArrayContext {
		return this.getTypedRuleContext(ArrayContext, 0) as ArrayContext;
	}
    public get ruleIndex(): number {
    	return titeParser.RULE_literal;
	}
	public enterRule(listener: titeListener): void {
	    if(listener.enterLiteral) {
	 		listener.enterLiteral(this);
		}
	}
	public exitRule(listener: titeListener): void {
	    if(listener.exitLiteral) {
	 		listener.exitLiteral(this);
		}
	}
}


export class BlockContext extends ParserRuleContext {
	constructor(parser?: titeParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public LF_list(): TerminalNode[] {
	    	return this.getTokens(titeParser.LF);
	}
	public LF(i: number): TerminalNode {
		return this.getToken(titeParser.LF, i);
	}
	public sequence(): SequenceContext {
		return this.getTypedRuleContext(SequenceContext, 0) as SequenceContext;
	}
    public get ruleIndex(): number {
    	return titeParser.RULE_block;
	}
	public enterRule(listener: titeListener): void {
	    if(listener.enterBlock) {
	 		listener.enterBlock(this);
		}
	}
	public exitRule(listener: titeListener): void {
	    if(listener.exitBlock) {
	 		listener.exitBlock(this);
		}
	}
}


export class ArrayContext extends ParserRuleContext {
	constructor(parser?: titeParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public LF_list(): TerminalNode[] {
	    	return this.getTokens(titeParser.LF);
	}
	public LF(i: number): TerminalNode {
		return this.getToken(titeParser.LF, i);
	}
	public sequence(): SequenceContext {
		return this.getTypedRuleContext(SequenceContext, 0) as SequenceContext;
	}
    public get ruleIndex(): number {
    	return titeParser.RULE_array;
	}
	public enterRule(listener: titeListener): void {
	    if(listener.enterArray) {
	 		listener.enterArray(this);
		}
	}
	public exitRule(listener: titeListener): void {
	    if(listener.exitArray) {
	 		listener.exitArray(this);
		}
	}
}


export class InputContext extends ParserRuleContext {
	constructor(parser?: titeParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public LF_list(): TerminalNode[] {
	    	return this.getTokens(titeParser.LF);
	}
	public LF(i: number): TerminalNode {
		return this.getToken(titeParser.LF, i);
	}
	public sequence(): SequenceContext {
		return this.getTypedRuleContext(SequenceContext, 0) as SequenceContext;
	}
    public get ruleIndex(): number {
    	return titeParser.RULE_input;
	}
	public enterRule(listener: titeListener): void {
	    if(listener.enterInput) {
	 		listener.enterInput(this);
		}
	}
	public exitRule(listener: titeListener): void {
	    if(listener.exitInput) {
	 		listener.exitInput(this);
		}
	}
}


export class DelimContext extends ParserRuleContext {
	constructor(parser?: titeParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public LF_list(): TerminalNode[] {
	    	return this.getTokens(titeParser.LF);
	}
	public LF(i: number): TerminalNode {
		return this.getToken(titeParser.LF, i);
	}
    public get ruleIndex(): number {
    	return titeParser.RULE_delim;
	}
	public enterRule(listener: titeListener): void {
	    if(listener.enterDelim) {
	 		listener.enterDelim(this);
		}
	}
	public exitRule(listener: titeListener): void {
	    if(listener.exitDelim) {
	 		listener.exitDelim(this);
		}
	}
}
