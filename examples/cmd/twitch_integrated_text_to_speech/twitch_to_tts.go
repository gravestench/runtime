package main

import (
	"math/rand"

	"github.com/gempir/go-twitch-irc/v2"

	"github.com/gravestench/runtime"
	"github.com/gravestench/runtime/examples/services/text_to_speech"
	"github.com/gravestench/runtime/pkg"
)

// this service will just connect the TTS to the twitch integration service
type glueService struct {
	tts text_to_speech.ConvertsTextToSpeech
}

func (g *glueService) OnTwitchPrivateMessage(message twitch.PrivateMessage) {
	voices := g.tts.Voices()

	randoVoice := voices[rand.Intn(len(voices))]

	g.tts.SetVoice("en-UK")
	g.tts.Speak(message.User.Name + " says: ")

	g.tts.SetVoice(randoVoice)
	g.tts.Speak(message.Message)
}

func (g *glueService) DependenciesResolved() bool {
	return g.tts != nil
}

func (g *glueService) ResolveDependencies(runtime runtime.R) {
	for _, service := range runtime.Services() {
		if candidate, ok := service.(text_to_speech.ConvertsTextToSpeech); ok {
			g.tts = candidate
		}
	}
}

func (g *glueService) Init(rt pkg.IsRuntime) {
	// do nothing
}

func (g *glueService) Name() string {
	return "glue service: tts <-> twitch integration"
}
