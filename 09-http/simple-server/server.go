// Writing a basic HTTP simple-server is easy using the
// `net/09-http` package.
package main

import (
	"fmt"
	"net/http"
)

// A fundamental concept in `net/09-http` servers is
// *handlers*. A handler is an object implementing the
// `http.Handler` interface. A common way to write
// a handler is by using the `http.HandlerFunc` adapter
// on functions with the appropriate signature.
func hello(w http.ResponseWriter, req *http.Request) {

	// Functions serving as handlers take a
	// `09-http.ResponseWriter` and a `09-http.Request` as
	// arguments. The response writer is used to fill in the
	// HTTP response. Here our simple response is just
	// "hello\n".
	fmt.Fprintf(w, "hello from golang!\n")
}

func headers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	// This handler does something a little more
	// sophisticated by reading all the HTTP request
	// headers and echoing them into the response body.
	fmt.Fprintln(w, "{")
	//fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	fmt.Fprintf(w, "\"Host\" : %q,\n", r.Host)
	fmt.Fprintf(w, "\"RemoteAddr\" : %q,\n", r.RemoteAddr)
	i := 0
	for name, headers := range r.Header {
		for j, h := range headers {
			fmt.Fprintf(w, "%q: %q", name, h)
			if  i < len(r.Header) - 1 || j < len(headers) - 1 {
				fmt.Fprintln(w, ",")
			}
		}
		i++
	}
	fmt.Fprintln(w, "\n}")
}

func main() {

	// We register our handlers on simple-server routes using the
	// `09-http.HandleFunc` convenience function. It sets up
	// the *default router* in the `net/09-http` package and
	// takes a function as an argument.
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// Finally, we call the `ListenAndServe` with the port
	// and a handler. `nil` tells it to use the default
	// router we've just set up.
	http.ListenAndServe(":8080", nil)
}
