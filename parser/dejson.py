import re
import json
from sys import argv

delims = r'[,\s\]\)]'

#literals
tupl = r'\(.*\)'
impl = r'\{.*\}'
string = '".*"'
number = r'\d+'

#symbols
word = r'\w+'
enum = r'\<.*\>'
struct = r'\[.*\]'
# for post-parsing /lazy parsing
datatype = f'(?:{word}|{enum}|{struct})'
signiture = f'(?:{tupl}){datatype}?'
value = '(?P<value>.+)'

box_pattern = f'(?P<name>{word}):(?P<type>{signiture}|{datatype}){delims}'
var_pattern = f'(?P<name>{word}):(?P<type>(?:{tupl})?{datatype}?)(?:{delims}|={value})'
def_pattern = f'(?P<name>{word})?(?P<type>{signiture})={value}'
# is function name necessary?

pattern = var_pattern 
print(pattern)

parser = re.compile(pattern)

if(len(argv) == 2):
    src = argv[1]
else:
    raise SystemExit(f"usage: python3 {argv[0]} <path/file.dej>")

source = open(src, mode='r')
code = source.readlines()
code = ''.join(code)

it = parser.finditer(code)

result = []

for data in it:
    result.append(data.groupdict())

with open('result.json', 'w', encoding='utf-8') as output:
    json.dump(result, output, indent="\t")

print("Conversion Complete. Please check 'result.json'")