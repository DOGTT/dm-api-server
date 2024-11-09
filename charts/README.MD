# Charts

## Creat charts repo
The chart please create repo to http://{{RepoBase}}/deploy/charts/xxxxx


## Submodule

### Add Charts

add charts repo by submodule
```shell
cd charts

git submodule add http://{{RepoBase}}/deploy/charts/xxxxx
```

update url for `.gitmodules`
```shell

[submodule "charts/xxxxx"]
	path = charts/xxxxx
	url = ../../deploy/charts/xxxxx.git
```

### Update Charts
update charts

```shell
cd charts/xxxxx

git checkout -b feat

# update charts

git add .

git commit -m "update charts"

git push origin feat
```