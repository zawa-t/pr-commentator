// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/zawa-t/pr/commentator/src/platform/github"
	"sync"
)

// Ensure, that ClientMock does implement github.Client.
// If this is not the case, regenerate this file with moq.
var _ github.Client = &ClientMock{}

// ClientMock is a mock implementation of github.Client.
//
//	func TestSomethingThatUsesClient(t *testing.T) {
//
//		// make and configure a mocked github.Client
//		mockedClient := &ClientMock{
//			CreateCommentFunc: func(ctx context.Context, data github.CommentData) error {
//				panic("mock out the CreateComment method")
//			},
//			CreateReviewFunc: func(ctx context.Context, data github.ReviewData) error {
//				panic("mock out the CreateReview method")
//			},
//		}
//
//		// use mockedClient in code that requires github.Client
//		// and then make assertions.
//
//	}
type ClientMock struct {
	// CreateCommentFunc mocks the CreateComment method.
	CreateCommentFunc func(ctx context.Context, data github.CommentData) error

	// CreateReviewFunc mocks the CreateReview method.
	CreateReviewFunc func(ctx context.Context, data github.ReviewData) error

	// calls tracks calls to the methods.
	calls struct {
		// CreateComment holds details about calls to the CreateComment method.
		CreateComment []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Data is the data argument value.
			Data github.CommentData
		}
		// CreateReview holds details about calls to the CreateReview method.
		CreateReview []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Data is the data argument value.
			Data github.ReviewData
		}
	}
	lockCreateComment sync.RWMutex
	lockCreateReview  sync.RWMutex
}

// CreateComment calls CreateCommentFunc.
func (mock *ClientMock) CreateComment(ctx context.Context, data github.CommentData) error {
	if mock.CreateCommentFunc == nil {
		panic("ClientMock.CreateCommentFunc: method is nil but Client.CreateComment was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Data github.CommentData
	}{
		Ctx:  ctx,
		Data: data,
	}
	mock.lockCreateComment.Lock()
	mock.calls.CreateComment = append(mock.calls.CreateComment, callInfo)
	mock.lockCreateComment.Unlock()
	return mock.CreateCommentFunc(ctx, data)
}

// CreateCommentCalls gets all the calls that were made to CreateComment.
// Check the length with:
//
//	len(mockedClient.CreateCommentCalls())
func (mock *ClientMock) CreateCommentCalls() []struct {
	Ctx  context.Context
	Data github.CommentData
} {
	var calls []struct {
		Ctx  context.Context
		Data github.CommentData
	}
	mock.lockCreateComment.RLock()
	calls = mock.calls.CreateComment
	mock.lockCreateComment.RUnlock()
	return calls
}

// CreateReview calls CreateReviewFunc.
func (mock *ClientMock) CreateReview(ctx context.Context, data github.ReviewData) error {
	if mock.CreateReviewFunc == nil {
		panic("ClientMock.CreateReviewFunc: method is nil but Client.CreateReview was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Data github.ReviewData
	}{
		Ctx:  ctx,
		Data: data,
	}
	mock.lockCreateReview.Lock()
	mock.calls.CreateReview = append(mock.calls.CreateReview, callInfo)
	mock.lockCreateReview.Unlock()
	return mock.CreateReviewFunc(ctx, data)
}

// CreateReviewCalls gets all the calls that were made to CreateReview.
// Check the length with:
//
//	len(mockedClient.CreateReviewCalls())
func (mock *ClientMock) CreateReviewCalls() []struct {
	Ctx  context.Context
	Data github.ReviewData
} {
	var calls []struct {
		Ctx  context.Context
		Data github.ReviewData
	}
	mock.lockCreateReview.RLock()
	calls = mock.calls.CreateReview
	mock.lockCreateReview.RUnlock()
	return calls
}
