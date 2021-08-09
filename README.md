# pixability

## Campaign Scheduler Exercise

This repo consists of the code designed/written for pixability exercise

## Code Organization

The code is written in pure Golang and can be started from campaign_scheduler.go under the folder pixability. You should be able to clone this folder as is under src folder of your GOPATH. Then you can either build it scripts or directly use command
  
  go build
  
Code can be viewed as:

      pixability
          |-- api
                |-- campaign_handler.go
                |-- campaign_helper.go
                |-- campaign_models.go
                |-- api_test.go
          |-- campaign_scheduler.go
          |-- build
          |-- Readme.md

## Testing the Code

To do this you've two options:

  1. Run the test cases under api package/folder. This can serve you quick debugging and learning how code works.
  2. Build the code and run locally, then test it as an API in Postman. Alternatively, you can also use CURL requests to test once API server is up.

## Implementation Specifics

There are 5 APIs implemented under campaign_scheduler.go

  1. Default API for route "/" : This API is just to onboard and tell how to use the basic API as a functionality requested in the question.
  2. /pixability/scheduleCampaigns : This API route adds the campaigns. The request is a POST request with a list of campaigns needed to be added:
    ``` JSON
    {
      "campaigns" : [
        {
          "campaign_name":"",
          "start_date":"",
          "end_date":""
        },
        {
          "campaign_name":"",
          "start_date":"",
          "end_date":""
        },
        {
          "campaign_name":"",
          "start_date":"",
          "end_date":""
        }
      ]
    }
    ```

  3. /pixability/getGaps : This API route fetch the gaps in the campaigns. This API is a GET request with optional parameter of year and month. If no year is specified it will fetch gaps for current year and all months. If a month is specified then it will fetch gaps for that particular month.
  4. /pixability/getCategorizedGaps : This API route fetch the gaps in the campaigns arranged categorically. This API is a GET request with an optional parameter of year. If no year is specified it will fetch gaps for current year. It'll give data for all months categorized by year and month.
  5. /debug/getCampaigns : This API route is a debug endpoint just to see current campaigns running under the server. It has no parameters and it gives full wholesome view of the running campaigns.
  
## Scripts

The clean script is used with build script to remove any temporary files generated. The build script is used to build binary for this server for MacOS. You can build it for linux as well by commenting the current build line and uncommenting the corrresponding line in build script.

## Assumptions

  1. The gaps are only calculated until the last day of the last campaign in a month.
  2. Campaigns can be added in any order.
  
All other information should be self explanatory through the code.
