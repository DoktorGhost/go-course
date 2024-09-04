package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v64/github"
)

func main() {
	client := github.NewClient(nil).WithAuthToken("token")
	g := NewGithubAdapter(client)
	fmt.Println(g.GetGists(context.Background(), "ptflp"))
	fmt.Println(g.GetRepos(context.Background(), "doktorghost"))
}

type RepoLister interface {
	List(ctx context.Context, username string, opt *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error)
}
type GistLister interface {
	List(ctx context.Context, username string, opt *github.GistListOptions) ([]*github.Gist,
		*github.Response, error)
}
type Githuber interface {
	GetGists(ctx context.Context, username string) ([]Item, error)
	GetRepos(ctx context.Context, username string) ([]Item, error)
}
type GithubAdapter struct {
	RepoList RepoLister
	GistList GistLister
}

func NewGithubAdapter(githubClient *github.Client) *GithubAdapter {
	g := &GithubAdapter{
		RepoList: githubClient.Repositories,
		GistList: githubClient.Gists,
	}
	return g
}

func (g *GithubAdapter) GetGists(ctx context.Context, username string) ([]Item, error) {
	gists, _, err := g.GistList.List(ctx, username, nil)
	if err != nil {
		return nil, err
	}

	var items []Item
	for _, gist := range gists {
		item := Item{
			Title:       gist.GetID(),
			Description: gist.GetDescription(),
			Link:        gist.GetHTMLURL(),
		}
		items = append(items, item)
	}

	return items, nil
}
func (g *GithubAdapter) GetRepos(ctx context.Context, username string) ([]Item, error) {
	repos, _, err := g.RepoList.List(ctx, username, nil)
	if err != nil {
		return nil, err
	}

	var items []Item
	for _, repo := range repos {
		item := Item{
			Title:       repo.GetName(),
			Description: repo.GetDescription(),
			Link:        repo.GetHTMLURL(),
		}
		items = append(items, item)
	}

	return items, nil
}

type Item struct {
	Title       string
	Description string
	Link        string
}
