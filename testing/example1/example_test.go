package example

import (
	"net/http"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestDownload(t *testing.T) {

	url := "https://www.goinggo.net/post/index.xml"
	statusCode := 200

	t.Log("given need to download content")
	{
		t.Logf("\tTest 0\twhen checking %q for status code %d", url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatalf("\t%s\tShould be able to make the call : %v", failed, err)
			}

			t.Logf("\t%s\tShould be able to make the call", succeed)

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t%s\tShould recieve a %d status code", succeed, statusCode)
			} else {
				t.Errorf("\t%s\tShould recieve a %d status code : %d", failed, statusCode, resp.StatusCode)
			}
		}
	}
}
