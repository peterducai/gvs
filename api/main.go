package api

import (
	"fmt"
	"net/http"

	"github.com/peterducai/jobdsigner/models"
)

//Main adds job to pipeline
func Main(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	fmt.Fprintf(w, "[main] GVS %d.%d.%d patch %s\nyour source versioning system.\nstart at %s",
		models.JVersion.MAJOR,
		models.JVersion.MINOR,
		models.JVersion.PATCH,
		models.JVersion.HASH,
		models.JVersion.Startime)
}
