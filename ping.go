package ping

import (
	"net"
	"os/exec"
)

//
// Result struct
//
type Result struct {
	Ip  net.IP
	Raw []byte
}

//
// Ping the given net.IP once
//
func Once(ip net.IP, args ...string) (*Result, error) {
	r := &Result{
		Ip: ip,
	}

	args = append([]string{ip.String(), "-c 1"}, args...)

	out, err := exec.Command("ping", args...).Output()
	if err != nil {
		return r, err
	}
	r.Raw = out

	return r, nil
}
