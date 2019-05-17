grammar Calc;

stmt: expr EOF;

expr:
	expr op = (MUL | DIV) expr		# MulDiv
	| expr op = ( ADD | SUB) expr	# AddSub
	| NUMBER						# Number;

MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
NUMBER: [0-9]+;
WHITESPACE: [ \n\r\t]+ -> skip;