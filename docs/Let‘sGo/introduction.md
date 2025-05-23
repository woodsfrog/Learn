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

``` git
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
