grammar tite;

program : sequence ;
sequence : declaration (delim declaration)* ;
declaration : (identifiers ':')? expression ;
identifiers : primary (',' primary)* ;
expression : '*'? (type | type? assign LF? expression) ;
type : condition | function ;
function : input type? ;

condition : disjunction | disjunction LF? '?' LF? (expression | otherwise) ;
otherwise : expression? LF? ':' LF? condition ;
disjunction : conjunction | disjunction LF? '||' LF? conjunction ;
conjunction : or | conjunction LF? '&&' LF? or ;
or : xor | or LF? '|' LF? xor ;
xor : and | xor LF? '^' LF? and ;
and : equality | and LF? '&' LF? equality ;
equality : relation | equality LF? ('=='|'!='|'==='|'!==') LF? relation ;
relation : shift | relation LF? ('<'|'>'|'<='|'>=') LF? shift ;
shift : additive | shift LF? ('<<'|'>>') LF? additive ;
additive : product | additive LF? ('+'|'-'|'..') LF? product ;
product : factor | product LF? ('*'|'/'|'%'|'//') LF? factor ;
factor : power | ('++'|'--'|'+'|'-'|'~'|'!') LF? factor | '|' factor '|' ;
power : postfix | postfix LF? '**' LF? factor ;
postfix : primary | postfix LF? ('++'|'--'|'.'IDENTIFIER|array|input) ;
primary : tag? IDENTIFIER | literal | '(' expression ')' ;

tag : '.' | '$' | '#' ;
assign : '='|'=>'|'<-'|'@'|'*='|'/='|'%='|'//='|'**='|'+='|'-='|'<<='|'>>='|'&='|'^='|'|=' ;
literal : INT | FLOAT | CHAR | STR | block | array ;
block : '{' LF? sequence? LF? '}' ;
array : '[' LF? sequence? LF? ']' ;
input : '(' LF? sequence? LF? ')' ;
delim: (LF | ',') LF? ;
IDENTIFIER : [a-zA-Z_][a-zA-Z_0-9]* ;
INT : [0-9]+ | '0x'[0-9a-fA-F]+ | ('0'|'0o')[0-7]+ | '0b'[01]+ ;
FLOAT : [0-9]+ '.' [0-9]* ;
CHAR : '\'' . '\'' ;
STR : '"' .*? '"' | '"""' LF? .*? (LF .*?)* '"""' | '`' .*? '`' ; // TODO: format string
WS : [ ;\t]+ -> skip ;
LF : [\r\n]+ ;