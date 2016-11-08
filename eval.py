import sys
from sklearn.metrics.cluster import adjusted_mutual_info_score, adjusted_rand_score


import pandas as pd

def eta(data, unit='natural'):
	base = { 'shannon' : 2.,
			'natural' : math.exp(1),
			'hartley' : 10. }
	if len(data) <= 1:
		return 0

	counts = Counter()
	for d in data:
		counts[d] += 1

	probs = [float(c) / len(data) for c in counts.values()]
	probs = [p for p in probs if p > 0.]
	ent = 0

	for p in probs:
		if p > 0.:
			ent -= p * math.log(p, base[unit])
	
	return ent


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

		unique_val = set(content)
		
		df = pd.DataFrame({ 'groundtruth': content1, 
							'predict': content })

		for i in unique_val:
			print(df[i])

	else:
		print("Lengths of two groups are not equal!!!")
