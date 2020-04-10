package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/peterducai/gvs/models"
	"github.com/peterducai/jobdsigner/api"
)

func main() {

	models.GvsVersion.MAJOR = 0
	models.GvsVersion.MINOR = 0
	models.GvsVersion.PATCH = 1
	models.GvsVersion.HASH = "a12c1"
	models.GvsVersion.Startime = time.Now().Format(time.RFC850)

	//models.PipelinePool = append(new(models.Pipeline{}))

	fmt.Printf("GVS %d.%d.%d %s\n", models.GvsVersion.MAJOR, models.GvsVersion.MINOR, models.GvsVersion.PATCH, models.GvsVersion.HASH)

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
