package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
	"log"
	"os"
)

type Item struct {
	Title       string
	Description string
	Link        string
}

// интерфейс стратегии
//
//go:generate mockgen -source=main.go -destination=mock_interfaces.go -package=mocks
type GithubLister interface {
	GetItems(ctx context.Context, username string) ([]Item, error)
}

// интерфейс с выбором стратегии
type GeneralGithubLister interface {
	GetItems(ctx context.Context, username string, strategy GithubLister) ([]Item, error)
}

type GeneralGithub struct {
	client *github.Client
}

func NewGeneralGithub(client *github.Client) GeneralGithub {
	return GeneralGithub{
		client: client,
	}
}

func (g GeneralGithub) GetItems(ctx context.Context, username string, strategy GithubLister) ([]Item, error) {
	return strategy.GetItems(ctx, username)
}

type GithubGist struct {
	service *github.GistsService
}

func NewGithubGist(client *github.Client) *GithubGist {
	return &GithubGist{
		service: client.Gists,
	}
}

func (g GithubGist) GetItems(ctx context.Context, username string) ([]Item, error) {
	gists, _, err := g.service.List(ctx, username, nil)
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

type GithubRepo struct {
	service *github.RepositoriesService
}

func NewGithubRepo(client *github.Client) *GithubRepo {
	return &GithubRepo{
		service: client.Repositories,
	}
}
func (g GithubRepo) GetItems(ctx context.Context, username string) ([]Item, error) {
	repos, _, err := g.service.List(ctx, username, nil)
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
func main() {
	ctx := context.Background()
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		fmt.Println("GITHUB_TOKEN is not set")
		return
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	gist := NewGithubGist(client)
	repo := NewGithubRepo(client)

	gg := NewGeneralGithub(client)

	data, err := gg.GetItems(context.Background(), "doktorghost", gist)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)
	data, err = gg.GetItems(context.Background(), "doktorghost", repo)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)

}
