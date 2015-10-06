package main

import (
	"bytes"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/rackspace/gophercloud/openstack"
	"github.com/rackspace/gophercloud/openstack/identity/v2/tokens"
	"net/http"
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

func PostPerfCeilometer(bearer_token *tokens.Token) error {

	m1 := Meter{
		Counter_name:   "osd.1.latency",
		Counter_type:   "gauge",
		Resource_id:    "1",
		Counter_unit:   "ms",
		Counter_volume: 10.665}

	meter := []Meter{m1}
	log.Debug("Meter json")

	m, err := json.Marshal(meter)
	log.Debug(string(m))
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://192.168.122.178:8777/v2/meters/osd.1.latency", bytes.NewBuffer(m))
	req.Header.Add("X-Auth-Token", bearer_token.ID)
	req.Header.Add("content-type", "application/json")
	resp, err := client.Do(req)
	log.Debug(resp)
	return err
}
