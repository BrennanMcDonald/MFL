import functools

OPLIST = ["+","/","-","*"]

def evaluate(l,m,r):
    if(m == "-"):
        return int(l) - int(r)
    elif (m == "+"):
        return int(l) + int(r)
    elif (m == "*"):
        return int(l) * int(r)
    elif (m == "/"):
        return int(l) // int(r)

def split(x):
    for c in OPLIST:
        if c in x:
            return c,x.split(c)

def parse2(x):
    x = x.replace(" ","")
    if not functools.reduce(lambda x,y: x + y,[i in x for i in OPLIST]) == 1:
        return False,0
    else:
        delim,lr = split(x)
        if not len(lr) == 2:
            return False,0
        else:
            return True,evaluate(lr[0],delim,lr[1])
