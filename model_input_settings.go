// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

// An array of objects that each describe an input file to be used to create the asset. As a shortcut, `input` can also be a string URL for a file when only one input file is used. See `input[].url` for requirements.
type InputSettings struct {
	// The URL of the file that Mux should download and use. * For the main input file, this should be the URL to the muxed file for Mux to download, for example an MP4, MOV, MKV, or TS file. Mux supports most audio/video file formats and codecs, but for fastest processing, you should [use standard inputs wherever possible](https://docs.mux.com/guides/minimize-processing-time). * For `audio` tracks, the URL is the location of the audio file for Mux to download, for example an M4A, WAV, or MP3 file. Mux supports most audio file formats and codecs, but for fastest processing, you should [use standard inputs wherever possible](https://docs.mux.com/guides/minimize-processing-time). * For `text` tracks, the URL is the location of subtitle/captions file. Mux supports [SubRip Text (SRT)](https://en.wikipedia.org/wiki/SubRip) and [Web Video Text Tracks](https://www.w3.org/TR/webvtt1/) formats for ingesting Subtitles and Closed Captions. * For Watermarking or Overlay, the URL is the location of the watermark image. The maximum size is 4096x4096. * When creating clips from existing Mux assets, the URL is defined with `mux://assets/{asset_id}` template where `asset_id` is the Asset Identifier for creating the clip from. The url property may be omitted on the first input object when providing asset settings for LiveStream and Upload objects, in order to configure settings related to the primary (live stream or direct upload) input.
	Url             string                       `json:"url,omitempty"`
	OverlaySettings InputSettingsOverlaySettings `json:"overlay_settings,omitempty"`
	// Generate subtitle tracks using automatic speech recognition with this configuration. This may only be provided for the first input object (the main input file). For direct uploads, this first input should omit the url parameter, as the main input file is provided via the direct upload. This will create subtitles based on the audio track ingested from that main input file. Note that subtitle generation happens after initial ingest, so the generated tracks will be in the `preparing` state when the asset transitions to `ready`.
	GeneratedSubtitles []AssetGeneratedSubtitleSettings `json:"generated_subtitles,omitempty"`
	// The time offset in seconds from the beginning of the video indicating the clip's starting marker. The default value is 0 when not included. This parameter is only applicable for creating clips when `input.url` has `mux://assets/{asset_id}` format.
	StartTime float64 `json:"start_time,omitempty"`
	// The time offset in seconds from the beginning of the video, indicating the clip's ending marker. The default value is the duration of the video when not included. This parameter is only applicable for creating clips when `input.url` has `mux://assets/{asset_id}` format.
	EndTime float64 `json:"end_time,omitempty"`
	// This parameter is required for `text` type tracks.
	Type string `json:"type,omitempty"`
	// Type of text track. This parameter only supports subtitles value. For more information on Subtitles / Closed Captions, [see this blog post](https://mux.com/blog/subtitles-captions-webvtt-hls-and-those-magic-flags/). This parameter is required for `text` type tracks.
	TextType string `json:"text_type,omitempty"`
	// The language code value must be a valid [BCP 47](https://tools.ietf.org/html/bcp47) specification compliant value. For example, `en` for English or `en-US` for the US version of English. This parameter is required for `text` and `audio` track types.
	LanguageCode string `json:"language_code,omitempty"`
	// The name of the track containing a human-readable description. This value must be unique within each group of `text` or `audio` track types. The HLS manifest will associate a subtitle text track with this value. For example, the value should be \"English\" for a subtitle text track with `language_code` set to `en`. This optional parameter should be used only for `text` and `audio` type tracks. This parameter can be optionally provided for the first video input to denote the name of the muxed audio track if present. If this parameter is not included, Mux will auto-populate based on the `input[].language_code` value.
	Name string `json:"name,omitempty"`
	// Indicates the track provides Subtitles for the Deaf or Hard-of-hearing (SDH). This optional parameter should be used for tracks with `type` of `text` and `text_type` set to `subtitles`.
	ClosedCaptions bool `json:"closed_captions,omitempty"`
	// This optional parameter should be used for tracks with `type` of `text` and `text_type` set to `subtitles`.
	Passthrough string `json:"passthrough,omitempty"`
}
