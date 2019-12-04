import re
import json
from sys import argv

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
signiture = f'(?P<param>{tupl})(?P<return>{datatype})?'

var_pattern = f'(?P<name>{word}):(?P<type>{datatype}|{signiture})'
def_pattern = f'(?P<name>{word})?(?P<type>{signiture})=(?P<value>.*)'

var_parser = re.compile(var_pattern)
def_parser = re.compile(def_pattern)

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
        'name': data.group('name'),
        'type': data['type'],
        'input': data.group('param'),
        'output': data.group('return'),
    })

it = def_parser.finditer(code)

for i, data in enumerate(it):
    kv = {}
    kv['name'] = data.group('name')
    kv['type'] = data.group('type')
    kv['input'] = data['param']
    kv['output'] = data['return']
    kv['value'] = data['value']
    result.append(kv)

with open('result.json', 'w', encoding='utf-8') as output:
    json.dump(result, output, indent="\t")

print("Conversion Complete. Please check 'result.json'")