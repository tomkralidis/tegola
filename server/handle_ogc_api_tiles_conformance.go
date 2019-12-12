package server

import (
	"encoding/json"
	"net/http"
)

type OgcApiTilesConformance struct {
    ConformsTo []string   `json:"conformsTo"`
	//Maps    []CapabilitiesMap `json:"maps"`
}

type HandleOgcApiTilesConformance struct{}

func (req HandleOgcApiTilesConformance) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// new capabilities struct
	conformance := OgcApiTilesConformance{
	}

    conformance.ConformsTo = []string{"http://www.opengis.net/spec/ogcapi-tiles-1/1.0/req/core",
                                      "http://www.opengis.net/spec/ogcapi-tiles-1/1.0/req/collections"}
    w.Header().Add("Content-Type", "application/json")

    // cache control headers (no-cache)
    w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
    w.Header().Add("Pragma", "no-cache")
    w.Header().Add("Expires", "0")


	// setup a new json encoder and encode our capabilities
	json.NewEncoder(w).Encode(conformance)
}
