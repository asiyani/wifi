package main

import (
	"fmt"
	"log"

	"github.com/asiyani/wifi"
)

func main() {
	out, err := wifi.GetInfo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", out)

	//output:-
	//{SSID:5GHz-wifi Mac:44:44:44:44:44:44 Security:wpa2-psk Channel:44,1 Frequency: SignalLevel:-70 MaxRate:300}
}
