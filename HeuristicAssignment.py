import sys, os
from Bio import SeqIO

fname = sys.argv[1]
fseq = sys.argv[2]

with open(fname) as f:
    content = f.readlines()

#print(len(content))

count = 0

# Extract significant names
for i in range(len(content)):
	if content[i].startswith("#") and not content[i+1].startswith("#"):
		print(content[i+1].rstrip())
		count += 1

print("# of seqs: ", count)



handle = open(fseq, "rU")
print(fseq)
for record in SeqIO.parse(handle, "fasta"):
    print(record.id)
handle.close()



