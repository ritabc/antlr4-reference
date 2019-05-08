grammar R;

prog: ( expr_or_assign (';' | NL) | NL)* EOF;

expr_or_assign: expr ('<-' | '=' | '<<-') expr_or_assign | expr;

NL: '\r'? '\n';

expr:
	expr '[[' sublist ']' ']' // '[[' follows R's yacc grammar
	| expr '[' sublist ']'
	| expr ('::' | ':::') expr
	| expr ('$' | '@') expr
	| expr '^' <assoc = right> expr
	| ('-' | '+') expr
	| expr ':' expr
	| expr USER_OP expr // anything wrapped in %, ie: '%' .* '%'
	| expr ('*' | '/') expr
	| expr ('+' | '-') expr
	| expr ('>' | '>=' | '<' | '<=' | '==' | '!=') expr
	| '!' expr
	| expr ('&' | '&&') expr
	| expr ('|' | '||') expr
	| '~' expr
	| expr ('->' | '->>' | ':= ') expr
	| 'function' '(' formlist? ')' expr // define function with formal argument definition lists
	| expr '(' sublist ')' // call function with call site argument expressions
	| '{' exprlist '}' // compound statement
	| 'if' '(' expr ')' expr
	| 'if' '(' expr ')' expr 'else' expr
	| 'for' '(' ID 'in' expr ')' expr
	| 'while' '(' expr ')' expr
	| 'repeat' expr
	| '?' expr // get help on expr, usually string or ID
	| 'next'
	| 'break'
	| '(' expr ')'
	| ID
	| INT;

exprlist: expr_or_assign ((';' | NL) expr_or_assign)* |;
formlist: form (',' form)*;
form: ID | ID '=' expr | '...';
sublist: sub (',' sub)*;
sub:
	expr
	| ID '='
	| ID '=' expr
	| STRING '='
	| STRING '=' expr
	| 'NULL' '='
	| 'NULL' '=' expr
	| '...'
	|;

ID:
	'.' (LETTER | '_' | '.') (LETTER | DIGIT | '_' | '.')*
	| LETTER (LETTER | DIGIT | '_' | '.')*;

fragment LETTER: [a-zA-Z];
fragment DIGIT: [0-9];

STRING: '"' ('\\"' | .)*? '"';
USER_OP: '%' .*? '%';
WS: [ \t]+ -> skip;
INT: [0-9]+;