grammar JSON;

json: object | array;

object:
	'{' pair (',' pair)* '}'	# AnObject
	| '{' '}'					# BlankObject;

array:
	'[' value (',' value)* ']'	# ArrayOfValues
	| '[' ']'					# BlankArray;

value:
	STRING		# String
	| NUMBER	# Atom
	| object	# ObjectValue
	| array		# ArrayValue
	| 'true'	# Atom
	| 'false'	# Atom
	| 'null'	# Atom;

pair: STRING ':' value;

STRING: '"' (~["\\])* '"';

NUMBER: '-'? INT '.' [0-9]+ | '-'? INT;

fragment INT: '0' | [1-9] [0-9]*;

WS: [ \t\n\r]+ -> skip;
