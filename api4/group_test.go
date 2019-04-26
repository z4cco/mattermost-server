// Copyright (c) 2018-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package api4

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mattermost/mattermost-server/model"
)

func TestGetGroup(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	id := model.NewId()
	g, err := th.App.CreateGroup(&model.Group{
		DisplayName: "dn_" + id,
		Name:        "name" + id,
		Source:      model.GroupSourceLdap,
		Description: "description_" + id,
		RemoteId:    model.NewId(),
	})
	assert.Nil(t, err)

	_, response := th.Client.GetGroup(g.Id, "")
	CheckNotImplementedStatus(t, response)

	_, response = th.SystemAdminClient.GetGroup(g.Id, "")
	CheckNotImplementedStatus(t, response)

	th.App.SetLicense(model.NewTestLicense("ldap"))

	group, response := th.SystemAdminClient.GetGroup(g.Id, "")
	CheckNoError(t, response)

	assert.Equal(t, g.DisplayName, group.DisplayName)
	assert.Equal(t, g.Name, group.Name)
	assert.Equal(t, g.Source, group.Source)
	assert.Equal(t, g.Description, group.Description)
	assert.Equal(t, g.RemoteId, group.RemoteId)
	assert.Equal(t, g.CreateAt, group.CreateAt)
	assert.Equal(t, g.UpdateAt, group.UpdateAt)
	assert.Equal(t, g.DeleteAt, group.DeleteAt)

	_, response = th.SystemAdminClient.GetGroup(model.NewId(), "")
	CheckNotFoundStatus(t, response)

	_, response = th.SystemAdminClient.GetGroup("12345", "")
	CheckBadRequestStatus(t, response)

	th.SystemAdminClient.Logout()
	_, response = th.SystemAdminClient.GetGroup(group.Id, "")
	CheckUnauthorizedStatus(t, response)
}

func TestPatchGroup(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	id := model.NewId()
	g, err := th.App.CreateGroup(&model.Group{
		DisplayName: "dn_" + id,
		Name:        "name" + id,
		Source:      model.GroupSourceLdap,
		Description: "description_" + id,
		RemoteId:    model.NewId(),
	})
	assert.Nil(t, err)

	updateFmt := "%s_updated"

	newName := fmt.Sprintf(updateFmt, g.Name)
	newDisplayName := fmt.Sprintf(updateFmt, g.DisplayName)
	newDescription := fmt.Sprintf(updateFmt, g.Description)

	gp := &model.GroupPatch{
		Name:        &newName,
		DisplayName: &newDisplayName,
		Description: &newDescription,
	}

	_, response := th.Client.PatchGroup(g.Id, gp)
	CheckNotImplementedStatus(t, response)

	_, response = th.SystemAdminClient.PatchGroup(g.Id, gp)
	CheckNotImplementedStatus(t, response)

	th.App.SetLicense(model.NewTestLicense("ldap"))

	group2, response := th.SystemAdminClient.PatchGroup(g.Id, gp)
	CheckOKStatus(t, response)

	group, response := th.SystemAdminClient.GetGroup(g.Id, "")
	CheckNoError(t, response)

	assert.Equal(t, *gp.DisplayName, group.DisplayName)
	assert.Equal(t, *gp.DisplayName, group2.DisplayName)
	assert.Equal(t, *gp.Name, group.Name)
	assert.Equal(t, *gp.Name, group2.Name)
	assert.Equal(t, *gp.Description, group.Description)
	assert.Equal(t, *gp.Description, group2.Description)

	assert.Equal(t, group2.UpdateAt, group.UpdateAt)

	assert.Equal(t, g.Source, group.Source)
	assert.Equal(t, g.Source, group2.Source)
	assert.Equal(t, g.RemoteId, group.RemoteId)
	assert.Equal(t, g.RemoteId, group2.RemoteId)
	assert.Equal(t, g.CreateAt, group.CreateAt)
	assert.Equal(t, g.CreateAt, group2.CreateAt)
	assert.Equal(t, g.DeleteAt, group.DeleteAt)
	assert.Equal(t, g.DeleteAt, group2.DeleteAt)

	_, response = th.SystemAdminClient.PatchGroup(model.NewId(), gp)
	CheckNotFoundStatus(t, response)

	th.SystemAdminClient.Logout()
	_, response = th.SystemAdminClient.PatchGroup(group.Id, gp)
	CheckUnauthorizedStatus(t, response)
}

