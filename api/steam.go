// Package handler proxy's Steam Community XML data into basic JSON format.
package handler

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// SteamProfile contains very basic information about
// a Steam community profile, including their current "ID",
// online state, and in-game info, if available.
type SteamProfile struct {
	SteamID     string       `xml:"steamID" json:"steamId"`
	OnlineState string       `xml:"onlineState" json:"onlineState"`
	InGameInfo  []InGameInfo `xml:"inGameInfo" json:"inGameInfo"`
}

// InGameInfo contains very basic data about a game currently
// being played.
type InGameInfo struct {
	Name string `xml:"gameName" json:"name"`
	Link string `xml:"gameLink" json:"link"`
	Icon string `xml:"gameIcon" json:"icon"`
	Logo string `xml:"gameLogo" json:"logo"`
}

func getSteamData() SteamProfile {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://steamcommunity.com/id/gmemstr?xml=1", nil)
	if err != nil {
		os.Stdout.Write([]byte(err.Error()))
		return SteamProfile{}
	}

	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Pragma", "no-cache")
	fmt.Println(req.Header)

	resp, err := client.Do(req)
	if err != nil {
		os.Stdout.Write([]byte(err.Error()))
		return SteamProfile{}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var steamData SteamProfile
	err = xml.Unmarshal([]byte(body), &steamData)
	if err != nil {
		os.Stdout.Write([]byte(err.Error()))
		return SteamProfile{}
	}

	return steamData
}

// Handler handles incoming requests and dispatches to fetch data.
func Handler(w http.ResponseWriter, r *http.Request) {
	data := getSteamData()
	b, err := json.Marshal(data)
	if err != nil {
		os.Stdout.Write([]byte(err.Error()))
		w.Write([]byte("{}"))
		return
	}
	w.Header().Set("Cache-Control", "no-cache")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
