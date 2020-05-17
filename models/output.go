package models

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type OutputFile struct {
	aspect                string
	resolution            string
	videoBitRate          string
	videoBitRateTolerance int
	videoMaxBitRate       int
	videoMinBitrate       int
	videoCodec            string
	vframes               int
	frameRate             int
	audioRate             int
	maxKeyframe           int
	minKeyframe           int
	keyframeInterval      int
	audioCodec            string
	audioBitrate          string
	audioChannels         int
	audioVariableBitrate  bool
	bufferSize            int
	preset                string
	tune                  string
	audioProfile          string
	videoProfile          string
	target                string
	duration              string
	durationInput         string
	seekTime              string
	qscale                uint32
	crf                   uint32
	strict                int
	muxDelay              string
	seekUsingTsInput      bool
	seekTimeInput         string
	movFlags              string
	hideBanner            bool
	outputPath            string
	outputFormat          string
	copyTs                bool
	nativeFramerateInput  bool
	inputInitialOffset    string
	rtmpLive              string
	hlsPlaylistType       string
	hlsListSize           int
	hlsSegmentDuration    int
	hlsMasterPlaylistName string
	hlsSegmentFilename    string
	httpMethod            string
	httpKeepAlive         bool
	hwaccel               string
	streamIds             map[int]string
	videoFilter           string
	audioFilter           string
	skipVideo             bool
	skipAudio             bool
	compressionLevel      int
	mapMetadata           string
	tags                  map[string]string
	encryptionKey         string
	movflags              string
	bframe                int
	pixFmt                string
}

/*** SETTERS ***/
func (o *OutputFile) SetAudioFilter(v string) {
	o.audioFilter = v
}

func (o *OutputFile) SetVideoFilter(v string) {
	o.videoFilter = v
}

// Deprecated: Use SetVideoFilter instead.
func (o *OutputFile) SetFilter(v string) {
	o.SetVideoFilter(v)
}

func (o *OutputFile) SetAspect(v string) {
	o.aspect = v
}

func (o *OutputFile) SetResolution(v string) {
	o.resolution = v
}

func (o *OutputFile) SetVideoBitRate(v string) {
	o.videoBitRate = v
}

func (o *OutputFile) SetVideoBitRateTolerance(v int) {
	o.videoBitRateTolerance = v
}

func (o *OutputFile) SetVideoMaxBitrate(v int) {
	o.videoMaxBitRate = v
}

func (o *OutputFile) SetVideoMinBitRate(v int) {
	o.videoMinBitrate = v
}

func (o *OutputFile) SetVideoCodec(v string) {
	o.videoCodec = v
}

func (o *OutputFile) SetVframes(v int) {
	o.vframes = v
}

func (o *OutputFile) SetFrameRate(v int) {
	o.frameRate = v
}

func (o *OutputFile) SetAudioRate(v int) {
	o.audioRate = v
}

func (o *OutputFile) SetAudioVariableBitrate() {
	o.audioVariableBitrate = true
}

func (o *OutputFile) SetMaxKeyFrame(v int) {
	o.maxKeyframe = v
}

func (o *OutputFile) SetMinKeyFrame(v int) {
	o.minKeyframe = v
}

func (o *OutputFile) SetKeyframeInterval(v int) {
	o.keyframeInterval = v
}

func (o *OutputFile) SetAudioCodec(v string) {
	o.audioCodec = v
}

func (o *OutputFile) SetAudioBitRate(v string) {
	o.audioBitrate = v
}

func (o *OutputFile) SetAudioChannels(v int) {
	o.audioChannels = v
}

func (o *OutputFile) SetPixFmt(v string) {
	o.pixFmt = v
}

func (o *OutputFile) SetBufferSize(v int) {
	o.bufferSize = v
}

func (o *OutputFile) SetPreset(v string) {
	o.preset = v
}

func (o *OutputFile) SetTune(v string) {
	o.tune = v
}

func (o *OutputFile) SetAudioProfile(v string) {
	o.audioProfile = v
}

func (o *OutputFile) SetVideoProfile(v string) {
	o.videoProfile = v
}

func (o *OutputFile) SetDuration(v string) {
	o.duration = v
}

func (o *OutputFile) SetDurationInput(v string) {
	o.durationInput = v
}

