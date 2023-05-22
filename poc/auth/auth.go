package auth

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"

	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
	authzed "github.com/authzed/authzed-go/v1"
	"github.com/authzed/grpcutil"
	"github.com/spf13/viper"
)

var (
	//go:embed permissions/*
	permissions embed.FS
)

func Migrations(ctx context.Context) error {
	viper.SetDefault("SPICE_DB_API", "shortlink.auth:50051")
	viper.SetDefault("SPICE_DB_COMMON_KEY", "secret-shortlink-preshared-key")

	client, err := authzed.NewClient(
		viper.GetString("SPICE_DB_API"),
		// NOTE: For SpiceDB behind TLS, use:
		// grpcutil.WithBearerToken("secret-shortlink-preshared-key"),
		// grpcutil.WithSystemCerts(grpcutil.VerifyCA),
		grpcutil.WithInsecureBearerToken(viper.GetString("SPICE_DB_COMMON_KEY")),
	)
	if err != nil {
		return err
	}

	permissionsData, err := GetPermissions(permissions)
	if err != nil {
		return err
	}

	for _, schema := range permissionsData {
		request := &pb.WriteSchemaRequest{Schema: schema}
		_, err = client.WriteSchema(ctx, request)
		if err != nil {
			return fmt.Errorf("Failed to write schema: %w", err)
		}
	}

	return nil
}

func GetPermissions(fsys fs.FS) ([][]byte, error) {
	var zedFileData [][]byte

	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && filepath.Ext(d.Name()) == ".zed" {
			fileData, err := fs.ReadFile(fsys, path)
			if err != nil {
				return fmt.Errorf("Failed to read file: %w", err)
			}

			zedFileData = append(zedFileData, fileData)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("Failed to walk directory: %w", err)
	}

	return zedFileData, nil
}