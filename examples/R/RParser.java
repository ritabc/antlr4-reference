// Generated from R.g4 by ANTLR 4.7.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class RParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.7.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		T__24=25, T__25=26, T__26=27, T__27=28, T__28=29, T__29=30, T__30=31, 
		T__31=32, T__32=33, T__33=34, T__34=35, T__35=36, T__36=37, T__37=38, 
		T__38=39, T__39=40, T__40=41, T__41=42, T__42=43, T__43=44, T__44=45, 
		T__45=46, T__46=47, T__47=48, T__48=49, NL=50, ID=51, STRING=52, USER_OP=53, 
		WS=54, INT=55;
	public static final int
		RULE_prog = 0, RULE_expr_or_assign = 1, RULE_expr = 2, RULE_exprlist = 3, 
		RULE_formlist = 4, RULE_form = 5, RULE_sublist = 6, RULE_sub = 7;
	private static String[] makeRuleNames() {
		return new String[] {
			"prog", "expr_or_assign", "expr", "exprlist", "formlist", "form", "sublist", 
			"sub"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "';'", "'<-'", "'='", "'<<-'", "'[['", "']'", "'['", "'::'", "':::'", 
			"'$'", "'@'", "'^'", "'-'", "'+'", "':'", "'*'", "'/'", "'>'", "'>='", 
			"'<'", "'<='", "'=='", "'!='", "'!'", "'&'", "'&&'", "'|'", "'||'", "'~'", 
			"'->'", "'->>'", "':= '", "'function'", "'('", "')'", "'{'", "'}'", "'if'", 
			"'else'", "'for'", "'in'", "'while'", "'repeat'", "'?'", "'next'", "'break'", 
			"','", "'...'", "'NULL'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, "NL", "ID", "STRING", "USER_OP", "WS", "INT"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "R.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public RParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class ProgContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(RParser.EOF, 0); }
		public List<Expr_or_assignContext> expr_or_assign() {
			return getRuleContexts(Expr_or_assignContext.class);
		}
		public Expr_or_assignContext expr_or_assign(int i) {
			return getRuleContext(Expr_or_assignContext.class,i);
		}
		public List<TerminalNode> NL() { return getTokens(RParser.NL); }
		public TerminalNode NL(int i) {
			return getToken(RParser.NL, i);
		}
		public ProgContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prog; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof RListener ) ((RListener)listener).enterProg(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof RListener ) ((RListener)listener).exitProg(this);
		}
	}

	public final ProgContext prog() throws RecognitionException {
		ProgContext _localctx = new ProgContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_prog);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(22);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__12) | (1L << T__13) | (1L << T__23) | (1L << T__28) | (1L << T__32) | (1L << T__33) | (1L << T__35) | (1L << T__37) | (1L << T__39) | (1L << T__41) | (1L << T__42) | (1L << T__43) | (1L << T__44) | (1L << T__45) | (1L << NL) | (1L << ID) | (1L << INT))) != 0)) {
				{
				setState(20);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case T__12:
				case T__13:
				case T__23:
				case T__28:
				case T__32:
				case T__33:
				case T__35:
				case T__37:
				case T__39:
				case T__41:
				case T__42:
				case T__43:
				case T__44:
				case T__45:
				case ID:
				case INT:
					{
					setState(16);
					expr_or_assign();
					setState(17);
					_la = _input.LA(1);
					if ( !(_la==T__0 || _la==NL) ) {
					_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					}
					break;
				case NL:
					{
					setState(19);
					match(NL);
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
				setState(24);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(25);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Expr_or_assignContext extends ParserRuleContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public Expr_or_assignContext expr_or_assign() {
			return getRuleContext(Expr_or_assignContext.class,0);
		}
		public Expr_or_assignContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr_or_assign; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof RListener ) ((RListener)listener).enterExpr_or_assign(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof RListener ) ((RListener)listener).exitExpr_or_assign(this);
		}
	}

	public final Expr_or_assignContext expr_or_assign() throws RecognitionException {
		Expr_or_assignContext _localctx = new Expr_or_assignContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_expr_or_assign);
		int _la;
		try {
			setState(32);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(27);
				expr(0);
				setState(28);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__1) | (1L << T__2) | (1L << T__3))) != 0)) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(29);
				expr_or_assign();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(31);
				expr(0);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExprContext extends ParserRuleContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public FormlistContext formlist() {
			return getRuleContext(FormlistContext.class,0);
		}
		public ExprlistContext exprlist() {
			return getRuleContext(ExprlistContext.class,0);
		}
		public TerminalNode ID() { return getToken(RParser.ID, 0); }
		public TerminalNode INT() { return getToken(RParser.INT, 0); }
		public TerminalNode USER_OP() { return getToken(RParser.USER_OP, 0); }
		public SublistContext sublist() {
			return getRuleContext(SublistContext.class,0);
		}
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof RListener ) ((RListener)listener).enterExpr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof RListener ) ((RListener)listener).exitExpr(this);
		}
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 4;
		enterRecursionRule(_localctx, 4, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(92);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				{
				setState(35);
				_la = _input.LA(1);
				if ( !(_la==T__12 || _la==T__13) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(36);
				expr(25);
				}
				break;
			case 2:
				{
				setState(37);
				match(T__23);
				setState(38);
				expr(19);
				}
				break;
			case 3:
				{
				setState(39);
				match(T__28);
				setState(40);
				expr(16);
				}
				break;
			case 4:
				{
				setState(41);
				match(T__32);
				setState(42);
				match(T__33);
				setState(44);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__47 || _la==ID) {
					{
					setState(43);
					formlist();
					}
				}

				setState(46);
				match(T__34);
				setState(47);
				expr(14);
				}
				break;
			case 5:
				{
				setState(48);
				match(T__35);
				setState(49);
				exprlist();
				setState(50);
				match(T__36);
				}
				break;
			case 6:
				{
				setState(52);
				match(T__37);
				setState(53);
				match(T__33);
				setState(54);
				expr(0);
				setState(55);
				match(T__34);
				setState(56);
				expr(11);
				}
				break;
			case 7:
				{
				setState(58);
				match(T__37);
				setState(59);
				match(T__33);
				setState(60);
				expr(0);
				setState(61);
				match(T__34);
				setState(62);
				expr(0);
				setState(63);
				match(T__38);
				setState(64);
				expr(10);
				}
				break;
			case 8:
				{
				setState(66);
				match(T__39);
				setState(67);
				match(T__33);
				setState(68);
				match(ID);
				setState(69);
				match(T__40);
				setState(70);
				expr(0);
				setState(71);
				match(T__34);
				setState(72);
				expr(9);
				}
				break;
			case 9:
				{
				setState(74);
				match(T__41);
				setState(75);
				match(T__33);
				setState(76);
				expr(0);
				setState(77);
				match(T__34);
				setState(78);
				expr(8);
				}
				break;
			case 10:
				{
				setState(80);
				match(T__42);
				setState(81);
				expr(7);
				}
				break;
			case 11:
				{
				setState(82);
				match(T__43);
				setState(83);
				expr(6);
				}
				break;
			case 12:
				{
				setState(84);
				match(T__44);
				}
				break;
			case 13:
				{
				setState(85);
				match(T__45);
				}
				break;
			case 14:
				{
				setState(86);
				match(T__33);
				setState(87);
				expr(0);
				setState(88);
				match(T__34);
				}
				break;
			case 15:
				{
				setState(90);
				match(ID);
				}
				break;
			case 16:
				{
				setState(91);
				match(INT);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(145);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(143);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(94);
						if (!(precpred(_ctx, 28))) throw new FailedPredicateException(this, "precpred(_ctx, 28)");
						setState(95);
						_la = _input.LA(1);
						if ( !(_la==T__7 || _la==T__8) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(96);
						expr(29);
						}
						break;
					case 2:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(97);
						if (!(precpred(_ctx, 27))) throw new FailedPredicateException(this, "precpred(_ctx, 27)");
						setState(98);
						_la = _input.LA(1);
						if ( !(_la==T__9 || _la==T__10) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(99);
						expr(28);
						}
						break;
					case 3:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(100);
						if (!(precpred(_ctx, 26))) throw new FailedPredicateException(this, "precpred(_ctx, 26)");
						setState(101);
						match(T__11);
						setState(102);
						expr(27);
						}
						break;
					case 4:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(103);
						if (!(precpred(_ctx, 24))) throw new FailedPredicateException(this, "precpred(_ctx, 24)");
						setState(104);
						match(T__14);
						setState(105);
						expr(25);
						}
						break;
					case 5:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(106);
						if (!(precpred(_ctx, 23))) throw new FailedPredicateException(this, "precpred(_ctx, 23)");
						setState(107);
						match(USER_OP);
						setState(108);
						expr(24);
						}
						break;
					case 6:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(109);
						if (!(precpred(_ctx, 22))) throw new FailedPredicateException(this, "precpred(_ctx, 22)");
						setState(110);
						_la = _input.LA(1);
						if ( !(_la==T__15 || _la==T__16) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(111);
						expr(23);
						}
						break;
					case 7:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(112);
						if (!(precpred(_ctx, 21))) throw new FailedPredicateException(this, "precpred(_ctx, 21)");
						setState(113);
						_la = _input.LA(1);
						if ( !(_la==T__12 || _la==T__13) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(114);
						expr(22);
						}
						break;
					case 8:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(115);
						if (!(precpred(_ctx, 20))) throw new FailedPredicateException(this, "precpred(_ctx, 20)");
						setState(116);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__17) | (1L << T__18) | (1L << T__19) | (1L << T__20) | (1L << T__21) | (1L << T__22))) != 0)) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(117);
						expr(21);
						}
						break;
					case 9:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(118);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(119);
						_la = _input.LA(1);
						if ( !(_la==T__24 || _la==T__25) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(120);
						expr(19);
						}
						break;
					case 10:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(121);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(122);
						_la = _input.LA(1);
						if ( !(_la==T__26 || _la==T__27) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(123);
						expr(18);
						}
						break;
					case 11:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(124);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(125);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__29) | (1L << T__30) | (1L << T__31))) != 0)) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(126);
						expr(16);
						}
						break;
					case 12:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(127);
						if (!(precpred(_ctx, 30))) throw new FailedPredicateException(this, "precpred(_ctx, 30)");
						setState(128);
						match(T__4);
						setState(129);
						sublist();
						setState(130);
						match(T__5);
						setState(131);
						match(T__5);
						}
						break;
					case 13:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(133);
						if (!(precpred(_ctx, 29))) throw new FailedPredicateException(this, "precpred(_ctx, 29)");
						setState(134);
						match(T__6);
						setState(135);
						sublist();
						setState(136);
						match(T__5);
						}
						break;
					case 14:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(138);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(139);
						match(T__33);
						setState(140);
						sublist();
						setState(141);
						match(T__34);
						}
						break;
					}
					} 
				}
				setState(147);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class ExprlistContext extends ParserRuleContext {
		public List<Expr_or_assignContext> expr_or_assign() {
			return getRuleContexts(Expr_or_assignContext.class);
		}
		public Expr_or_assignContext expr_or_assign(int i) {
			return getRuleContext(Expr_or_assignContext.class,i);
		}
		public List<TerminalNode> NL() { return getTokens(RParser.NL); }
		public TerminalNode NL(int i) {
			return getToken(RParser.NL, i);
		}
		public ExprlistContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_exprlist; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof RListener ) ((RListener)listener).enterExprlist(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof RListener ) ((RListener)listener).exitExprlist(this);
		}
	}

	public final ExprlistContext exprlist() throws RecognitionException {
		ExprlistContext _localctx = new ExprlistContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_exprlist);
		int _la;
		try {
			setState(157);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__12:
			case T__13:
			case T__23:
			case T__28:
			case T__32:
			case T__33:
			case T__35:
			case T__37:
			case T__39:
			case T__41:
			case T__42:
			case T__43:
			case T__44:
			case T__45:
			case ID:
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(148);
				expr_or_assign();
				setState(153);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__0 || _la==NL) {
					{
					{
					setState(149);
					_la = _input.LA(1);
					if ( !(_la==T__0 || _la==NL) ) {
					_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					setState(150);
					expr_or_assign();
					}
					}
					setState(155);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				break;
			case T__36:
				enterOuterAlt(_localctx, 2);
				{
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FormlistContext extends ParserRuleContext {
		public List<FormContext> form() {
			return getRuleContexts(FormContext.class);
		}
		public FormContext form(int i) {
			return getRuleContext(FormContext.class,i);
		}
		public FormlistContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_formlist; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof RListener ) ((RListener)listener).enterFormlist(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof RListener ) ((RListener)listener).exitFormlist(this);
		}
	}

	public final FormlistContext formlist() throws RecognitionException {
		FormlistContext _localctx = new FormlistContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_formlist);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(159);
			form();
			setState(164);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__46) {
				{
				{
				setState(160);
				match(T__46);
				setState(161);
				form();
				}
				}
				setState(166);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FormContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(RParser.ID, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public FormContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_form; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof RListener ) ((RListener)listener).enterForm(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof RListener ) ((RListener)listener).exitForm(this);
		}
	}

	public final FormContext form() throws RecognitionException {
		FormContext _localctx = new FormContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_form);
		try {
			setState(172);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(167);
				match(ID);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(168);
				match(ID);
				setState(169);
				match(T__2);
				setState(170);
				expr(0);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(171);
				match(T__47);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SublistContext extends ParserRuleContext {
		public List<SubContext> sub() {
			return getRuleContexts(SubContext.class);
		}
		public SubContext sub(int i) {
			return getRuleContext(SubContext.class,i);
		}
		public SublistContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sublist; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof RListener ) ((RListener)listener).enterSublist(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof RListener ) ((RListener)listener).exitSublist(this);
		}
	}

	public final SublistContext sublist() throws RecognitionException {
		SublistContext _localctx = new SublistContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_sublist);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(174);
			sub();
			setState(179);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__46) {
				{
				{
				setState(175);
				match(T__46);
				setState(176);
				sub();
				}
				}
				setState(181);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SubContext extends ParserRuleContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode ID() { return getToken(RParser.ID, 0); }
		public TerminalNode STRING() { return getToken(RParser.STRING, 0); }
		public SubContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sub; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof RListener ) ((RListener)listener).enterSub(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof RListener ) ((RListener)listener).exitSub(this);
		}
	}

	public final SubContext sub() throws RecognitionException {
		SubContext _localctx = new SubContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_sub);
		try {
			setState(200);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(182);
				expr(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(183);
				match(ID);
				setState(184);
				match(T__2);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(185);
				match(ID);
				setState(186);
				match(T__2);
				setState(187);
				expr(0);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(188);
				match(STRING);
				setState(189);
				match(T__2);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(190);
				match(STRING);
				setState(191);
				match(T__2);
				setState(192);
				expr(0);
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(193);
				match(T__48);
				setState(194);
				match(T__2);
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(195);
				match(T__48);
				setState(196);
				match(T__2);
				setState(197);
				expr(0);
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(198);
				match(T__47);
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 2:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 28);
		case 1:
			return precpred(_ctx, 27);
		case 2:
			return precpred(_ctx, 26);
		case 3:
			return precpred(_ctx, 24);
		case 4:
			return precpred(_ctx, 23);
		case 5:
			return precpred(_ctx, 22);
		case 6:
			return precpred(_ctx, 21);
		case 7:
			return precpred(_ctx, 20);
		case 8:
			return precpred(_ctx, 18);
		case 9:
			return precpred(_ctx, 17);
		case 10:
			return precpred(_ctx, 15);
		case 11:
			return precpred(_ctx, 30);
		case 12:
			return precpred(_ctx, 29);
		case 13:
			return precpred(_ctx, 13);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\39\u00cd\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\3\2\3\2\3\2\3\2"+
		"\7\2\27\n\2\f\2\16\2\32\13\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\5\3#\n\3\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4/\n\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\5\4_\n\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\7\4\u0092\n\4\f\4\16\4\u0095\13\4\3\5\3\5\3\5\7\5\u009a"+
		"\n\5\f\5\16\5\u009d\13\5\3\5\5\5\u00a0\n\5\3\6\3\6\3\6\7\6\u00a5\n\6\f"+
		"\6\16\6\u00a8\13\6\3\7\3\7\3\7\3\7\3\7\5\7\u00af\n\7\3\b\3\b\3\b\7\b\u00b4"+
		"\n\b\f\b\16\b\u00b7\13\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\t\5\t\u00cb\n\t\3\t\2\3\6\n\2\4\6\b\n\f\16\20"+
		"\2\f\4\2\3\3\64\64\3\2\4\6\3\2\17\20\3\2\n\13\3\2\f\r\3\2\22\23\3\2\24"+
		"\31\3\2\33\34\3\2\35\36\3\2 \"\2\u00f3\2\30\3\2\2\2\4\"\3\2\2\2\6^\3\2"+
		"\2\2\b\u009f\3\2\2\2\n\u00a1\3\2\2\2\f\u00ae\3\2\2\2\16\u00b0\3\2\2\2"+
		"\20\u00ca\3\2\2\2\22\23\5\4\3\2\23\24\t\2\2\2\24\27\3\2\2\2\25\27\7\64"+
		"\2\2\26\22\3\2\2\2\26\25\3\2\2\2\27\32\3\2\2\2\30\26\3\2\2\2\30\31\3\2"+
		"\2\2\31\33\3\2\2\2\32\30\3\2\2\2\33\34\7\2\2\3\34\3\3\2\2\2\35\36\5\6"+
		"\4\2\36\37\t\3\2\2\37 \5\4\3\2 #\3\2\2\2!#\5\6\4\2\"\35\3\2\2\2\"!\3\2"+
		"\2\2#\5\3\2\2\2$%\b\4\1\2%&\t\4\2\2&_\5\6\4\33\'(\7\32\2\2(_\5\6\4\25"+
		")*\7\37\2\2*_\5\6\4\22+,\7#\2\2,.\7$\2\2-/\5\n\6\2.-\3\2\2\2./\3\2\2\2"+
		"/\60\3\2\2\2\60\61\7%\2\2\61_\5\6\4\20\62\63\7&\2\2\63\64\5\b\5\2\64\65"+
		"\7\'\2\2\65_\3\2\2\2\66\67\7(\2\2\678\7$\2\289\5\6\4\29:\7%\2\2:;\5\6"+
		"\4\r;_\3\2\2\2<=\7(\2\2=>\7$\2\2>?\5\6\4\2?@\7%\2\2@A\5\6\4\2AB\7)\2\2"+
		"BC\5\6\4\fC_\3\2\2\2DE\7*\2\2EF\7$\2\2FG\7\65\2\2GH\7+\2\2HI\5\6\4\2I"+
		"J\7%\2\2JK\5\6\4\13K_\3\2\2\2LM\7,\2\2MN\7$\2\2NO\5\6\4\2OP\7%\2\2PQ\5"+
		"\6\4\nQ_\3\2\2\2RS\7-\2\2S_\5\6\4\tTU\7.\2\2U_\5\6\4\bV_\7/\2\2W_\7\60"+
		"\2\2XY\7$\2\2YZ\5\6\4\2Z[\7%\2\2[_\3\2\2\2\\_\7\65\2\2]_\79\2\2^$\3\2"+
		"\2\2^\'\3\2\2\2^)\3\2\2\2^+\3\2\2\2^\62\3\2\2\2^\66\3\2\2\2^<\3\2\2\2"+
		"^D\3\2\2\2^L\3\2\2\2^R\3\2\2\2^T\3\2\2\2^V\3\2\2\2^W\3\2\2\2^X\3\2\2\2"+
		"^\\\3\2\2\2^]\3\2\2\2_\u0093\3\2\2\2`a\f\36\2\2ab\t\5\2\2b\u0092\5\6\4"+
		"\37cd\f\35\2\2de\t\6\2\2e\u0092\5\6\4\36fg\f\34\2\2gh\7\16\2\2h\u0092"+
		"\5\6\4\35ij\f\32\2\2jk\7\21\2\2k\u0092\5\6\4\33lm\f\31\2\2mn\7\67\2\2"+
		"n\u0092\5\6\4\32op\f\30\2\2pq\t\7\2\2q\u0092\5\6\4\31rs\f\27\2\2st\t\4"+
		"\2\2t\u0092\5\6\4\30uv\f\26\2\2vw\t\b\2\2w\u0092\5\6\4\27xy\f\24\2\2y"+
		"z\t\t\2\2z\u0092\5\6\4\25{|\f\23\2\2|}\t\n\2\2}\u0092\5\6\4\24~\177\f"+
		"\21\2\2\177\u0080\t\13\2\2\u0080\u0092\5\6\4\22\u0081\u0082\f \2\2\u0082"+
		"\u0083\7\7\2\2\u0083\u0084\5\16\b\2\u0084\u0085\7\b\2\2\u0085\u0086\7"+
		"\b\2\2\u0086\u0092\3\2\2\2\u0087\u0088\f\37\2\2\u0088\u0089\7\t\2\2\u0089"+
		"\u008a\5\16\b\2\u008a\u008b\7\b\2\2\u008b\u0092\3\2\2\2\u008c\u008d\f"+
		"\17\2\2\u008d\u008e\7$\2\2\u008e\u008f\5\16\b\2\u008f\u0090\7%\2\2\u0090"+
		"\u0092\3\2\2\2\u0091`\3\2\2\2\u0091c\3\2\2\2\u0091f\3\2\2\2\u0091i\3\2"+
		"\2\2\u0091l\3\2\2\2\u0091o\3\2\2\2\u0091r\3\2\2\2\u0091u\3\2\2\2\u0091"+
		"x\3\2\2\2\u0091{\3\2\2\2\u0091~\3\2\2\2\u0091\u0081\3\2\2\2\u0091\u0087"+
		"\3\2\2\2\u0091\u008c\3\2\2\2\u0092\u0095\3\2\2\2\u0093\u0091\3\2\2\2\u0093"+
		"\u0094\3\2\2\2\u0094\7\3\2\2\2\u0095\u0093\3\2\2\2\u0096\u009b\5\4\3\2"+
		"\u0097\u0098\t\2\2\2\u0098\u009a\5\4\3\2\u0099\u0097\3\2\2\2\u009a\u009d"+
		"\3\2\2\2\u009b\u0099\3\2\2\2\u009b\u009c\3\2\2\2\u009c\u00a0\3\2\2\2\u009d"+
		"\u009b\3\2\2\2\u009e\u00a0\3\2\2\2\u009f\u0096\3\2\2\2\u009f\u009e\3\2"+
		"\2\2\u00a0\t\3\2\2\2\u00a1\u00a6\5\f\7\2\u00a2\u00a3\7\61\2\2\u00a3\u00a5"+
		"\5\f\7\2\u00a4\u00a2\3\2\2\2\u00a5\u00a8\3\2\2\2\u00a6\u00a4\3\2\2\2\u00a6"+
		"\u00a7\3\2\2\2\u00a7\13\3\2\2\2\u00a8\u00a6\3\2\2\2\u00a9\u00af\7\65\2"+
		"\2\u00aa\u00ab\7\65\2\2\u00ab\u00ac\7\5\2\2\u00ac\u00af\5\6\4\2\u00ad"+
		"\u00af\7\62\2\2\u00ae\u00a9\3\2\2\2\u00ae\u00aa\3\2\2\2\u00ae\u00ad\3"+
		"\2\2\2\u00af\r\3\2\2\2\u00b0\u00b5\5\20\t\2\u00b1\u00b2\7\61\2\2\u00b2"+
		"\u00b4\5\20\t\2\u00b3\u00b1\3\2\2\2\u00b4\u00b7\3\2\2\2\u00b5\u00b3\3"+
		"\2\2\2\u00b5\u00b6\3\2\2\2\u00b6\17\3\2\2\2\u00b7\u00b5\3\2\2\2\u00b8"+
		"\u00cb\5\6\4\2\u00b9\u00ba\7\65\2\2\u00ba\u00cb\7\5\2\2\u00bb\u00bc\7"+
		"\65\2\2\u00bc\u00bd\7\5\2\2\u00bd\u00cb\5\6\4\2\u00be\u00bf\7\66\2\2\u00bf"+
		"\u00cb\7\5\2\2\u00c0\u00c1\7\66\2\2\u00c1\u00c2\7\5\2\2\u00c2\u00cb\5"+
		"\6\4\2\u00c3\u00c4\7\63\2\2\u00c4\u00cb\7\5\2\2\u00c5\u00c6\7\63\2\2\u00c6"+
		"\u00c7\7\5\2\2\u00c7\u00cb\5\6\4\2\u00c8\u00cb\7\62\2\2\u00c9\u00cb\3"+
		"\2\2\2\u00ca\u00b8\3\2\2\2\u00ca\u00b9\3\2\2\2\u00ca\u00bb\3\2\2\2\u00ca"+
		"\u00be\3\2\2\2\u00ca\u00c0\3\2\2\2\u00ca\u00c3\3\2\2\2\u00ca\u00c5\3\2"+
		"\2\2\u00ca\u00c8\3\2\2\2\u00ca\u00c9\3\2\2\2\u00cb\21\3\2\2\2\17\26\30"+
		"\".^\u0091\u0093\u009b\u009f\u00a6\u00ae\u00b5\u00ca";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}