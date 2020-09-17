package server

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/yuriizinets/googletransx"
)

// decodeTP - decode translate params
func decodeTP(body io.ReadCloser) (googletransx.TranslateParams, error) {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return googletransx.TranslateParams{}, err
	}
	var params googletransx.TranslateParams
	err = json.Unmarshal(b, &params)
	if err != nil {
		return googletransx.TranslateParams{}, err
	}
	return params, nil
}

// decodeTPS - decode translate params slice
func decodeTPS(body io.ReadCloser) ([]googletransx.TranslateParams, error) {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return []googletransx.TranslateParams{}, err
	}
	var params []googletransx.TranslateParams
	err = json.Unmarshal(b, &params)
	if err != nil {
		return []googletransx.TranslateParams{}, err
	}
	return params, nil
}

// DetectHandler detects text's language
func DetectHandler(w http.ResponseWriter, r *http.Request) {
	// Read params (only Text is needed)
	params, err := decodeTP(r.Body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Decode error"))
		return
	}
	// Detect language
	detected, err := googletransx.Detect(params.Text)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Detection error"))
		return
	}
	// Write response
	respStr, _ := json.Marshal(detected)
	w.Write([]byte(respStr))
	return
}

// TranslateHandler is a main handler for processing
func TranslateHandler(w http.ResponseWriter, r *http.Request) {
	// Read input params
	params, err := decodeTP(r.Body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Decode error"))
		return
	}
	// Translate
	result, err := googletransx.Translate(params)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	// Write response
	respStr, _ := json.Marshal(result)
	w.Write([]byte(respStr))
	return
}

// BulkTranslateHandler is a separate handler for bulk processing
func BulkTranslateHandler(w http.ResponseWriter, r *http.Request) {
	// Read input params
	params, err := decodeTPS(r.Body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Decode error"))
		return
	}
	// Translate
	result, err := googletransx.BulkTranslate(params)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	// Write response
	respStr, _ := json.Marshal(result)
	w.Write([]byte(respStr))
	return
}
