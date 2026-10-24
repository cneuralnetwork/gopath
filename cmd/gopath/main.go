package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"gopath"
	"gopath/internal/content"
	"gopath/internal/runner"
	"gopath/internal/server"
	"gopath/internal/store"
)

func main() {
	port := flag.Int("port", 7777, "local HTTP port")
	dir := flag.String("data-dir", "", "progress directory")
	noBrowser := flag.Bool("no-browser", false, "do not open a browser")
	flag.Parse()
	if *dir == "" {
		home, e := os.UserHomeDir()
		if e != nil {
			log.Fatal(e)
		}
		*dir = filepath.Join(home, ".gopath-learn")
	}
	contentFS, e := fs.Sub(gopath.FS(), "content")
	if e != nil {
		log.Fatal(e)
	}
	catalog, e := content.Load(contentFS)
	if e != nil {
		log.Fatal(e)
	}
	progress, e := store.New(*dir)
	if e != nil {
		log.Fatal(e)
	}
	webFS, e := fs.Sub(gopath.FS(), "web/dist")
	if e != nil {
		log.Fatal(e)
	}
	goVersion, moduleVersion, _ := runner.Version()
	if moduleVersion == "" {
		moduleVersion = "1.22"
	}
	cache := filepath.Join(*dir, "gocache")
	_ = os.MkdirAll(cache, 0700)
	s := &server.Server{Catalog: catalog, Store: progress, Runner: runner.New(cache), Web: webFS, Version: "0.1.0", GoVersion: goVersion, GoModuleVersion: moduleVersion}
	address := fmt.Sprintf("127.0.0.1:%d", *port)
	url := "http://" + address
	log.Printf("Gopath at %s (Go: %s)", url, goVersion)
	if !*noBrowser {
		go func() {
			time.Sleep(300 * time.Millisecond)
			switch runtime.GOOS {
			case "darwin":
				_ = exec.Command("open", url).Start()
			case "windows":
				_ = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
			default:
				_ = exec.Command("xdg-open", url).Start()
			}
		}()
	}
	log.Fatal(http.ListenAndServe(address, s.Handler()))
}
