import re
import functools

OPLIST = ["+","/","-","*"]
RE = "([+-/*])"

OPCOUNTER = lambda x,y: x + y

def build_tree(x):
    element_list = re.split(RE,x)
    print((len(element_list)-1)/2)
