package lpbot

type LPBot struct {
	Hash string `json:"hash"`
	Bot  struct {
		ID                      string `json:"id"`
		Name                    string `json:"name"`
		ChatBotType             string `json:"chatBotType"`
		Status                  string `json:"status"`
		GetStartedButtonPayload string `json:"getStartedButtonPayload"`
		CreationTime            int64  `json:"creationTime"`
		ModificationTime        int64  `json:"modificationTime"`
		Demo                    bool   `json:"demo"`
		SkipNLP                 bool   `json:"skipNLP"`
		Language                string `json:"language"`
		BotAttributes           struct {
			EnableDebug      string `json:"ENABLE_DEBUG"`
			SmallTalkEnabled string `json:"SMALL_TALK_ENABLED"`
		} `json:"botAttributes"`
		SessionLength      int    `json:"sessionLength"`
		PassThroughMode    bool   `json:"passThroughMode"`
		TranscriptDisabled bool   `json:"transcriptDisabled"`
		Version            string `json:"version"`
		PublicBot          bool   `json:"publicBot"`
		TransferGroupID    string `json:"transferGroupId"`
		Channel            string `json:"channel"`
		ReadOnly           bool   `json:"readOnly"`
		SmallTalkEnabled   bool   `json:"smallTalkEnabled"`
	} `json:"bot"`
	Responder           []any `json:"responder"`
	RequiredContext     []any `json:"requiredContext"`
	ConversationMessage []struct {
		ID                string `json:"id"`
		ChatBotID         string `json:"chatBotId"`
		UserInputRequired bool   `json:"userInputRequired"`
		Name              string `json:"name"`
		Type              string `json:"type"`
		Content           struct {
			ContentType string `json:"contentType"`
			Results     struct {
				Type string `json:"type"`
				Tile struct {
					TileData []struct {
						Text           string `json:"text"`
						Buttons        []any  `json:"buttons"`
						QuickReplyList []any  `json:"quickReplyList"`
					} `json:"tileData"`
				} `json:"tile"`
			} `json:"results"`
		} `json:"content"`
		Group           string `json:"group"`
		Status          string `json:"status"`
		ResponseMatches []struct {
			Conditions        []any `json:"conditions"`
			ContextConditions []any `json:"contextConditions"`
			Action            struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"action"`
			ContextDataVariables []any `json:"contextDataVariables"`
		} `json:"responseMatches"`
		InteractionType   string   `json:"interactionType"`
		PrevMessageID     string   `json:"prevMessageId,omitempty"`
		Pattern           []string `json:"pattern,omitempty"`
		PreProcessMessage string   `json:"preProcessMessage,omitempty"`
		NextMessageID     string   `json:"nextMessageId,omitempty"`
	} `json:"conversationMessage"`
	Menus  []any `json:"menus"`
	Groups []struct {
		ChatBotID                      string `json:"chatBotId"`
		ID                             string `json:"id"`
		Name                           string `json:"name"`
		CreationTime                   string `json:"creationTime"`
		ModificationTime               string `json:"modificationTime"`
		DialogType                     string `json:"dialogType"`
		Status                         string `json:"status"`
		DisambiguteOnlySelectedDomains bool   `json:"disambiguteOnlySelectedDomains"`
	} `json:"groups"`
	Platforms []struct {
		ID              string `json:"id"`
		ChatbotID       string `json:"chatbotId"`
		IntegrationType string `json:"integrationType"`
		Platform        string `json:"platform"`
		Status          string `json:"status"`
	} `json:"platforms"`
	AssociatedDomains []any `json:"associatedDomains"`
	DialogTemplates   []any `json:"dialogTemplates"`
}