package internal

import (
	"bytes"
	"encoding/json"
	"github.com/MarkusAJacobsen/pgl/pkg"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const Connector = "connector-log"
const Application = "application-log"

func HandleError(w http.ResponseWriter, r *http.Request) {
	var errEntry pkg.ErrorReport
	GetPostData(r.Body, &errEntry, w)

	writeToLog(Application, errEntry, w)
}

func HandleTraffic(w http.ResponseWriter, r *http.Request) {
	var tEntry pkg.TrafficReport
	GetPostData(r.Body, &tEntry, w)

	writeToLog(Connector, tEntry, w)
}

func writeToLog(logType string, entry interface{}, w http.ResponseWriter) {
	root := getLogPath()

	abs, err := filepath.Abs(root)
	if err != nil {
		http.Error(w, "An error occurred", 500)
		return
	}

	os.Mkdir(root, os.ModePerm)
	path := filepath.Join(abs, logType)

	writer, err := getLogWriter(path)
	if err != nil {
		http.Error(w, "An error occurred", 500)
		return
	}
	defer writer.Close()

	logEntry := pkg.LogEntry{
		Timestamp: time.Now().Format(time.RFC850),
		Entry:     entry,
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
	logrus.Infoln(string(b))
}

func getLogWriter(path string) (w *rotatelogs.RotateLogs, err error) {
	w, err = rotatelogs.New(
		path+".%Y%m%d%H%M",
		rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(604800)*time.Second),
	)

	return
}
