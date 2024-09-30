package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type CardData struct {
	Image string
	Text  string
}

func main() {
	http.HandleFunc("/data/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/data/"):]

		println(id)
		var data CardData
		rand.Seed(time.Now().UnixNano())

		messages := []string{
			"You take the seed of a dream, a goal, a vision , a passion and nurture it to the point it grows",
			"You have the ingenuity to bring alive dreams and hopes",
			"i love your conversations",
			"I want to be told EVERYTHING",
			"You put energy into curating your environment",
			"You preemtively take care of yourself and those around you",
			"You are a lioness, you are fierce defending what is right",
			"Your hype and communial ",
		}
		urls := []string{
			"https://raw.githubusercontent.com/FredMunene/nduta/main/l1.jpg",
			"https://raw.githubusercontent.com/FredMunene/nduta/main/l2.jpg",
			"https://raw.githubusercontent.com/FredMunene/nduta/main/l3.jpg",
			"https://raw.githubusercontent.com/FredMunene/nduta/main/l4.jpg",
			"https://raw.githubusercontent.com/FredMunene/nduta/main/l5.jpg",
			"https://raw.githubusercontent.com/FredMunene/nduta/main/l6.jpg",
			"https://raw.githubusercontent.com/FredMunene/nduta/main/l7.jpg",
			"https://raw.githubusercontent.com/FredMunene/nduta/main/l8.jpg",
			"https://raw.githubusercontent.com/FredMunene/nduta/main/l9.jpg",
			"https://raw.githubusercontent.com/FredMunene/nduta/main/l10.jpg",
			"https://raw.githubusercontent.com/FredMunene/nduta/main/l11.jpg",
			"https://raw.githubusercontent.com/FredMunene/nduta/main/l12.jpg",
			"https://raw.githubusercontent.com/FredMunene/nduta/main/l13.jpg",
		}

		randomIndex := rand.Intn(len(urls))
		randomText := rand.Intn(len(messages))

		// // url, err := fetchImages(urls[randomIndex])
		// if err != nil {
		// 	log.Fatal(err)
		// 	return
		// }
		// println(url)

		data = CardData{Image: urls[randomIndex], Text: messages[randomText]}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.Handle("/style.css", http.FileServer(http.Dir(".")))

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// func fetchImages(url string) (string,error) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return "", err
// 	}

// 	defer resp.Body.Close()

// 	if resp.StatusCode != 200 {
// 		log.Fatalf("Failed to fetch URL; status Code: %d", resp.StatusCode)
// 		return "", fmt.Errorf("failed to fetch URL; status Code: %d", resp.StatusCode)
// 	}

// 	// load html
// 	doc, err := goquery.NewDocumentFromReader(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 		return "", err
// 	}

// 	links := []string{}

// 	doc.Find("img").Each(func(index int, item *goquery.Selection) {
// 		imgSrc, exists := item.Attr("src")

// 		re := regexp.MustCompile(`w\d+-h\d+-no`)
// 		newSize := "w600-h400-no"
// 		updatedURL := re.ReplaceAllString(imgSrc, newSize)
// 		if exists {
// 			links = append(links, updatedURL)
// 		}
// 	})

// 	return links[1], nil
// }
