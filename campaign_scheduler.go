package main

import (
	"fmt"
	"log"
	"net/http"
	api "pixability/api"
)

func main() {
	fmt.Println("Campaign Scheduling in Progress!")

	// Handlers for requests
	http.HandleFunc("/", api.HandleDefault)
	http.HandleFunc("/pixability/scheduleCampaigns", api.HandleScheduleCampaignsRequest)
	http.HandleFunc("/pixability/getGaps", api.HandleGetGapsRequest)
	http.HandleFunc("/pixability/getCategorizedGaps", api.HandleGetCategorizedGapsRequest)
	http.HandleFunc("/debug/getCampaigns", api.HandleGetCampaignsRequest)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
