package main

import (
	"context"

	"github.com/go-fed/activity/streams"
)

type FederateCallbacker struct {
}

func (cb *FederateCallbacker) Create(c context.Context, s *streams.Create) error {
	return nil
}
func (cb *FederateCallbacker) Update(c context.Context, s *streams.Update) error {
	return nil
}
func (cb *FederateCallbacker) Delete(c context.Context, s *streams.Delete) error {
	return nil
}
func (cb *FederateCallbacker) Add(c context.Context, s *streams.Add) error {
	return nil
}
func (cb *FederateCallbacker) Remove(c context.Context, s *streams.Remove) error {
	return nil
}
func (cb *FederateCallbacker) Like(c context.Context, s *streams.Like) error {
	return nil
}
func (cb *FederateCallbacker) Block(c context.Context, s *streams.Block) error {
	return nil
}
func (cb *FederateCallbacker) Follow(c context.Context, s *streams.Follow) error {
	return nil
}
func (cb *FederateCallbacker) Undo(c context.Context, s *streams.Undo) error {
	return nil
}
func (cb *FederateCallbacker) Accept(c context.Context, s *streams.Accept) error {
	return nil
}
func (cb *FederateCallbacker) Reject(c context.Context, s *streams.Reject) error {
	return nil
}
