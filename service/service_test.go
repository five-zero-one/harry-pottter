package service_test

import (
	"encoding/json"
	hp "harry-potter"
	"net/http"
	"net/http/httptest"
	"testing"

	"harry-potter/service"

	"github.com/hyphengolang/prelude/testing/is"
)

func TestService(t *testing.T) {
	is := is.New(t)

	h := service.New()
	srv := httptest.NewServer(h)

	t.Cleanup(func() { srv.Close() })

	t.Run("get all characters", func(t *testing.T) {
		// t.Skip()

		resp, err := srv.Client().Get(srv.URL + "/characters")
		is.NoErr(err)                            // create response
		is.Equal(resp.StatusCode, http.StatusOK) // get response
		defer resp.Body.Close()

		var cs []hp.Character
		err = json.NewDecoder(resp.Body).Decode(&cs)
		is.NoErr(err)         // decoding the body
		is.Equal(len(cs), 25) // length is 25
	})

	t.Run("get a character by key", func(t *testing.T) {
		// t.Skip()

		resp, err := srv.Client().Get(srv.URL + "/characters/110")
		is.NoErr(err)                            // create response
		is.Equal(resp.StatusCode, http.StatusOK) // get response
		defer resp.Body.Close()

		var c hp.Character
		err = json.NewDecoder(resp.Body).Decode(&c)
		is.NoErr(err)                 // decoding the body
		is.Equal(c.Name, "Cho Chang") // Character 110 has name Cho Chang
	})

	t.Run("query parameters for blood type", func(t *testing.T) {
		// t.Skip()

		resp, err := srv.Client().Get(srv.URL + "/characters?blood_type=half-blood")
		is.NoErr(err)                            // create response
		is.Equal(resp.StatusCode, http.StatusOK) // get response
		defer resp.Body.Close()

		var cs []hp.Character
		err = json.NewDecoder(resp.Body).Decode(&cs)
		is.NoErr(err)         // decoding the body
		is.Equal(len(cs), 11) // length is for "Half-blood"
	})

	t.Run("query parameters for birth month", func(t *testing.T) {
		// t.Skip()

		resp, err := srv.Client().Get(srv.URL + "/characters?birth_month=february")
		is.NoErr(err)                            // create response
		is.Equal(resp.StatusCode, http.StatusOK) // get response
		defer resp.Body.Close()

		var cs []hp.Character
		err = json.NewDecoder(resp.Body).Decode(&cs)
		is.NoErr(err)        // decoding the body
		is.Equal(len(cs), 2) // length is for "Feb"
	})

	t.Run("query parameters for name search", func(t *testing.T) {
		// t.Skip()

		resp, err := srv.Client().Get(srv.URL + "/characters?name=neville")
		is.NoErr(err)                            // create response
		is.Equal(resp.StatusCode, http.StatusOK) // get response
		defer resp.Body.Close()

		var cs []hp.Character
		err = json.NewDecoder(resp.Body).Decode(&cs)
		is.NoErr(err)        // decoding the body
		is.Equal(len(cs), 1) // length is for "Neville Longbottom"
	})

	t.Run("query parameters for blood type & birth month", func(t *testing.T) {
		// t.Skip()

		resp, err := srv.Client().Get(srv.URL + "/characters?blood_type=half-blood&birth_month=sep")
		is.NoErr(err)                            // create response
		is.Equal(resp.StatusCode, http.StatusOK) // get response
		defer resp.Body.Close()

		var cs []hp.Character
		err = json.NewDecoder(resp.Body).Decode(&cs)
		is.NoErr(err)        // decoding the body
		is.Equal(len(cs), 4) // length is for "Half-blood & Sep"
	})
}
