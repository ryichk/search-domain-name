package main

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/ryichk/search-domain-name/thesaurus"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf(".envの読み込みに失敗しました: %v\n", err)
	}
	apiKey := os.Getenv("BHT_APIKEY")
	thesaurus := &thesaurus.BigHuge{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalf("%qの類語検索に失敗しました: %v\n", word, err)
		}
		if len(syns) == 0 {
			log.Fatalf("%qに類語はありませんでした\n")
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
