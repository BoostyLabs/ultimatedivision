// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/zeebo/errs"

	"ultimatedivision/cardgenerator"
	"ultimatedivision/internal/logger/zaplog"
	"ultimatedivision/pkg/fileutils"
)

// Error is a default error type for card generator cli.
var Error = errs.Class("card generator cli error")

// Config contains configurable values for card generator project.
type Config struct {
	cardgenerator.Config `json:"config"`
}

// commands.
var (
	// card generator root cmd.
	rootCmd = &cobra.Command{
		Use:   "generator",
		Short: "cli for interacting with card generator project",
	}

	// card generator setup cmd.
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
	testCmd = &cobra.Command{
		Use:         "test",
		Short:       "tests the program",
		RunE:        testRun,
		Annotations: map[string]string{"type": "test"},
	}
	destroyCmd = &cobra.Command{
		Use:         "destroy",
		Short:       "deletes config folder",
		RunE:        cmdDestroy,
		Annotations: map[string]string{"type": "run"},
	}
	setupCfg Config
	runCfg   Config

	defaultConfigDir = fileutils.ApplicationDir(filepath.Join("ultimatedivision", "cardgenerator"))
)

func init() {
	rootCmd.AddCommand(setupCmd)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(testCmd)
	rootCmd.AddCommand(destroyCmd)
	rootCmd.PersistentFlags().StringVar(&defaultConfigDir, "config", "", "Config file path")
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

	if len(args) == 0 {
		log.Error("agrs are empty", Error.New("agrs are empty"))
		return Error.New("agrs are empty")
	}

	count, err := strconv.Atoi(args[0])
	if err != nil {
		log.Error("Error convert agrs to integer", Error.Wrap(err))
		return Error.Wrap(err)
	}

	peer, err := cardgenerator.New(log, runCfg.Config, count)
	if err != nil {
		log.Error("Error starting card generator bank service", Error.Wrap(err))
		return Error.Wrap(err)
	}

	if err := peer.Generate(ctx); err != nil {
		log.Error("could not generate cards with avatars", Error.Wrap(err))
		return Error.Wrap(err)
	}

	return nil
}

func testRun(cmd *cobra.Command, args []string) (err error) {
	ctx := context.Background()
	log := zaplog.NewLog()

	runCfg, err = readConfig()
	if err != nil {
		log.Error("Could not read config from default place", Error.Wrap(err))
		return Error.Wrap(err)
	}

	if len(args) == 0 {
		log.Error("agrs are empty", Error.New("agrs are empty"))
		return Error.New("agrs are empty")
	}

	count, err := strconv.Atoi(args[0])
	if err != nil {
		log.Error("Error convert agrs to integer", Error.Wrap(err))
		return Error.Wrap(err)
	}

	peer, err := cardgenerator.New(log, runCfg.Config, count)
	if err != nil {
		log.Error("Error starting card generator bank service", Error.Wrap(err))
		return Error.Wrap(err)
	}

	if err := peer.TestGenerate(ctx); err != nil {
		log.Error("could not generate cards with avatars", Error.Wrap(err))
		return Error.Wrap(err)
	}

	return nil
}

func cmdDestroy(cmd *cobra.Command, args []string) (err error) {
	return os.RemoveAll(defaultConfigDir)
}

// readConfig reads config from default config dir.
func readConfig() (config Config, err error) {
	configBytes, err := ioutil.ReadFile(path.Join(defaultConfigDir, "config.json"))
	if err != nil {
		return Config{}, err
	}

	return config, json.Unmarshal(configBytes, &config)
}
