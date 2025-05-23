# Inference Project Template

## TODO
- [x] HTTP(gin) and GRPC server
- [x] metric self define
- [x] metric grpc and gin 
- [x] lint 
- [x] build / dockerfile
- [x] auto install dep utils
- [ ] charts
- [ ] ut/cov
- [ ] mock
- [ ] git ci/cd
- [ ] call channel to units

# Feature

- [ ] 基础业务实现
- [ ] 回复通知
- [ ] 阅读位置记录
- [ ] 消息系统

## 规划

- 宠物关联机制
- 组队系统 
- 寻宠系统


频道是抽象灵活的概念，可以有不同的形式类型
帖子也是抽象概念，有不同的类型， 如足迹贴，互动贴，普通贴，频道下可以有不同类型的帖子

可以用于：
- 组合地点形成足迹频道，归类足迹和发帖，用于大家分享查询一个地点的动态，足迹频道在地图上发现；
- 组合功能形成功能频道，用于大家在这个频道体验一下有趣的东西，如ai头像生成；用户发送图片，服务器给予回复
- 组合主题形成交流频道，用于大家在这个频道分享，交流，如生日祝福，寻宠，游玩地点分享交流；

大的类型是足迹频道
足迹频道有小类型，说明这个足迹的类型？
类型属于足迹帖子
频道用于聚合足迹，没有足迹就没有足迹频道

发表是发表足迹，当下位置当前发表或者选择位置和时间发表。


足迹频道的生命周期
地点坐标是绝对核心，创建者并非管理员，创建者只是在这里发表了一个帖子，

1 用户创建频道，并发布第一个帖子
2 大家加入并浏览信息，可以发表攻略，动态
3 地点发生变化，不再活跃自动褪色
4 用户无法解散频道，频道无法再更新
5 管理权限基于实际打卡次数

关于类型：



## Dev Env
```
# setup git
vim ~/.netrc

# install protoc
make setup

# setup go utils
make init

```

## Build
```
make build

make package
```
