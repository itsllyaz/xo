package stream

import (
	"fmt"
	"net/http"

	"xo/internal"
)



func SampleStream() (*http.Response, error) {
	url := fmt.Sprintf("%s/tweets/sample/stream", internal.BaseURL)
	httpRequest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	internal.BearerOAUTH(httpRequest)

	client := http.Client{}
	res, err := client.Do(httpRequest)

	if err != nil {
		return nil, err
	}
	return res, nil
}
