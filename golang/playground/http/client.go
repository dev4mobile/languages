package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	request, err := http.NewRequest(http.MethodGet, "https://www.imooc.com", nil);
	request.Header.Add("user-agent","Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1)");
	//resp, err := http.DefaultClient.Do(request)
	//resp, err := http.Get("http://dev4mobiles.com")
	client := http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
		fmt.Println("Redirect:", req)
		return nil
	}}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", bytes)
}
