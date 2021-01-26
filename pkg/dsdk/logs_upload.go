package dsdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
        "time"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	_path "path"

	uuid "github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

var (
	EndLogging = make(chan bool)
)

type LogsUpload struct {
	Path string `json:"path,omitempty"`
}

type LogsUploadRequest struct {
	Ctxt  context.Context
	Files []string
}

func newLogsUpload(path string) *LogsUpload {
	return &LogsUpload{
		Path: _path.Join(path, "logs_upload"),
	}
}

func logsUpload(ctxt context.Context, file string) error {
	conn := GetConn(ctxt)
	tid, ok := ctxt.Value("tid").(string)
	if !ok {
		tid = "nil"
	}
	reqId := uuid.Must(uuid.NewRandom()).String()
	var err error
	if conn.apikey == "" {
		if _, err = conn.Login(ctxt); err != nil {
			return err
		}
	}
	key := conn.apikey
	gurl := *conn.baseUrl
	gurl.Path = _path.Join(conn.baseUrl.Path, "logs_upload")
	url := gurl.String()

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	err = w.WriteField("ecosystem", "kubernetes")
	if err != nil {
		return err
	}

        var dir string
        dir = "logs_" + string(time.Now().Format(time.RFC3339))
        err = w.WriteField("folder_name", dir)
        if err != nil {
                return err
        }

	fw, err := w.CreateFormFile("logs.tar.gz", file)
	if err != nil {
		return err
	}
	if _, err = io.Copy(fw, f); err != nil {
		return err
	}
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest(http.MethodPut, url, &b)
	if err != nil {
		return err
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Header.Set("Auth-Token", key)

	// Submit the request
	client := &http.Client{}
	sheaders, err := json.Marshal(req.Header)
	if err != nil {
		Log().Errorf("Couldn't stringify headers, %s", req.Header)
	}
	Log().WithFields(log.Fields{
		logTraceID:        tid,
		"request_id":      reqId,
		"request_method":  http.MethodPut,
		"request_url":     gurl.String(),
		"request_headers": sheaders,
	}).Debug("Datera SDK uploading logs")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	Log().Debugf("Status Code: %d", res.StatusCode)
	// Check the response
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		Log().Error(err)
		Log().Error(string(bodyBytes))
		return err
	}
	return nil
}

func rotateLogs(rule string) error {
	if _, err := RunCmd("logrotate", "-f", rule); err != nil {
		return err
	}
	return nil
}

func (e *LogsUpload) Upload(ro *LogsUploadRequest) (*LogsUpload, *ApiErrorResponse, error) {
	return nil, nil, logsUpload(ro.Ctxt, ro.Files[0])
}

func (e *LogsUpload) RotateUploadRemove(ctxt context.Context, rule, rotated string) error {
	if err := rotateLogs(rule); err != nil {
		return err
	}

	// Determine if filtered logs exist
	lf, err := os.Open(rotated)
	if err != nil {
		return err
	}
	defer os.Remove(rotated)
	defer lf.Close()
	fstat, err := lf.Stat()
	if err != nil {
		return err
	}
	// Even a single line of logs will be greater than 100 bytes
	if fstat.Size() > 100 {
		Log().Debug("Uploading logs")
		_, apierr, err := e.Upload(&LogsUploadRequest{
			Ctxt:  ctxt,
			Files: []string{rotated},
		})
		if apierr != nil {
			Log().Errorf("%s", Pretty(apierr))
		}
		if err != nil {
			Log().Error(err)
		}
	} else {
		Log().Debugf("No new filtered logs detected.  Size: %d", fstat.Size())
	}
	return nil
}
