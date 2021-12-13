// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zeebo/errs"

	"ultimatedivision"
	"ultimatedivision/database"
	"ultimatedivision/internal/logger/zaplog"
)

// Error is a default error type for ultimatedivision cli.
var Error = errs.Class("ultimatedivision cli error")

// Config contains configurable values for ultimatedivision project.
type Config struct {
	Database                string `json:"database"`
	ultimatedivision.Config `json:"config"`
}

// commands.
var (
	// ultimatedivision root cmd.
	rootCmd = &cobra.Command{
		Use:   "ultimatedivision",
		Short: "cli for interacting with ultimatedivision project",
	}

	// ultimatedivision setup cmd.
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
	seedCmd = &cobra.Command{
		Use:         "seed",
		Short:       "created test data in DB",
		RunE:        seedRun,
		Annotations: map[string]string{"type": "run"},
	}
	matchCmd = &cobra.Command{
		Use:         "match",
		Short:       "play matches",
		RunE:        matchRun,
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

	defaultConfigDir = ApplicationDir("ultimatedivision")
)

func init() {
	rootCmd.AddCommand(setupCmd)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(seedCmd)
	rootCmd.AddCommand(matchCmd)
	rootCmd.AddCommand(destroyCmd)
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

	//// TODO: remove for production.
	//err = db.CreateSchema(ctx)
	//if err != nil {
	//	log.Error("Error creating schema", Error.Wrap(err))
	//}

	peer, err := ultimatedivision.New(log, runCfg.Config, db)
	if err != nil {
		log.Error("Error starting ultimatedivision bank service", Error.Wrap(err))
		return Error.Wrap(err)
	}

	runError := peer.Run(ctx)
	closeError := peer.Close()

	return Error.Wrap(errs.Combine(runError, closeError))
}

// TODO: remove for production.
func seedRun(cmd *cobra.Command, args []string) (err error) {
	ctx := context.Background()
	runCfg, err = readConfig()
	if err != nil {
		return Error.Wrap(err)
	}
	conn, err := sql.Open("postgres", runCfg.Database)
	if err != nil {
		return Error.Wrap(err)
	}

	seedDB := database.NewSeedDB(conn)

	err = database.CreateUser(ctx, conn)
	if err != nil {
		return Error.Wrap(err)
	}
	err = database.CreateAdmin(ctx, conn)
	if err != nil {
		return Error.Wrap(err)
	}
	err = database.CreateDivisions(ctx, conn)
	if err != nil {
		return Error.Wrap(err)
	}
	err = database.CreateClubs(ctx, conn)
	if err != nil {
		return Error.Wrap(err)
	}
	err = database.CreateSquads(ctx, conn)
	if err != nil {
		return Error.Wrap(err)
	}
	err = seedDB.CreateSquadCards(ctx, conn, runCfg.Cards.Config, runCfg.LootBoxes.Config)
	if err != nil {
		return Error.Wrap(err)
	}

	return nil
}

// TODO: remove for production.
func matchRun(cmd *cobra.Command, args []string) (err error) {
	ctx := context.Background()
	runCfg, err = readConfig()
	if err != nil {
		return Error.Wrap(err)
	}
	conn, err := sql.Open("postgres", runCfg.Database)
	if err != nil {
		return Error.Wrap(err)
	}

	seedDB := database.NewSeedDB(conn)

	err = seedDB.CreateMatches(ctx, conn, runCfg.Matches.Config, runCfg.Cards.Config)
	if err != nil {
		return Error.Wrap(err)
	}

	return nil
}

func cmdDestroy(cmd *cobra.Command, args []string) (err error) {
	return os.RemoveAll(defaultConfigDir)
}

// readConfig reads config from default config dir.
func readConfig() (config Config, err error) {
	configBytes, err := ioutil.ReadFile("./config.json")
	if err != nil {
		return Config{}, err
	}

	return config, json.Unmarshal(configBytes, &config)
}

// ApplicationDir returns best base directory for specific OS.
func ApplicationDir(subdir ...string) string {
	for i := range subdir {
		if runtime.GOOS == "windows" || runtime.GOOS == "darwin" {
			subdir[i] = strings.Title(subdir[i])
		} else {
			subdir[i] = strings.ToLower(subdir[i])
		}
	}

	var appdir string

	home := os.Getenv("HOME")
	//
	switch runtime.GOOS {
	case "windows":
		// Windows standards: https://msdn.microsoft.com/en-us/library/windows/apps/hh465094.aspx?f=255&MSPPError=-2147217396
		for _, env := range []string{"AppData", "AppDataLocal", "UserProfile", "Home"} {
			val := os.Getenv(env)
			if val != "" {
				appdir = val
				break
			}
		}
	case "darwin":
		// Mac standards: https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/MacOSXDirectories/MacOSXDirectories.html
		appdir = filepath.Join(home, "Library", "Application Support")
	case "linux":
		fallthrough
	default:
		// Linux standards: https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html
		appdir = os.Getenv("XDG_DATA_HOME")
		if appdir == "" && home != "" {
			appdir = filepath.Join(home, ".local", "share")
		}
	}

	return filepath.Join(append([]string{appdir}, subdir...)...)
}
