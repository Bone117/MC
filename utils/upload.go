package utils

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"os/exec"
	"path"
	"server/global"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
)

func UploadFile(file *multipart.FileHeader, userid uint, jieCi uint) (string, error) {
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = MD5V([]byte(name))
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	// 创建路径

	storePath := path.Join(global.CONFIG.Local.StorePath, strconv.Itoa(int(jieCi)), strconv.Itoa(int(userid)))
	mkdirErr := os.MkdirAll(storePath, os.ModePerm)
	if mkdirErr != nil {
		global.LOG.Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
		return "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	p := storePath + "/" + filename
	filepath := storePath + "/" + filename

	f, openError := file.Open() // 读取文件
	if openError != nil {
		global.LOG.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(p)
	if createErr != nil {
		global.LOG.Error("function os.Create() Filed", zap.Any("err", createErr.Error()))
		return "", errors.New("function os.Create() Filed, err:" + createErr.Error())

	}
	defer out.Close() // 创建文件 defer 关闭

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		global.LOG.Error("function io.Copy() Filed", zap.Any("err", copyErr.Error()))
		return "", errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return filepath, nil
}

func DeleteFile(key string) error {
	//p := global.CONFIG.Local.StorePath + "/" + key
	//uploads/file/2/31/63f082743060b88e183d8da6e73d56c9_20230317010046.mp4
	//uploads/file/uploads/file/2/31/63f082743060b88e183d8da6e73d56c9_20230317010046.mp4
	if strings.Contains(key, global.CONFIG.Local.StorePath) {
		if err := os.Remove(key); err != nil {
			return errors.New("本地文件删除失败, err:" + err.Error())
		}
	}
	return nil
}

// GenerateThumbnail 生成缩略图
func GenerateThumbnail(videoPath string) (string, error) {
	thumbnailPath := fmt.Sprintf("%s_thumbnail.jpg", videoPath)

	cmd := exec.Command("ffmpeg", "-i", videoPath, "-ss", "00:00:01", "-vframes", "1", thumbnailPath)
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return thumbnailPath, nil
}
