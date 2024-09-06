package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGithubAdapter_GetRepos(t *testing.T) {
	// Создаем моковый объект RepoLister
	mockRepoLister := &MockRepoLister{}
	adapter := &GithubAdapter{RepoList: mockRepoLister}

	// Выполняем тестируемый метод
	repos, err := adapter.GetRepos(context.Background(), "user")

	// Проверяем отсутствие ошибок
	assert.NoError(t, err)

	// Проверяем количество репозиториев
	assert.Equal(t, 3, len(repos))

	// Проверяем данные первого репозитория
	assert.Equal(t, "Repo1", repos[0].Title)
	assert.Equal(t, "Description1", repos[0].Description)
	assert.Equal(t, "https://github.com/user/repo1", repos[0].Link)
}

func TestGithubAdapter_GetGists(t *testing.T) {
	// Создаем моковый объект GistLister
	mockGistLister := &MockGistLister{}
	adapter := &GithubAdapter{GistList: mockGistLister}

	// Выполняем тестируемый метод
	gists, err := adapter.GetGists(context.Background(), "user")

	// Проверяем отсутствие ошибок
	assert.NoError(t, err)

	// Проверяем количество гистов
	assert.Equal(t, 3, len(gists))

	// Проверяем данные первого гиста
	assert.Equal(t, "Gist1", gists[0].Title)
	assert.Equal(t, "Description1", gists[0].Description)
	assert.Equal(t, "https://gist.github.com/user/gist1", gists[0].Link)
}
