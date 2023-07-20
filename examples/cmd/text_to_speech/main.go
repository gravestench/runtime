package main

import (
	"math/rand"
	"time"

	"github.com/gravestench/runtime"
	"github.com/gravestench/runtime/examples/services/config_file"
	"github.com/gravestench/runtime/examples/services/text_to_speech"
)

func main() {
	rt := runtime.New()

	tts := &text_to_speech.Service{}

	rt.Add(&config_file.Service{RootDirectory: "~/.config/runtime/example/text_to_speech"})
	rt.Add(tts)

	go func() {
		for {
			tts.SetVoice("en-UK")
			tts.Speak(generateRandomPhrase())
		}
	}()

	rt.Run()
}

var (
	adjectives = []string{
		"happy", "sad", "brave", "kind", "smart", "funny", "silly",
		"crazy", "friendly", "honest", "curious", "energetic",
		"thoughtful", "creative", "patient", "generous",
	}

	nouns = []string{
		"dog", "cat", "book", "friend", "world", "car", "sun",
		"moon", "flower", "tree", "house", "coffee", "song",
		"smile", "dream", "mountain", "river", "ocean", "cloud",
	}

	verbs = []string{
		"run", "jump", "play", "sing", "dance", "read", "write",
		"learn", "explore", "create", "help", "love", "laugh",
		"think", "smile", "dream", "inspire", "imagine", "enjoy",
	}
)

func generateRandomPhrase() string {
	rand.Seed(time.Now().UnixNano())

	phrase := getRandomElement(adjectives) + " " +
		getRandomElement(nouns) + " " +
		getRandomElement(verbs) + " " +
		getRandomElement(adjectives) + " " +
		getRandomElement(nouns) + "."

	return phrase
}

func getRandomElement(list []string) string {
	index := rand.Intn(len(list))
	return list[index]
}
