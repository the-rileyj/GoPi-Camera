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
	latest, photoDemo, photoVerbose, raw, photoSignal, photoOutput bool
	timeOut, photoTimeLapse, enablePhotoKeypress                   bool
	//Video Specific
	videoKeypress, penc, videoSignal, timed                bool
	videoDemo, videoInline, videoVerbose                   bool
	enableBitRate, enableVideoTimeout, enableVideoKeypress bool
	//General Floats
	blueAWBG, redAWBG float64
	//General Int32
	brightness, cameraSelection                int32
	colorEffectU, colorEffectY, contrast, ev   int32
	iso, mode, previewOpacity, previewX        int32
	previewY                                   int32
	previewWidth, previewHeight, roiX, roiY    int32
	roiWidth, roiHeight, saturation, sharpness int32
	shutter                                    int32
	//Photo Specific Int32
	jpgQuality, photoDemoLength                 int32
	photoWidth, photoHeight, photoTimeOutLength int32
	quality, timeLapse, timeLength              int32
	//Video Specific Int32
	bitRate, frameRate, h264Prof, intraReferesh int32
	videoDemoLength, quantisation, rotation     int32
	segment, start, timeOn, timeOff, videoWidth int32
	videoHeight, wrap, videoCapTimed            int32
	videoTimeOutLength, videoWaitTimed          int32
	//General Strings
	annotate, awb, dynamicRangeCompression     string
	exposure, fileName, fileType, imageEffects string
	meteringMode, output, savePath             string
	simpleCommand                              string
	//Photo Specific Strings
	photoEncoding, latestFileName, thumbNail string
	exif, photoKeypressMode                  string
	//Video Specific Strings
	videoOutput, videoProfile, videoKeypressMode, initialStatevideoKeypressMode string
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

//Raspistill (photo) Specific Functions
/*
	-
*/
func (c *Camera) PhotoWidth(photoWidth int32) *Camera {
	c.photoWidth = photoWidth
	return c
}

func (c *Camera) PhotoHeight(photoHeight int32) *Camera {
	c.photoHeight = photoHeight
	return c
}

func (c *Camera) Quality(quality int32) *Camera {
	c.quality = quality
	return c
}

func (c *Camera) Raw(raw bool) *Camera {
	c.raw = raw
	return c
}

func (c *Camera) PhotoOutput(photoOutput bool) *Camera {
	c.photoOutput = photoOutput
	return c
}

func (c *Camera) Latest(latest bool) *Camera {
	c.latest = latest
	return c
}

func (c *Camera) LatestFileName(latestFileName string) *Camera {
	c.latestFileName = latestFileName
	return c
}

func (c *Camera) PhotoVerbose(photoVerbose bool) *Camera {
	c.photoVerbose = photoVerbose
	return c
}

func (c *Camera) EnableTimeout(timeOut bool) *Camera {
	c.timeOut = timeOut
	return c
}

func (c *Camera) Timeout(photoTimeOutLength int32) *Camera {
	c.photoTimeOutLength = photoTimeOutLength
	return c
}

func (c *Camera) Timelapse(timeLapse int32) *Camera {
	c.timeLapse = timeLapse
	return c
}

func (c *Camera) PhotoTimeLapse(photoTimeLapse bool) *Camera {
	c.photoTimeLapse = photoTimeLapse
	return c
}

func (c *Camera) ThumbNail(thumbNail string) *Camera {
	c.thumbNail = thumbNail
	return c
}

func (c *Camera) PhotoDemo(photoDemo bool) *Camera {
	c.photoDemo = photoDemo
	return c
}

func (c *Camera) PhotoDemoLength(photoDemoLength int32) *Camera {
	c.photoDemoLength = photoDemoLength
	return c
}

func (c *Camera) PhotoEncoding(photoEncoding string) *Camera {
	c.photoEncoding = photoEncoding
	return c
}

func (c *Camera) Exif(exif string) *Camera {
	c.exif = exif
	return c
}

func (c *Camera) FullScreenPreview(fullScreenPreview bool) *Camera {
	c.fullScreenPreview = fullScreenPreview
	return c
}

func (c *Camera) EnablePhotoKeypressMode(enableVideoKeypress bool) *Camera {
	c.enableVideoKeypress = enableVideoKeypress
	return c
}

func (c *Camera) PhotoKeypressMode(photoKeypressMode string) *Camera {
	c.photoKeypressMode = photoKeypressMode
	return c
}

func (c *Camera) PhotoSignal(photoSignal bool) *Camera {
	c.photoSignal = photoSignal
	return c
}

//Raspivid (video) Specific Functions
/*

 */
func (c *Camera) VideoWidth(videoWidth int32) *Camera {
	c.videoWidth = videoWidth
	return c
}

func (c *Camera) VideoHeight(videoHeight int32) *Camera {
	c.videoHeight = videoHeight
	return c
}

func (c *Camera) EnableBitRate(enableBitRate bool) *Camera {
	c.enableBitRate = enableBitRate
	return c
}

func (c *Camera) BitRate(bitRate int32) *Camera {
	c.bitRate = bitRate
	return c
}

func (c *Camera) VideoOutput(videoOutput string) *Camera {
	c.videoOutput = videoOutput
	return c
}

func (c *Camera) VideoVerbose(videoVerbose bool) *Camera {
	c.videoVerbose = videoVerbose
	return c
}

func (c *Camera) EnableVideoTimeout(enableVideoTimeout bool) *Camera {
	c.enableVideoTimeout = enableVideoTimeout
	return c
}

func (c *Camera) VideoTimeOutLength(videoTimeOutLength int32) *Camera {
	c.videoTimeOutLength = videoTimeOutLength
	return c
}

func (c *Camera) VideoDemo(videoDemo bool) *Camera {
	c.videoDemo = videoDemo
	return c
}

func (c *Camera) VideoDemoLength(videoDemoLength int32) *Camera {
	c.videoDemoLength = videoDemoLength
	return c
}

func (c *Camera) FrameRate(frameRate int32) *Camera {
	c.frameRate = frameRate
	return c
}

func (c *Camera) Penc(penc bool) *Camera {
	c.penc = penc
	return c
}

func (c *Camera) IntraReferesh(intraReferesh int32) *Camera {
	c.intraReferesh = intraReferesh
	return c
}

func (c *Camera) Quantisation(quantisation int32) *Camera {
	c.quantisation = quantisation
	return c
}

func (c *Camera) VideoProfile(videoProfile string) *Camera {
	c.videoProfile = videoProfile
	return c
}

func (c *Camera) VideoInline(videoInline bool) *Camera {
	c.videoInline = videoInline
	return c
}

func (c *Camera) videoTimed(capTime, waitTime int32) *Camera {
	c.videoCapTimed = capTime
	c.videoWaitTimed = waitTime
	return c
}

func (c *Camera) VideoCapTimed(capTime int32) *Camera {
	c.videoCapTimed = capTime
	return c
}

func (c *Camera) VideoKeypressMode(videoKeypressMode string) *Camera {
	c.videoKeypressMode = videoKeypressMode
	return c
}

func (c *Camera) VideoSignal(videoSignal bool) *Camera {
	c.videoSignal = videoSignal
	return c
}

func (c *Camera) Segment(segment int32) *Camera {
	c.segment = segment
	return c
}

func (c *Camera) Wrap(wrap int32) *Camera {
	c.wrap = wrap
	return c
}

func (c *Camera) Start(start int32) *Camera {
	c.start = start
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
	}3
	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	cmd.Wait()
	return fullPath, nil
}
