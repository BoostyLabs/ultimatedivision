// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package main

import (
	"context"
	"encoding/json"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/zeebo/errs"

	"ultimatedivision/currencysigner"
	"ultimatedivision/database"
	"ultimatedivision/internal/logger/zaplog"
	"ultimatedivision/pkg/fileutils"
)

// Error is a default error type for card currency signer cli.
var Error = errs.Class("currency signer cli error")

// Config contains configurable values for card generator project.
type Config struct {
	Database              string `json:"database"`
	currencysigner.Config `json:"config"`
}

// commands.
var (
	// currency signer root cmd.
	rootCmd = &cobra.Command{
		Use:   "signer",
		Short: "cli for interacting with currency signer project",
	}

	// currency signer setup cmd.
	setupCmd = &cobra.Command{
		Use:         "setup",
		Short:       "setups the program config",
		RunE:        cmdSetup,
		Annotations: map[string]string{"type": "setup"},
	}
	runCmd = &cobra.Command{
		Use:         "run",
		Short:       "runs the program",
		RunE:        cmdRun,
		Annotations: map[string]string{"type": "run"},
	}
	destroyCmd = &cobra.Command{
		Use:         "destroy",
		Short:       "deletes config folder",
		RunE:        cmdDestroy,
		Annotations: map[string]string{"type": "run"},
	}
	setupCfg Config
	runCfg   Config

	defaultConfigDir = fileutils.ApplicationDir(filepath.Join("ultimatedivision", "currencysigner"))
)

func init() {
	rootCmd.AddCommand(setupCmd)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(destroyCmd)
	rootCmd.PersistentFlags().StringVar(&defaultConfigDir, "config", defaultConfigDir, "Config file path")
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func cmdSetup(cmd *cobra.Command, args []string) (err error) {
	log := zaplog.NewLog()

	setupDir, err := filepath.Abs(defaultConfigDir)
	if err != nil {
		return Error.Wrap(err)
	}

	err = os.MkdirAll(setupDir, os.ModePerm)
	if err != nil {
		return Error.Wrap(err)
	}

	configFile, err := os.Create(path.Join(setupDir, "config.json"))
	if err != nil {
		log.Error("could not create config file", Error.Wrap(err))
		return Error.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, configFile.Close())
	}()

	jsonData, err := json.MarshalIndent(setupCfg, "", "    ")
	if err != nil {
		log.Error("could not marshal config", Error.Wrap(err))
		return Error.Wrap(err)
	}

	_, err = configFile.Write(jsonData)
	if err != nil {
		log.Error("could not write to config", Error.Wrap(err))
		return Error.Wrap(err)
	}

	return nil
}

func cmdRun(cmd *cobra.Command, args []string) (err error) {
	ctx := context.Background()
	log := zaplog.NewLog()

	runCfg, err = readConfig()
	if err != nil {
		log.Error("Could not read config from default place", Error.Wrap(err))
		return Error.Wrap(err)
	}

	db, err := database.New(runCfg.Database)
	if err != nil {
		log.Error("Error starting master database on ultimatedivision bank service", Error.Wrap(err))
		return Error.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, db.Close())
	}()

	err = db.CreateSchema(ctx)
	if err != nil {
		log.Error("Error creating schema", Error.Wrap(err))
	}

	peer, err := currencysigner.New(log, runCfg.Config, db)
	if err != nil {
		log.Error("Error starting currencysigner bank service", Error.Wrap(err))
		return Error.Wrap(err)
	}

	runError := peer.Run(ctx)

	return Error.Wrap(runError)
}

func cmdDestroy(cmd *cobra.Command, args []string) (err error) {
	return os.RemoveAll(defaultConfigDir)
}

// readConfig reads config from default config dir.
func readConfig() (config Config, err error) {
	configBytes, err := os.ReadFile(path.Join(defaultConfigDir, "config.json"))
	if err != nil {
		return Config{}, err
	}

	return config, json.Unmarshal(configBytes, &config)
}
