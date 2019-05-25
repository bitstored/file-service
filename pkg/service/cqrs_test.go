package service

import (
	"github.com/bitstored/file-service/pkg/service/events"
	"github.com/stretchr/testify/require"
	"github.com/bitstored/file-service/pkg/service/commands"
	"context"
	"testing"
)

func TestCreateDrive(t *testing.T) {
	s := NewService()
	command := commands.CreateDrive{UserID: t.Name()}
	r := s.LaunchCommand(context.TODO(), command )
	require.NotNil(t, r)
	rc, ok := r.(*events.DriveCreated)
	require.NoError(t, rc.Error)
	require.True(t, ok)
	require.Equal(t, t.Name(), rc.UserID)
}