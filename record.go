package bluesky

import (
	"context"

	"github.com/bluesky-social/indigo/api/atproto"
)

func (c *Client) GetRecord(did, collection, rkey string) *atproto.RepoGetRecord_Output {
	record, err := atproto.RepoGetRecord(context.TODO(), c.client, "", collection, did, rkey)
	if err != nil {
		panic(err)
	}
	return record
}
