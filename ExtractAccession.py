import sys, os

fn1 = sys.argv[1]
f1 = open(fn1)

fs = True
idx = 0
t = set()

while True:
	info= f1.readline()
	if info=='' or info==" ":
		break
	if info[0]=="#" and fs:
		idx += 1
		print(idx)
		fs = False		
	if info[0]!="#" and fs==False:
		fs = True
		for id in t:
			print(id)
		t = set()
	if info[0]!="#":
		#print(info.split())
		if not info.split():
			break
		t.add(info.split()[2])
		#print(t)

f1.close()