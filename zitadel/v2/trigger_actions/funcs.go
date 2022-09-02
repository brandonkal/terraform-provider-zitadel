package trigger_actions

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zitadel/zitadel-go/v2/pkg/client/zitadel/action"
	"github.com/zitadel/zitadel-go/v2/pkg/client/zitadel/management"

	"github.com/zitadel/terraform-provider-zitadel/zitadel/v2/helper"
)

func delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Info(ctx, "started delete")

	clientinfo, ok := m.(*helper.ClientInfo)
	if !ok {
		return diag.Errorf("failed to get client")
	}

	client, err := helper.GetManagementClient(clientinfo, d.Get(orgIDVar).(string))
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = client.SetTriggerActions(ctx, &management.SetTriggerActionsRequest{
		FlowType:    action.FlowType(action.FlowType_value[d.Get(flowTypeVar).(string)]),
		TriggerType: action.TriggerType(action.TriggerType_value[d.Get(triggerTypeVar).(string)]),
		ActionIds:   []string{},
	})
	if err != nil {
		return diag.Errorf("failed to delete trigger actions: %v", err)
	}
	return nil
}

func update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Info(ctx, "started update")

	clientinfo, ok := m.(*helper.ClientInfo)
	if !ok {
		return diag.Errorf("failed to get client")
	}

	client, err := helper.GetManagementClient(clientinfo, d.Get(orgIDVar).(string))
	if err != nil {
		return diag.FromErr(err)
	}

	actionsSet := d.Get(actionsVar).(*schema.Set)
	actions := make([]string, 0)
	for _, action := range actionsSet.List() {
		actions = append(actions, action.(string))
	}
	_, err = client.SetTriggerActions(ctx, &management.SetTriggerActionsRequest{
		FlowType:    action.FlowType(action.FlowType_value[d.Get(flowTypeVar).(string)]),
		TriggerType: action.TriggerType(action.TriggerType_value[d.Get(triggerTypeVar).(string)]),
		ActionIds:   actions,
	})
	if err != nil {
		return diag.Errorf("failed to update trigger actions: %v", err)
	}

	return nil
}

func create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Info(ctx, "started create")

	clientinfo, ok := m.(*helper.ClientInfo)
	if !ok {
		return diag.Errorf("failed to get client")
	}

	orgID := d.Get(orgIDVar).(string)
	client, err := helper.GetManagementClient(clientinfo, orgID)
	if err != nil {
		return diag.FromErr(err)
	}

	actionsSet := d.Get(actionsVar).(*schema.Set)
	actions := make([]string, 0)
	for _, action := range actionsSet.List() {
		actions = append(actions, action.(string))
	}
	flowType := d.Get(flowTypeVar).(string)
	triggerType := d.Get(triggerTypeVar).(string)
	_, err = client.SetTriggerActions(ctx, &management.SetTriggerActionsRequest{
		FlowType:    action.FlowType(action.FlowType_value[flowType]),
		TriggerType: action.TriggerType(action.TriggerType_value[triggerType]),
		ActionIds:   actions,
	})
	d.SetId(getTriggerActionsID(orgID, flowType, triggerType))

	return nil
}

func read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Info(ctx, "started read")

	orgID := d.Get(orgIDVar).(string)
	flowType := d.Get(flowTypeVar).(string)
	triggerType := d.Get(triggerTypeVar).(string)
	d.SetId(getTriggerActionsID(orgID, flowType, triggerType))
	return nil
}

func getTriggerActionsID(orgID, flowType string, triggerType string) string {
	return orgID + "_" + flowType + "_" + triggerType
}
