# Lottery-Period

使用golang开发的福彩、体彩的数字彩种期数算法及API接口


## 特性

* 可输出API，供第三方程序调用；
* 也可修改代码，提供GRPC接口调用；
* 增加彩种或新的一年时，只需要修改项目中yaml内的配置文件；
* docker image(TODO)；

## 安装

```sh
go get -u github.com/zituocn/lottery-period
```

```sh
cd lottery-period
go run .
```

或

```sh
cd lottery-period
go build
./lottery-period # windows下执行 lottery-period.exe
```



## API路由

> 接口路由说明

```go
func APIRouter(r *gow.Engine) {
	r.NoRoute(handler.ErrorHandler)
	v1 := r.Group("/api")
	{
		v1.GET("/time", handler.GetOneByTime)     //根据时间取配置
		v1.GET("/period", handler.GetOneByPeriod) //根据期数取配置
		v1.GET("/today", handler.Today)           //今天开奖信息
		v1.GET("/game", handler.Game)             //彩种配置输出
		v1.GET("/year", handler.Year)             //年份配置输出
		v1.GET("/game/yaml", handler.GameYaml)    //彩种配置yaml输出
		v1.GET("/year/yaml", handler.YearYaml)    //年份配置yaml输出
	}
}
```

## 接口说明

分别就提供的接口做出的说明

### 根据时间取期数配置

*接口*

```sh
GET /api/time
```

*参数*

```sh
gamecode  int 彩种id
ts        int 时间戳，不传为当前时间
typeid    int -1=上一期 0=当前期 1=下一期
```

*带参数*


```sh
GET /api/time?gamecode=1&ts=1652001618&typeid=0
```

*返回值*

```json
{
  "code": 0,
  "msg": "success",
  "time": 1657100708,
  "data": {
    "period": 2022118,
    "bt": "2022-05-08 21:15:00",
    "et": "2022-05-08 21:15:00",
    "at": "2022-05-08 21:15:00"
  }
}
```

### 根据期数取期数配置

*接口*

```sh
GET /api/period
```

*参数*

```sh
gamecode  int 彩种id
period    int 期数
```

*带参数*

```sh
GET /api/period?gamecode=1&period=2022150
```

*返回值*

```json
{
  "code": 0,
  "msg": "success",
  "time": 1657100824,
  "data": {
    "period": 2022150,
    "bt": "2022-06-09 21:15:00",
    "et": "2022-06-09 21:15:00",
    "at": "2022-06-09 21:15:00"
  }
}
```

### 查询所有彩种配置

可查询 gamecode

*接口*

```sh
GET  /api/game
```

*返回值*

```json
{
  "code": 0,
  "msg": "success",
  "time": 1657100429,
  "data": [
    {
      "gamecode": 1,
      "oid": 1,
      "typeid": 1,
      "name": "福彩3D",
      "sname": "福彩3D",
      "ename": "fcsd",
      "opendate": "Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday",
      "opentime": "21:15",
      "state": 0,
      "islocal": 0,
      "province": "",
      "formatqi": "YYYY,%.3d",
      "selist": null
    }...
  ]
}

```