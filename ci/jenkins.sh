#!/bin/bash

# latest_prod_version returns the latest version deployed in production, e.g. v0.9.0.
function latest_prod_version() {
  git tag --list "v*.*.*" --sort "version:refname" | tail -n 1 | cut -f 2 -d '-'
}

# minor_version_bump bumps a given app version in the format of "v<major>.<minor>.<patch>" to
# "v<major>.<minor+1>.0"
function minor_version_bump() {
  tag=$1

  echo ${tag} | awk -F \. '{ print $1"."$2+1".0" }'
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

# do_tag_work starts the main process to tag vungo release. It is different from VSL that the basic development branch
# is master, not dev. So it is relative simpler than VSL. No need to create new pull request, no approve needed.
function do_tag_work() {
  CURRENT_VERSION=$(latest_prod_version)
  NEXT=$(minor_version_bump $(latest_prod_version))

  git log --graph --decorate --oneline -5
  echo "start to tag current version is ${CURRENT_VERSION}"
  push_git_tag ${NEXT} git@github.com:Vungle/vungo.git || true
  echo "push tag to remote repo ${NEXT}"
}

do_tag_work