// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

type CreateLiveStreamRequest struct {
	PlaybackPolicy   []PlaybackPolicy   `json:"playback_policy,omitempty"`
	NewAssetSettings CreateAssetRequest `json:"new_asset_settings,omitempty"`
	// When live streaming software disconnects from Mux, either intentionally or due to a drop in the network, the Reconnect Window is the time in seconds that Mux should wait for the streaming software to reconnect before considering the live stream finished and completing the recorded asset. Defaults to 60 seconds on the API if not specified.  If not specified directly, Standard Latency streams have a Reconnect Window of 60 seconds; Reduced and Low Latency streams have a default of 0 seconds, or no Reconnect Window. For that reason, we suggest specifying a value other than zero for Reduced and Low Latency streams.  Reduced and Low Latency streams with a Reconnect Window greater than zero will insert slate media into the recorded asset while waiting for the streaming software to reconnect or when there are brief interruptions in the live stream media. When using a Reconnect Window setting higher than 60 seconds with a Standard Latency stream, we highly recommend enabling slate with the `use_slate_for_standard_latency` option.
	ReconnectWindow float32 `json:"reconnect_window,omitempty"`
	// By default, Standard Latency live streams do not have slate media inserted while waiting for live streaming software to reconnect to Mux. Setting this to true enables slate insertion on a Standard Latency stream.
	UseSlateForStandardLatency bool `json:"use_slate_for_standard_latency,omitempty"`
	// The URL of the image file that Mux should download and use as slate media during interruptions of the live stream media. This file will be downloaded each time a new recorded asset is created from the live stream. If this is not set, the default slate media will be used.
	ReconnectSlateUrl string `json:"reconnect_slate_url,omitempty"`
	Passthrough       string `json:"passthrough,omitempty"`
	// Force the live stream to only process the audio track when the value is set to true. Mux drops the video track if broadcasted.
	AudioOnly bool `json:"audio_only,omitempty"`
	// Describe the embedded closed caption contents of the incoming live stream.
	EmbeddedSubtitles []LiveStreamEmbeddedSubtitleSettings `json:"embedded_subtitles,omitempty"`
	// Configure the incoming live stream to include subtitles created with automatic speech recognition. Each Asset created from a live stream with `generated_subtitles` configured will automatically receive two text tracks. The first of these will have a `text_source` value of `generated_live`, and will be available with `ready` status as soon as the stream is live. The second text track will have a `text_source` value of `generated_live_final` and will contain subtitles with improved accuracy, timing, and formatting. However, `generated_live_final` tracks will not be available in `ready` status until the live stream ends. If an Asset has both `generated_live` and `generated_live_final` tracks that are `ready`, then only the `generated_live_final` track will be included during playback.
	GeneratedSubtitles []LiveStreamGeneratedSubtitleSettings `json:"generated_subtitles,omitempty"`
	// This field is deprecated. Please use `latency_mode` instead. Latency is the time from when the streamer transmits a frame of video to when you see it in the player. Set this if you want lower latency for your live stream. Read more here: https://mux.com/blog/reduced-latency-for-mux-live-streaming-now-available/
	ReducedLatency bool `json:"reduced_latency,omitempty"`
	// This field is deprecated. Please use `latency_mode` instead. Latency is the time from when the streamer transmits a frame of video to when you see it in the player. Setting this option will enable compatibility with the LL-HLS specification for low-latency streaming. This typically has lower latency than Reduced Latency streams, and cannot be combined with Reduced Latency.
	LowLatency bool `json:"low_latency,omitempty"`
	// Latency is the time from when the streamer transmits a frame of video to when you see it in the player. Set this as an alternative to setting low latency or reduced latency flags.
	LatencyMode string `json:"latency_mode,omitempty"`
	// Marks the live stream as a test live stream when the value is set to true. A test live stream can help evaluate the Mux Video APIs without incurring any cost. There is no limit on number of test live streams created. Test live streams are watermarked with the Mux logo and limited to 5 minutes. The test live stream is disabled after the stream is active for 5 mins and the recorded asset also deleted after 24 hours.
	Test             bool                           `json:"test,omitempty"`
	SimulcastTargets []CreateSimulcastTargetRequest `json:"simulcast_targets,omitempty"`
	// The time in seconds a live stream may be continuously active before being disconnected. Defaults to 12 hours.
	MaxContinuousDuration int32 `json:"max_continuous_duration,omitempty"`
}
