package handler

import "context"

func (h *Handler) CreateChat(ctx context.Context, usernames []string) (int64, error) {
	chatID, err := h.chatClient.CreateChat(ctx, usernames)
	if err != nil {
		return 0, err
	}

	return chatID, nil
}
