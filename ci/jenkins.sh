#!/bin/bash

type hub >/dev/null 2>&1 || { echo >&2 "hub is required. See https://hub.github.com/"; exit 1; }

# latest_prod_version returns the latest version deployed in production, e.g. v0.9.0.
function latest_prod_version() {
  git tag --list "v*.*.*" --sort "version:refname" | tail -n 1 | cut -f 2 -d '-'
}

# app_version returns the version of the main vungo service.
function app_version() {
  local image=$(awk '{print $NF}' ci/prod/version)
  echo ${image}
}

# minor_version_bump bumps a given app version in the format of "v<major>.<minor>.<patch>" to
# "v<major>.<minor+1>.0"
function minor_version_bump() {
  tag=$1

  echo ${tag} | awk -F \. '{ print $1"."$2+1".0" }'
}

# should_adjust_stage_version returns whether the app version should be adjusted based on the state
# of the development cycle.
function should_adjust_stage_version() {
  tag=$1

  # Staging is always one development cycle ahead of production, but depending on when the latest
  # release is deployed to production, the latest tag can be 1 or 2 minor versions behind. But, we
  # know for sure that if we have a `release-*` branch at the latest tag + 1 minor, we know stage
  # needs to be 1 minor version ahead of the latest release.
  [ -n "$(git branch -r | grep "origin\/$(release_branch_name ${tag})")" ]
}

# release_branch_name returns the name of release branch with the given app version.
function release_branch_name() {
  echo vungo-release-$1
}

# update_version updates the version for the given project and app version and
# creates the release branch locally.
function update_version() {
  local tag=$1
  echo "start update version"
  echo ${tag} > ci/prod/version

  # Create git branch and commit the change
  # git diff
  # git checkout -b $(release_branch_name ${tag})
  git add ci/prod/version
  echo "after add"
  git config user.email "platform-bj@vungle.com"
  git config user.name "Vungo CI by Jenkins"
  git commit -m "chore(vungo,release): Vungo version bump"
  git status
  echo "update version OK"
}

# push_git_tag pushes the app version to GitHub.
function push_git_tag() {
  tag=$1
  repo=$2 # Repo URL in HTTPS or SSH.

  git config user.email "platform-bj@vungle.com"
  git config user.name "Vungo CI by Jenkins"
  git tag -a ${tag} -m "Vungo Release ${tag}"
  git push -q ${repo} ${tag} || exit 1
}

function do_release_work() {
  git log --graph --decorate --oneline -5
  if [ $(latest_prod_version) != $(app_version) ]; then
    echo "Latest deployed version is different from the configured version, aborting release process!"
    exit 0
  fi

  NEXT=$(minor_version_bump $(latest_prod_version))
  update_version ${NEXT}
  # send_pull_request ${NEXT} git@github.com:Vungle/vungo.git
  # push_git_tag $(app_version) git@github.com:Vungle/vungo.git || true
  #https://${GITHUB_AUTH}@github.com/Vungle/vungo.git
}

do_release_work