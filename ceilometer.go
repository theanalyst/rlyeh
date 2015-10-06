package main

import (
	"github.com/rackspace/gophercloud/openstack"
	"github.com/rackspace/gophercloud/openstack/identity/v2/tokens"
)

func get_token() (*tokens.Token, error) {

	authOpts, err := openstack.AuthOptionsFromEnv()
	// authOpts := gophercloud.AuthOptions{
	// 	IdentityEndpoint: endpoint,
	// 	Username:         username,
	// 	Password:         pass,
	// 	TenantID:         tenantid,
	// }

	provider, err := openstack.AuthenticatedClient(authOpts)

	client := openstack.NewIdentityV2(provider)

	opts := tokens.AuthOptions{authOpts}

	bearer_token, err := tokens.Create(client, opts).ExtractToken()
	if err != nil {
		return nil, err
	}
	return bearer_token, nil
}
