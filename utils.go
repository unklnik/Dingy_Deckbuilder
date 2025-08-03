package main

import (
	"math"
	"math/rand/v2"
	"os"
	"path/filepath"
	"strings"

	z "github.com/gen2brain/raylib-go/raylib"
)

var (
	scrollONOFF bool
	scrollSPD   float32 = 4
)

// MARK: SLICES
func REM[T any](s []T, i int) []T {
	if i < 0 || i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}

// MARK: REC
func moveRECtoV2topleftandRECCOLLIS(r, rC z.Rectangle, newPos z.Vector2) (z.Rectangle, z.Rectangle) {
	diffx := rC.X - r.X
	diffy := rC.Y - r.Y
	r.X = newPos.X
	r.Y = newPos.Y
	rC.X = r.X + diffx
	rC.Y = r.Y + diffy
	return r, rC
}
func moveRECCOLLIStoV2topleftandREC(r, rC z.Rectangle, newPos z.Vector2) (z.Rectangle, z.Rectangle) {
	diffx := rC.X - r.X
	diffy := rC.Y - r.Y
	rC.X = newPos.X
	rC.Y = newPos.Y
	r.X = rC.X - diffx
	r.Y = rC.Y - diffy
	return r, rC
}
func qResizeRecDOUBLE(r z.Rectangle) z.Rectangle {
	return z.NewRectangle(r.X-r.Width/2, r.Y-r.Height/2, r.Width*2, r.Height*2)
}
func qResizeRecTRIPLE(r z.Rectangle) z.Rectangle {
	return z.NewRectangle(r.X-r.Width, r.Y-r.Height, r.Width*3, r.Height*3)
}
func qShadowREC(r z.Rectangle, offset float32) z.Rectangle {
	return z.NewRectangle(r.X-offset, r.Y+offset, r.Width, r.Height)
}
func DREC(r z.Rectangle) z.Rectangle {
	return z.NewRectangle(r.X+r.Width/2, r.Y+r.Height/2, r.Width, r.Height)
}
func qResizeRecNARROWERcnt(r z.Rectangle, offset float32) z.Rectangle {
	return z.NewRectangle(r.X+offset, r.Y, r.Width-offset*2, r.Height)
}
func qResizeRecNARROWERright(r z.Rectangle, offset float32) z.Rectangle {
	return z.NewRectangle(r.X, r.Y, r.Width-offset*2, r.Height)
}
func qResizeRecSHORTERcnt(r z.Rectangle, offset float32) z.Rectangle {
	return z.NewRectangle(r.X, r.Y+offset, r.Width, r.Height-offset*2)
}
func qResizeRecSHORTERtop(r z.Rectangle, offset float32) z.Rectangle {
	return z.NewRectangle(r.X, r.Y+offset*2, r.Width, r.Height-offset*2)
}
func qResizeRecSHORTERbottom(r z.Rectangle, offset float32) z.Rectangle {
	return z.NewRectangle(r.X, r.Y, r.Width, r.Height-offset*2)
}
func qHeightProportional(w, h, newW float32) float32 {
	return (newW * h) / w
}
func qWidthProportional(w, h, newH float32) float32 {
	return (newH * w) / h
}
func qResizeRecHALF(r z.Rectangle) z.Rectangle {
	return z.NewRectangle(r.X+r.Width/4, r.Y+r.Height/4, r.Width/2, r.Height/2)
}
func qResizeRecTHREEQUARTER(r z.Rectangle) z.Rectangle {
	return z.NewRectangle(r.X+r.Width/8, r.Y+r.Height/8, (r.Width/4)*3, (r.Height/4)*3)
}
func qResizeRecProportionalH(r z.Rectangle, w float32) z.Rectangle {
	h := (w / r.Width) * r.Height
	return z.NewRectangle(float32(r.X), float32(r.Y), w, h)
}
func qResizeRecOffsetSMLR(r z.Rectangle, offset float32) z.Rectangle {
	return z.NewRectangle(r.X+offset, r.Y+offset, r.Width-offset*2, r.Height-offset*2)
}
func qResizeRecOffsetLRGR(r z.Rectangle, offset float32) z.Rectangle {
	return z.NewRectangle(r.X-offset, r.Y-offset, r.Width+offset*2, r.Height+offset*2)
}
func qRecPointsREC(r z.Rectangle) []z.Vector2 {
	v1 := z.NewVector2(r.X, r.Y)
	v2 := v1
	v2.X += r.Width
	v3 := v2
	v3.Y += r.Height
	v4 := v1
	v4.Y += r.Height
	return []z.Vector2{v1, v2, v3, v4}
}
func qRecPointsCNT(cnt z.Vector2, w, h float32) []z.Vector2 {
	v1 := z.NewVector2(cnt.X-w/2, cnt.Y-h/2)
	v2 := v1
	v2.X += w
	v3 := v2
	v3.Y += h
	v4 := v1
	v4.Y += h
	return []z.Vector2{v1, v2, v3, v4}
}
func qResizeRecXYZERO(r z.Rectangle, zoom float32) z.Rectangle {
	return z.NewRectangle(r.X, r.Y, r.Width*zoom, r.Height*zoom)
}
func qORIGIN(r z.Rectangle) z.Vector2 {
	return z.NewVector2(r.Width/2, r.Height/2)
}
func qRecCNT(r z.Rectangle) z.Vector2 {
	return z.NewVector2(r.X+r.Width/2, r.Y+r.Height/2)
}
func qRecCNTXONLY(r z.Rectangle) float32 {
	return r.X + r.Width/2
}
func qRecCNTYONLY(r z.Rectangle) float32 {
	return r.Y + r.Height/2
}

