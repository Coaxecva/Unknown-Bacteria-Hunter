import sys, os

fn1 = sys.argv[1]
f1 = open(fn1)

fs = True
idx = 0
#t = set()

while True:
	info= f1.readline()
	if info=='' or info==" ":
		break
	if info[0]=="#":
		idx += 1		
		while f1.readline()[0] == "#": continue
		while True:
			info = f1.readline()
			if not info.split() or info[0]=="#":
				break
			print(info.split())
			while f1.readline()[0]!="#" or f1.readline()[0]!='' or f1.readline()[0]!=" ": continue
			break
		
print(idx)

f1.close()
