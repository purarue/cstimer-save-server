package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

const defaultPort = 8553

var defaultFilename string = "cstimer.json"

type config struct {
	port        int
	saveTo      string
	timestamped bool
	secretKey   string
}

func parseFlags() *config {
	port := flag.Int("port", defaultPort, "port to serve server on")
	saveTo := flag.String("save-to", "", "path to save datafile to.")
	timestamped := flag.Bool("timestamped", false, fmt.Sprintf("instead of writing to the same '%s' file, write to a new file each time", defaultFilename))
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: cstimer-save-server [FLAG...]\nFor instructions, see https://github.com/seanbreckenridge/cstimer-save-server\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if *saveTo == "" {
		log.Fatalf("Error: must provide -save-to, a directory to save files to")
	}
	fileInfo, err := os.Stat(*saveTo)
	if err != nil {
		log.Fatalf("Error: Folder to save files to, '%s' does not exist\n", *saveTo)
	}
	if !fileInfo.IsDir() {
		log.Fatalf("Error: Path '%s' is not a directory", *saveTo)
	}
	secretKey := os.Getenv("CSTIMER_SECRET")
	return &config{
		port:        *port,
		saveTo:      *saveTo,
		timestamped: *timestamped,
		secretKey:   secretKey,
	}
}

func getEpochTime() string {
	return strconv.FormatInt(time.Now().UnixMilli(), 10)
}

func main() {
	config := parseFlags()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method != http.MethodPost {
			http.Error(w, "Not allowed", http.StatusMethodNotAllowed)
			return
		}

		auth := r.URL.Query().Get("auth")
		if config.secretKey != "" && config.secretKey != auth {
			http.Error(w, "Secret key doesnt not match", http.StatusForbidden)
			return
		}
		if auth != "" && config.secretKey == "" {
			fmt.Fprintf(os.Stderr, "Warning: received auth key but no CSTIMER_SECRET set in environment")
		}

		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Could not read body", http.StatusBadRequest)
			return
		}
		if len(b) == 0 {
			http.Error(w, "Body was empty", http.StatusBadRequest)
			return
		}
		var target string
		if config.timestamped {
			target = path.Join(config.saveTo, fmt.Sprintf("%s.json", getEpochTime()))
		} else {
			target = path.Join(config.saveTo, defaultFilename)
		}
		log.Printf("Saving data to '%s'\n", target)
		ioutil.WriteFile(target, b, 0644)
	})
	log.Printf("cstimer-save-server saving to '%s' on port %d\n", config.saveTo, config.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.port), nil))
}
