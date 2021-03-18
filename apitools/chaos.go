package apitools

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

//Middleware basic gin middleware interface
type Middleware interface {
	Handle(c *gin.Context)
}

type chaoticMiddleware struct {
	chaosStrategy   chaosStrategy
	changeFactor    int
	mu              *sync.Mutex
	chaosStrategies []chaosStrategy
	avgRT           int
}

func (m *chaoticMiddleware) avgTime() time.Duration {
	return time.Duration(m.avgRT + rand.Intn(50)*(1-rand.Intn(3)))
}

func (m *chaoticMiddleware) Handle(c *gin.Context) {
	time.Sleep(time.Millisecond * m.avgTime())

	if !strings.Contains(c.FullPath(), "check") {
		m.chaosStrategy.MakeChaos(c)

		if rand.Intn(100) < m.changeFactor {
			m.changeStrategy()
		}
	}
}

func (m *chaoticMiddleware) changeStrategy() {
	m.mu.Lock()
	m.chaosStrategy = m.chaosStrategies[rand.Int()%len(m.chaosStrategies)]
	m.mu.Unlock()
}

//NewChaoticMiddleware instanciate a chaotic middleware
func NewChaoticMiddleware(factor int, avgResponseTime int) *chaoticMiddleware {
	chaosStrategies := []chaosStrategy{
		new(noChaos),
		new(errorChaos),
		new(slowChaos),
	}

	return &chaoticMiddleware{
		avgRT:           avgResponseTime,
		chaosStrategy:   chaosStrategies[0],
		changeFactor:    factor,
		mu:              &sync.Mutex{},
		chaosStrategies: chaosStrategies,
	}
}

type chaosStrategy interface {
	MakeChaos(*gin.Context)
}

type noChaos struct{}

func (s *noChaos) MakeChaos(c *gin.Context) {
	fmt.Println("[noChaos] be safe")
}

type errorChaos struct{}

var errorChaosResponse = gin.H{"message": "internal server error"}

func (s *errorChaos) MakeChaos(c *gin.Context) {
	fmt.Println("[errorChaos] making error!")
	c.AbortWithStatusJSON(http.StatusInternalServerError, errorChaosResponse)
}

type slowChaos struct{}

func (s *slowChaos) MakeChaos(c *gin.Context) {
	fmt.Println("[slowChaos] making it slow...")
	time.Sleep(2 * time.Second)
}
