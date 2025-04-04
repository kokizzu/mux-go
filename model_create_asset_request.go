// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

type CreateAssetRequest struct {
	// Deprecated. Use `inputs` instead, which accepts an identical type.
	Input []InputSettings `json:"input,omitempty"`
	// An array of objects that each describe an input file to be used to create the asset. As a shortcut, input can also be a string URL for a file when only one input file is used. See `input[].url` for requirements.
	Inputs []InputSettings `json:"inputs,omitempty"`
	// Deprecated. Use `playback_policies` instead, which accepts an identical type.
	PlaybackPolicy []PlaybackPolicy `json:"playback_policy,omitempty"`
	// An array of playback policy names that you want applied to this asset and available through `playback_ids`. Options include:  * `\"public\"` (anyone with the playback URL can stream the asset). * `\"signed\"` (an additional access token is required to play the asset).  If no `playback_policies` are set, the asset will have no playback IDs and will therefore not be playable. For simplicity, a single string name can be used in place of the array in the case of only one playback policy.
	PlaybackPolicies []PlaybackPolicy `json:"playback_policies,omitempty"`
	// An array of playback policy objects that you want applied to this asset and available through `playback_ids`. `advanced_playback_policies` must be used instead of `playback_policies` when creating a DRM playback ID.
	AdvancedPlaybackPolicies []CreatePlaybackIdRequest `json:"advanced_playback_policies,omitempty"`
	PerTitleEncode           bool                      `json:"per_title_encode,omitempty"`
	// You can set this field to anything you want. It will be included in the asset details and related webhooks. If you're looking for more structured metadata, such as `title` or `external_id`, you can use the `meta` object instead. **Max: 255 characters**.
	Passthrough string `json:"passthrough,omitempty"`
	// Deprecated. See the [Static Renditions API](https://www.mux.com/docs/guides/enable-static-mp4-renditions) for the updated API.  Specify what level of support for mp4 playback. You may not enable both `mp4_support` and  `static_renditions`.  * The `capped-1080p` option produces a single MP4 file, called `capped-1080p.mp4`, with the video resolution capped at 1080p. This option produces an `audio.m4a` file for an audio-only asset. * The `audio-only` option produces a single M4A file, called `audio.m4a` for a video or an audio-only asset. MP4 generation will error when this option is specified for a video-only asset. * The `audio-only,capped-1080p` option produces both the `audio.m4a` and `capped-1080p.mp4` files. Only the `capped-1080p.mp4` file is produced for a video-only asset, while only the `audio.m4a` file is produced for an audio-only asset.  The `standard`(deprecated) option produces up to three MP4 files with different levels of resolution (`high.mp4`, `medium.mp4`, `low.mp4`, or `audio.m4a` for an audio-only asset).  MP4 files are not produced for `none` (default).  In most cases you should use our default HLS-based streaming playback (`{playback_id}.m3u8`) which can automatically adjust to viewers' connection speeds, but an mp4 can be useful for some legacy devices or downloading for offline playback. See the [Download your videos guide](https://docs.mux.com/guides/enable-static-mp4-renditions) for more information.
	Mp4Support string `json:"mp4_support,omitempty"`
	// Normalize the audio track loudness level. This parameter is only applicable to on-demand (not live) assets.
	NormalizeAudio bool `json:"normalize_audio,omitempty"`
	// Specify what level (if any) of support for master access. Master access can be enabled temporarily for your asset to be downloaded. See the [Download your videos guide](https://docs.mux.com/guides/enable-static-mp4-renditions) for more information.
	MasterAccess string `json:"master_access,omitempty"`
	// Marks the asset as a test asset when the value is set to true. A Test asset can help evaluate the Mux Video APIs without incurring any cost. There is no limit on number of test assets created. Test asset are watermarked with the Mux logo, limited to 10 seconds, deleted after 24 hrs.
	Test bool `json:"test,omitempty"`
	// Max resolution tier can be used to control the maximum `resolution_tier` your asset is encoded, stored, and streamed at. If not set, this defaults to `1080p`.
	MaxResolutionTier string `json:"max_resolution_tier,omitempty"`
	// This field is deprecated. Please use `video_quality` instead. The encoding tier informs the cost, quality, and available platform features for the asset. The default encoding tier for an account can be set in the Mux Dashboard. [See the video quality guide for more details.](https://docs.mux.com/guides/use-video-quality-levels)
	EncodingTier string `json:"encoding_tier,omitempty"`
	// The video quality controls the cost, quality, and available platform features for the asset. The default video quality for an account can be set in the Mux Dashboard. This field replaces the deprecated `encoding_tier` value. [See the video quality guide for more details.](https://docs.mux.com/guides/use-video-quality-levels)
	VideoQuality string `json:"video_quality,omitempty"`
	// An array of static renditions to create for this asset. You may not enable both `static_renditions` and `mp4_support (the latter being deprecated)`
	StaticRenditions []CreateStaticRenditionRequest `json:"static_renditions,omitempty"`
	Meta             AssetMetadata                  `json:"meta,omitempty"`
}
