package spankwire

type SpankwireEmbedCode struct {
	Embed SpankwireCode `json:"embed,omitempty"`
}

type SpankwireCode struct {
	Code string `json:"code,omitempty"`
}
