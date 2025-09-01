package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

/*
A clean terminal chat room over Redis Pub/Sub.

Highlights:
- Interactive prompts if flags are omitted (username/room).
- Rooms (channels), message history (LPUSH/LTRIM), and presence (SET with TTL).
- JSON messages (type/user/room/text/ts).
- Commands: /help, /users, /history [N], /quit
- Graceful shutdown with join/leave announcements.
*/

var (
	// Redis data model
	historyKeyFmt  = "chat:history:%s"     // LPUSH newest, LTRIM to max
	channelFmt     = "chat:room:%s"        // Pub/Sub channel per room
	presenceKeyFmt = "chat:presence:%s:%s" // SET presence with TTL (room:user)

	// Tunables
	historyMaxLen  = 200
	presenceTTL    = 25 * time.Second
	heartbeatEvery = 10 * time.Second
)

type Msg struct {
	Type string `json:"type"` // chat | join | leave | system
	Room string `json:"room"`
	User string `json:"user"`
	Text string `json:"text"`
	TS   string `json:"ts"` // RFC3339
}

func nowISO() string {
	return time.Now().UTC().Format(time.RFC3339Nano)
}

func serialize(mt, room, user, text string) string {
	b, _ := json.Marshal(Msg{
		Type: mt, Room: room, User: user, Text: text, TS: nowISO(),
	})
	return string(b)
}

func pushHistory(ctx context.Context, rdb *redis.Client, room, jsonMsg string) {
	key := fmt.Sprintf(historyKeyFmt, room)
	pipe := rdb.TxPipeline()
	pipe.LPush(ctx, key, jsonMsg)
	pipe.LTrim(ctx, key, 0, int64(historyMaxLen-1))
	_, _ = pipe.Exec(ctx)

}
