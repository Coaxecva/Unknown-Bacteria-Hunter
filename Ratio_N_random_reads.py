import random, sys, string

def ShuffleArray(fname):
	with open(fname) as f:
		content = [read.strip() for read in f]

	#print(content)
	#print(len(content))
	random.shuffle(content)
	return content

if __name__ == '__main__':

	read_f = sys.argv[1]
	arr = ShuffleArray(read_f)
	#print(arr)
	#print(len(arr))

	
