package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v64/github"
)

func main() {
	client := github.NewClient(nil).WithAuthToken("token")
	a := NewGithubAdapter(client)
	g := NewGithubProxy(*a)

	fmt.Println(g.GetGists(context.Background(), "ptflp"))
	fmt.Println(g.GetRepos(context.Background(), "doktorghost"))

	fmt.Println(g.cache)

	fmt.Println(g.GetGists(context.Background(), "ptflp"))
	fmt.Println(g.GetRepos(context.Background(), "doktorghost"))

	fmt.Println(g.GetGists(context.Background(), "ptflp"))
	fmt.Println(g.GetRepos(context.Background(), "doktorghost"))

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

type GithubProxy struct {
	github Githuber
	cache  map[string][]Item
}

func NewGithubProxy(github GithubAdapter) *GithubProxy {
	return &GithubProxy{
		github: &github,
		cache:  make(map[string][]Item),
	}
}

func (g *GithubProxy) GetGists(ctx context.Context, username string) ([]Item, error) {
	if gists, ok := g.cache["gists_"+username]; ok {
		return gists, nil
	}

	gists, err := g.github.GetGists(ctx, username)
	if err != nil {
		return nil, err
	}
	g.cache["gists_"+username] = gists

	return gists, nil
}
func (g *GithubProxy) GetRepos(ctx context.Context, username string) ([]Item, error) {
	if repos, ok := g.cache["repos_"+username]; ok {
		return repos, nil
	}

	repos, err := g.github.GetRepos(ctx, username)
	if err != nil {
		return nil, err
	}
	g.cache["repos_"+username] = repos

	return repos, nil
}
