package azuretexttospeech

// AudioOutput defines supported audio formats
// Each incorporates a bitrate and encoding type. Azure Speech Service supports 24-KHz, 16-KHz, and 8-KHz audio outputs
// See https://docs.microsoft.com/en-us/azure/cognitive-services/speech-service/rest-apis#audio-outputs
type AudioOutput string

const (
	RIFF8Bit8kHzMonoPCM          AudioOutput = "riff-8khz-8bit-mono-mulaw"
	RIFF16Bit16kHzMonoPCM        AudioOutput = "riff-16khz-16bit-mono-pcm"
	RIFF16khz16kbpsMonoSiren     AudioOutput = "riff-16khz-16kbps-mono-siren"
	RIFF24khz16bitMonoPcm        AudioOutput = "riff-24khz-16bit-mono-pcm"
	RAW8Bit8kHzMonoMulaw         AudioOutput = "raw-8khz-8bit-mono-mulaw"
	RAW16Bit16kHzMonoMulaw       AudioOutput = "raw-16khz-16bit-mono-pcm"
	RAW24khz16bitMonoPcm         AudioOutput = "raw-24khz-16bit-mono-pcm"
	Ssml16khz16bitMonoTts        AudioOutput = "ssml-16khz-16bit-mono-tts"
	Audio16khz16kbpsMonoSiren    AudioOutput = "audio-16khz-16kbps-mono-siren"
	Audio16khz32kbitrateMonoMp3  AudioOutput = "audio-16khz-32kbitrate-mono-mp3"
	Audio16khz64kbitrateMonoMp3  AudioOutput = "audio-16khz-64kbitrate-mono-mp3"
	Audio16khz128kbitrateMonoMp3 AudioOutput = "audio-16khz-128kbitrate-mono-mp3"
	Audio24khz48kbitrateMonoMp3  AudioOutput = "audio-24khz-48kbitrate-mono-mp3"
	Audio24khz96kbitrateMonoMp3  AudioOutput = "audio-24khz-96kbitrate-mono-mp3"
)

// Gender type for the digitized language
type Gender string

const (
	// Male , Female are the static Gender constants for digitized voices.
	// See Gender in https://docs.microsoft.com/en-us/azure/cognitive-services/speech-service/language-support#standard-voices for breakdown
	Male   Gender = "Male"
	Female Gender = "Female"
)

// Region references the language or locale for text-to-speech.
// See "locale" in https://docs.microsoft.com/en-us/azure/cognitive-services/speech-service/language-support#standard-voices
type Region string

const (
	ArEG Region = "ar-EG"
	ArSA Region = "as-SA"
	BgBG Region = "bg-BG"
	CaES Region = "ca-ES"
	CsCZ Region = "cs-CZ"
	DaDK Region = "da-DK"
	DeAT Region = "de-AT"
	DeCH Region = "de-CH"
	DeDE Region = "de-DE"
	ElGR Region = "el-GR"
	EnAU Region = "en-AU"
	EnCA Region = "en-CA"
	EnGB Region = "en-GB"
	EnIE Region = "en-IE"
	EnIN Region = "en-IN"
	EnUS Region = "en-US"
	EsES Region = "es-ES"
	EsMX Region = "es-MX"
	FiFI Region = "fi-FI"
	FrCA Region = "fr-CA"
	FrCH Region = "fr-CH"
	FrFR Region = "fr-FR"
	HeIL Region = "he-IL"
	HiIN Region = "hi-IN"
	HrHR Region = "hr-HR"
	HuHU Region = "hu-HU"
	IdID Region = "id-ID"
	ItIT Region = "it-IT"
	JaJP Region = "ja-JP"
	KoKR Region = "ko-KR"
	MsMY Region = "ms-MY"
	NbNO Region = "nb-NO"
	NlNL Region = "nl-NL"
	PlPL Region = "pl-PL"
	PtBR Region = "pt-BR"
	PtPT Region = "pt-PT"
	RoRO Region = "ro-RO"
	RuRU Region = "ru-RU"
	SkSK Region = "sk-SK"
	SlSL Region = "sl-SL"
	SvSE Region = "sv-SE"
	TaIN Region = "ta-IN"
	TeIN Region = "te-IN"
	ThTH Region = "th-TH"
	TrTR Region = "tr-TR"
	ViVN Region = "vi-VN"
	ZhCN Region = "zh-CN"
	ZhHK Region = "zh-HK"
	ZhTW Region = "zh-TW"
)
