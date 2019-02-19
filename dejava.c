#include"dejava.h"

bool_t boolConversion(char* statement){
	char str[256];
	char*tok;
	char op[3] = "  ";
	char* symbols[4]; //necessary?
	void* context = &idx; 
	bool_t result; //first token
	memcpy(str, statement, strlen(statement)+1);
	
	if( strstr(str,"=") ){
		op[1] = '=';
	}
	tok = strtok(str, DELIM); //idx
	symbols[0] = tok;
	op[0] = str[strlen(tok)+1];// get a op after null. duplicate =.
	
	while( True ){
		tok = strtok(NULL, DELIM);
		if ( !tok ) return True;
		
		if( parseOperand(&result, op, tok) != True ){
			symbols[0] = tok; //to get 'not or' 'not and'? verify usage
			if( !strcmp( tok, "or" ) ){
				strtok(NULL, DELIM);
				result = parseOperand(context, op, tok)? result || context : result; 
				continue;
			}
			return False; //wrong expression.
		}
		//TODO: multiple statements. of | and &, is and or, xor, not something. 
		// or if(a<10) && if(b> 10)? how bout it?
	}
}//TODO: way to call a symbol by string name. replace idx.

bool_t parseOperand(const void* context, char op[], char*tok ){
	int result = *(int*)context;
	int value = atoi(tok); // except number 0. faults if tok is NULL
	
	if( value || strcmp(tok,"0") == 0 ) {
		//if we can transfile, just return here a string of ("%s %s %d", symbol, op, result);
		switch(op[0]){
			case '!':
				result = idx != value;
			case '>':
				result = (op[1]=='=')? result >= value : result > value;
			case '<':
				result = (op[1]=='=')? result <= value : result < value;
			case '=':
				result = idx == value;
		}
		println("idx%s%s ? %s",op,tok,result?"True":"False");
	} else {
		return False;
	}
	
	return True;
}




