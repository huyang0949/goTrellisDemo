package mongo

import (
	"fmt"
	"net/url"
	"strings"
)

func databaseName(uri string) (string, error) {
	parsed, err := url.Parse(uri)
	if err != nil {
		return "", fmt.Errorf("parse mongo uri: %w", err)
	}

	name := strings.Trim(parsed.Path, "/")
	if name == "" {
		return "", fmt.Errorf("mongo uri must include database name")
	}

	return name, nil
}
