package parser

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

// this parses the "CurrentServerTime" param into a Day, Hour string
// the server time is a number based on "ticks", where there are
// 24,000 ticks per day. So a time value of 14500 means it's 14:50 on the first day
func ParseGameTime(currentServerTime string) string {
	t, _ := strconv.Atoi(currentServerTime)
	day := int(t / 24000)
	hour := int(t % 24000 / 1000)
	return fmt.Sprintf("Day %d Hour %d", day, hour)
}

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
	time, ok := serverInfo["CurrentServerTime"]
	if ok {
		serverInfo["DayHour"] = ParseGameTime(time)
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
