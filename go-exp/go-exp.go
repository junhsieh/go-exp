/*
 * Use the following command to generate the private key and self-signed certificate.
 * openssl req -newkey rsa:2048 -nodes -subj "/C=US/ST=New York/L=Brooklyn/O=Company Name/CN=mydomain.com" -keyout mydomain.com.key -x509 -days 365 -out mydomain.com.crt
 */
package main

import (
	"fmt"
	"github.com/junhsieh/go-exp/lib"
	"net/http"
	"time"
)

const (
	HTTPS_PORT       = ":443"
	PRIVATE_KEY_FILE = "./bin/cent-exp.local.key"
	CERTIFICATE_FILE = "./bin/cent-exp.local.crt"
)

func rootHander(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Fprint(w, "Test: "+t.Format(time.RFC1123)+" test")
}

func main() {
	http.HandleFunc("/", rootHander)

	err := http.ListenAndServeTLS(HTTPS_PORT, CERTIFICATE_FILE, PRIVATE_KEY_FILE, nil)

	if err != nil {
		fmt.Printf("main(): %s\n", err)
		lib.HelpGenTLSKeys()
	}
}
