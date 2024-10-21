package main

import (
	"fmt"
	"github.com/teamnsrg/go-perspectiveapi/perspective"
	"strconv"
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
	Archived         bool    `json:"archived"`
	Author           string  `json:"author"`
	AuthorCreatedUTC uint64 `json:"author_created_utc"`
	Body             string  `json:"body"`
	CreatedUTC       uint64 `json:"created_utc"`
	Permalink        string  `json:"permalink"`
	Subreddit        string  `json:"subreddit"`
	SubredditType    string  `json:"subreddit_type"`
	Id               string
	ParentId         string
	NestLevel        int64
	LinkId           string
}

type CommentWithToxicity struct {
	Comment  *RedditComment                      `json:"comment"`
	Toxicity *perspective.AnalyzeCommentResponse `json:"toxicity"`
}

func getToxicScoreFromAttribute(t *perspective.AnalyzeCommentResponse, attribute string) string {
	attributeScores := t.AttributeScores
	if attributeScores == nil {
		return "-1.0"
	}
	spanScores := *attributeScores[attribute].SpanScores
	return fmt.Sprintf("%f", spanScores[0].Score.Value)
}

func (cwt *CommentWithToxicity) buildOutput(lang string) []string {
	out := []string{}
	out = append(out, cwt.Comment.Author)
	out = append(out, cwt.Comment.Body)
	out = append(out, cwt.Comment.Subreddit)

	createdStr := strconv.FormatUint(cwt.Comment.CreatedUTC, 10)
	userCreatedStr := strconv.FormatUint(cwt.Comment.AuthorCreatedUTC, 10)

	out = append(out, createdStr)
	out = append(out, userCreatedStr)
	out = append(out, cwt.Comment.Permalink)
	out = append(out, cwt.Comment.SubredditType)
	out = append(out, lang)

	out = append(out, getToxicScoreFromAttribute(cwt.Toxicity, "ATTACK_ON_AUTHOR"))
	out = append(out, getToxicScoreFromAttribute(cwt.Toxicity, "IDENTITY_ATTACK"))
	out = append(out, getToxicScoreFromAttribute(cwt.Toxicity, "INSULT"))
	out = append(out, getToxicScoreFromAttribute(cwt.Toxicity, "PROFANITY"))
	out = append(out, getToxicScoreFromAttribute(cwt.Toxicity, "SEVERE_TOXICITY"))
	out = append(out, getToxicScoreFromAttribute(cwt.Toxicity, "SEXUALLY_EXPLICIT"))
	out = append(out, getToxicScoreFromAttribute(cwt.Toxicity, "THREAT"))
	out = append(out, getToxicScoreFromAttribute(cwt.Toxicity, "TOXICITY"))

	return out
}
