# 介绍
从头到尾构建一个名为Greenlight的应用，用于检索和管理电影相关信息的JSON API。

<table style="width: 100%;">
    <tr>
        <th align="center">Method</th>
        <th align="center">URL Pateern</th>
        <th align="center">Action</th>
    </tr>
    <tr>
        <th align="center">GET</th>
        <th align="center">/v1/healthcheck</th>
        <th align="center">展示应用运行情况和版本信息</th>
    </tr>
    <tr>
        <th align="center">GET</th>
        <th align="center">/v1/movies</th>
        <th align="center">显示所有电影的详细信息</th>
    </tr>
    <tr>
        <th align="center">POST</th>
        <th align="center">/v1/movies</th>
        <th align="center">创造新的影片</th>
    </tr>
    <tr>
        <th align="center">GET</th>
        <th align="center">/v1/movies/:id</th>
        <th align="center">显示特定电影的详细信息</th>
    </tr>
    <tr>
        <th align="center">PATCH</th>
        <th align="center">/v1/movies/:id</th>
        <th align="center">更新特定电影的详细信息</th>
    </tr>
    <tr>
        <th align="center">DELETE</th>
        <th align="center">/v1/movies/:id</th>
        <th align="center">删除特定电影</th>
    </tr>
    <tr>
        <th align="center">POST</th>
        <th align="center">/v1/users</th>
        <th align="center">注册一个新用户</th>
    </tr>
    <tr>
        <th align="center">PUT</th>
        <th align="center">/v1/users/activated</th>
        <th align="center">激活特定用户</th>
    </tr>
    <tr>
        <th align="center">PUT</th>
        <th align="center">/v1/users/password</th>
        <th align="center">更新特定用户的密码</th>
    </tr>
    <tr>
        <th align="center">POST</th>
        <th align="center">/v1/tokens/authentication</th>
        <th align="center">生成新的身份验证令牌</th>
    </tr>
    <tr>
        <th align="center">POST</th>
        <th align="center">/v1/tokens/password-reset</th>
        <th align="center">生成新的密码重置令牌</th>
    </tr>
    <tr>
        <th align="center">GET</th>
        <th align="center">/debug/vars</th>
        <th align="center">显示应用程序指标</th>
    </tr>
</table>

在本项目完成时，通过GET获取电影信息，会返回以下内容响应。

``` 
$ curl -H "Authorization: Bearer RIDBIAE3AMMK5716IAEBUGA7XQ" localhost:4000/v1/movies/1
{
    "movie":{
        "id": 1,
        "title": "Moana",
        "year": 2016,
        "runtime": "107 mins",
        "genres": [
            "animation",
            "adventure"
        ],
        "version": 1
    }
}
```

本书使用PostgreSQL作为存储的数据库。

这个项目作者是Alex Edwards，一位全栈式 Web 开发人员和作家。下面是他的blog和github网页地址。
<div align="center">
    <a href="https://github.com/alexedwards/">github</a>·
    <a href="https://www.alexedwards.net/blog">blog</a>
</div>

# 运行环境软件的安装
- pycharm运行项目实例
- Go的主要版本（1.23版）「可通过输入go version查询」
- curl工具，用于处理终端HTTP请求和响应。「可输入curl --version查询」
- hey工具，用于执行负载测试
- git版本控制系统
- chrome浏览器，易于开发的web浏览器
- 文本编辑器