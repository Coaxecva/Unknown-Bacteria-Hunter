# My runner for Bacteria hunter

import sys, os

samtools = ""
ref_path = ""
read_path = ""
bowtie_path = ""
bowtie_index_path = ""


# concatenate refs
os.system(ref_path+"cat * > mutltigenomes.fa")

#
os.system(bowtie_path + "bowtie2-build " + ref_path + "multigenomes.fa " + bowtie_index_path + "mutltigenome-bowtie2")
