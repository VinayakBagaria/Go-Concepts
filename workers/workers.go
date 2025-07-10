package workers

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

type Job struct {
	filename string
	content  string
}

func generateFiles(fileCount int) map[string]string {
	files := make(map[string]string)

	words := []string{"the", "quick", "brown", "fox", "jumps", "over", "lazy", "dog",
		"pack", "my", "box", "with", "five", "dozen", "liquor", "jugs"}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 1; i <= fileCount; i++ {
		fileName := fmt.Sprintf("file_%d", i)
		var content []string

		// Generate files with varying sizes (10-50 words)
		wordCount := r.Intn(41) + 10

		for j := 0; j < wordCount; j++ {
			content = append(content, words[r.Intn(len(words))])
		}

		files[fileName] = strings.Join(words, " ")
	}

	return files
}

func processFiles(files map[string]string, numWorkers int) (int, time.Duration) {
	numFiles := len(files)
	jobs := make(chan Job, numFiles)
	results := make(chan int, numFiles)

	var wg sync.WaitGroup
	start := time.Now()

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for job := range jobs {
				// Simulate I/O
				time.Sleep(100 * time.Millisecond)
				count := len(strings.Fields(job.content))
				results <- count
			}
		}()
	}

	for filename, content := range files {
		jobs <- Job{filename, content}
	}
	close(jobs)

	// Wait for workers and close results
	go func() {
		wg.Wait()
		close(results)
	}()

	total := 0
	for count := range results {
		total += count
	}

	return total, time.Since(start)
}

func sequential(files map[string]string) (int, time.Duration) {
	start := time.Now()
	total := 0
	for _, content := range files {
		time.Sleep(100 * time.Millisecond)
		total += len(strings.Fields(content))
	}
	return total, time.Since(start)
}

func DoWork() {
	files := generateFiles(100)

	seqTotal, seqTime := sequential(files)
	fmt.Printf("Sequential: %d words in %v\n", seqTotal, seqTime)

	// Try different worker pool sizes
	workerCounts := []int{1, 5, 10, 20, 50}
	for _, workers := range workerCounts {
		total, duration := processFiles(files, workers)
		speedup := float64(seqTime) / float64(duration)
		fmt.Printf("%2d workers: %d words in %v (%.2fx faster)\n", workers, total, duration, speedup)
	}
}