func TestLinkGroupTeam(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	id := model.NewId()
	g, err := th.App.CreateGroup(&model.Group{
		DisplayName: "dn_" + id,
		Name:        "name" + id,
		Source:      model.GroupSourceLdap,
		Description: "description_" + id,
		RemoteId:    model.NewId(),
	})
	assert.Nil(t, err)

	patch := &model.GroupSyncablePatch{
		AutoAdd: model.NewBool(true),
	}

	_, response := th.Client.LinkGroupSyncable(g.Id, th.BasicTeam.Id, model.GroupSyncableTypeTeam, patch)
	CheckNotImplementedStatus(t, response)

	_, response = th.SystemAdminClient.LinkGroupSyncable(g.Id, th.BasicTeam.Id, model.GroupSyncableTypeTeam, patch)
	CheckNotImplementedStatus(t, response)

	th.App.SetLicense(model.NewTestLicense("ldap"))

	groupTeam, response := th.SystemAdminClient.LinkGroupSyncable(g.Id, th.BasicTeam.Id, model.GroupSyncableTypeTeam, patch)
	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assert.NotNil(t, groupTeam)
}

func TestLinkGroupChannel(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	id := model.NewId()
	g, err := th.App.CreateGroup(&model.Group{
		DisplayName: "dn_" + id,
		Name:        "name" + id,
		Source:      model.GroupSourceLdap,
		Description: "description_" + id,
		RemoteId:    model.NewId(),
	})
	assert.Nil(t, err)

	patch := &model.GroupSyncablePatch{
		AutoAdd: model.NewBool(true),
	}

	_, response := th.Client.LinkGroupSyncable(g.Id, th.BasicChannel.Id, model.GroupSyncableTypeChannel, patch)
	CheckNotImplementedStatus(t, response)

	_, response = th.SystemAdminClient.LinkGroupSyncable(g.Id, th.BasicChannel.Id, model.GroupSyncableTypeChannel, patch)
	CheckNotImplementedStatus(t, response)

	th.App.SetLicense(model.NewTestLicense("ldap"))

	_, response = th.SystemAdminClient.LinkGroupSyncable(g.Id, th.BasicChannel.Id, model.GroupSyncableTypeChannel, patch)
	assert.Equal(t, http.StatusCreated, response.StatusCode)
}

func TestUnlinkGroupTeam(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	id := model.NewId()
	g, err := th.App.CreateGroup(&model.Group{
		DisplayName: "dn_" + id,
		Name:        "name" + id,
		Source:      model.GroupSourceLdap,
		Description: "description_" + id,
		RemoteId:    model.NewId(),
	})
	assert.Nil(t, err)

	patch := &model.GroupSyncablePatch{
		AutoAdd: model.NewBool(true),
	}

	th.App.SetLicense(model.NewTestLicense("ldap"))

	_, response := th.SystemAdminClient.LinkGroupSyncable(g.Id, th.BasicTeam.Id, model.GroupSyncableTypeTeam, patch)
	assert.Equal(t, http.StatusCreated, response.StatusCode)

	th.App.SetLicense(nil)

	response = th.Client.UnlinkGroupSyncable(g.Id, th.BasicTeam.Id, model.GroupSyncableTypeTeam)
	CheckNotImplementedStatus(t, response)

	response = th.SystemAdminClient.UnlinkGroupSyncable(g.Id, th.BasicTeam.Id, model.GroupSyncableTypeTeam)
	CheckNotImplementedStatus(t, response)

	th.App.SetLicense(model.NewTestLicense("ldap"))

	response = th.SystemAdminClient.UnlinkGroupSyncable(g.Id, th.BasicTeam.Id, model.GroupSyncableTypeTeam)
	CheckOKStatus(t, response)
}

