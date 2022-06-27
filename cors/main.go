package main

import (
	"fmt"
	"net/http"
)

func main() {
	go server(":8083")
	go server(":8081")
	server(":8082")
}

/*
<?php
header('Content-type: application/json');
//获取回调函数名
$jsoncallback = htmlspecialchars($_REQUEST ['jsoncallback']);
//json数据
$json_data = '["customername1","customername2"]';
//输出jsonp格式的数据
echo $jsoncallback . "(" . $json_data . ")";
?>
*/

//https://www.w3cschool.cn/ajax/ajax-k68g2or5.html
func server(port string) {
	handler := http.NewServeMux()
	handler.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch port {
		case ":8082":
			/*
				client := &http.Client{}
				  //生成要访问的url
				  url := "http://somesite/somepath/"
				  //提交请求
				  reqest, err := http.NewRequest("GET", url, nil)

				  //增加header选项
				  reqest.Header.Add("Cookie", "xxxxxx")
			*/
			//    使用通配符 * ，表示当前服务端 返回的信息允许所有源访问,也可指定可信任的域名来接收响应信息
			//https://blog.csdn.net/weixin_39634022/article/details/111613779
			w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8083")
			//w.Header().Set("Access-control-Allow-Origin", "*")
			w.Header().Set("access-control-allow-methods", "PUT,PATCH")
			//    响应头设置为ajax提交
			w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			//    允许携带 用户认证凭据（也就是允许客户端发送的请求携带Cookie）
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			fmt.Println(r.Method)
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
		fmt.Fprint(w, `<!DOCTYPE html>
		<html>
		<head>
		<meta charset="utf-8">
		<script>
		function loadXMLDoc(v)
		{
			var xmlhttp;
			if (window.XMLHttpRequest)
			{
				//  IE7+, Firefox, Chrome, Opera, Safari 浏览器执行代码
				xmlhttp=new XMLHttpRequest();
			}
			else
			{
				// IE6, IE5 浏览器执行代码
				xmlhttp=new ActiveXObject("Microsoft.XMLHTTP");
			}
			xmlhttp.onreadystatechange=function()
			{
				if (xmlhttp.readyState==4 && xmlhttp.status==200)
				{
					document.getElementById("myDiv").innerHTML=xmlhttp.responseText;
				}
			}
			switch (v){
			case 1:
				xmlhttp.open("GET","HTTP://127.0.0.1:8082",true);
				xmlhttp.send();
				break;
				case 2:	
				xmlhttp.open("PUT","HTTP://127.0.0.1:8082",true);
				xmlhttp.send();
				break;
				case 3:
					xmlhttp.open("GET","HTTP://127.0.0.1:8081",true);
					xmlhttp.send();
					break;	
			}
		}
		</script>
		</head>
		<body>
		
		<div id="myDiv"><h2>测试请求</h2></div>
		<button type="button" onclick="loadXMLDoc(1)">修改内容1</button>
		<button type="button" onclick="loadXMLDoc(2)">修改内容2</button>
		<button type="button" onclick="loadXMLDoc(3)">修改内容3</button>
		
		</body>
		</html>`,
		/*`
				<!DOCTYPE html>
		<html>
		<head>
		<meta charset="utf-8">
		<title>JSONP 实例</title>
		</head>
		<body>
		<div id="divCustomers"></div>
		<script type="text/javascript">
		function callbackFunction(result, methodName)
		{
		    var html = '<ul>';
		    for(var i = 0; i < result.length; i++)
		    {
		        html += '<li>' + result[i] + '</li>';
		    }
		    html += '</ul>';
		    document.getElementById('divCustomers').innerHTML = html;
		}
		</script>
		<script type="text/javascript" src="https://www.runoob.com/try/ajax/jsonp.php?jsoncallback=callbackFunction"></script>
		</body>
		</html>
		`*/
		/*	`
			<!DOCTYPE html>
			<html>
			<head>
				<meta charset="utf-8">
				<title>JSONP 实例</title>
				<script src="https://cdn.static.runoob.com/libs/jquery/1.8.3/jquery.js"></script>
			</head>
			<body>
			<div id="divCustomers"></div>
			<script>
			$.getJSON("https://www.runoob.com/try/ajax/jsonp.php?jsoncallback=?", function(data) {

				var html = '<ul>';
				for(var i = 0; i < data.length; i++)
				{
					html += '<li>' + data[i] + '</li>';
				}
				html += '</ul>';

				$('#divCustomers').html(html);
			});
			</script>
			</body>
			</html>
		*/
		)
	}))
	http.ListenAndServe(port, handler)
}

//https://www.runoob.com/ajax/ajax-json.html
