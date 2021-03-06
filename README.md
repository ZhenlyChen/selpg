# selpg 

A simple CLI program

## Usage

```bash
Usage of selpg:
  -e, --end int     End page (default -1)
  -f, --find        Find page breaks
  -l, --leap int    The count of row in each page (default 72)
  -s, --start int   Start page (default 1)
```

example: 

```bash
selpg -s10 -e20 -l66
selpg -s10 -e20 -l66 ./input_file
selpg -s10 -e20 -f
```



**`-sNumber`和`-eNumber`强制选项：**
selpg 要求用户用两个命令行参数`-sNumber`（例如，`-s10`表示从第 10 页开始）和`-eNumber`（例如，`-e20`表示在第 20 页结束）指定要抽取的页面范围的起始页和结束页。selpg 对所给的页号进行合理性检查；换句话说，它会检查两个数字是否为有效的正整数以及结束页是否不小于起始页。这两个选项，`-sNumber`和`-eNumber`是强制性的，而且必须是命令行上在命令名 selpg 之后的头两个参数：

```bash
$ selpg -s10 -e20 ...
```



**`-lNumber`和`-lNumber`可选选项：**
selpg 可以处理两种输入文本：

*类型 1：*该类文本的页行数固定。这是缺省类型，因此不必给出选项进行说明。也就是说，如果既没有给出`-lNumber`也没有给出`-f`选项，则 selpg 会理解为页有固定的长度（每页 72 行）。

选择 72 作为缺省值是因为在行打印机上这是很常见的页长度。这样做的意图是将最常见的命令用法作为缺省值，这样用户就不必输入多余的选项。该缺省值可以用“-lNumber”选项覆盖，如下所示：

```bash
$ selpg -s10 -e20 -l66 ...
```

这表明页有固定长度，每页为 66 行。

*类型 2：*该类型文本的页由 ASCII 换页字符（十进制数值为 12，在 C 中用“\f”表示）定界。该格式与“每页行数固定”格式相比的好处在于，当每页的行数有很大不同而且文件有很多页时，该格式可以节省磁盘空间。在含有文本的行后面，类型 2 的页只需要一个字符 ― 换页 ― 就可以表示该页的结束。打印机会识别换页符并自动根据在新的页开始新行所需的行数移动打印头。

类型 2 格式由“-f”选项表示，如下所示：

```bash
$ selpg -s10 -e20 -f ...
```

该命令告诉 selpg 在输入中寻找换页符，并将其作为页定界符处理。

注：`-lNumber`和`-f`选项是互斥的。



**`-dDestination`可选选项：**
selpg 还允许用户使用`-dDestination`选项将选定的页直接发送至打印机。这里，`Destination`应该是 lp 命令`-d`选项（请参阅`man lp`）可接受的打印目的地名称。该目的地应该存在 ― selpg 不检查这一点。在运行了带`-d`选项的 selpg 命令后，若要验证该选项是否已生效，请运行命令`lpstat -t`。该命令应该显示添加到`Destination`打印队列的一项打印作业。如果当前有打印机连接至该目的地并且是启用的，则打印机应打印该输出。这一特性是用 `popen()` 系统调用实现的，该系统调用允许一个进程打开到另一个进程的管道，将管道用于输出或输入。在下面的示例中，我们打开到命令

```bash
$ lp -dDestination
```

的管道以便输出，并写至该管道而不是标准输出：

```bash
selpg -s10 -e20 -dlp1
```

## 设计说明

本程序使用了`pflag`解析命令



## 测试

测试文件：

`test.txt`

```tx
Hello
Hi

World
BB

GG
kk

```

### 1

读取第一页，每页两行

```bash
$  selpg .\test.txt -l2 -s1 -e1
Hello
Hi
```

### 2

读取1-2页，每页两行

```bash
$ selpg .\test.txt -l2 -s1 -e2
Hello
Hi
♀
World
```

### 3

读取1-2页，以换页符分割

```bash
$ selpg .\test.txt -f -s1 -e2
Hello
Hi
♀
World
BB
♀
```

### 4

读取1-2页，使用`-dlp1`输出

```bash
$ selpg ./test.txt -f -s1 -e1 -dlp1
```



