# wc-clone 📝

A lightweight and fast `wc` (word count) clone written in Go — supporting line, word, and byte counting from files or standard input, just like the Unix `wc` command.

## 🚀 Features

- ✅ Count **lines**, **words**, and **bytes**
- ✅ Works with **multiple files**
- ✅ Supports **stdin** (piped input)
- ✅ Compatible with common `wc` flags: `-l`, `-w`, `-c`
- ✅ Outputs aligned just like GNU `wc`
- ✅ Clean and idiomatic Go structure

## 📦 Installation

```bash
git clone git@github.com:PCoelho07/wc-clone.git
cd wc-clone
go build -o ccwc cmd/main.go