func TestUnlinkGroupChannel(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	id := model.NewId()
	g, err := th.App.CreateGroup(&model.Group{
		DisplayName: "dn_" + id,
		Name:        "name" + id,
		Source:      model.GroupSourceLdap,
		Description: "description_" + id,
		RemoteId:    model.NewId(),
	})
	assert.Nil(t, err)

	patch := &model.GroupSyncablePatch{
		AutoAdd: model.NewBool(true),
	}

	th.App.SetLicense(model.NewTestLicense("ldap"))

	_, response := th.SystemAdminClient.LinkGroupSyncable(g.Id, th.BasicChannel.Id, model.GroupSyncableTypeChannel, patch)
	assert.Equal(t, http.StatusCreated, response.StatusCode)

	th.App.SetLicense(nil)

	response = th.Client.UnlinkGroupSyncable(g.Id, th.BasicChannel.Id, model.GroupSyncableTypeChannel)
	CheckNotImplementedStatus(t, response)

	response = th.SystemAdminClient.UnlinkGroupSyncable(g.Id, th.BasicChannel.Id, model.GroupSyncableTypeChannel)
	CheckNotImplementedStatus(t, response)

	th.App.SetLicense(model.NewTestLicense("ldap"))

	response = th.SystemAdminClient.UnlinkGroupSyncable(g.Id, th.BasicChannel.Id, model.GroupSyncableTypeChannel)
	CheckOKStatus(t, response)
}

func TestGetGroupTeam(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	id := model.NewId()
	g, err := th.App.CreateGroup(&model.Group{
		DisplayName: "dn_" + id,
		Name:        "name" + id,
		Source:      model.GroupSourceLdap,
		Description: "description_" + id,
		RemoteId:    model.NewId(),
	})
	assert.Nil(t, err)

	_, response := th.Client.GetGroupSyncable(g.Id, th.BasicTeam.Id, model.GroupSyncableTypeTeam, "")
	CheckNotImplementedStatus(t, response)

	_, response = th.SystemAdminClient.GetGroupSyncable(g.Id, th.BasicTeam.Id, model.GroupSyncableTypeTeam, "")
	CheckNotImplementedStatus(t, response)

	th.App.SetLicense(model.NewTestLicense("ldap"))

	patch := &model.GroupSyncablePatch{
		AutoAdd: model.NewBool(true),
	}

	_, response = th.SystemAdminClient.LinkGroupSyncable(g.Id, th.BasicTeam.Id, model.GroupSyncableTypeTeam, patch)
	assert.Equal(t, http.StatusCreated, response.StatusCode)

	groupSyncable, response := th.SystemAdminClient.GetGroupSyncable(g.Id, th.BasicTeam.Id, model.GroupSyncableTypeTeam, "")
	CheckOKStatus(t, response)
	assert.NotNil(t, groupSyncable)

	assert.Equal(t, g.Id, groupSyncable.GroupId)
	assert.Equal(t, th.BasicTeam.Id, groupSyncable.SyncableId)
	assert.Equal(t, *patch.AutoAdd, groupSyncable.AutoAdd)

	_, response = th.SystemAdminClient.GetGroupSyncable(model.NewId(), th.BasicTeam.Id, model.GroupSyncableTypeTeam, "")
	CheckNotFoundStatus(t, response)

	_, response = th.SystemAdminClient.GetGroupSyncable(g.Id, model.NewId(), model.GroupSyncableTypeTeam, "")
	CheckNotFoundStatus(t, response)

	_, response = th.SystemAdminClient.GetGroupSyncable("asdfasdfe3", th.BasicTeam.Id, model.GroupSyncableTypeTeam, "")
	CheckBadRequestStatus(t, response)

	_, response = th.SystemAdminClient.GetGroupSyncable(g.Id, "asdfasdfe3", model.GroupSyncableTypeTeam, "")
	CheckBadRequestStatus(t, response)

	th.SystemAdminClient.Logout()
	_, response = th.SystemAdminClient.GetGroupSyncable(g.Id, th.BasicTeam.Id, model.GroupSyncableTypeTeam, "")
	CheckUnauthorizedStatus(t, response)
}

