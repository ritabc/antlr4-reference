grammar Rows;

@parser::members { // add members to generated RowsParser
    int col;
    public RowsParser(TokenStream input, int col) { // custom constructor
        this(input);
        this.col = col;
    }
}

file: (row NL)+;

row
	locals[int i=0]:
	(
		STUFF {
            $i++;
            if ( $i == col ) System.out.println($STUFF.text);
        }
	)+;

TAB: '\t' -> skip; // match but doesn't pass to parser
NL: '\r'? '\n'; // match and pass to the parser
STUFF: ~[\t\r\n]+; // match any chars except tab, newLine