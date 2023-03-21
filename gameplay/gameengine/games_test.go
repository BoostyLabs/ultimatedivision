package gameengine_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision"
	"ultimatedivision/database/dbtesting"
	"ultimatedivision/gameplay/gameengine"
)

func TestMatches(t *testing.T) {
	matchID := uuid.New()
	card1 := gameengine.Card{
		CardID:   uuid.New(),
		Position: 1,
	}
	card2 := gameengine.Card{
		CardID:   uuid.New(),
		Position: 2,
	}
	card3 := gameengine.Card{
		CardID:   uuid.New(),
		Position: 3,
	}

	testGame := gameengine.Game{
		MatchID: matchID,
		GameInfo: []gameengine.Card{
			card1,
			card2,
			card3,
		},
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repositoryGames := db.Games()

		t.Run("Create", func(t *testing.T) {
			err := repositoryGames.Create(ctx, testGame)
			require.NoError(t, err)
		})
		t.Run("Get", func(t *testing.T) {
			game, err := repositoryGames.Get(ctx, matchID)
			fmt.Println("gameMatchID ---> ", game.MatchID)
			fmt.Println("gameInfo ---> ", game.GameInfo)
			compareGame(t, testGame, game)
			require.NoError(t, err)
		})

		//t.Run("Create", func(t *testing.T) {
		//	_, err := repositoryGames.List(ctx)
		//	//compareGame(t, testGame, games)
		//	require.NoError(t, err)
		//})
		//t.Run("Create", func(t *testing.T) {
		//	err := repositoryGames.Create(ctx, testGame)
		//	require.NoError(t, err)
		//})

	})
}

func compareGame(t *testing.T, expectedGame, actualGame gameengine.Game) {
	assert.Equal(t, expectedGame.MatchID, actualGame.MatchID)
	//for i, card := range actualGame.GameInfo {
	//	assert.Equal(t, expectedGame.GameInfo[i].CardID, card.CardID)
	//	assert.Equal(t, expectedGame.GameInfo[i].Position, card.Position)
	//}
}
