package server

import (
	"encoding/json"
	"net/http"
    "net/url"
)

type OgcApiTiles struct {
    Title string            `json:"title"`
    Description string      `json:"description"`
    Links []LinkMap         `json:"links"`
}

type LinkMap struct {
	Href string            `json:"href"`
    Rel string             `json:"rel"`
	Type string            `json:"type"`
    Title string           `json:"title"`
}

type HandleOgcApiTiles struct{}

func (req HandleOgcApiTiles) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// new capabilities struct
	tiles := OgcApiTiles{
        Title: "OGC-API-Tiles",
        Description: "OGC API Tiles",
	}

    // parse our query string
    //var query = r.URL.Query()

    // iterate our registered maps

    debugQuery := url.Values{}

    conformanceLink := LinkMap{
        Href:       buildCapabilitiesURL(r, []string{"ogc-api-tiles", "conformance"}, debugQuery),
        Rel:        "conformance",
        Type:       "application/json",
        Title:      "the list of conformance classes implemented by this API",
    }
    tiles.Links = append(tiles.Links, conformanceLink)

    collectionsLink := LinkMap{
        Href:       buildCapabilitiesURL(r, []string{"ogc-api-tiles", "collections"}, debugQuery),
        Rel:        "data",
        Type:       "application/json",
        Title:      "The collections in the dataset in JSON",
    }
    tiles.Links = append(tiles.Links, collectionsLink)

    mapsLink := LinkMap{
        Href:       buildCapabilitiesURL(r, []string{"ogc-api-tiles", "tiles"}, debugQuery),
        Rel:        "tiles",
        Type:       "application/json",
        Title:      "Access the data as multi-layer vector tiles",
    }
    tiles.Links = append(tiles.Links, mapsLink)

    stylesLink := LinkMap{
        Href:       buildCapabilitiesURL(r, []string{"ogc-api-tiles", "styles"}, debugQuery),
        Rel:        "tiles",
        Type:       "application/json",
        Title:      "Styles to render the data",
    }
    tiles.Links = append(tiles.Links, stylesLink)

    w.Header().Add("Content-Type", "application/json")

    // cache control headers (no-cache)
    w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
    w.Header().Add("Pragma", "no-cache")
    w.Header().Add("Expires", "0")

	// setup a new json encoder and encode our capabilities
	json.NewEncoder(w).Encode(tiles)
}
