package server

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"net/http"
)

const (
	pdfPath = "/pdf"
)

func pdfMux() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		page(PageProps{
			Title: "Test PDF",
		},
			Div(
				Class("p-4 bg-gray-100 rounded shadow"),
				Text("hello world"),
			),
		).Render(w)
	})

	return http.StripPrefix(pdfPath, m)
}
