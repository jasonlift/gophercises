package utils

import (
	"bytes"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

type Uri struct {
	Key string
	Val url.URL
}

func (u *Uri) String() string {
	val := u.Val.String()
	if val == "" {
		return fmt.Sprintf("%s", u.Key)
	}
	return fmt.Sprintf("%s:%s", u.Key, val)
}

func (u *Uri) Set(value string) error {
	s := strings.SplitN(value, ":", 2)
	if s[0] == "" {
		return fmt.Errorf("missing uri key in '%s'", value)
	}
	u.Key = s[0]
	if len(s) > 1 && s[1] != "" {
		e := os.ExpandEnv(s[1])
		uri, err := url.Parse(e)
		if err != nil {
			return err
		}
		u.Val = *uri
	}
	return nil
}

type Uris []Uri

func (us *Uris) String() string {
	var b bytes.Buffer
	b.WriteString("[")
	for i, u := range *us {
		if i > 0 {
			b.WriteString(" ")
		}
		b.WriteString(u.String())
	}
	b.WriteString("]")
	return b.String()
}

func (us *Uris) Set(value string) error {
	var u Uri
	if err := u.Set(value); err != nil {
		return err
	}
	*us = append(*us, u)
	return nil
}

type ProgramOptions struct {
	NodeParam string
	NodeList  []string
	Port      int
	DebugMode bool
	SinkType  string
	SinkHost  string
	SinkPort  int
}

func NewProgramOptions() *ProgramOptions {
	return &ProgramOptions{}
}

func (opt *ProgramOptions) AddFlags(fs *flag.FlagSet) {
	fs.StringVar(&opt.NodeParam, "nodeList", "localhost", "node list that Program should collect metrics from")
	fs.IntVar(&opt.Port, "port", 8082, "port used by Program")
	fs.BoolVar(&opt.DebugMode, "debug", false, "whether it is debug mode")
	fs.StringVar(&opt.SinkType, "sinkType", "influxdb", "sink type")
	fs.StringVar(&opt.SinkHost, "sinkHost", "localhost", "sink host")
	fs.IntVar(&opt.SinkPort, "sinkPort", 8086, "sink port")
}

func (opt *ProgramOptions) ParseNodeList() {
	opt.NodeList = strings.Split(opt.NodeParam, ",")
}

func IsDebug(options *ProgramOptions) bool {
	if options.DebugMode {
		return true
	} else {
		return false
	}
}
