// Code generated by kds. DO NOT EDIT.
// source: TODO: Source File

package kds;

import (
	"github.com/iakud/krocher/kds/kdsc/example/kdspb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type syncableCity struct {
	PlayerId int64
	PlayerBasicInfo *PlayerBasicInfo
	CityInfo *CityBaseInfo
}

type City struct {
	id int64
	syncable syncableCity

	dirty uint64
}

func NewCity() *City {
	x := new(City)
	x.dirty = 1
	x.id = 0 // FIXME: gen nextId()
	x.setPlayerBasicInfo(NewPlayerBasicInfo())
	x.setCityInfo(NewCityBaseInfo())
	return x
}

func (x *City) Id() int64 {
	return x.id
}

func (x *City) GetPlayerId() int64 {
	return x.syncable.PlayerId
}

func (x *City) SetPlayerId(v int64) {
	if v == x.syncable.PlayerId {
		return
	}
	x.syncable.PlayerId = v
	x.markDirty(uint64(0x01) << 1)
}

func (x *City) GetPlayerBasicInfo() *PlayerBasicInfo {
	return x.syncable.PlayerBasicInfo
}

func (x *City) setPlayerBasicInfo(v *PlayerBasicInfo) {
	if v != nil && v.dirtyParent != nil {
		panic("the component should be removed or evicted from its original place first")
	}
	if v == x.syncable.PlayerBasicInfo {
		return
	}
	if x.syncable.PlayerBasicInfo != nil {
		x.syncable.PlayerBasicInfo.dirtyParent = nil
	}
	x.syncable.PlayerBasicInfo = v
	v.dirtyParent = func() {
		x.markDirty(uint64(0x01) << 2)
	}
	x.markDirty(uint64(0x01) << 2)
	if v != nil {
		v.markDirty(uint64(0x01))
	}
}

func (x *City) GetCityInfo() *CityBaseInfo {
	return x.syncable.CityInfo
}

func (x *City) setCityInfo(v *CityBaseInfo) {
	if v != nil && v.dirtyParent != nil {
		panic("the component should be removed or evicted from its original place first")
	}
	if v == x.syncable.CityInfo {
		return
	}
	if x.syncable.CityInfo != nil {
		x.syncable.CityInfo.dirtyParent = nil
	}
	x.syncable.CityInfo = v
	v.dirtyParent = func() {
		x.markDirty(uint64(0x01) << 3)
	}
	x.markDirty(uint64(0x01) << 3)
	if v != nil {
		v.markDirty(uint64(0x01))
	}
}

func (x *City) DumpChange() *kdspb.City {
	if x.checkDirty(uint64(0x01)) {
		return x.DumpFull()
	}
	m := new(kdspb.City)
	if x.checkDirty(uint64(0x01) << 1) {
		m.PlayerId = x.syncable.PlayerId
	}
	if x.checkDirty(uint64(0x01) << 2) {
		m.PlayerBasicInfo = x.syncable.PlayerBasicInfo.DumpChange()
	}
	if x.checkDirty(uint64(0x01) << 3) {
		m.CityInfo = x.syncable.CityInfo.DumpChange()
	}
	return m
}

func (x *City) DumpFull() *kdspb.City {
	m := new(kdspb.City)
	m.PlayerId = x.syncable.PlayerId
	m.PlayerBasicInfo = x.syncable.PlayerBasicInfo.DumpFull()
	m.CityInfo = x.syncable.CityInfo.DumpFull()
	return m
}

func (x *City) markAll() {
	x.dirty = uint64(0x01)
}

func (x *City) markDirty(n uint64) {
	if x.dirty & n == n {
		return
	}
	x.dirty |= n
}

func (x *City) clearAll() {
	x.syncable.PlayerBasicInfo.clearDirty()
	x.syncable.CityInfo.clearDirty()
	x.dirty = 0
}

func (x *City) clearDirty() {
	if x.dirty == 0 {
		return
	}
	if x.dirty & uint64(0x01) != 0 {
		x.clearAll()
		return
	}
	if x.dirty & uint64(0x01) << 2 != 0 {
		x.syncable.PlayerBasicInfo.clearDirty()
	}
	if x.dirty & uint64(0x01) << 3 != 0 {
		x.syncable.CityInfo.clearDirty()
	}
	x.dirty = 0
}

func (x *City) checkDirty(n uint64) bool {
	return x.dirty & n != 0
}

type syncableCityBaseInfo struct {
	Positions []*Vector
	Troops map[int32]struct{}
	BuildInfo []byte
}

type dirtyParentFunc_CityBaseInfo func()

func (f dirtyParentFunc_CityBaseInfo) invoke() {
	if f == nil {
		return
	}
	f()
}

type CityBaseInfo struct {
	syncable syncableCityBaseInfo

	dirty uint64
	dirtyParent dirtyParentFunc_CityBaseInfo
}

