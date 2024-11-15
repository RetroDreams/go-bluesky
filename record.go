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

func (c *Client) DeleteRecord(did, collection, rkey string) *atproto.RepoDeleteRecord_Output {
	output, err := atproto.RepoDeleteRecord(context.TODO(), c.client, &atproto.RepoDeleteRecord_Input{
		Collection: collection,
		Repo:       did,
		Rkey:       rkey,
	})
	if err != nil {
		panic(err)
	}
	return output
}
