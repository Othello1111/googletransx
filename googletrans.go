package googletransx

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"text/scanner"
	"time"

	"golang.org/x/net/html"

	"github.com/mind1949/googletrans/tk"
	"github.com/mind1949/googletrans/tkk"
	"github.com/mind1949/googletrans/transcookie"
)

const (
	defaultServiceURL = "https://translate.google.cn"
)

var (
	emptyTranlated     = Translated{}
	emptyDetected      = Detected{}
	emptyRawTranslated = rawTranslated{}

	defaultTranslator = New()
)

// Translate uses defaultTranslator to translate params.text
func Translate(params TranslateParams) (Translated, error) {
	return defaultTranslator.Translate(params)
}

// BulkTranslate uses defaultTranslator to bulk translate with goroutines
func BulkTranslate(params []TranslateParams) ([]Translated, error) {
	return defaultTranslator.BulkTranslate(params)
}

// Detect uses defaultTranslator to detect language
func Detect(text string) (Detected, error) {
	return defaultTranslator.Detect(text)
}

// Append appends serviceURLs to defaultTranslator's serviceURLs
func Append(serviceURLs ...string) {
	defaultTranslator.Append(serviceURLs...)
}

// TranslateParams represents translate params
type TranslateParams struct {
	Src      string `json:"src"`  // source language (default: auto)
	Dest     string `json:"dest"` // destination language
	Text     string `json:"text"` // text for translating
	MimeType string
}

// Translated represents translated result
type Translated struct {
	Params        TranslateParams `json:"params"`
	Text          string          `json:"text"`          // translated text
	Pronunciation string          `json:"pronunciation"` // pronunciation of translated text
}

// Detected represents language detection result
type Detected struct {
	Lang       string  `json:"lang"`       // detected language
	Confidence float64 `json:"confidence"` // the confidence of detection result (0.00 to 1.00)
}

type rawTranslated struct {
	translated struct {
		text          string
		pronunciation string
	}
	detected struct {
		originalLanguage string
		confidence       float64
	}
}

// Translator is responsible for translation
type Translator struct {
	clt         *http.Client
	serviceURLs []string
	tkkCache    tkk.Cache
}

// New initializes a Translator
func New(serviceURLs ...string) *Translator {
	var has bool
	for i := 0; i < len(serviceURLs); i++ {
		if serviceURLs[i] == defaultServiceURL {
			has = true
			break
		}
	}
	if !has {
		serviceURLs = append(serviceURLs, defaultServiceURL)
	}

	return &Translator{
		clt:         &http.Client{},
		serviceURLs: serviceURLs,
		tkkCache:    tkk.NewCache(random(serviceURLs)),
	}
}

// Translate translates text from src language to dest language
func (t *Translator) Translate(params TranslateParams) (Translated, error) {
	if params.Src == "" {
		params.Src = "auto"
	}

	if params.MimeType == "text/plain" || params.MimeType == "" {
		// Default case
		transData, err := t.do(params)
		if err != nil {
			return emptyTranlated, err
		}
		return Translated{
			Params:        params,
			Text:          transData.translated.text,
			Pronunciation: transData.translated.pronunciation,
		}, nil
	} else if params.MimeType == "text/html" { // HTML case
		// Extract texts from html
		texts := ExtractTextsFromHTML(params.Text)
		// Build parameters for bulk
		totranslate := []TranslateParams{}
		for _, t := range texts {
			totranslate = append(totranslate, TranslateParams{
				Text: t,
				Src:  t,
				Dest: params.Dest,
			})
		}
		// Translate built parameters
		translated, err := t.BulkTranslate(totranslate)
		if err != nil {
			return Translated{}, err
		}
		// Replace source HTML with translated parts
		text := params.Text
		for i := 0; i < len(totranslate); i++ {
			text = strings.ReplaceAll(text, totranslate[i].Text, translated[i].Text)
		}
		// Return
		return Translated{
			Params: params,
			Text:   text,
		}, nil
	}
	return Translated{}, errors.New("MimeType not supported")
}

// BulkTranslate translates texts to dest language with goroutines
func (t *Translator) BulkTranslate(params []TranslateParams) ([]Translated, error) {
	// Final results store
	results := []Translated{}

	// Translate with goroutines
	var wg sync.WaitGroup
	rchan := make(chan Translated, len(params))
	for _, p := range params {
		wg.Add(1)
		go func(wg *sync.WaitGroup, p TranslateParams, r chan<- Translated) {
			defer wg.Done()
			result, err := t.Translate(p)
			if err != nil {
				panic(err)
			}
			r <- result
		}(&wg, p, rchan)
	}
	wg.Wait()
	close(rchan)

	// Extract from chan
	for r := range rchan {
		results = append(results, r)
	}

	// Return
	return results, nil
}

// Detect detects text's language
func (t *Translator) Detect(text string) (Detected, error) {
	transData, err := t.do(TranslateParams{
		Src:  "auto",
		Dest: "en",
		Text: text,
	})
	if err != nil {
		return emptyDetected, err
	}
	return Detected{
		Lang:       transData.detected.originalLanguage,
		Confidence: transData.detected.confidence,
	}, nil
}

