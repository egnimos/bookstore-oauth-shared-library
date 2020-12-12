package oauth

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

//create your own mockup
func TestMain(m *testing.M) {
	fmt.Println("about to start oauth test")

	// rest.StartMockupServer()
	os.Exit(m.Run())
}

//test the constant
func TestOauthConstant(t *testing.T) {

}

//test IsPublic
func TestIsPublicNilRequest(t *testing.T) {
	assert.True(t, IsPublic(nil))
}

func TestIsPublicNoError(t *testing.T) {

}
