* simplify fsnotify. support only inotify, kqueue and ReadDirectoryChangesW
* diff per line, diff per function and both combined (for merging)


check diffkit.org and data diffing
see https://en.wikipedia.org/wiki/VCDIFF


data diffing for YAML, JSON, INI

https://en.wikipedia.org/wiki/INI_file

implement 

> git bundle create your_name.bundle --all





<!-- * **strict RBAC** (requires all access through roles, and permissions are connected only to roles, not directly to users.)
* integrated **tickets** and **wiki/docs**

* tracking of **large binary data**
* **tracking of directories**
* **tracking of folder and file permissions**

* write or hide on commit, do not delete!
* Archive and send a collection of changesets as a series of patch by emails
* diff lines OR functions. 
* data diffing for YAML, JSON, INI -->



<!-- use/investigate https://github.com/Microsoft/language-server-protocol to specify languages
or https://github.com/AnanthaRajuCprojects/Reserved-Key-Words-list-of-various-programming-languages 


PYTHON: 
  def STRING( STRING? ):
    STRING
    return
C:
  STRING STRING( STRING? ) { STRING }
GO:
  func STRING( STRING? ) STRING? { STRING }
JAVA:
  STRING(?) STRING( STRING? ) {  -->


<!-- ## Examples  TODO: make better workflow examples

### Simple workflow (single person)

```shell
gvs #will init gvs
gvs name "initial save"
#change some files
gvs name "my temp save $(date)"
gvs stat/log/history
gvs archive "my-last-fix_dev_0.3"
gvs push
# will push to master
```

### Multi workflow (several persons)

*developer1*

```shell
gvs clone <url>
#change some files
gvs name "my save"
gvs branch "br1"
gvs push
```

*developer2*

```shell
gvs clone <url>
#change some files
gvs name "my save"
gvs branch "br2"
gvs push
```
<!-- 
*merging lines*

```shell
gvs branch list
# br1
# br2
# br3-playground
gvs branch compare br1 br2
# br1.1 (file1)
#   changed_lines: 1,5, 22-55, 75
# br2.1 (file1)
#   changed_lines: 4,5, 22-55, 75
#
# br1.2 (file2)
#   changed_lines: 15, 25-55, 99
gvs branch from br1 br2 
gvs branch "merged br1 br2"
gvs pick br1.1

```

*merging functions*

```shell
gvs branch list
# br1
# br2
gvs branch compare br1 br2
# br1.1
#   -
#
#
# br2.1
#
# br2.2
#

``` -->

<!--
## Other

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

--> -->
