# ping - a go wrapper for ping command

### Example
```
import (
	"log"
	"net"

	"github.com/xellio/ping"
)

func main() {
	ip := net.ParseIP("8.8.8.8")
	res, err := ping.Once(ip)
	if err != nil {
		log.Println(err)
	}

	log.Println(string(res.Raw))
}
```