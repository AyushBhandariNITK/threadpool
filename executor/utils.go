package executor

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/google/uuid"
)

var UUIDNameMap sync.Map

func GenerateUUID() string {
	return uuid.New().String()
}

var adjectives = []string{
	"efficient", "scalable", "responsive", "dynamic", "parallel",
	"concurrent", "robust", "synchronized", "threadsafe", "highperformance",
}

var notableTerms = []string{
	"scheduler", "executor", "dispatcher", "worker", "queue",
	"organiser", "manager", "handler", "processor", "coordinator",
}

func GenerateName(uuid ...string) string {
	UUID := ""
	if len(uuid) > 0 {
		UUID = uuid[0]
	} else {
		UUID = GenerateUUID()
	}
	adjective := adjectives[rand.Intn(len(adjectives))]
	term := notableTerms[rand.Intn(len(notableTerms))]
	Name := fmt.Sprintf("%s_%s", adjective, term)
	UUIDNameMap.Store(UUID, Name)
	return Name
}
