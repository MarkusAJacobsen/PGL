package internal

import (
	"bufio"
	"encoding/json"
	"github.com/MarkusAJacobsen/pgl/pkg"
	"github.com/sirupsen/logrus"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type report struct {
	PageTitle string
	Entries   []pkg.LogEntry
}

const ReportTemplate = "./web/template/report.html"

func HandleReport(w http.ResponseWriter, _ *http.Request) {
	e := getLogFiles()

	abs, err := filepath.Abs(ReportTemplate)
	if err != nil {
		http.Error(w, "Internal error", 500)
		logrus.Errorln(err)
		return
	}

	t := template.Must(template.ParseFiles(abs))
	data := &report{
		PageTitle: "Report",
		Entries:   e,
	}

	t.Execute(w, data)
}

func getLogFiles(r ...string) (e []pkg.LogEntry) {
	var root string
	if len(r) > 0 {
		root = r[0]
	} else {
		root = getLogPath()
	}

	files, err := ioutil.ReadDir(root)
	if err != nil {
		logrus.Errorln(err)
		return
	}

	var entries []pkg.LogEntry
	for _, f := range files {
		if f.IsDir() {
			abs, err := filepath.Abs(root + string(filepath.Separator) + f.Name())
			if err != nil {
				logrus.Errorln(err)
				return
			}
			entries = append(entries, getLogFiles(abs)...)
		} else {
			abs, err := filepath.Abs(root + string(filepath.Separator) + f.Name())
			if err != nil {
				logrus.Errorln(err)
				return
			}
			entries = append(entries, getData(abs)...)
		}
	}
	return entries
}

func getData(p string) (entries []pkg.LogEntry) {
	f, err := os.Open(p)
	if err != nil {
		logrus.Errorln("Could not open file:", p)
		return entries
	}
	defer f.Close()

	// https://stackoverflow.com/a/41741702
	reader := bufio.NewReader(f)
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}

		var entry pkg.LogEntry
		err = json.Unmarshal([]byte(line), &entry)
		if err != nil {
			break
		}

		entries = append(entries, entry)
	}

	if err != io.EOF {
		logrus.Errorln(err)
	}

	return entries
}
