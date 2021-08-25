package real

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	Timeout   time.Duration
}

func (r *Retriever) Post(url string, form map[string]string) string {
	fmt.Println(url, form)
	return ""
}

func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	return string(result)
}
