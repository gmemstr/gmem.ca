const fetchSteamStatus = () => {
    fetch("/api/steam")
      .then((response) => {return response.json()})
      .then((data) => {
          let experiment = new Experiment("steam")
          experiment.html`SteamID: ${data.SteamID} Online State: ${data.OnlineState}`
       })
}

const loadApp = () => {
    fetchSteamStatus()
}
window.onload = loadApp

class Experiment {
    constructor(id) {
        this.id = id
    }

    html(strings, ...things) {
        const element = document.getElementById(this.id)
        let x = document.createRange().createContextualFragment(
          strings.reduce(
            (markup, string, index) => {
              markup += string
      
              if (things[index]) {
                markup += things[index]
              }
      
              return markup
            },
            ''
          )
        )
        element.innerHTML = ""
        element.append(x)
    }
}
