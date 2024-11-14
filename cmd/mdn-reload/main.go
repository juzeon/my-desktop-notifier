package main

import (
	"github.com/samber/lo"
	"io"
	"log/slog"
	"net/http"
)

func main() {
	resp, err := http.Post("http://127.0.0.1:7888/", "", nil)
	if err != nil {
		slog.Error("Error", "err", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		slog.Error("Status", "v", resp.StatusCode)
		slog.Error("Body", "v", string(lo.Must(io.ReadAll(resp.Body))))
		return
	}
	slog.Info("OK")
}
