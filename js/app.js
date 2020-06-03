const fetchSteamStatus = async () => {
    let response = await fetch("/api/steam")
    let data = await response.json()
    let experiment = new Experiment("steam")

    experiment.html`
    <h2>Steam Status</h2>
    <p>SteamID: ${data.steamId} Online State: ${data.onlineState}</p>
    ${data.onlineState === 'in-game' && data.inGameInfo !== null ? `<img src=${data.inGameInfo.logo}>${data.inGameInfo.name}` : ''}
    `
}

const loadApp = async () => {
    await fetchSteamStatus()
}
window.onload = loadApp

class Experiment {
    constructor(id) {
        this.id = id
        this.element = document.getElementById(this.id)
    }

    html(strings, ...things) {
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
        this.content = x
        this.element.innerHTML = ""
        this.element.append(x)
    }
}
