package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// WordCountPair struct
type WordCountPair struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

// ResponseBodyStruct struct
type ResponseBodyStruct struct {
	WordCountPairs []WordCountPair `json:"wordCountPairs"`
	WordCount      int             `json:"wordCount"`
}

// Upload for files
func Upload(w http.ResponseWriter, r *http.Request) {
	var str string
	counts := map[string]int{}

	var countPairs []WordCountPair

	if r.Method == http.MethodPost {
		f, _, err := r.FormFile("usrfile")
		if err != nil {
			log.Println(err)
			http.Error(w, "Error uploading file", http.StatusInternalServerError)
			return
		}
		defer f.Close()

		bs, err := ioutil.ReadAll(f)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}
		str = string(bs)
	}

	for _, word := range strings.Fields(str) {
		counts[word]++
	}

	for word, count := range counts {
		var toAppend WordCountPair
		toAppend.Word = word
		toAppend.Count = count
		countPairs = append(countPairs, toAppend)
	}

	// countPairsJSON, _ := json.MarshalIndent(countPairs, "", " ")
	wordCount := len(strings.Fields(str))

	var responseBody ResponseBodyStruct
	responseBody.WordCount = wordCount
	responseBody.WordCountPairs = countPairs

	w.Header().Set("CONTENT-TYPE", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	responseJSON, _ := json.MarshalIndent(responseBody, "", "")

	// writeToDisk({})

	w.Write(responseJSON)
}

// WordCount counts the frequency of each word
func WordCount(str string) map[string]int {
	counts := map[string]int{}
	for _, word := range strings.Fields(str) {
		counts[word]++
	}
	return counts
}

// TotalWordCount counts total number of words
func TotalWordCount(s string) int {
	words := strings.Fields(s)
	return len(words)
}

func main() {
	http.HandleFunc("/", Upload)
	log.Println("Listening on Port 8000")
	http.ListenAndServe(":8000", nil)
}
