package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"


  "github.com/mikkergimenez/distmon/docker"
  "github.com/mikkergimenez/distmon/proc"
)

type StatsJson struct {
	Docker docker.Docker
	Proc proc.Proc
}

type Hosts struct {
	Hosts []StatsJson
}

type Peer struct {
	Hostname string
}

type StatsHTTP struct {
  Peers []Peer
}

func (s StatsHTTP) Handler(w http.ResponseWriter, r *http.Request) {
		var t *template.Template

		t = template.New("DistMon Index Template") // Create a template.
		t, _ = t.ParseFiles("tmpl/index.html", "docker/tmpl/main.html", "proc/tmpl/main.html")  // Parse template file.

		procData := proc.Proc{}
		dockData := docker.Docker{}

		hostStats := StatsJson{
			Proc: procData.Get(),
			Docker: dockData.Get(),
		}

		stats := []StatsJson{hostStats}

		for _, peer := range s.Peers {
			url := fmt.Sprintf("http://%s:55556/json", peer.Hostname)

			fmt.Printf("Getting stats from %s\n" + url)
			res, err := http.Get(url)
			if err != nil {
         panic(err.Error())
		  }

			body, err := ioutil.ReadAll(res.Body)

	 		if err != nil {
			 	panic(err.Error())
	 		}

			var statsJson StatsJson
			json.Unmarshal(body, &statsJson)

			stats = append(stats, statsJson)
		}

		hosts := Hosts{
			Hosts: stats,
		}

		t.ExecuteTemplate(w, "layout", hosts)
}

func (s StatsHTTP) HostHandler(w http.ResponseWriter, r *http.Request) {
		var t *template.Template

		t = template.New("Host Template") // Create a template.
		t, _ = t.ParseFiles("tmpl/index.html", "docker/tmpl/main.html", "proc/tmpl/main.html")  // Parse template file.

		procData := proc.Proc{}
		dockData := docker.Docker{}

		stats := StatsJson{
			Proc: procData.Get(),
			Docker: dockData.Get(),
		}

		t.ExecuteTemplate(w, "layout", stats)
}

func (s StatsHTTP) JSONHandler(w http.ResponseWriter, r *http.Request) {
		procData := proc.Proc{}
		dockData := docker.Docker{}

		stats := StatsJson{
			Proc: procData.Get(),
			Docker: dockData.Get(),
		}

		b, err := json.Marshal(stats)

    if err != nil {
        fmt.Println(err)
        return
    }

		fmt.Fprintf(w, string(b))

}

func main() {
		var peers = []Peer {
			Peer {
				Hostname: "edge0.kbrns.tae.io",
			},
			Peer {
				Hostname: "edge1.kbrns.tae.io",
			},
			Peer {
				Hostname: "edge2.kbrns.tae.io",
			},
		}

		statsHTTP := StatsHTTP{Peers: peers}
		http.HandleFunc("/", statsHTTP.Handler)
		http.HandleFunc("/host", statsHTTP.HostHandler)
		http.HandleFunc("/json", statsHTTP.JSONHandler)

    fmt.Println("Listening on http://localhost:55556")
    http.ListenAndServe(":55556", nil)
}
