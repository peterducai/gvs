package api

import (
	"fmt"
	"net/http"

	"github.com/peterducai/jobdsigner/models"
)

//JobAdd adds job to pipeline
func JobAdd(w http.ResponseWriter, r *http.Request) {
	var job = new(models.Job)
	job.ID = 100
	job.Title = "example job1"
	job.Interaction = "run"
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	fmt.Fprintf(w, "job add: %s\n", job.Title)
}