func NewCityBaseInfo() *CityBaseInfo {
	x := new(CityBaseInfo)
	x.dirty = 1
	x.syncable.Troops = make(map[int32]struct{})
	return x
}

func (x *CityBaseInfo) GetBuildInfo() []byte {
	return x.syncable.BuildInfo
}

func (x *CityBaseInfo) SetBuildInfo(v []byte) {
	if v != nil || x.syncable.BuildInfo != nil {
		return
	}
	x.syncable.BuildInfo = v
	x.markDirty(uint64(0x01) << 3)
}

func (x *CityBaseInfo) DumpChange() *kdspb.CityBaseInfo {
	if x.checkDirty(uint64(0x01)) {
		return x.DumpFull()
	}
	m := new(kdspb.CityBaseInfo)
	if x.checkDirty(uint64(0x01) << 1) {
		for _, v := range x.syncable.Positions {
			m.Positions = append(m.Positions, v.DumpChange())
		}
	}
	if x.checkDirty(uint64(0x01) << 2) {
		for k := range x.syncable.Troops {
			m.Troops[k] = new(emptypb.Empty)
		}
	}
	if x.checkDirty(uint64(0x01) << 3) {
		m.BuildInfo = x.syncable.BuildInfo
	}
	return m
}

func (x *CityBaseInfo) DumpFull() *kdspb.CityBaseInfo {
	m := new(kdspb.CityBaseInfo)
	for _, v := range x.syncable.Positions {
		m.Positions = append(m.Positions, v.DumpFull())
	}
	for k := range x.syncable.Troops {
		m.Troops[k] = new(emptypb.Empty)
	}
	m.BuildInfo = x.syncable.BuildInfo
	return m
}

func (x *CityBaseInfo) markAll() {
	x.dirty = uint64(0x01)
}

func (x *CityBaseInfo) markDirty(n uint64) {
	if x.dirty & n == n {
		return
	}
	x.dirty |= n
	x.dirtyParent.invoke()
}

func (x *CityBaseInfo) clearAll() {
	for i := 0; i < len(x.syncable.Positions); i++ {
		x.syncable.Positions[i].clearDirty()
	}
	x.dirty = 0
}

func (x *CityBaseInfo) clearDirty() {
	if x.dirty == 0 {
		return
	}
	if x.dirty & uint64(0x01) != 0 {
		x.clearAll()
		return
	}
	if x.dirty & uint64(0x01) << 1 != 0 {
		for i := 0; i < len(x.syncable.Positions); i++ {
			x.syncable.Positions[i].clearDirty()
		}
	}
	x.dirty = 0
}

func (x *CityBaseInfo) checkDirty(n uint64) bool {
	return x.dirty & n != 0
}

type syncableVector struct {
	X int32
	Y int32
}

type dirtyParentFunc_Vector func()

func (f dirtyParentFunc_Vector) invoke() {
	if f == nil {
		return
	}
	f()
}

type Vector struct {
	syncable syncableVector

	dirty uint64
	dirtyParent dirtyParentFunc_Vector
}

func NewVector() *Vector {
	x := new(Vector)
	x.dirty = 1
	return x
}

func (x *Vector) GetX() int32 {
	return x.syncable.X
}

func (x *Vector) SetX(v int32) {
	if v == x.syncable.X {
		return
	}
	x.syncable.X = v
	x.markDirty(uint64(0x01) << 1)
}

func (x *Vector) GetY() int32 {
	return x.syncable.Y
}

func (x *Vector) SetY(v int32) {
	if v == x.syncable.Y {
		return
	}
	x.syncable.Y = v
	x.markDirty(uint64(0x01) << 2)
}

func (x *Vector) DumpChange() *kdspb.Vector {
	if x.checkDirty(uint64(0x01)) {
		return x.DumpFull()
	}
	m := new(kdspb.Vector)
	if x.checkDirty(uint64(0x01) << 1) {
		m.X = x.syncable.X
	}
	if x.checkDirty(uint64(0x01) << 2) {
		m.Y = x.syncable.Y
	}
	return m
}

func (x *Vector) DumpFull() *kdspb.Vector {
	m := new(kdspb.Vector)
	m.X = x.syncable.X
	m.Y = x.syncable.Y
	return m
}

func (x *Vector) markAll() {
	x.dirty = uint64(0x01)
}

func (x *Vector) markDirty(n uint64) {
	if x.dirty & n == n {
		return
	}
	x.dirty |= n
	x.dirtyParent.invoke()
}

func (x *Vector) clearAll() {
	x.dirty = 0
}

func (x *Vector) clearDirty() {
	if x.dirty == 0 {
		return
	}
	if x.dirty & uint64(0x01) != 0 {
		x.clearAll()
		return
	}
	x.dirty = 0
}

func (x *Vector) checkDirty(n uint64) bool {
	return x.dirty & n != 0
}