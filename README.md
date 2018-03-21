# wifi
use wifi to get information about connected wireless in your go package.
Following information is available.

```go
type Info struct {
	SSID        string
	Mac         string
	Security    string
	Channel     string
	Frequency   string
	SignalLevel string
	MaxRate     string
}
```

# Install
```
go get github.com/asiyani/wifi
```

# Example

```go
	out, err := wifi.GetInfo()
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%+v\n", out)

	//output:-
	//{SSID:5GHz-wifi Mac:44:44:44:44:44:44 Security:wpa2-psk Channel:44,1 Frequency: SignalLevel:-70 MaxRate:300}
```

# Todo
* Add support for Linux
* calculate Frequency from Channel

# Licence 

MIT