package idp

import (
	"math/rand"
	"os"
	"testing"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
	"github.com/moovit-sp-gmbh/hcloud-sdk-go/entities"
)

var email = os.Getenv("EMAIL")
var client = &hcloud.Config{
	Server: os.Getenv("SERVER"),
	Token:  os.Getenv("TOKEN"),
}

func setupSuite(t *testing.T) {
	if email == "" {
		t.Errorf("missing EMAIL env")
		os.Exit(1)
	}
	if client.Server == "" {
		t.Errorf("missing SERVER env")
		os.Exit(1)
	}
	if client.Token == "" {
		t.Errorf("missing TOKEN env")
		os.Exit(1)
	}
}

func TestIGetUserEmail(t *testing.T) {
	setupSuite(t)

	client := hcloud.New(client)
	i := New(client)
	userEmail, _, err := i.GetUserEmail()

	if err != nil {
		t.Error(err)
	}

	if userEmail.Email != email {
		t.Errorf("Invalid response, wanted %s got %s", email, userEmail.Email)
	}
}

func TestIGetUser(t *testing.T) {
	setupSuite(t)

	client := hcloud.New(client)
	i := New(client)
	user, _, err := i.GetUser()

	if err != nil {
		t.Error(err)
	}

	if user.Email != email {
		t.Errorf("Invalid response, wanted %s got %s", email, user.Email)
	}
}

func TestIGetUserNats(t *testing.T) {
	setupSuite(t)

	client := hcloud.New(client)
	i := New(client)
	userNats, _, err := i.GetUserNats()

	if err != nil {
		t.Error(err)
	}

	if userNats.Email != email {
		t.Errorf("Invalid response, wanted %s got %s", email, userNats.Email)
	}

	if len(userNats.Permissions.Subscribe.Allow) < 1 {
		t.Errorf("Invalid response, wanted >=1 got %d", len(userNats.Permissions.Subscribe.Allow))
	}
}

// TestIPatchUser patches the users Company
func TestIPatchUser(t *testing.T) {
	setupSuite(t)
	t.Skip("Skipping testing as backend not ready")

	client := hcloud.New(client)
	i := New(client)
	randStringBytes := randStringBytes(12)
	patchedUser, _, err := i.PatchUser(&entities.PatchUser{Company: randStringBytes})

	if err != nil {
		t.Error(err)
	}

	if patchedUser.Email != email {
		t.Errorf("Invalid response, wanted %s got %s", email, patchedUser.Email)
	}

	if patchedUser.Company != randStringBytes {
		t.Errorf("Invalid response, wanted %s got %s", randStringBytes, patchedUser.Company)
	}
}

func TestIGetUserOrgs(t *testing.T) {
	setupSuite(t)

	client := hcloud.New(client)
	i := New(client)
	userOrgs, total, _, err := i.GetUserOrganizations(1, 0)

	if err != nil {
		t.Error(err)
	}

	if total < 1 {
		t.Errorf("Invalid response, wanted >=1 got %d", total)
	}

	if len(*userOrgs) != 1 {
		t.Errorf("Invalid response, wanted 1 got %d", len(*userOrgs))
	}
}

func TestISearchUserOrgs(t *testing.T) {
	setupSuite(t)
	t.Skip("Skipping testing as backend not ready")

	client := hcloud.New(client)
	i := New(client)
	userOrgs, total, _, err := i.SearchUserOrganizations(&entities.SearchOrganizationFilter{}, 1, 0)

	if err != nil {
		t.Error(err)
	}

	if total < 1 {
		t.Errorf("Invalid response, wanted >=1 got %d", total)
	}

	if len(*userOrgs) != 1 {
		t.Errorf("Invalid response, wanted 1 got %d", len(*userOrgs))
	}
}

// randStringBytes helper function to generate random string with length n
func randStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
