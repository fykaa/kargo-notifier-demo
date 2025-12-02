package github

import (
    "context"
    "fmt"

    gh "github.com/google/go-github/v61/github"
)

type MergeQueueClient struct {
    gh *gh.Client
}

func NewMergeQueueClient(c *gh.Client) *MergeQueueClient {
    return &MergeQueueClient{gh: c}
}

func (m *MergeQueueClient) EnqueuePullRequest(
    ctx context.Context,
    owner, repo string,
    prNumber int,
    mergeMethod string,
) error {
    req := &gh.MergeQueueEntryRequest{
        PullRequestNumber: prNumber,
        MergeMethod:       mergeMethod,
    }

    _, _, err := m.gh.MergeQueue.CreateEntry(ctx, owner, repo, req)
    if err != nil {
        return fmt.Errorf("enqueue failed: %w", err)
    }

    return nil
}

func (m *MergeQueueClient) GetQueueState(
    ctx context.Context,
    owner, repo string,
) ([]*gh.MergeQueueEntry, error) {
    entries, _, err := m.gh.MergeQueue.ListEntries(ctx, owner, repo)
    if err != nil {
        return nil, fmt.Errorf("list failed: %w", err)
    }

    return entries, nil
}

func (m *MergeQueueClient) DequeuePullRequest(
    ctx context.Context,
    owner, repo string,
    entryID int64,
) error {
    _, err := m.gh.MergeQueue.DeleteEntry(ctx, owner, repo, entryID)
    if err != nil {
        return fmt.Errorf("dequeue failed: %w", err)
    }

    return nil
}
