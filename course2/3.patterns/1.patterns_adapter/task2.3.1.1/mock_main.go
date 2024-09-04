package main

import (
	"context"
	"github.com/google/go-github/v64/github"
)

// MockRepoLister - мок для интерфейса RepoLister
type MockRepoLister struct{}

func (m *MockRepoLister) List(ctx context.Context, username string, opt *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error) {
	repos := []*github.Repository{
		{Name: github.String("Repo1"), Description: github.String("Description1"), HTMLURL: github.String("https://github.com/user/repo1")},
		{Name: github.String("Repo2"), Description: github.String("Description2"), HTMLURL: github.String("https://github.com/user/repo2")},
		{Name: github.String("Repo3"), Description: github.String("Description3"), HTMLURL: github.String("https://github.com/user/repo3")},
	}
	return repos, nil, nil
}

// MockGistLister - мок для интерфейса GistLister
type MockGistLister struct{}

func (m *MockGistLister) List(ctx context.Context, username string, opt *github.GistListOptions) ([]*github.Gist, *github.Response, error) {
	gists := []*github.Gist{
		{ID: github.String("Gist1"), Description: github.String("Description1"), HTMLURL: github.String("https://gist.github.com/user/gist1")},
		{ID: github.String("Gist2"), Description: github.String("Description2"), HTMLURL: github.String("https://gist.github.com/user/gist2")},
		{ID: github.String("Gist3"), Description: github.String("Description3"), HTMLURL: github.String("https://gist.github.com/user/gist3")},
	}
	return gists, nil, nil
}
