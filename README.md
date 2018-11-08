# golearn
Learn golang.


# web开发
- 包放到$GOPATH/src中，用import cblog/models的形式导入，不喜欢import github.com/caitinggui/cblog/models这一串
- glide 要设置好mirror，并且在一个空白项目中先glide install 这些官方包，目的是为了建立缓存
```
glide mirror set https://golang.org/x/mobile https://github.com/golang/mobile --vcs git
glide mirror set https://golang.org/x/crypto https://github.com/golang/crypto --vcs git
glide mirror set https://golang.org/x/net https://github.com/golang/net --vcs git
glide mirror set https://golang.org/x/tools https://github.com/golang/tools --vcs git
glide mirror set https://golang.org/x/text https://github.com/golang/text --vcs git
glide mirror set https://golang.org/x/image https://github.com/golang/image --vcs git
glide mirror set https://golang.org/x/sys https://github.com/golang/sys --vcs git
glide mirror set https://golang.org/x/lint https://github.com/golang/lint --vcs git
glide get github.com/golang/mobile github.com/golang/crypto github.com/golang/net github.com/golang/tools github.com/golang/text github.com/golang/image github.com/golang/sys github.com/golang/lint 
```
- 在glide设置好ignore等.ignore表示忽略的包，excludeDirs表示忽略的目录,比如
```
package: cblog
ignore:
- models
- golang.org/x/sys/unix
- database/sql
excludeDirs:
- models
import:
- package: github.com/cihub/seelog
  version: v2.6
- package: github.com/gin-gonic/gin
  version: v1.3.0
- package: github.com/gin-contrib/sessions
  subpackages:
  - cookie
```
