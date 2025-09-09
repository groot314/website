package content

import (
	_ "embed"
	"net/http"
)

//go:embed html/home.html
var homeHTML []byte

func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Write(homeHTML)
}
