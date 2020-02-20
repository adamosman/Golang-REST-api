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
	var wordCount int
	var countPairs []WordCountPair
	counts := map[string]int{}

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
		if strings.Contains(word, "blue") {
			continue
		}
		wordCount++
		word = strings.Trim(word, ".,\"!-/:-@[-`{-~")
		counts[word]++
	}

	for word, count := range counts {
		var toAppend WordCountPair
		toAppend.Word = word
		toAppend.Count = count
		countPairs = append(countPairs, toAppend)
	}

	var responseBody ResponseBodyStruct
	responseBody.WordCount = wordCount
	responseBody.WordCountPairs = countPairs

	w.Header().Set("CONTENT-TYPE", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	responseJSON, _ := json.Marshal(responseBody)

	w.Write(responseJSON)
	// fmt.Println(string(responseJSON)) // checking output of JSON
}

func main() {
	http.HandleFunc("/", Upload)
	log.Println("Listening on Port 8000")
	http.ListenAndServe(":8000", nil)

}

// blah blah blah