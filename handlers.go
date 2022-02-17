package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

/**
 * This function handles a call for index, the top level
 * default page
 */
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index here")

	/**** testing
	c := Community{
		uuid:           CreateUUID(),
		fullname:       "The Construct",
		shortname:      "const",
		date_created:   DBNow(),
		date_lastvisit: 0,
		post_count:     0,
		feature_mask:   0,
		creator:        0,
	}

	CommunityCreate(c)
	*/
	c := Community{}
	err := CommunityGet(2, &c)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Fprintln(w, err)
			return
		}
	}

	//fmt.Fprintf(w, "%s - %s", c.fullname, c.uuid)
	fmt.Fprintln(w, c)
}
