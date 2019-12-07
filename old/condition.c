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
	}
	//For iterates param1 gets param2 increases byparm3. this is temprorary expression.
	//Tentative Expr: "for idx++:5{". assign 5 to increasing variable(iterator) idx.
	
	For(2,0,-1)
		printf("\ndejava loves %s.", dejava++? "Human" Else "Inline Statement" );
	printf("\n"); // Instead, you can use below.
	println();
	
	/* Humanic condition statement using 'If' boolean conversion*/
	For( 0, 6, +1 )
	{
		If(idx = 0)
			println("'If' enables to use just one equal sign.")
			
		Else If(idx = 1)
			println("Now you don't need to use '?'.\n This does both bool conversion & condition branch.")

		Else If(idx == 2)
			println("In the If call, double equal also available. \nIt may reduce your mistake.")
							
		Else If(idx <= 3)
			println("Inequality signs are still compatible. \nAlso supports other high level condition.")
		
		Else idx == 4 ? //possible to combine ternary and if statement.
			println("Remember! without 'If' you should use '?', and can't use single equal sign.")
			
		Else
			println("Again, welcome to Dejava."); //can we remove this semi colon?
	}
	//Tentative Expr: Add support 'and', 'or', 'xor', 'equals', 'has' keywords lifeIf(idx is 1)
}

//TODO: enable to use multiple lines in ternary operations by instant inline function.
// scope formated by {}
