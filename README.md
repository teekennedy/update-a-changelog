# Update a Changelog

Automates the addition of CHANGELOG.md entries according to [Keep a Changelog].
Maintain a Changelog lets you focus on the part of keeping a changelog that
matters - documenting notable changes as they are introduced. Never deal with
changelog-related merge conflicts, approved PRs needing re-review after
resolving these conflicts, or changelog entries getting inserted under the
wrong version again!

Specifically, this action enables the following workflow:

1. A user makes a Pull Request and documents the noteworthy changes in the PR
   as a list of typed entries:
   * Added: for new features.
   * Changed: for changes in existing functionality.
   * Deprecated: for soon-to-be removed features.
   * Removed: for now removed features.
   * Fixed: for any bug fixes.
   * Security: in case of vulnerabilities.
1. When the PR is merged, the Maintain a Changelog action will parse the
   description and add changelog entries to the appropriate section of
   CHANGELOG.md.

It's recommended that you install this action along with [Keep a Changelog -
New Release]. [Keep a Changelog - New Release] will automatically update
CHANGELOG.md when a new version tag is pushed. Between these two actions, you
can ensure your CHANGELOG.md is always correct and up-to-date without ever
having to edit it directly.

## Motivation

I wanted a changelog management solution that:

- Removes the pain points of manual changelog management.
- Is fully compatible with [Keep a Changelog], both in rules and in spirit.
- Maintains the same level of quality as a manually curated changelog.
- Works with a variety of development workflows and contribution processes.
- Does not require contributors to install software or run scripts.

### Manual changelog management

The [Keep a Changelog] process is a great way for users and contributors to
keep track of all notable changes have been made to a project between each
release. However, in practice it has a couple pain points that require manual
intervention.

First, anytime you have more than one Pull Request that adds a new changelog
entry, and one of them gets merged, you often end up with merge conflicts in
the other open Pull Requests. Resolving these conflicts is a chore that, in the
worst case, scales _quadratically_ with respect to the number of open PRs.

If the repository requires PR reviews and has enabled the setting to dismiss
stale PR reviews when new commits are pushed, any open PRs that have been
previously approved will have to be re-approved after the merge conflicts are
fixed. This expands the maintenance burden to involve multiple people.

Second, anytime a new version is released, all unreleased changes in
CHANGELOG.md need to be moved to a new release version section as well as
copied into the GitHub release notes for that version. Once again, all open
Pull Requests that contain changelog entries will need to be updated, otherwise
the entries will be incorrectly added under the released section of
CHANGELOG.md due to the way git handles the merge.

Update a Changelog avoids changes to any git-managed piece of information,
which makes all of the above problems irrelevant.

### Other GitHub Actions

Other changelog-related GitHub Actions automate the management of CHANGELOG.md
by generating it from other sources, such as commit messages or issue titles.
While it may be tempting to build your changelog from already available
information, it is against the spirit of [Keep a Changelog].

The use of commit messages in changelogs is specifically called out as a [bad
practice]:

> Using commit log diffs as changelogs is a bad idea: they're full of noise.
> Things like merge commits, commits with obscure titles, documentation
> changes, etc.
>
> The purpose of a commit is to document a step in the evolution of the source
> code. Some projects clean up commits, some don't.

Using issue or PR titles isn't that much better. They may contain multiple
noteworthy changes, or none at all. By generating your changelog from
potentially unrelated sources, you'd be allowing the quality of your release
notes to suffer for the sake of convenience.

Release notes are for humans, and communicating effectively with humans
requires a conscious effort. The perfect time to think about how your changes
should be communicated is when drafting notes in a Pull Request. You're already
in the mindset of reflecting on your freshly minted code, and are likely
summarizing the changes for reviewers anyway. Update a Changelog capitalizes
on this mindset as an opportunity to craft some quality changelog entries.

[Keep a Changelog - New Release]: https://github.com/marketplace/actions/keep-a-changelog-new-release
[Keep a Changelog]: https://keepachangelog.com/en/1.1.0/
[Semantic Version]: https://semver.org/
[bad practices]: https://keepachangelog.com/en/1.1.0/#bad-practices
