package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type WebhookHandler func(eventname string, payload *GitHubPayload, req *http.Request) error

func Handler(secret string, fn WebhookHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		event := req.Header.Get("x-github-event")
		delivery := req.Header.Get("x-github-delivery")
		signature := req.Header.Get("x-hub-signature")

		// Utility funcs
		_fail := func(err error) {
			fail(w, event, err)
		}
		_succeed := func() {
			succeed(w, event)
		}

		// Ensure headers are all there
		if event == "" || delivery == "" {
			_fail(fmt.Errorf("Missing x-github-* and x-hub-* headers"))
			return
		}

		// No secret provided to github
		if signature == "" && secret != "" {
			_fail(fmt.Errorf("GitHub isn't providing a signature, whilst a secret is being used (please give github's webhook the secret)"))
			return
		}

		// Ensure secret is a sha1 hash
		signature_parts := strings.SplitN(signature, "=", 2)
		if len(signature_parts) != 2 {
			_fail(fmt.Errorf("Signature: '%s' does not contain two parts (hash type and hash)", signature))
			return
		}
		signature_type := signature_parts[0]
		signature_hash := signature_parts[1]
		if signature_type != "sha1" {
			_fail(fmt.Errorf("Signature should be a 'sha1' hash not '%s'", signature_type))
			return
		}

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			_fail(err)
			return
		}

		// Check that payload came from github
		// skip check if empty secret provided
		if secret != "" && !IsValidPayload(secret, signature_hash, body) {
			_fail(fmt.Errorf("Payload did not come from GitHub"))
			return
		}

		// Get payload
		payload := GitHubPayload{}
		if err := json.Unmarshal(body, &payload); err != nil {
			_fail(fmt.Errorf("Could not deserialize payload"))
			return
		}

		// Do something with payload
		if err := fn(event, &payload, req); err == nil {
			_succeed()
		} else {
			_fail(err)
		}
	})
}

func succeed(w http.ResponseWriter, event string) {
	render(w, PayloadPong{
		Ok:    true,
		Event: event,
	})
}

func fail(w http.ResponseWriter, event string, err error) {
	w.WriteHeader(500)
	render(w, PayloadPong{
		Ok:    false,
		Event: event,
		Error: err.Error(),
	})
}

func render(w http.ResponseWriter, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(data)
}
