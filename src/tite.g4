grammar tite;

program : declaration (delim declaration)* ;
declaration : (tag? IDENTIFIER)? ':' (type | expression) ;
tag : '.'|'$'|'#' ;
parameter : LF? declaration (delim declaration)* LF? ;
function : '*'? '(' parameter? ')' type? ;
type : condition | function ;

primary : IDENTIFIER | literal | '(' expression ')' ;
postfix : primary | postfix ('++'|'--'|'.'IDENTIFIER|argument|array) ; // TODO: struct unpacking
unary : postfix | ('++'|'--'|'+'|'-'|'~'|'!'|'@'|tag) unary | '|' unary '|' ;
product : unary | product ('*'|'/'|'%'|'**'|'..') unary ;
additive : product | additive ('+'|'-') product ;
shift : additive | shift ('<<'|'>>') additive ;
relation : shift | relation ('<'|'>'|'<='|'>=') shift ;
equality : relation | equality ('=='|'!='|'==='|'!==') relation ;
and : equality | and '&' equality ;
xor : and | xor '^' and ;
or : xor | or '|' xor ;
logic_and : or | logic_and LF? '&&' LF? or ;
logic_or : logic_and | logic_or LF? '||' LF? logic_and ;
condition : logic_or | logic_or '?' LF? (expression | otherwise) ;
otherwise : expression? LF? ':' LF? condition ;
assignment : '='|'=>'|'*='|'/='|'%='|'+='|'-='|'<<='|'>>='|'&='|'^='|'|=' ;
expression : condition | type? assignment LF? expression ;

element : declaration | expression ;
sequence : LF? element (delim element)* LF? ;
argument : '(' sequence? ')' ;
structure : '{' sequence? '}' ;
array : '[' sequence? ']' ; // TODO: should limit to literal key & value?
literal : INT | FLOAT | CHAR | STRING | structure | array ;
delim: (LF | ',') LF? ;
IDENTIFIER : [a-zA-Z_][a-zA-Z_0-9]* ;
INT : [0-9]+ ;
FLOAT : [0-9]+ '.' [0-9]* ;
CHAR : '\'' . '\'' ;
STRING : '"' .*? '"' ;
WS : [ ;\t]+ -> skip ;
LF : [\r\n]+ ;