package main

import (
	"context"
	"crypto"
	"net/http"
	"net/url"

	"github.com/go-fed/activity/vocab"
	"github.com/go-fed/httpsig"

	"github.com/go-fed/activity/pub"
)

type CombinedApplication struct {
	*Application
	*SocialAPI
	*FederateAPI
}
type SocialApp struct {
	*Application
	*SocialAPI
}

type FederateApp struct {
	*Application
	*FederateAPI
}
type Application struct {
}

func (a *Application) Owns(c context.Context, id *url.URL) bool {
	return false
}
func (a *Application) Get(c context.Context, id *url.URL, rw pub.RWType) (pub.PubObject, error) {
	return nil, nil
}
func (a *Application) GetAsVerifiedUser(c context.Context, id, authdUser *url.URL, rw pub.RWType) (pub.PubObject, error) {
	return nil, nil
}
func (a *Application) Has(c context.Context, id *url.URL) (bool, error) {
	return false, nil
}
func (a *Application) Set(c context.Context, o pub.PubObject) error {
	return nil
}
func (a *Application) GetInbox(c context.Context, r *http.Request, rw pub.RWType) (vocab.OrderedCollectionType, error) {
	return nil, nil
}
func (a *Application) GetOutbox(c context.Context, r *http.Request, rw pub.RWType) (vocab.OrderedCollectionType, error) {
	return nil, nil
}
func (a *Application) NewId(c context.Context, t pub.Typer) *url.URL {
	return nil
}
func (a *Application) GetPublicKey(c context.Context, publicKeyId string) (pubKey crypto.PublicKey, algo httpsig.Algorithm, user *url.URL, err error) {
	return nil, "SHA256", nil, nil
}
func (a *Application) CanAdd(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool {
	return false
}
func (a *Application) CanRemove(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool {
	return false
}
