package probes

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/rajch/ttws/pkg/webserver"
)

// probe represents an HTTP get endpoint
type probe struct {
	name             string
	called           int
	calledtotal      int
	failafter        int
	failed           bool
	lastfailedat     int
	recoverafter     int
	failafterFlag    *string
	recoverafterFlag *string
}

func (p *probe) handler(w http.ResponseWriter, r *http.Request) {
	p.called++
	p.calledtotal++

	if p.failafter == 0 {
		return
	}

	if p.failed && p.recoverafter > 0 {
		if p.called >= p.recoverafter {
			p.failed = false
			p.called = 0

			return // Failure recovered, return success.
		}
	}

	if p.called >= p.failafter {
		p.failed = true
		p.lastfailedat = p.calledtotal

		p.called = 0
	}

	if p.failed {
		http.Error(w, p.name+" probe failed.", 500)
		return // Failed, return 500.
	}

	return // Return success.
}

func (p *probe) status() string {
	return fmt.Sprintf(
		"%s probe:\n  Total calls: %v\n  Fail Threshold: %v calls\n  Failed: %v\n  Last failed at: %v calls\n  Recover after: %v calls\n",
		p.name,
		p.calledtotal,
		p.failafter,
		p.failed,
		p.lastfailedat,
		p.recoverafter,
	)
}

func (p *probe) parseflags() {
	p.failafter = parseflag(
		p.failafterFlag,
		p.name,
		"fail",
		p.failafter,
	)
	p.recoverafter = parseflag(
		p.recoverafterFlag,
		p.name,
		"recover",
		p.recoverafter,
	)

}

// NewProbe creates a new probe.
// The probe adds a handler to the path:
// <probesdefaultpath>/<probename>
// . The probe name is converted to lowercase in the path.
// The probe handler keeps a count of calls made to it.
// Initially, it returns a 200 status code. After <failafter>
// calls, it returns a 500 status code on each subseqent call.
// After doing this <recoverafter> times, it goes back to
// returning 200 status codes, and the cycle repeats.
func NewProbe(name string, failafter int, recoverafter int) {
	newprobe := &probe{
		name:             name,
		called:           0,
		calledtotal:      0,
		failafter:        failafter,
		failed:           false,
		lastfailedat:     0,
		recoverafter:     recoverafter,
		failafterFlag:    makeflag(name, "fail", failafter, "Fail %s probe after n calls"),
		recoverafterFlag: makeflag(name, "recover", recoverafter, "Recover %s  probe n calls after failure"),
	}
	probes[name] = newprobe

	newprobepath := path.Join(DefaultPath, strings.ToLower(name))
	webserver.AddHandler(newprobepath, newprobe.handler)
}
