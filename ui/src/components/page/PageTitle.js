import React from "react";

import { EuiFlexGroup, EuiFlexItem, EuiIcon, EuiTitle } from "@elastic/eui";

import { useConfig } from "config";

export const PageTitle = ({ icon, title, size, postpend }) => {
  const { appConfig } = useConfig();

  return (
    <EuiFlexGroup direction="row" gutterSize="s" alignItems="center">
      <EuiFlexItem grow={false}>
        <EuiTitle size={size || "l"}>
          <span>
            <EuiIcon type={icon ? icon : appConfig.appIcon} size="xl" />
            &nbsp;{title}
          </span>
        </EuiTitle>
      </EuiFlexItem>
      {!!postpend && <EuiFlexItem grow={false}>{postpend}</EuiFlexItem>}
    </EuiFlexGroup>
  );
};
