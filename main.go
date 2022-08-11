package main

import (
	"crypto/rand"
	"fmt"
	mt "github.com/tommytim0515/go-merkletree"
	"math"
	"runtime"
	"time"
)

const (
	testSize = 1024
	trial    = 8096
)

// Binding is the binding of personal information hash and certificate information hash
type Binding struct {
	PersonInfoHash []byte `json:"person_info_hash"`
	CertInfoHash   []byte `json:"certificate"`
}

func (b *Binding) Serialize() ([]byte, error) {
	return append(b.PersonInfoHash, b.CertInfoHash...), nil
}

func genTestBindings(size int) []mt.DataBlock {
	var bindings []mt.DataBlock
	for i := 0; i < size; i++ {
		bindings = append(bindings, &Binding{
			PersonInfoHash: make([]byte, 32),
			CertInfoHash:   make([]byte, 32),
		})
		_, err := rand.Read(bindings[i].(*Binding).PersonInfoHash)
		handleError(err)
		_, err = rand.Read(bindings[i].(*Binding).CertInfoHash)
		handleError(err)
	}
	return bindings
}

func main() {
	config := &mt.Config{
		AllowDuplicates: true,
		RunInParallel:   true,
		NumRoutines:     runtime.NumCPU(),
	}
	tree := mt.NewMerkleTree(config)
	bindings := genTestBindings(testSize)
	var totalTime float64
	times := make([]float64, trial)
	for i := 0; i < trial; i++ {
		startTime := time.Now()
		err := tree.Build(bindings)
		handleError(err)
		timeInterval := time.Since(startTime).Seconds()
		totalTime += timeInterval
		times[i] = timeInterval
	}
	avgTime := totalTime / float64(trial)
	// standard deviation
	var sum float64
	for i := 0; i < trial; i++ {
		sum += (times[i] - avgTime) * (times[i] - avgTime)
	}
	std := stdDev(times, avgTime)
	fmt.Println("avg:", avgTime)
	fmt.Println("std:", std)
}

func stdDev(numbers []float64, mean float64) float64 {
	total := 0.0
	for _, number := range numbers {
		total += math.Pow(number-mean, 2)
	}
	variance := total / float64(len(numbers)-1)
	return math.Sqrt(variance)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
