package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// HandleDefault : Default handler for showing options
func HandleDefault(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	fmt.Fprintf(response, `{

"message":"Welcome to Campaign Scheduler. You've reached to the default usage. Use one of the two options:

	1. /pixability/scheduleCampaigns/: To add/schedule campaigns.
	
	2. /pixability/getGaps/: To get gaps in campaigns from the scheduler.

	OR

	Read the Readme for more about usage."

}`)
}

// HandleScheduleCampaignsRequest : Handler to add campaigns
func HandleScheduleCampaignsRequest(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	if request.Method != "POST" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(response, `{"error":"Wrong HTTP method specified. Only POST method allowed for this request"}`)
		return
	}
	campaignsSent := ScheduleRequest{}
	err := json.NewDecoder(request.Body).Decode(&campaignsSent)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(response, `{"error":"`+err.Error()+`"`)
		return
	}

	campaignScheduler, successCount, err := addCampaigns(campaignsSent.Campaigns)
	if err != nil {
		fmt.Println("Something went wrong. ", err.Error())
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, `{"error":"Something went wrong while sending the response.`+err.Error()+`"}`)
	}
	campaignScheduler.Message = "Success in adding " + fmt.Sprintf("%d", successCount) + " campaign(s)."
	if err != nil {
		fmt.Println("Something went wrong. ", err.Error())
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, `{"error":"Something went wrong while sending the response.`+err.Error()+`"}`)
	} else {
		response.WriteHeader(http.StatusOK)
		fmt.Fprintf(response, `{"message":"`+campaignScheduler.Message+`"}`)
	}
	return
}

// HandleGetGapsRequest : Handler to find gaps
func HandleGetGapsRequest(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	if request.Method != "GET" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(response, `{"error":"Wrong HTTP method specified. Only GET method allowed for this request"}`)
		return
	}

	today := strings.Split(time.Now().String(), " ")[0]
	year := request.URL.Query().Get("year")
	month := request.URL.Query().Get("month")
	if year == "" {
		year = strings.Split(today, "-")[0]
	}

	result := getAllGaps(year, month)
	gaps := GapsResult{CampaignGaps: result}
	jsonResponse, err := json.Marshal(gaps)
	if err != nil {
		fmt.Println("Something went wrong. ", err.Error())
		response.WriteHeader(500)
		fmt.Fprintf(response, `{"error":"Something went wrong while sending the response."}`)
	} else {
		response.WriteHeader(200)
		fmt.Fprintf(response, "%s", string(jsonResponse[:]))
	}
	return
}

// HandleGetCategorizedGapsRequest : Handler to find gaps categorized by year and month
func HandleGetCategorizedGapsRequest(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	if request.Method != "GET" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(response, `{"error":"Wrong HTTP method specified. Only GET method allowed for this request"}`)
		return
	}

	today := strings.Split(time.Now().String(), " ")[0]
	year := request.URL.Query().Get("year")
	if year == "" {
		year = strings.Split(today, "-")[0]
	}

	result := getCategorizedGaps(year)
	jsonResponse, err := json.Marshal(result)
	if err != nil {
		fmt.Println("Something went wrong. ", err.Error())
		response.WriteHeader(500)
		fmt.Fprintf(response, `{"error":"Something went wrong while sending the response."}`)
	} else {
		response.WriteHeader(200)
		fmt.Fprintf(response, "%s", string(jsonResponse[:]))
	}
	return
}

// HandleGetCampaignsRequest : Handler to a debug endpoint to verify campaigns edition
func HandleGetCampaignsRequest(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(campaignScheduler.CampaignsMap)
	if err != nil {
		fmt.Println("Something went wrong. ", err.Error())
		response.WriteHeader(500)
		fmt.Fprintf(response, `{"error":"Something went wrong while sending the response."}`)
	} else {
		response.WriteHeader(200)
		fmt.Fprintf(response, "%s", string(jsonResponse[:]))
	}
	return
}
