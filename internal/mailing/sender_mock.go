// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mailing

import (
	"sync"
)

// Ensure, that SenderMock does implement Sender.
// If this is not the case, regenerate this file with moq.
var _ Sender = &SenderMock{}

// SenderMock is a mock implementation of Sender.
//
// 	func TestSomethingThatUsesSender(t *testing.T) {
//
// 		// make and configure a mocked Sender
// 		mockedSender := &SenderMock{
// 			SendFunc: func(to string, subject string, html string, text string) error {
// 				panic("mock out the Send method")
// 			},
// 		}
//
// 		// use mockedSender in code that requires Sender
// 		// and then make assertions.
//
// 	}
type SenderMock struct {
	// SendFunc mocks the Send method.
	SendFunc func(to string, subject string, html string, text string) error

	// calls tracks calls to the methods.
	calls struct {
		// Send holds details about calls to the Send method.
		Send []struct {
			// To is the to argument value.
			To string
			// Subject is the subject argument value.
			Subject string
			// HTML is the html argument value.
			HTML string
			// Text is the text argument value.
			Text string
		}
	}
	lockSend sync.RWMutex
}

// Send calls SendFunc.
func (mock *SenderMock) Send(to string, subject string, html string, text string) error {
	if mock.SendFunc == nil {
		panic("SenderMock.SendFunc: method is nil but Sender.Send was just called")
	}
	callInfo := struct {
		To      string
		Subject string
		HTML    string
		Text    string
	}{
		To:      to,
		Subject: subject,
		HTML:    html,
		Text:    text,
	}
	mock.lockSend.Lock()
	mock.calls.Send = append(mock.calls.Send, callInfo)
	mock.lockSend.Unlock()
	return mock.SendFunc(to, subject, html, text)
}

// SendCalls gets all the calls that were made to Send.
// Check the length with:
//     len(mockedSender.SendCalls())
func (mock *SenderMock) SendCalls() []struct {
	To      string
	Subject string
	HTML    string
	Text    string
} {
	var calls []struct {
		To      string
		Subject string
		HTML    string
		Text    string
	}
	mock.lockSend.RLock()
	calls = mock.calls.Send
	mock.lockSend.RUnlock()
	return calls
}
