package main

import (
	"fmt"
	"github.com/teamnsrg/go-perspectiveapi/perspective"
	"strconv"
  "time"
)

/*
{
  "archived": false,
  "author": "TistedLogic",
  "author_created_utc": 1312615878,
  "author_flair_background_color": null,
  "author_flair_css_class": null,
  "author_flair_richtext": [],
  "author_flair_template_id": null,
  "author_flair_text": null,
  "author_flair_text_color": null,
  "author_flair_type": "text",
  "author_fullname": "t2_5mk6v",
  "author_patreon_flair": false,
  "body": "Is it still r/BoneAppleTea worthy if it's the opposite?",
  "can_gild": true,
  "can_mod_post": false,
  "collapsed": false,
  "collapsed_reason": null,
  "controversiality": 0,
  "created_utc": 1538352000,
  "distinguished": null,
  "edited": false,
  "gilded": 0,
  "gildings": {
    "gid_1": 0,
    "gid_2": 0,
    "gid_3": 0
  },
  "id": "e6xucdd",
  "is_submitter": false,
  "link_id": "t3_9ka1hp",
  "no_follow": true,
  "parent_id": "t1_e6xu13x",
  "permalink": "/r/Unexpected/comments/9ka1hp/jesus_fking_woah/e6xucdd/",
  "removal_reason": null,
  "retrieved_on": 1539714091,
  "score": 2,
  "send_replies": true,
  "stickied": false,
  "subreddit": "Unexpected",
  "subreddit_id": "t5_2w67q",
  "subreddit_name_prefixed": "r/Unexpected",
  "subreddit_type": "public"
}


*/

//author,body,subreddit,created_utc,is_submitter,permalink,subreddit_type,removal_reason,language,attack_on_author,
// identity_attack,insult,profanity,severe_toxicity,sexually_explicity,threat,toxicity

type RedditComment struct {
	Archived         bool   `json:"archived"`
	Author           string `json:"author"`
	AuthorCreatedUTC uint64 `json:"author_created_utc"`
	Body             string `json:"body"`
	CreatedUTC       uint64 `json:"created_utc"`
	Permalink        string `json:"permalink"`
	Subreddit        string `json:"subreddit"`
	SubredditType    string `json:"subreddit_type"`
	Id               string `json:"id"`
	ParentId         string `json:"parent_id"`
	LinkId           string `json:"link_id"`
}

type CommentWithToxicity struct {
	Comment *RedditComment `json:"comment"`
	// Toxicity *perspective.AnalyzeCommentResponse `json:"toxicity"`
}

func getToxicScoreFromAttribute(t *perspective.AnalyzeCommentResponse, attribute string) string {
	attributeScores := t.AttributeScores
	if attributeScores == nil {
		return "-1.0"
	}
	spanScores := *attributeScores[attribute].SpanScores
	return fmt.Sprintf("%f", spanScores[0].Score.Value)
}

func (cwt *RedditComment) buildOutput(lang string) []string {
	out := []string{}
	out = append(out, cwt.Author)
	out = append(out, cwt.Body)
	out = append(out, cwt.Subreddit)

  tm := time.Unix(int64(cwt.CreatedUTC), 0)

	out = append(out, strconv.FormatInt(int64(tm.Year()), 10))
  out = append(out, strconv.FormatInt(int64(tm.Month()), 10))
  out = append(out, strconv.FormatInt(int64(tm.Day()), 10))
  out = append(out, tm.Format("15:04:05"))

  createdStr := strconv.FormatUint(cwt.CreatedUTC, 10)
	userCreatedStr := strconv.FormatUint(cwt.AuthorCreatedUTC, 10)

  if userCreatedStr == "0" {
    out = append(out, "")
    out = append(out, "")
    out = append(out, "")
    out = append(out, "")
  } else {
    user_tm := time.Unix(int64(cwt.CreatedUTC), 0)
    out = append(out, strconv.FormatInt(int64(user_tm.Year()), 10))
    out = append(out, strconv.FormatInt(int64(user_tm.Month()), 10))
    out = append(out, strconv.FormatInt(int64(user_tm.Day()), 10))
    out = append(out, user_tm.Format("15:04:05"))
  }

  out = append(out, createdStr)
	out = append(out, userCreatedStr)
	out = append(out, cwt.Permalink)
	out = append(out, cwt.SubredditType)
	out = append(out, lang)
  out = append(out, cwt.Id)
  out = append(out, cwt.ParentId)
  out = append(out, cwt.LinkId)

	return out
}
