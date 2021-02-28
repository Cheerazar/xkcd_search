package main

//func FindSearchTerms(searchTerms string) (matches int, filesSearched int) {
//	lowerTerms := strings.ToLower(searchTerms)
//	files, err := os.ReadDir(StorageDir)
//
//	if err != nil {
//		log.Fatalf("Failed to read directory in FindSearchTerms: %s\n", err)
//	}
//
//	for _, file := range files {
//		data, rfErr := os.ReadFile(file.Name())
//
//		if rfErr != nil {
//			log.Fatalf("Failed to open file: %s\nWith error: %s\n", file.Name(), rfErr)
//		}
//
//		var comicInfo ComicContents
//
//		unmarshalErr := json.Unmarshal(data, &comicInfo)
//
//		if unmarshalErr != nil {
//			log.Fatalf("Failed to unmarshal file: %s\nError: %s\n", file.Name(), unmarshalErr)
//		}
//
//
//		if strings.Contains(strings.ToLower(comicInfo.Alt), lowerTerms) {
//			matches++
//			fmt.Printf("Match %d\nTranscript: %s\nComic URL: %s\n", matches, comicInfo.Transcript, BaseUrl+ strconv.Itoa(comicInfo.Num))
//			continue
//		}
//
//		if strings.Contains(strings.ToLower(comicInfo.SafeTitle), lowerTerms) {
//			matches++
//			fmt.Printf("Match %d\nTranscript: %s\nComic URL: %s\n", matches, comicInfo.SafeTitle, BaseUrl+ strconv.Itoa(comicInfo.Num))
//			continue
//		}
//
//		if strings.Contains(strings.ToLower(comicInfo.Title), lowerTerms) {
//			matches++
//			fmt.Printf("Match %d\nTranscript: %s\nComic URL: %s\n", matches, comicInfo.Title, BaseUrl+ strconv.Itoa(comicInfo.Num))
//			continue
//		}
//
//		if strings.Contains(strings.ToLower(comicInfo.Transcript), lowerTerms) {
//			matches++
//			fmt.Printf("Match %d\nTranscript: %s\nComic URL: %s\n", matches, comicInfo.Transcript, BaseUrl+ strconv.Itoa(comicInfo.Num))
//			continue
//		}
//	}
//
//	return matches, len(files)
//}