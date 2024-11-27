// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/zawa-t/pr/src/platform/http"
	"sync"
)

// Ensure, that ClientMock does implement http.Client.
// If this is not the case, regenerate this file with moq.
var _ http.Client = &ClientMock{}

// ClientMock is a mock implementation of http.Client.
//
//	func TestSomethingThatUsesClient(t *testing.T) {
//
//		// make and configure a mocked http.Client
//		mockedClient := &ClientMock{
//			SendFunc: func(ctx context.Context, req *http.Request) (*http.Response, error) {
//				panic("mock out the Send method")
//			},
//		}
//
//		// use mockedClient in code that requires http.Client
//		// and then make assertions.
//
//	}
type ClientMock struct {
	// SendFunc mocks the Send method.
	SendFunc func(ctx context.Context, req *http.Request) (*http.Response, error)

	// calls tracks calls to the methods.
	calls struct {
		// Send holds details about calls to the Send method.
		Send []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Req is the req argument value.
			Req *http.Request
		}
	}
	lockSend sync.RWMutex
}

// Send calls SendFunc.
func (mock *ClientMock) Send(ctx context.Context, req *http.Request) (*http.Response, error) {
	if mock.SendFunc == nil {
		panic("ClientMock.SendFunc: method is nil but Client.Send was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Req *http.Request
	}{
		Ctx: ctx,
		Req: req,
	}
	mock.lockSend.Lock()
	mock.calls.Send = append(mock.calls.Send, callInfo)
	mock.lockSend.Unlock()
	return mock.SendFunc(ctx, req)
}

// SendCalls gets all the calls that were made to Send.
// Check the length with:
//
//	len(mockedClient.SendCalls())
func (mock *ClientMock) SendCalls() []struct {
	Ctx context.Context
	Req *http.Request
} {
	var calls []struct {
		Ctx context.Context
		Req *http.Request
	}
	mock.lockSend.RLock()
	calls = mock.calls.Send
	mock.lockSend.RUnlock()
	return calls
}