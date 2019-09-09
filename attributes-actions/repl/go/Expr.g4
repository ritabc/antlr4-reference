grammar Expr;
import ExprLexerGrammar;

@header {
    import (
    )
}

@parser::members {
   /** "memory" for our calculator; variable/value pairs go here */
   var memory = map[string]int{}
   func eval(left, op, right int) int {
       switch op {
       case ExprParserMULT:
           return left * right
       case ExprParserDIV:
           return left / right
       case ExprParserPLUS:
           return left + right
       case ExprParserMINUS:
           return left - right
       default:
           return 0
       }
   }
}

stat : e NEWLINE {fmt.Println($e.v)}
    | ID '=' e NEWLINE {memory[$ID.text] = $e.v}
    | NEWLINE ;


e returns [int v]
    : a=e op=('*'|'/') b=e {$v = eval($a.v, $op.type, $b.v)}
    | a=e op=('+'|'-') b=e {$v = eval($a.v, $op.type, $b.v)}
    | INT                  {$v = $INT.int}
    | ID                   {
                                var id = $ID.text
                                value, found := memory[id]
                                if found {
                                    $v = value
                                } else {
                                    $v = 0
                                }
                           }
    | '(' e ')'            {$v = $e.v}
    ;

ID  :   [a-zA-Z]+ ;      // match identifiers
INT :   [0-9]+ ;         // match integers
NEWLINE:'\r'? '\n' ;     // return newlines to parser (is end-statement signal)
WS  :   [ \t]+ -> skip ; // toss out whitespace

