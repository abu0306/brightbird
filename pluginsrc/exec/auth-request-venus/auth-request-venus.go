package main

import (
	"context"
	"fmt"

	"github.com/hunjixin/brightbird/env/plugin"

	"github.com/hunjixin/brightbird/types"

	"github.com/filecoin-project/venus-auth/auth"
	"github.com/filecoin-project/venus-auth/jwtclient"
	chain "github.com/filecoin-project/venus/venus-shared/api/chain/v1"
	"github.com/hunjixin/brightbird/env"
	"github.com/hunjixin/brightbird/utils"
	"github.com/hunjixin/brightbird/version"
	"go.uber.org/fx"
)

func main() {
	plugin.SetupPluginFromStdin(Info, Exec)
}

var Info = types.PluginInfo{
	Name:        "auth_request_venus",
	Version:     version.Version(),
	PluginType:  types.TestExec,
	Description: "auth request venus",
}

type TestCaseParams struct {
	fx.In
	Params struct {
		//Permission string `json:"permission"`
	} `optional:"true"`

	K8sEnv    *env.K8sEnvDeployer `json:"-"`
	VenusAuth env.IDeployer       `json:"-" svcname:"VenusAuth"`
	Venus     env.IDeployer       `json:"-" svcname:"Venus"`
}

func Exec(ctx context.Context, params TestCaseParams) (env.IExec, error) {
	venusAuthPods, err := params.VenusAuth.Pods(ctx)
	if err != nil {
		return nil, err
	}

	svc, err := params.VenusAuth.Svc(ctx)
	if err != nil {
		return nil, err
	}

	endpoint, err := params.VenusAuth.SvcEndpoint()
	if err != nil {
		return nil, err
	}
	if env.Debug {
		var err error
		endpoint, err = params.K8sEnv.PortForwardPod(ctx, venusAuthPods[0].GetName(), int(svc.Spec.Ports[0].Port))
		if err != nil {
			return nil, err
		}
	}

	adminToken, err := params.VenusAuth.Param("AdminToken")
	if err != nil {
		return nil, err
	}
	authAPIClient, err := jwtclient.NewAuthClient(endpoint.ToHttp(), adminToken.String())
	if err != nil {
		return nil, err
	}

	_, err = authAPIClient.CreateUser(ctx, &auth.CreateUserRequest{
		Name:    "admin",
		Comment: utils.StringPtr("comment admin"),
		State:   0,
	})
	if err != nil {
		return nil, err
	}

	token, err := authAPIClient.GenerateToken(ctx, "admin", "admin", "")
	if err != nil {
		return nil, err
	}
	fmt.Println(token)

	err = checkPermission(ctx, token, params)
	if err != nil {
		return nil, err
	}
	return env.NewSimpleExec(), nil
}

func checkPermission(ctx context.Context, token string, params TestCaseParams) error {
	endpoint, err := params.Venus.SvcEndpoint()
	if err != nil {
		return err
	}
	if env.Debug {
		venusPods, err := params.Venus.Pods(ctx)
		if err != nil {
			return err
		}

		svc, err := params.Venus.Svc(ctx)
		if err != nil {
			return err
		}
		endpoint, err = params.K8sEnv.PortForwardPod(ctx, venusPods[0].GetName(), int(svc.Spec.Ports[0].Port))
		if err != nil {
			return err
		}
	}
	chainRpc, closer, err := chain.DialFullNodeRPC(ctx, endpoint.ToMultiAddr(), token, nil)
	if err != nil {
		return err
	}
	defer closer()

	_, err = chainRpc.ChainHead(ctx)
	if err != nil {
		return err
	}

	return nil
}
