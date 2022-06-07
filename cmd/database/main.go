// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zeebo/errs"

	"ultimatedivision/database"
	"ultimatedivision/internal/logger/zaplog"
	"ultimatedivision/pkg/fileutils"
)

// Error is a default error type for database cli.
var Error = errs.Class("database cli error")

// Config contains configurable values for card generator project.
type Config struct {
	Database       string `json:"database"`
	MigrationsPath string `json:"migrationsPath"`
}

// commands.
var (
	// database root cmd.
	rootCmd = &cobra.Command{
		Use:   "database",
		Short: "cli for interacting with database project",
	}

	// create database schema.
	createMigrationCmd = &cobra.Command{
		Use:         "create-migration",
		Short:       "creates a new migration",
		RunE:        cmdCreateMigration,
		Annotations: map[string]string{"type": "run"},
	}

	// execute database migrations.
	migrateCmd = &cobra.Command{
		Use:         "migrate",
		Short:       "executes migrations",
		RunE:        cmdMigrate,
		Annotations: map[string]string{"type": "run"},
	}

	runCfg     Config
	configPath = fileutils.ApplicationDir(filepath.Join("ultimatedivision", "database"))
)

func init() {
	rootCmd.AddCommand(createMigrationCmd)
	rootCmd.AddCommand(migrateCmd)
	rootCmd.PersistentFlags().StringVar(&configPath, "config", "", "config file path")
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// cmdCreateSchema creates schema for all tables and databases.
func cmdCreateMigration(cmd *cobra.Command, args []string) (err error) {
	log := zaplog.NewLog()

	if len(args) == 0 {
		log.Error("migration name is required", Error.New("invalid arguments"))
		return Error.New("invalid arguments")
	}

	runCfg, err = readConfig()
	if err != nil {
		log.Error("Could not read config from default place", Error.Wrap(err))
		return Error.Wrap(err)
	}

	fExt := ".sql$"
	curVer := 0

	files, err := ioutil.ReadDir(runCfg.MigrationsPath)
	if err != nil {
		log.Error("Could not read config dir", Error.Wrap(err))
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		fName := f.Name()
		r, err := regexp.MatchString(fExt, fName)
		if err != nil {
			log.Error("Could not process existing transaction file names", Error.Wrap(err))
			return Error.Wrap(err)
		}
		if r == false {
			continue
		}

		parts := strings.Split(fName, "_")
		ver, err := strconv.Atoi(parts[0])
		if err != nil {
			// Looks like that file name is without a numeric prefix
			continue
		}
		if ver > curVer {
			curVer = ver
		}
	}

	migName := fmt.Sprintf("%06d_%s", curVer+1, args[0])
	fNames := [2]string{
		migName + ".up.sql",
		migName + ".down.sql",
	}
	for _, fName := range fNames {
		isExist, err := isFileExist(runCfg.MigrationsPath, fName)
		if err != nil {
			log.Error("Could not check file existence ", Error.Wrap(err))
			return Error.Wrap(err)
		}
		if isExist {
			errMsg := fmt.Sprintf("File '%s' is already exists", fName)
			log.Error(errMsg, Error.New("file exists"))
			return Error.New("file exists")
		}
	}

	for _, fName := range fNames {
		if err := createFile(runCfg.MigrationsPath, fName); err != nil {
			errMsg := fmt.Sprintf("Could not crate file '%s'", fName)
			log.Error(errMsg, Error.Wrap(err))
		} else {
			fmt.Printf("New file: %s\n", fName)
		}
	}

	return nil
}

// cmdMigrate executes migrations by path in database.
func cmdMigrate(cmd *cobra.Command, args []string) (err error) {
	var isUp bool
	ctx := context.Background()
	log := zaplog.NewLog()

	if len(args) == 0 {
		log.Error("at least 1 argument is required", Error.New("invalid arguments"))
		return Error.New("invalid arguments")
	}

	runCfg, err = readConfig()
	if err != nil {
		log.Error("Could not read config from default place", Error.Wrap(err))
		return Error.Wrap(err)
	}

	db, err := database.New(runCfg.Database)
	if err != nil {
		log.Error("Error starting master database", Error.Wrap(err))
		return Error.Wrap(err)
	}
	defer func() {
		err = Error.Wrap(errs.Combine(err, db.Close()))
	}()

	switch args[0] {
	case "up":
		isUp = true
		if err = db.ExecuteMigrations(ctx, runCfg.MigrationsPath, isUp); err != nil {
			log.Error("Error migrations up", Error.Wrap(err))
		}
	case "down":
		isUp = false
		if err = db.ExecuteMigrations(ctx, runCfg.MigrationsPath, isUp); err != nil {
			log.Error("Error migrations down", Error.Wrap(err))
		}
	default:
		err = errs.New("invalid arguments")
		log.Error("Invalid 1 argument", Error.Wrap(err))
	}

	return Error.Wrap(err)
}

// readConfig reads config from default config dir.
func readConfig() (config Config, err error) {
	configBytes, err := ioutil.ReadFile(path.Join(configPath, "/config.json"))
	if err != nil {
		return Config{}, err
	}

	return config, json.Unmarshal(configBytes, &config)
}

// isFileExist checks if file with given name exists in path
func isFileExist(path, fName string) (bool, error) {
	name := path
	if name[len(name)-1:] != "/" {
		name += "/"
	}

	name += fName
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

// createFile creates a new file in path
func createFile(path, fName string) error {
	name := path
	if name[len(name)-1:] != "/" {
		name += "/"
	}

	name += fName
	f, err := os.Create(name)
	if err != nil {
		return err
	}

	_ = f.Close()
	return nil
}
