package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.Info("Started web service...")

	http.HandleFunc("/", Homepage)
	http.HandleFunc("/vote", Vote)
	http.ListenAndServe(":8090", nil)
}

func Vote(w http.ResponseWriter, req *http.Request) {
	log.Infof("Received POST for Vote")

	err := req.ParseForm()
	if err != nil {
		log.Error("Cannot parse form!")
	}

	vote := req.PostFormValue("contact")

	log.Infof("Vote is %v", vote)

	ProcessVote(vote)

	fmt.Fprintf(w, "Thank you for voting!")
}

func ProcessVote(vote string) int {
	log.Warnf("Into ProcessVote with param = %v", vote)

	var voteInt int
	switch vote {
	case "1":
		log.Debug("User didn't like it")
		voteInt = 100
	case "2":
		voteInt = 200
	case "3":
		voteInt = 300
	case "4":
		voteInt = 400
	case "5":
		voteInt = 500
	default:
		log.Errorf("I don't know what vote %s is!!!", vote)
		voteInt = 0
	}
	return voteInt
}

func Homepage(w http.ResponseWriter, req *http.Request) {
	log.Warnf("Someone connected: %v", req)
	fmt.Fprintf(w, `<h1>Voting machine</h1>
    How do you like this seminar?
    
    <form action="/vote" method="post">
      <p>Please vote with a number:</p>
      <div>
        <input type="radio" id="vote1" name="contact" value="1">
        <label for="contactChoice1">1</label>
    
        <input type="radio" id="vote2" name="contact" value="2">
        <label for="contactChoice1">2</label>
    
        <input type="radio" id="vote3" name="contact" value="3">
        <label for="contactChoice1">3</label>
    
        <input type="radio" id="vote4" name="contact" value="4">
        <label for="contactChoice1">4</label>
    
        <input type="radio" id="vote5" name="contact" value="5">
        <label for="contactChoice1">5</label>
      </div>
      <div>
        <button type="submit">Submit</button>
      </div>
    </form>
`)

}
