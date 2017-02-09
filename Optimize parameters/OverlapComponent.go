package main

import (
	"bufio"
	"fmt"
	"strings"
	"log"
	"os"
	"flag"
	"crypto/sha256"
	//"./mapset"
	"./unionfind"
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

// read REFINED unmapped reads
func ReadSequence1(file string) []DNARead{
	reads := make([]DNARead,0)
	// Open the file.
    f, _ := os.Open(file)
    // Create a new Scanner for the file.
    scanner := bufio.NewScanner(f)
    // Loop over all lines in the file and print them.
    for scanner.Scan() {
		line := scanner.Text()
		// line = strings.Split(line,"\t")[9]
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

 // Write DNAreads into a file
func WriteDNAreads(lines []DNARead, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, string(line.data))
	}
	return w.Flush()
}

// ExactOverlapp
 func ExactOverlapString(a, b string) int {
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

// Count the number of bits that are different
func BitsDifference(h1, h2 *[sha256.Size]byte) int {
    n := 0
    for i := range h1 {
        for b := h1[i] ^ h2[i]; b != 0; b &= b - 1 {
            n++
        }
    }
    return n
}

// Hamming distance
func ComputeHammingDistance(str1, str2 string) int {
	c := 0
	for i := range str1 {
		if str1[i] != str2[i]{
			c++
		}
	}
	return c
}

func ApproximateHamming(overlap int, str1, str2 string) int {
	return overlap - ComputeHammingDistance(str1, str2)
}

func ApproximateLevenshtein(overlap int, str1, str2 string) int {
	return overlap - ComputeLevenshteinDistance(str1, str2)
}

// Approximate overlaps on Hamming
func ApproximateHammingOverlap(a, b string) int {

	maxOverlap := Min(len(a)-1, len(b)-1)
	appr := 0

	println(maxOverlap)

	//println(a, " ", len(a))
	//println(b, " ", len(b))

 	if len(b)>len(a) {
 		SwapString(&a,&b)
 	}

  	// Start with maximum possible overlap and work down until a max is found
 	for (maxOverlap>0) {
 		//fmt.Println(a[len(a)-maxOverlap:])
 		//fmt.Println(b[0:maxOverlap])
		// fmt.Println(b[len(b)-maxOverlap:])
		// fmt.Println(a[0:maxOverlap])
		// fmt.Println("-----")

 		right := ApproximateHamming(maxOverlap, a[len(a)-maxOverlap:], b[0:maxOverlap])
 		left := ApproximateHamming(maxOverlap, b[len(b)-maxOverlap:], a[0:maxOverlap]) 		

		//fmt.Println(right, left)

 		if  right > appr {
 			appr = right
 		}
 		
 		if  left > appr {
 			appr = left
 		}
 		
 		maxOverlap -= 1
 		//fmt.Println("maxoverlap ", maxOverlap)
 		if maxOverlap < appr {
 			break
 		}
 	}
 	pct := appr/maxOverlap;
 	fmt.Println(pct)
 	//fmt.Println(appr, " ", maxOverlap)
 	return appr	
}

// Aproximate overlaps on Levenshtein
func ApproximateLevenshteinOverlap(a, b string) int {
	maxOverlap := Min(len(a)-1, len(b)-1)
	appr := 0

 	if len(b)>len(a) {
 		SwapString(&a,&b)
 	}

 	// Start with maximum possible overlap and work down until a max is found
 	for (maxOverlap>0) {
		// fmt.Println(a[len(a)-maxOverlap:])
		// fmt.Println(b[0:maxOverlap])
		// fmt.Println(b[len(b)-maxOverlap:])
		// fmt.Println(a[0:maxOverlap])
		// fmt.Println("-----")

 		right := ApproximateLevenshtein(maxOverlap, a[len(a)-maxOverlap:], b[0:maxOverlap])
 		left := ApproximateLevenshtein(maxOverlap, b[len(b)-maxOverlap:], a[0:maxOverlap])

		// fmt.Println(right, left)

 		if  right > appr {
 			appr = right
 		}
 		
 		if  left > appr {
 			appr = left
 		}

 		maxOverlap -= 1
 		if maxOverlap < appr {
 			break
 		}
 	}
 	return appr	
}

// opt==1: exact overlap
// opt==2: hamming distance overlap
// opt==3: edit distance overlap
func Condition(str1, str2 string, opt int, threshold int) bool {
	first_rev := string(ReverseComplement([]byte(str1)))
	second_rev := string(ReverseComplement([]byte(str2)))
	switch opt {
		case 1: {
			return (ExactOverlapString(str1,str2)>threshold || 
			ExactOverlapString(str1, second_rev)>threshold || 
			ExactOverlapString(first_rev,str2)>threshold)
		}
		case 2: {
			return (ApproximateHammingOverlap(str1,str2)>threshold || 
			ApproximateHammingOverlap(str1, second_rev)>threshold || 
			ApproximateHammingOverlap(first_rev,str2)>threshold) 
		}
		case 3: {
			return (ApproximateLevenshteinOverlap(str1,str2)>threshold || 
			ApproximateLevenshteinOverlap(str1, second_rev)>threshold || 
			ApproximateLevenshteinOverlap(first_rev,str2)>threshold)
		}
	}
	return false
}


// main
func main() {
	log.Printf("START!!! \n")
	log.Printf("go run OverlapComponent.go -qf <unmapped reads> -sr <distance selector> -ot <threshold>")

	var queries_file = flag.String("qf", "", "queries file")
	var selector = flag.Int("sr", 1, "measuring distance selector")
	var overlap_threshold = flag.Int("ot", 1, "overlap two substring threshold")

	flag.Parse()
	//fmt.Println(*queries_file)

	//s1 := "AAAACBC"
	//s2 := "BCEAAA"
	//fmt.Println(string(ReverseComplement([]byte("ATGGCCTTAAA"))))
	//fmt.Println(OverlappString(s1,s2))
	//fmt.Println(LevenshteinDistance(s1, s2))
	//fmt.Println((ComputeEditDistance(s1,s2,len(s1)-1,len(s2)-1)))

	if *queries_file != "" {

		// DNAreads := ReadSequence(*queries_file)
		// DNA reads refined
		DNAreads := ReadSequence1(*queries_file)

		// if err := WriteDNAreads(DNAreads, "DNAreads.txt"); err != nil {
		// 	log.Fatalf("writeLines: %s", err)
		// }

		// PrintDNAreads(DNAreads)
		// fmt.Println(len(DNAreads))

		th := *overlap_threshold
		var uf *unionfind.UnionFind
		uf = unionfind.New(len(DNAreads))		
		
		for i:=0; i<len(DNAreads); i++ {
			for j:=i+1; j<len(DNAreads); j++  {				
				// if Condition(string(DNAreads[i].data),string(DNAreads[j].data), *selector, *overlap_threshold) {
				if Condition(string(DNAreads[i].data),string(DNAreads[j].data), *selector, th) {
					//fmt.Println(i, j)
					uf.Union(i,j)						
				}
			}
		}

		// fmt.Println(uf.GetNumClusters())
		//uf.PrintClusters()
		// PrintDNAreads(DNAreads)					
				
	}		
}
