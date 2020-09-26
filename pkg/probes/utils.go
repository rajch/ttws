package probes

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/rajch/ttws/pkg/webserver"
)

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

func makeflag(probename string, flagname string, defaultvalue int, messageformat string) *string {
	return flag.String(
		fmt.Sprintf("%s%safter", strings.ToLower(probename), flagname),
		strconv.Itoa(defaultvalue),
		fmt.Sprintf(messageformat, probename),
	)
}

func parseflag(flag *string, probename string, flagname string, defaultvalue int) int {
	return getintoption(
		flag,
		fmt.Sprintf("%s_%s_AFTER", strings.ToUpper(probename), strings.ToUpper(flagname)),
		strconv.Itoa(defaultvalue),
		defaultvalue,
	)
}
