grammar JSON;

json: object | array;

object: '{' pair (',' pair)* '}' | '{' '}'; // empty object

pair: STRING ':' value;

array: '[' value (',' value)* ']' | '[' ']'; // empty array

value:
	STRING
	| NUMBER
	| object // recursion 
	| array // recursion
	| 'true' // keywords 
	| 'false' // keywords
	| 'null';

WS: [ \t\n\r]+ -> skip;
STRING: '"' (ESC | ~["\\])* '"';
NUMBER:
	'-'? INT '.' [0-9]+ EXP? // 1.35, 1.35E-9, 0.3, -4.5
	| '-'? INT EXP // 1E10, -3e4
	| '-'? INT;
// -3, 45

fragment INT: '0' | [1-9] [0-9]*;
// no leading zeros
fragment EXP: [Ee] [+\-]? INT;
// \- since - means 'range' inside [...]
fragment ESC: '\\' (["\\/bfnrt] | UNICODE);
fragment UNICODE: 'u' HEX HEX HEX HEX;
fragment HEX: [0-9a-fA-F];