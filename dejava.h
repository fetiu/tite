
#include<stdio.h> 
#include<string.h>
#include<stdlib.h>
#define println(fmt,...) printf(fmt"\n", ##__VA_ARGS__)

/* Equality keyworlds */
/*
#define equal	=
#define not !
#define equals equal equal
*/
#define equals ==
#define not_equal !=
#define bigger >
#define lesser <
#define greater bigger
#define smaller lesser


#define If(x) boolConversion(#x"?")?
#define Else :
#define For(i,d,x) idx = (int)i-x; while( d != (idx+=x) )
#define DELIM "=!<>? "

int idx;
typedef enum { False, True } bool_t;
bool_t boolConversion(char* statement);
bool_t parseOperand(const void* context, char op[] , char*tok);

