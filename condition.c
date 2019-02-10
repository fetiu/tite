/*
	Prototype of Condition Statement rule using Ternary operator
	Property of the Project Dejava. Copyright 2019, Jake Fetiu.	
*/
#include"dejava.h"

int main(){
	bool_t dejava = False;
	
	/* basic condition statement. notice that there's no 'if' */
	For( 0, 6, +1 )
	{ 
		idx == 0 || idx == 4?  //idx is temporary reserved word.
			println("hello world")
		Else idx <= 3 ? 
			println("hello friend %d", idx)
		Else 
			println("welcome to the world of Dejava"); //end of condition
	}//For iterates 0->5. this is temprorary expression.
	
	
	For(2,0,-1)
		printf("\ndejava loves %s.", dejava++? "Human" Else "Inline Statement" );
	printf("\n"); // Instead, you can use below.
	println();
	
	/* Humanic condition statement using 'If' boolean conversion*/
	For( 0, 6, +1 )
	{
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