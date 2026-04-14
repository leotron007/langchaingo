package llms

// CallOption is a function that configures a CallOptions.
type CallOption func(*CallOptions)

// CallOptions is a set of options for calling models. Not all models support
// all options.
type CallOptions struct {
	// Model is the model to use.
	Model string `json:"model"`
	// MaxTokens is the maximum number of tokens to generate.
	MaxTokens int `json:"max_tokens"`
	// Temperature is the temperature for sampling, between 0 and 1.
	// A value of 0 makes the model deterministic; 1 makes it more creative.
	Temperature float64 `json:"temperature"`
	// StopWords is a list of words to stop on.
	StopWords []string `json:"stop_words"`
	// StreamingFunc is a function to be called for each chunk of a streaming response.
	// Return an error to stop streaming early.
	StreamingFunc func(ctx context.Context, chunk []byte) error `json:"-"`
	// TopK is the number of tokens to consider for top-k sampling.
	TopK int `json:"top_k"`
	// TopP is the cumulative probability for top-p sampling.
	TopP float64 `json:"top_p"`
	// Seed is a seed for deterministic sampling.
	Seed int `json:"seed"`
	// MinLength is the minimum length of the generated text.
	MinLength int `json:"min_length"`
	// MaxLength is the maximum length of the generated text.
	MaxLength int `json:"max_length"`
	// RepetitionPenalty is the repetition penalty for sampling.
	RepetitionPenalty float64 `json:"repetition_penalty"`
	// FrequencyPenalty penalizes tokens that appear frequently in the output.
	// Range is typically -2.0 to 2.0; positive values reduce repetition.
	FrequencyPenalty float64 `json:"frequency_penalty"`
	// PresencePenalty penalizes tokens that have already appeared in the output.
	// Range is typically -2.0 to 2.0; positive values encourage new topics.
	PresencePenalty float64 `json:"presence_penalty"`
	// JSONMode enables JSON output mode if supported by the model.
	JSONMode bool `json:"json_mode"`
}

// WithModel specifies the model name to use for generation.
func WithModel(model string) CallOption {
	return func(o *CallOptions) {
		o.Model = model
	}
}

// WithMaxTokens specifies the maximum number of tokens to generate.
func WithMaxTokens(maxTokens int) CallOption {
	return func(o *CallOptions) {
		o.MaxTokens = maxTokens
	}
}

// WithTemperature specifies the model temperature, a value between 0 and 1.
func WithTemperature(temperature float64) CallOption {
	return func(o *CallOptions) {
		o.Temperature = temperature
	}
}

// WithStopWords specifies a list of words to stop generation on.
func WithStopWords(stopWords []string) CallOption {
	return func(o *CallOptions) {
		o.StopWords = stopWords
	}
}

// WithStreamingFunc specifies the streaming function to use for streaming responses.
func WithStreamingFunc(streamingFunc func(ctx context.Context, chunk []byte) error) CallOption {
	return func(o *CallOptions) {
		o.StreamingFunc = streamingFunc
	}
}

// WithTopK specifies the top-k sampling parameter.
func WithTopK(topK int) CallOption {
	return func(o *CallOptions) {
		o.TopK = topK
	}
}

// WithTopP specifies the top-p sampling parameter.
func WithTopP(topP float64) CallOption {
	return func(o *CallOptions) {
		o.TopP = topP
	}
}

// WithSeed specifies the seed for deterministic sampling.
func WithSeed(seed int) CallOption {
	return func(o *CallOptions) {
		o.Seed = seed
	}
}
