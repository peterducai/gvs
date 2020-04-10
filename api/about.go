package api

import (
	"fmt"
	"net/http"

	"github.com/peterducai/gvs/models"
)

//About adds job to pipeline
func About(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	fmt.Fprintf(w, "about  %d.%d.%d %s\n", models.GvsVersion.MAJOR, models.GvsVersion.MINOR, models.GvsVersion.PATCH, models.GvsVersion.HASH)
}
