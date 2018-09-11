package dsdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"text/template"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

// No vowels so no accidental profanity :P
const letterBytes = "bcdfghjklmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

// Args have the form "name=value"
func parseTemplate(fstring string, args ...interface{}) (string, error) {
	tpl, err := template.New("format").Parse(fstring)
	if err != nil {
		return "", err
	}
	argm := make(map[string]string)
	switch t := args[0].(type) {
	default:
		fmt.Println("Error")
	case string:
		for _, i := range args {
			arg := i.(string)
			x := strings.Split(arg, "=")
			if len(x) == 2 {
				argm[x[0]] = x[1]
			}
		}
	case map[string]string:
		argm = t
	}
	for k := range argm {
		if !strings.Contains(fstring, "{{."+k+"}}") {
			err := fmt.Errorf("Could not find arg: '%s' in template: '%s'", fstring, k)
			return "", err
		}
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, argm)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// SETS

type StringSet struct {
	data map[string]struct{}
}

func newStringSet(d ...string) *StringSet {
	result := StringSet{data: map[string]struct{}{}}
	for _, i := range d {
		result.data[i] = struct{}{}
	}
	return &result
}

func (s *StringSet) Contains(ns string) bool {
	if _, ok := s.data[ns]; !ok {
		return false
	}
	return true
}

func (s *StringSet) Add(ns string) {
	s.data[ns] = struct{}{}
}

func (s *StringSet) Delete(ns string) {
	delete(s.data, ns)
}

func (s *StringSet) Union(ss *StringSet) *StringSet {
	result := *s
	for k := range ss.data {
		result.data[k] = struct{}{}
	}
	return &result
}

func (s *StringSet) Stringersection(ss *StringSet) *StringSet {
	result := StringSet{}
	for k := range s.data {
		if _, ok := ss.data[k]; ok {
			result.data[k] = struct{}{}
		}
	}
	return &result
}

func (s *StringSet) Difference(ss *StringSet) *StringSet {
	result := *s
	for k := range s.data {
		if _, ok := ss.data[k]; !ok {
			result.data[k] = struct{}{}
		}
	}
	return &result
}

func (s *StringSet) SymDifference(ss *StringSet) *StringSet {
	result := *s
	for k := range s.data {
		if _, ok := ss.data[k]; !ok {
			result.data[k] = struct{}{}
		}
	}
	for k := range ss.data {
		if _, ok := s.data[k]; !ok {
			result.data[k] = struct{}{}
		}
	}
	return &result
}

type IntSet struct {
	data map[int]struct{}
}

func newIntSet(d ...int) *IntSet {
	result := IntSet{data: map[int]struct{}{}}
	for _, i := range d {
		result.data[i] = struct{}{}
	}
	return &result
}

func (s *IntSet) Contains(ns int) bool {
	if _, ok := s.data[ns]; !ok {
		return false
	}
	return true
}

func (s *IntSet) Add(ns int) {
	s.data[ns] = struct{}{}
}

func (s *IntSet) Delete(ns int) {
	delete(s.data, ns)
}

func (s *IntSet) Union(ss *IntSet) *IntSet {
	result := *s
	for k := range ss.data {
		result.data[k] = struct{}{}
	}
	return &result
}

func (s *IntSet) Intersection(ss *IntSet) *IntSet {
	result := IntSet{}
	for k := range s.data {
		if _, ok := ss.data[k]; ok {
			result.data[k] = struct{}{}
		}
	}
	return &result
}

func (s *IntSet) Difference(ss *IntSet) *IntSet {
	result := *s
	for k := range s.data {
		if _, ok := ss.data[k]; !ok {
			result.data[k] = struct{}{}
		}
	}
	return &result
}

func (s *IntSet) SymDifference(ss *IntSet) *IntSet {
	result := *s
	for k := range s.data {
		if _, ok := ss.data[k]; !ok {
			result.data[k] = struct{}{}
		}
	}
	for k := range ss.data {
		if _, ok := s.data[k]; !ok {
			result.data[k] = struct{}{}
		}
	}
	return &result
}

// From https://stackoverflow.com/a/31832326/4408885
func RandString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func FillStruct(m map[string]interface{}, s interface{}) error {
	return mapstructure.Decode(m, s)
}

type LogFormatter struct {
}

func (f *LogFormatter) Format(entry *log.Entry) ([]byte, error) {
	msg := entry.Message
	level := entry.Level
	t := entry.Time
	return []byte(fmt.Sprintf("%s %s %s", t.Format(time.RFC3339), strings.ToUpper(level.String()), string(msg))), nil
}

func GetConn(ctxt context.Context) *ApiConnection {
	defer recoverConn()
	conn := ctxt.Value("conn")
	return conn.(*ApiConnection)
}

func recoverConn() {
	if r := recover(); r != nil {
		panic("You MUST provide a context object containing a *ApiConnection for requests." +
			"Use sdk.Context() to obtain the context object")
	}
}

func Pretty(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func init() {
	log.SetFormatter(&LogFormatter{})
	log.SetLevel(log.DebugLevel)
}
