import sys
from sklearn.metrics.cluster import adjusted_mutual_info_score, adjusted_rand_score, homogeneity_score, completeness_score
from sklearn.metrics import jaccard_similarity_score

import math
from collections import Counter
import pandas as pd
import numpy as np

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

	predict = sys.argv[2]
	groundtruth = sys.argv[1]

	print("predict: ", predict)
	print("groundtruth: ", groundtruth)

	with open(predict) as f:
		content = [line.rstrip() for line in f]

	with open(groundtruth) as f:
		content1 = [line.rstrip() for line in f]

	if len(content1)==len(content):		
		print("Mutual info score: ", adjusted_mutual_info_score(content1, content))
		print("Rand Index: ", adjusted_rand_score(content1, content))
		print("Homogeneity Score: ", homogeneity_score(content1, content))
		print("Completeness Score: ", completeness_score(content1, content))
		print("Jaccard Similarity Score: ", jaccard_similarity_score(content1, content))


		unique_val = set(content)
		print(len(unique_val))
		#print(unique_val)
		
		df = pd.DataFrame({ 'groundtruth': content1, 
							'predict': content })

		en_arr = []
		for label in unique_val:
			#print(label)
			#print(df[df['predict']==label])
			group = list(df[df['predict']==label]['groundtruth'])
			en, ok = eta(group, "shannon")
			if ok != -1:
				print(en)
				en_arr.append(en)
				for key, value in ok.items():
					print("\t", key, value)
				#print(en, "\t", ok)
			print(np.average(en_arr))
			print(np.median(en_arr))

	else:
		print("Lengths of two groups are not equal!!!")
