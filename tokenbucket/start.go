package tokenbucket

import (
	"math"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type TokenBucket struct {
	rate           int64
	maxTokens      int64
	currentTokens  int64
	lastRefillTime time.Time
	mutex          sync.Mutex
}

var clientBucket *TokenBucket

func NewTokenBucket(rate, maxTokens int64) *TokenBucket {
	return &TokenBucket{
		rate:           rate,
		maxTokens:      maxTokens,
		lastRefillTime: time.Now(),
		currentTokens:  maxTokens,
	}
}

func (t *TokenBucket) IsRequestAllowed(tokens int64) bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.refill()

	if t.currentTokens >= tokens {
		t.currentTokens -= tokens
		return true
	}

	return false
}

func (t *TokenBucket) refill() {
	now := time.Now()
	difference := now.Sub(t.lastRefillTime)
	tokensToBeAdded := t.rate * int64(difference.Seconds())
	t.currentTokens = int64(math.Min(float64(t.currentTokens+tokensToBeAdded), float64(t.maxTokens)))
	t.lastRefillTime = now
}

func rateLimitMiddleware(c *gin.Context) {
	if clientBucket == nil {
		clientBucket = NewTokenBucket(5, 10)
	}
	if clientBucket.IsRequestAllowed(5) {
		c.Next()
		return
	}

	c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
		"message":       "Try again after some time",
		"rate":          clientBucket.rate,
		"currentTokens": clientBucket.currentTokens,
	})
}

func DoWork() {
	router := gin.Default()
	router.GET("/hello", rateLimitMiddleware, sayHello)

	router.Run(":8000")
}

func sayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}
