grammar PropertyFile;

file: prop+;
prop: ID '=' STRING '\n';

ID: LETTER (LETTER | DIGIT)*;
fragment LETTER: [a-zA-Z\u0080-\u00FF_];
fragment DIGIT: [0-9];
STRING: '"' ('""' | ~'"')* '"';