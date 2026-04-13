// Package llms provides interfaces and types for interacting with
// large language models (LLMs) in a provider-agnostic way.
package llms

import "context"

// ContentType represents the type of content in a message part.
type ContentType string

const (
	// ContentTypeText represents plain text content.
	ContentTypeText ContentType = "text"
	// ContentTypeImageURL represents an image URL content.
	ContentTypeImageURL ContentType = "image_url"
	// ContentTypeBinary represents binary content.
	ContentTypeBinary ContentType = "binary"
)

// ContentPart represents a single part of a message content.
type ContentPart interface {
	GetType() ContentType
}

// TextContent represents a text content part.
type TextContent struct {
	Text string
}

// GetType returns the content type for TextContent.
func (t TextContent) GetType() ContentType { return ContentTypeText }

// ImageURLContent represents an image URL content part.
type ImageURLContent struct {
	URL    string
	Detail string // "low", "high", or "auto"
}

// GetType returns the content type for ImageURLContent.
func (i ImageURLContent) GetType() ContentType { return ContentTypeImageURL }

// BinaryContent represents binary (e.g. raw image bytes) content.
type BinaryContent struct {
	MIMEType string
	Data     []byte
}

// GetType returns the content type for BinaryContent.
func (b BinaryContent) GetType() ContentType { return ContentTypeBinary }

// MessageContent holds the role and content parts of a chat message.
type MessageContent struct {
	Role  ChatMessageType
	Parts []ContentPart
}

// ContentResponse is the response returned by a model for a content generation request.
type ContentResponse struct {
	// Choices holds the generated response choices.
	Choices []*ContentChoice
}

// ContentChoice represents a single generated response choice.
type ContentChoice struct {
	// Content is the generated text.
	Content string
	// StopReason is the reason the model stopped generating.
	// Common values: "stop", "length", "tool_calls", "content_filter".
	StopReason string
	// GenerationInfo contains additional metadata about the generation.
	GenerationInfo map[string]any
	// FuncCall holds function call information if the model requested a tool call.
	FuncCall *FunctionCall
	// ToolCalls holds all tool/function calls requested by the model in this choice.
	// This is preferred over FuncCall when multiple tool calls may be present.
	ToolCalls []ToolCall
}

// ToolCall represents a single tool/function call requested by the model.
// Unlike FunctionCall, this struct is designed to support multiple concurrent
// tool calls returned in a single response choice.
type ToolCall struct {
	ID        string
	Type      string
	FuncCall  *FunctionCall
}

// FunctionCall represents a function/tool call requested by the model.
type FunctionCall struct {
	Name      string
	Arguments string
}

// Model is the interface that all LLM providers must implement.
type Model interface {
	// Call generates a response for a simple text prompt.
	// Deprecated: use GenerateContent instead.
	Call(ctx context.Context, prompt string, options ...CallOption) (string, error)

	// GenerateContent generates a response for a list of message contents,
	// supporting multimodal and multi-turn conversations.
	GenerateContent(ctx context.Context, messages []MessageContent, options ...CallOption) (*ContentResponse, error)
}

// ChatMessageType represents the role of a chat message sender.
type ChatMessageT