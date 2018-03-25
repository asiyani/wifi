# wifi
use wifi to get information about connected wireless in go app.
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
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", out)

	//output:-
	//{SSID:5GHz-wifi Mac:44:44:44:44:44:44 Security:wpa2-psk Channel:44,1 Frequency:5220 SignalLevel:-70 MaxRate:300}
```



# Licence 

MIT