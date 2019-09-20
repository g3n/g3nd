package other

import (
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
	"github.com/g3n/engine/util"
	"github.com/g3n/engine/window"
	"github.com/g3n/g3nd/app"
	"time"
)

func init() {
	app.DemoMap["other.tank"] = &TankTest{}
}

const (
	CMD_FORWARD = iota
	CMD_BACKWARD
	CMD_LEFT
	CMD_RIGHT
	CMD_CANNON_LEFT
	CMD_CANNON_RIGHT
	CMD_CANNON_UP
	CMD_CANNON_DOWN
	CMD_LAST
)

type TankTest struct {
	a        *app.App
	velocity float32        // linear velocity (m/s)
	rotvel   float32        // rotation velocity (rad/s)
	model    *TankModel     // tank model
	commands [CMD_LAST]bool // commands states
}

// Start is called once at the start of the demo.
func (t *TankTest) Start(a *app.App) {

	t.a = a
	a.Orbit().SetEnabled(^camera.OrbitKeys)

	// Add directional white light
	l1 := light.NewDirectional(&math32.Color{1, 1, 1}, 1.0)
	l1.SetPosition(10, 10, 10)
	a.Scene().Add(l1)

	// Show grid helper
	grid := util.NewGridHelper(100, 1, &math32.Color{0.4, 0.4, 0.4})
	a.Scene().Add(grid)

	// Sets camera position
	a.Camera().SetPosition(0, 4, 10)
	pos := a.Camera().Position()
	a.Camera().UpdateSize(pos.Length())
	a.Camera().LookAt(&math32.Vector3{0, 0, 0}, &math32.Vector3{0, 1, 0})

	// Add help label
	label1 := gui.NewLabel("Use ASDW to drive tank\nUse JKLI to move cannon")
	label1.SetFontSize(16)
	label1.SetPosition(10, 10)
	a.DemoPanel().Add(label1)

	// Creates tank model
	t.model = t.newTankModel()
	t.velocity = 5.0
	t.rotvel = 0.8
	a.Scene().Add(t.model.node)

	// Subscribe to key events
	a.Subscribe(window.OnKeyDown, t.onKey)
	a.Subscribe(window.OnKeyUp, t.onKey)
}

// Update is called every frame.
func (t *TankTest) Update(a *app.App, deltaTime time.Duration) {

	if t.commands[CMD_LEFT] || t.commands[CMD_RIGHT] {
		// Calculates angle delta to rotate
		angle := t.rotvel * float32(deltaTime.Seconds())
		if t.commands[CMD_RIGHT] {
			angle = -angle
		}
		t.model.node.RotateY(angle)
		// Rotate the wheel caps
		for i := 0; i < len(t.model.caps); i++ {
			wcap := t.model.caps[i]
			if i%2 == 0 {
				wcap.RotateZ(3 * angle)
			} else {
				wcap.RotateZ(-3 * angle)
			}
		}
	}

	if t.commands[CMD_FORWARD] || t.commands[CMD_BACKWARD] {
		// Calculates the distance to move
		dist := t.velocity * float32(deltaTime.Seconds())
		// Calculates wheel rotation
		var rot = -dist / 0.5

		// Get tank world direction
		var quat math32.Quaternion
		t.model.node.WorldQuaternion(&quat)
		direction := math32.Vector3{1, 0, 0}
		direction.ApplyQuaternion(&quat)
		direction.Normalize()
		direction.MultiplyScalar(dist)
		if t.commands[CMD_BACKWARD] {
			direction.Negate()
			rot = -rot
		}
		// Get tank world position
		var position math32.Vector3
		t.model.node.WorldPosition(&position)
		position.Add(&direction)
		t.model.node.SetPositionVec(&position)
		// Rotate whell caps
		for _, wcap := range t.model.caps {
			wcap.RotateZ(rot)
		}
	}

	if t.commands[CMD_CANNON_LEFT] {
		t.model.nodeTop.RotateY(0.01)
	}
	if t.commands[CMD_CANNON_RIGHT] {
		t.model.nodeTop.RotateY(-0.01)
	}
	if t.commands[CMD_CANNON_UP] || t.commands[CMD_CANNON_DOWN] {
		// Get cannon world direction
		var quat math32.Quaternion
		t.model.meshCannon.WorldQuaternion(&quat)
		direction := math32.Vector3{1, 0, 0}
		direction.ApplyQuaternion(&quat)
		direction.Normalize()
		// Calculates angle with Y vector
		cosElevation := direction.Dot(&math32.Vector3{0, 1, 0})
		elevation := math32.Acos(cosElevation)
		if t.commands[CMD_CANNON_UP] {
			if elevation <= math32.Pi/4 {
				return
			}
			t.model.meshCannon.RotateZ(0.01)
		} else {
			if elevation >= math32.Pi/2-0.01 {
				return
			}
			t.model.meshCannon.RotateZ(-0.01)
		}
	}
}

