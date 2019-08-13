grammar Cymbol;

file: (functionDecl | varDecl)+;

varDecl: cymbolType ID ('=' expr)? ';';

cymbolType: 'float' | 'int' | 'void' ;

functionDecl: cymbolType ID '(' formalParameters? ')' block;

formalParameters: formalParameter (',' formalParameter)*;

formalParameter : cymbolType ID;

block: '{' stat* '}' ;

stat:
    block
    | varDecl
    | 'if' expr 'then' stat ('else' stat)?
    | 'return' expr? ';'
    | expr '=' expr ';'
    | expr ';' // function call
    ;

expr: ID '(' exprList? ')' # Call
    | expr '[' expr ']'    # Index
    | '-' expr             # Negate
    | '!' expr             # Not
    | expr '*' expr        # Mult
    | expr ('+'|'-') expr  # AddSub
    | expr '==' expr       # Equal
    | ID                   # Var
    | INT                  # Int
    | '(' expr ')'         # Parens
    ;

exprList: expr (',' expr)* ;

ID: LETTER (LETTER | [0-9])* ;
fragment LETTER : [a-zA-Z];

INT: [0-9]+;

WS : [ \t\r\n]+ -> skip;

SL_COMMENT : '//' .*? '\n' -> skip;