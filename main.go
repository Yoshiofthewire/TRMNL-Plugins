package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ComicData holds the extracted information for the JSON response.
type ComicData struct {
	ImageURL string `json:"imageUrl"` // URL of the comic image
	Title    string `json:"title"`    // Caption/title of the comic
}

const targetURL = "https://www.thefarside.com/"

// fetchComicData fetches and parses the Far Side website to extract comic info.
func fetchComicData() (*ComicData, error) {
	// Make HTTP GET request
	res, err := http.Get(targetURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL %s: %w", targetURL, err)
	}
	defer res.Body.Close() // Ensure body is closed

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"request failed with status code %d",
			res.StatusCode,
		)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	// Find the first element with class "card-body"
	cardBody := doc.Find(".card-body").First()
	if cardBody.Length() == 0 {
		return nil, fmt.Errorf("could not find element with class 'card-body'")
	}

	// Find the image within the card-body and get its data-src
	img := cardBody.Find("img")
	imgSrc, exists := img.Attr("data-src")
	if !exists || imgSrc == "" {
		// Fallback: Sometimes it might use 'src' directly if JS hasn't run
		// or if the structure changes slightly. Check 'src' as well.
		imgSrc, exists = img.Attr("src")
		if !exists || imgSrc == "" {
			return nil, fmt.Errorf("could not find img with 'data-src' or 'src' attribute within card-body")
		}
	}

	// Find the figure caption within the card-body and get its text
	caption := cardBody.Find(".figure-caption").First()
	if caption.Length() == 0 {
		return nil, fmt.Errorf("could not find element with class 'figure-caption' within card-body")
	}
	title := strings.TrimSpace(caption.Text()) // Trim whitespace

	data := &ComicData{
		ImageURL: imgSrc,
		Title:    title,
	}

	return data, nil
}

// farsideHandler handles incoming web requests.
func farsideHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request for %s from %s", r.URL.Path, r.RemoteAddr)

	// Fetch the latest comic data
	comicData, err := fetchComicData()
	if err != nil {
		log.Printf("Error fetching comic data: %v", err)
		http.Error(
			w,
			"Failed to retrieve Far Side comic data.",
			http.StatusInternalServerError,
		)
		return
	}

	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode the data to JSON and write to response
	err = json.NewEncoder(w).Encode(comicData)
	if err != nil {
		// This error is less likely but possible (e.g., write error)
		log.Printf("Error encoding JSON response: %v", err)
		// Don't write another http.Error if headers are already sent
		return
	}
	log.Println("Successfully served comic data.")
}

func main() {
	// Register the handler function for the root path "/"
	http.HandleFunc("/", farsideHandler)

	port := "8123"
	log.Printf("Starting Far Side JSON server on port %s...", port)
	log.Printf("Access it at: http://localhost:%s", port)

	// Start the HTTP server
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
