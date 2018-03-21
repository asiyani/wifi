package main

import (
	"fmt"

	"github.com/asiyani/wifi"
)

func main() {
	out, err := wifi.GetInfo()
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%+v\n", out)

	//output:-
	//{SSID:5GHz-wifi Mac:44:44:44:44:44:44 Security:wpa2-psk Channel:44,1 Frequency: SignalLevel:-70 MaxRate:300}
}
