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
