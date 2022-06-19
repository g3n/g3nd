package movement

import (
	"github.com/g3n/engine/math32"
)

//save some garbage collection
var (
	usePos              math32.Vector3
	flDifference, dtime = float32(0.0), float32(0.0)
	//vecUpHat            = math32.NewVector3(0, 1, 0)
	incRot, incLinear, incAcceleration        = float32(0.0), float32(0.0), float32(2)
	vecSphere, vecCube, vecGopher, vecS, vecC math32.Vector3
	//these are the vectors that makes flying / looking where you're running possible
	vecViewTmp, vecViewForward, vecViewRight, vecViewUp math32.Vector3
)

//"constant" vars
var (
	zeroVector   math32.Vector3 = *math32.NewVector3(0, 0, 0)
	cameraVector math32.Vector3 = *math32.NewVector3(15, 4, -2)
)

const (
	incrementRotFly       = float32(0.004)
	incrementRotTranslate = float32(0.02)
	incrementLinear       = float32(0.005)
)
