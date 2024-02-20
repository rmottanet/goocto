package util_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"goocto/cli/pkg/util"
	"goocto/tests/pkg/mock"
)

func TestFetchData(t *testing.T) {
	srv := mock.NewTestServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test response"))
	}))
	defer srv.Close()

	client := &mock.Client{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return srv.Client().Do(req)
		},
	}

	// Test FetchData
	responseData, err := util.FetchData(srv.URL, util.RequestOptions{Client: client})
	if err != nil {
		t.Errorf("FetchData returned unexpected error: %v", err)
	}

	expectedResponse := "test response"
	if string(responseData) != expectedResponse {
		t.Errorf("FetchData returned unexpected response. Expected: %s, Got: %s", expectedResponse, string(responseData))
	}
}



func TestSubmitData(t *testing.T) {
	srv := mock.NewTestServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Expected POST request", http.StatusBadRequest)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error reading request body: %v", err), http.StatusInternalServerError)
			return
		}

		expectedBody := []byte("test data")
		if !bytes.Equal(body, expectedBody) {
			http.Error(w, fmt.Sprintf("Unexpected request body. Expected: %s, Got: %s", expectedBody, body), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test response"))
	}))
	defer srv.Close()

	client := &mock.Client{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return srv.Client().Do(req)
		},
	}

	data := []byte("test data")

	// Test SubmitData
	responseData, err := util.SubmitData(srv.URL, data, util.RequestOptions{Client: client})
	if err != nil {
		t.Errorf("SubmitData returned unexpected error: %v", err)
	}

	expectedResponse := "test response"
	if string(responseData) != expectedResponse {
		t.Errorf("SubmitData returned unexpected response. Expected: %s, Got: %s", expectedResponse, string(responseData))
	}
}


func TestUpdateData(t *testing.T) {
	srv := mock.NewTestServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Expected PUT request", http.StatusBadRequest)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error reading request body: %v", err), http.StatusInternalServerError)
			return
		}

		expectedBody := []byte("test data")
		if !bytes.Equal(body, expectedBody) {
			http.Error(w, fmt.Sprintf("Unexpected request body. Expected: %s, Got: %s", expectedBody, body), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test response"))
	}))
	defer srv.Close()

	client := &mock.Client{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return srv.Client().Do(req)
		},
	}

	data := []byte("test data")

	// Test UpdateData
	responseData, err := util.UpdateData(srv.URL, data, util.RequestOptions{Client: client})
	if err != nil {
		t.Errorf("UpdateData returned unexpected error: %v", err)
	}

	expectedResponse := "test response"
	if string(responseData) != expectedResponse {
		t.Errorf("UpdateData returned unexpected response. Expected: %s, Got: %s", expectedResponse, string(responseData))
	}
}


func TestPatchData(t *testing.T) {
	srv := mock.NewTestServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			http.Error(w, "Expected PATCH request", http.StatusBadRequest)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error reading request body: %v", err), http.StatusInternalServerError)
			return
		}

		expectedBody := []byte("test data")
		if !bytes.Equal(body, expectedBody) {
			http.Error(w, fmt.Sprintf("Unexpected request body. Expected: %s, Got: %s", expectedBody, body), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test response"))
	}))
	defer srv.Close()

	client := &mock.Client{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return srv.Client().Do(req)
		},
	}

	data := []byte("test data")

	// Test PatchData
	responseData, err := util.PatchData(srv.URL, data, util.RequestOptions{Client: client})
	if err != nil {
		t.Errorf("PatchData returned unexpected error: %v", err)
	}

	expectedResponse := "test response"
	if string(responseData) != expectedResponse {
		t.Errorf("PatchData returned unexpected response. Expected: %s, Got: %s", expectedResponse, string(responseData))
	}
}


func TestDeleteData(t *testing.T) {
	srv := mock.NewTestServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Expected DELETE request", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	client := &mock.Client{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return srv.Client().Do(req)
		},
	}

	// Test DeleteData
	err := util.DeleteData(srv.URL, util.RequestOptions{Client: client})
	if err != nil {
		t.Errorf("DeleteData returned unexpected error: %v", err)
	}
}
