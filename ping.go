package ping

import (
	"net"
	"os/exec"
	"strings"
)

//
// Result struct
//
type Result struct {
	Ip        net.IP
	Raw       []byte
	Meta      string
	Data      []string
	Statistic []string
}

//
// Ping the given net.IP once
//
func Once(ip net.IP, args ...string) (r *Result, err error) {
	r = &Result{
		Ip: ip,
	}

	args = append([]string{ip.String(), "-c 1"}, args...)
	err = r.execute(args)
	return
}

func (r *Result) execute(args []string) error {
	path, err := exec.LookPath("ping")
	if err != nil {
		return err
	}

	out, err := exec.Command(path, args...).Output()
	if err != nil {
		return err
	}
	r.Raw = out

	err = r.parseRaw()
	if err != nil {
		return err
	}

	return nil
}

func (r *Result) parseRaw() error {

	lines := strings.Split(string(r.Raw), "\n")

	statsBlockStart := len(lines) - 4

	for key, line := range lines {
		switch {
		case key == 0:
			r.Meta = line
			break
		case key >= statsBlockStart:
			r.Statistic = append(r.Statistic, line)
			break
		case key == len(lines):
			break
		default:
			r.Data = append(r.Data, line)
		}
	}

	return nil
}
