使用go创建web应用
课程主要内容:
处理请求
模版
中间件
存储数据
HTTPS，HTTP2
测试
部署

go mod init github.com/tiamxu/01web

2】处理Handle请求
process 处理过程
如何处理（Handle） Web请求
  http.Handle函数
  http.HandleFunc函数
创建Web Server 
  http.ListenAndServer()
    接受两个参数，第一个网络地址，第二个handler：nil表示默认为DefaultServeMux
  http.Server是一个struct
  server.ListenAndServer()
Handler:
  Handler是一个接口interface
  Handler定义了一个方法ServeHTTP()
    HTTPResponseWriter
    指向Request结构体的指针
  DefaultServeMux 多路复用器
请求->DefaultServeMux -> handler...

多个Handler-http.Handle
 不指定Server struct里面的handler值
 使用http.Handle将某个Handler注册到DefaultServeMux
    http包里有handle函数
    ServerMux struce里也有handle方法
 如果你调用http.Handle函数，实际上调用的是DefaultServeMux上的Handle方法
   DefaultServeMux是ServerMux的指针变量

http.Handle 第二个参数是Handler
http.HandleFunc 第二个参数是一个Handler函数
   http.HandlerFunc 可以把Handler函数转换为handler
   内部调用的还是http.Handle函数

3】内置的handler
 NotFoundHandler
 RedirectHandler
 StripPrefix
 TimeoutHandler
 FileServer
 http.FileServer(http.Dir("wwwroot"))

 4】请求Rquest
  http请求
  Request 
    Request是struct，代表了客户端发送的HTTP请求消息
    重要字段：
      URL、Header、Body、From、PostForm、MultiPartForm
    也可以通过Request的方法访问请求中的Cookie、URL、User Agent等信息
    Request既可代表发送到服务器的请求，又可代表客户端发出的请求
  URL
    URL Query http://www.xxx.com/post?id=123&thread_id=456
    r.URL.RawQuery会提供实际查询的原始字符串
    r.URL.Query 会提供查询字符串对应的map[string][]string
       GET方法，返回切片第一个值

  Header 
    请求和响应的header是通过Header类型描述的，是一个map
  Body 

 # 表单
   通过表单发送数据
    表单里的数据会以name-value对的形式，通过post请求发出去
    他的数据内容会放在POST请求的Body里面
    但name-value对在Body里面的格式是什么样的
    【表单Post请求的数据格式】
    通过POST发送的name-value数据对的格式可以通过表单的Context Type来指定，也就是enctype属性
    如何选择：enctype="multipart/form-data"  enctype="application/x-www-form-urlencoded" 
      简单文本：表单URL编码 enctype="application/x-www-form-urlencoded"
      大量数据：例如上传文件：mutipart-MIME
   Form字段
     Request上的函数允许我们从URL或/和Body中提取数据，通过这些字段
     Form PostForm MultipartForm 
     Form里面数据是key-value对 
     通常是先调用ParseForm或ParseMultipareForm来解析Request，然后相应的访问Form、PostForm或者MultipartForm字段
   PostForm字段
     对于URL和表单存在相同的key，那么他们都会放在一个slice里，表单的值靠前，PostForm只会返回表单数据
     PostForm 只支持application/x-www-form-urlencoded

   MultipartForm字段
      对于enctype="multipart/form-data"类型，需要使用ParseMultipartForm解析，返回一个结构体包含两个map，使用MultipartForm获取字段

   FormValue PostFormValue方法
     FormValue 方法会返回Form字段中指定的key对应的第一个value
     PostFormValue 一样，但是只返回表单的数据
      FormValue PostFormValue都会调用ParseMultipartForm方法
   文件Files
    enctype="multipart/form-data" 用来上传文件
    MulitpartForm.File 
    FormFile方法，无需调用ParseMultipartForm方法，返回指定key对应的第一个value，同时返回Fiel和FileHeader、错误信息。
   POST JSON
     不是所有的POST请求都来自Form
     客户端框架会以不同的方式对POST请求编码
        jQuery通常用application/x-www-form-urlencoded
        Angular是applicaton/json 
      ParseForm方法无法处理applicaton/json 

ResponseWriter
  从服务器向客户端返回响应需要使用ResponseWriter
  写入到ResponseWriter
    Write方法接收一个byte切片，然后写入到HTTP响应的Body里面
      如果在Write方法被调用时，header里面没有设定content type，那数据前512字节用来检测content type
    WriteHeader方法
      接收一个状态码作为参数，并把他作为HTTP响应的状态码返回
      如果没有显示调用，那么第一次调用Write方法前，会隐式调用WriteHeader(http.StatusOk)
      所以WriteHeader主要用来发送错误类状态码 
      调用完WriteHeader方法之后，仍然可以写入到ResponseWrite，但无法修改header了

