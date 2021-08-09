package api

import (
	"strings"
	"testing"
)

func TestAddCampaignsSingle(test *testing.T) {
	campaigns := []Campaign{}
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_1", StartDate: "2021-01-01", EndDate: "2021-01-05"})
	_, count, err := addCampaigns(campaigns)
	if err != nil || len(campaigns) != count {
		test.Log(test.Name() + " FAILED\n")
	}
}

func TestAddCampaignsMultiple(test *testing.T) {
	campaigns := []Campaign{}
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_1", StartDate: "2021-01-01", EndDate: "2021-01-05"})
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_2", StartDate: "2021-01-10", EndDate: "2021-01-15"})
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_3", StartDate: "2021-01-20", EndDate: "2021-01-25"})
	_, count, err := addCampaigns(campaigns)
	if err != nil || len(campaigns) != count {
		test.Log(test.Name() + " FAILED\n")
	} else {
		test.Log(test.Name() + " PASSED\n\n")
	}
}

func TestAddCampaignsNoCampaignName(test *testing.T) {
	campaigns := []Campaign{}
	campaigns = append(campaigns, Campaign{CampaignName: "", StartDate: "2021-01-01", EndDate: "2021-01-05"})
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_2", StartDate: "2021-01-10", EndDate: "2021-01-15"})
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_3", StartDate: "2021-01-20", EndDate: "2021-01-25"})
	_, count, err := addCampaigns(campaigns)
	if err != nil || len(campaigns)-1 != count {
		test.Log(test.Name() + " FAILED\n")
	} else {
		test.Log(test.Name() + " PASSED\n")
	}
}

func TestAddCampaignsBadStartDate(test *testing.T) {
	campaigns := []Campaign{}
	campaigns = append(campaigns, Campaign{CampaignName: "", StartDate: "2021-125/10001", EndDate: "2021-01-05"})
	_, count, err := addCampaigns(campaigns)
	if err != nil || len(campaigns)-1 != count {
		test.Log(test.Name() + " PASSED\n")
	} else {
		test.Log(test.Name() + " FAILED\n")
	}
}

func TestAddCampaignsBadEndDate(test *testing.T) {
	campaigns := []Campaign{}
	campaigns = append(campaigns, Campaign{CampaignName: "", StartDate: "2021-01-01", EndDate: "2021-1201-201//5"})
	_, count, err := addCampaigns(campaigns)
	if err != nil || len(campaigns)-1 != count {
		test.Log(test.Name() + " PASSED\n")
	} else {
		test.Log(test.Name() + " FAILED\n")
	}
}

func TestAddCampaignsInvalidDate(test *testing.T) {
	campaigns := []Campaign{}
	campaigns = append(campaigns, Campaign{CampaignName: "", StartDate: "2021-02-01", EndDate: "2021-01-01"})
	_, count, err := addCampaigns(campaigns)
	if err != nil || len(campaigns)-1 != count {
		test.Log(test.Name() + " PASSED\n")
	} else {
		test.Log(test.Name() + " FAILED\n")
	}
}

func TestAddCampaignsMultipleMonths(test *testing.T) {
	campaigns := []Campaign{}
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_1", StartDate: "2021-01-01", EndDate: "2021-01-05"})
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_2", StartDate: "2021-01-10", EndDate: "2021-01-15"})
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_3", StartDate: "2021-01-20", EndDate: "2021-01-25"})

	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_1", StartDate: "2021-02-01", EndDate: "2021-02-05"})
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_2", StartDate: "2021-02-10", EndDate: "2021-02-15"})
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_3", StartDate: "2021-02-20", EndDate: "2021-02-25"})

	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_1", StartDate: "2021-03-01", EndDate: "2021-03-05"})
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_2", StartDate: "2021-03-10", EndDate: "2021-03-15"})
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_3", StartDate: "2021-03-20", EndDate: "2021-03-25"})

	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_1", StartDate: "2021-04-01", EndDate: "2021-04-05"})
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_2", StartDate: "2021-04-10", EndDate: "2021-04-15"})
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_3", StartDate: "2021-04-20", EndDate: "2021-04-25"})

	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_1", StartDate: "2021-11-01", EndDate: "2021-11-05"})
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_2", StartDate: "2021-11-10", EndDate: "2021-11-15"})
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_3", StartDate: "2021-11-20", EndDate: "2021-11-25"})

	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_1", StartDate: "2021-12-01", EndDate: "2021-12-05"})
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_2", StartDate: "2021-12-10", EndDate: "2021-12-15"})
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_3", StartDate: "2021-12-20", EndDate: "2021-12-25"})
	_, count, err := addCampaigns(campaigns)
	if err != nil || len(campaigns) != count {
		test.Log(test.Name() + " FAILED\n")
	} else {
		test.Log(test.Name() + " PASSED\n")
	}
}

