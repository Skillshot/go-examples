package server_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go-examples/logger"
	"go-examples/mocks"
	"go-examples/server"
	"testing"
)

func TestServer_RenderAdminPage(t *testing.T) {
	type args struct {
		user string
	}
	tests := []struct {
		name string
		//fields fields
		args         args
		expectedPage string
	}{
		{
			name: "Should return Admin Page when user is admin",
			args: args{
				user: "admin",
			},
			expectedPage: "Admin Page",
		},
		{
			name: "Should return Access Denied when user is not admin",
			args: args{
				user: "hacker",
			},
			expectedPage: "Access Denied",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			s := server.New(logger.StandardLogger{})
			// Act
			pageContent := s.RenderAdminPage(tt.args.user)
			// Assert
			require.Equal(t, tt.expectedPage, pageContent)
		})
	}
}

func TestServer_RenderAdminPageWithLogger(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctrl := gomock.NewController(t)
	l := mocks.NewMockLogger(ctrl)

	type args struct {
		user string
	}
	tests := []struct {
		name                 string
		args                 args
		expectedPage         string
		expectedInfoMessage  string
		expectedErrorMessage string
	}{
		{
			name: "Should return Admin Page when user is admin",
			args: args{
				user: "admin",
			},
			expectedPage:        "Admin Page",
			expectedInfoMessage: "User admin accessed the admin page",
		},
		{
			name: "Should return Access Denied when user is not admin",
			args: args{
				user: "hacker",
			},
			expectedPage:         "Access Denied",
			expectedErrorMessage: "User hacker tried to access the admin page",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			if tt.expectedInfoMessage != "" {
				l.EXPECT().Info(tt.expectedInfoMessage).Times(1)
			}
			if tt.expectedErrorMessage != "" {
				l.EXPECT().Error(tt.expectedErrorMessage).Times(1)
			}
			s := server.New(l)
			// Act
			pageContent := s.RenderAdminPage(tt.args.user)
			// Assert
			require.Equal(t, tt.expectedPage, pageContent)
		})
	}
}
