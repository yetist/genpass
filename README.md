# genpass

[文艺青年如何管理密码](https://linuxtoy.org/archives/art-of-password.html) 的golang 实现。

线性同余方法中的A、B、M常量使用了gcc中的常数。

## 特性

[x] 命令行版本

[x] web版本

## 命令行版本

```
NAME:
   genpass - Generate Password

USAGE:
   genpass [global options]

VERSION:
   0.1

COMMANDS:
   server	Run http server to use web generate password.
   help, h	Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --primary, -p 'primary'		Primary password, or use user name about website.
   --description, -d 'description.com'	Description about the password, or use the website domain.
   --flag, -f 'alnum'			Which chars should include in password, valid option is:
					upper|lower|alpha|digit|punct|xdigit|alpha|alnum|graph
   --extra, -e 				Which extra chars can used for part of password.
   --method, -m 'sha256'		Which method should use, valid options is: md5|sha1|sha256|sha512
   --reversion, -r '0'			Password version, for update password.
   --length, -l '8'			Password length, default is 8.
   --help, -h				show help
   --version, -v			print the version
```

## Web版本

[Demo](http://moses.zhcn.cc:9999)