func TestGetGroupChannel(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	id := model.NewId()
	g, err := th.App.CreateGroup(&model.Group{
		DisplayName: "dn_" + id,
		Name:        "name" + id,
		Source:      model.GroupSourceLdap,
		Description: "description_" + id,
		RemoteId:    model.NewId(),
	})
	assert.Nil(t, err)

	_, response := th.Client.GetGroupSyncable(g.Id, th.BasicChannel.Id, model.GroupSyncableTypeChannel, "")
	CheckNotImplementedStatus(t, response)

	_, response = th.SystemAdminClient.GetGroupSyncable(g.Id, th.BasicChannel.Id, model.GroupSyncableTypeChannel, "")
	CheckNotImplementedStatus(t, response)

	th.App.SetLicense(model.NewTestLicense("ldap"))

	patch := &model.GroupSyncablePatch{
		AutoAdd: model.NewBool(true),
	}

	_, response = th.SystemAdminClient.LinkGroupSyncable(g.Id, th.BasicChannel.Id, model.GroupSyncableTypeChannel, patch)
	assert.Equal(t, http.StatusCreated, response.StatusCode)

	groupSyncable, response := th.SystemAdminClient.GetGroupSyncable(g.Id, th.BasicChannel.Id, model.GroupSyncableTypeChannel, "")
	CheckOKStatus(t, response)
	assert.NotNil(t, groupSyncable)

	assert.Equal(t, g.Id, groupSyncable.GroupId)
	assert.Equal(t, th.BasicChannel.Id, groupSyncable.SyncableId)
	assert.Equal(t, *patch.AutoAdd, groupSyncable.AutoAdd)

	_, response = th.SystemAdminClient.GetGroupSyncable(model.NewId(), th.BasicChannel.Id, model.GroupSyncableTypeChannel, "")
	CheckNotFoundStatus(t, response)

	_, response = th.SystemAdminClient.GetGroupSyncable(g.Id, model.NewId(), model.GroupSyncableTypeChannel, "")
	CheckNotFoundStatus(t, response)

	_, response = th.SystemAdminClient.GetGroupSyncable("asdfasdfe3", th.BasicChannel.Id, model.GroupSyncableTypeChannel, "")
	CheckBadRequestStatus(t, response)

	_, response = th.SystemAdminClient.GetGroupSyncable(g.Id, "asdfasdfe3", model.GroupSyncableTypeChannel, "")
	CheckBadRequestStatus(t, response)

	th.SystemAdminClient.Logout()
	_, response = th.SystemAdminClient.GetGroupSyncable(g.Id, th.BasicChannel.Id, model.GroupSyncableTypeChannel, "")
	CheckUnauthorizedStatus(t, response)
}

func TestGetGroupTeams(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	id := model.NewId()
	g, err := th.App.CreateGroup(&model.Group{
		DisplayName: "dn_" + id,
		Name:        "name" + id,
		Source:      model.GroupSourceLdap,
		Description: "description_" + id,
		RemoteId:    model.NewId(),
	})
	assert.Nil(t, err)

	th.App.SetLicense(model.NewTestLicense("ldap"))

	patch := &model.GroupSyncablePatch{
		AutoAdd: model.NewBool(true),
	}

	for i := 0; i < 10; i++ {
		team := th.CreateTeam()
		_, response := th.SystemAdminClient.LinkGroupSyncable(g.Id, team.Id, model.GroupSyncableTypeTeam, patch)
		assert.Equal(t, http.StatusCreated, response.StatusCode)
	}

	th.App.SetLicense(nil)

	_, response := th.Client.GetGroupSyncables(g.Id, model.GroupSyncableTypeTeam, "")
	CheckNotImplementedStatus(t, response)

	_, response = th.SystemAdminClient.GetGroupSyncables(g.Id, model.GroupSyncableTypeTeam, "")
	CheckNotImplementedStatus(t, response)

	th.App.SetLicense(model.NewTestLicense("ldap"))

	_, response = th.Client.GetGroupSyncables(g.Id, model.GroupSyncableTypeTeam, "")
	assert.Equal(t, http.StatusForbidden, response.StatusCode)

	groupSyncables, response := th.SystemAdminClient.GetGroupSyncables(g.Id, model.GroupSyncableTypeTeam, "")
	CheckOKStatus(t, response)

	assert.Len(t, groupSyncables, 10)

	th.SystemAdminClient.Logout()
	_, response = th.SystemAdminClient.GetGroupSyncables(g.Id, model.GroupSyncableTypeTeam, "")
	CheckUnauthorizedStatus(t, response)
}

