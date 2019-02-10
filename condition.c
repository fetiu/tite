/*
	Prototype of Condition Statement rule using Ternary operator
	Property of the Project Dejava. Copyright 2019, Jake Fetiu.	
*/

#include<stdio.h> 
#include<string.h>
#include<stdlib.h>
#define println(fmt,...) printf(fmt"\n", ##__VA_ARGS__)
#define If(x) boolConversion(#x"?")
#define Else :
#define For(i,d,x) idx = (int)i-x; while( d != (idx+=x) )
#define DELIM "=!<>? "

static int idx;
typedef enum { False, True } bool_t;
bool_t boolConversion(char* statement);

int main(){
	/* basic condition statement. notice that there's no if */
	bool_t dejava = False;
	
	For( 0, 6, +1 ){ //iterates 0->5
		idx == 0 || idx == 4?  
			println("hello world")
		Else idx <= 3 ? 
			println("hello friend %d", idx)
		Else 
			println("welcome to the world of Dejava"); //end of condition
	}
	println();
	println("dejava loves %s.", dejava++? "Human" Else "Inline Statement" );
	println("dejava loves %s.", dejava++? "Human" Else "Inline Statement" );
	println();
	
	/* Humanic condition statement using 'If' boolean conversion*/
	For( 0, 6, +1 ){
		If (idx = 0) ? 
			println("'If' enables to use just one equal sign")
			
		Else If(idx = 1)?
			println("Do not forget to use '?' still, this does only a bool conversion")
			
		Else If(idx <= 2)? 
			println("Inequality operations are still compatible")
			
		Else idx == 3 ?
			println("Remember! without 'If', you can't use single equal sign")
			
		Else If(idx == 4)?
			println("In the If call, double equal also available. It may reduce your mistake")
			
		Else 
			println("Again, welcome to the world of Dejava"); //can we remove this semi colon?
	}
}


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
				case '!': //!=
					return idx != result;
				case '>':
					return (op[1]=='=')? 
						idx >= result 
						: idx > result;
				case '<':
					return (op[1]=='=')? 
						idx <= result 
						: idx < result;
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