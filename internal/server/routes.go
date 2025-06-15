package server

func (s *Server) setupRoutes() {
	//s.mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	_, _ = w.Write([]byte("Hello"))
	//}))

	Static(s.mux)

	// Register the PDF handler
	s.mux.Handle(pdfPath+"/", pdfMux())
}