func (o *OutputFile) SetSeekTime(v string) {
	o.seekTime = v
}

func (o *OutputFile) SetSeekTimeInput(v string) {
	o.seekTimeInput = v
}

// Q Scale must be integer between 1 to 31 - https://trac.ffmpeg.org/wiki/Encode/MPEG-4
func (o *OutputFile) SetQScale(v uint32) {
	o.qscale = v
}

func (o *OutputFile) SetCRF(v uint32) {
	o.crf = v
}

func (o *OutputFile) SetStrict(v int) {
	o.strict = v
}

func (o *OutputFile) SetSeekUsingTsInput(val bool) {
	o.seekUsingTsInput = val
}

func (o *OutputFile) SetCopyTs(val bool) {
	o.copyTs = val
}

/*** GETTERS ***/

// Deprecated: Use VideoFilter instead.
func (o *OutputFile) Filter() string {
	return o.VideoFilter()
}

func (o *OutputFile) VideoFilter() string {
	return o.videoFilter
}

func (o *OutputFile) AudioFilter() string {
	return o.audioFilter
}

func (o *OutputFile) Aspect() string {
	return o.aspect
}

func (o *OutputFile) Resolution() string {
	return o.resolution
}

func (o *OutputFile) VideoBitrate() string {
	return o.videoBitRate
}

func (o *OutputFile) VideoBitRateTolerance() int {
	return o.videoBitRateTolerance
}

func (o *OutputFile) VideoMaxBitRate() int {
	return o.videoMaxBitRate
}

func (o *OutputFile) VideoMinBitRate() int {
	return o.videoMinBitrate
}

func (o *OutputFile) VideoCodec() string {
	return o.videoCodec
}

func (o *OutputFile) Vframes() int {
	return o.vframes
}

func (o *OutputFile) FrameRate() int {
	return o.frameRate
}

func (o *OutputFile) GetPixFmt() string {
	return o.pixFmt
}

func (o *OutputFile) AudioRate() int {
	return o.audioRate
}

func (o *OutputFile) MaxKeyFrame() int {
	return o.maxKeyframe
}

func (o *OutputFile) MinKeyFrame() int {
	return o.minKeyframe
}

func (o *OutputFile) KeyFrameInterval() int {
	return o.keyframeInterval
}

func (o *OutputFile) AudioCodec() string {
	return o.audioCodec
}

func (o *OutputFile) AudioBitrate() string {
	return o.audioBitrate
}

func (o *OutputFile) AudioChannels() int {
	return o.audioChannels
}

func (o *OutputFile) BufferSize() int {
	return o.bufferSize
}

func (o *OutputFile) Target() string {
	return o.target
}

func (o *OutputFile) Duration() string {
	return o.duration
}

func (o *OutputFile) DurationInput() string {
	return o.durationInput
}

func (o *OutputFile) SeekTime() string {
	return o.seekTime
}

func (o *OutputFile) Preset() string {
	return o.preset
}

func (o *OutputFile) AudioProfile() string {
	return o.audioProfile
}

func (o *OutputFile) VideoProfile() string {
	return o.videoProfile
}

func (o *OutputFile) Tune() string {
	return o.tune
}

func (o *OutputFile) SeekTimeInput() string {
	return o.seekTimeInput
}

func (o *OutputFile) QScale() uint32 {
	return o.qscale
}

func (o *OutputFile) CRF() uint32 {
	return o.crf
}

func (o *OutputFile) Strict() int {
	return o.strict
}

func (o *OutputFile) MuxDelay() string {
	return o.muxDelay
}

func (o *OutputFile) SeekUsingTsInput() bool {
	return o.seekUsingTsInput
}

func (o *OutputFile) CopyTs() bool {
	return o.copyTs
}

func (o *OutputFile) MovFlags() string {
	return o.movFlags
}

func (o *OutputFile) HideBanner() bool {
	return o.hideBanner
}

func (o *OutputFile) OutputPath() string {
	return o.outputPath
}

func (o *OutputFile) OutputFormat() string {
	return o.outputFormat
}

func (o *OutputFile) NativeFramerateInput() bool {
	return o.nativeFramerateInput
}

func (o *OutputFile) RtmpLive() string {
	return o.rtmpLive
}

func (o *OutputFile) HlsListSize() int {
	return o.hlsListSize
}