// Cleanup is called once at the end of the demo.
func (t *TankTest) Cleanup(a *app.App) {}

// Process key events
func (t *TankTest) onKey(evname string, ev interface{}) {

	var state bool
	if evname == window.OnKeyDown {
		state = true
	} else {
		state = false
	}
	kev := ev.(*window.KeyEvent)
	switch kev.Key {
	case window.KeyW:
		t.commands[CMD_FORWARD] = state
	case window.KeyS:
		t.commands[CMD_BACKWARD] = state
	case window.KeyA:
		t.commands[CMD_LEFT] = state
	case window.KeyD:
		t.commands[CMD_RIGHT] = state
	case window.KeyJ:
		t.commands[CMD_CANNON_LEFT] = state
	case window.KeyL:
		t.commands[CMD_CANNON_RIGHT] = state
	case window.KeyI:
		t.commands[CMD_CANNON_UP] = state
	case window.KeyK:
		t.commands[CMD_CANNON_DOWN] = state
	}
}

type TankModel struct {
	node       *core.Node // node with all tank meshes
	meshBase   *graphic.Mesh
	caps       []*graphic.Mesh
	nodeTop    *core.Node
	meshCannon *graphic.Mesh
}

// Builds and returns a new tank model with separate meshes
// for the wheel caps, base, turret and cannon
func (t *TankTest) newTankModel() *TankModel {

	const EPS = 0.01
	const BASE_WIDTH = 3.0
	const BASE_HEIGHT = 0.6
	const BASE_LENGTH = 4.4
	var BASE_COLOR = math32.NewColorHex(0x10c010)
	const WHEEL_RADIUS = 0.5
	const WHEEL_WIDTH = (BASE_WIDTH / 4) * 0.9
	var WHEEL_COLOR = math32.NewColorHex(0x808080)
	const EMBED = 0.5
	const CAP_RADIUS = WHEEL_RADIUS - 0.1
	const CAP_DZ = 0.01
	var CAP_COLOR = math32.NewColorHex(0x707070)
	const TRACK_HEIGHT = 0.05
	var TRACK_COLOR = WHEEL_COLOR
	const TUR_TOP_RADIUS = 1
	const TUR_BOTTOM_RADIUS = BASE_WIDTH / 2
	const TUR_HEIGHT = 1.0
	var TUR_COLOR = math32.NewColorHex(0x10f010)
	const CANNON_TOP_RADIUS = 0.2
	const CANNON_BOTTOM_RADIUS = 0.3
	const CANNON_LENGTH = 3
	var CANNON_COLOR = TUR_COLOR
	var CANNON_TOP_COLOR = math32.NewColorHex(0x102010)

	// Creates tank model
	model := new(TankModel)
	model.node = core.NewNode()

	// Loads tank wheel texture
	texfile := t.a.DirData() + "/images/wheel.png"
	tex, err := texture.NewTexture2DFromImage(texfile)
	if err != nil {
		t.a.Log().Fatal("Error:%s loading texture:%s", err, texfile)
	}

	// Create the tank wheels and add them to the group
	matWheel := material.NewStandard(WHEEL_COLOR)
	matCap := material.NewStandard(CAP_COLOR)
	matCap.AddTexture(tex)
	//defer matWheel.Dispose()
	//defer matCap.Dispose()

	var zWheel float32 = BASE_WIDTH/2 - WHEEL_WIDTH/2 - 0.1
	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			// Creates wheel mesh
			geomWheel := geometry.NewCylinder(WHEEL_RADIUS, WHEEL_RADIUS, WHEEL_WIDTH+EPS, 20, 20, 0, 2*math32.Pi, true, true)
			meshWheel := graphic.NewMesh(geomWheel, matWheel)
			var zdir float32 = 1.0
			if j%2 == 0 {
				zdir = -1.0
			}
			meshWheel.SetPosition(
				-(BASE_LENGTH/2)+WHEEL_RADIUS+0.2+float32(i)*(WHEEL_RADIUS*2),
				WHEEL_RADIUS,
				zdir*zWheel,
			)
			meshWheel.SetRotationX(math32.Pi / 2)
			model.node.Add(meshWheel)

			// Creates wheel cap
			geomCap := geometry.NewCircle(CAP_RADIUS, 20)
			meshCap := graphic.NewMesh(geomCap, matCap)
			meshCap.SetPositionX(-(BASE_LENGTH / 2) + WHEEL_RADIUS + 0.2 + float32(i)*(WHEEL_RADIUS*2))
			meshCap.SetPositionY(WHEEL_RADIUS)
			zWheelCap := zWheel + WHEEL_WIDTH/2 + CAP_DZ
			if j%2 != 0 {
				meshCap.SetPositionZ(zWheelCap)
			} else {
				meshCap.SetPositionZ(-zWheelCap)
			}
			// Rotate the wheel cap circle geometry
			if j%2 == 0 {
				geomCap.ApplyMatrix(math32.NewMatrix4().MakeRotationX(-math32.Pi))
			}
			model.caps = append(model.caps, meshCap)
			model.node.Add(meshCap)
		}
	}

	// Creates the wheel tracks
	for i := 0; i < 2; i++ {
		geomTrack := geometry.NewBox(WHEEL_RADIUS*6, WHEEL_WIDTH+EPS, TRACK_HEIGHT)
		matTrack := material.NewStandard(TRACK_COLOR)
		meshTrack := graphic.NewMesh(geomTrack, matTrack)
		var zdir float32 = 1.0
		if i%2 == 0 {
			zdir = -1.0
		}
		meshTrack.SetPosition(0, TRACK_HEIGHT/2, zdir*zWheel)
		meshTrack.SetRotationX(math32.Pi / 2)
		model.node.Add(meshTrack)
	}

	// Creates the tank base chassis and adds it to the group
	geomBase := geometry.NewBox(BASE_LENGTH, BASE_WIDTH, BASE_HEIGHT)
	matBase := material.NewStandard(BASE_COLOR)
	model.meshBase = graphic.NewMesh(geomBase, matBase)
	model.meshBase.SetPosition(
		0,
		2*WHEEL_RADIUS+(BASE_HEIGHT/2)-EMBED,
		0,
	)
	model.meshBase.SetRotationX(math32.Pi / 2)
	model.node.Add(model.meshBase)

	// Create node top for the turret and cannon
	model.nodeTop = core.NewNode()

	// Create the turret mesh and add it to the top group
	geomTurret := geometry.NewCylinder(TUR_TOP_RADIUS, TUR_BOTTOM_RADIUS, TUR_HEIGHT, 32, 32, 0, 2*math32.Pi, true, true)
	matTurret := material.NewStandard(TUR_COLOR)
	meshTurret := graphic.NewMesh(geomTurret, matTurret)
	meshTurret.SetPositionY(model.meshBase.Position().Y + TUR_HEIGHT/2)
	model.nodeTop.Add(meshTurret)

	// Create the cannon geometry
	geomCannon := geometry.NewCylinder(CANNON_TOP_RADIUS, CANNON_BOTTOM_RADIUS, CANNON_LENGTH, 20, 20, 0, 2*math32.Pi, true, true)
	matCannon := material.NewStandard(CANNON_COLOR)
	matCannonTop := material.NewStandard(CANNON_TOP_COLOR)
	model.meshCannon = graphic.NewMesh(geomCannon, nil)
	model.meshCannon.AddGroupMaterial(matCannon, 0)
	model.meshCannon.AddGroupMaterial(matCannonTop, 1)
	// translate the cannon geometry so its origin is at its start (not the middle)
	// from which it will rotate.
	geomCannon.ApplyMatrix(math32.NewMatrix4().MakeTranslation(0, CANNON_LENGTH/2, 0))
	model.meshCannon.SetPosition(
		0,
		meshTurret.Position().Y+0.1,
		0,
	)
	geomCannon.ApplyMatrix(math32.NewMatrix4().MakeRotationZ(-math32.Pi / 2))
	model.nodeTop.Add(model.meshCannon)

	model.node.Add(model.nodeTop)
	return model
}
