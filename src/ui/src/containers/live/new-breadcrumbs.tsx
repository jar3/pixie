/*
 * Copyright 2018- The Pixie Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

import * as React from 'react';

import {
  Theme, withStyles,
} from '@material-ui/core/styles';
import { createStyles } from '@material-ui/styles';
import Paper from '@material-ui/core/Paper';
import { GQLClusterStatus as ClusterStatus } from '@pixie-labs/api';
import { useListClusters, useAutocompleteFieldSuggester } from '@pixie-labs/api-react';

import {
  Breadcrumbs, BreadcrumbOptions,
  PixieCommandIcon, StatusCell,
} from '@pixie-labs/components';
import { ClusterContext } from 'common/cluster-context';
import { argVariableMap, argTypesForVis } from 'utils/args-utils';
import { SCRATCH_SCRIPT, ScriptsContext } from 'containers/App/scripts-context';
import { ScriptContext } from 'context/new-script-context';
import { pxTypeToEntityType, entityStatusGroup } from 'containers/command-input/autocomplete-utils';
import { clusterStatusGroup } from 'containers/admin/utils';
import ExecuteScriptButton from './new-execute-button';
import { parseVisSilently, Variable } from './vis';

const styles = (({ shape, palette, spacing }: Theme) => createStyles({
  root: {
    display: 'flex',
    alignItems: 'center',
    marginTop: spacing(1),
    marginRight: spacing(3.5),
    marginLeft: spacing(3),
    marginBottom: spacing(1),
    background: palette.background.three,
    // This adds a scroll to the breadcrumbs on overflow,
    // but it's hard for the user to know it exists. Perhaps we can
    // consider adding a scroll effect or something to make it easier to
    // discover.
    borderRadius: shape.borderRadius,
    boxShadow: '4px 4px 4px rgba(0, 0, 0, 0.25)',
    border: palette.border.unFocused,
  },
  spacer: {
    flex: 1,
  },
  verticalLine: {
    borderLeft: `2px solid ${palette.foreground.grey1}`,
    height: spacing(2.7),
    padding: 0,
  },
  pixieIcon: {
    color: palette.primary.main,
    marginRight: spacing(1),
    marginLeft: spacing(0.5),
  },
  iconContainer: {
    padding: 0,
  },
  breadcrumbs: {
    width: '100%',
    overflow: 'hidden',
    display: 'flex',
  },
}));

const LiveViewBreadcrumbs = ({ classes }) => {
  const [clusters, loading, error] = useListClusters();
  const { selectedCluster, setCluster, selectedClusterUID } = React.useContext(ClusterContext);
  const { scripts } = React.useContext(ScriptsContext);

  const {
    args, script, setScriptAndArgs,
  } = React.useContext(ScriptContext);

  const getCompletions = useAutocompleteFieldSuggester(selectedClusterUID);

  const scriptIds = React.useMemo(() => [...scripts.keys()], [scripts]);
  // eslint-disable-next-line react-hooks/exhaustive-deps
  const clusterIds = React.useMemo(() => clusters?.map((c) => c.id), [clusters]);
  const collapsedClusterIds = clusterIds?.join(',');
  const collapsedScriptIds = scriptIds?.join(',');

  type MemoCrumbs = { entityBreadcrumbs: BreadcrumbOptions[]; argBreadcrumbs: BreadcrumbOptions[] };
  const { entityBreadcrumbs, argBreadcrumbs }: MemoCrumbs = React.useMemo(() => {
    // eslint-disable-next-line @typescript-eslint/no-shadow
    const entityBreadcrumbs: BreadcrumbOptions[] = [];
    // eslint-disable-next-line @typescript-eslint/no-shadow
    const argBreadcrumbs: BreadcrumbOptions[] = [];

    if (loading || !clusters || error) return { entityBreadcrumbs, argBreadcrumbs };

    // Cluster always goes first in breadcrumbs.
    const clusterName = clusters.find((c) => c.id === selectedCluster)?.prettyClusterName || 'unknown cluster';
    const clusterNameToID: Record<string, string> = {};
    clusters.forEach((c) => {
      clusterNameToID[c.prettyClusterName] = c.id;
    });
    entityBreadcrumbs.push({
      title: 'cluster',
      value: clusterName,
      selectable: true,
      // eslint-disable-next-line
      getListItems: async (input) => (clusters.filter((c) => c.status !== ClusterStatus.CS_DISCONNECTED
        && c.prettyClusterName.includes(input))
        .map((c) => ({ value: c.prettyClusterName, icon: <StatusCell statusGroup={clusterStatusGroup(c.status)} /> }))
      ),
      onSelect: (input) => {
        setCluster(clusterNameToID[input]);
      },
      requireCompletion: true,
    });

    // Add args to breadcrumbs.
    const argTypes = argTypesForVis(script?.vis);
    const variables = argVariableMap(script?.vis);

    for (const [argName, argVal] of Object.entries(args)) {
      // Only add suggestions if validValues are specified. Otherwise, the dropdown is populated with autocomplete
      // entities or has no elements and the user must manually type in values.
      const variable: Variable = variables[argName];

      const argProps = {
        title: argName,
        value: argVal?.toString(),
        selectable: true,
        allowTyping: true,
        onSelect: (newVal) => {
          setScriptAndArgs(script, { ...args, [argName]: newVal });
        },
        getListItems: null,
        requireCompletion: false,
        placeholder: variable?.description,
      };

      if (variable && typeof variable.defaultValue === 'undefined') {
        argProps.title += '*';
      }

      if (variable?.validValues?.length) {
        argProps.getListItems = async (input) => (variable.validValues
          .filter((suggestion) => input === '' || suggestion.indexOf(input) >= 0)
          .map((suggestion) => ({
            value: suggestion,
            description: '',
          })));

        argProps.requireCompletion = true;
      }

      const entityType = pxTypeToEntityType(argTypes[argName]);
      if (entityType !== 'AEK_UNKNOWN') {
        argProps.getListItems = async (input) => (
          (await getCompletions(input, entityType)).map((suggestion) => ({
            value: suggestion.name,
            description: suggestion.description,
            icon: <StatusCell statusGroup={entityStatusGroup(suggestion.state)} />,
          })));
      }

      // TODO(michelle): Ideally we should just be able to use the entityType to determine whether the
      // arg appears in the argBreadcrumbs. However, some entities still don't have a corresponding entity type
      // (such as nodes), since they are not yet supported in autocomplete. Until that is fixed, this is hard-coded
      // for now.
      if (argName === 'start_time' || argName === 'start') {
        argBreadcrumbs.push(argProps);
      } else {
        entityBreadcrumbs.push(argProps);
      }
    }

    // Add script at end of breadcrumbs.
    entityBreadcrumbs.push({
      title: 'script',
      value: script?.id,
      selectable: true,
      allowTyping: true,
      getListItems: async (input) => {
        // Turns "  px/be_spoke-  " into "px/bespoke"
        const normalize = (s: string) => s.toLowerCase().replace(/[^a-z0-9/]+/gi, '');
        const normalizedInput = normalize(input);
        const normalizedScratchId = normalize(SCRATCH_SCRIPT.id);

        const ids = !input ? [...scriptIds] : scriptIds.filter((s) => {
          const ns = normalize(s);
          return ns === normalizedScratchId || ns.indexOf(normalizedInput) >= 0;
        });

        // The scratch script should always appear at the top of the list for visibility. It doesn't get auto-selected
        // unless it's the only thing in the list.
        const scratchIndex = scriptIds.indexOf(SCRATCH_SCRIPT.id);
        if (scratchIndex !== -1) {
          scriptIds.splice(scratchIndex, 1);
          scriptIds.unshift(SCRATCH_SCRIPT.id);
        }

        return ids.map((scriptId) => ({
          value: scriptId,
          description: scripts.get(scriptId).description,
          autoSelectPriority: scriptId === SCRATCH_SCRIPT.id ? -1 : 0,
        }));
      },
      onSelect: (newVal) => {
        const newScript = scripts.get(newVal);
        const selectedVis = parseVisSilently(newScript.vis);
        setScriptAndArgs({
          ...newScript,
          visString: newScript.vis,
          vis: selectedVis,
        }, args);
      },
    });

    return { entityBreadcrumbs, argBreadcrumbs };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [loading, script?.id, collapsedClusterIds, collapsedScriptIds, args, selectedCluster, getCompletions]);

  if (loading) {
    return (<div>Loading...</div>);
  }

  return (
    <Paper className={classes.root} elevation={2}>
      <PixieCommandIcon fontSize='large' className={classes.pixieIcon} />
      <div className={classes.verticalLine} />
      <div className={classes.breadcrumbs}>
        <Breadcrumbs
          breadcrumbs={entityBreadcrumbs}
        />
        <div className={classes.spacer} />
        <div>
          <Breadcrumbs
            breadcrumbs={argBreadcrumbs}
          />
        </div>
      </div>
      <ExecuteScriptButton />
    </Paper>
  );
};

export default withStyles(styles)(LiveViewBreadcrumbs);
