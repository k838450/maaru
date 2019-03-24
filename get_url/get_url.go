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
		fmt.Println(p_url)
    }

	return p_url
}
