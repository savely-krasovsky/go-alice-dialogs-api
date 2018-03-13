package aliceapi

type Req struct {
	Meta    meta    `json:"meta"`
	Request Request `json:"request"`
	Session session `json:"session"`
	Version string  `json:"version"`
}

type Request struct {
	Type              string      `json:"type"`
	Markup            markup      `json:"markup"`
	Command           string      `json:"command"`
	OriginalUtterance string      `json:"original_utterance"`
	Payload           interface{} `json:"payload"`
}

type markup struct {
	DangerousContext bool `json:"dangerous_context"`
}

type Res struct {
	Response Response `json:"response"`
	Session  session  `json:"session"`
	Version  string      `json:"version"`
}

type Response struct {
	Text       string   `json:"text"`
	TTS        string   `json:"tts"`
	Buttons    []Button `json:"buttons"`
	EndSession bool     `json:"end_session"`
}

type Button struct {
	Title   string      `json:"title"`
	Payload interface{} `json:"payload"`
	URL     string      `json:"url"`
	Hide    bool        `json:"hide"`
}

type meta struct {
	Locale   string `json:"locale"`
	Timezone string `json:"timezone"`
	ClientID string `json:"client_id"`
}

type session struct {
	New       bool   `json:"new"`
	SessionID string `json:"session_id"`
	MessageID int    `json:"message_id"`
	SkillID   string `json:"skill_id"`
	UserID    string `json:"user_id"`
}