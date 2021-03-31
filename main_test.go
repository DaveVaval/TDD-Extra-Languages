package elevator_media

import (
	"testing"
	"regexp"
)

func TestGetContent(t *testing.T) {
	t.Run("greets GoCon", func(t *testing.T){
		got := Hello("GoCon")
		expected := "Hello GoCon!"
		
		if got != expected{
			t.Errorf("Got: %s, Expected: %s", got, expected)
		}
	})
	
	streamer := newStreamer()
	
	t.Run("game Test", func(t *testing.T){
		got := streamer.getContent("game")
		expected := format(streamer.game)

		if got != expected{
			t.Errorf("Got: %s, Expected: %s", got, expected)
		}
	})

	t.Run("ada Test", func(t *testing.T){
		got := streamer.getContent("ada")
		expected := format(streamer.ada)

		if got != expected{
			t.Errorf("Got: %s, Expected: %s", got, expected)
		}
	})

	t.Run("quote Test", func(t *testing.T){
		got := streamer.getContent("quote")
		expected, _:= regexp.MatchString("<div>", got)

		if expected != true {
			t.Errorf("Got: %s, Match: false", got)
		}
	})
}

func format(data string) string {
	return "<div> " + data + " </div>"
}

