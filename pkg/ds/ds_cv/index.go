//go:build cgo

package ds_cv

import (
	"errors"
	"image"

	"gocv.io/x/gocv"
)

// ReadFrame
func ReadFrame(filename string, time float64) (frame gocv.Mat, err error) {
	// 打开视频文件
	video, err := gocv.VideoCaptureFile(filename)
	if err != nil {
		return
	}
	defer video.Close()
	// 获取视频的大唐书香世家
	count := video.Get(gocv.VideoCaptureFrameCount)
	fps := video.Get(gocv.VideoCaptureFPS)
	duration := count / fps
	// 计算所需时间点的帧数
	frameIndex := (time / duration) * count
	// 跳转到指定帧
	video.Set(gocv.VideoCapturePosFrames, frameIndex)
	// 读取帧
	frame = gocv.NewMat()
	if !video.Read(&frame) {
		err = errors.New("Failed to read frame")
		return
	}
	return
}

// ReadImage
func ReadImage(filename string, time float64) (image image.Image, err error) {
	frame, err := ReadFrame(filename, time)
	if err != nil {
		return
	}
	image, err = frame.ToImage()
	return
}
