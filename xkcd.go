package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// If search terms are provided, this will download all of the xkcd
// comics, if they're not already downloaded, and then perform a case
// insensitive search of the search therms through alt, safe_title,
// title, and transcript of each comic. It logs out each match as it's
// found and then prints the number of matches found, if any, along with
// the search terms that were used.
func main()  {
	searchTerms := os.Args[1:]
	if len(searchTerms) == 0 {
		fmt.Println("Error: xkcd search terms go here.")
		os.Exit(1)
	}

	joinedTerms := strings.Join(searchTerms, " ")

	fmt.Printf("Submitted search terms %s\n", joinedTerms)
	err := DownloadComics(); if err != nil {
		log.Fatal(err)
	}

	numOfMatches, numFilesSearched := FindSearchTerms(joinedTerms)

	if numOfMatches == 0 {
		fmt.Printf("0 matches for search terms: %s\nSearched %d xkcd comics.\n", joinedTerms, numFilesSearched)
		os.Exit(0)
	}

	fmt.Printf("Found %d matches for search terms: %s\nSearched %d xkcd comics\n", numOfMatches, joinedTerms, numFilesSearched)
	os.Exit(0)
}


/*

Get the input args, anything past the command is the search query
If no data
  Pull it from xkcd_fetch starting at https://xkcd.com/1/info.0.json and go until there's an error
    Save the json response locally (num, safe_title, transcript, alt, and title)
count matches
Search through each JSON file
  If search term matches
    Print transcript and link

If matches is 0
  Print no matches found for search term searchTerm
*/

func FindSearchTerms(searchTerms string) (matches int, filesSearched int) {
	lowerTerms := strings.ToLower(searchTerms)
	files, err := os.ReadDir(StorageDir)

	if err != nil {
		log.Fatalf("Failed to read directory in FindSearchTerms: %s\n", err)
	}

	for _, file := range files {
		data, rfErr := os.ReadFile(StorageDir + "/" + file.Name())

		if rfErr != nil {
			log.Fatalf("Failed to open file: %s\nWith error: %s\n", file.Name(), rfErr)
		}

		var comicInfo ComicContents

		unmarshalErr := json.Unmarshal(data, &comicInfo)

		if unmarshalErr != nil {
			log.Fatalf("Failed to unmarshal file: %s\nError: %s\n", file.Name(), unmarshalErr)
		}


		if strings.Contains(strings.ToLower(comicInfo.Alt), lowerTerms) {
			matches++
			fmt.Printf("Match %d\nTranscript: %s\nComic URL: %s\n\n", matches, comicInfo.Transcript, BaseUrl+ strconv.Itoa(comicInfo.Num))
			continue
		}

		if strings.Contains(strings.ToLower(comicInfo.SafeTitle), lowerTerms) {
			matches++
			fmt.Printf("Match %d\nTranscript: %s\nComic URL: %s\n\n", matches, comicInfo.SafeTitle, BaseUrl+ strconv.Itoa(comicInfo.Num))
			continue
		}

		if strings.Contains(strings.ToLower(comicInfo.Title), lowerTerms) {
			matches++
			fmt.Printf("Match %d\nTranscript: %s\nComic URL: %s\n\n", matches, comicInfo.Title, BaseUrl+ strconv.Itoa(comicInfo.Num))
			continue
		}

		if strings.Contains(strings.ToLower(comicInfo.Transcript), lowerTerms) {
			matches++
			fmt.Printf("Match %d\nTranscript: %s\nComic URL: %s\n\n", matches, comicInfo.Transcript, BaseUrl+ strconv.Itoa(comicInfo.Num))
			continue
		}
	}

	return matches, len(files)
}

const (
	BaseUrl = "https://xkcd.com/"
	JsonApiEnding = "info.0.json"
	StorageDir = "./storage"
)

type ComicContents struct {
	Alt string
	Num int
	SafeTitle string `json:"safe_title"`
	Title string
	Transcript string
}

// When fetching the current comic, comicNum is "",
// otherwise it's the number of the comic to fetch
func fetchComic(comicNum string) (ComicContents, error) {
	var comic string
	if comicNum == "" {
		comic = ""
	} else {
		comic = comicNum + "/"
	}

	resp, err := http.Get(BaseUrl + comic + JsonApiEnding)

	if err != nil {
		return ComicContents{}, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return ComicContents{}, fmt.Errorf("fetch of comic %s failed: %s", comicNum, resp.Status)
	}

	var contents ComicContents
	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		resp.Body.Close()
		return ComicContents{}, err
	}

	resp.Body.Close()

	return contents, nil
}

func getCurrentNumberOfComics() (int, error){
	// get https://xkcd.com/info.0.json, as this the current comic
	// which should have the maximum number that the for loop
	// will need to go up to
	comicContents, err := fetchComic("")

	if err != nil {
		return -1, err
	}

	return comicContents.Num, nil
}

func DownloadComics() error {
	files, err := ioutil.ReadDir(StorageDir)
	if err != nil {
		log.Fatal("Failure to read storage directory")
	}

	numOfComics, numErr := getCurrentNumberOfComics()

	if numErr != nil {
		log.Fatal("Failure to get current number of comics")
	}

	if len(files) < numOfComics {
		var i int
		if numOfComics - len(files) != numOfComics {
			// 1 to go past the last file and one more to account for the 404 page
			i = len(files) + 1 + 1
		} else {
			i = 1
		}

		for ; i <= numOfComics; i++ {
			// There's no comic 404
			if i == 404 {
				continue
			}

			log.Printf("Fetching comic %d\n", i)
			comicContents, err := fetchComic(strconv.Itoa(i))

			if err != nil {
				log.Fatalf("Failed to fetch comic %d: %s", i, err)
			}
			log.Printf("Fetched comic %d\n", i)

			file, marshalErr := json.MarshalIndent(comicContents, "", "  ")

			if marshalErr != nil {
				return marshalErr
			}

			ioutil.WriteFile(StorageDir+ "/xkcd." + strconv.Itoa(i) + ".json", file, os.FileMode(0444))
			log.Printf("Saved comic %d\n", i)
			time.Sleep(1 * time.Second)
		}
	}

	return nil
}
