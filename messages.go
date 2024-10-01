package proxyapi

// Register your server with the proxy
type RegisterJSON struct {
	PublicKey string `json:"publicKey"`
	Network   string `json:"network"` // either "IPv4" or "IPv6", will give you a network on the private IPv4 or IPv6 network
}

// Response to Register
type RegisterResponseJSON struct {
	ProxyPublicKey  string `json:"proxyPublicKey"`  // The proxy's public key (eg "ZO0qmRbISuPHSBIoZnC8sSDBkWrxsLxbiNXgGZIhKEE=")
	ProxyVpnIP      string `json:"proxyVpnIP"`      // The proxy's IP in the VPN (eg "10.6.0.0" or "fdce:c10b:5ca1:1::1", depending on the network choice)
	ProxyListenPort int    `json:"proxyListenPort"` // The port that the proxy's Wireguard listens on (eg 51820)
	ServerVpnIP     string `json:"serverVpnIP"`     // Your IP in the VPN (eg "10.7.3.213" or "fdce:c10b:5ca1:2::13", depending on the network choice)
}

// Remove your server from the proxy
type RemoveJSON struct {
	PublicKey string `json:"publicKey"`
}
