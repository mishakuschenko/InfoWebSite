package internal

import (
	"fmt"
	"net/http"
)

func Home_page(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "Let's GO!")
}

func About_page(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "about page")
}
