// DO NOT EDIT - THIS FILE IS GENERATED
package protocol

import "encoding/binary"

//////////////////////////////////////
// Enums

//////////////////////////////////////

type ButtonGesture uint16

const BUTTON_GESTURE_PRESS ButtonGesture = 1
const BUTTON_GESTURE_HOLD ButtonGesture = 2
const BUTTON_GESTURE_PRESS_PRESS ButtonGesture = 3
const BUTTON_GESTURE_PRESS_HOLD ButtonGesture = 4
const BUTTON_GESTURE_HOLD_HOLD ButtonGesture = 5

//////////////////////////////////////

type ButtonTargetType uint16

const BUTTON_TARGET_TYPE_RELAYS ButtonTargetType = 2
const BUTTON_TARGET_TYPE_DEVICE ButtonTargetType = 3
const BUTTON_TARGET_TYPE_LOCATION ButtonTargetType = 4
const BUTTON_TARGET_TYPE_GROUP ButtonTargetType = 5
const BUTTON_TARGET_TYPE_SCENE ButtonTargetType = 6
const BUTTON_TARGET_TYPE_DEVICE_RELAYS ButtonTargetType = 7

//////////////////////////////////////

type DeviceService uint8

const DEVICE_SERVICE_UDP DeviceService = 1

//////////////////////////////////////

type LightLastHevCycleResult uint8

const LIGHT_LAST_HEV_CYCLE_RESULT_SUCCESS LightLastHevCycleResult = 0
const LIGHT_LAST_HEV_CYCLE_RESULT_BUSY LightLastHevCycleResult = 1
const LIGHT_LAST_HEV_CYCLE_RESULT_INTERRUPTED_BY_RESET LightLastHevCycleResult = 2
const LIGHT_LAST_HEV_CYCLE_RESULT_INTERRUPTED_BY_HOMEKIT LightLastHevCycleResult = 3
const LIGHT_LAST_HEV_CYCLE_RESULT_INTERRUPTED_BY_LAN LightLastHevCycleResult = 4
const LIGHT_LAST_HEV_CYCLE_RESULT_INTERRUPTED_BY_CLOUD LightLastHevCycleResult = 5
const LIGHT_LAST_HEV_CYCLE_RESULT_NONE LightLastHevCycleResult = 255

//////////////////////////////////////

type LightWaveform uint8

const LIGHT_WAVEFORM_SAW LightWaveform = 0
const LIGHT_WAVEFORM_SINE LightWaveform = 1
const LIGHT_WAVEFORM_HALF_SINE LightWaveform = 2
const LIGHT_WAVEFORM_TRIANGLE LightWaveform = 3
const LIGHT_WAVEFORM_PULSE LightWaveform = 4

//////////////////////////////////////

type MultiZoneApplicationRequest uint8

const MULTI_ZONE_APPLICATION_REQUEST_NO_APPLY MultiZoneApplicationRequest = 0
const MULTI_ZONE_APPLICATION_REQUEST_APPLY MultiZoneApplicationRequest = 1
const MULTI_ZONE_APPLICATION_REQUEST_APPLY_ONLY MultiZoneApplicationRequest = 2

//////////////////////////////////////

type MultiZoneEffectType uint8

const MULTI_ZONE_EFFECT_TYPE_OFF MultiZoneEffectType = 0
const MULTI_ZONE_EFFECT_TYPE_MOVE MultiZoneEffectType = 1

//////////////////////////////////////

type MultiZoneExtendedApplicationRequest uint8

const MULTI_ZONE_EXTENDED_APPLICATION_REQUEST_NO_APPLY MultiZoneExtendedApplicationRequest = 0
const MULTI_ZONE_EXTENDED_APPLICATION_REQUEST_APPLY MultiZoneExtendedApplicationRequest = 1
const MULTI_ZONE_EXTENDED_APPLICATION_REQUEST_APPLY_ONLY MultiZoneExtendedApplicationRequest = 2

//////////////////////////////////////

type TileEffectType uint8

const TILE_EFFECT_TYPE_OFF TileEffectType = 0
const TILE_EFFECT_TYPE_MORPH TileEffectType = 2
const TILE_EFFECT_TYPE_FLAME TileEffectType = 3

//////////////////////////////////////
// Fields

