package elevator_media

import (
	"net/http"
	"io/ioutil"
	"log"
)

type Streamer struct {
	game string
	ada string
}

func newStreamer() Streamer {
	s := Streamer{}
	s.game = "<iframe width=100% height='100%' src=https://www.addictinggames.com/embed/html5-games/23635 scrolling=no></iframe>"
	s.ada = "<div class=nomics-ticker-widget data-name=Cardano data-base=ADA data-quote=USD></div><script src=https://widget.nomics.com/embed.js></script>"
	return s
}

func (s *Streamer) getContent(data string) string {
	if data == "game" {
		return s.format(s.game)
	} else if data == "ada" {
		return s.format(s.ada)
	} else if data == "quote" {
		return s.format(s.quote())
	}
	return ""
}

func (s *Streamer) format(data string) string {
	return "<div> " + data + " </div>"
}

func (s *Streamer) quote() string {
	resp, err := http.Get("http://getmeaquote.designedbyaturtle.com/")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}