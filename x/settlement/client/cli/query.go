package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/Finschia/finschia-sdk/client"
	// "github.com/Finschia/finschia-sdk/client/flags"
	// sdk "github.com/Finschia/finschia-sdk/types"

	"github.com/Finschia/finschia/x/settlement/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group settlement queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	// this line is used by starport scaffolding # 1

	return cmd
}
