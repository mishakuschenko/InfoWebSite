package internal

import "net/http"

func Pages() {

	http.HandleFunc("/", Home_page)
	http.HandleFunc("/about/", About_page)
	http.ListenAndServe(":8000", nil)
}
						