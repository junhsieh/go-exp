package lib

import (
	"fmt"
)

func HelpGenTLSKeys() {
	str := "openssl req -newkey rsa:2048 -nodes -subj \"/C=US/ST=New York/L=Brooklyn/O=Company Name/CN=mydomain.com\" -keyout mydomain.com.key -x509 -days 365 -out mydomain.com.crt"

	fmt.Printf("Run the following command to generate the private key and the self-signed certificate:\n" + str + "\n")
}
