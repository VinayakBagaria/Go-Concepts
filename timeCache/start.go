package timecache

import (
	"fmt"
	"sync"
	"time"
)

type User struct {
	Id        int
	FirstName string
}

type CachedUser struct {
	user     User
	expireAt int
}

type LocalCache struct {
	mu    sync.RWMutex
	users map[int]CachedUser
	stop  chan struct{}
	wg    sync.WaitGroup
}

func NewLocalCache(cleanupInterval time.Duration) *LocalCache {
	lc := &LocalCache{
		users: make(map[int]CachedUser),
		stop:  make(chan struct{}),
	}

	lc.wg.Add(1)
	go func(cleanupInterval time.Duration) {
		defer lc.wg.Done()
		lc.cleanup(cleanupInterval)
	}(cleanupInterval)

	return lc
}

func (lc *LocalCache) cleanup(interval time.Duration) {
	t := time.NewTicker(interval)
	defer t.Stop()

	for {
		select {
		case <-lc.stop:
			return
		case <-t.C:
			lc.mu.Lock()
			defer lc.mu.Unlock()

			// if there are users with less expiry time than current, delete them
			for userId, cachedUser := range lc.users {
				if cachedUser.expireAt <= int(time.Now().Unix()) {
					delete(lc.users, userId)
				}
			}
		}
	}
}

func (lc *LocalCache) StopSystem() {
	close(lc.stop)
	lc.wg.Wait()
}

func (lc *LocalCache) Update(user User, expireAt int) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	lc.users[user.Id] = CachedUser{user, expireAt}
}

func (lc *LocalCache) Read(userId int) (User, error) {
	lc.mu.RLock()
	defer lc.mu.RUnlock()

	cachedUser, ok := lc.users[userId]
	if !ok {
		return User{}, fmt.Errorf("No user found")
	}

	return cachedUser.user, nil
}

func (lc *LocalCache) Delete(userId int) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	delete(lc.users, userId)
}
