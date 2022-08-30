package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

// Структура блока блокчейна
type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

// Переменная, содержащая блокчейн
var Blockchain []Block

// Берет данные Block(прошлого блока) и создает для них новый хэш SHA256
func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + strconv.Itoa(block.BPM) + block.PrevHash // Объединение в одну строку данных из прошлого блока
	h := sha256.New()                                                                                // Объявление нового хэша
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// Создание блока
func generateBlock(oldBlock Block, BPM int) (Block, error) {
	var newBlock Block
	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.Hash = calculateHash(newBlock)
	newBlock.PrevHash = oldBlock.Hash

	return newBlock, nil
}

// Проверка блока на валидность
func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index { // Проверка индекса
		return false
	}
	if oldBlock.Hash != newBlock.PrevHash { // Проверка хэша предыдущего блока в новом блоке
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash { // Проверка хэша
		return false
	}
	return true
}

func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}
