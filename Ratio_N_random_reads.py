import random, sys, string
import pandas as pd
import math
from collections import Counter
import numpy as np

def ShuffleArray(fname):
	with open(fname) as f:
		content = [read.strip() for read in f]

	#print(content)
	#print(len(content))
	random.shuffle(content)
	return content

def ShuffleArray1(fname, c):
	# fname of corrected clusters
	num_lines = sum(1 for line in open(fname))
	#print(num_lines)
	#print(num_lines/c)
	arr = []

	for i in range(c):
		for j in range(num_lines//c):
			arr.append(i)	

	for j in range(num_lines - len(arr)):
		arr.append(i)

	# shuffle labels
	random.shuffle(arr)
	return arr

def eta(data, unit='natural'):
	base = { 'shannon' : 2.,
			'natural' : math.exp(1),
			'hartley' : 10. }
	if len(data) < 1:
		return 0, -1

	counts = Counter()
	for d in data:
		counts[d] += 1

	probs = [float(c) / len(data) for c in counts.values()]
	probs = [p for p in probs if p > 0.]
	ent = 0

	for p in probs:
		if p > 0.:
			ent -= p * math.log(p, base[unit])
	
	return ent, counts


if __name__ == '__main__':

	#read_f = sys.argv[1]
	#arr = ShuffleArray(read_f)
	#print(arr)
	#print(len(arr))

	fname = sys.argv[1]  # correct clusters
	num_unknown = sys.argv[2]

	# shuffle array
	arr = ShuffleArray1(fname, int(num_unknown))
	#print(arr)
	
	with open(fname) as f:
		content = [line.rstrip() for line in f]

	# Compute entropy
	unique_val = set(arr)
	#print(len(unique_val))

	df = pd.DataFrame({ 'groundtruth': content, 
							'predict': arr })

	en_arr = []
	for label in unique_val:
		#print(label)
		#print(df[df['predict']==label])
		group = list(df[df['predict']==label]['groundtruth'])
		en, ok = eta(group, "shannon")
		if ok != -1:
			print(en)
			en_arr.append(en)
			#for key, value in ok.items():
			#	print("\t", key, value)

	print(np.average(en_arr))
	print(np.median(en_arr))
