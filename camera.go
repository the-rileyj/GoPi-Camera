package camera

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"time"
)

const (
	PHOTO      = "raspistill"
	VIDEO      = "raspivid"
	HFLIP      = "-hf"
	VFLIP      = "-vf"
	OUTFLAG    = "-o"
	TIME_STAMP = "2006-01-02_15:04::05"
)

type Camera struct {
	//General Booleans
	AWBG, fullScreenPreview, horizontalFlip bool
	preview, verticalFlip                   bool
	videoStablization, roi, simpleCapture   bool
	stats                                   bool
	//Photo Specific
	latest, photoDemo, photoVerbose, raw bool
	timeOut                              bool
	//Video Specific
	keypress, penc, signal, timed, videoDemo, videoInline bool
	videoVerbose                                          bool
	//General Floats
	blueAWBG, redAWBG float64
	//General Int32
	brightness, cameraSelection                   int32
	colorEffectU, colorEffectY, contrast, ev      int32
	iso, mode, previewOpacity, previewX, previewY int32
	previewWidth, previewHeight, roiX, roiY       int32
	roiWidth, roiHeight, saturation, sharpness    int32
	shutter                                       int32
	//Photo Specific Int32
	photoWidth, photoHeight, jpgQuality int32
	timeLength, timeOutLength           int32
	//Video Specific Int32
	bitRate, frameRate, videoWidth         int32
	videoHeight, h264Prof, intraReferesh   int32
	quantisation, rotation, segment, start int32
	timeOn, timeOff, wrap                  int32
	//General Strings
	annotate, awb, dynamicRangeCompression     string
	exposure, fileName, fileType, imageEffects string
	meteringMode, output, savePath             string
	simpleCommand                              string
	//Photo Specific Strings
	photoEncoding, latestFileName, thumb string
	exif                                 string
	//Video Specific Strings
	videoProfile, initialState string
}

func New(path, name, fType string) *Camera {
	if name == "" {
		name = time.Now().Format(TIME_STAMP)
	}
	if fType == "" {
		fType = ".jpg"
	}
	return &Camera{horizontalFlip: false, verticalFlip: false, fileName: name, fileType: fType, savePath: path}
}

func (c *Camera) Preview(b bool) *Camera {
	c.preview = b
	return c
}

func (c *Camera) PreviewSize(previewX, previewY, previewWidth, previewHeight int32) *Camera {
	c.previewX = previewX
	c.previewY = previewY
	c.previewWidth = previewWidth
	c.previewHeight = previewHeight
	return c
}

func (c *Camera) PreviewOpacity(opacity int32) *Camera {
	c.previewOpacity = opacity
	return c
}

func (c *Camera) Sharpness(sharpness int32) *Camera {
	c.sharpness = sharpness
	return c
}

func (c *Camera) Contrast(contrast int32) *Camera {
	c.contrast = contrast
	return c
}

func (c *Camera) Brightness(brightness int32) *Camera {
	c.brightness = brightness
	return c
}

func (c *Camera) Saturation(saturation int32) *Camera {
	c.saturation = saturation
	return c
}

func (c *Camera) ISO(iso int32) *Camera {
	c.iso = iso
	return c
}

func (c *Camera) VideoStablization(videoStablization bool) *Camera {
	c.videoStablization = videoStablization
	return c
}

func (c *Camera) EV(ev int32) *Camera {
	c.ev = ev
	return c
}

func (c *Camera) Exposure(exposure string) *Camera {
	c.exposure = exposure
	return c
}

func (c *Camera) AWB(awb string) *Camera {
	c.awb = awb
	return c
}

func (c *Camera) ImageEffects(imageEffect string) *Camera {
	c.imageEffects = imageEffect
	return c
}

func (c *Camera) ColorEffects(colorEffectU, colorEffectY int32) *Camera {
	c.colorEffectU = colorEffectU
	c.colorEffectY = colorEffectY
	return c
}

func (c *Camera) ColorEffectU(colorEffectU int32) *Camera {
	c.colorEffectU = colorEffectU
	return c
}

func (c *Camera) ColorEffectY(colorEffectY int32) *Camera {
	c.colorEffectY = colorEffectY
	return c
}

func (c *Camera) Rotation(rotation int32) *Camera {
	c.rotation = rotation
	return c
}

func (c *Camera) HorizonalFlip(horizontalFlip bool) *Camera {
	c.horizontalFlip = horizontalFlip
	return c
}

func (c *Camera) VerticalFlip(verticalFlip bool) *Camera {
	c.verticalFlip = verticalFlip
	return c
}

func (c *Camera) ROI(roi bool) *Camera {
	c.roi = roi
	return c
}

func (c *Camera) ROICoordinates(roiX, roiY, roiWidth, roiHeight int32) *Camera {
	c.roiX, c.roiY, c.roiWidth, c.roiHeight = roiX, roiY, roiWidth, roiHeight
	return c
}

func (c *Camera) Shutter(shutter int32) *Camera {
	c.shutter = shutter
	return c
}

func (c *Camera) DynamicRangeCompression(dynamicRangeCompression string) *Camera {
	c.dynamicRangeCompression = dynamicRangeCompression
	return c
}

func (c *Camera) Stats(stats bool) *Camera {
	c.stats = stats
	return c
}

func (c *Camera) AWBGains(blueAWBG, redAWBG float64) *Camera {
	c.blueAWBG, c.redAWBG = blueAWBG, redAWBG
	return c
}

func (c *Camera) BlueAWBG(blueAWBG float64) *Camera {
	c.blueAWBG = blueAWBG
	return c
}

func (c *Camera) RedAWBG(redAWBG float64) *Camera {
	c.redAWBG = redAWBG
	return c
}

func (c *Camera) Mode(mode int32) *Camera {
	c.mode = mode
	return c
}

func (c *Camera) CameraSelection(cameraSelection int32) *Camera {
	c.cameraSelection = cameraSelection
	return c
}

func (c *Camera) Annotate(annotate string) *Camera {
	c.annotate = annotate
	return c
}

func (c *Camera) Output(output string) *Camera {
	c.output = output
	return c
}

func (c *Camera) Capture() (string, error) {
	args := make([]string, 0)
	args = append(args, OUTFLAG)
	fullPath := c.fileName
	if c.savePath != "" {
		fullPath = filepath.Join(c.savePath, c.fileName)
	}
	args = append(args, fullPath)
	if c.horizontalFlip {
		args = append(args, HFLIP)
	}
	if c.verticalFlip {
		args = append(args, VFLIP)
	}
	cmd := exec.Command(PHOTO, args...)
	_, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	cmd.Wait()
	return fullPath, nil
}
