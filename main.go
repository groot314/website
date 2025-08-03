package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/starfederation/datastar-go/datastar"
)

//go:embed index.html
var HTML []byte

//go:embed about.html
var ABOUT []byte

const port = 8080

var userCount int

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(HTML)
	})

	r.Get("/home", func(w http.ResponseWriter, r *http.Request) {
		sse := datastar.NewSSE(w, r)

		lineMsg := "hello user" + strconv.Itoa(userCount)
		lineMsg2 := "the time is " + time.Now().UTC().Format("2006-01-02 15:04:05 UTC")
		if err := sse.PatchElements(`<div id="content-main">` + consoleMessages([]string{lineMsg, lineMsg2}) + `</div>`); err != nil {
			return
		}
	})

	r.Get("/console-hello", func(w http.ResponseWriter, r *http.Request) {
		sse := datastar.NewSSE(w, r)

		userCount++
		lineMsg := "hello user" + strconv.Itoa(userCount)
		for i := range len(lineMsg) {
			if err := sse.PatchElements(consoleMessages([]string{lineMsg[:i+1]})); err != nil {
				return
			}
			time.Sleep(100 * time.Millisecond)
		}

		lineMsg2 := "the time is " + time.Now().UTC().Format("2006-01-02 15:04:05 UTC")
		for i := range len(lineMsg2) {
			if err := sse.PatchElements(consoleMessages([]string{lineMsg, lineMsg2[:i+1]})); err != nil {
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	})

	r.Get("/about", func(w http.ResponseWriter, r *http.Request) {
		sse := datastar.NewSSE(w, r)

		element := string(ABOUT)
		if err := sse.PatchElements(element); err != nil {
			return
		}
	})

	log.Printf("Starting server on http://localhost:%d", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), r); err != nil {
		panic(err)
	}
}

func consoleMessages(lines []string) string {

	result := `
		<div
		  id="console"
		  class="block bg-black border-2 border-gray-500 rounded-sm max-w-xl h-48"
		>
		`
	messageConsoleLine := func(msg string) string {
		start := `<div class="flex">
			  <p class="font-bold mx-2">></p>
			  <p class="text-green-500">`
		end := `</p></div>`

		return start + msg + end

	}
	for _, line := range lines {
		result += messageConsoleLine(line)
	}

	return result + "</div>"

}
