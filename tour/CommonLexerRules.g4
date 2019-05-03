lexer grammar CommonLexerRules;

// match identifiers 
ID: [a-zA-Z]+;
// return newlines to parser (its end-statement signal) (I think if we come across a return)
NEWLINE: '\r'? '\n';
// match integers 
INT: [0-9]+;
// toss out whitespace
WS: [ \t]+ -> skip;