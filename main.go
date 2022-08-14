package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math"
	"runtime"
	"time"

	mt "github.com/txaty/go-merkletree"
)

// Transaction is the structure used as local/global chain transaction
type Transaction struct {
	Hash []byte `json:"hash"`
}

func (t *Transaction) Serialize() ([]byte, error) {
	return t.Hash, nil
}

func genTestBindings(size int) []mt.DataBlock {
	var txList []mt.DataBlock
	for i := 0; i < size; i++ {
		randHash := make([]byte, 32)
		_, err := rand.Read(randHash)
		handleError(err)
		txList = append(txList, &Transaction{randHash})
	}
	return txList
}

func main() {
	depth := flag.Int("d", 2, "depth of the merkle tree")
	trial := flag.Int("t", 10000, "number of trials")
	parallel := flag.Bool("p", false, "number of parallel threads")
	flag.Parse()
	testSize := 1 << uint(*depth)
	fmt.Println("Test")
	fmt.Println("size:", testSize, ", depth:", *depth, ", trials:", *trial, ", parallel:", *parallel)
	config := &mt.Config{
		AllowDuplicates: true,
	}
	if *parallel {
		config.RunInParallel = true
		config.NumRoutines = runtime.NumCPU()
	}
	bindings := genTestBindings(testSize)
	//tree, err := mt.New(config, bindings)
	//handleError(err)
	var totalTime float64
	times := make([]float64, *trial)
	for i := 0; i < *trial; i++ {
		startTime := time.Now()
		_, err := mt.New(config, bindings)
		//_, err := tree.Verify(bindings[i%testSize], tree.Proofs[i%testSize])
		handleError(err)
		timeInterval := time.Since(startTime).Seconds()
		totalTime += timeInterval
		times[i] = timeInterval
	}
	avgTime := totalTime / float64(*trial)
	// standard deviation
	var sum float64
	for i := 0; i < *trial; i++ {
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
