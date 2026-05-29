This file is a merged representation of the entire codebase, combined into a single document by Repomix.

# File Summary

## Purpose
This file contains a packed representation of the entire repository's contents.
It is designed to be easily consumable by AI systems for analysis, code review,
or other automated processes.

## File Format
The content is organized as follows:
1. This summary section
2. Repository information
3. Directory structure
4. Repository files (if enabled)
5. Multiple file entries, each consisting of:
  a. A header with the file path (## File: path/to/file)
  b. The full contents of the file in a code block

## Usage Guidelines
- This file should be treated as read-only. Any changes should be made to the
  original repository files, not this packed version.
- When processing this file, use the file path to distinguish
  between different files in the repository.
- Be aware that this file may contain sensitive information. Handle it with
  the same level of security as you would the original repository.

## Notes
- Some files may have been excluded based on .gitignore rules and Repomix's configuration
- Binary files are not included in this packed representation. Please refer to the Repository Structure section for a complete list of file paths, including binary files
- Files matching patterns in .gitignore are excluded
- Files matching default ignore patterns are excluded
- Files are sorted by Git change count (files with more changes are at the bottom)

# Directory Structure
```
js-map/webview/assets/src/auth/use-codex-cloud-access.ts
```

# Files

## File: js-map/webview/assets/src/auth/use-codex-cloud-access.ts
```typescript
import { useQuery } from "@tanstack/react-query";
import { useAtom } from "jotai";
import type { ExtendedAuthMethod } from "protocol";
import { useCallback, useEffect } from "react";

import { useEnvironments } from "@/codex-api";
import { useGate } from "@/statsig/statsig";
import { logger } from "@/utils/logger";
import { persistedAtom } from "@/utils/persisted-atom";
import { QUERY_STALE_TIME } from "@/utils/query-stale-times";
import { CodexRequest } from "@/utils/request";
import type { AccountPlanType } from "@/utils/skus";
import { isEnterpriseyPlan } from "@/utils/skus";
import { useFetchFromVSCode } from "@/vscode-api";
import { FetchError } from "@/web-fetch-wrapper";

import { useAuth } from "./use-auth";

export type CodexCloudAccess =
  | "loading"
  | "enabled"
  | "enabled_needs_setup"
  | "disabled"
  | "error";

/**
 * Override this locally in dev or in stories to test components with different Codex access states.
 *
 * Do not commit this with anything other than undefined.
 */
const STORYBOOK_DEV_CODEX_ACCESS_OVERRIDE: CodexCloudAccess | undefined =
  undefined;

export const aCodexCloudAccessAtom = persistedAtom<CodexCloudAccess | null>(
  "codexCloudAccess",
  null,
);

export function useCodexCloudAccess(): {
  access: CodexCloudAccess;
  refetch: () => Promise<void>;
} {
  const { authMethod } = useAuth();
  const [cachedCodexCloudAccess, setCachedCodexCloudAccess] = useAtom(
    aCodexCloudAccessAtom,
  );
  const storybookOverride =
    (__STORYBOOK__ || __DEV__) && STORYBOOK_DEV_CODEX_ACCESS_OVERRIDE
      ? STORYBOOK_DEV_CODEX_ACCESS_OVERRIDE
      : undefined;
  const {
    data: accountInfo,
    isLoading: accountInfoLoading,
    isError: accountInfoError,
    refetch: refetchAccountInfo,
  } = useFetchFromVSCode("account-info", {
    queryConfig: {
      enabled: authMethod === "chatgpt",
    },
  });

  const plan = (accountInfo?.plan ?? undefined) as AccountPlanType | undefined;
  const enterprisey = isEnterpriseyPlan(plan);
  const {
    data: workspaceSettings,
    isLoading: workspaceSettingsLoading,
    isError: workspaceSettingsError,
    refetch: refetchWorkspaceSettings,
  } = useQuery({
    queryKey: ["accounts", "settings", accountInfo?.accountId],
    enabled:
      !!accountInfo?.accountId && enterprisey && authMethod === "chatgpt",
    queryFn: async () =>
      CodexRequest.safeGet("/accounts/{account_id}/settings", {
        parameters: { path: { account_id: accountInfo?.accountId ?? "" } },
      }),
    staleTime: QUERY_STALE_TIME.ONE_MINUTE,
  });

  const {
    data: environments,
    isLoading: environmentsLoading,
    error: environmentsError,
    refetch: refetchEnvironments,
  } = useEnvironments({
    enabled: authMethod === "chatgpt",
  });
  const isWorktreeSnapshotsEnabled = useGate(
    __statsigName("codex_worktree_snapshots"),
  );

  const refetch = useCallback(async (): Promise<void> => {
    await Promise.all([
      refetchAccountInfo(),
      enterprisey && refetchWorkspaceSettings(),
      refetchEnvironments(),
    ]);
  }, [
    refetchAccountInfo,
    refetchWorkspaceSettings,
    refetchEnvironments,
    enterprisey,
  ]);

  const isLoading =
    accountInfoLoading || workspaceSettingsLoading || environmentsLoading;
  // API may 404 instead of returning an empty list if there are no errors.
  const isTaskError404 =
    environmentsError instanceof FetchError && environmentsError.status === 404;
  const hasErrors =
    accountInfoError ||
    (enterprisey && workspaceSettingsError) ||
    (!!environmentsError && !isTaskError404);
  const needsOnboarding = isWorktreeSnapshotsEnabled
    ? isTaskError404
    : environments?.length === 0 || isTaskError404;
  const hasWorkspaceEnabledCodex =
    !enterprisey || (workspaceSettings?.beta_settings?.wham_access ?? false);
  const access = codexCloudAccess(plan, authMethod, {
    isLoading,
    hasErrors,
    needsOnboarding,
    hasWorkspaceEnabledCodex,
  });

  useEffect(() => {
    if (access !== "loading" && access !== "error") {
      setCachedCodexCloudAccess(access);
    }
  }, [access, setCachedCodexCloudAccess]);

  if (storybookOverride) {
    return { access: storybookOverride, refetch };
  }

  if ((access === "loading" || access === "error") && cachedCodexCloudAccess) {
    return { access: cachedCodexCloudAccess, refetch };
  }

  return { access, refetch };
}

export function codexCloudAccess(
  planType: AccountPlanType | undefined,
  authMethod: ExtendedAuthMethod | null,
  {
    isLoading,
    hasErrors,
    needsOnboarding,
    hasWorkspaceEnabledCodex,
    hasLoggedDisabledRef,
  }: {
    isLoading: boolean;
    hasErrors: boolean;
    needsOnboarding: boolean;
    hasWorkspaceEnabledCodex: boolean;
    hasLoggedDisabledRef?: React.RefObject<boolean>;
  },
): CodexCloudAccess {
  const enterprisey = isEnterpriseyPlan(planType);
  const maybeLogOnce = (message: string): void => {
    if (hasLoggedDisabledRef && !hasLoggedDisabledRef.current) {
      logger.info(message);
      hasLoggedDisabledRef.current = true;
    }
  };

  if (authMethod !== "chatgpt") {
    maybeLogOnce(
      "Codex Cloud access disabled because user is not logged in via ChatGPT.",
    );
    return "disabled";
  } else if (isLoading) {
    return "loading";
  } else if (hasErrors) {
    return "error";
  } else if (enterprisey && !hasWorkspaceEnabledCodex) {
    maybeLogOnce(
      "Codex Cloud access disabled because workspace has not enabled Codex.",
    );
    return "disabled";
  }
  return needsOnboarding ? "enabled_needs_setup" : "enabled";
}
```
