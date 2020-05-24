package main

import (
       "bufio"
       "fmt"
       "os"
)

/*
type Container interface {
     
}
*/

func main() {
     if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
	fmt.Println("Usage: txtcmp [OPTION]... DOC_1_FILEPATH DOC_2_FILEPATH")  
     	return
     }

     if len(os.Args) < 3 {
     	fmt.Println("txtcmp: missing file operands")
	fmt.Println("Usage: txtcmp [OPTION]... DOC_1_FILEPATH DOC_2_FILEPATH")  
     	return
     }


     file_1, error := os.Open(os.Args[1])
     if error != nil {
     	panic(error)
     }
     
     file_2, error := os.Open(os.Args[2])
     if error != nil {
     	panic(error)
     }

     var file1_frequencies map[string]int = make(map[string]int)
     var file2_frequencies map[string]int = make(map[string]int)

     file1_num_words := 0
     file2_num_words := 0

     scanner := bufio.NewScanner(file_1)
     scanner.Split(bufio.ScanWords)
     for scanner.Scan() {
     	 file1_frequencies[scanner.Text()]++
	 file1_num_words++
     }

     scanner = bufio.NewScanner(file_2)
     scanner.Split(bufio.ScanWords)
     for scanner.Scan() {
     	 file2_frequencies[scanner.Text()]++
	 file2_num_words++
     }

     var shorter, longer *map[string]int
     var shorter_len int 
     if file1_num_words < file2_num_words {
     	shorter = &file1_frequencies
	shorter_len = file1_num_words
	longer = &file2_frequencies
     } else {
     	shorter = &file2_frequencies
	shorter_len = file2_num_words
	longer = &file1_frequencies  
     }

     differences := 0
     for word, frequency := range *shorter {
     	 differences += Max(frequency - (*longer)[word], 0) 
     }

     matches := shorter_len - differences
     comparison_result := float64(matches) / float64(shorter_len) * 100.0
     fmt.Printf("Documents are %.2f%% similar\n", comparison_result)
}

func Max(x, y int) int {
     if x > y {
     	return x
     } else {
       return y
     }
}