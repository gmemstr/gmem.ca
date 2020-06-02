// Proxy's Steam Community XML data into basic JSON format.
package handler

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type SteamProfile struct {
	SteamID     string       `xml:"steamID"`
	OnlineState string       `xml:"onlineState"`
	InGameInfo  []InGameInfo `xml:"inGameInfo"`
}

type InGameInfo struct {
	Name string `xml:"gameName"`
	Link string `xml:"gameLink"`
	Icon string `xml:"gameIcon"`
	Logo string `xml:"gameLogo"`
}

func getSteamData() SteamProfile {
	resp, err := http.Get("https://steamcommunity.com/id/gmemstr?xml=1")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var steamData SteamProfile
	err = xml.Unmarshal([]byte(body), &steamData)
	if err != nil {
		// handle error
	}

	return steamData
}

// Handler handles incoming requests and dispatches to fetch data.
func Handler(w http.ResponseWriter, r *http.Request) {
	data := getSteamData()
	b, _ := json.Marshal(data)
	w.Write(b)
}
