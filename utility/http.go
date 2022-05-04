package utility

import (
	"bytes"
	"net/http"
)

// HttpRequest ...
// making the http request upload or get data from external vendor
func HttpRequest(method string, url string, body []byte, header map[string]string, queryString map[string]string) (*http.Response, error) {
	var req *http.Request
	var res *http.Response
	var err error
	for {
		// creating request for uploading data
		if req, err = http.NewRequest(method, url, bytes.NewBuffer(body)); err != nil {
			break
		}
		if len(header) > 0 {
			// header for request
			for key, value := range header {
				req.Header.Set(key, value)
			}
		}
		// setting query string params
		if len(queryString) > 0 {
			queryParam := req.URL.Query()
			for key, value := range queryString {
				queryParam.Add(key, value)
			}
			req.URL.RawQuery = queryParam.Encode()
		}
		// response from vendor
		res, err = http.DefaultClient.Do(req)
		break
	}
	return res, err
}
