grammar Simple;

prog: classDef+; // match one or more class definitions

classDef
    :   'class' ID '{' member+ '}' // a class has one or more members
        { fmt.Println("class " + $ID.text) }
    ;

member
    :   'int' ID ';'                        // field definition
        { fmt.Println("var " + $ID.text) }
    |   'int' f=ID '(' ID ')' '{' stat '}'  // method definition
        { fmt.Println("method: " + $f.text) }
    ;

stat
    :   expr ';'          { fmt.Println("found expr: " + $expr.text) }
    |   ID '=' expr ';'   { fmt.Println("found assign: " + $expr.text) }
    ;

expr
    :   INT
    |   ID '(' INT ')'
    ;


INT : [0-9]+ ;
ID  : [a-zA-Z]+ ;
WS  : [ \t\r\n]+ -> skip;