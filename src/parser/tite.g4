grammar tite;

program : sequence ;
sequence : declaration (delim declaration)* ;
declaration : (tag? primary ':' LF?)? expression ;
expression : '*'? (set? access expression | set) ;
set : input set? | condition ;

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
term : factor | term ('*'|'/'|'%'|'//') LF? factor ;
factor : power | ('++'|'--'|'+'|'-'|'~'|'!') LF? factor | '|' factor '|' ;
power : postfix | postfix LF? '**' LF? factor ;
postfix : range | postfix ('++'|'--'|input) ;
range : primary | primary? LF? ('..'|'..<') LF? primary ;
primary : IDENTIFIER | literal ;

tag : '.' | '$' | '#' ;
assign : '='|'=>'|'<-'|'*='|'/='|'%='|'//='|'**='|'+='|'-='|'<<='|'>>='|'&='|'^='|'|=' ;
access : (assign | LF? '.' | '@') LF? ;
literal : INT | FLOAT | CHAR | STR | block | array | tuple ;
input : array | tuple | macro ;
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