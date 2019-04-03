package internal

import (
	"bytes"
	"encoding/json"
	"github.com/lestrrat-go/file-rotatelogs"
	"net/http"
	"path/filepath"
	"pgl/pkg"
	"time"
)

const CONNECTOR = "connector-log"
const APPLICATION = "application-log"
const ROOT = "logs"

func HandleError(w http.ResponseWriter, r *http.Request) {
	var errEntry pkg.ErrorReport
	GetPostData(r.Body, &errEntry, w)

	abs, _ := filepath.Abs(".")
	path := filepath.Join(abs, APPLICATION)

	writer, err := getLogWriter(path)
	if err != nil {
		http.Error(w, "An error occurred", 500)
		return
	}
	defer writer.Close()

	logEntry := pkg.LogEntry{
		Timestamp: time.Now().Format(time.RFC850),
		Entry:     errEntry,
	}

	var buffer bytes.Buffer
	b, err := json.Marshal(logEntry)
	if err != nil {
		http.Error(w, "An error occurred", 500)
		return
	}
	buffer.Write(b)
	buffer.WriteString("\n")

	writer.Write(buffer.Bytes())
}

func HandleTraffic(w http.ResponseWriter, r *http.Request) {

}

func getLogWriter(path string) (w *rotatelogs.RotateLogs, err error) {
	w, err = rotatelogs.New(
		path+".%Y%m%d%H%M",
		rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(604800)*time.Second),
	)

	return
}
