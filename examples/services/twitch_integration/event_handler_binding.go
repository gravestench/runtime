package twitch_integration

import (
	"time"

	"github.com/gravestench/runtime"
)

// loopBindHandlers iterates over all services every second and binds event
// handlers for services that implement interfaces found in abstract.go
func (s *Service) loopBindHandlers(rt runtime.R) {
	// we will keep track service we've already bound
	bound := make(map[string]any)

	for {
		time.Sleep(time.Second * 1)

		for _, service := range rt.Services() {
			// if already bound, do nothing
			if _, isBound := bound[service.Name()]; isBound {
				continue
			}

			// otherwise bind, add to list
			s.bindService(service)
			bound[service.Name()] = service
		}
	}
}

func (s *Service) bindService(service runtime.Service) {
	if service == s {
		return
	}

	if handler, ok := service.(OnConnect); ok {
		s.twitchIrcClient.OnConnect(handler.OnTwitchConnect)
	}

	if handler, ok := service.(OnWhisperMessage); ok {
		s.twitchIrcClient.OnWhisperMessage(handler.OnTwitchWhisperMessage)
	}

	if handler, ok := service.(OnPrivateMessage); ok {
		s.twitchIrcClient.OnPrivateMessage(handler.OnTwitchPrivateMessage)
	}

	if handler, ok := service.(OnClearChatMessage); ok {
		s.twitchIrcClient.OnClearChatMessage(handler.OnTwitchClearChatMessage)
	}

	if handler, ok := service.(OnClearMessage); ok {
		s.twitchIrcClient.OnClearMessage(handler.OnTwitchClearMessage)
	}

	if handler, ok := service.(OnRoomStateMessage); ok {
		s.twitchIrcClient.OnRoomStateMessage(handler.OnTwitchRoomStateMessage)
	}

	if handler, ok := service.(OnUserNoticeMessage); ok {
		s.twitchIrcClient.OnUserNoticeMessage(handler.OnTwitchUserNoticeMessage)
	}

	if handler, ok := service.(OnUserStateMessage); ok {
		s.twitchIrcClient.OnUserStateMessage(handler.OnTwitchUserStateMessage)
	}

	if handler, ok := service.(OnGlobalUserStateMessage); ok {
		s.twitchIrcClient.OnGlobalUserStateMessage(handler.OnTwitchGlobalUserStateMessage)
	}

	if handler, ok := service.(OnNoticeMessage); ok {
		s.twitchIrcClient.OnNoticeMessage(handler.OnTwitchNoticeMessage)
	}

	if handler, ok := service.(OnUserJoinMessage); ok {
		s.twitchIrcClient.OnUserJoinMessage(handler.OnTwitchUserJoinMessage)
	}

	if handler, ok := service.(OnUserPartMessage); ok {
		s.twitchIrcClient.OnUserPartMessage(handler.OnTwitchUserPartMessage)
	}

	if handler, ok := service.(OnReconnectMessage); ok {
		s.twitchIrcClient.OnReconnectMessage(handler.OnTwitchReconnectMessage)
	}

	if handler, ok := service.(OnNamesMessage); ok {
		s.twitchIrcClient.OnNamesMessage(handler.OnTwitchNamesMessage)
	}

	if handler, ok := service.(OnPingMessage); ok {
		s.twitchIrcClient.OnPingMessage(handler.OnTwitchPingMessage)
	}

	if handler, ok := service.(OnPongMessage); ok {
		s.twitchIrcClient.OnPongMessage(handler.OnTwitchPongMessage)
	}

	if handler, ok := service.(OnUnsetMessage); ok {
		s.twitchIrcClient.OnUnsetMessage(handler.OnTwitchUnsetMessage)
	}

	if handler, ok := service.(OnPingSent); ok {
		s.twitchIrcClient.OnPingSent(handler.OnTwitchPingSent)
	}
}
