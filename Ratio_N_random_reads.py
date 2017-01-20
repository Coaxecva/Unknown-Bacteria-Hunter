import random, sys, string

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
	print(num_lines/c)

if __name__ == '__main__':

	#read_f = sys.argv[1]
	#arr = ShuffleArray(read_f)
	#print(arr)
	#print(len(arr))

	fname = sys.argv[1]
	num_unknown = sys.argv[2]
	ShuffleArray1(fname, int(num_unknown))


	