模版主要内容：
简介
  什么是模版？ Web模版就是预先设计好的HTML页面，可以被模版引擎反复使用，来产生HTML页面
  GO的标准库提供了text/template 、html/template两个模版库
模版引擎 
  模版引擎可以合并模版与上下文数据、产生最终的HTML
  模版+数据 -->模版引擎 -->HTML文件
  GO的模版引擎使用text/template，HTML相关的部分使用了html/template。
  GO模版引擎工作原理：
    在web应用中，通常是有handler来触发模版引擎
    handler调用模版引擎，并将使用的模版传递给引擎，通常是一组模版文件和动态数据
    模版引擎生成HTML，并将其写入到ResponseWriter
    ResponWriter再将他加入到HTTP响应中，返回给客户端。
  关于模版：对于Web应用，通常是HTML
  text/template是通用模版引擎，html/template是HTML模版引擎
  使用模版引擎：
    1、解析模版源（可以是字符串或模版文件），从而创建一个解析好的模版的struct
    2、执行解析好的模版，并传入ResponseWriter和数据
       会触发模版引擎组合解析好的模版和数据，来产生最终的HTML文件，并传递给ResponseWriter
  解析模版：
    ParseFiles 
      解析模版文件，并创建一个解析好的模版struct，后续可以被执行
      ParseFiles函数是Template struct上ParseFiles方法的简便调用
      调用ParseFiles后，会创建一个新的模版，模版名字是文件名
      New函数
      ParseFiles的参数数量可变，但只返回一个模版，当解析多个文件时，第一个文件作为返回的模版（名、内容），其余内容作为map，供后续使用
    ParseGlob
      使用模式匹配来解析特定的文件
    Parse 
      可以解析字符串模版，其他方式最终都会调用Parse 
    Lookup方法：
      通过模版名来寻找模版，如果没找到就返回nil
    Must函数
      可以包裹一个函数，返回到一个模版的指针和一个错误，如果错误不为nil，那么就panic
Action 
  Action就是Go模版中嵌入的命令，位于两组花括号之间{{ xxx }}
  . 就是一个Action，而且是最重要的一个，他代表了传入模版的数据
  Action主要分为五类：
    条件类
     {{ if }}
     {{ else }}
     {{ end }}
    迭代/遍历类
    {{ range }}
    {{ end }}
    设置类
    {{ with }}
    包含类
      include 包含action形式{{ template "name" . }}
    定义类
    {{ define }}
参数、变量、管道
函数 
模版组合
  执行模版
     Execute：参数是ResponseWriter、数据，单模版很适用，模版集只用第一个模版
     ExecuteTemplate：
       参数是ResponseWrite、模版名、数据
       模版集适用


【路由】
  Controller的角色
    main():设置类工作
    controller：
      静态资源
      把不同请求送到不同controller进行处理
路由参数
  静态路由： 一个路径对于一个页面
   /about
   /home 
  带参数的路由：根据路由参数，创建出一组不同页面
    /companies/123 
    /companies/microsoft
第三方路由器
  gorilla/mux：灵活性高，功能强大，性能差点
  httprouter：注重性能、功能简单
  编写自己的路由规则

  JSON:
    JSON 对于 go Struct 
    类型映射
    Go bool： JSON boolean
    Go float64：JSON 数值 
    Go string：JSON strings 
    Go nil：JSON null
  对于未知结构的JSON
  map[string]interface{}可以存储任意JSON对象
  []interface{}可以存储任意的JSON数组
  读取JSON
    需要一个解码器：dec := json.NewDecoder(r.Body) 参数需要实现Reader接口
    在解码器上进行解码：dec.Decode(&query)
  写入JSON
    需要一个编码器： enc := json.NewEncoder(w) 需要实现Writer接口
    编码： enc.Encode(results)
Marshal和Unmarshal
 Marshal（编码）：把go struct转化为json格式数据
    MarshalIndent 增加缩进
 Unmarshal解码：把json转化为go struct 

 两种方式区别：
   针对string 和bytes： 
      Marshal=>string
      Unmarshal <= string 
  针对stream：
     Encode => Stream 把数据写入io.Writer 
     Decode <= Stream 从io.Reader读取数据


【中间件】

【请求上下文】Request Context 
  Context()
  WitchContext()