package goOpenstackAuth

import (
	"fmt"
	"strings"
	"testing"
)

func resetAuthentication() {
	AuthenticationV3 = NewMockAuthenticationV3
}

func TestAuthenticationTokenSuccess(t *testing.T) {
	resetAuthentication()

	options := AuthOptions{
		IdentityEndpoint: "http://some_test_url",
		UserId:           "miau",
		Password:         "123456789",
		ProjectId:        "bup",
	}

	a := AuthenticationV3(options)
	token, err := a.GetToken()
	if err != nil {
		t.Error(fmt.Sprint(`Expected to not get an error. `, err.Error()))
		return
	}

	if !strings.Contains(token.ID, "test_token_id") {
		diffString := StringDiff(token.ID, "test_token_id")
		t.Error(fmt.Sprintf("Token does not match. \n \n %s", diffString))
	}
}

func TestAuthenticationEndpointSuccess(t *testing.T) {
	resetAuthentication()

	options := AuthOptions{
		IdentityEndpoint: "http://some_test_url",
		UserId:           "miau",
		Password:         "123456789",
		ProjectId:        "bup",
	}

	a := AuthenticationV3(options)
	endpoint, err := a.GetServiceEndpoint("arc", "staging", "public")
	if err != nil {
		t.Error(fmt.Sprint(`Expected to not get an error. `, err.Error()))
		return
	}

	if !strings.Contains(endpoint, "https://arc.staging.***REMOVED***/public") {
		diffString := StringDiff(endpoint, "https://arc.staging.***REMOVED***/public")
		t.Error(fmt.Sprintf("Endpoint does not match. \n \n %s", diffString))
	}
}

func TestAuthenticationEndpointNotGivenARegion(t *testing.T) {
	resetAuthentication()

	options := AuthOptions{
		IdentityEndpoint: "http://some_test_url",
		UserId:           "miau",
		Password:         "123456789",
		ProjectId:        "bup",
	}

	a := AuthenticationV3(options)
	endpoint, err := a.GetServiceEndpoint("arc", "", "public")
	if err != nil {
		t.Error(fmt.Sprint(`Expected to not get an error. `, err.Error()))
		return
	}

	if !strings.Contains(endpoint, "https://arc.staging.***REMOVED***/public") {
		diffString := StringDiff(endpoint, "https://arc.staging.***REMOVED***/public")
		t.Error(fmt.Sprintf("Endpoint does not match. \n \n %s", diffString))
	}
}

func TestAuthenticationGivenARegion(t *testing.T) {
	resetAuthentication()

	options := AuthOptions{
		IdentityEndpoint: "http://some_test_url",
		UserId:           "miau",
		Password:         "123456789",
		ProjectId:        "bup",
	}

	a := AuthenticationV3(options)
	endpoint, err := a.GetServiceEndpoint("arc", "production", "internal")
	if err != nil {
		t.Error(fmt.Sprint(`Expected to not get an error. `, err.Error()))
		return
	}

	if !strings.Contains(endpoint, "https://arc.production.***REMOVED***/internal") {
		diffString := StringDiff(endpoint, "https://arc.production.***REMOVED***/internal")
		t.Error(fmt.Sprintf("Endpoint does not match. \n \n %s", diffString))
	}
}

func TestAuthenticationGivenAWrongRegion(t *testing.T) {
	resetAuthentication()

	options := AuthOptions{
		IdentityEndpoint: "http://some_test_url",
		UserId:           "miau",
		Password:         "123456789",
		ProjectId:        "bup",
	}

	a := AuthenticationV3(options)
	endpoint, err := a.GetServiceEndpoint("arc", "non_exisitng_region", "internal")
	if err != nil {
		t.Error(fmt.Sprint(`Expected to not get an error. `, err.Error()))
		return
	}

	if endpoint != "" {
		t.Error("Endpoint should be empty")
	}
}

func TestAuthenticationProject(t *testing.T) {
	resetAuthentication()

	options := AuthOptions{
		IdentityEndpoint: "http://some_test_url",
		UserId:           "miau",
		Password:         "123456789",
		ProjectId:        "bup",
	}

	a := AuthenticationV3(options)
	project, err := a.GetProject()
	if err != nil {
		t.Error(fmt.Sprint(`Expected to not get an error. `, err.Error()))
		return
	}

	if project == nil {
		t.Error(`Expected to not get an empty project. `)
		return
	}

	if project.ID != "p-9597d2775" {
		diffString := StringDiff(project.ID, "p-9597d2775")
		t.Error(fmt.Sprintf("Project id does not match. \n \n %s", diffString))
	}

	if project.DomainID != "o-monsoon2" {
		diffString := StringDiff(project.DomainID, "o-monsoon2")
		t.Error(fmt.Sprintf("Project id does not match. \n \n %s", diffString))
	}

	if project.Name != "Arc_Development" {
		diffString := StringDiff(project.Name, "Arc_Development")
		t.Error(fmt.Sprintf("Project id does not match. \n \n %s", diffString))
	}
}