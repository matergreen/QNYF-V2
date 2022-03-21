# 本 README 为Web界面的使用

## Start

使用一下命令编译
```shell
go build web.go
```
使用 `./web` 将在本地的 `8077端口` 开启一个 web 服务.如果您需要更换启动端口,可以修改`web.go`中的`第70行`用户只需要按照下图所示的提示填入指定内容,点击登录就好.
![web](../img/web.png)

> 注: 前端比较菜,使用的模板,可能登录失败后需要手动返回到界面再次进行信息填写.
> 登录成功后会将相关信息写在当前目录下的 uidList 文件里,此文件请不要删除