func TestGetGroupChannels(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	id := model.NewId()
	g, err := th.App.CreateGroup(&model.Group{
		DisplayName: "dn_" + id,
		Name:        "name" + id,
		Source:      model.GroupSourceLdap,
		Description: "description_" + id,
		RemoteId:    model.NewId(),
	})
	assert.Nil(t, err)

	th.App.SetLicense(model.NewTestLicense("ldap"))

	patch := &model.GroupSyncablePatch{
		AutoAdd: model.NewBool(true),
	}

	for i := 0; i < 10; i++ {
		channel := th.CreatePublicChannel()
		_, response := th.SystemAdminClient.LinkGroupSyncable(g.Id, channel.Id, model.GroupSyncableTypeChannel, patch)
		assert.Equal(t, http.StatusCreated, response.StatusCode)
	}

	th.App.SetLicense(nil)

	_, response := th.Client.GetGroupSyncables(g.Id, model.GroupSyncableTypeChannel, "")
	CheckNotImplementedStatus(t, response)

	_, response = th.SystemAdminClient.GetGroupSyncables(g.Id, model.GroupSyncableTypeChannel, "")
	CheckNotImplementedStatus(t, response)

	th.App.SetLicense(model.NewTestLicense("ldap"))

	_, response = th.Client.GetGroupSyncables(g.Id, model.GroupSyncableTypeChannel, "")
	assert.Equal(t, http.StatusForbidden, response.StatusCode)

	groupSyncables, response := th.SystemAdminClient.GetGroupSyncables(g.Id, model.GroupSyncableTypeChannel, "")
	CheckOKStatus(t, response)

	assert.Len(t, groupSyncables, 10)

	th.SystemAdminClient.Logout()
	_, response = th.SystemAdminClient.GetGroupSyncables(g.Id, model.GroupSyncableTypeChannel, "")
	CheckUnauthorizedStatus(t, response)
}

