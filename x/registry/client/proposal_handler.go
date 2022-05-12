package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"kyve/x/registry/client/cli"
	"kyve/x/registry/client/rest"
)

var CreatePoolHandler = govclient.NewProposalHandler(cli.CmdSubmitCreatePoolProposal, rest.ProposalCreatePoolRESTHandler)
var UpdatePoolHandler = govclient.NewProposalHandler(cli.CmdSubmitUpdatePoolProposal, rest.ProposalUpdatePoolRESTHandler)
var PausePoolHandler = govclient.NewProposalHandler(cli.CmdSubmitPausePoolProposal, rest.ProposalPausePoolRESTHandler)
var UnpausePoolHandler = govclient.NewProposalHandler(cli.CmdSubmitUnpausePoolProposal, rest.ProposalUnpausePoolRESTHandler)
