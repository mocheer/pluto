package img

import "image"

// Thumbnail 缩略图
// 不同于Resize，当图片比缩略图小，不会缩放
func (p *Img) Thumbnail(width, height int) *Img {
	return &Img{Image: Thumbnail(p.Image, width, height)}
}

// Thumbnail
func Thumbnail(target image.Image, maxWidth, maxHeight int) image.Image {
	origBounds := target.Bounds()
	origWidth := origBounds.Dx()
	origHeight := origBounds.Dy()
	newWidth, newHeight := origWidth, origHeight

	// Return original image if it have same or smaller size as constraints
	if maxWidth >= origWidth && maxHeight >= origHeight {
		return target
	}

	// Preserve aspect ratio
	if origWidth > maxWidth {
		newHeight = origHeight * maxWidth / origWidth
		if newHeight < 1 {
			newHeight = 1
		}
		newWidth = maxWidth
	}

	if newHeight > maxHeight {
		newWidth = newWidth * maxHeight / newHeight
		if newWidth < 1 {
			newWidth = 1
		}
		newHeight = maxHeight
	}
	return Resize(target, newWidth, newHeight)
}
