package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/stripe/stripe-cli/pkg/login"
	"github.com/stripe/stripe-cli/pkg/stripe"
	"github.com/stripe/stripe-cli/pkg/validators"
)

type loginCmd struct {
	cmd              *cobra.Command
	interactive      bool
	dashboardBaseURL string
}

func newLoginCmd() *loginCmd {
	lc := &loginCmd{}

	lc.cmd = &cobra.Command{
		Use:   "login",
		Args:  validators.NoArgs,
		Short: "Login to your Stripe account",
		Long:  `Login to your Stripe account to write your configuration file`,
		RunE:  lc.runLoginCmd,
	}
	lc.cmd.Flags().BoolVarP(&lc.interactive, "interactive", "i", false, "interactive configuration mode")

	// Hidden configuration flags, useful for dev/debugging
	lc.cmd.Flags().StringVar(&lc.dashboardBaseURL, "dashboard-base", stripe.DefaultDashboardBaseURL, "Sets the dashboard base URL")
	lc.cmd.Flags().MarkHidden("dashboard-base") // #nosec G104

	return lc
}

func (lc *loginCmd) runLoginCmd(cmd *cobra.Command, args []string) error {
	if lc.interactive {
		return login.InteractiveLogin(&Config)
	}

	return login.Login(lc.dashboardBaseURL, &Config, os.Stdin)
}
