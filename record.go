package bluesky

import (
	"context"

	"github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/lex/util"
)

func (c *Client) GetRecord(did, collection, rkey string) (*atproto.RepoGetRecord_Output, error) {
	output, err := atproto.RepoGetRecord(context.TODO(), c.client, "", collection, did, rkey)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (c *Client) DeleteRecord(did, collection, rkey string) (*atproto.RepoDeleteRecord_Output, error) {
	output, err := atproto.RepoDeleteRecord(context.TODO(), c.client, &atproto.RepoDeleteRecord_Input{
		Collection: collection,
		Repo:       did,
		Rkey:       rkey,
	})
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (c *Client) CreateRecord(collection, repo string, record util.LexiconTypeDecoder) (*atproto.RepoCreateRecord_Output, error) {
	output, err := atproto.RepoCreateRecord(context.TODO(), c.client, &atproto.RepoCreateRecord_Input{
		Collection: collection,
		Repo:       repo,
		Record:     &record,
	})
	if err != nil {
		return nil, err
	}
	return output, nil
}
