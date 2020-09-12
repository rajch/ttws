package probes

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/rajch/ttws/pkg/webserver"
)

var readinesscalled = 0
var readinesscalledtotal = 0
var readinessfailafter = 10
var readinessfailed = false
var readinesslastfailedat = 0
var readinessrecoverafter = 0

var readinessfailafterFlag = flag.String("readinessfailafter", "10", "Fail readiness probe after n calls")
var readinessrecoverafterFlag = flag.String("readinessrecoverafter", "0", "Recover readiness probe n calls after failure")

var livenesscalled int = 0
var livenesscalledtotal = 0
var livenessfailafter = 0
var livenessfailed = false
var livenesslastfailedat = 0
var livenessrecoverafter = 0

var livenessfailafterFlag = flag.String("livenessfailafter", "10", "Fail liveness probe after n calls")
var livenessrecoverafterFlag = flag.String("livenessrecoverafter", "0", "Recover liveness probe n calls after failure")

var delaystartupby int = 0

// DefaultPath is the default path to invoke the probes module - /probes
const DefaultPath = "/probes"

func probereadinesshandler(w http.ResponseWriter, r *http.Request) {
	readinesscalled++
	readinesscalledtotal++

	if readinessfailafter == 0 {
		return
	}

	if readinessfailed && readinessrecoverafter > 0 {
		if readinesscalled >= readinessrecoverafter {
			readinessfailed = false
			readinesscalled = 0

			return // Failure recovered, return success.
		}
	}

	if readinesscalled >= readinessfailafter {
		readinessfailed = true
		readinesslastfailedat = readinesscalledtotal

		readinesscalled = 0
	}

	if readinessfailed {
		http.Error(w, "Readiness probe failed.", 500)
		return // Failed, return 500.
	}

	return // Return success.
}

func probeslivenesshandler(w http.ResponseWriter, r *http.Request) {
	livenesscalled++
	livenesscalledtotal++

	if livenessfailafter == 0 {
		return
	}

	if livenessfailed && livenessrecoverafter > 0 {
		if livenesscalled >= livenessrecoverafter {
			livenessfailed = false
			livenesscalled = 0

			return // Failure recovered, return success.
		}
	}

	if livenesscalled >= livenessfailafter {
		livenessfailed = true
		livenesslastfailedat = livenesscalledtotal
		livenesscalled = 0
	}

	if livenessfailed {
		http.Error(w, "Liveness probe failed.", 500)
		return // Failed, return 500.
	}

	return // Return success.
}

func probesstatushandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "PROBE STATUS")
	fmt.Fprintln(w, "Readiness probe:")
	fmt.Fprintf(
		w,
		"  Total calls: %v\n  Fail Threshold: %v calls\n  Failed: %v\n  Last failed at: %v calls\n  Recover after: %v calls\n",
		readinesscalledtotal,
		readinessfailafter,
		readinessfailed,
		readinesslastfailedat,
		readinessrecoverafter,
	)
	fmt.Fprintln(w, "Liveness probe:")
	fmt.Fprintf(
		w,
		"  Total calls: %v\n  Fail Threshold: %v calls\n  Failed: %v\n  Last failed at: %v calls\n  Recover after: %v calls\n",
		livenesscalledtotal,
		livenessfailafter,
		livenessfailed,
		livenesslastfailedat,
		livenessrecoverafter,
	)
}

func getintoption(flag *string, envvar string, defaultstring string, defaultint int) int {
	stringvalue := webserver.GetOption(
		flag,
		envvar,
		defaultstring,
	)
	log.Printf("%v = %v", envvar, stringvalue)

	intvalue, err := strconv.Atoi(stringvalue)
	if err != nil {
		log.Printf("Int parse error")
		return defaultint
	}

	return intvalue
}

func init() {
	readinessfailafter = getintoption(
		readinessfailafterFlag,
		"READINESS_FAIL_AFTER",
		"10",
		10,
	)
	readinessrecoverafter = getintoption(
		readinessrecoverafterFlag,
		"READINESS_RECOVER_AFTER",
		"0",
		0,
	)

	livenessfailafter = getintoption(
		livenessfailafterFlag,
		"LIVENESS_FAIL_AFTER",
		"10",
		10,
	)
	livenessrecoverafter = getintoption(
		livenessrecoverafterFlag,
		"LIVENESS_RECOVER_AFTER",
		"0",
		0,
	)

	readinessPath := path.Join(DefaultPath, "/readiness")
	livenessPath := path.Join(DefaultPath, "/liveness")

	webserver.AddHandler(readinessPath, probereadinesshandler)
	webserver.AddHandler(livenessPath, probeslivenesshandler)
	webserver.AddHandler(DefaultPath, probesstatushandler)
}
