package wifi

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

// Info the information about connected Wi-Fi
type Info struct {
	SSID        string
	Mac         string
	Security    string
	Channel     string
	Frequency   string
	SignalLevel string
	MaxRate     string
}

// GetInfo will return information about current WiFi connected.
func GetInfo() (Info, error) {
	platform := runtime.GOOS
	if platform == "darwin" {
		return getOsxInfo(), nil
	} else if platform == "windows" {
		return getWinInfo(), nil
	} else if platform == "linux" {
		return getLinuxInfo(), nil
	}

	return Info{}, fmt.Errorf("error %q plateform is not supported", platform)
}

func getOsxInfo() Info {

	osxCmd := "/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport"
	osxArg := "-I"

	var i Info
	cmd := exec.Command(osxCmd, osxArg)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	re := regexp.MustCompile(`(\w+): (\S+)`)
	subMatch := re.FindAllStringSubmatch(out.String(), -1)

	for _, sub := range subMatch {
		switch true {
		case sub[1] == "agrCtlRSSI":
			i.SignalLevel = strings.TrimSpace(sub[2])
		case sub[1] == "SSID":
			i.SSID = strings.TrimSpace(sub[2])
		case sub[1] == "maxRate":
			i.MaxRate = strings.TrimSpace(sub[2])
		case sub[1] == "BSSID":
			i.Mac = strings.TrimSpace(sub[2])
		case sub[1] == "channel":
			i.Channel = strings.Split(strings.TrimSpace(sub[2]), ",")[0]
		case sub[1] == "auth":
			i.Security = strings.TrimSpace(sub[2])
		}
	}
	i.Frequency = mapToFreq(i.Channel)

	return i
}

func getWinInfo() Info {

	winCmd := "Netsh"

	var i Info
	cmd := exec.Command(winCmd, "WLAN", "show", "interfaces")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	re := regexp.MustCompile(`(.*?): (.*?)(\r\n|\r|\n)`)
	subMatch := re.FindAllStringSubmatch(out.String(), -1)

	for _, sub := range subMatch {
		switch true {
		case strings.TrimSpace(sub[1]) == "Signal":
			i.SignalLevel = strings.TrimSpace(sub[2])
		case strings.TrimSpace(sub[1]) == "SSID":
			i.SSID = strings.TrimSpace(sub[2])
		case strings.TrimSpace(sub[1]) == "Receive rate (Mbps)":
			i.MaxRate = strings.TrimSpace(sub[2])
		case strings.TrimSpace(sub[1]) == "BSSID":
			i.Mac = strings.TrimSpace(sub[2])
		case strings.TrimSpace(sub[1]) == "Channel":
			i.Channel = strings.TrimSpace(sub[2])
		case strings.TrimSpace(sub[1]) == "Authentication":
			i.Security = strings.TrimSpace(sub[2])
		}
	}
	i.Frequency = mapToFreq(i.Channel)

	return i
}

func getLinuxInfo() Info {
	//TODO: Need to get information for linux
	return Info{}
}

func mapToFreq(ch string) string {
	if ch == "14" {
		return "2484"
	}
	chInt, err := strconv.Atoi(ch)
	if err != nil {
		return ""
	}
	if 1 <= chInt && chInt <= 13 {
		freq := ((chInt - 1) * 5) + 2412
		return strconv.Itoa(freq)
	}
	freqMap := map[int]int{
		36: 5180, 38: 5190, 40: 5200, 42: 5210, 44: 5220, 46: 5230, 48: 5240, 50: 5250, 52: 5260, 54: 5270, 56: 5280, 58: 5290,
		60: 5300, 62: 5310, 64: 5320, 100: 5500, 102: 5510, 104: 5520, 106: 5530, 108: 5540, 110: 5550, 112: 5560, 114: 5570,
		116: 5580, 118: 5590, 120: 5600, 122: 5610, 124: 5620, 126: 5630, 128: 5640, 132: 5660, 134: 5670, 136: 5680, 138: 5690,
		140: 5700, 142: 5710, 144: 5720, 149: 5745, 151: 5755, 153: 5765, 155: 5775, 157: 5785, 159: 5795, 161: 5805, 165: 5825,
		169: 5845, 173: 5865, 183: 4915, 184: 4920, 185: 4925, 187: 4935, 188: 4940, 189: 4945, 192: 4960, 196: 4980,
	}

	freq, ok := freqMap[chInt]
	if ok {
		return strconv.Itoa(freq)
	}

	return ""
}
