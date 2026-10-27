package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/Pwatermelon/owl-ontology-versioning/internal/compare"
	"github.com/Pwatermelon/owl-ontology-versioning/internal/ui"
)

// VERSION подставляется при сборке: -ldflags "-X main.VERSION=1.0.0"
var VERSION = "0.3.0"

func main() {
	root := &cobra.Command{
		Use:   "owl-versioning",
		Short: "CLI версионирования OWL-онтологий: preview / diff двух снимков",
		Long: `CLI для версионирования OWL-онтологий: preview / diff двух снимков.

Примеры:
  owl-versioning preview base.owl incoming.owl
  owl-versioning preview base.owl incoming.owl --json-stdout

Бинарники Linux / Windows / macOS публикуются в GitHub Releases.`,
	}

	var (
		jsonPath   string
		limit      int
		jsonStdout bool
	)

	previewCmd := &cobra.Command{
		Use:   "preview <base.owl> <incoming.owl>",
		Short: "Сравнить две версии онтологии (added / removed / conflicts)",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			report, err := compare.CompareFiles(args[0], args[1])
			if err != nil {
				ui.RenderError(err.Error())
				return err
			}

			if jsonStdout || jsonPath != "" {
				raw, err := report.JSON()
				if err != nil {
					return err
				}
				if jsonPath != "" {
					if err := os.WriteFile(jsonPath, raw, 0o644); err != nil {
						return err
					}
					if !jsonStdout {
						fmt.Fprintf(os.Stdout, "JSON сохранён: %s\n", jsonPath)
					}
				}
				if jsonStdout {
					fmt.Println(string(raw))
					return nil
				}
			}

			ui.RenderReport(report, limit)
			return nil
		},
	}
	previewCmd.Flags().StringVar(&jsonPath, "json", "", "Записать полный отчёт в JSON-файл")
	previewCmd.Flags().BoolVar(&jsonStdout, "json-stdout", false, "Печать JSON в stdout (для HTTP-обёртки)")
	previewCmd.Flags().IntVar(&limit, "limit", 15, "Сколько строк сущностей показать в TUI")

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Версия CLI",
		Run: func(cmd *cobra.Command, args []string) {
			ui.BannerVersion(VERSION)
		},
	}

	root.AddCommand(previewCmd, versionCmd)
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
