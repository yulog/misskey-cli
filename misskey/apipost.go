package misskey

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func (c *Client) apiPost(jsonByte []byte, endpoint string) error {

	endpoint, err := url.JoinPath(c.InstanceInfo.Host, "api", endpoint)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"POST",
		endpoint,
		bytes.NewBuffer(jsonByte),
	)
	if err != nil {
		return err
	}

	req.Header.Set("content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	c.resBuf = new(bytes.Buffer)
	if _, err = io.Copy(c.resBuf, resp.Body); err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 300 {
		fmt.Fprintln(os.Stderr, resp.StatusCode, c.resBuf)
		os.Exit(1)
	}
	defer resp.Body.Close()
	return nil
}
