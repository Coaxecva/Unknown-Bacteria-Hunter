import sys, os

fname = sys.argv[1]

with open(fname) as f:
    content = f.readlines()

for i in xrange(len(rcontent)):
	if line[i].startswith("#") and not line[i+1].startswith("#"):
		print(line[i+1])

