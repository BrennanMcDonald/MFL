from optparse import OptionParser
import itertools

parser = OptionParser()
parser.add_option("-f", "--file", dest="filename")

(options, args) = parser.parse_args()

file = options.filename

to_parse = open(file)

with open("BNF") as f:
    content = f.readlines()

content = [x.strip() for x in content]

rules = {x.split(":")[0]:list(itertools.starmap((lambda r: r.trim()), x.split(":")[1].split("|"))) for x in content}
print(rules)
