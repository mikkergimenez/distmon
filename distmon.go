package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
  "github.com/mikkergimenez/distmon/docker"
  "github.com/mikkergimenez/distmon/proc"
)

type StatsJson struct {
	Hostname string
	Docker docker.Docker
	Proc proc.Proc
}

type Hosts struct {
	Hosts []StatsJson
}

type Peers []string

type StatsHTTP struct {
  Peers Peers
	Hostname string
}

type Config struct {
	Hostname string
	Peers 	 []string
}

func (s StatsHTTP) Handler(w http.ResponseWriter, r *http.Request) {
		var t *template.Template

		t = template.New("DistMon Index Template") // Create a template.
		t, _ = t.ParseFiles("tmpl/index.html")  // Parse template file.

		stats := []StatsJson{}

		for i, peer := range s.Peers {
			url := fmt.Sprintf("http://%s:55556/json", peer)

			fmt.Printf("%d) Getting stats from %s\n", i, url)

			res, err := http.Get(url)

			if err != nil {
	       log.Println(err.Error())
				 continue
		  }

			body, err := ioutil.ReadAll(res.Body)
	 		if err != nil {
			 	fmt.Println(err.Error())
				continue
	 		}

			var statsJson StatsJson
			json.Unmarshal(body, &statsJson)

			stats = append(stats, statsJson)
		}

		hosts := Hosts{
			Hosts: stats,
		}

		if (os.Getenv("DISTMON_DEBUG") == "TRUE") {
			fmt.Printf("%+v", hosts)
		}

		t.ExecuteTemplate(w, "layout", hosts)
}

func (s StatsHTTP) HostHandler(w http.ResponseWriter, r *http.Request) {
		var t *template.Template

		t = template.New("Host Template") // Create a template.

		t, _ = t.ParseFiles("tmpl/host.html", "tmpl/docker.html", "tmpl/proc.html"	)  // Parse template file.

		procData := proc.Proc{}
		dockData := docker.Docker{}

		stats := StatsJson{
			Proc: procData.Get(),
			Docker: dockData.Get(),
		}

		t.ExecuteTemplate(w, "layout", stats)
}

func getHostname(confHostname string) string {
		value := os.Getenv("DISTMON_FQDN")
    if (len(value) == 0) {
				if (os.Getenv("DISTMON_DEBUG") == "TRUE") {
					fmt.Println("Getting hostname from config.")
				}
        return confHostname
    }
		if (os.Getenv("DISTMON_DEBUG") == "TRUE") {
			fmt.Println("Getting hostname from environment variable.")
		}
    return value
}

func (s StatsHTTP) JSONHandler(w http.ResponseWriter, r *http.Request) {
		procData := proc.Proc{}
		dockData := docker.Docker{}

		hostname := getHostname(s.Hostname)

		if (len(hostname) == 0) {
			panic("No Hostname Set!  Check out readme at https://github.com/mikkergimenez/distmon/ for more information.")
		}

		stats := StatsJson{
			Hostname: hostname,
			Proc: procData.Get(),
			Docker: dockData.Get(),
		}

		b, err := json.Marshal(stats)

    if err != nil {
				fmt.Println("Error Marshaling Stats: %s", err)
        return
    }

		fmt.Fprintf(w, string(b))

}

func main() {
		var conf Config
		if _, err := toml.DecodeFile("/etc/distmon.toml", &conf); err != nil {
			fmt.Println(err)
			panic("Can't Read Config File at /etc/distmon.toml")
		}
		peers := conf.Peers

		statsHTTP := StatsHTTP{Peers: peers, Hostname: conf.Hostname}
		http.HandleFunc("/", statsHTTP.Handler)
		http.HandleFunc("/host", statsHTTP.HostHandler)
		http.HandleFunc("/json", statsHTTP.JSONHandler)

    fmt.Println("Listening on http://localhost:55556")
    http.ListenAndServe(":55556", nil)
}
