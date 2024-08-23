grammar tite;

program : sequence ;
sequence : declaration (delim declaration)* ;
declaration : (tag? primary ':' LF?)? expression ;
expression : '*'? (type | type? assign expression) ;
type : input type? | condition ;
input : array | tuple | macro ;

condition : disjunction | disjunction LF? '?' LF? (expression | otherwise) ;
otherwise : expression? LF? ':' LF? condition ;
disjunction : conjunction | disjunction LF? '||' LF? conjunction ;
conjunction : or | conjunction LF? '&&' LF? or ;
or : xor | or LF? '|' LF? xor ;
xor : and | xor LF? '^' LF? and ;
and : equality | and LF? '&' LF? equality ;
equality : relation | equality LF? ('=='|'!='|'==='|'!==') LF? relation ;
relation : shift | relation LF? ('<'|'>'|'<='|'>=') LF? shift ;
shift : sum | shift LF? ('<<'|'>>') LF? sum ;
sum : term | sum ('+'|'-') LF? term ;
term : factor | term LF? ('*'|'/'|'%'|'//') LF? factor ;
factor : power | ('++'|'--'|'+'|'-'|'~'|'!') LF? factor | '|' factor '|' ;
power : postfix | postfix LF? '**' LF? factor ;
postfix : range | postfix ('++'|'--'|access|input) ;
access : '?'? LF? '.' IDENTIFIER | '?' LF? '.' input ;
range : primary | primary? LF? ('..'|'..<') LF? primary ;
primary : IDENTIFIER | literal ;

tag : '.' | '$' | '#' ;
mutate : '*='|'/='|'%='|'//='|'**='|'+='|'-='|'<<='|'>>='|'&='|'^='|'|=' ;
assign : LF? ('='|'=>'|'<-'|'@'|mutate) LF? ;
literal : INT | FLOAT | CHAR | STR | block | array | tuple ;
block : '{' LF? sequence? LF? '}' ;
array : '[' LF? sequence? LF? ']' ;
tuple : '(' LF? sequence? LF? ')' ;
macro : '<' LF? sequence? LF? '>' ;
delim: (LF | ',') LF? ;
IDENTIFIER : [a-zA-Z_][a-zA-Z_0-9]* ;
INT : [0-9]+ | '0x'[0-9a-fA-F]+ | ('0'|'0o')[0-7]+ | '0b'[01]+ ;
FLOAT : [0-9]* '.' [0-9]+ ;
CHAR : '\'' . '\'' ;
STR : '"' .*? '"' | '"""' LF? .*? (LF .*?)* '"""' | '`' .*? '`' ; // TODO: format string
WS : [ ;\t]+ -> skip ;
LF : [\r\n]+ ;