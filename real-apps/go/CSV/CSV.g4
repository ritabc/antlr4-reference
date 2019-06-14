grammar CSV;

file: hdr row+;
hdr: row;

row: fieldWWS (COMMA fieldWWS)* '\r'? '\n';

fieldWWS: (WS)* field (WS)*;

field: TEXT # text | STRING # string | # empty;

TEXT: ~[,\n\r"]+;
STRING: '"' ('""' | ~'"')* '"';
// Find out how to account for text in the double double quotes like the blah in: ,January,"blah""zippo"""

COMMA: ',';
WS: [ \t];