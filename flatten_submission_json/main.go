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
		
		body := cwt.SelfText
		// clean deleted or removed content
		if body == "[deleted]" {
			continue
		}
		if body == "[removed]" {
			continue
		}

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

	header := []string{"year", "month", "day", "time", "created_utc",
						"meta_removal_type", "meta_was_deleted", 
						"all_awarding", "allow_live_comments", "approved_utc", "approved_by",
						"author", "archived", "author_flair_text", "author_fullname", "author_is_blocked", "author_patreon_flair", "author_preminum",
						"banned_utc", "banned_by", "can_mod_post", "category", "clicked", "content_category", "contest_mode",
						"discussion_type", "distinguished", "domain", "downs", "edited", "hidden", "hide_score", "id", "is_created_from_ad", "is_crosspostable", "is_meta",
						"is_original_content", "is_self", "is_video", "link_flair_template", "link_flair_text", "locked",
						"media_description", "media_html", "media_provider", "media_title", "media_type",
						"modnote_action", "modnote_created_at", "modnote_description", "modnote_details", "modnote_id", "modnote_label", "modnote_moderator", "modnote_note", "modnote_reddit_id", "modnote_subreddit", "modnote_type", "modnote_user",
						"mod_reason", "mod_reason_title", "mod_reports", "name", "no_follow", "num_comment", "num_crosspost", "over_18", "permalink", "pinned", "post_hint", "quarantine",
						"removal_reason_id", "removal_reason_message", "removal_reason_title",
						"removed_by", "removed_by_category", "retrieved_on", "saved", "score", "selftext",
						"spolier", "stickied", "subreddit", "subreddit_type", "title", "total_awards", "update_utc", "upvote_ratio", "ups", "url", "user_reports", "view_count", "whitelist_status", "lang" }
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

	for i := 0; i < 1; i++ {
		commentWg.Add(1)
		go processComment(jsonChannel, outputObjectChannel, &commentWg, &mutex, id)
	}
	for i := 0; i < 31; i++ {
		commentWg.Add(1)
		go processComment(jsonChannel, outputObjectChannel, &commentWg, &mutex, id)
	}

	commentWg.Wait()
	close(outputObjectChannel)
	wg.Wait()
}