func TestMonthlyGaps(test *testing.T) {
	campaigns := []Campaign{}
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_1", StartDate: "2021-01-01", EndDate: "2021-01-05"})
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_2", StartDate: "2021-01-10", EndDate: "2021-01-15"})
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_3", StartDate: "2021-01-20", EndDate: "2021-01-25"})

	_, count, err := addCampaigns(campaigns)
	if err != nil || len(campaigns) != count {
		test.Log(test.Name() + " FAILED\n")
	}

	testMap := make(map[int]int)
	testMap[1] = 1
	testMap[2] = 2
	testMap[3] = 3
	testMap[4] = 4
	testMap[5] = 5
	testMap[10] = 10
	testMap[11] = 11
	testMap[12] = 12
	testMap[13] = 13
	testMap[14] = 14
	testMap[15] = 15
	testMap[20] = 20
	testMap[21] = 21
	testMap[22] = 22
	testMap[23] = 23
	testMap[24] = 24
	testMap[25] = 25
	mockResult := "2021/01/06-2021/01/09,2021/01/16-2021/01/19"
	result := getMonthlyGaps(testMap, "2021", "01")
	if strings.Join(result, ",") != mockResult {
		test.Log(test.Name() + " FAILED\n")
	} else {
		test.Log(test.Name() + " PASSED\n")
	}
}

func TestMonthlyGapsMultipleMonths(test *testing.T) {
	campaigns := []Campaign{}
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_1", StartDate: "2021-01-01", EndDate: "2021-01-05"})
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_2", StartDate: "2021-01-10", EndDate: "2021-01-15"})
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_3", StartDate: "2021-01-20", EndDate: "2021-01-25"})

	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_1", StartDate: "2021-02-01", EndDate: "2021-02-05"})
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_2", StartDate: "2021-02-10", EndDate: "2021-02-15"})
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_3", StartDate: "2021-02-20", EndDate: "2021-02-25"})

	_, count, err := addCampaigns(campaigns)
	if err != nil || len(campaigns) != count {
		test.Log(test.Name() + " FAILED\n")
	}

	mockResult := "2021/01/06-2021/01/09,2021/01/16-2021/01/19,2021/02/06-2021/01/09,2021/01/16-2021/02/19"
	result := getAllGaps("2021", "")
	if strings.Join(result, ",") != mockResult {
		test.Log(test.Name() + " FAILED\n")
	} else {
		test.Log(test.Name() + " PASSED\n")
	}
}

func TestMonthlyGapsMultipleMonthsWithChoice(test *testing.T) {
	campaigns := []Campaign{}
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_1", StartDate: "2021-01-01", EndDate: "2021-01-05"})
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_2", StartDate: "2021-01-10", EndDate: "2021-01-15"})
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_3", StartDate: "2021-01-20", EndDate: "2021-01-25"})

	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_1", StartDate: "2021-02-01", EndDate: "2021-02-05"})
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_2", StartDate: "2021-02-10", EndDate: "2021-02-15"})
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_3", StartDate: "2021-02-20", EndDate: "2021-02-25"})

	_, count, err := addCampaigns(campaigns)
	if err != nil || len(campaigns) != count {
		test.Log(test.Name() + " FAILED\n")
	}

	mockResult := "2021/02/06-2021/01/09,2021/01/16-2021/02/19"
	result := getAllGaps("2021", "02")
	if strings.Join(result, ",") != mockResult {
		test.Log(test.Name() + " FAILED\n")
	} else {
		test.Log(test.Name() + " PASSED\n")
	}
}

func TestCategorizedGaps(test *testing.T) {
	campaigns := []Campaign{}
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_1", StartDate: "2021-01-01", EndDate: "2021-01-05"})
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_2", StartDate: "2021-01-10", EndDate: "2021-01-15"})
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_3", StartDate: "2021-01-20", EndDate: "2021-01-25"})

	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_1", StartDate: "2021-02-01", EndDate: "2021-02-05"})
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_2", StartDate: "2021-02-10", EndDate: "2021-02-15"})
	campaigns = append(campaigns, Campaign{CampaignName: "mytestcampaign_3", StartDate: "2021-02-20", EndDate: "2021-02-25"})

	_, count, err := addCampaigns(campaigns)
	if err != nil || len(campaigns) != count {
		test.Log(test.Name() + " FAILED\n")
	}

	result := getCategorizedGaps("2021")
	for _, value := range result {
		test1 := strings.Join(value["01"], ",")
		test2 := "2021/01/06-2021/01/09,2021/01/16-2021/01/19"
		if test1 != test2 {
			test.Log(test.Name() + " FAILED\n")
		}
		test1 = strings.Join(value["02"], ",")
		test2 = "2021/02/06-2021/01/09,2021/01/16-2021/02/19"
		if test1 != test2 {
			test.Log(test.Name() + " FAILED\n")
		}
	}
	test.Log(test.Name() + " PASSED\n")
}
