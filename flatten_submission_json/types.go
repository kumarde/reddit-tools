package main

import (
	"fmt"
	"strconv"
  "time"
)

type ModNote struct { //https://github.com/praw-dev/praw/blob/master/praw/models/mod_note.py#L9
  Action    string  `json:"action"`
  CreatedAt uint64  `json:"created_at"`
  Description string  `json:"description"`
  Details   string  `json:"details"`
  Id        string  `json:"id"`
  Label     string  `json:"label"`
  Moderator string  `json:"moderator"`
  Note      string  `json:"note"`
  RedditId  string  `json:"reddit_id"`
  Subreddit string  `json:"subreddit"`
  Type      string  `json:"type"`
  User      string  `json:"user"`
}

type RemoveReason struct {
  Id  string  `json:"id"`
  Messeage string `json:"message"`
  Title    string  `json:"title"`
}

type Media struct {
  MediaDetail struct {
    Description string  `json:"description"`
    HTML        string  `json:"html"`
    Provider    string  `json:"provider_name"`
    Title       string  `json:"title"`
  }   `json:"oembed"`
  Type     string `json:type`
}

type RedditComment struct {
	Meta struct {
    RemovalType     string  `json:"removal_type"`
    WasDeleted      bool    `json:"was_deleted_later"`
  }  `json:"_meta"`
  AllAwardings      []string  `json:"all_awardings"`
  AllowLiveComments bool    `json:"allow_live_comments"`
  ApprovedUTC       uint64  `json:"approved_at_utc"`
  ApprovedBy        string  `json:"approved_by"`
	Author            string  `json:"author"`
  Archived          bool    `json:"archived"`
  AuthorFlairText   string  `json:"author_flair_text"`
  AuthorFullName    string  `json:"author_fullname"`
  AuthorIsBlocked   bool    `json:"author_is_blocked"`
  AuthorPatreonFlair  bool  `json:"author_patreon_flair"`
  AuthorPremium     bool  `json:"author_premium"`
  BannedUTC         uint64  `json:"banned_at_utc"`
  BannedBy          string  `json:"banned_by"`
  CanModPost        bool    `json:"can_mod_post"`
  Category          []string  `json:"category"`
  Clicked           bool    `json:"clicked"`
  ContentCategory   []string  `json:"content_categories"`
  ContestMode       bool    `json:"contest_mode"`
  CreatedUTC        uint64  `json:"created_utc"`
  DiscussionType    string  `json:"discussion_type"`
  Distinguished     string  `json:"distinguished"`
  Domain            string  `json:"domain"`
  Downs             uint64  `json:"downs"`
  Edited            bool    `json:"edited"`
  Hidden            bool    `json:"hidden"`
  HideScore         bool    `json:"hide_score"`
  Id                string  `json:"id"`
  IsCreatedFromAd   bool    `json:"is_created_from_ad_ui"`
  IsCrosspostable   bool    `json:"is_crosspostable"`
  IsMeta            bool    `json:"is_meta"`
  IsOriginalContent bool    `json:"is_original_content"`
  IsSelf            bool    `json:"is_self"`
  IsVideo           bool    `json:"is_video"`
  LinkFlairTemplate string  `json:"link_flair_template_id"`
  LinkFlairText     string  `json:"link_flair_text"`
  Locked            bool    `json:"locked"`
  Media             Media   `json:"media"`
  ModNote           ModNote  `json:"mod_note`
  ModReason         string  `json:"mod_reason_by"` // TODO check type - likely to be dict
  ModReasonTitle    string  `json:"mod_reason_title"`
  ModReports        []string  `json:"mod_reports"`
  Name              string  `json:"name"`
  NoFollow          bool    `json:"no_follow"`
  NumComment        uint64  `json:"num_comments"`
  NumCrossPost      uint64  `json:"num_crossposts"`
  Over18            bool    `json:"over_18"`
	Permalink         string  `json:"permalink"`
  Pinned            bool    `json:"pinned"`
  PostHint          string  `json:"post_hint"`
  Quarantine        bool    `json:"quarantine"`
  RemovalReason     RemoveReason  `json:"removal_reason"`
  RemovedBy         string  `json:"removed_by"`
  RemovedByCat      string  `json:"removed_by_category"`
  RetrievedOn       uint64  `json:"retrieved_on"`
  Saved             bool    `json:"saved"`
  Score             int64   `json:"score"`
  SelfText          string  `json:"selftext"`
  Spoiler           bool    `json:"spolier"`
  Stickied          bool    `json:"stickied"`
	Subreddit         string  `json:"subreddit"`
	SubredditType     string  `json:"subreddit_type"`
  Title             string  `json:"title"`
  TotalAwards       uint64  `json:"total_awards_received"`
  UpdateUTC         uint64  `json:"updated_on"`
  UpvoteRatio       float64 `json:"upvote_ratio"`
  Ups               uint64  `json:"ups"`
  URL               string  `json:"url"`
  UserReports       []string  `json:"user_reports"`
  ViewCount         uint64  `json:"view_count"`
  WhitelistStatus   string  `json:"whitelist_status"`
}

