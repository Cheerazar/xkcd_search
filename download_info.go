package main

//const (
//	BaseUrl = "https://xkcd.com/"
//	JsonApiEnding = "info.0.json"
//	StorageDir = "./storage"
//)
//
//type ComicContents struct {
//	Alt string
//	Num int
//	SafeTitle string `json:"safe_title"`
//	Title string
//	Transcript string
//}
//
//// When fetching the current comic, comicNum is "",
//// otherwise it's the number of the comic to fetch
//func fetchComic(comicNum string) (ComicContents, error) {
//	var comic string
//	if comicNum == "" {
//		comic = ""
//	} else {
//		comic = comicNum + "/"
//	}
//
//	resp, err := http.Get(BaseUrl + comic + JsonApiEnding)
//
//	if err != nil {
//		return ComicContents{}, err
//	}
//
//	if resp.StatusCode != http.StatusOK {
//		resp.Body.Close()
//		return ComicContents{}, fmt.Errorf("fetch of comic %s failed: %s", comicNum, resp.Status)
//	}
//
//	var contents ComicContents
//	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
//		resp.Body.Close()
//		return ComicContents{}, err
//	}
//
//	resp.Body.Close()
//
//	return contents, nil
//}
//
//func getCurrentNumberOfComics() (int, error){
//	// get https://xkcd.com/info.0.json, as this the current comic
//	// which should have the maximum number that the for loop
//	// will need to go up to
//	comicContents, err := fetchComic("")
//
//	if err != nil {
//		return -1, err
//	}
//
//	return comicContents.Num, nil
//}
//
//func DownloadComics() error {
//	files, err := ioutil.ReadDir(StorageDir)
//	if err != nil {
//		log.Fatal("Failure to read storage directory")
//	}
//
//	numOfComics, numErr := getCurrentNumberOfComics()
//
//	if numErr != nil {
//		log.Fatal("Failure to get current number of comics")
//	}
//
//	if len(files) < numOfComics {
//		var i int
//		if numOfComics - len(files) != numOfComics {
//			// 1 to go past the last file and one more to account for the 404 page
//			i = len(files) + 1 + 1
//		} else {
//			i = 1
//		}
//
//		for ; i <= numOfComics; i++ {
//			// There's no comic 404
//			if i == 404 {
//				continue
//			}
//
//			log.Printf("Fetching comic %d\n", i)
//			comicContents, err := fetchComic(strconv.Itoa(i))
//
//			if err != nil {
//				log.Fatalf("Failed to fetch comic %d: %s", i, err)
//			}
//			log.Printf("Fetched comic %d\n", i)
//
//			file, marshalErr := json.MarshalIndent(comicContents, "", "  ")
//
//			if marshalErr != nil {
//				return marshalErr
//			}
//
//			ioutil.WriteFile(StorageDir+ "/xkcd." + strconv.Itoa(i) + ".json", file, os.FileMode(0444))
//			log.Printf("Saved comic %d\n", i)
//			time.Sleep(1 * time.Second)
//		}
//	}
//
//	return nil
//}