package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
)

var (
	listenPort = flag.String("port", "8888", "HTTP Listen port")
	listenSSL  = flag.Bool("ssl", false, "FLAG: run SSL mode")
	SSLCert    = flag.String("cert", "", "SSL Cert")
	SSLKey     = flag.String("key", "", "SSL Key")
)

type Config struct {
	ListenPort string `yaml:"port" envconfig:"listen_port"`
	ListenSSL  bool   `yaml:"ssl"  envconfig:"listen_ssl"`
	SSLCert    string `yaml:"cert" envconfig:"ssl_cert"`
	SSLKey     string `yaml:"key"  envconfig:"ssl_key"`
}

func main() {
	var err error
	var config Config

	flag.Parse()

	if err = envconfig.Process("", &config); err != nil {
		err = fmt.Errorf("error parsing ENVIRONMENT configuration: %v", err)
		return
	}
	if config.ListenPort == "" {
		config.ListenPort = *listenPort
	}
	if *listenSSL {
		config.ListenSSL = true
	}
	if config.SSLKey == "" {
		config.SSLKey = *SSLKey
	}
	if config.SSLCert == "" {
		config.SSLCert = *SSLCert
	}

	http.HandleFunc("/", HTTPHandler)

	if config.ListenSSL {
		fmt.Println("Listening for SSL connection on port:", config.ListenPort, ", SSL Cert:", config.SSLCert, ", SSL Key:", config.SSLKey)
		err = http.ListenAndServeTLS(fmt.Sprintf(":%s", config.ListenPort), config.SSLCert, config.SSLKey, nil)
	} else {
		fmt.Println("Listening connections on port: ", config.ListenPort)
		err = http.ListenAndServe(fmt.Sprintf(":%s", config.ListenPort), nil)
	}

	if err != nil {
		log.Fatal("Error listening:", err)
	}

}

type RespStruct struct {
	Host       string
	Method     string
	URL        string
	RemoteAddr string
	Headers    map[string][]string
}

func HTTPHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	resp := RespStruct{
		Host:       r.Host,
		Method:     r.Method,
		URL:        r.RequestURI,
		RemoteAddr: r.RemoteAddr,
		Headers:    map[string][]string{},
	}
	for hName, hVal := range r.Header {
		for _, val := range hVal {
			resp.Headers[hName] = append(resp.Headers[hName], val)
		}
	}
	out, err := json.Marshal(resp)
	if err != nil {
		fmt.Fprintln(w, "Error generating output:", err)
		return
	}
	fmt.Fprintln(w, string(out))
}