func (cwt *RedditComment) buildOutput(lang string) []string {
	out := []string{}
  
  tm := time.Unix(int64(cwt.CreatedUTC), 0)
  fmt.Println(tm)
  fmt.Println(cwt.CreatedUTC)
	out = append(out, strconv.FormatInt(int64(tm.Year()), 10))
  out = append(out, strconv.FormatInt(int64(tm.Month()), 10))
  out = append(out, strconv.FormatInt(int64(tm.Day()), 10))
  out = append(out, tm.Format("15:04:05"))
  out = append(out, strconv.FormatUint(cwt.CreatedUTC, 10))

	out = append(out, cwt.Meta.RemovalType)
  out = append(out, strconv.FormatBool(cwt.Meta.WasDeleted))

	out = append(out, fmt.Sprint(cwt.AllAwardings))
	out = append(out, strconv.FormatBool(cwt.AllowLiveComments))
  out = append(out, strconv.FormatUint(cwt.ApprovedUTC, 10))
  out = append(out, cwt.ApprovedBy)

  out = append(out, cwt.Author)
  out = append(out, strconv.FormatBool(cwt.Archived))
  out = append(out, cwt.AuthorFlairText)
  out = append(out, cwt.AuthorFullName)
  out = append(out, strconv.FormatBool(cwt.AuthorIsBlocked))
  out = append(out, strconv.FormatBool(cwt.AuthorPatreonFlair))
  out = append(out, strconv.FormatBool(cwt.AuthorPremium))

  out = append(out, strconv.FormatUint(cwt.BannedUTC, 10))
  out = append(out, cwt.BannedBy)
  out = append(out, strconv.FormatBool(cwt.CanModPost))
  out = append(out, fmt.Sprint(cwt.Category))
  out = append(out, strconv.FormatBool(cwt.Clicked))
  out = append(out, fmt.Sprint(cwt.ContentCategory))
  out = append(out, strconv.FormatBool(cwt.ContestMode))

	out = append(out, cwt.DiscussionType)
  out = append(out, cwt.Distinguished)
	out = append(out, cwt.Domain)
	out = append(out, strconv.FormatUint(cwt.Downs, 10))
  out = append(out, strconv.FormatBool(cwt.Edited))
  out = append(out, strconv.FormatBool(cwt.Hidden))
  out = append(out, strconv.FormatBool(cwt.HideScore))
  out = append(out, cwt.Id)
  out = append(out, strconv.FormatBool(cwt.IsCreatedFromAd))
  out = append(out, strconv.FormatBool(cwt.IsCrosspostable))
  out = append(out, strconv.FormatBool(cwt.IsMeta))
  out = append(out, strconv.FormatBool(cwt.IsOriginalContent))
  out = append(out, strconv.FormatBool(cwt.IsSelf))
  out = append(out, strconv.FormatBool(cwt.IsVideo))
  out = append(out, cwt.LinkFlairTemplate)
  out = append(out, cwt.LinkFlairText)
  out = append(out, strconv.FormatBool(cwt.Locked))

  out = append(out, cwt.Media.MediaDetail.Description)
  out = append(out, cwt.Media.MediaDetail.HTML)
  out = append(out, cwt.Media.MediaDetail.Provider)
  out = append(out, cwt.Media.MediaDetail.Title)
  out = append(out, cwt.Media.Type)

  out = append(out, cwt.ModNote.Action)
  out = append(out, strconv.FormatUint(cwt.ModNote.CreatedAt, 10))
  out = append(out, cwt.ModNote.Description)
  out = append(out, cwt.ModNote.Details)
  out = append(out, cwt.ModNote.Id)
  out = append(out, cwt.ModNote.Label)
  out = append(out, cwt.ModNote.Moderator)
  out = append(out, cwt.ModNote.Note)
  out = append(out, cwt.ModNote.RedditId)
  out = append(out, cwt.ModNote.Subreddit)
  out = append(out, cwt.ModNote.Type)
  out = append(out, cwt.ModNote.User)

  out = append(out, cwt.ModReason)
  out = append(out, cwt.ModReasonTitle)
  out = append(out, fmt.Sprint(cwt.ModReports))
  out = append(out, cwt.Name)
  out = append(out, strconv.FormatBool(cwt.NoFollow))
  out = append(out, strconv.FormatUint(cwt.NumComment, 10))
  out = append(out, strconv.FormatUint(cwt.NumCrossPost, 10))
  out = append(out, strconv.FormatBool(cwt.Over18))
  out = append(out, cwt.Permalink)
  out = append(out, strconv.FormatBool(cwt.Pinned))
  out = append(out, cwt.PostHint)
  out = append(out, strconv.FormatBool(cwt.Quarantine))
  
  out = append(out, cwt.RemovalReason.Id)
  out = append(out, cwt.RemovalReason.Messeage)
  out = append(out, cwt.RemovalReason.Title)

  out = append(out, cwt.RemovedBy)
  out = append(out, cwt.RemovedByCat)

  out = append(out, strconv.FormatUint(cwt.RetrievedOn, 10))
  out = append(out, strconv.FormatBool(cwt.Saved))
  out = append(out, strconv.FormatInt(cwt.Score, 10))
  
  out = append(out, cwt.SelfText)

  out = append(out, strconv.FormatBool(cwt.Spoiler))
  out = append(out, strconv.FormatBool(cwt.Stickied))

  out = append(out, cwt.Subreddit)
  out = append(out, cwt.SubredditType)
  out = append(out, cwt.Title)

  out = append(out, strconv.FormatUint(cwt.TotalAwards, 10))
  out = append(out, strconv.FormatUint(cwt.UpdateUTC, 10))
  out = append(out, strconv.FormatFloat(cwt.UpvoteRatio, 'f', 3, 64))
  out = append(out, strconv.FormatUint(cwt.Ups, 10))

  out = append(out, cwt.URL)
  out = append(out, fmt.Sprint(cwt.UserReports))
  out = append(out, strconv.FormatUint(cwt.ViewCount, 10))
  out = append(out, cwt.WhitelistStatus)

  out = append(out, lang)

	return out
}
