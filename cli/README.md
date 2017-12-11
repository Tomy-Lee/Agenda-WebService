
## agenda help

列出命令说明

## agenda register

注册 Agenda 账号

### 语法

agenda register [flags]

- -m, --email string      邮箱
- -h, --help              帮助
- -p, --password string   密码
- -n, --phone string      电话
- -u, --username string   用户名

### 实例

agenda register -u test -p testpass -m test@test.com -n 12345678909

## agenda login

登录 Agenda 账号

### 语法

agenda login [flags]

- -h, --help              帮助
- -p, --password string   密码
- -u, --username string   用户名

### 实例

agenda login -u test -p testpass

## agenda logout

登出 Agenda 账号

### 语法

agenda logout [flags]

- -h, --help              帮助

### 实例

agenda logout

## agenda list

登录后可查看已注册的所有用户的用户名、邮箱及电话信息

### 语法

agenda list [flags]

- -h, --help              帮助

### 实例

agenda list

## agenda delete-account

删除本用户账户

### 语法

agenda delete-account [flags]

- -h, --help              帮助

### 实例

agenda delete-account

## agenda add-meeting

添加会议

### 语法

agenda add-meeting [flags]

- -t, --title string          会议主题
- -p, --participator string   参加者（用多个-p 添加多个参加者）
- -s, --startTime string      开始时间（格式XXXX-XXX-XX/XX:XX:XX, 24小时制）
- -e, --endTime string        结束时间（格式XXXX-XXX-XX/XX:XX:XX, 24小时制）
- -h, --help                  帮助

### 实例

agenda add-meeting -t Shadowsocks -p clowwindy -s 2017-01-01/12:00:00 -e 2017-01-01/13:00:00

## agenda add-participator

添加会议参与者

### 语法

agenda add-participator [flags]

- -t, --title string          会议主题
- -p, --participator string   参加者（用多个-p 添加多个参加者）
- -h, --help                  帮助

### 实例

agenda add-participator -t Shadowsocks -p clowwindy

## agenda delete-participator

删除会议参与者

### 语法

agenda delete-participator [flags]

- -t, --title string          会议主题
- -p, --participator string   参加者（用多个-p 添加多个参加者）
- -h, --help                  帮助

### 实例

agenda delete-participator -t Shadowsocks -p clowwindy

## agenda show

查询会议

### 语法

agenda show [flags]

- -s, --startTime string      开始时间（格式XXXX-XXX-XX/XX:XX:XX, 24小时制）
- -e, --endTime string        结束时间（格式XXXX-XXX-XX/XX:XX:XX, 24小时制）
- -h, --help                  帮助

### 实例

agenda -s 2017-01-01/12:00:00 -e 2017-01-01/13:00:00

## agenda cancel-meeting

取消会议

### 语法

agenda cancel-meeting [flags]

- -t, --title string          会议主题
- -h, --help                  帮助

### 实例

agenda cancel-meeting -t Shadowsocks

### agenda quit-meeting

退出会议

### 语法

agenda quit-meeting [flags]

- -t, --title string          会议主题
- -h, --help                  帮助

### 实例

agenda quit-meeting -t Shadowsocks

## agenda clear

清空自己发起的所有会议安排

### 语法

agenda clear [flags]

- -h, --help                  帮助

### 实例

agenda clear
