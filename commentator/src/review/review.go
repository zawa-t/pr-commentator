//go:generate moq -rm -out $GOPATH/app/src/test/mock/$GOFILE -pkg mock . Reviewer
package review

import (
	"context"
	"fmt"
)

// Reviewer ...
type Reviewer interface {
	Review(ctx context.Context, data Data) error
}

type ID string

func NewID(filePath string, lineNum uint, message Message) ID {
	return ID(fmt.Sprintf("%s:%d:%s", filePath, lineNum, message.String()))
}

func ReNewID(filePath string, lineNum uint, message string) ID {
	return ID(fmt.Sprintf("%s:%d:%s", filePath, lineNum, message))
}

type Message string

func (m Message) String() string {
	return string(m)
}

// Data ...
type Data struct {
	Name     string
	Contents []Content
}

// Content ...
type Content struct {
	ID        ID
	Linter    string
	FilePath  string
	LineNum   uint
	ColumnNum uint
	CodeLine  string
	Indicator string
	Message   Message
}

func DefaultMessage(filePath string, lineNum uint, linter string, text string) Message {
	return Message(fmt.Sprintf("[Automatic PR Comment]  \n・File: %s（%d）  \n・Linter: %s  \n・Details: %s", filePath, lineNum, linter, text)) // NOTE: 改行する際には、「空白2つ+`/n`（  \n）」が必要な点に注意
}

func CustomMessage(customText string) Message {
	return Message(fmt.Sprintf("[Automatic PR Comment]  \n%s", customText))
}
