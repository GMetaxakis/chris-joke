package main

import (
    "fmt"
    "log"
    "net/http"
)
var Jokes []string

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/chris-joke", returnAllChrisJokes)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func populateJokes() {
    Jokes = []{
        "“Kostas will join in a bit” -“if he joins in 8 bits, it’ll be a byte late”",
        "- I love the smell of conflicts in the morning :seferlis:",
        "- I’m such a good presenter",
        "- There’s an issue if there are no issues",
        "- Goodies and Mcdonaldsss",
        "- I guess they should rename their company to offFido",
        "- Chryssa: Are you the prime minister? Chris: No, I am not. I am prime, though",
        "- I’m here to ask the hard questions",
        "- It is easy, it is :electric_plug: and :pray:",
        "- Only peasants can’t write SQL!!!!!!!!! SQL is for the lords (arxontous) not for cheap cloths (linatses)",
    }
}

func returnAllChrisJokes(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllChrisJokes")
    json.NewEncoder(w).Encode(Jokes)
}

func main() {
    populateJokes()
    handleRequests()
}