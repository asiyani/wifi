package wifi

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

// Info contains all the information about connected Wi-Fi
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
			i.Channel = strings.TrimSpace(sub[2])
		case sub[1] == "auth":
			i.Security = strings.TrimSpace(sub[2])
		}
	}

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

	return i
}

func getLinuxInfo() Info {
	//TODO: Need to get information for linux
	return Info{}
}
