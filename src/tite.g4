grammar tite;

program : declaration+ ;
declaration : tag? IDENTIFIER ':' (type | type? (assignment | function));
tag : '$' | '.' | '#' ;
type : IDENTIFIER | expression | return? parameters;
return : IDENTIFIER ;
expression : IDENTIFIER | constant | execution | access | structure | hashmap | '(' production ')' ;
assignment : '=' expression ;
function : '=>' (expression | production | '{' (declarations ',')? expressions? '}') ;
production : expression ('*'|'/'|'**'|'+') expression ;
execution : IDENTIFIER arguments;
access: IDENTIFIER '.' expression ;
declarations : declaration (',' declaration)* ;
expressions : expression (',' expression)* ;
parameters : '(' declarations? ')' ;
arguments : '(' (declarations | expressions)? ')' ;
structure : '{' (declarations | expressions)? '}' ;
hashmap : '[' expressions? ',' declarations? ']' ;
constant : INT | FLOAT | STRING;
IDENTIFIER : [a-zA-Z_][a-zA-Z_0-9]* ;
INT : [0-9]+ ;
FLOAT : [0-9]+ '.' [0-9]* ;
STRING : '"' .*? '"' ;
WS : [ \t]+ -> skip ;
// NL : [\r\n]+ -> ',' ;
// TODO: remove ',' on structure and function definition