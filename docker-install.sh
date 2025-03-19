#!/bin/bash

# 检查是否为 macOS
if [[ "$(uname)" != "Darwin" ]]; then
  echo "This script is only for macOS."
  exit 1
fi

# 检查是否有管理员权限
if [[ "$(id -u)" != "0" ]]; then
  echo "This script requires administrator privileges. Please run with sudo."
  exit 1
fi

# 卸载 Docker
echo "Uninstalling Docker..."

# 停止 Docker 服务
osascript -e 'quit app "Docker"'

# 删除 Docker 应用程序
rm -rf /Applications/Docker.app

# 删除 Docker 相关配置文件和缓存
rm -rf ~/Library/Containers/com.docker.docker
rm -rf ~/Library/Application\ Support/Docker
rm -rf ~/Library/Group\ Containers/group.com.docker
rm -rf ~/Library/Preferences/com.docker.docker.plist
rm -rf ~/Library/Saved\ Application\ State/com.docker.docker.savedState
rm -rf ~/Library/Caches/com.docker.docker
rm -rf ~/Library/Logs/Docker

echo "Docker has been uninstalled."

# 重新安装 Docker
echo "Downloading and installing Docker..."

# # 下载最新版本的 Docker for macOS
# DOCKER_DMG_URL="https://desktop.docker.com/mac/main/amd64/Docker.dmg"
# DOCKER_DMG="/tmp/Docker.dmg"

# curl -L -o "$DOCKER_DMG" "$DOCKER_DMG_URL"

# # 挂载 DMG 文件
# MOUNT_POINT=$(hdiutil attach "$DOCKER_DMG" | grep -E '/Volumes/Docker' | awk '{print $3}')

# # 安装 Docker
# sudo cp -rf "$MOUNT_POINT/Docker.app" /Applications/

# # 卸载 DMG 文件
# hdiutil detach "$MOUNT_POINT"

# # 删除下载的 DMG 文件
# rm "$DOCKER_DMG"

# echo "Docker has been installed. Please open Docker.app to complete the setup."