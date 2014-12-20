package github

import (
	"encoding/json"
	"testing"
)

func TestDeserialize(t *testing.T) {
	payload := _unmarshal(t)
	if payload.Repository.CreatedAt.String() != "2014-05-20 11:51:35 +0200 CEST" {
		t.Error("Timestamp is not parsed as expected: %s", payload.Repository.CreatedAt.String())
	}

	if payload.HeadCommit.Timestamp.String() != "2014-11-23 02:19:56 +0100 CET" {
		t.Errorf("Time is not parsed as expected: %s", payload.HeadCommit.Timestamp.String())
	}

	if payload.Repository.DefaultBranch != "master" {
		t.Errorf("DefaultBranch is not parsed as expected: %s", payload.Repository.DefaultBranch)
	}

	if payload.Repository.FullName != "GitbookIO/documentation" {
		t.Errorf("FullName is not parsed as expected: %s", payload.Repository.DefaultBranch)
	}
}

func _unmarshal(t *testing.T) GitHubPayload {
	payload := GitHubPayload{}
	if err := json.Unmarshal([]byte(jsonPayload), &payload); err != nil {
		t.Error(err)
	}
	return payload
}

const jsonPayload string = `
{
	"ref": "refs/heads/master",
	"before": "421eadcf4c46b860e7b62b899e15c61233a3351d",
	"after": "14a7480c37bff138ae2629c38d2e1c7e9f1324f9",
	"created": false,
	"deleted": false,
	"forced": false,
	"base_ref": null,
	"compare": "https://github.com/GitbookIO/documentation/compare/421eadcf4c46...14a7480c37bf",
	"commits": [
		{
			"id": "14a7480c37bff138ae2629c38d2e1c7e9f1324f9",
			"distinct": true,
			"message": "Update format.md",
			"timestamp": "2014-11-23T02:19:56+01:00",
			"url": "https://github.com/GitbookIO/documentation/commit/14a7480c37bff138ae2629c38d2e1c7e9f1324f9",
			"author": {
				"name": "Aaron O&#39;Mullan",
				"email": "aaron.omullan@gmail.com",
				"username": "AaronO"
			},
			"committer": {
				"name": "Aaron O&#39;Mullan",
				"email": "aaron.omullan@gmail.com",
				"username": "AaronO"
			},
			"added": [

			],
			"removed": [

			],
			"modified": [
				"book/format.md"
			]
		}
	],
	"head_commit": {
		"id": "14a7480c37bff138ae2629c38d2e1c7e9f1324f9",
		"distinct": true,
		"message": "Update format.md",
		"timestamp": "2014-11-23T02:19:56+01:00",
		"url": "https://github.com/GitbookIO/documentation/commit/14a7480c37bff138ae2629c38d2e1c7e9f1324f9",
		"author": {
			"name": "Aaron O&#39;Mullan",
			"email": "aaron.omullan@gmail.com",
			"username": "AaronO"
		},
		"committer": {
			"name": "Aaron O&#39;Mullan",
			"email": "aaron.omullan@gmail.com",
			"username": "AaronO"
		},
		"added": [

		],
		"removed": [

		],
		"modified": [
			"book/format.md"
		]
	},
	"repository": {
		"id": 19975935,
		"name": "documentation",
		"full_name": "GitbookIO/documentation",
		"owner": {
			"name": "GitbookIO",
			"email": "contact@gitbook.io"
		},
		"private": false,
		"html_url": "https://github.com/GitbookIO/documentation",
		"description": "Documentation for GitBook and gitbook.io",
		"fork": false,
		"url": "https://github.com/GitbookIO/documentation",
		"forks_url": "https://api.github.com/repos/GitbookIO/documentation/forks",
		"keys_url": "https://api.github.com/repos/GitbookIO/documentation/keys{/key_id}",
		"collaborators_url": "https://api.github.com/repos/GitbookIO/documentation/collaborators{/collaborator}",
		"teams_url": "https://api.github.com/repos/GitbookIO/documentation/teams",
		"hooks_url": "https://api.github.com/repos/GitbookIO/documentation/hooks",
		"issue_events_url": "https://api.github.com/repos/GitbookIO/documentation/issues/events{/number}",
		"events_url": "https://api.github.com/repos/GitbookIO/documentation/events",
		"assignees_url": "https://api.github.com/repos/GitbookIO/documentation/assignees{/user}",
		"branches_url": "https://api.github.com/repos/GitbookIO/documentation/branches{/branch}",
		"tags_url": "https://api.github.com/repos/GitbookIO/documentation/tags",
		"blobs_url": "https://api.github.com/repos/GitbookIO/documentation/git/blobs{/sha}",
		"git_tags_url": "https://api.github.com/repos/GitbookIO/documentation/git/tags{/sha}",
		"git_refs_url": "https://api.github.com/repos/GitbookIO/documentation/git/refs{/sha}",
		"trees_url": "https://api.github.com/repos/GitbookIO/documentation/git/trees{/sha}",
		"statuses_url": "https://api.github.com/repos/GitbookIO/documentation/statuses/{sha}",
		"languages_url": "https://api.github.com/repos/GitbookIO/documentation/languages",
		"stargazers_url": "https://api.github.com/repos/GitbookIO/documentation/stargazers",
		"contributors_url": "https://api.github.com/repos/GitbookIO/documentation/contributors",
		"subscribers_url": "https://api.github.com/repos/GitbookIO/documentation/subscribers",
		"subscription_url": "https://api.github.com/repos/GitbookIO/documentation/subscription",
		"commits_url": "https://api.github.com/repos/GitbookIO/documentation/commits{/sha}",
		"git_commits_url": "https://api.github.com/repos/GitbookIO/documentation/git/commits{/sha}",
		"comments_url": "https://api.github.com/repos/GitbookIO/documentation/comments{/number}",
		"issue_comment_url": "https://api.github.com/repos/GitbookIO/documentation/issues/comments/{number}",
		"contents_url": "https://api.github.com/repos/GitbookIO/documentation/contents/{+path}",
		"compare_url": "https://api.github.com/repos/GitbookIO/documentation/compare/{base}...{head}",
		"merges_url": "https://api.github.com/repos/GitbookIO/documentation/merges",
		"archive_url": "https://api.github.com/repos/GitbookIO/documentation/{archive_format}{/ref}",
		"downloads_url": "https://api.github.com/repos/GitbookIO/documentation/downloads",
		"issues_url": "https://api.github.com/repos/GitbookIO/documentation/issues{/number}",
		"pulls_url": "https://api.github.com/repos/GitbookIO/documentation/pulls{/number}",
		"milestones_url": "https://api.github.com/repos/GitbookIO/documentation/milestones{/number}",
		"notifications_url": "https://api.github.com/repos/GitbookIO/documentation/notifications{?since,all,participating}",
		"labels_url": "https://api.github.com/repos/GitbookIO/documentation/labels{/name}",
		"releases_url": "https://api.github.com/repos/GitbookIO/documentation/releases{/id}",
		"created_at": 1400579495,
		"updated_at": "2014-11-22T23:50:39Z",
		"pushed_at": 1416705597,
		"git_url": "git://github.com/GitbookIO/documentation.git",
		"ssh_url": "git@github.com:GitbookIO/documentation.git",
		"clone_url": "https://github.com/GitbookIO/documentation.git",
		"svn_url": "https://github.com/GitbookIO/documentation",
		"homepage": "http://help.gitbook.io/",
		"size": 1336,
		"stargazers_count": 37,
		"watchers_count": 37,
		"language": null,
		"has_issues": true,
		"has_downloads": true,
		"has_wiki": false,
		"has_pages": false,
		"forks_count": 35,
		"mirror_url": null,
		"open_issues_count": 2,
		"forks": 35,
		"open_issues": 2,
		"watchers": 37,
		"default_branch": "master",
		"stargazers": 37,
		"master_branch": "master",
		"organization": "GitbookIO"
	},
	"pusher": {
		"name": "AaronO",
		"email": "aaron.omullan@gmail.com"
	},
	"organization": {
		"login": "GitbookIO",
		"id": 7111340,
		"url": "https://api.github.com/orgs/GitbookIO",
		"repos_url": "https://api.github.com/orgs/GitbookIO/repos",
		"events_url": "https://api.github.com/orgs/GitbookIO/events",
		"members_url": "https://api.github.com/orgs/GitbookIO/members{/member}",
		"public_members_url": "https://api.github.com/orgs/GitbookIO/public_members{/member}",
		"avatar_url": "https://avatars.githubusercontent.com/u/7111340?v=3"
	},
	"sender": {
		"login": "AaronO",
		"id": 949223,
		"avatar_url": "https://avatars.githubusercontent.com/u/949223?v=3",
		"gravatar_id": "",
		"url": "https://api.github.com/users/AaronO",
		"html_url": "https://github.com/AaronO",
		"followers_url": "https://api.github.com/users/AaronO/followers",
		"following_url": "https://api.github.com/users/AaronO/following{/other_user}",
		"gists_url": "https://api.github.com/users/AaronO/gists{/gist_id}",
		"starred_url": "https://api.github.com/users/AaronO/starred{/owner}{/repo}",
		"subscriptions_url": "https://api.github.com/users/AaronO/subscriptions",
		"organizations_url": "https://api.github.com/users/AaronO/orgs",
		"repos_url": "https://api.github.com/users/AaronO/repos",
		"events_url": "https://api.github.com/users/AaronO/events{/privacy}",
		"received_events_url": "https://api.github.com/users/AaronO/received_events",
		"type": "User",
		"site_admin": false
	}
}
`
