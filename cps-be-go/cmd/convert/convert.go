// convert.go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ForgeArtifact struct {
	ABI      json.RawMessage `json:"abi"`
	Bytecode struct {
		Object string `json:"object"`
	} `json:"bytecode"`
}

type CombinedJSON struct {
	Contracts map[string]struct {
		ABI json.RawMessage `json:"abi"`
		Bin string          `json:"bin"`
	} `json:"contracts"`
}

/**
 * 这个工具将 Foundry 的 artifact 转换为 combined-json 格式，方便 Go 代码使用。
 * 运行前请确保已经编译了合约，并且路径正确。
 * 输出文件将保存在 ../contracts/ 目录下，命名为 <ContractName>-combined.json。
 * 用法: `go run utils/convert.go MyToken` 或在 utils 目录下 `go run convert.go MyToken`
 */
func main() {
	// contractName 从命令行第一个参数获取：
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "缺少参数,命令行用法: go run convert.go <ContractName>")
		os.Exit(2)
	}
	contractName := os.Args[1]
	data, err := os.ReadFile("../../../cps-contracts/out/" + contractName + ".sol/" + contractName + ".json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to read artifact:", err)
		os.Exit(1)
	}
	var artifact ForgeArtifact
	if err := json.Unmarshal(data, &artifact); err != nil {
		fmt.Fprintln(os.Stderr, "failed to unmarshal artifact:", err)
		os.Exit(1)
	}

	// 转换为 combined-json 格式
	combined := CombinedJSON{
		Contracts: map[string]struct {
			ABI json.RawMessage `json:"abi"`
			Bin string          `json:"bin"`
		}{
			contractName + ".sol:" + contractName: {
				ABI: artifact.ABI,
				Bin: artifact.Bytecode.Object,
			},
		},
	}

	output, err := json.MarshalIndent(combined, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to marshal combined json:", err)
		os.Exit(1)
	}
	// 写出 combined json 文件
	if err := os.WriteFile("../../contracts/"+contractName+"-combined.json", output, 0644); err != nil {
		fmt.Fprintln(os.Stderr, "failed to write output file:", err)
		os.Exit(1)
	}
	// 另外写出单独的 ABI 文件，内容为 artifact.ABI（原始 JSON bytes）
	if err := os.WriteFile("../../contracts/"+contractName+".abi", artifact.ABI, 0644); err != nil {
		fmt.Fprintln(os.Stderr, "failed to write abi file:", err)
		os.Exit(1)
	}
	fmt.Println("wrote sucess!! path: ../../contracts/" + contractName + "-combined.json and ../../contracts/" + contractName + ".abi")
}
