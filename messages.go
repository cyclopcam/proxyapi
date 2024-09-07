package proxyapi

// Register your server with the proxy
type RegisterJSON struct {
	PublicKey string `json:"publicKey"`
}

// Response to Register
type RegisterResponseJSON struct {
	ProxyPublicKey  string `json:"proxyPublicKey"`  // The proxy's public key (eg "ZO0qmRbISuPHSBIoZnC8sSDBkWrxsLxbiNXgGZIhKEE=")
	ProxyVpnIP      string `json:"proxyVpnIP"`      // The proxy's IP in the VPN (eg "10.6.0.0")
	ProxyListenPort int    `json:"proxyListenPort"` // The port that the proxy's Wireguard listens on (eg 51820)
	ServerVpnIP     string `json:"serverVpnIP"`     // Your IP in the VPN (eg "10.7.3.213")
}

// Remove your server from the proxy
type RemoveJSON struct {
	PublicKey string `json:"publicKey"`
}
