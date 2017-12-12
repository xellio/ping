package ping

import (
	"net"
	"os/exec"
)

type Result struct {
	Ip     net.IP
	Raw    []byte
	Output map[string]string
}

func Once(ip net.IP) (*Result, error) {
	r := &Result{
		Ip: ip,
	}
	out, err := exec.Command("ping", "-c 1", ip.String()).Output()
	if err != nil {
		return r, err
	}
	r.Raw = out

	return r, nil
}
