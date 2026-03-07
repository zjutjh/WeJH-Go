package register

import (
	"wejh-go/register/router"

	"github.com/spf13/cobra"
	"github.com/zjutjh/mygo/foundation/command"
	"github.com/zjutjh/mygo/foundation/httpserver"
)

func Command(root *cobra.Command) {
	command.Add("server", httpserver.CommandRegister(router.Route))
}
