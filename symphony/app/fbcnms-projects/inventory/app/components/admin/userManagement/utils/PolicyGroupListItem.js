/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import type {
  PermissionsPolicy,
  UserPermissionsGroup,
} from './UserManagementUtils';
import type {ToggleButtonDisplay} from './ListItem';

import * as React from 'react';
import MemberListItem from './MemberListItem';
import {useCallback, useEffect, useState} from 'react';
import {useUserManagement} from '../UserManagementContext';

export type AssigenmentButtonProp = $ReadOnly<{|
  assigmentButton?: ?ToggleButtonDisplay,
|}>;

type Props = $ReadOnly<{|
  group: UserPermissionsGroup,
  isMember?: ?boolean,
  policy?: ?PermissionsPolicy,
  className?: ?string,
  ...AssigenmentButtonProp,
|}>;

const checkIsGroupInPolicy = (
  group: UserPermissionsGroup,
  policy?: ?PermissionsPolicy,
) => policy == null || policy.groups.find(g => g.id === group.id) != null;

export default function PolicyGroupListItem(props: Props) {
  const {group, policy, assigmentButton, className} = props;
  const [isGroupInPolicy, setIsGroupInPolicy] = useState(false);
  useEffect(() => setIsGroupInPolicy(checkIsGroupInPolicy(group, policy)), [
    group,
    policy,
  ]);
  const userManagement = useUserManagement();

  const toggleAssigment = useCallback(
    (group, shouldAssign) => {
      if (policy == null) {
        return Promise.resolve();
      }
      const add = shouldAssign ? [group.id] : [];
      const remove = shouldAssign ? [] : [group.id];
      return userManagement
        .updatePolicyGroups(policy, add, remove)
        .then(() => setIsGroupInPolicy(shouldAssign));
    },
    [policy, userManagement],
  );

  return (
    <MemberListItem
      member={{
        item: group,
        isMember: isGroupInPolicy,
      }}
      className={className}
      assigmentButton={assigmentButton}
      onAssignToggle={() => toggleAssigment(group, !isGroupInPolicy)}>
      {group.name}
    </MemberListItem>
  );
}
