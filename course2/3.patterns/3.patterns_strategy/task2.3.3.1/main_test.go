package main

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v53/github"
	"github.com/stretchr/testify/assert"
)

func TestGetItemsWithGistStrategy(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGist := NewMockGithubLister(ctrl)

	expectedItems := []Item{
		{Title: "gist1", Description: "Description1", Link: "http://example.com/gist1"},
		{Title: "gist2", Description: "Description2", Link: "http://example.com/gist2"},
	}

	mockGist.EXPECT().GetItems(gomock.Any(), "doktorghost").Return(expectedItems, nil)

	client := &github.Client{} // Здесь не важно, так как мы мокаем
	gg := NewGeneralGithub(client)

	items, err := gg.GetItems(context.Background(), "doktorghost", mockGist)
	assert.NoError(t, err)
	assert.Equal(t, expectedItems, items)
}

func TestGetItemsWithRepoStrategy(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockGithubLister(ctrl)

	expectedItems := []Item{
		{Title: "repo1", Description: "Description1", Link: "http://example.com/repo1"},
		{Title: "repo2", Description: "Description2", Link: "http://example.com/repo2"},
	}

	mockRepo.EXPECT().GetItems(gomock.Any(), "doktorghost").Return(expectedItems, nil)

	client := &github.Client{} // Здесь не важно, так как мы мокаем
	gg := NewGeneralGithub(client)

	items, err := gg.GetItems(context.Background(), "doktorghost", mockRepo)
	assert.NoError(t, err)
	assert.Equal(t, expectedItems, items)
}
