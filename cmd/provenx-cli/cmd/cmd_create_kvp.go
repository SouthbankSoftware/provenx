/*
 * @Author: guiguan
 * @Date:   2019-09-16T16:21:53+10:00
 * @Last modified by:   guiguan
 * @Last modified time: 2020-03-18T15:01:18+11:00
 */

package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/SouthbankSoftware/provenx-cli/pkg/api"
	"github.com/SouthbankSoftware/provenx-cli/pkg/colorcli"
	apiPB "github.com/SouthbankSoftware/provenx-cli/pkg/protos/api"
	"github.com/SouthbankSoftware/provenx-cli/pkg/strutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	nameTriePath = "trie-path"

	viperKeyCreateKvpTriePath   = nameCreate + "." + nameKvp + "." + nameTriePath
	viperKeyCreateKvpOutputPath = nameCreate + "." + nameKvp + "." + nameOutputPath
)

var cmdCreateKvp = &cobra.Command{
	Use:   fmt.Sprintf("%v <key ...>", nameKvp),
	Short: "Create a key-values proof",
	Long: fmt.Sprintf(`Create a key-values proof (%v) out from the given trie (%v). The proof proves a subset of key-values of the trie independently

Each <key> must be a valid key from the output of "%s/%s %s"
`, api.FileExtensionKeyValuesProof, api.FileExtensionTrie, nameCreate, nameVerify, nameTrie),
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// from this point, we should silence usage if error happens
		cmd.SilenceUsage = true

		keyMap := map[string]struct{}{}
		filter := &apiPB.KeyValuesFilter{}

		for _, a := range args {
			if _, ok := keyMap[a]; ok {
				// already exists, skip
				continue
			}
			keyMap[a] = struct{}{}

			colorcli.Printf("%s\n", a)

			filter.Keys = append(filter.Keys, &apiPB.Key{
				Key: api.NormalizeKey(strutil.Bytes(a)),
			})
		}

		triePath := viper.GetString(viperKeyCreateKvpTriePath)
		_, err := os.Stat(triePath)
		if err != nil {
			return err
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		creds, err := getCreds(ctx)
		if err != nil {
			return err
		}

		return api.WithAPIClient(
			viper.GetString(viperKeyAPIHostPort),
			creds,
			func(cli apiPB.APIServiceClient) error {
				return api.WithImportedTrie(ctx, cli, "", triePath,
					func(id, root string) error {
						tp, err := api.GetTrieProof(ctx, cli, id, "", root)
						if err != nil {
							return err
						}

						kvpOutputPath := viper.GetString(viperKeyCreateKvpOutputPath)

						err = api.CreateKeyValuesProof(ctx, cli, id, tp.GetId(), filter,
							kvpOutputPath)
						if err != nil {
							return err
						}

						colorcli.Oklnf("the key-values proof has successfully been created at %s with %s key-values and merkle root %s, which is anchored to %s in block %v with transaction %s at %s, which can be viewed at %s",
							colorcli.Green(kvpOutputPath),
							colorcli.Green(len(filter.Keys), " or more"),
							colorcli.Green(tp.GetProofRoot()),
							colorcli.Green(tp.GetAnchorType()),
							colorcli.Green(tp.GetBlockNumber()),
							colorcli.Green(tp.GetTxnId()),
							colorcli.Green(time.Unix(int64(tp.GetBlockTime()), 0).Format(time.UnixDate)),
							tp.GetTxnUri())

						return nil
					})
			})
	},
}

func init() {
	cmdCreate.AddCommand(cmdCreateKvp)

	cmdCreateKvp.Flags().StringP(nameTriePath, "t", "", "specify the trie path")
	err := cmdCreateKvp.MarkFlagRequired(nameTriePath)
	if err != nil {
		panic(err)
	}
	viper.BindPFlag(viperKeyCreateKvpTriePath, cmdCreateKvp.Flags().Lookup(nameTriePath))

	cmdCreateKvp.Flags().StringP(nameOutputPath, "p",
		defaultKvpPath, "specify the proof output path")
	viper.BindPFlag(viperKeyCreateKvpOutputPath, cmdCreateKvp.Flags().Lookup(nameOutputPath))
}
