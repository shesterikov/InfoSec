package demo

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"
	"strings"
)

func parseSafe(data []byte) (string, error) {
	decoder := xml.NewDecoder(bytes.NewReader(data))
	// Token inspection rejects DTD explicitly; encoding/xml performs no external fetches.
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		if directive, ok := token.(xml.Directive); ok && strings.HasPrefix(strings.ToUpper(strings.TrimSpace(string(directive))), "DOCTYPE") {
			return "", errors.New("DTD forbidden")
		}
	}
	var document struct {
		XMLName xml.Name `xml:"message"`
		Value   string   `xml:"value"`
	}
	if err := xml.Unmarshal(data, &document); err != nil {
		return "", err
	}
	return document.Value, nil
}