func (t *Translator) do(params TranslateParams) (rawTranslated, error) {
	req, err := t.buildTransRequest(params)
	if err != nil {
		return emptyRawTranslated, err
	}

	transService := req.URL.Scheme + "://" + req.URL.Hostname()
	var resp *http.Response
	for try := 0; try < 3; try++ {
		cookie, err := transcookie.Get(transService)
		if err != nil {
			return emptyRawTranslated, err
		}
		req.AddCookie(&cookie)
		resp, err = t.clt.Do(req)
		if err != nil {
			return emptyRawTranslated, err
		}

		if resp.StatusCode == http.StatusOK {
			break
		}

		if resp.StatusCode == http.StatusTooManyRequests {
			_, err = transcookie.Update(transService, 3*time.Second)
			if err != nil {
				return emptyRawTranslated, err
			}
		}
	}
	if resp.StatusCode != http.StatusOK {
		return emptyRawTranslated, fmt.Errorf("failed to get translation result, err: %s", resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return emptyRawTranslated, err
	}
	resp.Body.Close()

	result, err := t.parseRawTranslated(data)
	if err != nil {
		return emptyRawTranslated, err
	}

	return result, nil
}

func (t *Translator) buildTransRequest(params TranslateParams) (request *http.Request, err error) {
	tkk, err := t.tkkCache.Get()
	if err != nil {
		return nil, err
	}
	tk, _ := tk.Get(params.Text, tkk)

	u, err := url.Parse(t.randomServiceURL() + "/translate_a/single")
	if err != nil {
		return nil, err
	}

	if params.Src == "" {
		params.Src = "auto"
	}
	queries := url.Values{}
	for k, v := range map[string]string{
		"client": "webapp",
		"sl":     params.Src,
		"tl":     params.Dest,
		"hl":     params.Dest,
		"ie":     "UTF-8",
		"oe":     "UTF-8",
		"otf":    "1",
		"ssel":   "0",
		"tsel":   "0",
		"kc":     "7",
		"tk":     tk,
	} {
		queries.Add(k, v)
	}
	dts := []string{"at", "bd", "ex", "ld", "md", "qca", "rw", "rm", "ss", "t"}
	for i := 0; i < len(dts); i++ {
		queries.Add("dt", dts[i])
	}

	q := url.Values{}
	q.Add("q", params.Text)

	// If the length of the url of the get request exceeds 2000, change to a post request
	if len(u.String()+"?"+queries.Encode()+q.Encode()) >= 2000 {
		u.RawQuery = queries.Encode()
		request, err = http.NewRequest(http.MethodPost, u.String(), strings.NewReader(q.Encode()))
		if err != nil {
			return nil, err
		}
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	} else {
		queries.Add("q", params.Text)
		u.RawQuery = queries.Encode()
		request, err = http.NewRequest(http.MethodGet, u.String(), nil)
		if err != nil {
			return nil, err
		}
	}

	return request, nil
}

func (*Translator) parseRawTranslated(data []byte) (result rawTranslated, err error) {
	var s scanner.Scanner
	s.Init(bytes.NewReader(data))
	var (
		coord       = []int{-1}
		textBuilder strings.Builder
	)
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		switch tok {
		case '[':
			coord[len(coord)-1]++
			coord = append(coord, -1)
		case ']':
			coord = coord[:len(coord)-1]
		case ',':
			// no-op
		default:
			tokText := s.TokenText()
			coord[len(coord)-1]++

			if len(coord) == 4 && coord[1] == 0 && coord[3] == 0 {
				if tokText != "null" {
					textBuilder.WriteString(tokText[1 : len(tokText)-1])
				}
			}
			if len(coord) == 4 && coord[0] == 0 && coord[1] == 0 && coord[2] == 1 && coord[3] == 2 {
				if tokText != "null" {
					result.translated.pronunciation = tokText[1 : len(tokText)-1]
				}
			}
			if len(coord) == 4 && coord[0] == 0 && coord[1] == 0 && coord[3] == 2 {
				if tokText != "null" {
					result.translated.pronunciation = tokText[1 : len(tokText)-1]
				}
			}
			if len(coord) == 2 && coord[0] == 0 && coord[1] == 2 {
				result.detected.originalLanguage = tokText[1 : len(tokText)-1]
			}
			if len(coord) == 2 && coord[0] == 0 && coord[1] == 6 {
				result.detected.confidence, _ = strconv.ParseFloat(s.TokenText(), 64)
			}
		}
	}
	result.translated.text = textBuilder.String()

	return result, nil
}

// Append appends serviceURLS to  t's serviceURLs
func (t *Translator) Append(serviceURLs ...string) {
	t.serviceURLs = append(t.serviceURLs, serviceURLs...)
}

func (t *Translator) randomServiceURL() (serviceURL string) {
	return random(t.serviceURLs)
}

func ExtractTextsFromHTML(htmlsource string) []string {
	texts := []string{}
	tokenizer := html.NewTokenizer(bytes.NewBufferString(htmlsource))
	for {
		tt := tokenizer.Next()
		exit := false
		switch {
		case tt == html.ErrorToken:
			exit = true
		case tt == html.TextToken:
			texts = append(texts, string(tokenizer.Text()))
		}
		if exit {
			break
		}
	}
	return texts
}

func random(list []string) string {
	i := rand.Intn(len(list))
	return list[i]
}
