# webhookgo

# github设置
进入github项目->Settings->Webhooks  
* Payload URL：github更新时会主动post消息到此url，我们仅处理/webhook方法，格式：url/webhook
* Content Type：选择application/json
* Secret：密码随意
* 最后选择Just the push event，勾选Active
点击按钮Update webhook
