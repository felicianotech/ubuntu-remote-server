package api

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

/*
 * Display / GET
 */
func (this *App) index(w http.ResponseWriter, r *http.Request) {

	output, err := exec.Command("amixer", "-D", "pulse", "sget", "Master", "|", "grep", "-m", "1", "-o", "-E", "[[:digit:]]+%").Output()
	if err != nil {
		log.Fatal(err)
	}

	log.Info(output)
}

/*
 * Return volume and mute status /volume GET
 */
func (a *App) volumeGet(w http.ResponseWriter, r *http.Request) {

	command := "amixer -D pulse sget Master"
	volOutput, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		log.Fatalf("Get volume failed: %+v", err)
	}

	var vol string
	var muted bool
	lines := strings.Split(string(volOutput), "\n")
	for _, line := range lines {

		if strings.Contains(line, "Front Left:") {
			start := strings.Index(line, "[") + 1
			end := strings.Index(line, "]")

			vol = line[start:end]

			start = strings.LastIndex(line, "[") + 1
			end = strings.LastIndex(line, "]")

			muteOutput := line[start:end]
			if muteOutput == "off" {
				muted = true
			} else if muteOutput == "on" {
				muted = false
			} else {
				log.Fatalf("Muted parsing failed: %+v", err)
			}

			break
		}
	}

	respondWithJSON(w, 200, &Volume{Value: vol, Muted: muted})
}

/*
 * Set volume and mute status /volume POST
 */
func (a *App) volumePost(w http.ResponseWriter, r *http.Request) {

	var vol Volume

	err := json.NewDecoder(r.Body).Decode(&vol)
	if err != nil {

		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	command := "amixer -D pulse sset Master " + vol.Value
	_, err = exec.Command("bash", "-c", command).Output()
	if err != nil {
		log.Fatalf("Set volume failed: %+v", err)
	}

	var mute string
	if vol.Muted {
		mute = "off"
	} else {
		mute = "on"
	}

	command = "amixer -D pulse sset Master " + mute
	_, err = exec.Command("bash", "-c", command).Output()
	if err != nil {
		log.Fatalf("Set volume failed: %+v", err)
	}
}
