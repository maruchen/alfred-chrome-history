# alfred-chrome-history

Search Chrome history from Alfred and open in browser.

## Features

- Search Chrome history (title and URL)
- Support another Chrome profile

## Installation

Clone and `make dist` or just download [binary releases](https://github.com/pasela/alfred-chrome-history/releases).

```sh
git clone https://github.com/pasela/alfred-chrome-history.git
cd alfred-chrome-history
make dist
open alfred-chrome-history.alfredworkflow
```

## Usage

in Alfred:

```
ch {query}
```

## Use another Chrome profile

1. Open workflow `Chrome History` in Alfred Workflows tab.
2. Open Workflow Configuration dialog by upper right side button.
3. Set `CHROME_PROFILE` variable with your Chrome profile directory name or path such as `Profile 1`.

## License

MIT

## Author

Yuki (a.k.a. pasela)

## 开发说明
`make build`编译输出到`build/alfred-chrome-history`
`make dist`会先执行`make build`，然后拷贝到当前目录`alfred-chrome-history.alfredworkflow`

可执行文件既可以单独运行，也可以以alfred插件的形式运行。后者依赖 `github.com/deanishe/awgo`，一个alfred的go开发库。

## 原理
从环境变量`CHROME_PROFILE`或者启动参数获取Chrome Profile目录地址，从中找到history文件。  
history文件是 sqlite3 数据库，包含 id、url、title、visit_count、typed_count、last_visit_time等字段


## TODO
[] feishu/docs 和 sheet 对?后面的去重
[] 保存favicon
