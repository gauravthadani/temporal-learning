package serviceworkflow

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GreetInSpanish(ctx context.Context, name string) (string, error) {
	base := "http://localhost:9999/get-spanish-greeting?name=%s"
	url := fmt.Sprintf(base, url.QueryEscape(name))

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	translation := string(body)
	status := resp.StatusCode
	if status >= 400 {
		return "", fmt.Errorf("HTTP Error %d: %s", status, translation)
	}

	return translation, nil
}