func (o *OutputFile) HlsSegmentDuration() int {
	return o.hlsSegmentDuration
}

func (o *OutputFile) HlsMasterPlaylistName() string {
	return o.hlsMasterPlaylistName
}

func (o *OutputFile) HlsSegmentFilename() string {
	return o.hlsSegmentFilename
}

func (o *OutputFile) HlsPlaylistType() string {
	return o.hlsPlaylistType
}

func (o *OutputFile) InputInitialOffset() string {
	return o.inputInitialOffset
}

func (o *OutputFile) HttpMethod() string {
	return o.httpMethod
}

func (o *OutputFile) HttpKeepAlive() bool {
	return o.httpKeepAlive
}

func (o *OutputFile) HardwareAcceleration() string {
	return o.hwaccel
}

func (o *OutputFile) StreamIds() map[int]string {
	return o.streamIds
}

func (o *OutputFile) SkipVideo() bool {
	return o.skipVideo
}

func (o *OutputFile) SkipAudio() bool {
	return o.skipAudio
}

func (o *OutputFile) CompressionLevel() int {
	return o.compressionLevel
}

func (o *OutputFile) MapMetadata() string {
	return o.mapMetadata
}

func (o *OutputFile) Tags() map[string]string {
	return o.tags
}

func (o *OutputFile) SetEncryptionKey(v string) {
	o.encryptionKey = v
}

func (o *OutputFile) EncryptionKey() string {
	return o.encryptionKey
}

func (o *OutputFile) ObtainAudioFilter() []string {
	if o.audioFilter != "" {
		return []string{"-af", o.audioFilter}
	}
	return nil
}

func (o *OutputFile) ObtainVideoFilter() []string {
	if o.videoFilter != "" {
		return []string{"-vf", o.videoFilter}
	}
	return nil
}

func (o *OutputFile) ObtainAspect() []string {
	// Set aspect
	if o.resolution != "" {
		resolution := strings.Split(o.resolution, "x")
		if len(resolution) != 0 {
			width, _ := strconv.ParseFloat(resolution[0], 64)
			height, _ := strconv.ParseFloat(resolution[1], 64)
			return []string{"-aspect", fmt.Sprintf("%f", width/height)}
		}
	}

	if o.aspect != "" {
		return []string{"-aspect", o.aspect}
	}
	return nil
}

func (o *OutputFile) ObtainHardwareAcceleration() []string {
	if o.hwaccel != "" {
		return []string{"-hwaccel", o.hwaccel}
	}
	return nil
}

func (o *OutputFile) ObtainMovFlags() []string {
	if o.movFlags != "" {
		return []string{"-movflags", o.movFlags}
	}
	return nil
}

func (o *OutputFile) ObtainHideBanner() []string {
	if o.hideBanner {
		return []string{"-hide_banner"}
	}
	return nil
}

func (o *OutputFile) ObtainNativeFramerateInput() []string {
	if o.nativeFramerateInput {
		return []string{"-re"}
	}
	return nil
}

func (o *OutputFile) ObtainOutputPath() []string {
	if o.outputPath != "" {
		return []string{o.outputPath}
	}
	return nil
}

func (o *OutputFile) ObtainVideoCodec() []string {
	if o.videoCodec != "" {
		return []string{"-c:v", o.videoCodec}
	}
	return nil
}

func (o *OutputFile) ObtainVframes() []string {
	if o.vframes != 0 {
		return []string{"-vframes", fmt.Sprintf("%d", o.vframes)}
	}
	return nil
}

func (o *OutputFile) ObtainFrameRate() []string {
	if o.frameRate != 0 {
		return []string{"-r", fmt.Sprintf("%d", o.frameRate)}
	}
	return nil
}

func (o *OutputFile) ObtainAudioRate() []string {
	if o.audioRate != 0 {
		return []string{"-ar", fmt.Sprintf("%d", o.audioRate)}
	}
	return nil
}

func (o *OutputFile) ObtainResolution() []string {
	if o.resolution != "" {
		return []string{"-s", o.resolution}
	}
	return nil
}

func (o *OutputFile) ObtainVideoBitRate() []string {
	if o.videoBitRate != "" {
		return []string{"-b:v", o.videoBitRate}
	}
	return nil
}

