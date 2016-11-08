package main

import (
	"bufio"
	"fmt"
	"strings"
	"log"
	"os"
	"flag"
)

// DNA read
type DNARead struct {
	data []byte
}

func InitDNARead(data []byte) *DNARead{
	read := new(DNARead)
	read.data = data
	return read
}

// Min
func Min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

// Max
func Max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

// read unmapped reads
func ReadSequence(file string) []DNARead{
	reads := make([]DNARead,0)
	// Open the file.
    f, _ := os.Open(file)
    // Create a new Scanner for the file.
    scanner := bufio.NewScanner(f)
    // Loop over all lines in the file and print them.
    for scanner.Scan() {
		line := scanner.Text()
		line = strings.Split(line,"\t")[9]
		// read info
		// fmt.Println(strings.Split(line,"\t")[9])
		if !StringInSlice(line, reads) {
			read := InitDNARead([]byte(line))
			reads = append(reads, *read)
		}
    }
    return reads
}

// reverse complement of a read
func ReverseComplement(read []byte) []byte{
	read_len := len(read)
	rev_comp_read := make([]byte, read_len)
	for i, elem := range read {		
		if elem == 'A' {
			rev_comp_read[read_len-1-i] = 'T'
		} else if elem == 'T' {
			rev_comp_read[read_len-1-i] = 'A'
		} else if elem == 'C' {
			rev_comp_read[read_len-1-i] = 'G'
		} else if elem == 'G' {
			rev_comp_read[read_len-1-i] = 'C'
		} else {
			rev_comp_read[read_len-1-i] = elem
		}
	}
	return rev_comp_read
}

// Swap two strings
func SwapString(str1, str2 *string) {
    *str1, *str2 = *str2, *str1
}

// Is a read in bool?
func StringInSlice(read string, reads []DNARead) bool {
 	for _, v := range reads {
 		if strings.Compare(string(v.data), string(read))==0 {
 			return true
 		}
 	}
 	return false
 }

 // Print out
 func PrintDNAreads(li []DNARead) {
 	fmt.Println(len(li))
 	for i := range(li) {
			fmt.Println(string(li[i].data))
		}
 }

// Overlapp
 func OverlappString(a, b string) int {
 	maxOverlap := Min(len(a)-1, len(b)-1)

 	if len(b)>len(a) {
 		SwapString(&a,&b)
 	}

 	// Start with maximum possible overlap and work down until a match is found
 	for (!strings.HasSuffix(a, b[0:maxOverlap]) && !strings.HasSuffix(b, a[0:maxOverlap])) {
 		//fmt.Println(a, b, b[0:maxOverlap], a[0:maxOverlap], maxOverlap)
 		maxOverlap -= 1
 	}
 	return maxOverlap
 }

// Edit distance
 func ComputeEditDistance(x, y string, i, j int) int {
	if j==-1 {
		return i+1
	}
	if i==-1 {
		return j+1
	}
	if x[i]==y[j]{
		return ComputeEditDistance(x,y,i-1,j-1)
	} else {
		return Min3(1+ComputeEditDistance(x,y,i-1,j-1), 1+ComputeEditDistance(x,y,i-1,j), 1+ComputeEditDistance(x,y,i,j-1))
	}
 }
 
// The Levenshtein distance between two strings is defined as the minimum
// number of edits needed to transform one string into the other, with the
// allowable edit operations being insertion, deletion, or substitution of
// a single character
func ComputeLevenshteinDistance(str1, str2 string) int {
	var cost, lastdiag, olddiag int
	s1 := []rune(str1)
	s2 := []rune(str2)

	len_s1 := len(s1)
	len_s2 := len(s2)

	column := make([]int, len_s1+1)

	for y := 1; y <= len_s1; y++ {
		column[y] = y
	}

	for x := 1; x <= len_s2; x++ {
		column[0] = x
		lastdiag = x - 1
		for y := 1; y <= len_s1; y++ {
			olddiag = column[y]
			cost = 0
			if s1[y-1] != s2[x-1] {
				cost = 1
			}
			column[y] = Min3(
				column[y]+1,
				column[y-1]+1,
				lastdiag+cost)
			lastdiag = olddiag
		}
	}
	return column[len_s1]
}

func Min3(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}


// main
func main() {
	log.Printf("START!!! \n")
	log.Printf("go run OverlapComponent.go -q <unmapped reads>")

	var queries_file = flag.String("q", "", "queries file")
	flag.Parse()

	//s1 := "AAAACBC"
	//s2 := "BCEAAA"
	//fmt.Println(string(ReverseComplement([]byte("ATGGCCTTAAA"))))
	//fmt.Println(OverlappString(s1,s2))
	//fmt.Println(LevenshteinDistance(s1, s2))
	//fmt.Println((ComputeEditDistance(s1,s2,len(s1)-1,len(s2)-1)))


	//fmt.Println(*queries_file)
	if *queries_file != "" {
		DNAreads := ReadSequence(*queries_file)

		for i:=0; i<len(DNAreads); i++ {
			for j:=i+1; j<len(DNAreads); j++  {
				if OverlappString(string(DNAreads[i].data),string(DNAreads[j].data))>7 ||
				OverlappString(string(DNAreads[i].data),string(ReverseComplement(DNAreads[j].data)))>7 || 
				OverlappString(string(ReverseComplement(DNAreads[i].data)),string(DNAreads[j].data))>7 {
					fmt.Println(i, j)
				}
			}
		}

		// PrintDNAreads(DNAreads)
		
	}	
}