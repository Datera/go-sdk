package dsdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os/exec"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"text/template"
	"time"

	greq "github.com/levigross/grequests"
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

var (
	src         = rand.NewSource(time.Now().UnixNano())
	execCommand = exec.Command
)

func Log() *log.Entry {
	return DecorateRuntimeContext(log.WithFields(log.Fields{}))
}

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
	m    *sync.Mutex
	data map[string]struct{}
}

func NewStringSet(size int, d ...string) *StringSet {
	result := StringSet{
		m:    &sync.Mutex{},
		data: make(map[string]struct{}, size),
	}
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
	s.m.Lock()
	defer s.m.Unlock()
	s.data[ns] = struct{}{}
}

func (s *StringSet) Delete(ns string) {
	s.m.Lock()
	defer s.m.Unlock()
	delete(s.data, ns)
}

func (s *StringSet) Union(ss *StringSet) *StringSet {
	result := StringSet{
		m:    &sync.Mutex{},
		data: make(map[string]struct{}, len(s.data)+len(ss.data)),
	}
	for k := range s.data {
		result.data[k] = struct{}{}
	}
	for k := range ss.data {
		result.data[k] = struct{}{}
	}
	return &result
}

func (s *StringSet) Intersection(ss *StringSet) *StringSet {
	result := StringSet{
		m:    &sync.Mutex{},
		data: make(map[string]struct{}, len(s.data)),
	}
	for k := range s.data {
		if _, ok := ss.data[k]; ok {
			result.data[k] = struct{}{}
		}
	}
	return &result
}

func (s *StringSet) Difference(ss *StringSet) *StringSet {
	result := StringSet{
		m:    &sync.Mutex{},
		data: make(map[string]struct{}, len(s.data)),
	}
	for k := range s.data {
		if _, ok := ss.data[k]; !ok {
			result.data[k] = struct{}{}
		}
	}
	return &result
}

func (s *StringSet) SymDifference(ss *StringSet) *StringSet {
	result := StringSet{
		m:    &sync.Mutex{},
		data: make(map[string]struct{}, len(s.data)),
	}
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

func (s *StringSet) List() []string {
	keys := make([]string, len(s.data))
	i := 0
	for k := range s.data {
		keys[i] = k
		i++
	}
	return keys
}

type IntSet struct {
	m    *sync.Mutex
	data map[int]struct{}
}

func NewIntSet(size int, d ...int) *IntSet {
	result := IntSet{
		m:    &sync.Mutex{},
		data: make(map[int]struct{}, size),
	}
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
	s.m.Lock()
	defer s.m.Unlock()
	s.data[ns] = struct{}{}
}

func (s *IntSet) Delete(ns int) {
	s.m.Lock()
	defer s.m.Unlock()
	delete(s.data, ns)
}

func (s *IntSet) Union(ss *IntSet) *IntSet {
	result := IntSet{
		m:    &sync.Mutex{},
		data: make(map[int]struct{}, len(s.data)+len(ss.data)),
	}
	for k := range s.data {
		result.data[k] = struct{}{}
	}
	for k := range ss.data {
		result.data[k] = struct{}{}
	}
	return &result
}

func (s *IntSet) Intersection(ss *IntSet) *IntSet {
	result := IntSet{
		m:    &sync.Mutex{},
		data: make(map[int]struct{}, len(s.data)),
	}
	for k := range s.data {
		if _, ok := ss.data[k]; ok {
			result.data[k] = struct{}{}
		}
	}
	return &result
}

func (s *IntSet) Difference(ss *IntSet) *IntSet {
	result := IntSet{
		m:    &sync.Mutex{},
		data: make(map[int]struct{}, len(s.data)),
	}
	for k := range s.data {
		if _, ok := ss.data[k]; !ok {
			result.data[k] = struct{}{}
		}
	}
	return &result
}

func (s *IntSet) SymDifference(ss *IntSet) *IntSet {
	result := IntSet{
		m:    &sync.Mutex{},
		data: make(map[int]struct{}, len(s.data)),
	}
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

func (s *IntSet) List() []int {
	keys := make([]int, len(s.data))
	i := 0
	for k := range s.data {
		keys[i] = k
		i++
	}
	return keys
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

type LogFormatter struct {
}

func DecorateRuntimeContext(logger *log.Entry) *log.Entry {
	if pc, file, line, ok := runtime.Caller(2); ok {
		fName := runtime.FuncForPC(pc).Name()
		return logger.WithField("file", file).WithField("line", line).WithField("func", fName)
	} else {
		return logger
	}
}

func (f *LogFormatter) Format(entry *log.Entry) ([]byte, error) {
	msg := entry.Message
	level := entry.Level
	t := entry.Time
	fields, err := json.Marshal(entry.Data)
	if err != nil {
		fmt.Printf("Error marshalling fields during logging: %s", err)
	}
	return []byte(fmt.Sprintf("%s %s %s %s\n",
		t.Format(time.RFC3339),
		strings.ToUpper(level.String()),
		string(msg),
		fields),
	), nil
}

func RunCmd(cmd ...string) (string, error) {
	ncmd := []string{}
	for _, c := range cmd {
		c = strings.TrimSpace(c)
		if c != "" {
			ncmd = append(ncmd, c)
		}
	}
	Log().Debugf("Running command: [%s]", strings.Join(ncmd, " "))
	prefix := ncmd[0]
	ncmd = ncmd[1:]
	c := execCommand(prefix, ncmd...)
	out, err := c.CombinedOutput()
	sout := string(out)
	Log().Debug(sout)
	return sout, err
}

func init() {
	log.SetFormatter(&LogFormatter{})
	log.SetLevel(log.DebugLevel)
}

func formatQueryParams(gro *greq.RequestOptions, v reflect.Value, t reflect.Type) {
	// Formats the Query Params of the Request Option to include
	// all the fields (name - value) as query params in the URL
	numFields := t.NumField()
	params := make(map[string]string)
	for i := 0; i < numFields; i++ {
		if t.Field(i).Name == "Ctxt" {
			continue
		}

		json := t.Field(i).Tag.Get("json")
		hasOmitEmpty := strings.Contains(json, "omitempty")

		key := t.Field(i).Tag.Get("mapstructure")
		ifc := fmt.Sprintf("%v", v.Field(i).Interface())
		if ifc == "" && hasOmitEmpty {
			continue
		}

		params[key] = ifc
	}
	gro.Params = params
}
