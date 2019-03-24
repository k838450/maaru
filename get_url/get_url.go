/*
  get redirect url using http.Client in Go
  
*/

package get_url

import (
    "errors"
    "fmt"
    "net/http"
    "net/url"
    "time"
	"strings"
)

var RedirectAttemptedError = errors.New("redirect")

func GetUrl(short_url string) string{
    target_url := short_url

    client := &http.Client{
        Timeout: time.Duration(3) * time.Second,
        CheckRedirect: func(req *http.Request, via []*http.Request) error {
            return RedirectAttemptedError
        },
    }
    resp, err := client.Head(target_url)
    defer resp.Body.Close()

	p_url := "aaa"

	if urlError, ok := err.(*url.Error); ok && urlError.Err == RedirectAttemptedError {
        p_url = strings.Join(resp.Header["Location"],"")
		//fmt.Println(resp.Header["Location"])
		fmt.Println(p_url)
        //fmt.Println(strings.Trim(string[url],"[]"))
    }

	//err := exec.Command("youtube-dl",url).Run

	//if err != nil{
		//fmt.Println("Youtube Dl Error")
	//}

	return p_url
}
