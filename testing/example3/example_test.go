package example3_test

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

// feed is mocking the XML document we expect to receive.
var feed = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
<channel>
    <title>Going Go Programming</title>
    <description>Golang : https://github.com/goinggo</description>
    <link>http://www.goinggo.net/</link>
    <item>
        <pubDate>Sun, 15 Mar 2015 15:04:00 +0000</pubDate>
        <title>Object Oriented Programming Mechanics</title>
        <description>Go is an object oriented language.</description>
        <link>http://www.goinggo.net/2015/03/object-oriented</link>
    </item>
</channel>
</rss>`

type Item struct {
	XMLNAME     xml.Name `xml:"item"`
	Title       string   `xml:"title:`
	Description string   `xml:"description`
	Link        string   `xml:"link"`
}

type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
	Items       []Item   `xml:"item"`
}

type Document struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
	URI     string
}

func mockServer() *httptest.Server {

	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, feed)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}

func TestDownload(t *testing.T) {

	statusCode := http.StatusOK

	server := mockServer()
	defer server.Close()

	t.Logf("Given the need to test downloading content.")
	{
		t.Logf("\tTest : 0\tWhen Checking %q for status code %d", server.URL, statusCode)
		{

			resp, err := http.Get(server.URL)
			if err != nil {
				t.Fatalf("\t%s\tShould be able to make the GET call : %v", failed, err)
			}

			t.Logf("\t%s\tShould be able to make the Get call.", succeed)

			defer resp.Body.Close()

			if resp.StatusCode != statusCode {
				t.Fatalf("\t%s\tShould receive a %d status code : %v", failed, statusCode, resp.StatusCode)
			}
			t.Logf("\t%s\tShould receive a %d status code", succeed, statusCode)

			var d Document
			if err := xml.NewDecoder(resp.Body).Decode(&d); err != nil {
				t.Fatalf("\t%s\tShould be able to unmarshal the response : %v", failed, err)
			}
			t.Logf("\t%s\tShould be able to unmarshal the response.", succeed)

			if len(d.Channel.Items) == 1 {
				t.Logf("\t%s\tShould have 1 item in the feed.", succeed)
			} else {
				t.Errorf("\t%s\tShould have 1 item in the feed : %d", failed, len(d.Channel.Items))
			}
		}
	}
}
