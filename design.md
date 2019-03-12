# Design

## Outcome

* Stand-alone executable.. but modular code base!
* only PUSH/PULL.. no clone/checkout/fork.. these are just buzzwords for same thing.
* every change is per file and it's called PATCH
* commit can be one or several PATCHes.. if several, it's called PATCHLIST
* only SNAPSHOTS.. every commit, branch or fork is actually snapshot.
* PERMPATCH is patch containing changes in file permissions
* every SNAPSHOT contains TIMESTAMP, USER, HOST, PATCHLIST and PERMPATCH
* DEDUPLICATION! to minimize size of repo.
* you can pull specific PATCH from specific SNAPSHOT
* after mergin SNAPSHOT, don't delete/detach info about it.
* *gvs push all* will push to all servers (external repos)
* Cathedral vs. Bazaar... why not both?


## Ideas

* BRANCH is actually just SNAPSHOT in time
* commit info should contain TIMESTAMP, USER, diff HASH
* needs better model than SemVer (tracking functions? addition and deletions)
* no more.. error: You have local changes to 'frotz'; not switching branches.

## Ideas from others

* SourceDepot couldn't handle a project the size of Windows, so rather than having the whole operating system in a single repository, the Windows code was actually divided among 65 different repositories, with a kind of virtualization layer on top to produce a unified view of all the code. Some of these 65 repos contained nicely isolated, standalone components; others took vertical or horizontal slices through the operating system; others were just grab bags of different code. 

* autosync mode on: that means simply running fossil commit or fossil update will automatically push or pull files to the server. 
* hg serve
* a virtual file system that only downloads files to local storage as they are needed (GVFS)
* integrated code review 
* Easy collaboration by e-mail. If you want to add a feature or bugfix to some project, you can make a local clone, apply your changes, then send the patches by email (darcs send). The project’s maintainers can decide whether to accept or reject the patches. This way, you do not need commit privileges to contribute.
Parallel development. Let’s say you follow the development of an open-source project, and you have some controversial patches that aren’t accepted by the official maintainers. No problem – make your changes and publish your own repository. It’s a fork, of sorts, but it’s still connected to the mainline. Whenever the official project makes changes, you can do a darcs pull to get them, and resolve any conflicts. This way, your fork is kept up to date.
Cherry-picking. If you’ve ever worked on a team, you will know that somebody often has a change you want, but which can’t be committed to the trunk yet. With Darcs you can grab just the one change by pulling it into your repository.
Interactivity. Darcs enables you to be precise and say “yes” or “no” to every change that you can include in your patch. Thus you can really create minimal patches, or separate your work in several patches, each one doing a consistent change. Other commands, like darcs pull and darcs push, behave the same.
* The merging system is based on a pair of 3-way merges: one set-oriented one at the changeset level to resolve differences in tree layout (file and directory renames, for instance), and one line-oriented one at the file level, to resolve differences in concurrent edits to the same file. If either of these fail, they are passed off to a user-provided hook function, which invokes emacs ediff mode by default (but can be overridden).

Alternately, the merge conflicts can be saved to a file, and resolved asynchronously, then used to perform the merge.

It is important to note that a 3-way merge is not the same as simply "applying patches" in one order or another: we locate the least common ancestor of the merged children in our ancestry graph, calculate the edits on the left and right edges, adjust the right edge's edit coordinates based on the left edge's edits, and only then do we concatenate the left and right edges (ignoring identical changes, and rejecting conflicting ones).


* Fossil and Git promote different development styles. Git promotes a "bazaar" development style in which numerous anonymous developers make small and sometimes haphazard contributions. Fossil promotes a "cathedral" development model in which the project is closely supervised by an highly engaged architect and implemented by a clique of developers.
* One commentator has mused that Git records history according to the victors, whereas Fossil records history as it actually happened.
* The fossil push, fossil pull, and fossil sync commands do not provide the capability to push or pull individual branches. Pushing and pulling in Fossil is all or nothing. This is in keeping with Fossil's emphasis on maintaining a complete record and on sharing everything between all developers.