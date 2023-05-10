package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"regexp"

	"github.com/hunjixin/brightbird/env/plugin"

	"github.com/filecoin-project/venus-auth/auth"
	"github.com/filecoin-project/venus-auth/jwtclient"
	"github.com/filecoin-project/venus/venus-shared/api/messager"
	"github.com/hunjixin/brightbird/env"
	"github.com/hunjixin/brightbird/types"
	"github.com/hunjixin/brightbird/utils"
	"github.com/hunjixin/brightbird/version"
	"go.uber.org/fx"
)

func main() {
	plugin.SetupPluginFromStdin(Info, Exec)
}

var Info = types.PluginInfo{
	Name:        "message-broadcast",
	Version:     version.Version(),
	PluginType:  types.TestExec,
	Description: "verity message if normal",
}

type TestCaseParams struct {
	fx.In
	K8sEnv       *env.K8sEnvDeployer `json:"-"`
	VenusAuth    env.IDeployer       `json:"-" svcname:"VenusAuth"`
	VenusMessage env.IDeployer       `json:"-" svcname:"VenusMessage"`
}

func Exec(ctx context.Context, params TestCaseParams) (env.IExec, error) {
	authToken, err := CreateAuthToken(ctx, params)
	if err != nil {
		fmt.Printf("create auth token failed: %v\n", err)
		return nil, err
	}

	err = CreateMessage(ctx, params, authToken)
	if err != nil {
		fmt.Printf("get message failed: %v\n", err)
		return nil, err
	}

	return env.NewSimpleExec(), nil
}

func CreateAuthToken(ctx context.Context, params TestCaseParams) (string, error) {
	adminToken, err := params.VenusAuth.Param("AdminToken")
	if err != nil {
		return "", err
	}

	endpoint, err := params.VenusAuth.SvcEndpoint()
	if err != nil {
		return "", err
	}

	if env.Debug {
		pods, err := params.VenusAuth.Pods(ctx)
		if err != nil {
			return "", err
		}

		svc, err := params.VenusAuth.Svc(ctx)
		if err != nil {
			return "", err
		}
		endpoint, err = params.K8sEnv.PortForwardPod(ctx, pods[0].GetName(), int(svc.Spec.Ports[0].Port))
		if err != nil {
			return "", err
		}
	}

	authAPIClient, err := jwtclient.NewAuthClient(endpoint.ToHttp(), adminToken.String())
	if err != nil {
		return "", err
	}
	_, err = authAPIClient.CreateUser(ctx, &auth.CreateUserRequest{
		Name:    "admin",
		Comment: utils.StringPtr("comment admin"),
		State:   0,
	})
	if err != nil {
		return "", err
	}

	return authAPIClient.GenerateToken(ctx, "admin", "admin", "")
}

func CreateMessage(ctx context.Context, params TestCaseParams, authToken string) error {
	pods, err := params.VenusMessage.Pods(ctx)
	if err != nil {
		return err
	}

	svc, err := params.VenusMessage.Svc(ctx)
	if err != nil {
		return err
	}

	endpoint, err := params.VenusMessage.SvcEndpoint()
	if err != nil {
		return err
	}

	if env.Debug {
		var err error
		endpoint, err = params.K8sEnv.PortForwardPod(ctx, pods[0].GetName(), int(svc.Spec.Ports[0].Port))
		if err != nil {
			return err
		}
	}

	client, closer, err := messager.DialIMessagerRPC(ctx, endpoint.ToHttp(), authToken, nil)
	if err != nil {
		return err
	}
	defer closer()

	// Get message IDs.
	precommitID, procommitID, wdpostID, faultID, err := readLogForMsgIds()
	if err != nil {
		return fmt.Errorf("failed to get message IDs: %v", err)
	}

	_, err = client.GetMessageByUid(ctx, precommitID)
	if err != nil {
		return err
	}

	_, err = client.GetMessageByUid(ctx, procommitID)
	if err != nil {
		return err
	}

	_, err = client.GetMessageByUid(ctx, wdpostID)
	if err != nil {
		return err
	}

	_, err = client.GetMessageByUid(ctx, faultID)
	if err != nil {
		return err
	}

	return nil
}

func readLogForMsgIds() (string, string, string, string, error) {
	file, err := os.Open("log.txt")
	if err != nil {
		return "", "", "", "", err
	}
	defer file.Close()

	var msgPrecommitId, msgProcommitId, msgWdpostId, msgFaultId string
	rePrecommit := regexp.MustCompile(`"stage": "pre-commit", "msg-cid": "(\w+)"`)
	reProcommit := regexp.MustCompile(`"stage": "prove-commit", "msg-cid": "(\w+)"`)
	reWdpost := regexp.MustCompile(`Submitted window post: (\w+)`)
	reFault := regexp.MustCompile(`declare faults recovered message published\s+{"message-id": "(\w+)"}`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if match := rePrecommit.FindStringSubmatch(line); len(match) > 0 {
			msgPrecommitId = match[1]
		} else if match := reProcommit.FindStringSubmatch(line); len(match) > 0 {
			msgProcommitId = match[1]
		} else if match := reWdpost.FindStringSubmatch(line); len(match) > 0 {
			msgWdpostId = match[1]
		} else if match := reFault.FindStringSubmatch(line); len(match) > 0 {
			msgFaultId = match[1]
		}
		if msgPrecommitId != "" && msgProcommitId != "" && msgWdpostId != "" && msgFaultId != "" {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return "", "", "", "", err
	}

	fmt.Printf("msg_precommit_id: %v\nmsg_procommit_id: %v\nmsg_wdpost_id: %v\nmsg_fault_id: %v\n",
		msgPrecommitId, msgProcommitId, msgWdpostId, msgFaultId)
	return msgPrecommitId, msgProcommitId, msgWdpostId, msgFaultId, nil
}
