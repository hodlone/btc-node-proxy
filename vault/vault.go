package vault

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/vault/api"
)

var (
	// Client is a pointer to the Vault connection available application-wide
	Client        *api.Client
	err           error
	vaultAddress  string = os.Getenv("VAULT_ADDR")
	vaultRoleID   string = os.Getenv("VAULT_ROLE_ID")
	vaultSecretID string = os.Getenv("VAULT_SECRET_ID")
)

// Start creates a connection to the Vault
func Start() {
	conf := api.DefaultConfig()
	Client, err = api.NewClient(conf)
	Client.SetAddress(vaultAddress)

	resp, err := Client.Logical().Write("auth/approle/login", map[string]interface{}{
		"role_id":   vaultRoleID,
		"secret_id": vaultSecretID,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Vault token: %s\n", resp.Auth.ClientToken)
}
