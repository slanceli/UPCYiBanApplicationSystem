+ **/login**  
    该接口用于登录  
    请求方法：POST  
    发送表单，格式如下：
    ```
    key1:"username"
    value1:"学号"
    key2:"password"
    value2:"数字石大密码"
    ```
    登陆成功后返回字符串“successful”，失败返回“failed”
+ **/application**  
    该接口用于提交报名信息  
    请求方法：POST  
    发送json，格式如下：  
    ```json
    {
      "name": "姓名",
      "phome_num": "电话号码",
      "gender":"性别",
      "mail_address": "邮箱地址",
      "political_face": "政治面貌",
      "class": "专业班级",
      "first_volunteer": "第一志愿",
      "second_volunteer": "第二志愿",
      "transfers": "是否服从调剂",
      "profile": "个人简介",
      "advantage": "个人对该岗位的优势",
      "cognition": "个人对该岗位的认知和思路",
      "review_comments": "审核意见（可以提交空白）"
    }
    ```
    **以上字段的类型都是字符串类型。**  
    若未登录则返回字符串“未登录”  
    成功返回字符串“successful”  
    失败返回字符串“failed”  
    