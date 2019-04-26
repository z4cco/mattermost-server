// Copyright (c) 2016-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package store

import "github.com/mattermost/mattermost-server/model"
import "context"

type LayeredStoreSupplierResult struct {
	StoreResult
}

func NewSupplierResult() *LayeredStoreSupplierResult {
	return &LayeredStoreSupplierResult{}
}

type LayeredStoreSupplier interface {
	//
	// Control
	//
	SetChainNext(LayeredStoreSupplier)
	Next() LayeredStoreSupplier

	//
	// Reactions
	//), hints ...LayeredStoreHint)
	ReactionSave(ctx context.Context, reaction *model.Reaction, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	ReactionDelete(ctx context.Context, reaction *model.Reaction, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	ReactionGetForPost(ctx context.Context, postId string, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	ReactionDeleteAllWithEmojiName(ctx context.Context, emojiName string, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	ReactionPermanentDeleteBatch(ctx context.Context, endTime int64, limit int64, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	ReactionsBulkGetForPosts(ctx context.Context, postIds []string, hints ...LayeredStoreHint) *LayeredStoreSupplierResult

	// Roles
	RoleSave(ctx context.Context, role *model.Role, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	RoleGet(ctx context.Context, roleId string, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	RoleGetAll(ctx context.Context, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	RoleGetByName(ctx context.Context, name string, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	RoleGetByNames(ctx context.Context, names []string, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	RoleDelete(ctx context.Context, roldId string, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	RolePermanentDeleteAll(ctx context.Context, hints ...LayeredStoreHint) *LayeredStoreSupplierResult

	// Schemes
	SchemeSave(ctx context.Context, scheme *model.Scheme, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	SchemeGet(ctx context.Context, schemeId string, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	SchemeGetByName(ctx context.Context, schemeName string, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	SchemeDelete(ctx context.Context, schemeId string, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	SchemeGetAllPage(ctx context.Context, scope string, offset int, limit int, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	SchemePermanentDeleteAll(ctx context.Context, hints ...LayeredStoreHint) *LayeredStoreSupplierResult

	// Groups
	GroupCreate(ctx context.Context, group *model.Group, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	GroupGet(ctx context.Context, groupID string, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	GroupGetByRemoteID(ctx context.Context, remoteID string, groupSource model.GroupSource, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	GroupGetAllBySource(ctx context.Context, groupSource model.GroupSource, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	GroupUpdate(ctx context.Context, group *model.Group, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	GroupDelete(ctx context.Context, groupID string, hints ...LayeredStoreHint) *LayeredStoreSupplierResult

	GroupGetMemberUsers(ctx context.Context, groupID string, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	GroupGetMemberUsersPage(ctx context.Context, groupID string, offset int, limit int, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	GroupGetMemberCount(ctx context.Context, groupID string, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	GroupCreateOrRestoreMember(ctx context.Context, groupID string, userID string, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	GroupDeleteMember(ctx context.Context, groupID string, userID string, hints ...LayeredStoreHint) *LayeredStoreSupplierResult

	GroupCreateGroupSyncable(ctx context.Context, groupSyncable *model.GroupSyncable, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	GroupGetGroupSyncable(ctx context.Context, groupID string, syncableID string, syncableType model.GroupSyncableType, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	GroupGetAllGroupSyncablesByGroup(ctx context.Context, groupID string, syncableType model.GroupSyncableType, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	GroupUpdateGroupSyncable(ctx context.Context, groupSyncable *model.GroupSyncable, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	GroupDeleteGroupSyncable(ctx context.Context, groupID string, syncableID string, syncableType model.GroupSyncableType, hints ...LayeredStoreHint) *LayeredStoreSupplierResult

	TeamMembersToAdd(ctx context.Context, since int64, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	ChannelMembersToAdd(ctx context.Context, since int64, hints ...LayeredStoreHint) *LayeredStoreSupplierResult

	TeamMembersToRemove(ctx context.Context, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	ChannelMembersToRemove(ctx context.Context, hints ...LayeredStoreHint) *LayeredStoreSupplierResult

	GetGroupsByChannel(ctx context.Context, channelId string, page, perPage int, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	GetGroupsByTeam(ctx context.Context, teamId string, page, perPage *int, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
	GetGroupsPage(ctx context.Context, page, perPage int, opts model.GroupSearchOpts, hints ...LayeredStoreHint) *LayeredStoreSupplierResult
}
