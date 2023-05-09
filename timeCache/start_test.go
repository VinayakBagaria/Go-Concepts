// https://hackernoon.com/in-memory-caching-in-golang
package timecache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_TimeCache(t *testing.T) {
	lc := NewLocalCache(1 * time.Minute)

	testUser := User{Id: 1, FirstName: "Alice"}
	lc.Update(testUser, int(time.Now().Add(1*time.Hour).Unix()))

	// Read call
	user, err := lc.Read(testUser.Id)
	assert.NoError(t, err)
	assert.Equal(t, testUser, user)

	// Update call
	testUser.Id = 100
	lc.Update(testUser, int(time.Now().Add(1*time.Hour).Unix()))

	// Read again
	user, err = lc.Read(testUser.Id)
	assert.NoError(t, err)
	assert.Equal(t, testUser, user)

	// Delete user
	lc.Delete(testUser.Id)

	// Since deleted now, we won't be able to find it
	user, err = lc.Read(testUser.Id)
	assert.Error(t, err)
	assert.Equal(t, User{}, user)
}
