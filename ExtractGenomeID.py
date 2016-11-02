import sys
from Bio.Seq import Seq

GenomeID = {}

# Extract GenomeID from fastq
def ExtractFASTQ(fname):
	with open(fname) as f:
		content = f.readlines()

	#print(len(content), len(content[0::4]), len(content[1::4]))

	for (line1, line2) in zip(content[0::4], content[1::4]):
		#print(line1.rstrip().split()[1])
		#print(line2.rstrip())
		GenomeID[line2.rstrip()] = line1.rstrip().split()[1] 


# Retrieve GenomeID from unknow genome reads
def RetriveGenome(fname):
	with open(fname) as f:
		content = [line.rstrip() for line in f]
		
	# Open a file
	fo1 = open("correct_reads.txt", "wb")
	fo2 = open("correct_clusters.txt", "wb")

	
	for i, read in enumerate(content):
		if read in GenomeID:
			fo1.write(read);
			fo2.write(GenomeID[read]);
		#	print(GenomeID[read])
		# else:
		# 	seq = Seq(read)
		# 	print(i, read, str(seq.reverse_complement()), str(seq.reverse_complement()) in GenomeID)

	# Close opend file
	fo1.close()
	fo2.close()

if __name__ == '__main__':

	fastq = sys.argv[1]
	txt = sys.argv[2]

	ExtractFASTQ(fastq)
	RetriveGenome(txt)
