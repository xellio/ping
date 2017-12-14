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

	return nil
}
