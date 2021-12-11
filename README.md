# GECK

## Elevator Pitch

A dotfile manager that tracks changes to local files and other resources that you care about and keeps them in sync with a remote repo.
If you need to nuke your distro, geck will restore the dotfiles that you told it to care about from the remote repo afterwards.

Basically a barebones Ansible with the following differences:

It only configures the host that it is installed on.
No ssh, no agents.

It configures itself based on the resources that you ask it to track, instead of you writing yaml that gets deployed.
For workstations and dev machines, this simplifies the configuration process. You just configure your workstation as you would normally, tell it which files and resources are important, and it does the rest.

Because you should never have to manually edit the geck repo, it is encrypted using pbkdf2 by default using a password that you specify.

## Envisioned usage:

### Repository Management
```shell
geck init|set ~/geck/ 
geck clone ~/geck/ https://github.com/foo/bar
```
geck init creates a new local geck repo. If the specified directory does not exist, it will be created.
If the specified directory exists and is not empty, geck will refuse to use it.

geck set configures geck to use an existing directory as its geck repo

geck clone will create a new local repo by cloning the specified upstream.

### Resource Management
```shell
geck file|user|group|plugin|<plugin> add|remove
```

```shell
geck file add ~/.zshrc
```
Geck adds the given file to a list of files, which it syncs from your system to a specified git repo. Geck tracks the contents, mode and ownership of the file.

```shell
geck user add foo
```
Geck adds the given user to a list of users, which it syncs from /etc/passwd and /etc/shadow. It tracks the name, gecos, uid, gid and other state of the user.

```shell
geck group add bar
```
Geck adds the given group to a list of groups, which it syncs from /etc/group.

### Deployment
```shell
geck read|write|log|show
```

```shell
geck read
```
Geck iterates through the resources it was told to care about and ensures that it's repo is up to date with what is currently on the system

```shell
geck write
```
Geck deploys the resources it knows about to your system.

## Useful Resources

* [git2go tutorial](https://blog.gopheracademy.com/advent-2014/git2go-tutorial/)