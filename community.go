package main

import (
	"log"
)

/**
 * Insert a new community into the database from
 * the argument structure
 */
func CommunityCreate(c Community) {
	sql := "INSERT INTO community " +
		"(uuid, fullname, shortname, date_created, feature_mask, creator) VALUES " +
		"(?, ?, ?, NOW(), ?, ?)"
	st, err := database.Prepare(sql)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = st.Exec(c.uuid, c.fullname, c.shortname, c.feature_mask, c.creator)
	if err != nil {
		log.Println(err)
	}
	st.Close()
}

/**
 * Pull a community from the database and populate
 * a struct with the data
 */
func CommunityGet(uuid string, c *Community) error {
	sql := "SELECT * FROM community WHERE uuid = ? LIMIT 1"

	row := database.QueryRow(sql, uuid)
	err := row.Scan(
		&c.id,
		&c.uuid,
		&c.fullname,
		&c.shortname,
		&c.date_created,
		&c.date_lastvisit,
		&c.post_count,
		&c.feature_mask,
		&c.creator,
	)

	if err != nil {
		return err
	}

	return nil
}
