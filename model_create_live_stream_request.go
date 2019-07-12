// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

type CreateLiveStreamRequest struct {
	PlaybackPolicy   []PlaybackPolicy   `json:"playback_policy,omitempty"`
	NewAssetSettings CreateAssetRequest `json:"new_asset_settings,omitempty"`
	// When live streaming software disconnects from Mux, either intentionally or due to a drop in the network, the Reconnect Window is the time in seconds that Mux should wait for the streaming software to reconnect before considering the live stream finished and completing the recorded asset. Default: 60 seconds
	ReconnectWindow float32 `json:"reconnect_window,omitempty"`
	Passthrough     string  `json:"passthrough,omitempty"`
}
