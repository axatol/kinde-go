//go:build e2e
// +build e2e

package applications_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/axatol/kinde-go/api/applications"
	"github.com/axatol/kinde-go/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestE2EList(t *testing.T) {
	client := applications.New(testutil.DefaultE2EClient(t))
	res, err := client.List(context.TODO(), applications.ListParams{})
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestE2ECreateGetUpdateDelete(t *testing.T) {
	client := applications.New(testutil.DefaultE2EClient(t))
	tempID := fmt.Sprintf("test-%d", time.Now().UnixMilli())

	res, err := client.Create(context.TODO(), applications.CreateParams{Name: tempID, Type: applications.TypeRegular})
	assert.NoError(t, err)
	require.NotNil(t, res)
	require.NotEmpty(t, res.ID)

	id := res.ID

	t.Logf("created test application: %s\n", id)

	res, err = client.Get(context.TODO(), id)
	assert.NoError(t, err)
	require.NotNil(t, res)

	t.Logf("got test application: %+v\n", res)

	updateParams := applications.UpdateParams{
		Name:         tempID,
		LoginURI:     "https://example.com",
		HomepageURI:  "https://example.com",
		LogoutURIs:   []string{"https://example.com"},
		RedirectURIs: []string{"https://example.com"},
	}

	err = client.Update(context.TODO(), id, updateParams)
	assert.NoError(t, err)

	t.Logf("updated test application: %+v\n", res)

	err = client.Delete(context.TODO(), id)
	assert.NoError(t, err)

	t.Logf("deleted test application: %+v\n", res)
}
