package middlewares

import (
	"fmt"
	"net/http"
)

func SecurityHeaders(next http.Handler) http.Handler {
	fmt.Println("Security Header Middleware")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Security Header Middleware Returning...")

		// Controls DNS prefetching to reduce unnecessary DNS lookups.
		w.Header().Set("X-DNS-Prefetch-Control", "off")

		// Prevents clickjacking by stopping other websites from embedding your site in iframes.
		w.Header().Set("X-Frame-Options", "DENY") // REMEMBER

		// Deprecated: Old browser XSS filter. Modern protection should use CSP.
		w.Header().Set("X-XSS-Protection", "1; mode=block")

		// Prevents browsers from MIME-sniffing responses.
		w.Header().Set("X-Content-Type-Options", "nosniff") // REMEMBER

		// Forces browsers to use HTTPS for the specified duration.
		// Only use over HTTPS. Enable subdomains/preload only when ready.
		w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")

		// Restricts which sources the browser can load resources from.
		// Helps mitigate XSS and content injection.
		w.Header().Set("Content-Security-Policy", "default-src 'self'") // REMEMBER

		// Prevents the browser from sending referrer information.
		w.Header().Set("Referrer-Policy", "no-referrer")

		// Removes framework/version information that could reveal implementation details.
		// These headers may also need to be removed at the reverse proxy.
		// w.Header().Del("X-Powered-By")

		// Fakes the Frameworks
		w.Header().Set("X-Powered-By", "Django")

		w.Header().Del("Server")

		// Prevents Adobe Flash/PDF cross-domain policy files from granting access.
		w.Header().Set("X-Permitted-Cross-Domain-Policies", "none")

		// Prevents browser caching of sensitive responses.
		// Apply selectively: this can disable useful caching for static resources.
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")

		// Prevents other origins from using this response as a resource.
		w.Header().Set("Cross-Origin-Resource-Policy", "same-origin")

		// Isolates the browsing context from cross-origin windows.
		w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")

		// Requires cross-origin resources to explicitly grant permission.
		// Can break third-party scripts, images and other resources.
		// Enable only if your application needs cross-origin isolation.
		w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")

		// // CORS: Specifies which request headers browsers may send.
		// // Configure according to your API requirements.
		// w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// // CORS: Specifies which HTTP methods browsers may use.
		// w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")

		// // CORS: Allows credentialed cross-origin requests.
		// // Requires an explicit trusted origin in Access-Control-Allow-Origin.
		// w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Restricts browser features such as geolocation and microphone.
		w.Header().Set("Permissions-Policy", "geolocation=(self), microphone=()")

		next.ServeHTTP(w, r)
		fmt.Println("Security Header Middleware Done")
	})
}

// Skeleton
// func securityHeaders(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// next.ServeHTTP(w, r)

// 	})
// }