// MARK: OBJ
func IM2OBJ(ims []IM) []OBJ {
	var o []OBJ
	o2 := OBJ{}
	for i := range ims {
		o2.im = ims[i]
		o = append(o, o2)
	}
	return o
}
func remOBJoff(o []OBJ) []OBJ {
	var o2 []OBJ
	for i := range o {
		if !o[i].off {
			o2 = append(o2, o[i])
		}
	}
	return o2
}

// MARK: COLLISIONS
func cMSREC(r z.Rectangle) bool {
	return MS.X >= r.X && MS.X <= r.X+r.Width && MS.Y >= r.Y && MS.Y <= r.Y+r.Height
}
func cPointRec(p z.Vector2, r z.Rectangle) bool {
	return p.X >= r.X && p.X <= r.X+r.Width && p.Y >= r.Y && p.Y <= r.Y+r.Height
}
func cRecIMSheetRecDrawCollis(r z.Rectangle, ims []IM) bool {
	collis := false
	for i := range ims {
		if z.CheckCollisionRecs(r, ims[i].rD) {
			collis = true
			break
		}
	}
	return collis
}
func cRECvsRECslice(r z.Rectangle, r2 []z.Rectangle) bool {
	collis := false
	for i := range r2 {
		if z.CheckCollisionRecs(r, r2[i]) {
			collis = true
			break
		}
	}
	return collis
}

// MARK: VECTOR2
func Angl2Points(v1, v2 z.Vector2) float32 {
	deltaX := v2.X - v1.X
	deltaY := v2.Y - v1.Y
	radians := math.Atan2(float64(deltaY), float64(deltaX))
	degrees := float32((radians * 180) / math.Pi)
	for degrees >= 360 {
		degrees -= 360
	}
	for degrees < 0 {
		degrees += 360
	}
	return degrees
}
func PointOnCirc(radius, angl float32, cntr z.Vector2) z.Vector2 {
	angleInRadians := float64(angl * math.Pi / 180.0)
	x := float64(radius)*math.Cos(angleInRadians) + float64(cntr.X)
	y := float64(radius)*math.Sin(angleInRadians) + float64(cntr.Y)
	return z.NewVector2(float32(x), float32(y))
}
func VELXY(p1 z.Vector2, p2 z.Vector2, maxspeed float32) (xvel float32, yvel float32) {
	xvel = p2.X - p1.X
	yvel = p2.Y - p1.Y
	distance := float32(math.Sqrt(float64(xvel*xvel + yvel*yvel)))
	if distance > maxspeed {
		scale := maxspeed / distance
		xvel *= scale
		yvel *= scale
	}
	return xvel, yvel
}

// MARK: CAM
func SCROLL() {
	siz := float32(32)
	_, up := dSQLineMouseOver(float32(SCRW)-siz, 0, siz, z.Orange, z.Magenta)
	if up && CAM.Target.Y > CNT.Y {
		CAM.Target.Y -= scrollSPD
	}
	_, down := dSQLineMouseOver(float32(SCRW)-siz, float32(SCRH)-siz, siz, z.Orange, z.Magenta)
	if down {
		CAM.Target.Y += scrollSPD
	}
}

// MARK: FILES
func qImagePaths(directoryPath string) []string {
	// Open the directory and list its entries.
	entries, _ := os.ReadDir(directoryPath)
	var imagePaths []string
	for _, entry := range entries {
		// Skip if the entry is a directory.
		if entry.IsDir() {
			continue
		}
		// Convert the filename to lowercase for case-insensitive comparison.
		lowerName := strings.ToLower(entry.Name())
		// Check if the file ends with .jpg or .png.
		if strings.HasSuffix(lowerName, ".jpg") || strings.HasSuffix(lowerName, ".png") {
			fullPath := filepath.Join(directoryPath, entry.Name())
			imagePaths = append(imagePaths, fullPath)
		}
	}
	return imagePaths
}

// MARK: RANDOM NUMBERS
func RPICKINT(nums []int) int {
	return nums[rand.IntN(len(nums))]
}
func RINT(min, max int) int {
	return min + rand.IntN(max-min)
}
func RF32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}
func FLIPCOIN() bool {
	return rand.IntN(2) == 0
}
func RF64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
func RUINT8(min, max int) uint8 {
	return uint8(min + rand.IntN(max-min))
}
func ROLL6() int {
	return RINT(1, 7)
}
func ROLL12() int {
	return RINT(1, 13)
}
func ROLL18() int {
	return RINT(1, 19)
}
func ROLL24() int {
	return RINT(1, 25)
}
func ROLL30() int {
	return RINT(1, 31)
}
func ROLL36() int {
	return RINT(1, 37)
}

// MARK:COLORS
func cRAN() z.Color {
	return z.NewColor(RUINT8(0, 256), RUINT8(0, 256), RUINT8(0, 256), 255)
}
