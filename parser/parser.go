package parser

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func Parse(resp string) map[string]string {

	serverInfo := map[string]string{}
	tokens := strings.Split(resp, ";")
	for _, token := range tokens {
		splitToken := strings.Split(token, ":")
		k := splitToken[0]
		v := ""
		if len(splitToken) > 1 {
			v = splitToken[1]
		}
		serverInfo[k] = v
	}
	return serverInfo
}

// addr should be a host:port combined string
func QueryServer(addr string) (map[string]string, error) {
	var nilMap map[string]string
	conn, err := net.DialTimeout("tcp", addr, 8e9) // 8 secs or 8B nanoseconds
	if err != nil {
		fmt.Print(err.Error())
		return nilMap, err
	}
	reader := bufio.NewReader(conn)
	resp, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(err.Error())
		return nilMap, err
	}
	fmt.Print(resp)
	respMap := Parse(resp)
	return respMap, nil
}
