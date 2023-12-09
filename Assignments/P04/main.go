// Package declaration
package main

// Import required packages
import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	// Slice of strings that contain URLs of images
	urls := []string{
		"https://unsplash.com/photos/hvdnff_bieQ/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		"https://unsplash.com/photos/HQaZKCDaax0/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		"https://images.unsplash.com/photo-1698778573682-346d219402b5?ixlib=rb-4.0.3&q=85&fm=jpg&crop=entropy&cs=srgb&w=640",
		"https://unsplash.com/photos/Bs2jGUWu4f8/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		
		// Add more image URLs
		"https://cdn.stocksnap.io/img-thumbs/960w/christmas-ribbon_FJRZC2S6XO.jpg",
		"https://cdn.stocksnap.io/img-thumbs/960w/flat%20lay-spa_2HEUUNACGX.jpg",
		"https://cdn.stocksnap.io/img-thumbs/960w/autumn-foliage_YVAKQAUZJ6.jpg",
		"https://cdn.stocksnap.io/img-thumbs/960w/travel-explore_RMN2RSOXPL.jpg",
		"https://cdn.stocksnap.io/img-thumbs/960w/food-food%20photography_GNAFKFNM2R.jpg",
	}

	// Print Program Heading
	fmt.Println("P04 - Image Downloader")

	// Define the number of workers for concurrent downloads
	numWorkers := 4

	// Sequential download
	start := time.Now()
	downloadImagesSequential(urls)
	fmt.Printf("Sequential download took: %v\n", time.Since(start))

	// Concurrent download
	start = time.Now()
	downloadImagesConcurrent(urls, numWorkers)
	fmt.Printf("Concurrent download took: %v\n", time.Since(start))
}

// Helper function to download and save a single image.
func downloadImage(url, filename string, wg *sync.WaitGroup) {
	// Decrement the WaitGroup counter by one; done when this function returns.
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	// Create the directories if they don't exist
	if err := os.MkdirAll(filepath.Dir(filename), os.ModePerm); err != nil {
		log.Println("Error creating directories:", err)
		return
	}

	// Create a new `http.Request` object.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Create a new `http.Client` object.
	client := &http.Client{}

	// Do the request and get the response.
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	// Check the response status code.
	if resp.StatusCode != http.StatusOK {
		log.Println("Response status code:", resp.StatusCode)
		return
	}

	// Create a new file to save the image to.
	f, err := os.Create(filename)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close() // Ensure the file is closed even in case of an error.

	// Copy the image from the response body to the file.
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		log.Println(err)
	}
}

func downloadImagesConcurrent(urls []string, numWorkers int) {
	// Used for synchronization between goroutines.
	var wg sync.WaitGroup

	// Create a buffered channel to limit the number of concurrent downloads
	ch := make(chan struct{}, numWorkers)

	// Loop through slice of urls
	for i, url := range urls {
		// Add one to wait group for each goroutine
		wg.Add(1)

		// Use an empty struct to represent a worker
		ch <- struct{}{}

		go func(index int, u string) {
			// Decrement the WaitGroup counter by one; done when this function returns.
			defer func() {
				<-ch // Release a worker when done
				wg.Done()
			}()

			// Download the image
			filename := fmt.Sprintf("images/Image%d.jpg", index+1)
			downloadImage(u, filename, nil)
		}(i, url)
	}

	// Wait until all goroutines are done
	wg.Wait()
}

func downloadImagesSequential(urls []string) {
	// Loop through slice of urls, start downloading images sequentially
	for i, url := range urls {
		filename := fmt.Sprintf("images/Image%d.jpg", i+1)
		downloadImage(url, filename, nil)
	}
}
