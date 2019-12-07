#include"dejava.h"

bool_t boolConversion(char* statement){
	char str[256];
	char*tok;
	char op[3] = "  ";
	int result;
	memcpy(str, statement, strlen(statement)+1);
	
	if( strstr(str,"=") ){
		op[1] = '=';
	}
	tok = strtok(str, DELIM);
	op[0] = str[strlen(tok)+1];// get a op after null. duplicate =.
	
	while( tok ){
		tok = strtok(NULL, DELIM);
		result = atoi(tok); // except number 0
		
		if( result || strcmp(tok,"0") == 0 ) {
			//if we can transfile, just return here a string of ("idx %s %d", op, result);
			switch(op[0]){
				case '!':
					return idx != result;
				case '>':
					return (op[1]=='=')? idx >= result : idx > result;
				case '<':
					return (op[1]=='=')? idx <= result : idx < result;
				case '=':
					return idx == result;
			}
		} else {
			return False; //wrong expression.
			continue;
		}
		//TODO: multiple statements. of | and &, is and or, xor, not something. 
		// or if(a<10) && if(b> 10)? how bout it?
	}
}//TODO: way to call a symbol by string name.