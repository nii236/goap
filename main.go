package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"

	"github.com/go-fed/activity/pub"
)

type Clock struct {
}

func (c *Clock) Now() time.Time {
	return time.Now()
}

func main() {
	maxDeliveryDepth := 1
	maxForwardingDepth := 1
	clock := &Clock{}
	app := &CombinedApplication{}
	p := pub.NewPubber(
		clock,
		app,
		&SocialCallbacker{},
		&FederateCallbacker{},
		&Deliverer{},
		&HttpClient{},
		"Goap",
		maxDeliveryDepth,
		maxForwardingDepth,
	)

	serveObjectHandler := pub.ServeActivityPubObject(app, clock)
	r := chi.NewRouter()
	c := &Controller{p, serveObjectHandler}
	r.Get("/get_inbox", c.GetInbox)
	r.Get("/get_outbox", c.GetOutbox)
	r.Get("/post_inbox", c.PostInbox)
	r.Get("/post_outbox", c.PostOutbox)
	r.Get("/serve_object", c.ServeObject)

	fmt.Println("Running server on :8080")
	http.ListenAndServe(":8080", r)
}

type Controller struct {
	p                  pub.Pubber
	serveObjectHandler pub.HandlerFunc
}

func (c *Controller) ServeObject(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	result, err := c.serveObjectHandler(ctx, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Println(result, err)

}
func (c *Controller) GetInbox(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	result, err := c.p.GetInbox(ctx, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Println(result, err)
}
func (c *Controller) GetOutbox(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	result, err := c.p.GetOutbox(ctx, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Println(result, err)

}
func (c *Controller) PostInbox(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	result, err := c.p.PostInbox(ctx, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Println(result, err)
}
func (c *Controller) PostOutbox(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	result, err := c.p.PostOutbox(ctx, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Println(result, err)
}
