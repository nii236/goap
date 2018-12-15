package main

import (
	"context"
	"crypto"
	"net/url"

	"github.com/go-fed/httpsig"

	"github.com/go-fed/activity/vocab"

	"github.com/go-fed/activity/pub"

	"github.com/go-fed/activity/streams"
)

type FederateAPI struct {
}

func (f *FederateAPI) OnFollow(c context.Context, s *streams.Follow) pub.FollowResponse {
	return 0

}
func (f *FederateAPI) Unblocked(c context.Context, actorIRIs []*url.URL) error {
	return nil
}
func (f *FederateAPI) FilterForwarding(c context.Context, activity vocab.ActivityType, iris []*url.URL) ([]*url.URL, error) {
	return nil, nil
}
func (f *FederateAPI) NewSigner() (httpsig.Signer, error) {
	return nil, nil
}
func (f *FederateAPI) PrivateKey(boxIRI *url.URL) (privKey crypto.PrivateKey, pubKeyId string, err error) {
	return nil, "", nil
}