func (o *OutputFile) ObtainAudioCodec() []string {
	if o.audioCodec != "" {
		return []string{"-c:a", o.audioCodec}
	}
	return nil
}

func (o *OutputFile) ObtainAudioBitRate() []string {
	switch {
	case !o.audioVariableBitrate && o.audioBitrate != "":
		return []string{"-b:a", o.audioBitrate}
	case o.audioVariableBitrate && o.audioBitrate != "":
		return []string{"-q:a", o.audioBitrate}
	case o.audioVariableBitrate:
		return []string{"-q:a", "0"}
	default:
		return nil
	}
}

func (o *OutputFile) ObtainAudioChannels() []string {
	if o.audioChannels != 0 {
		return []string{"-ac", fmt.Sprintf("%d", o.audioChannels)}
	}
	return nil
}

func (o *OutputFile) ObtainVideoMaxBitRate() []string {
	if o.videoMaxBitRate != 0 {
		return []string{"-maxrate", fmt.Sprintf("%dk", o.videoMaxBitRate)}
	}
	return nil
}

func (o *OutputFile) ObtainVideoMinBitRate() []string {
	if o.videoMinBitrate != 0 {
		return []string{"-minrate", fmt.Sprintf("%dk", o.videoMinBitrate)}
	}
	return nil
}

func (o *OutputFile) ObtainBufferSize() []string {
	if o.bufferSize != 0 {
		return []string{"-bufsize", fmt.Sprintf("%dk", o.bufferSize)}
	}
	return nil
}

func (o *OutputFile) ObtainVideoBitRateTolerance() []string {
	if o.videoBitRateTolerance != 0 {
		return []string{"-bt", fmt.Sprintf("%dk", o.videoBitRateTolerance)}
	}
	return nil
}

func (o *OutputFile) ObtainTarget() []string {
	if o.target != "" {
		return []string{"-target", o.target}
	}
	return nil
}

func (o *OutputFile) ObtainDuration() []string {
	if o.duration != "" {
		return []string{"-t", o.duration}
	}
	return nil
}

func (o *OutputFile) ObtainDurationInput() []string {
	if o.durationInput != "" {
		return []string{"-t", o.durationInput}
	}
	return nil
}

func (o *OutputFile) ObtainKeyframeInterval() []string {
	if o.keyframeInterval != 0 {
		return []string{"-g", fmt.Sprintf("%d", o.keyframeInterval)}
	}
	return nil
}

func (o *OutputFile) ObtainSeekTime() []string {
	if o.seekTime != "" {
		return []string{"-ss", o.seekTime}
	}
	return nil
}

func (o *OutputFile) ObtainSeekTimeInput() []string {
	if o.seekTimeInput != "" {
		return []string{"-ss", o.seekTimeInput}
	}
	return nil
}

func (o *OutputFile) ObtainPreset() []string {
	if o.preset != "" {
		return []string{"-preset", o.preset}
	}
	return nil
}

func (o *OutputFile) ObtainTune() []string {
	if o.tune != "" {
		return []string{"-tune", o.tune}
	}
	return nil
}

func (o *OutputFile) ObtainCRF() []string {
	if o.crf != 0 {
		return []string{"-crf", fmt.Sprintf("%d", o.crf)}
	}
	return nil
}

func (o *OutputFile) ObtainQScale() []string {
	if o.qscale != 0 {
		return []string{"-qscale", fmt.Sprintf("%d", o.qscale)}
	}
	return nil
}

func (o *OutputFile) ObtainStrict() []string {
	if o.strict != 0 {
		return []string{"-strict", fmt.Sprintf("%d", o.strict)}
	}
	return nil
}

func (o *OutputFile) ObtainVideoProfile() []string {
	if o.videoProfile != "" {
		return []string{"-profile:v", o.videoProfile}
	}
	return nil
}

func (o *OutputFile) ObtainAudioProfile() []string {
	if o.audioProfile != "" {
		return []string{"-profile:a", o.audioProfile}
	}
	return nil
}

func (o *OutputFile) ObtainCopyTs() []string {
	if o.copyTs {
		return []string{"-copyts"}
	}
	return nil
}

func (o *OutputFile) ObtainOutputFormat() []string {
	if o.outputFormat != "" {
		return []string{"-f", o.outputFormat}
	}
	return nil
}

