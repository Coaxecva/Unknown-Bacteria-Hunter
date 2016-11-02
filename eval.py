import sys
from sklearn.metrics.cluster import adjusted_mutual_info_score

if __name__ == '__main__':

	predict = sys.argv[1]
	groundtruth = sys.argv[2]

	with open(predict) as f:
		content = [line.rstrip() for line in f]

	with open(groundtruth) as f:
		content1 = [line.rstrip() for line in f]

	if len(content1)==len(content):
		print(adjusted_mutual_info_score(content1, content))
	else:
		print("Lengths of two groups are not equal!!!")
