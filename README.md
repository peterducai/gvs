# gvs

go/global/general versioning system

* ask user details if missing (no more git config or bazaar whoami)
* Dead simple & single binary with tons of features
* AutoSync - Reduces needless Merging and Forking
* integrated wiki, ticketing & bug tracking, embedded documentation, and Technical notes. 
* easily manage users and access to your repos

## Other

* Manage commit access to parts of a repo using control lists
* Edit, fold, drop changesets in the style of git rebase --interactive
* Track large binary files
* Send email to subscribed addresses to notify repository changes
* Send a collection of changesets as a series of patch emails
* Purge all files and dirs in the repository that are not being tracked
* Handle nested repositories
* Allow commands to affect multiple repositories simultaneously
* Count lines of source code
* List TODOs
* Contributor branch


what is different from GIT:

* integrated SSL server.. should be as easy as 
> python -m SimpleHTTPServer
* integrated RBAC
* send the patches by email (darcs send)

## Issues with other vcs

* Poor handling of binary data
* Submodules are very difficult to work with effectively, and are limited to including an entire
* Steep learning curve
* no diffs on binary files
* The repository is a bunch of files that can easily be ruined by some accident.
* Inconsistent command line interface
* Revert features with ease 