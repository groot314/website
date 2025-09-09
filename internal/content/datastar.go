package content

import (
	_ "embed"

	"net/http"
	"strconv"
	"time"
	"website/internal/elements"

	"github.com/starfederation/datastar-go/datastar"
)

var userCount int

func HelloConsole(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)

	userCount++
	lineMsg := "hello user" + strconv.Itoa(userCount)
	for i := range len(lineMsg) {
		if err := sse.PatchElements(elements.ConsoleMessages([]string{lineMsg[:i+1]})); err != nil {
			return
		}
		time.Sleep(100 * time.Millisecond)
	}

	lineMsg2 := "the time is " + time.Now().UTC().Format("2006-01-02 15:04:05 UTC")
	for i := range len(lineMsg2) {
		if err := sse.PatchElements(elements.ConsoleMessages([]string{lineMsg, lineMsg2[:i+1]})); err != nil {
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
}

//go:embed html/about.html
var aboutHTML []byte

func GetAbout(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)

	element := string(aboutHTML)
	if err := sse.PatchElements(element); err != nil {
		return
	}
}

func GetHomeContent(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)

	lineMsg := "hello user" + strconv.Itoa(userCount)
	lineMsg2 := "the time is " + time.Now().UTC().Format("2006-01-02 15:04:05 UTC")
	if err := sse.PatchElements(`<div id="content-main">` + elements.ConsoleMessages([]string{lineMsg, lineMsg2}) + `</div>`); err != nil {
		return
	}
}