func (o *OutputFile) ObtainMuxDelay() []string {
	if o.muxDelay != "" {
		return []string{"-muxdelay", o.muxDelay}
	}
	return nil
}

func (o *OutputFile) ObtainSeekUsingTsInput() []string {
	if o.seekUsingTsInput {
		return []string{"-seek_timestamp", "1"}
	}
	return nil
}

func (o *OutputFile) ObtainRtmpLive() []string {
	if o.rtmpLive != "" {
		return []string{"-rtmp_live", o.rtmpLive}
	} else {
		return nil
	}
}

func (o *OutputFile) ObtainHlsPlaylistType() []string {
	if o.hlsPlaylistType != "" {
		return []string{"-hls_playlist_type", o.hlsPlaylistType}
	} else {
		return nil
	}
}

func (o *OutputFile) ObtainInputInitialOffset() []string {
	if o.inputInitialOffset != "" {
		return []string{"-itsoffset", o.inputInitialOffset}
	} else {
		return nil
	}
}

func (o *OutputFile) ObtainHlsListSize() []string {
	if o.hlsListSize != 0 {
		return []string{"-hls_list_size", fmt.Sprintf("%d", o.hlsListSize)}
	}
	return nil
}

func (o *OutputFile) ObtainHlsSegmentDuration() []string {
	if o.hlsSegmentDuration != 0 {
		return []string{"-hls_time", fmt.Sprintf("%d", o.hlsSegmentDuration)}
	} else {
		return nil
	}
}

func (o *OutputFile) ObtainHlsMasterPlaylistName() []string {
	if o.hlsMasterPlaylistName != "" {
		return []string{"-master_pl_name", fmt.Sprintf("%s", o.hlsMasterPlaylistName)}
	} else {
		return nil
	}
}

func (o *OutputFile) ObtainHlsSegmentFilename() []string {
	if o.hlsSegmentFilename != "" {
		return []string{"-hls_segment_filename", fmt.Sprintf("%s", o.hlsSegmentFilename)}
	} else {
		return nil
	}
}

func (o *OutputFile) ObtainHttpMethod() []string {
	if o.httpMethod != "" {
		return []string{"-method", o.httpMethod}
	} else {
		return nil
	}
}

func (o *OutputFile) ObtainPixFmt() []string {
	if o.pixFmt != "" {
		return []string{"-pix_fmt", o.pixFmt}
	} else {
		return nil
	}
}

func (o *OutputFile) ObtainHttpKeepAlive() []string {
	if o.httpKeepAlive {
		return []string{"-multiple_requests", "1"}
	} else {
		return nil
	}
}

func (o *OutputFile) ObtainSkipVideo() []string {
	if o.skipVideo {
		return []string{"-vn"}
	} else {
		return nil
	}
}

func (o *OutputFile) ObtainSkipAudio() []string {
	if o.skipAudio {
		return []string{"-an"}
	} else {
		return nil
	}
}

func (o *OutputFile) ObtainStreamIds() []string {
	if o.streamIds != nil && len(o.streamIds) != 0 {
		result := []string{}
		for i, val := range o.streamIds {
			result = append(result, []string{"-streamid", fmt.Sprintf("%d:%s", i, val)}...)
		}
		return result
	}
	return nil
}

func (o *OutputFile) ObtainCompressionLevel() []string {
	if o.compressionLevel != 0 {
		return []string{"-compression_level", fmt.Sprintf("%d", o.compressionLevel)}
	}
	return nil
}

func (o *OutputFile) ObtainMapMetadata() []string {
	if o.mapMetadata != "" {
		return []string{"-map_metadata", o.mapMetadata}
	}
	return nil
}

func (o *OutputFile) ObtainEncryptionKey() []string {
	if o.encryptionKey != "" {
		return []string{"-hls_key_info_file", o.encryptionKey}
	}
	return nil
}

func (o *OutputFile) ObtainBframe() []string {
	if o.bframe != 0 {
		return []string{"-bf", fmt.Sprintf("%d", o.bframe)}
	}
	return nil
}

func (o *OutputFile) ObtainTags() []string {
	if o.tags != nil && len(o.tags) != 0 {
		result := []string{}
		for key, val := range o.tags {
			result = append(result, []string{"-metadata", fmt.Sprintf("%s=%s", key, val)}...)
		}
		return result
	}
	return nil
}

