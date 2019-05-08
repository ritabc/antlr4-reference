// Generated from /Users/ritabc/src/Projects/ANTLR/examples/JSON/JSON.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class JSONLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		WS=10, STRING=11, NUMBER=12;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	public static final String[] ruleNames = {
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
		"WS", "STRING", "NUMBER", "INT", "EXP", "ESC", "UNICODE", "HEX"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'{'", "','", "'}'", "':'", "'['", "']'", "'true'", "'false'", "'null'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, null, null, null, null, null, null, null, null, "WS", "STRING", 
		"NUMBER"
	};
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


	public JSONLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "JSON.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\16\u0088\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\7\3\7\3\b\3\b\3\b\3\b"+
		"\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\13\6\13C\n\13\r\13"+
		"\16\13D\3\13\3\13\3\f\3\f\3\f\7\fL\n\f\f\f\16\fO\13\f\3\f\3\f\3\r\5\r"+
		"T\n\r\3\r\3\r\3\r\6\rY\n\r\r\r\16\rZ\3\r\5\r^\n\r\3\r\5\ra\n\r\3\r\3\r"+
		"\3\r\3\r\5\rg\n\r\3\r\5\rj\n\r\3\16\3\16\3\16\7\16o\n\16\f\16\16\16r\13"+
		"\16\5\16t\n\16\3\17\3\17\5\17x\n\17\3\17\3\17\3\20\3\20\3\20\5\20\177"+
		"\n\20\3\21\3\21\3\21\3\21\3\21\3\21\3\22\3\22\2\2\23\3\3\5\4\7\5\t\6\13"+
		"\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\2\35\2\37\2!\2#\2\3\2\n\5\2"+
		"\13\f\17\17\"\"\4\2$$^^\3\2\62;\3\2\63;\4\2GGgg\4\2--//\n\2$$\61\61^^"+
		"ddhhppttvv\5\2\62;CHch\2\u0090\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t"+
		"\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2"+
		"\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\3%\3\2\2\2\5\'\3\2\2\2\7)\3"+
		"\2\2\2\t+\3\2\2\2\13-\3\2\2\2\r/\3\2\2\2\17\61\3\2\2\2\21\66\3\2\2\2\23"+
		"<\3\2\2\2\25B\3\2\2\2\27H\3\2\2\2\31i\3\2\2\2\33s\3\2\2\2\35u\3\2\2\2"+
		"\37{\3\2\2\2!\u0080\3\2\2\2#\u0086\3\2\2\2%&\7}\2\2&\4\3\2\2\2\'(\7.\2"+
		"\2(\6\3\2\2\2)*\7\177\2\2*\b\3\2\2\2+,\7<\2\2,\n\3\2\2\2-.\7]\2\2.\f\3"+
		"\2\2\2/\60\7_\2\2\60\16\3\2\2\2\61\62\7v\2\2\62\63\7t\2\2\63\64\7w\2\2"+
		"\64\65\7g\2\2\65\20\3\2\2\2\66\67\7h\2\2\678\7c\2\289\7n\2\29:\7u\2\2"+
		":;\7g\2\2;\22\3\2\2\2<=\7p\2\2=>\7w\2\2>?\7n\2\2?@\7n\2\2@\24\3\2\2\2"+
		"AC\t\2\2\2BA\3\2\2\2CD\3\2\2\2DB\3\2\2\2DE\3\2\2\2EF\3\2\2\2FG\b\13\2"+
		"\2G\26\3\2\2\2HM\7$\2\2IL\5\37\20\2JL\n\3\2\2KI\3\2\2\2KJ\3\2\2\2LO\3"+
		"\2\2\2MK\3\2\2\2MN\3\2\2\2NP\3\2\2\2OM\3\2\2\2PQ\7$\2\2Q\30\3\2\2\2RT"+
		"\7/\2\2SR\3\2\2\2ST\3\2\2\2TU\3\2\2\2UV\5\33\16\2VX\7\60\2\2WY\t\4\2\2"+
		"XW\3\2\2\2YZ\3\2\2\2ZX\3\2\2\2Z[\3\2\2\2[]\3\2\2\2\\^\5\35\17\2]\\\3\2"+
		"\2\2]^\3\2\2\2^j\3\2\2\2_a\7/\2\2`_\3\2\2\2`a\3\2\2\2ab\3\2\2\2bc\5\33"+
		"\16\2cd\5\35\17\2dj\3\2\2\2eg\7/\2\2fe\3\2\2\2fg\3\2\2\2gh\3\2\2\2hj\5"+
		"\33\16\2iS\3\2\2\2i`\3\2\2\2if\3\2\2\2j\32\3\2\2\2kt\7\62\2\2lp\t\5\2"+
		"\2mo\t\4\2\2nm\3\2\2\2or\3\2\2\2pn\3\2\2\2pq\3\2\2\2qt\3\2\2\2rp\3\2\2"+
		"\2sk\3\2\2\2sl\3\2\2\2t\34\3\2\2\2uw\t\6\2\2vx\t\7\2\2wv\3\2\2\2wx\3\2"+
		"\2\2xy\3\2\2\2yz\5\33\16\2z\36\3\2\2\2{~\7^\2\2|\177\t\b\2\2}\177\5!\21"+
		"\2~|\3\2\2\2~}\3\2\2\2\177 \3\2\2\2\u0080\u0081\7w\2\2\u0081\u0082\5#"+
		"\22\2\u0082\u0083\5#\22\2\u0083\u0084\5#\22\2\u0084\u0085\5#\22\2\u0085"+
		"\"\3\2\2\2\u0086\u0087\t\t\2\2\u0087$\3\2\2\2\20\2DKMSZ]`fipsw~\3\b\2"+
		"\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}