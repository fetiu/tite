import re
import json
from sys import argv

def capture(k,v):
    return f'(P?<{k}>{v})'

delims = r'[\s,]'

#symbols
word = r'\w+'
enum = r'\<.+\>'
struct = r'\[.+\]'

#literals
tupl = r'\(.*\)'
impl = r'\{.*\}'
string = '".*"'
number = r'\d+'

datatype = f'{word}|{enum}|{struct}'
signiture = f'(?:{tupl})?(?:{datatype})?'

var_pattern = f'(?P<name>{word}):(?P<type>(?:{tupl})?(?:{datatype}))=?(?P<value>.*)'
def_pattern = f'(?P<name>{word})?(?P<param>{tupl})(?P<return>{datatype}|)?=(?P<value>.*)'
#유지보수가 거의 불가능한 코드네
pattern = capture('name',word) + '(:?)' + capture('type',signiture)+ '(?:{delims}|=(.+))'

var_parser = re.compile(var_pattern)
def_parser = re.compile(def_pattern)
parser = re.compile(pattern)

if(len(argv) == 2):
    src = argv[1]
else:
    raise SystemExit(f"usage: python3 {argv[0]} <path/file.dej>")

source = open(src, mode='r')
code = source.readlines()
code = ''.join(code)

# variables = var_parser.findall(code)
# functions = def_parser.findall(code)

result = []

it = var_parser.finditer(code)

for i, data in enumerate(it):
    result.append({
        'name': data['name'],
        'type': data['type'],
        'type': data['value']
    })

it = def_parser.finditer(code)

for i, data in enumerate(it):
    result.append({
        'name': data['name'],
        'type': data['param'],
        'type': data['return'],
        'type': data['value']
    })

with open('result.json', 'w', encoding='utf-8') as output:
    json.dump(result, output, indent="\t")

print("Conversion Complete. Please check 'result.json'")