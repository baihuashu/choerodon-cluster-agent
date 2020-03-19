package gitops

import (
	"context"
	"errors"
	"github.com/choerodon/choerodon-cluster-agent/pkg/agent/model"
	"github.com/choerodon/choerodon-cluster-agent/pkg/util/command"
)

func DoSync(opts *command.Opts, cmd *model.Packet) ([]*model.Packet, *model.Packet) {
	ctx, cancel := context.WithTimeout(context.Background(), opts.GitTimeout)
	if !opts.Namespaces.Contain(cmd.Namespace()) {
		return nil, command.NewResponseError(cmd.Key, model.GitOpsSyncFailed, errors.New("cluster env not init"))
	}
	if opts.GitRepos[cmd.Namespace()] == nil {

		return nil, command.NewResponseError(cmd.Key, model.GitOpsSyncFailed, errors.New("git repo not init"))
	}
	err := opts.GitRepos[cmd.Namespace()].Refresh(ctx)
	cancel()
	if err != nil {
		return nil, command.NewResponseError(cmd.Key, model.GitOpsSyncFailed, err)
	}
	return nil, &model.Packet{
		Key:  cmd.Key,
		Type: model.GitOpsSync,
	}
}
