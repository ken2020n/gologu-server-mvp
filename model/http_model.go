package model

import "time"

type Http struct {
	CreatedAt           time.Time `json:"createdAt"`
	App                 string    `json:"app"`
	Type                int       `json:"type"`
	Url                 string    `json:"url"`
	RequestedAt         time.Time `json:"requestedAt"`
	RequestQueryString  string    `json:"requestQueryString"`
	RequestHeader       string    `json:"requestHeader"`
	RequestContentType  string    `json:"requestContentType"`
	RequestCookies      string    `json:"requestCookies"`
	RequestBody         string    `json:"requestBody"`
	ResponseAt          time.Time `json:"responseAt"`
	ResponseBody        string    `json:"responseBody"`
	ResponseStatusCode  int       `json:"responseStatusCode"`
	ResponseContentType string    `json:"responseContentType"`
	ResponseHeader      string    `json:"responseContentHeader"`
}
