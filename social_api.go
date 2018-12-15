package main

import (
	"context"
	"crypto"
	"net/http"
	"net/url"

	"github.com/go-fed/httpsig"

	"github.com/go-fed/activity/pub"
)

type SocialAPI struct {
}

func (s *SocialAPI) ActorIRI(c context.Context, r *http.Request) (*url.URL, error) {
	return nil, nil
}
func (s *SocialAPI) GetSocialAPIVerifier(c context.Context) pub.SocialAPIVerifier {
	return nil
}
func (s *SocialAPI) GetPublicKeyForOutbox(c context.Context, publicKeyId string, boxIRI *url.URL) (crypto.PublicKey, httpsig.Algorithm, error) {
	return nil, "SHA256", nil
}

type SocialAPIVerifier struct {
}

func (s *SocialAPIVerifier) Verify(r *http.Request) (authenticatedUser *url.URL, authn, authz bool, err error) {
	return nil, false, false, nil
}
func (s *SocialAPIVerifier) VerifyForOutbox(r *http.Request, outbox *url.URL) (authn, authz bool, err error) {
	return false, false, nil
}
