package cmd

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/hirosassa/bqiam/metadata"
	"github.com/spf13/cobra"
)

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web [port]",
	Short: "run web ui use input port",
	Long: `
This subcommand run web server on input port.
that show user or service account is able to access.
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("allow only port")
		}
		return nil
	},
	RunE: runCmdWeb,
}

func runCmdWeb(cmd *cobra.Command, args []string) error {
	port := "3000"
	if len(args) > 0 {
		port = args[0]
	}

	http.HandleFunc("/", createHandler(cmd))

	http.ListenAndServe(":"+port, nil)
	return nil
}

func createHandler(cmd *cobra.Command) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		user := r.URL.Query().Get("user")
		if user == "" {
			fmt.Fprintf(w, "Please input user")
			return
		}
		refreshCache(cmd) // refresh cache if needed

		var ms metadata.Metas
		if err := ms.Load(config.CacheFile); err != nil {
			fmt.Fprintf(w, "CacheFile load error")
			return
		}

		for _, m := range ms.Metas {
			if m.Entity == user {
				fmt.Fprintln(w, m.Project, m.Dataset, m.Role)
			}
		}
	}

}

func init() {
	rootCmd.AddCommand(webCmd)
}
