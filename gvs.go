package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/peterducai/gvs/api"
	"github.com/peterducai/gvs/models"
)

func main() {

	models.JVersion.MAJOR = 0
	models.JVersion.MINOR = 0
	models.JVersion.PATCH = 1
	models.JVersion.HASH = "a12b"
	models.JVersion.Startime = time.Now().Format(time.RFC850)

	//models.PipelinePool = append(new(models.Pipeline{}))

	fmt.Printf("Job dSigner %d.%d.%d %s\n", models.JVersion.MAJOR, models.JVersion.MINOR, models.JVersion.PATCH, models.JVersion.HASH)

	mux := http.NewServeMux()

	//HANDLERS
	mux.HandleFunc("/", api.Main)
	mux.HandleFunc("/job/add", api.JobAdd)
	mux.HandleFunc("/about", api.About)

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	srv := &http.Server{
		Addr:         ":8443",
		Handler:      mux,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	fmt.Println("server is running at https://localhost:8443")
	log.Fatal(srv.ListenAndServeTLS("tls.crt", "tls.key"))

}
