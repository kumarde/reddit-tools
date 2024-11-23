package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/abadojack/whatlanggo"
	log "github.com/sirupsen/logrus"
	"os"
	"sync"
)

var atomicCounter = NewCounter()

func processComment(jsonChannel chan string, outputObjectChannel chan []string, wg *sync.WaitGroup,
	mutex *sync.Mutex, id map[string]bool) {
	defer wg.Done()

	for jsonBlob := range jsonChannel {
		cwt := RedditComment{}
		json.Unmarshal([]byte(jsonBlob), &cwt)
		body := cwt.Body

		// clean deleted or removed content
		//if body == "[deleted]" {
		//	continue
		//}
		//if body == "[removed]" {
		//	continue
		//}

		//check for duplicate
		mutex.Lock()
		if id[cwt.Id] {
			mutex.Unlock()
			continue
		}
		id[cwt.Id] = true
		mutex.Unlock()

		// detect language of body
		langInfo := whatlanggo.Detect(body)
		langString := langInfo.Lang.String()
		// Build Output array
		outputObjectChannel <- cwt.buildOutput(langString)
	}
	fmt.Println("finished process comment")
}

func openFile(filename string) *os.File {
	f, err := os.Create(filename)
	if err != nil {
		log.Errorf("could not open input file %s for writing", filename)
	}
	return f
}

func writer(in <-chan []string, outDir string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Open file handlers
	commentFile := openFile(outDir + ".csv")

	defer commentFile.Close()
	writer := csv.NewWriter(commentFile)
	defer writer.Flush()

	header := []string{"author", "body", "subreddit", "year", "month", "day", "clock", "created_utc", "user_year", "user_month", "user_day", "user_clock",
				"author_created_utc", "permalink", "subreddit_type", "lang", "id", "parent_id", "link_id"}
	writer.Write(header)

	for object := range in {
		err := writer.Write(object)
		if err != nil {
			log.Fatal("failed to write object.")
		}
	}
}

func reader(infile string, jsonChannel chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	inFile, err := os.Open(infile)
	if err != nil {
		log.Fatal("failed to open input file for reading: ", err)
	}
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)

	// Default buffer is too small to store blobs, must increase size
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 2048*2048)

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		jsonChannel <- scanner.Text()
	}
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}
	close(jsonChannel)
}

func main() {
	infile := flag.String("input", os.Args[1] + ".json", "JSON file containing Stream file.")
	outDir := flag.String("outdir", os.Args[1], "Output directory")
	flag.Parse()

	jsonChannel := make(chan string, 100000)
	outputObjectChannel := make(chan []string, 10000)

	var mutex sync.Mutex
	var id = make(map[string]bool)

	//limiter := time.Tick(100 * time.Millisecond)

	wg := sync.WaitGroup{}
	commentWg := sync.WaitGroup{}

	wg.Add(2)
	go reader(*infile, jsonChannel, &wg)
	go writer(outputObjectChannel, *outDir, &wg)

	for i := 0; i < 31; i++ {
		commentWg.Add(1)
		go processComment(jsonChannel, outputObjectChannel, &commentWg, &mutex, id)
	}

	commentWg.Wait()
	close(outputObjectChannel)
	wg.Wait()
}
