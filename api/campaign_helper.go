package api

import (
	"fmt"
	"strings"
	"time"
)

func addCampaigns(campaigns []Campaign) (*CampaignScheduler, int, error) {
	scheduler := GetCampaignScheduler()
	successfullyAddedCampaigns := 0

	for _, campaign := range campaigns {
		campaignKey := campaign.CampaignName + campaign.StartDate + campaign.EndDate
		if _, ok := addedCampaigns[campaignKey]; ok {
			continue
		}
		if campaign.CampaignName == "" {
			errorString := "Problem in adding the campaign. No Campaign Name defined in request. "
			scheduler.Errors = append(scheduler.Errors, errorString)
			continue
		}
		start, err := time.Parse("2006-01-02", campaign.StartDate)
		if err != nil {
			errorString := "Problem in adding the campaign. start_date not correctly defined in request. " + err.Error()
			scheduler.Errors = append(scheduler.Errors, errorString)
			continue
		}
		end, err := time.Parse("2006-01-02", campaign.EndDate)
		if err != nil {
			errorString := "Problem in adding the campaign. end_date not correctly defined in request. " + err.Error()
			scheduler.Errors = append(scheduler.Errors, errorString)
			continue
		}
		if end.Before(start) {
			errorString := "Problem in adding the campaign. end_date cannot be earlier than start_date."
			scheduler.Errors = append(scheduler.Errors, errorString)
			continue
		}
		dateParts := strings.Split(campaign.StartDate, "-")
		year := dateParts[0]
		month := dateParts[1]
		dayBegin := start.Day()
		dayEnd := end.Day()

		campaignMap := make(map[int]int)
		for index := dayBegin; index <= dayEnd; index++ {
			campaignMap[index] = index
			if myMonthlyMap, ok := scheduler.ForGapsMap[year]; ok {
				if monthData, present := myMonthlyMap[month]; present {
					monthData[index] = index
					myMonthlyMap[month] = monthData
					scheduler.ForGapsMap[year] = myMonthlyMap
				} else {
					myMap := make(map[int]int)
					myMap[index] = index
					myMonthlyMap[month] = myMap
					scheduler.ForGapsMap[year] = myMonthlyMap
				}
			} else {
				myMonthlyMap := make(map[string]map[int]int)
				myMap := make(map[int]int)
				myMap[index] = index
				myMonthlyMap[month] = myMap
				scheduler.ForGapsMap[year] = myMonthlyMap
			}
		}

		yearlyData := YearlyArranged{}
		ok, monthFound := false, false
		if yearlyData, ok = scheduler.CampaignsMap[year]; ok {
			for index, info := range yearlyData.Info {
				if month == info.Month {
					monthFound = true
					info.Campaigns = append(info.Campaigns, campaign)
					yearlyData.Info[index] = info
				}
			}
			if !monthFound {
				monthly := CampaignsInfo{Month: month}
				monthly.Campaigns = append(monthly.Campaigns, campaign)
				yearlyData.Info = append(yearlyData.Info, monthly)
			}
		} else {
			yearlyData = YearlyArranged{Year: year}
			monthly := CampaignsInfo{Month: month}
			monthly.Campaigns = append(monthly.Campaigns, campaign)
			yearlyData.Info = append(yearlyData.Info, monthly)
		}

		scheduler.CampaignsMap[year] = yearlyData
		addedCampaigns[campaignKey] = campaign.StartDate + " to " + campaign.EndDate
		successfullyAddedCampaigns++
	}
	return scheduler, successfullyAddedCampaigns, nil
}

func getAllGaps(year, month string) []string {
	result := []string{}
	monthlyData := campaignScheduler.ForGapsMap[year]
	if month == "" {
		for index := 1; index <= 12; index++ {
			var monthKey string
			if index >= 1 && index < 10 {
				monthKey = fmt.Sprintf("0%d", index)
			} else {
				monthKey = fmt.Sprintf("%d", index)
			}
			if value, ok := monthlyData[monthKey]; ok {
				monthGaps := getMonthlyGaps(value, year, monthKey)
				result = append(result, monthGaps...)
			} else {
				continue
			}
		}
	} else {
		result = getMonthlyGaps(monthlyData[month], year, month)
	}
	return result
}

func getCategorizedGaps(year string) map[string]map[string][]string {
	result := make(map[string]map[string][]string)
	monthlyGapMap := make(map[string][]string)
	monthlyData := campaignScheduler.ForGapsMap[year]

	for index := 1; index <= 12; index++ {
		var monthKey string
		if index >= 1 && index < 10 {
			monthKey = fmt.Sprintf("0%d", index)
		} else {
			monthKey = fmt.Sprintf("%d", index)
		}
		if value, ok := monthlyData[monthKey]; ok {
			monthlyGaps := getMonthlyGaps(value, year, monthKey)
			monthlyGapMap[monthKey] = monthlyGaps
		} else {
			continue
		}
	}

	result[year] = monthlyGapMap
	return result
}

func getMonthlyGaps(monthlyValues map[int]int, year, month string) []string {
	result := []string{}
	beginFound, endFound := false, false
	gapBegin, gapEnd := "", ""
	daysInMonth := monthlyMap[month]
	for index := 1; index <= daysInMonth; index++ {
		if _, ok := monthlyValues[index]; ok {
			if beginFound && index >= 1 && index < 10 {
				gapEnd = year + "/" + month + "/" + "0" + fmt.Sprintf("%d", index-1)
				endFound = true
			} else if beginFound && index >= 10 {
				gapEnd = year + "/" + month + "/" + fmt.Sprintf("%d", index-1)
				endFound = true
			}
		} else {
			if index >= 1 && index < 10 && !beginFound {
				gapBegin = year + "/" + month + "/" + "0" + fmt.Sprintf("%d", index)
				beginFound = true
			} else if index >= 10 && !beginFound {
				gapBegin = year + "/" + month + "/" + fmt.Sprintf("%d", index)
				beginFound = true
			}
		}
		if beginFound && endFound {
			result = append(result, gapBegin+"-"+gapEnd)
			beginFound = false
			endFound = false
		}
	}
	return result
}