func (o *OutputFile) SetMovFlags(val string) {
	o.movFlags = val
}

func (o *OutputFile) SetHideBanner(val bool) {
	o.hideBanner = val
}

func (o *OutputFile) SetMuxDelay(val string) {
	o.muxDelay = val
}

func (o *OutputFile) SetOutputPath(val string) {
	o.outputPath = val
}

func (o *OutputFile) SetOutputFormat(val string) {
	o.outputFormat = val
}

func (o *OutputFile) SetNativeFramerateInput(val bool) {
	o.nativeFramerateInput = val
}

func (o *OutputFile) SetRtmpLive(val string) {
	o.rtmpLive = val
}

func (o *OutputFile) SetHlsListSize(val int) {
	o.hlsListSize = val
}

func (o *OutputFile) SetHlsSegmentDuration(val int) {
	o.hlsSegmentDuration = val
}

func (o *OutputFile) SetHlsPlaylistType(val string) {
	o.hlsPlaylistType = val
}

func (o *OutputFile) SetHlsMasterPlaylistName(val string) {
	o.hlsMasterPlaylistName = val
}

func (o *OutputFile) SetHlsSegmentFilename(val string) {
	o.hlsSegmentFilename = val
}

func (o *OutputFile) SetHttpMethod(val string) {
	o.httpMethod = val
}

func (o *OutputFile) SetHttpKeepAlive(val bool) {
	o.httpKeepAlive = val
}

func (o *OutputFile) SetHardwareAcceleration(val string) {
	o.hwaccel = val
}

func (o *OutputFile) SetInputInitialOffset(val string) {
	o.inputInitialOffset = val
}

func (o *OutputFile) SetStreamIds(val map[int]string) {
	o.streamIds = val
}

func (o *OutputFile) SetSkipVideo(val bool) {
	o.skipVideo = val
}

func (o *OutputFile) SetSkipAudio(val bool) {
	o.skipAudio = val
}

func (o *OutputFile) SetCompressionLevel(val int) {
	o.compressionLevel = val
}

func (o *OutputFile) SetMapMetadata(val string) {
	o.mapMetadata = val
}

func (o *OutputFile) SetTags(val map[string]string) {
	o.tags = val
}

func (o *OutputFile) SetBframe(v int) {
	o.bframe = v
}

func (o *OutputFile) ToStrCommand() []string {
	var strCommand []string

	opts := []string{
		"SeekTimeInput",
		"SeekUsingTsInput",
		"NativeFramerateInput",
		"DurationInput",
		"RtmpLive",
		"InputInitialOffset",
		"HardwareAcceleration",
		"HideBanner",
		"Aspect",
		"Resolution",
		"FrameRate",
		"AudioRate",
		"VideoCodec",
		"Vframes",
		"VideoBitRate",
		"VideoBitRateTolerance",
		"VideoMaxBitRate",
		"VideoMinBitRate",
		"VideoProfile",
		"SkipVideo",
		"AudioCodec",
		"AudioBitRate",
		"AudioChannels",
		"AudioProfile",
		"SkipAudio",
		"CRF",
		"QScale",
		"Strict",
		"BufferSize",
		"MuxDelay",
		"KeyframeInterval",
		"Preset",
		"PixFmt",
		"Tune",
		"Target",
		"SeekTime",
		"Duration",
		"CopyTs",
		"StreamIds",
		"MovFlags",
		"OutputFormat",
		"HlsListSize",
		"HlsSegmentDuration",
		"HlsPlaylistType",
		"HlsMasterPlaylistName",
		"HlsSegmentFilename",
		"AudioFilter",
		"VideoFilter",
		"HttpMethod",
		"HttpKeepAlive",
		"CompressionLevel",
		"MapMetadata",
		"Tags",
		"EncryptionKey",
		"Bframe",
		"MovFlags",
		"OutputPath",
	}

	for _, name := range opts {
		opt := reflect.ValueOf(o).MethodByName(fmt.Sprintf("Obtain%s", name))
		if (opt != reflect.Value{}) {
			result := opt.Call([]reflect.Value{})

			if val, ok := result[0].Interface().([]string); ok {
				strCommand = append(strCommand, val...)
			}
		}
	}

	return strCommand
}
