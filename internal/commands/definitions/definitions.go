package definitions

import (
	"github.com/spf13/cobra"
)

type Command interface {
	CMD() *cobra.Command
}
