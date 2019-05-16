// Generated from /Users/ritabc/src/Projects/ANTLR/decoupling/PropertyFile.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class PropertyFileLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, ID=3, STRING=4;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	public static final String[] ruleNames = {
		"T__0", "T__1", "ID", "LETTER", "DIGIT", "STRING"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'='", "'\n'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, null, "ID", "STRING"
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


	public PropertyFileLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "PropertyFile.g4"; }

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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\6*\b\1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\3\2\3\2\3\3\3\3\3\4\3\4\3\4\7\4"+
		"\27\n\4\f\4\16\4\32\13\4\3\5\3\5\3\6\3\6\3\7\3\7\3\7\3\7\7\7$\n\7\f\7"+
		"\16\7\'\13\7\3\7\3\7\2\2\b\3\3\5\4\7\5\t\2\13\2\r\6\3\2\5\6\2C\\aac|\u0082"+
		"\u0101\3\2\62;\3\2$$\2+\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\r\3\2\2"+
		"\2\3\17\3\2\2\2\5\21\3\2\2\2\7\23\3\2\2\2\t\33\3\2\2\2\13\35\3\2\2\2\r"+
		"\37\3\2\2\2\17\20\7?\2\2\20\4\3\2\2\2\21\22\7\f\2\2\22\6\3\2\2\2\23\30"+
		"\5\t\5\2\24\27\5\t\5\2\25\27\5\13\6\2\26\24\3\2\2\2\26\25\3\2\2\2\27\32"+
		"\3\2\2\2\30\26\3\2\2\2\30\31\3\2\2\2\31\b\3\2\2\2\32\30\3\2\2\2\33\34"+
		"\t\2\2\2\34\n\3\2\2\2\35\36\t\3\2\2\36\f\3\2\2\2\37%\7$\2\2 !\7$\2\2!"+
		"$\7$\2\2\"$\n\4\2\2# \3\2\2\2#\"\3\2\2\2$\'\3\2\2\2%#\3\2\2\2%&\3\2\2"+
		"\2&(\3\2\2\2\'%\3\2\2\2()\7$\2\2)\16\3\2\2\2\7\2\26\30#%\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}