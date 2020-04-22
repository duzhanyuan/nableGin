# nableGin Install

1、git clone https://github.com/kokobing/nableGin.git

2、cd nableGin & config config.yml for mysql and redis

2、god mod init nablegin

3、go env -w GO111MODULE=on

4、go env -w GOPROXY=https://goproxy.cn,direct

5、go run main.go

6、http://localhost:8088/admin/login/    admin   123456


First bulid you need open nableGin/app/route.go 18-25 for casbin permit :

	e.LoadPolicy()
	e.AddPolicy("anonymous", "/admin", "GET")
	e.AddPolicy("anonymous", "/admin/login/", "GET")
	e.AddPolicy("anonymous", "/admin/login/", "POST")
	e.AddPolicy("admin", "/*", "*")
	if err := e.SavePolicy(); err != nil {
		panic(err)
	}



# Screenshot

![](https://github.com/Kokolpb/PHP-WEBSITE-TEMPLATE-L1/blob/master/Screenshot_1.png)  
![](https://github.com/Kokolpb/PHP-WEBSITE-TEMPLATE-L1/blob/master/Screenshot_2.png)  
![](https://github.com/Kokolpb/PHP-WEBSITE-TEMPLATE-L1/blob/master/Screenshot_3.png)  
![](https://github.com/Kokolpb/PHP-WEBSITE-TEMPLATE-L1/blob/master/Screenshot_4.png)  
![](https://github.com/Kokolpb/PHP-WEBSITE-TEMPLATE-L1/blob/master/Screenshot_5.png)  
![](https://github.com/Kokolpb/PHP-WEBSITE-TEMPLATE-L1/blob/master/Screenshot_6.png)  
![](https://github.com/Kokolpb/PHP-WEBSITE-TEMPLATE-L1/blob/master/Screenshot_7.png)  
![](https://github.com/Kokolpb/PHP-WEBSITE-TEMPLATE-L1/blob/master/Screenshot_8.png)  
![](https://github.com/Kokolpb/PHP-WEBSITE-TEMPLATE-L1/blob/master/Screenshot_9.png)  
![](https://github.com/Kokolpb/PHP-WEBSITE-TEMPLATE-L1/blob/master/Screenshot_10.png)  
![](https://github.com/Kokolpb/PHP-WEBSITE-TEMPLATE-L1/blob/master/Screenshot_11.png)  
![](https://github.com/Kokolpb/PHP-WEBSITE-TEMPLATE-L1/blob/master/Screenshot_12.png)  
![](https://github.com/Kokolpb/PHP-WEBSITE-TEMPLATE-L1/blob/master/Screenshot_13.png)  
![](https://github.com/Kokolpb/PHP-WEBSITE-TEMPLATE-L1/blob/master/Screenshot_14.png)  
