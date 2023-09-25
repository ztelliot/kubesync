package internal

import (
	"bytes"
	"encoding/json"

	"github.com/CQUPTMirror/kubesync/api/v1beta1"
)

type MirrorStatus struct {
	ID    string             `json:"id"`
	Alias string             `json:"alias"`
	Desc  string             `json:"desc"`
	Url   string             `json:"url"`
	Type  v1beta1.MirrorType `json:"type"`

	v1beta1.JobStatus
}

type MirrorConfig struct {
	ID string `json:"id"`

	v1beta1.JobSpec
}

type MirrorSchedule struct {
	NextSchedule int64 `json:"next_schedule"`
}

// A CmdVerb is an action to a job or worker
type CmdVerb uint8

const (
	// CmdStart start a job
	CmdStart CmdVerb = iota
	// CmdStop stop syncing, but keep the job
	CmdStop
	// CmdRestart restart a syncing job
	CmdRestart
	// CmdPing ensures the goroutine is alive
	CmdPing
	// CmdUpdate update size
	CmdUpdate
)

func (c CmdVerb) String() string {
	mapping := map[CmdVerb]string{
		CmdStart:   "start",
		CmdStop:    "stop",
		CmdRestart: "restart",
		CmdPing:    "ping",
	}
	return mapping[c]
}

func NewCmdVerbFromString(s string) CmdVerb {
	mapping := map[string]CmdVerb{
		"start":   CmdStart,
		"stop":    CmdStop,
		"restart": CmdRestart,
		"ping":    CmdPing,
	}
	return mapping[s]
}

// Marshal and Unmarshal for CmdVerb
func (s CmdVerb) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(s.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *CmdVerb) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*s = NewCmdVerbFromString(j)
	return nil
}

// A ClientCmd is the command message send from client
// to the manager
type ClientCmd struct {
	Cmd   CmdVerb `json:"cmd"`
	Force bool    `json:"force"`
}
