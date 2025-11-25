package lsp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

func send(msg any) {
	data, _ := json.Marshal(msg)
	fmt.Printf("Content-Length: %d\r\n\r\n%s", len(data), data)
}

func readMessage(r *bufio.Reader) ([]byte, error) {
	headers := map[string]string{}
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			return nil, err
		}

		line = strings.TrimSpace(line)
		if line == "" {
			break
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}

	length := 0
	fmt.Sscanf(headers["Content-Length"], "%d", &length)

	body := make([]byte, length)
	_, err := io.ReadFull(r, body)
	return body, err
}
