package main

import (
	"context"

	"github.com/go-fed/activity/streams"
)

type SocialCallbacker struct {
}

func (cb *SocialCallbacker) Create(c context.Context, s *streams.Create) error {
	return nil
}

func (cb *SocialCallbacker) Update(c context.Context, s *streams.Update) error {
	return nil
}
func (cb *SocialCallbacker) Delete(c context.Context, s *streams.Delete) error {
	return nil
}
func (cb *SocialCallbacker) Add(c context.Context, s *streams.Add) error {
	return nil
}
func (cb *SocialCallbacker) Remove(c context.Context, s *streams.Remove) error {
	return nil
}
func (cb *SocialCallbacker) Like(c context.Context, s *streams.Like) error {
	return nil
}
func (cb *SocialCallbacker) Block(c context.Context, s *streams.Block) error {
	return nil
}
func (cb *SocialCallbacker) Follow(c context.Context, s *streams.Follow) error {
	return nil
}
func (cb *SocialCallbacker) Undo(c context.Context, s *streams.Undo) error {
	return nil
}
func (cb *SocialCallbacker) Accept(c context.Context, s *streams.Accept) error {
	return nil
}
func (cb *SocialCallbacker) Reject(c context.Context, s *streams.Reject) error {
	return nil
}
