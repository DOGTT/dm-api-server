#!/bin/bash

# 定义密钥文件名
PRIVATE_KEY_FILE="configs/private_key.pem"
PUBLIC_KEY_FILE="configs/public_key.pem"

# 检查是否安装了 OpenSSL
check_openssl() {
    if ! command -v openssl &>/dev/null; then
        echo "OpenSSL is not installed. Installing OpenSSL..."
        install_openssl
    else
        echo "OpenSSL is already installed."
    fi
}

# 安装 OpenSSL（适用于 macOS）
install_openssl() {
    if command -v brew &>/dev/null; then
        brew install openssl
    else
        echo "Homebrew is not installed. Please install Homebrew first."
        exit 1
    fi
}

# 生成 RSA 公私密钥
generate_keys() {
    if [[ -f $PRIVATE_KEY_FILE ]] && [[ -f $PUBLIC_KEY_FILE ]]; then
        echo "Keys already exist:"
        echo " - Private Key: $PRIVATE_KEY_FILE"
        echo " - Public Key: $PUBLIC_KEY_FILE"
        return
    fi

    echo "Generating RSA keys..."
    openssl genrsa -out "$PRIVATE_KEY_FILE" 2048
    openssl rsa -in "$PRIVATE_KEY_FILE" -pubout -out "$PUBLIC_KEY_FILE"

    if [[ -f $PRIVATE_KEY_FILE ]] && [[ -f $PUBLIC_KEY_FILE ]]; then
        echo "Keys generated successfully:"
        echo " - Private Key: $PRIVATE_KEY_FILE"
        echo " - Public Key: $PUBLIC_KEY_FILE"
    else
        echo "Failed to generate keys."
        exit 1
    fi
}

# 主程序入口
main() {
    check_openssl
    generate_keys
}

# 执行主程序
main