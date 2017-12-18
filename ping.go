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
	Meta      *ResultMeta
	Data      []*ResultData
	Statistic *ResultStatistic
}

//
// Struct for storing meta information in Result.Meta
//
type ResultMeta struct {
	raw   string
	Host  string
	Ip    net.IP
	Bytes string
}

//
// Struct for storing statistic information in Result.Statistic
//
type ResultStatistic struct {
	raw []string
}

//
// Struct for storing data information in Result.Data
//
type ResultData struct {
	raw string
}

//
// Ping the given net.IP once
//
func Once(ip net.IP, args ...string) (r *Result, err error) {
	r = &Result{
		Ip:        ip,
		Meta:      &ResultMeta{},
		Statistic: &ResultStatistic{},
	}

	args = append([]string{ip.String(), "-c 1"}, args...)
	err = r.execute(args)
	return
}

//
// Executes the ping command and calls r.parseRaw func
//
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

//
// Parse and convert r.Raw
//
func (r *Result) parseRaw() error {

	lines := strings.Split(string(r.Raw), "\n")

	statsBlockStart := len(lines) - 4

	for key, line := range lines {
		switch {
		case key == 0:
			splitted := strings.Split(line, " ")

			parsedIp := net.ParseIP(splitted[2][1 : len(splitted[2])-1])
			r.Meta = &ResultMeta{
				raw:   line,
				Host:  splitted[1],
				Ip:    parsedIp,
				Bytes: splitted[3],
			}
			break
		case key >= statsBlockStart:
			r.Statistic.raw = append(r.Statistic.raw, line)
			break
		case key == len(lines):
			break
		default:
			r.Data = append(r.Data, &ResultData{
				raw: line,
			})
		}
	}

	return nil
}

//
// Returns r.Meta as string
//
func (m *ResultMeta) String() string {
	return m.raw
}

//
// Returns r.Statistic as string
//
func (s *ResultStatistic) String() string {

	return strings.Join(s.raw, "\n")
}

//
// Returns r.Raw as string
//
func (r *Result) String() string {
	return string(r.Raw)
}
