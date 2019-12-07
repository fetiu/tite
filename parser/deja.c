#include <stdio.h>
#include <string.h>
#include <stdbool.h>

bool parse(char* line);

int main(int argc, char *argv[])
{
    if (argc != 2)
    {
        printf("usage: %s <file>\n", argv[0]);
        return -1;
    }

    const char *src = argv[1];
    FILE* file = fopen(src, "r");

    if (NULL == file)
    {
        printf("No such file: %s\n", src);
        return -1;
    }

    char line[256];
    bool result = true;
    
    while (true != feof(file) && false != result)
    {
        fgets(line, sizeof(line), file);
        result = parse(line);
        memset(line, 0, sizeof(line));
    }
    
    return 0;
}

typedef enum
{
    VARIABLE,
    FUNCTION,
    LITERAL,
    TYPE,
    COMMAND,
    UNKNOWN,
} label_t;

typedef struct
{
    char* name;
    char* type;
    char* value;
} token_t;

label_t match(char a)
{
    switch(a)
    {
        case ':':
            return VARIABLE;
        case '(':
            return FUNCTION;
        case '=':
            return LITERAL;
        case '[':
            return TYPE;
        default:
            return UNKNOWN;
    }
}

char* lookup[]=
{
    [VARIABLE] = "variable",
    [FUNCTION] = "function",
    [LITERAL]  = "literal",
    [TYPE]  = "type",
    [COMMAND] = "command",
    [UNKNOWN] = "unknown",
};

bool parse(char* line)
{
    const int max = strlen(line);
    label_t label;

    for(int pos= 0; pos != max; pos++)
    {
        // printf("%s\n", lookup[match(line[pos])]);

        char a = line[pos];
        label = match(a);
        if(label != UNKNOWN)
        {
            char* tok = strtok(line, &a);
            printf("%s:%s\n", lookup[label], tok);
        }

    }
    return true;
}

//label은 상태값으로써 지니는게 나을지도? 값들마다 델리미터가 다르니...