func TestPatchGroupTeam(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	id := model.NewId()
	g, err := th.App.CreateGroup(&model.Group{
		DisplayName: "dn_" + id,
		Name:        "name" + id,
		Source:      model.GroupSourceLdap,
		Description: "description_" + id,
		RemoteId:    model.NewId(),
	})
	assert.Nil(t, err)

	patch := &model.GroupSyncablePatch{
		AutoAdd: model.NewBool(true),
	}

	th.App.SetLicense(model.NewTestLicense("ldap"))

	groupSyncable, response := th.SystemAdminClient.LinkGroupSyncable(g.Id, th.BasicTeam.Id, model.GroupSyncableTypeTeam, patch)
	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assert.NotNil(t, groupSyncable)
	assert.True(t, groupSyncable.AutoAdd)

	_, response = th.Client.PatchGroupSyncable(g.Id, th.BasicTeam.Id, model.GroupSyncableTypeTeam, patch)
	assert.Equal(t, http.StatusForbidden, response.StatusCode)

	th.App.SetLicense(nil)

	_, response = th.SystemAdminClient.PatchGroupSyncable(g.Id, th.BasicTeam.Id, model.GroupSyncableTypeTeam, patch)
	CheckNotImplementedStatus(t, response)

	th.App.SetLicense(model.NewTestLicense("ldap"))

	patch.AutoAdd = model.NewBool(false)
	groupSyncable, response = th.SystemAdminClient.PatchGroupSyncable(g.Id, th.BasicTeam.Id, model.GroupSyncableTypeTeam, patch)
	CheckOKStatus(t, response)
	assert.False(t, groupSyncable.AutoAdd)

	assert.Equal(t, g.Id, groupSyncable.GroupId)
	assert.Equal(t, th.BasicTeam.Id, groupSyncable.SyncableId)
	assert.Equal(t, model.GroupSyncableTypeTeam, groupSyncable.Type)

	patch.AutoAdd = model.NewBool(true)
	groupSyncable, response = th.SystemAdminClient.PatchGroupSyncable(g.Id, th.BasicTeam.Id, model.GroupSyncableTypeTeam, patch)
	CheckOKStatus(t, response)

	_, response = th.SystemAdminClient.PatchGroupSyncable(model.NewId(), th.BasicTeam.Id, model.GroupSyncableTypeTeam, patch)
	CheckNotFoundStatus(t, response)

	_, response = th.SystemAdminClient.PatchGroupSyncable(g.Id, model.NewId(), model.GroupSyncableTypeTeam, patch)
	CheckNotFoundStatus(t, response)

	_, response = th.SystemAdminClient.PatchGroupSyncable("abc", th.BasicTeam.Id, model.GroupSyncableTypeTeam, patch)
	CheckBadRequestStatus(t, response)

	_, response = th.SystemAdminClient.PatchGroupSyncable(g.Id, "abc", model.GroupSyncableTypeTeam, patch)
	CheckBadRequestStatus(t, response)

	th.SystemAdminClient.Logout()
	_, response = th.SystemAdminClient.PatchGroupSyncable(g.Id, th.BasicTeam.Id, model.GroupSyncableTypeTeam, patch)
	CheckUnauthorizedStatus(t, response)
}

func TestPatchGroupChannel(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	id := model.NewId()
	g, err := th.App.CreateGroup(&model.Group{
		DisplayName: "dn_" + id,
		Name:        "name" + id,
		Source:      model.GroupSourceLdap,
		Description: "description_" + id,
		RemoteId:    model.NewId(),
	})
	assert.Nil(t, err)

	patch := &model.GroupSyncablePatch{
		AutoAdd: model.NewBool(true),
	}

	th.App.SetLicense(model.NewTestLicense("ldap"))

	groupSyncable, response := th.SystemAdminClient.LinkGroupSyncable(g.Id, th.BasicChannel.Id, model.GroupSyncableTypeChannel, patch)
	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assert.NotNil(t, groupSyncable)
	assert.True(t, groupSyncable.AutoAdd)

	_, response = th.Client.PatchGroupSyncable(g.Id, th.BasicChannel.Id, model.GroupSyncableTypeChannel, patch)
	assert.Equal(t, http.StatusForbidden, response.StatusCode)

	th.App.SetLicense(nil)

	_, response = th.SystemAdminClient.PatchGroupSyncable(g.Id, th.BasicChannel.Id, model.GroupSyncableTypeChannel, patch)
	CheckNotImplementedStatus(t, response)

	th.App.SetLicense(model.NewTestLicense("ldap"))

	patch.AutoAdd = model.NewBool(false)
	groupSyncable, response = th.SystemAdminClient.PatchGroupSyncable(g.Id, th.BasicChannel.Id, model.GroupSyncableTypeChannel, patch)
	CheckOKStatus(t, response)
	assert.False(t, groupSyncable.AutoAdd)

	assert.Equal(t, g.Id, groupSyncable.GroupId)
	assert.Equal(t, th.BasicChannel.Id, groupSyncable.SyncableId)
	assert.Equal(t, model.GroupSyncableTypeChannel, groupSyncable.Type)

	patch.AutoAdd = model.NewBool(true)
	groupSyncable, response = th.SystemAdminClient.PatchGroupSyncable(g.Id, th.BasicChannel.Id, model.GroupSyncableTypeChannel, patch)
	CheckOKStatus(t, response)

	_, response = th.SystemAdminClient.PatchGroupSyncable(model.NewId(), th.BasicChannel.Id, model.GroupSyncableTypeChannel, patch)
	CheckNotFoundStatus(t, response)

	_, response = th.SystemAdminClient.PatchGroupSyncable(g.Id, model.NewId(), model.GroupSyncableTypeChannel, patch)
	CheckNotFoundStatus(t, response)

	_, response = th.SystemAdminClient.PatchGroupSyncable("abc", th.BasicChannel.Id, model.GroupSyncableTypeChannel, patch)
	CheckBadRequestStatus(t, response)

	_, response = th.SystemAdminClient.PatchGroupSyncable(g.Id, "abc", model.GroupSyncableTypeChannel, patch)
	CheckBadRequestStatus(t, response)

	th.SystemAdminClient.Logout()
	_, response = th.SystemAdminClient.PatchGroupSyncable(g.Id, th.BasicChannel.Id, model.GroupSyncableTypeChannel, patch)
	CheckUnauthorizedStatus(t, response)
}

