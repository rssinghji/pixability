package api

import (
	"fmt"
	"time"
)

// A singleton pointer to return same object every time - An overkill for this exercise but anyway.
var campaignScheduler *CampaignScheduler

// Map to store days for a month
var monthlyMap = make(map[string]int)

// Map to store already added campaigns to avoid redundancy
var addedCampaigns = make(map[string]string)

func init() {
	campaignScheduler = new(CampaignScheduler)
	campaignScheduler.CampaignsMap = make(map[string]YearlyArranged)
	campaignScheduler.ForGapsMap = make(map[string]map[string]map[int]int)
	monthlyMap["01"] = 31
	monthlyMap["02"] = 28
	monthlyMap["03"] = 31
	monthlyMap["04"] = 30
	monthlyMap["05"] = 31
	monthlyMap["06"] = 30
	monthlyMap["07"] = 31
	monthlyMap["08"] = 31
	monthlyMap["09"] = 30
	monthlyMap["10"] = 31
	monthlyMap["11"] = 30
	monthlyMap["12"] = 31
	go updateMap()
}

// An update thread to update number of days in case if a leap year
func updateMap() {
	for {
		now := time.Now()
		year := now.Year()
		if year%4 == 0 {
			monthlyMap["02"] = 29
		}
		nextYear := year + 1
		checkDate := fmt.Sprintf("%d", nextYear) + "-" + "01" + "-" + "01"
		nextUpdate, _ := time.Parse("2006-01-02", checkDate)
		duration := nextUpdate.Sub(now)
		time.Sleep(duration)
	}
}

// GetCampaignScheduler : A singleton class to Get the pointer to the class scheduler
func GetCampaignScheduler() *CampaignScheduler {
	if campaignScheduler == nil {
		campaignScheduler = new(CampaignScheduler)
	}
	return campaignScheduler
}

// CampaignsInfo : Struct to store monthly campaigns
type CampaignsInfo struct {
	Month     string     `json:"month"`
	Campaigns []Campaign `json:"campaigns"`
}

// YearlyArranged : Struct to store campaigns arranged yearly-monthly
type YearlyArranged struct {
	Year string          `json:"-"`
	Info []CampaignsInfo `json:"campaigns"`
}

// Campaign ; Struct to store the basic Campaign as per question
type Campaign struct {
	CampaignName string `json:"campaign_name"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
}

// CampaignScheduler : Object to store all APIs related data
type CampaignScheduler struct {
	CampaignsMap map[string]YearlyArranged         `json:"campaigns,omitempty"`
	ForGapsMap   map[string]map[string]map[int]int `json:"-"`
	Errors       []string                          `json:"errors,omitempty"`
	Message      string                            `json:"message,omitempty"`
}

// GapsResult : struct to store default gaps API result
type GapsResult struct {
	CampaignGaps []string `json:"gaps,omitempty"`
}

// ScheduleRequest : Struct to parse a schedule request
type ScheduleRequest struct {
	Campaigns []Campaign `json:"campaigns"`
}
