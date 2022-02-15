package main

import (
	"fmt"
	"net/http"
)

/**
 * This function handles a call for index, the top level
 * default page
 */
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index here")
}
