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
