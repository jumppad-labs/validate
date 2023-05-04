package cmd

import (
	"encoding/json"
	"net/http"

	"github.com/jumppad-labs/validate/data"
	"github.com/kr/pretty"
	"github.com/oliveagle/jsonpath"
	"github.com/spf13/cobra"
)

var jsonCmd = &cobra.Command{
	Use: "json",
	Run: func(cmd *cobra.Command, args []string) {

		resp, err := http.DefaultClient.Get("https://api.releases.hashicorp.com/v1/releases/terraform/latest")
		if err != nil {
			data.Error(err)
		}

		// body, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	data.Error(err)
		// }

		var v interface{}
		json.NewDecoder(resp.Body).Decode(&v)

		pat, _ := jsonpath.Compile(`$.builds[*].[?(@.arch == "amd64")]`)
		res, err := pat.Lookup(v)
		if err != nil {
			data.Error(err)
		}

		// value, err := jsonpath.JsonPathLookup(v, `$.builds[?@.arch=="amd64"]`)
		// if err != nil {
		// data.Error(err)
		// }

		pretty.Println(res)
	},
}
