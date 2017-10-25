package example2_test

import (
	"net/http"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestDownload(t *testing.T) {

	tests := []struct {
		url        string
		statusCode int
	}{
		{"https://www.goinggo.net/post/index.xml", http.StatusOK},
		{"http://rss.cnn.com/rss/cnn_topstorie.rss", http.StatusNotFound},
	}

	t.Log("Given the need to test downloading different content.")
	{
		for i, tt := range tests {

			t.Logf("\tTest: %d\tWhen checking %q for status code %d", i, tt.url, tt.statusCode)
			{
				resp, err := http.Get(tt.url)
				if err != nil {
					t.Fatalf("\t%s\tShould be able to make the GET call : %v", failed, err)
				}
				t.Logf("\t%s\tShould be able to make the GET call.", succeed)

				defer resp.Body.Close()

				if resp.StatusCode == tt.statusCode {
					t.Logf("\t%s\tShould recieved a %d status code.", succeed, tt.statusCode)
				} else {
					t.Errorf("\t%s\tShould recieved a %d status code : %v", failed, tt.statusCode, resp.StatusCode)
				}
			}
		}
	}
}
