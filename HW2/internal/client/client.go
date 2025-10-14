package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"hw2/internal/decstr"
	"io"
	"net/http"
	"time"
)

func StartClient(url string, input string) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		IdleConnTimeout:       30 * time.Second,
		TLSHandshakeTimeout:   30 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := &http.Client{
		Timeout:   15 * time.Second,
		Transport: tr,
	}

	resp, err := client.Get(url + "/version")
	if err != nil {
		fmt.Println("Error in GET /version request: ", err)
	} else {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	}

	code := base64.StdEncoding.EncodeToString([]byte(input))
	request := decstr.DecodeRequest{
		InputString: code,
	}
	str, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Json convertation failed: ", err)
	}

	resp, err = client.Post(url+"/decode", "application/json", bytes.NewBuffer(str))

	if err != nil {
		fmt.Println("Error in Post /decode request: ", err)
	} else {
		defer resp.Body.Close()
		var ans decstr.DecodeResponse
		err := json.NewDecoder(resp.Body).Decode(&ans)
		if err != nil {
			fmt.Println("Json decoding failed: ", err)
		}
		fmt.Println(ans.OutputString)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url+"/hard-op", nil)
	resp, err = client.Do(req)
	if errors.Is(err, context.DeadlineExceeded) {
		fmt.Println("false")
	} else if err != nil {
		fmt.Println("Error in GET /hard-op request: ", err)
	} else {
		defer resp.Body.Close()
		fmt.Println("true,", resp.StatusCode)
	}

}
