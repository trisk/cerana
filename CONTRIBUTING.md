# Contributing to Cerana

## Reporting Bugs and Issues

All bugs, issues, and feature requests are tracked in [Github Issues](https://github.com/cerana/cerana/issues). For a better overview of the project management, you can use the [Zenhub](https://www.zenhub.io/) browser extension.

Check whether the bug of feature has already been reported by searching the [issue database](https://github.com/cerana/cerana/issues). If an open issue is found, click the `+1` button on the first post and comment if you have additional information to share. If this has not yet been reported, follow the template provided when creating a [new issue](https://github.com/cerana/cerana/issues/new)

## Patches

1. Make sure an [issue](https://github.com/cerana/cerana/issues) exists, or create one. It is best to discuss new features and enhancements before implementing.
2. Fork the repository and create a feature branch named `xxxx-something`, where xxxx is the number of the corresponding issue. Branches should be made off of `master`.
3. Ensure the style and testing guidelines are met
4. Use `git rebase -i` (and `git push -f`) to clean up the branch history into working, logical chunks (e.g. remove WIP commits, combine commits and subsequent minor fixups, etc.).
4. Create a Pull Request via [Github](https://github.com/cerana/cerana) following the template provided.
5. Check that the PR merges cleanly. If not use `git rebase master` and fix any issues.
6. Make sure the [TravisCI](https://travis-ci.org/cerana/cerana/pull_requests) tests pass for the PR.
7. Follow along with code review discussion. Comment in the PR after pushing any fixes.
8. When all is in order and the maintainers agree, the PR will be merged.

## Conventions

### Documentation

All exported methods and variables should have an accompanying comment following the [godoc](http://blog.golang.org/godoc-documenting-go-code) format. Regenerate READMEs using `make godocdown`.

### Code
For code consistency and catching minor issues, please run the following on all Go code and fix accordingly:

* `gofmt -s ./...`
* `goimports ./...`
* `golint ./...`
* `errcheck ./...`
* `go vet ./...`
* `go vet --shadow ./...`

### Testing
Submit unit tests for all changes. Before submitting a pull request, make sure the full test suite passes, using `go test ./...`

TravisCI is configured to run a build on pull request. Results can be found [here](https://travis-ci.org/cerana/cerana)

### Dependencies

All go dependencies are vendored and managed using [glide](https://github.com/Masterminds/glide). Use `glide install` to get the current set of dependencies. When introducing a new dependency `glide get [package name]` and commit both the updated `glide.yaml` and `glide.lock`. See the `glide` README for for more information.

## Contact

* Internet Relay Chat (IRC) - `#cerana` on `irc.freenode.net`