package wmic_command

import (
	"encoding/xml"
	"fmt"
	"os/exec"

	"github.com/go-zero-boilerplate/wmic-xml-command/response"
)

func Run(wmicArgsButNotFormat []string) (*response.CommandResponseXml, error) {
	args := []string{}
	args = append(args, wmicArgsButNotFormat...)
	args = append(args, "/FORMAT:RAWXML")

	out, err := exec.Command("wmic", args...).CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("WMIC command failed. Error: %s. OUTPUT: %s", err.Error(), string(out))
	}

	c := &response.CommandResponseXml{}
	if err = xml.Unmarshal(out, c); err != nil {
		parseErr := &XmlParseError{XmlContent: out, UnmarshalErr: err}
		return nil, parseErr
	}

	return c, nil
}
