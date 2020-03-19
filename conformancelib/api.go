package conformance

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/uc-cdis/mariner/wflib"
)

// some structs matching JSON request/responses to/from mariner API

// WorkflowRequest ..
type WorkflowRequest struct {
	Workflow *wflib.WorkflowJSON
	Input    map[string]interface{}
	Tags     map[string]string
}

// AccessToken response from fence /access_token endpoint
type AccessToken struct {
	Token string `json:"access_token"`
}

// fixme: would be nice to be able to import the mariner lib safely importable
// so that you can just access the type definitions without duplicating them here
// requires a fix in mariner server/engine
// probable only issue: marinerlib always tried to read in the configmap, no matter what
// clearly this is problematic and needs to be fixed
// for now.. duplicating the type definitions here

// RunLog ..
type RunLog struct {
	Path      string           `json:"-"`
	Request   *WorkflowRequest `json:"request"`
	Main      *Log             `json:"main"`
	ByProcess map[string]*Log  `json:"byProcess"`
}

// LogJSON ..
type LogJSON struct {
	Log *RunLog `json:"log"`
}

// StatusJSON ..
type StatusJSON struct {
	Status string `json:"status"`
}

// RunIDJSON ..
type RunIDJSON struct {
	RunID string `json:"runID"`
}

// Log ..
type Log struct {
	Created        string                 `json:"created,omitempty"`
	CreatedObj     time.Time              `json:"-"`
	LastUpdated    string                 `json:"lastUpdated,omitempty"`
	LastUpdatedObj time.Time              `json:"-"`
	JobID          string                 `json:"jobID,omitempty"`
	JobName        string                 `json:"jobName,omitempty"`
	Status         string                 `json:"status"`
	Stats          *Stats                 `json:"stats"`
	Event          EventLog               `json:"eventLog,omitempty"`
	Input          map[string]interface{} `json:"input"`
	Scatter        map[int]*Log           `json:"scatter,omitempty"`
	Output         map[string]interface{} `json:"output"` // okay
}

// Stats holds performance stats for a given process
// recorded for tasks as well as workflows
// Runtime for a workflow is the sum of runtime of that workflow's steps
type Stats struct {
	CPUReq        ResourceRequirement `json:"cpuReq"` // in-progress
	MemoryReq     ResourceRequirement `json:"memReq"` // in-progress
	ResourceUsage ResourceUsage       `json:"resourceUsage"`
	Duration      float64             `json:"duration"`  // okay - currently measured in minutes
	DurationObj   time.Duration       `json:"-"`         // okay
	NFailures     int                 `json:"nfailures"` // TODO
	NRetries      int                 `json:"nretries"`  // TODO
}

type ResourceRequirement struct {
	Min int64 `json:"min"`
	Max int64 `json:"max"`
}

type ResourceUsage struct {
	Series         ResourceUsageSeries `json:"data"`
	SamplingPeriod int                 `json:"samplingPeriod"`
}

type ResourceUsageSeries []ResourceUsageSamplePoint

type ResourceUsageSamplePoint struct {
	CPU    int64 `json:"cpu"`
	Memory int64 `json:"mem"`
}

type EventLog []string

const (
	// of course, avoid hardcoding
	// could pass commons url as param
	tokenEndpoint = "https://mattgarvin1.planx-pla.net/user/credentials/api/access_token"
	runEndpt      = "https://mattgarvin1.planx-pla.net/ga4gh/wes/v1/runs"
	fstatusEndpt  = "https://mattgarvin1.planx-pla.net/ga4gh/wes/v1/runs/%v/status"
	flogsEndpt    = "https://mattgarvin1.planx-pla.net/ga4gh/wes/v1/runs/%v"
)

func token(creds string) (string, error) {
	body, err := apiKey(creds)
	if err != nil {
		return "", err
	}
	resp, err := http.Post(tokenEndpoint, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	t := &AccessToken{}
	err = json.Unmarshal(b, t)
	if err != nil {
		return "", err
	}
	return t.Token, nil
}

// add auth header, make request, return response
func (r *Runner) request(method string, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", r.Token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *Runner) status(url string) (string, error) {
	resp, err := r.request("GET", url, nil)
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	s := &StatusJSON{}
	if err = json.Unmarshal(b, s); err != nil {
		return "", err
	}
	return s.Status, nil
}

// return output JSON from test run with given runID
func (r *Runner) fetchRunLog(runID *RunIDJSON) (*RunLog, error) {
	url := fmt.Sprintf(flogsEndpt, runID.RunID)
	resp, err := r.request("GET", url, nil)
	if err != nil {
		return nil, err
	}

	j := &LogJSON{}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	/*
		fmt.Println("response body:")
		fmt.Println(string(b))
	*/

	err = json.Unmarshal(b, j)
	if err != nil {
		return nil, err
	}

	/*
		fmt.Println("main log of test run:")
		printJSON(j)
	*/

	return j.Log, nil
}

func (r *Runner) requestRun(wf *wflib.WorkflowJSON, in map[string]interface{}, tags map[string]string) (*http.Response, error) {
	b, err := wfBytes(wf, in, tags)
	if err != nil {
		return nil, err
	}

	resp, err := r.request("POST", runEndpt, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	return resp, nil
}
