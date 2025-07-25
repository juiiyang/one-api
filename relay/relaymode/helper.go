package relaymode

import "strings"

func GetByPath(path string) int {
	relayMode := Unknown
	if strings.HasPrefix(path, "/v1/oneapi/proxy") {
		relayMode = Proxy
	} else if strings.HasPrefix(path, "/v1/responses") {
		relayMode = ResponseAPI
	} else if strings.HasPrefix(path, "/v1/messages") {
		relayMode = ClaudeMessages
	} else if strings.HasPrefix(path, "/v1/chat/completions") {
		relayMode = ChatCompletions
	} else if strings.HasPrefix(path, "/v1/completions") {
		relayMode = Completions
	} else if strings.HasPrefix(path, "/v1/embeddings") {
		relayMode = Embeddings
	} else if strings.HasPrefix(path, "/v1/rerank") {
		relayMode = Rerank
	} else if strings.HasSuffix(path, "/rerank") {
		relayMode = Rerank
	} else if strings.HasSuffix(path, "/rerankers") {
		relayMode = Rerank
	} else if strings.HasSuffix(path, "embeddings") {
		relayMode = Embeddings
	} else if strings.HasPrefix(path, "/v1/moderations") {
		relayMode = Moderations
	} else if strings.HasPrefix(path, "/v1/images/generations") {
		relayMode = ImagesGenerations
	} else if strings.HasPrefix(path, "/v1/edits") {
		relayMode = Edits
	} else if strings.HasPrefix(path, "/v1/audio/speech") {
		relayMode = AudioSpeech
	} else if strings.HasPrefix(path, "/v1/audio/transcriptions") {
		relayMode = AudioTranscription
	} else if strings.HasPrefix(path, "/v1/audio/translations") {
		relayMode = AudioTranslation
	} else if strings.HasPrefix(path, "/v1/images/edits") {
		relayMode = ImagesEdits
	}

	return relayMode
}
