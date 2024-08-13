package lpbot
type BotAttributes struct {
	EnableDebug      string `json:"ENABLE_DEBUG"`
	SmallTalkEnabled string `json:"SMALL_TALK_ENABLED"`
}

type ContentTileData struct {
	Text           string   `json:"text"`
	Buttons        []string `json:"buttons"`
	QuickReplyList []string `json:"quickReplyList"`
}

type ContentResults struct {
	Type string `json:"type"`
	Tile struct {
		TileData []ContentTileData `json:"tileData"`
	} `json:"tile"`
}

type Content struct {
	ContentType string        `json:"contentType"`
	Results     ContentResults `json:"results"`
}

type ResponseMatch struct {
	Conditions          []string `json:"conditions"`
	ContextConditions   []string `json:"contextConditions"`
	Action              struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"action"`
	ContextDataVariables []string `json:"contextDataVariables"`
}

type ConversationMessage struct {
	ID              string          `json:"id"`
	ChatBotId       string          `json:"chatBotId"`
	UserInputRequired bool          `json:"userInputRequired"`
	Name            string          `json:"name"`
	Type            string          `json:"type"`
	Content         Content         `json:"content"`
	Group           string          `json:"group"`
	Status          string          `json:"status"`
	PrevMessageId   string          `json:"prevMessageId,omitempty"`
	ResponseMatches []ResponseMatch `json:"responseMatches"`
	InteractionType string          `json:"interactionType"`
	Pattern         []string        `json:"pattern,omitempty"`
	PreProcessMessage string        `json:"preProcessMessage,omitempty"`
	NextMessageId   string          `json:"nextMessageId,omitempty"`
}

type Group struct {
	ChatBotId                    string `json:"chatBotId"`
	ID                           string `json:"id"`
	Name                         string `json:"name"`
	CreationTime                 string `json:"creationTime"`
	ModificationTime             string `json:"modificationTime"`
	DialogType                   string `json:"dialogType"`
	Status                       string `json:"status"`
	DisambiguteOnlySelectedDomains bool  `json:"disambiguteOnlySelectedDomains"`
}

type Platform struct {
	ID             string `json:"id"`
	ChatbotId      string `json:"chatbotId"`
	IntegrationType string `json:"integrationType"`
	Platform       string `json:"platform"`
	Status         string `json:"status"`
}

type Bot struct {
	ID                string       `json:"id"`
	Name              string       `json:"name"`
	ChatBotType       string       `json:"chatBotType"`
	Status            string       `json:"status"`
	GetStartedButtonPayload string `json:"getStartedButtonPayload"`
	CreationTime      int64        `json:"creationTime"`
	ModificationTime  int64        `json:"modificationTime"`
	Demo              bool         `json:"demo"`
	SkipNLP           bool         `json:"skipNLP"`
	Language          string       `json:"language"`
	BotAttributes     BotAttributes `json:"botAttributes"`
	SessionLength     int          `json:"sessionLength"`
	PassThroughMode   bool         `json:"passThroughMode"`
	TranscriptDisabled bool        `json:"transcriptDisabled"`
	Version           string       `json:"version"`
	PublicBot         bool         `json:"publicBot"`
	TransferGroupId   string       `json:"transferGroupId"`
	Channel           string       `json:"channel"`
	ReadOnly          bool         `json:"readOnly"`
	SmallTalkEnabled  bool         `json:"smallTalkEnabled"`
}

type LPBot struct {
	Hash               string               `json:"hash"`
	Bot                Bot                  `json:"bot"`
	Responder          []string             `json:"responder"`
	RequiredContext    []string             `json:"requiredContext"`
	ConversationMessage []ConversationMessage `json:"conversationMessage"`
	Menus              []string             `json:"menus"`
	Groups             []Group              `json:"groups"`
	Platforms          []Platform           `json:"platforms"`
	AssociatedDomains  []string             `json:"associatedDomains"`
	DialogTemplates    []string             `json:"dialogTemplates"`
}