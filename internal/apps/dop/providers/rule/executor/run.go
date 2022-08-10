package executor

import (
	"encoding/json"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/internal/apps/dop/providers/rule/actions/api"
	"github.com/erda-project/erda/pkg/parser/pipelineyml"
)

func buildPipelineYml(c *RuleConfig) (string, error) {
	var stage pipelineyml.Stage
	for _, a := range c.APIConfigs {
		action, err := buildRuleAction(c, a)
		if err != nil {
			return "", err
		}
		stage.Actions = append(stage.Actions, action)
	}

	spec := pipelineyml.Spec{
		Version: "1.1",
		Stages:  []*pipelineyml.Stage{&stage},
	}
	yml, err := pipelineyml.GenerateYml(&spec)
	if err != nil {
		return "", err
	}
	return string(yml), nil
}

func buildRuleAction(config *RuleConfig, apiConfig *api.APIConfig) (map[pipelineyml.ActionType]*pipelineyml.Action, error) {
	info := apistructs.APIInfoV2{
		URL:    apiConfig.URL,
		Method: apiConfig.Method,
	}
	for i, v := range apiConfig.Headers {
		info.Headers = append(info.Headers, apistructs.APIHeader{
			Key:   i,
			Value: v[0],
		})
	}

	b, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}
	var params map[string]interface{}
	if err := json.Unmarshal(b, &params); err != nil {
		return nil, err
	}
	action := pipelineyml.Action{
		Alias:   pipelineyml.ActionAlias(config.RuleID),
		Version: "2.0",
		Type:    "api-test",
		Params:  params,
	}

	return map[pipelineyml.ActionType]*pipelineyml.Action{
		action.Type: &action,
	}, nil
}

// func (e *Executor) ExecuteAPIByAPITest(c *RuleConfig) (*apistructs.PipelineDTO, error) {
// 	yml, err := buildPipelineYml(c)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req := apistructs.PipelineCreateRequest{
// 		PipelineYmlName:   c.RuleID,
// 		PipelineYmlSource: apistructs.PipelineSourceRuleAPI,
// 	}
// }
