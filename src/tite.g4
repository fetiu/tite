grammar tite;

program : declaration (delim declaration)* ;
identifiers : primary (',' primary)* ;
declaration : (identifiers ':'| identifiers? '@') (type | expression) ;
parameter : declaration (delim declaration)* ;
function : '(' LF? parameter? LF? ')' type? ;
type : condition | function ;
tag : '.' | '$' | '#' ;

primary : tag? IDENTIFIER | literal | '(' expression ')' ;
postfix : primary | postfix ('++'|'--'|'.'IDENTIFIER|argument|brackets) ;
power : postfix | postfix '**' factor ;
factor : power | ('++'|'--'|'+'|'-'|'~'|'!') factor | '|' factor '|' ;
product : factor | product ('*'|'/'|'%'|'//') factor ;
additive : product | additive ('+'|'-'|'..') product ;
shift : additive | shift ('<<'|'>>') additive ;
relation : shift | relation ('<'|'>'|'<='|'>=') shift ;
equality : relation | equality ('=='|'!='|'==='|'!==') relation ;
and : equality | and '&' equality ;
xor : and | xor '^' and ;
or : xor | or '|' xor ;
conjunction : or | conjunction LF? '&&' LF? or ;
disjunction : conjunction | disjunction LF? '||' LF? conjunction ;
condition : disjunction | disjunction '?' (expression | otherwise) ;
otherwise : expression? ':' condition ;
assignment : '='|'=>'|'<-'|'*='|'/='|'%='|'//='|'**='|'+='|'-='|'<<='|'>>='|'&='|'^='|'|=' ;
expression : '*'? (condition | type? assignment LF? expression) ;

element : declaration | expression ;
sequence : element (delim element)* ;
compound : '{' LF? sequence? LF? '}' ;
argument : '(' LF? sequence? LF? ')' ;
brackets : '[' LF? sequence? LF? ']' ; // TODO: should limit to literal key & value?
literal : INT | FLOAT | CHAR | STR | compound | brackets ;
delim: (LF | ',') LF? ;
IDENTIFIER : [a-zA-Z_][a-zA-Z_0-9]* ;
INT : [0-9]+ | '0x'[0-9a-fA-F]+ | ('0'|'0o')[0-7]+ | '0b'[01]+ ;
FLOAT : [0-9]+ '.' [0-9]* ;
CHAR : '\'' . '\'' ;
STR : '"' .*? '"' | '"""' LF? .*? (LF .*?)* '"""' | '`' .*? '`' ; // TODO: format string
WS : [ ;\t]+ -> skip ;
LF : [\r\n]+ ;
