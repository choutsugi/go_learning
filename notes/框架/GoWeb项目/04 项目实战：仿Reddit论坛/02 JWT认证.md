

Access Token与Refresh Token

- Access Token：
  - 5分钟有效期。
  - 存储在客户端。
  - 如果Access Token过期，服务器通过检查数据库中Refresh Token是否有效以重新生成AccessToken。

- Refresh Token：
  - 14天有效期。
  - 存储在数据库。
  - 每重新生成AccessToken后更新RefreshToken有效期（可选）。
  - 用户登出后删除Refresh Token。
  - 黑名单用户删除Refresh Token。







