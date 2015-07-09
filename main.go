package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"flag"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
)

func init() {
	var s int64
	if err := binary.Read(crand.Reader, binary.BigEndian, &s); err != nil {
		log.Fatal(err)
	}
	rand.Seed(s)
}

var indexTemplate *template.Template

func main() {
	listen := flag.String("listen", ":8090", "host:port to listen on")
	assets := flag.String("assets", "assets", "path to assets directory")
	templates := flag.String("templates", "templates", "path to template directory")
	flag.Parse()

	var err error
	indexTemplate, err = template.ParseFiles(*templates + "/index.html")
	if err != nil {
		log.Fatalf("unable to parse index template: %s", err)
	}

	http.Handle("/assets/", http.StripPrefix("/assets/",
		http.FileServer(http.Dir(*assets))))
	http.HandleFunc("/", home)

	log.Fatal(http.ListenAndServe(*listen, nil))
}

func pick(segment []string) string {
	return segment[rand.Intn(len(segment))]
}

func insult() string {
	return fmt.Sprintf("Thou %s %s %s!", pick(columnA), pick(columnB), pick(columnC))
}

func home(w http.ResponseWriter, r *http.Request) {
	s := insult()
	indexTemplate.Execute(w, s)
}

// Word lists from: http://web.mit.edu/dryfoo/Funny-pages/shakespeare-insult-kit.html
var columnA = []string{
	"artless", "bawdy", "beslubbering", "bootless", "churlish",
	"cockered", "clouted", "craven", "currish", "dankish", "dissembling", "droning",
	"errant", "fawning", "fobbing", "froward", "frothy", "gleeking", "goatish",
	"gorbellied", "impertinent", "infectious", "jarring", "loggerheaded", "lumpish",
	"mammering", "mangled", "mewling", "paunchy", "pribbling", "puking", "puny",
	"qualling", "rank", "reeky", "roguish", "ruttish", "saucy", "spleeny", "spongy",
	"surly", "tottering", "unmuzzled", "vain", "venomed", "villainous", "warped",
	"wayward", "weedy", "yeasty",
}

var columnB = []string{
	"base-court", "bat-fowling", "beef-witted", "beetle-headed", "boil-brained",
	"clapper-clawed", "clay-brained", "common-kissing", "crook-pated", "dismal-dreaming",
	"dizzy-eyed", "doghearted", "dread-bolted", "earth-vexing", "elf-skinned",
	"fat-kidneyed", "fen-sucked", "flap-mouthed", "fly-bitten", "folly-bitten",
	"fool-born", "full-gorged", "guts-griping", "half-faced", "hasty-witted",
	"hedge-born", "hell-hated", "idle-headed", "ill-breeding", "ill-nurtured",
	"knotty-pated", "milk-livered", "motley-minded", "onion-eyed", "plume-plucked",
	"pottle-deep", "pox-marked", "reeling-ripe", "rough-hewn", "rude-growing", "rump-fed",
	"shard-borne", "sheep-biting", "spur-galled", "swag-bellied", "tardy-gaited",
	"tickle-brained", "toad-spotted", "unchin-snouted", "weather-bitten",
}

var columnC = []string{
	"apple-john", "baggage", "barnacle", "bladder", "boar-pig", "bugbear", "bum-bailey",
	"canker-blossom", "clack-dish", "clotpole", "coxcomb", "codpiece", "death-token",
	"dewberry", "flap-dragon", "flax-wench", "flirt-gill", "foot-licker", "fustilarian",
	"giglet", "gudgeon", "haggard", "harpy", "hedge-pig", "horn-best", "hugger-mugger",
	"joithead", "lewdster", "lout", "maggot-pie", "malt-worm", "mammet", "measle", "minnow",
	"miscreant", "moldwarp", "mumble-news", "nut-hook", "pidgeon-egg", "pignut", "puttock", "pumpion",
	"ratsbane", "scut", "skainsmate", "strumpet", "varlot", "vassal", "whey-faced", "wagtail",
}
