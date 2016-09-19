from Bio.Blast import NCBIWWW
from Bio import SeqIO
import os, sys
#from Bio import SearchIO
from Bio.Blast import NCBIXML

BLAST_Bacteria = 'http://blast.ncbi.nlm.nih.gov/Blast.cgi?PAGE_TYPE=BlastSearch&PROG_DEF=blastn&BLAST_PROG_DEF=megaBlast&SHOW_DEFAULTS=on&BLAST_SPEC=MicrobialGenomes'

if __name__ == "__main__":

	in_file = sys.argv[1]
	xml_file = sys.argv[2]
	
	#
	print("Hello, BACTERIA HUNTER: LOADING....")

	#
	records = list(SeqIO.parse(in_file, "fasta"))
	print("# of seqs: ", len(records))

	for i in range(len(records)):
		#on file
	 	print("%s %i" % (records[i].id, len(records[i])))
	 	

	fasta_string = "\n".join([rec.format("fasta") for rec in records])
	result_handle = NCBIWWW.qblast("blastn", "nt", fasta_string, url_base=BLAST_Bacteria, hitlist_size=10)
	#result_handle = NCBIWWW.qblast("blastn", "nr", records.format("fasta"), hitlist_size=10)

	# store results in a file
	save_file = open(xml_file, "w")
	save_file.write(result_handle.read())
	save_file.close()
	result_handle.close()

	# get BLAST results from file
	result_handle = open(xml_file)
	blast_records = NCBIXML.parse(result_handle)
	
	# inspect results
	for record in blast_records:
		#print("new record...")
		# 
		if not record.alignments:
			print("N/A \n")
			continue
		for alignment in record.alignments:
			print(alignment.title)
		#
		best_alignment = record.alignments[0]
		best_hit = best_alignment.hsps[0]
		if best_hit.expect < 0.04:
			print("%s\n%.3e, %i" % (best_alignment.title, best_hit.expect, best_alignment.length))

	#
	#blast_records = NCBIXML.parse(result_handle)
	#blast_record = next(blast_records)

	sys.exit(0)

	#result_handle = NCBIWWW.qblast("blastn", "nt", record.format("fasta"))

	#Blast results
	#save_file = open("my_blast.xml", "w")
	#save_file.write(result_handle.read())
	#save_file.close()
	#result_handle.close()

	#Search result
	#blast_qresult = SearchIO.read('my_blast.xml', 'blast-xml')
	#for hit in blast_qresult[:5]:   # id and sequence length of the first five hits
		#print(hit.id, hit.description)
	#print(blast_qresult)