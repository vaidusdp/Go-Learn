package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
)

func loadClientCAs() *x509.CertPool {
	clientCAs := x509.NewCertPool()
	caCert, err := os.ReadFile("cert.pem")
	if err != nil {
		log.Fatalln("Could not load client CA:", err)
	}

	clientCAs.AppendCertsFromPEM(caCert)
	return clientCAs
}

func main() {
	http.HandleFunc("/orders", func(w http.ResponseWriter,r *http.Request) {
		logRequestDetails(r)
		fmt.Fprintf(w, "Handleing Incoming Orders")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter,r *http.Request) {
		logRequestDetails(r)
		fmt.Fprintf(w, "Handleing Users")
	})

	port := 3000

	cert := "cert.pem"
	key := "key.pem"

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		// enforcing mTLS
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs: loadClientCAs(),
	}

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: nil,
		TLSConfig: tlsConfig,
	}

	http2.ConfigureServer(server, &http2.Server{})

	fmt.Println("Server is running on port:", port)

	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Could not start server", err)
	}

	// HTTP1.1 Without TLS
	// err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	// if err != nil {
	// 	log.Fatalln("Could not start server", err)
	// }
}

func logRequestDetails(r *http.Request) {
	httpVersion := r.Proto
	fmt.Println("HTTP Version:", httpVersion)

	if r.TLS != nil {
		tlsVersion := getTLSVersionName(r.TLS.Version)
		fmt.Println("TLS Version:", tlsVersion)
	} else {
		fmt.Println("Recieved request without TLS")
	}
}

func getTLSVersionName(version uint16) string {
	switch version {
	case tls.VersionTLS10:
		return "TLS: 1.0"
	case tls.VersionTLS12:
		return "TLS: 2.0"
	case tls.VersionTLS13:
		return "TLS: 3.0"
	default:
		return "Unknown TLS Version"
	}
}