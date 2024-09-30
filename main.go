package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type CardData struct {
	Image string
	Text  string
}

func main() {
	http.HandleFunc("/data/", func(w http.ResponseWriter, r *http.Request) {
		// id : r.URL.Path[len("/data/"):]

		var data CardData
		rand.Seed(time.Now().UnixNano())

		messages := []string{
			"You take the seed of a dream, a goal, a vision , a passion and nurture it to the point it grows",
			"You have the ingenuity to bring alive dreams and hopes",
			"I love your conversations",
			"I love your voice. I want to be told EVERYTHING, every story",
			"You put energy into curating your environment",
			"You preemtively take care of yourself and those around you",
			"You are a lioness, you are fierce. You stand right",
			"You are hype and happiness ",
			"Your mind is amazing",
			"You are a babygirl, a cute one",
			"I love how feminine you are",
			"You trust in your guts",
			"You have faith, you believe tremendously",
			"You are kind and compassionate ",
			"You are funny, you make me laugh",
			" You are passionate about your goals",
			"You are intelligent, smart, my Korean baby",
		}
		urls := []string{
			"https://raw.githubusercontent.com/FredMunene/nduta/main/img/l1.jpg",
			"https://raw.githubusercontent.com/FredMunene/nduta/main/img/l2.jpg",
			"https://raw.githubusercontent.com/FredMunene/nduta/main/img/l3.jpg",
			"https://raw.githubusercontent.com/FredMunene/nduta/main/img/l4.jpg",
			"https://raw.githubusercontent.com/FredMunene/nduta/main/img/l5.jpg",
			"https://raw.githubusercontent.com/FredMunene/nduta/main/img/l6.jpg",
			"https://raw.githubusercontent.com/FredMunene/nduta/main/img/l7.jpg",
			"https://raw.githubusercontent.com/FredMunene/nduta/main/img/l8.jpg",
			"https://raw.githubusercontent.com/FredMunene/nduta/main/img/l9.jpg",
			"https://raw.githubusercontent.com/FredMunene/nduta/main/img/l11.jpeg",
			"https://raw.githubusercontent.com/FredMunene/nduta/main/img/l12.jpg",
			"https://raw.githubusercontent.com/FredMunene/nduta/main/img/l13.jpg",
		}

		randomIndex := rand.Intn(len(urls))
		randomText := rand.Intn(len(messages))

		// // url, err := fetchImages(urls[randomIndex])
		// if err != nil {
		// 	log.Fatal(err)
		// 	return
		// }
		// println(url)

		// println(randomIndex)

		data = CardData{Image: urls[randomIndex], Text: messages[randomText]}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.Handle("/style.css", http.FileServer(http.Dir(".")))

	port := os.Getenv("PORT")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
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
