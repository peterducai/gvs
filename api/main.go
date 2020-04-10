package api

import (
	"fmt"
	"net/http"

	"github.com/peterducai/gvs/models"
)

//Main adds job to pipeline
func Main(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	fmt.Fprintf(w, "GVS %d.%d.%d patch %s\nyour source versioning system.\nstarttime: %s",
		models.GvsVersion.MAJOR,
		models.GvsVersion.MINOR,
		models.GvsVersion.PATCH,
		models.GvsVersion.HASH,
		models.GvsVersion.Startime)
}