func TestGetGroupsByChannel(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	id := model.NewId()
	group, err := th.App.CreateGroup(&model.Group{
		DisplayName: "dn_" + id,
		Name:        "name" + id,
		Source:      model.GroupSourceLdap,
		Description: "description_" + id,
		RemoteId:    model.NewId(),
	})
	assert.Nil(t, err)

	_, err = th.App.CreateGroupSyncable(&model.GroupSyncable{
		AutoAdd:    true,
		SyncableId: th.BasicChannel.Id,
		Type:       model.GroupSyncableTypeChannel,
		GroupId:    group.Id,
	})
	assert.Nil(t, err)

	_, response := th.SystemAdminClient.GetGroupsByChannel("asdfasdf", 0, 60)
	CheckBadRequestStatus(t, response)

	th.App.SetLicense(nil)

	_, response = th.SystemAdminClient.GetGroupsByChannel(th.BasicChannel.Id, 0, 60)
	CheckNotImplementedStatus(t, response)

	th.App.SetLicense(model.NewTestLicense("ldap"))

	_, response = th.Client.GetGroupsByChannel(th.BasicChannel.Id, 0, 60)
	CheckForbiddenStatus(t, response)

	groups, response := th.SystemAdminClient.GetGroupsByChannel(th.BasicChannel.Id, 0, 60)
	assert.Nil(t, response.Error)
	assert.ElementsMatch(t, []*model.Group{group}, groups)

	groups, response = th.SystemAdminClient.GetGroupsByChannel(model.NewId(), 0, 60)
	assert.Nil(t, response.Error)
	assert.Empty(t, groups)
}

func TestGetGroupsByTeam(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	id := model.NewId()
	group, err := th.App.CreateGroup(&model.Group{
		DisplayName: "dn_" + id,
		Name:        "name" + id,
		Source:      model.GroupSourceLdap,
		Description: "description_" + id,
		RemoteId:    model.NewId(),
	})
	assert.Nil(t, err)

	_, err = th.App.CreateGroupSyncable(&model.GroupSyncable{
		AutoAdd:    true,
		SyncableId: th.BasicTeam.Id,
		Type:       model.GroupSyncableTypeTeam,
		GroupId:    group.Id,
	})
	assert.Nil(t, err)

	page := model.NewInt(0)
	perPage := model.NewInt(60)

	_, response := th.SystemAdminClient.GetGroupsByTeam("asdfasdf", page, perPage)
	CheckBadRequestStatus(t, response)

	th.App.SetLicense(nil)

	_, response = th.SystemAdminClient.GetGroupsByTeam(th.BasicTeam.Id, page, perPage)
	CheckNotImplementedStatus(t, response)

	th.App.SetLicense(model.NewTestLicense("ldap"))

	_, response = th.Client.GetGroupsByTeam(th.BasicTeam.Id, page, perPage)
	CheckForbiddenStatus(t, response)

	groups, response := th.SystemAdminClient.GetGroupsByTeam(th.BasicTeam.Id, page, perPage)
	assert.Nil(t, response.Error)
	assert.ElementsMatch(t, []*model.Group{group}, groups)

	groups, response = th.SystemAdminClient.GetGroupsByTeam(model.NewId(), page, perPage)
	assert.Nil(t, response.Error)
	assert.Empty(t, groups)
}