// ////////////////////////////////////
// Button is 101 bytes large
type Button struct {
	ActionsCount uint8           // 1 Bytes
	Actions      [5]ButtonAction // 100 Bytes
}

func (d Button) Write(w io.Writer) error {
	return binary.Write(w, binary.LittleEndian, d)
}
func (d *Button) Read(r io.Reader) error {
	return binary.Read(r, binary.LittleEndian, d)
}

// ////////////////////////////////////
// ButtonAction is 20 bytes large
type ButtonAction struct {
	Gesture    ButtonGesture    // 2 Bytes
	TargetType ButtonTargetType // 2 Bytes
	Target     ButtonTarget     // 16 Bytes
}

func (d ButtonAction) Write(w io.Writer) error {
	return binary.Write(w, binary.LittleEndian, d)
}
func (d *ButtonAction) Read(r io.Reader) error {
	return binary.Read(r, binary.LittleEndian, d)
}

// ////////////////////////////////////
// ButtonBacklightHsbk is 8 bytes large
type ButtonBacklightHsbk struct {
	Hue        uint16 // 2 Bytes
	Saturation uint16 // 2 Bytes
	Brightness uint16 // 2 Bytes
	Kelvin     uint16 // 2 Bytes
}

func (d ButtonBacklightHsbk) Write(w io.Writer) error {
	return binary.Write(w, binary.LittleEndian, d)
}
func (d *ButtonBacklightHsbk) Read(r io.Reader) error {
	return binary.Read(r, binary.LittleEndian, d)
}

// ////////////////////////////////////
// ButtonTargetDevice is 16 bytes large
type ButtonTargetDevice struct {
	Serial   [6]byte  // 6 Bytes
	Reserved [10]byte // 10 Bytes
}

func (d ButtonTargetDevice) Write(w io.Writer) error {
	return binary.Write(w, binary.LittleEndian, d)
}
func (d *ButtonTargetDevice) Read(r io.Reader) error {
	return binary.Read(r, binary.LittleEndian, d)
}

// ////////////////////////////////////
// ButtonTargetDeviceRelays is 16 bytes large
type ButtonTargetDeviceRelays struct {
	Serial      [6]byte  // 6 Bytes
	RelaysCount uint8    // 1 Bytes
	Relays      [9]uint8 // 9 Bytes
}

func (d ButtonTargetDeviceRelays) Write(w io.Writer) error {
	return binary.Write(w, binary.LittleEndian, d)
}
func (d *ButtonTargetDeviceRelays) Read(r io.Reader) error {
	return binary.Read(r, binary.LittleEndian, d)
}

// ////////////////////////////////////
// ButtonTargetRelays is 16 bytes large
type ButtonTargetRelays struct {
	RelaysCount uint8     // 1 Bytes
	Relays      [15]uint8 // 15 Bytes
}

func (d ButtonTargetRelays) Write(w io.Writer) error {
	return binary.Write(w, binary.LittleEndian, d)
}
func (d *ButtonTargetRelays) Read(r io.Reader) error {
	return binary.Read(r, binary.LittleEndian, d)
}

// ////////////////////////////////////
// LightHsbk is 8 bytes large
type LightHsbk struct {
	Hue        uint16 // 2 Bytes
	Saturation uint16 // 2 Bytes
	Brightness uint16 // 2 Bytes
	Kelvin     uint16 // 2 Bytes
}

func (d LightHsbk) Write(w io.Writer) error {
	return binary.Write(w, binary.LittleEndian, d)
}
func (d *LightHsbk) Read(r io.Reader) error {
	return binary.Read(r, binary.LittleEndian, d)
}

// ////////////////////////////////////
// MultiZoneEffectParameter is 32 bytes large
type MultiZoneEffectParameter struct {
	Parameter0 uint32 // 4 Bytes
	Parameter1 uint32 // 4 Bytes
	Parameter2 uint32 // 4 Bytes
	Parameter3 uint32 // 4 Bytes
	Parameter4 uint32 // 4 Bytes
	Parameter5 uint32 // 4 Bytes
	Parameter6 uint32 // 4 Bytes
	Parameter7 uint32 // 4 Bytes
}

func (d MultiZoneEffectParameter) Write(w io.Writer) error {
	return binary.Write(w, binary.LittleEndian, d)
}
func (d *MultiZoneEffectParameter) Read(r io.Reader) error {
	return binary.Read(r, binary.LittleEndian, d)
}

