# genpass

[文艺青年如何管理密码](https://linuxtoy.org/archives/art-of-password.html) 的golang 实现。

线性同余方法中的A、B、M常量使用了gcc中的常数。

## 特性

[x] 命令行版本
[ ] web版本

## 命令行版本

NAME:
   gen - Generate password now

USAGE:
   command gen [command options] [arguments...]

DESCRIPTION:
   Generate your password.

OPTIONS:
   --user, -u 'username'	User name about website.
   --domain, -d 'baidu.com'	The domain about your password used for.
   --flag, -f 'alpha'		Which chars should include in password, valid option is:
				upper|lower|alpha|digit|punct|xdigit|alpha|alnum|graph
   --extra, -e 			Which extra chars can used for part of password.
   --method, -m 'sha256'	Which method should use, valid options is: md5|sha1|sha256|sha512
   --reversion, -r '0'		Password version, for update password.
   --length, -l '8'		Password length, default is 8.
