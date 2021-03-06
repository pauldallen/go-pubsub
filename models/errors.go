package models

import "github.com/pkg/errors"

// topic errors
var (
	ErrAlreadyExistTopic = errors.New("already exist topic")
)

// subscription errors
var (
	ErrAlreadyExistSubscription = errors.New("already exist subscription")
	ErrNotFoundAckID            = errors.New("not found message dependent to ack id")
	ErrInvalidEndpoint          = errors.New("invalid endpoint URL format")
)

// message errors
var (
	ErrEmptyMessage       = errors.New("empty message")
	ErrNotYetReceivedAck  = errors.New("not yet received ack")
	ErrAlreadyReadMessage = errors.New("already read message")
)

// datastore errors
var (
	ErrNotFoundEntry             = errors.New("not found entry")
	ErrInvalidEntry              = errors.New("invalid entry")
	ErrNotMatchTypeMessage       = errors.New("not match type message")
	ErrNotMatchTypeMessageStatus = errors.New("not match type message status")
	ErrNotMatchTypeSubscription  = errors.New("not match type subscription")
	ErrNotMatchTypeTopic         = errors.New("not match type topic")
	ErrNotSupportOperation       = errors.New("not support operation")
	ErrNotSupportDriver          = errors.New("not support driver")
)
