## Git工具使用总结
rebase
## Git工具使用
create local branch
`git checkout -b your_branch`
track the remote branch by, so that you can pull the next time
`git push -u origin remote_branch`//track the remote branch with the same name
`git push --set-upstream origin remote_branch`
`git branch --set-upstream-to=origin/remote_branch`
you can also push to the anthoer untracked branch by
`git push origin HEAD:refs/heads/your_branch`

rebase is interesting
rebase相比merge最主要的作用是让分支更清晰，不把master的多个commit加到feature分支上，所以master像feature一般使用rebase
但是rebase和merge的共同作用是”合并”
合并过程中可能会有冲突，应该在本地rebase解决，而不是remote merge时候解决冲突；本地解决后push到remote feature branch，再merge到 remote master branch