// ////////////////////////////////////
// MultiZoneEffectSettings is 59 bytes large
type MultiZoneEffectSettings struct {
	Instanceid uint32                   // 4 Bytes
	Type       MultiZoneEffectType      // 1 Bytes
	_          [2]byte                  // 2 Bytes
	Speed      uint32                   // 4 Bytes
	Duration   uint64                   // 8 Bytes
	_          [4]byte                  // 4 Bytes
	_          [4]byte                  // 4 Bytes
	Parameter  MultiZoneEffectParameter // 32 Bytes
}

func (d MultiZoneEffectSettings) Write(w io.Writer) error {
	return binary.Write(w, binary.LittleEndian, d)
}
func (d *MultiZoneEffectSettings) Read(r io.Reader) error {
	return binary.Read(r, binary.LittleEndian, d)
}

// ////////////////////////////////////
// TileAccelMeas is 6 bytes large
type TileAccelMeas struct {
	X int16 // 2 Bytes
	Y int16 // 2 Bytes
	Z int16 // 2 Bytes
}

func (d TileAccelMeas) Write(w io.Writer) error {
	return binary.Write(w, binary.LittleEndian, d)
}
func (d *TileAccelMeas) Read(r io.Reader) error {
	return binary.Read(r, binary.LittleEndian, d)
}

// ////////////////////////////////////
// TileBufferRect is 4 bytes large
type TileBufferRect struct {
	_     [1]byte // 1 Bytes
	X     uint8   // 1 Bytes
	Y     uint8   // 1 Bytes
	Width uint8   // 1 Bytes
}

func (d TileBufferRect) Write(w io.Writer) error {
	return binary.Write(w, binary.LittleEndian, d)
}
func (d *TileBufferRect) Read(r io.Reader) error {
	return binary.Read(r, binary.LittleEndian, d)
}

// ////////////////////////////////////
// TileEffectParameter is 32 bytes large
type TileEffectParameter struct {
	Parameter0 uint32 // 4 Bytes
	Parameter1 uint32 // 4 Bytes
	Parameter2 uint32 // 4 Bytes
	Parameter3 uint32 // 4 Bytes
	Parameter4 uint32 // 4 Bytes
	Parameter5 uint32 // 4 Bytes
	Parameter6 uint32 // 4 Bytes
	Parameter7 uint32 // 4 Bytes
}

func (d TileEffectParameter) Write(w io.Writer) error {
	return binary.Write(w, binary.LittleEndian, d)
}
func (d *TileEffectParameter) Read(r io.Reader) error {
	return binary.Read(r, binary.LittleEndian, d)
}

// ////////////////////////////////////
// TileEffectSettings is 186 bytes large
type TileEffectSettings struct {
	Instanceid   uint32              // 4 Bytes
	Type         TileEffectType      // 1 Bytes
	Speed        uint32              // 4 Bytes
	Duration     uint64              // 8 Bytes
	_            [4]byte             // 4 Bytes
	_            [4]byte             // 4 Bytes
	Parameter    TileEffectParameter // 32 Bytes
	PaletteCount uint8               // 1 Bytes
	Palette      [16]LightHsbk       // 128 Bytes
}

func (d TileEffectSettings) Write(w io.Writer) error {
	return binary.Write(w, binary.LittleEndian, d)
}
func (d *TileEffectSettings) Read(r io.Reader) error {
	return binary.Read(r, binary.LittleEndian, d)
}

// ////////////////////////////////////
// TileStateDevice is 55 bytes large
type TileStateDevice struct {
	AccelMeas     TileAccelMeas           // 6 Bytes
	_             [2]byte                 // 2 Bytes
	UserX         float32                 // 4 Bytes
	UserY         float32                 // 4 Bytes
	Width         uint8                   // 1 Bytes
	Height        uint8                   // 1 Bytes
	_             [1]byte                 // 1 Bytes
	DeviceVersion DeviceStateVersion      // 12 Bytes
	Firmware      DeviceStateHostFirmware // 20 Bytes
	_             [4]byte                 // 4 Bytes
}

func (d TileStateDevice) Write(w io.Writer) error {
	return binary.Write(w, binary.LittleEndian, d)
}
func (d *TileStateDevice) Read(r io.Reader) error {
	return binary.Read(r, binary.LittleEndian, d)
}
