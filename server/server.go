package server

import "net/http"

// BuildServer builds server with default params, except of provided by args
// Used to simplify runtime instance creation
func BuildServer(addr string) *http.Server {
	// Register handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/detect", DetectHandler)
	mux.HandleFunc("/translate", TranslateHandler)
	mux.HandleFunc("/translate/bulk", BulkTranslateHandler)
	// Build server with defaults
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	// Return
	return server
}
