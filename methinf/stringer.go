package methinf // import "github.com/workspace7/tourofgo/methinf"

import (
	"strconv"
	"strings"
)

//IPAddr holds the IP address bits as array
type IPAddr [4]byte

// String prints the IP address in  string format.
func (ip IPAddr) String() string {

	strIps := make([]string, 0)

	for _, v := range ip {
		strIps = append(strIps, strconv.Itoa(int(v)))
	}
	return strings.Join(strIps, ".")
}

// func main() {
// 	hosts := map[string]IPAddr{
// 		"loopback":  {127, 0, 0, 1},
// 		"googleDNS": {8, 8, 8, 8},
// 	}
// 	for name, ip := range hosts {
// 		fmt.Printf("%v: %v\n", name, ip)
// 	}
// }
