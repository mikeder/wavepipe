package api

import (
	"log"
	"net/http"

	"github.com/mdlayher/wavepipe/common"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// StatusResponse represents the JSON response for /api/status
type StatusResponse struct {
	Error   *Error          `json:"error"`
	Status  *common.Status  `json:"status"`
	Metrics *common.Metrics `json:"metrics"`
}

// GetStatus returns the current server status, with an HTTP status and JSON
func GetStatus(res http.ResponseWriter, req *http.Request) {
	// Retrieve render
	r := context.Get(req, CtxRender).(*render.Render)

	// Output struct for songs request
	out := StatusResponse{}

	// Check API version
	if version, ok := mux.Vars(req)["version"]; ok {
		// Check if this API call is supported in the advertised version
		if !apiVersionSet.Has(version) {
			r.JSON(res, 400, errRes(400, "unsupported API version: "+version))
			return
		}
	}

	// Retrieve current server status
	status, err := common.ServerStatus()
	if err != nil {
		log.Println(err)
		r.JSON(res, 500, serverErr)
		return
	}

	// Copy into response
	out.Status = status

	// If requested, fetch additional metrics (not added by default due to full table scans in database)
	if req.URL.Query().Get("metrics") == "true" {
		metrics, err := common.ServerMetrics()
		if err != nil {
			log.Println(err)
			r.JSON(res, 500, serverErr)
			return
		}

		// Return metrics
		out.Metrics = metrics
	}

	// HTTP 200 OK with JSON
	out.Error = nil
	r.JSON(res, 200, out)
	return
}
