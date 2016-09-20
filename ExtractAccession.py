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
		t = set()
		while f1.readline()[0] == "#": continue
		while True:
			info = f1.readline()
			if not info.split() or info[0]=="#":
				break
			#print(info.split())
			t.add(info.split()[2])
		print(idx)
		for accs in t:
			print(accs)
f1.close()
