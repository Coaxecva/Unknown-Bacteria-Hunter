import sys
from Bio import Entrez

#define email for entrez login
db           = "nuccore"
Entrez.email = "some_email@somedomain.com"
batchSize    = 100
retmax       = 10**9

def is_number(s):
    try:
        float(s)
        return True
    except ValueError:
        return False

def retrieve(accs):
	#first get GI for query accesions
	#sys.stderr.write( "\nFetching %s entries from GenBank: %s\n" % (len(accs), ", ".join(accs[:10])))
	query  = " ".join(accs)
	handle = Entrez.esearch( db=db,term=query,retmax=retmax )
	giList = Entrez.read(handle)['IdList']
	#sys.stderr.write( "Found %s GI: %s\n" % (len(giList), ", ".join(giList[:10])))

	#post NCBI query
	search_handle     = Entrez.epost(db=db, id=",".join(giList))
	search_results    = Entrez.read(search_handle)
	webenv,query_key  = search_results["WebEnv"], search_results["QueryKey"] 


	#fecth all results in batch of batchSize entries at once
	for start in range( 0,len(giList),batchSize ):
		sys.stderr.write( " %9i" % (start+1,))
		#fetch entries in batch
		handle = Entrez.efetch(db=db, rettype="gb", retstart=start, retmax=batchSize, webenv=webenv, query_key=query_key, retmode="xml")
		#print output to stdout
		#sys.stdout.write(handle.read())
		records = Entrez.parse(handle)
		for record in records:
			print(record['GBSeq_source'])
		handle.close()

#main
if __name__ == '__main__':

	fn1 = sys.argv[1]	
	Table = {}
	key = -1

	for line in open(fn1):		
		if is_number(line):		
			key = int(line)			
			Table[key]=[]
		else:
			Table[key].append(line.rstrip("\n"))

	#print(Table)
	for k in Table:
		print(k)
		retrieve(Table[k])
