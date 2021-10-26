# common

公共模块

- 微信小程序登录与授权获取用户信息
    - 客户端调用 [wx.login](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/login/wx.login.html) 获取到code;
    - 将code传回给服务器,服务器调用 [auth.code2Session](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html)
    方法,获取到openID和sessionKey;
    - sessionKey保存在服务器,服务器返回一个token给客户端,保证sessionKey与token的映射关系;
    - 客户端获取用户授权后,将token和[UserInfo](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/user-info/wx.getUserInfo.html)
    用户信息返回给服务器,服务器根据sessionKey和用户数据中的信息进行解密
    - 对解密后的用户,没有的话注册用户...