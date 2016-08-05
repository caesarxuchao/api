/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// ************************************************************
// DO NOT EDIT.
// THIS FILE IS AUTO-GENERATED BY codecgen.
// ************************************************************

package policy

import (
	"errors"
	"fmt"
	codec1978 "github.com/ugorji/go/codec"
	pkg3_api "k8s.io/kubernetes/pkg/api"
	pkg2_unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	pkg4_types "k8s.io/kubernetes/pkg/types"
	pkg1_intstr "k8s.io/kubernetes/pkg/util/intstr"
	"reflect"
	"runtime"
	time "time"
)

const (
	// ----- content types ----
	codecSelferC_UTF81234 = 1
	codecSelferC_RAW1234  = 0
	// ----- value types used ----
	codecSelferValueTypeArray1234 = 10
	codecSelferValueTypeMap1234   = 9
	// ----- containerStateValues ----
	codecSelfer_containerMapKey1234    = 2
	codecSelfer_containerMapValue1234  = 3
	codecSelfer_containerMapEnd1234    = 4
	codecSelfer_containerArrayElem1234 = 6
	codecSelfer_containerArrayEnd1234  = 7
)

var (
	codecSelferBitsize1234                         = uint8(reflect.TypeOf(uint(0)).Bits())
	codecSelferOnlyMapOrArrayEncodeToStructErr1234 = errors.New(`only encoded map or array can be decoded into a struct`)
)

type codecSelfer1234 struct{}

func init() {
	if codec1978.GenVersion != 5 {
		_, file, _, _ := runtime.Caller(0)
		err := fmt.Errorf("codecgen version mismatch: current: %v, need %v. Re-generate file: %v",
			5, codec1978.GenVersion, file)
		panic(err)
	}
	if false { // reference the types, but skip this branch at build/run time
		var v0 pkg3_api.ObjectMeta
		var v1 pkg2_unversioned.LabelSelector
		var v2 pkg4_types.UID
		var v3 pkg1_intstr.IntOrString
		var v4 time.Time
		_, _, _, _, _ = v0, v1, v2, v3, v4
	}
}

func (x *PodDisruptionBudgetSpec) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			var yyq2 [2]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			yyq2[0] = true
			yyq2[1] = x.Selector != nil
			var yynn2 int
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(2)
			} else {
				yynn2 = 0
				for _, b := range yyq2 {
					if b {
						yynn2++
					}
				}
				r.EncodeMapStart(yynn2)
				yynn2 = 0
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[0] {
					yy4 := &x.MinAvailable
					yym5 := z.EncBinary()
					_ = yym5
					if false {
					} else if z.HasExtensions() && z.EncExt(yy4) {
					} else if !yym5 && z.IsJSONHandle() {
						z.EncJSONMarshal(yy4)
					} else {
						z.EncFallback(yy4)
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("minAvailable"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy6 := &x.MinAvailable
					yym7 := z.EncBinary()
					_ = yym7
					if false {
					} else if z.HasExtensions() && z.EncExt(yy6) {
					} else if !yym7 && z.IsJSONHandle() {
						z.EncJSONMarshal(yy6)
					} else {
						z.EncFallback(yy6)
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[1] {
					if x.Selector == nil {
						r.EncodeNil()
					} else {
						yym9 := z.EncBinary()
						_ = yym9
						if false {
						} else if z.HasExtensions() && z.EncExt(x.Selector) {
						} else {
							z.EncFallback(x.Selector)
						}
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("selector"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					if x.Selector == nil {
						r.EncodeNil()
					} else {
						yym10 := z.EncBinary()
						_ = yym10
						if false {
						} else if z.HasExtensions() && z.EncExt(x.Selector) {
						} else {
							z.EncFallback(x.Selector)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *PodDisruptionBudgetSpec) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym11 := z.DecBinary()
	_ = yym11
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct12 := r.ContainerType()
		if yyct12 == codecSelferValueTypeMap1234 {
			yyl12 := r.ReadMapStart()
			if yyl12 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl12, d)
			}
		} else if yyct12 == codecSelferValueTypeArray1234 {
			yyl12 := r.ReadArrayStart()
			if yyl12 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl12, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *PodDisruptionBudgetSpec) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys13Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys13Slc
	var yyhl13 bool = l >= 0
	for yyj13 := 0; ; yyj13++ {
		if yyhl13 {
			if yyj13 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys13Slc = r.DecodeBytes(yys13Slc, true, true)
		yys13 := string(yys13Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys13 {
		case "minAvailable":
			if r.TryDecodeAsNil() {
				x.MinAvailable = pkg1_intstr.IntOrString{}
			} else {
				yyv14 := &x.MinAvailable
				yym15 := z.DecBinary()
				_ = yym15
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv14) {
				} else if !yym15 && z.IsJSONHandle() {
					z.DecJSONUnmarshal(yyv14)
				} else {
					z.DecFallback(yyv14, false)
				}
			}
		case "selector":
			if r.TryDecodeAsNil() {
				if x.Selector != nil {
					x.Selector = nil
				}
			} else {
				if x.Selector == nil {
					x.Selector = new(pkg2_unversioned.LabelSelector)
				}
				yym17 := z.DecBinary()
				_ = yym17
				if false {
				} else if z.HasExtensions() && z.DecExt(x.Selector) {
				} else {
					z.DecFallback(x.Selector, false)
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys13)
		} // end switch yys13
	} // end for yyj13
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *PodDisruptionBudgetSpec) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj18 int
	var yyb18 bool
	var yyhl18 bool = l >= 0
	yyj18++
	if yyhl18 {
		yyb18 = yyj18 > l
	} else {
		yyb18 = r.CheckBreak()
	}
	if yyb18 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.MinAvailable = pkg1_intstr.IntOrString{}
	} else {
		yyv19 := &x.MinAvailable
		yym20 := z.DecBinary()
		_ = yym20
		if false {
		} else if z.HasExtensions() && z.DecExt(yyv19) {
		} else if !yym20 && z.IsJSONHandle() {
			z.DecJSONUnmarshal(yyv19)
		} else {
			z.DecFallback(yyv19, false)
		}
	}
	yyj18++
	if yyhl18 {
		yyb18 = yyj18 > l
	} else {
		yyb18 = r.CheckBreak()
	}
	if yyb18 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		if x.Selector != nil {
			x.Selector = nil
		}
	} else {
		if x.Selector == nil {
			x.Selector = new(pkg2_unversioned.LabelSelector)
		}
		yym22 := z.DecBinary()
		_ = yym22
		if false {
		} else if z.HasExtensions() && z.DecExt(x.Selector) {
		} else {
			z.DecFallback(x.Selector, false)
		}
	}
	for {
		yyj18++
		if yyhl18 {
			yyb18 = yyj18 > l
		} else {
			yyb18 = r.CheckBreak()
		}
		if yyb18 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj18-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *PodDisruptionBudgetStatus) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym23 := z.EncBinary()
		_ = yym23
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep24 := !z.EncBinary()
			yy2arr24 := z.EncBasicHandle().StructToArray
			var yyq24 [4]bool
			_, _, _ = yysep24, yyq24, yy2arr24
			const yyr24 bool = false
			var yynn24 int
			if yyr24 || yy2arr24 {
				r.EncodeArrayStart(4)
			} else {
				yynn24 = 4
				for _, b := range yyq24 {
					if b {
						yynn24++
					}
				}
				r.EncodeMapStart(yynn24)
				yynn24 = 0
			}
			if yyr24 || yy2arr24 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				yym26 := z.EncBinary()
				_ = yym26
				if false {
				} else {
					r.EncodeBool(bool(x.PodDisruptionAllowed))
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("disruptionAllowed"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				yym27 := z.EncBinary()
				_ = yym27
				if false {
				} else {
					r.EncodeBool(bool(x.PodDisruptionAllowed))
				}
			}
			if yyr24 || yy2arr24 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				yym29 := z.EncBinary()
				_ = yym29
				if false {
				} else {
					r.EncodeInt(int64(x.CurrentHealthy))
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("currentHealthy"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				yym30 := z.EncBinary()
				_ = yym30
				if false {
				} else {
					r.EncodeInt(int64(x.CurrentHealthy))
				}
			}
			if yyr24 || yy2arr24 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				yym32 := z.EncBinary()
				_ = yym32
				if false {
				} else {
					r.EncodeInt(int64(x.DesiredHealthy))
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("desiredHealthy"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				yym33 := z.EncBinary()
				_ = yym33
				if false {
				} else {
					r.EncodeInt(int64(x.DesiredHealthy))
				}
			}
			if yyr24 || yy2arr24 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				yym35 := z.EncBinary()
				_ = yym35
				if false {
				} else {
					r.EncodeInt(int64(x.ExpectedPods))
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("expectedPods"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				yym36 := z.EncBinary()
				_ = yym36
				if false {
				} else {
					r.EncodeInt(int64(x.ExpectedPods))
				}
			}
			if yyr24 || yy2arr24 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *PodDisruptionBudgetStatus) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym37 := z.DecBinary()
	_ = yym37
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct38 := r.ContainerType()
		if yyct38 == codecSelferValueTypeMap1234 {
			yyl38 := r.ReadMapStart()
			if yyl38 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl38, d)
			}
		} else if yyct38 == codecSelferValueTypeArray1234 {
			yyl38 := r.ReadArrayStart()
			if yyl38 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl38, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *PodDisruptionBudgetStatus) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys39Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys39Slc
	var yyhl39 bool = l >= 0
	for yyj39 := 0; ; yyj39++ {
		if yyhl39 {
			if yyj39 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys39Slc = r.DecodeBytes(yys39Slc, true, true)
		yys39 := string(yys39Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys39 {
		case "disruptionAllowed":
			if r.TryDecodeAsNil() {
				x.PodDisruptionAllowed = false
			} else {
				x.PodDisruptionAllowed = bool(r.DecodeBool())
			}
		case "currentHealthy":
			if r.TryDecodeAsNil() {
				x.CurrentHealthy = 0
			} else {
				x.CurrentHealthy = int32(r.DecodeInt(32))
			}
		case "desiredHealthy":
			if r.TryDecodeAsNil() {
				x.DesiredHealthy = 0
			} else {
				x.DesiredHealthy = int32(r.DecodeInt(32))
			}
		case "expectedPods":
			if r.TryDecodeAsNil() {
				x.ExpectedPods = 0
			} else {
				x.ExpectedPods = int32(r.DecodeInt(32))
			}
		default:
			z.DecStructFieldNotFound(-1, yys39)
		} // end switch yys39
	} // end for yyj39
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *PodDisruptionBudgetStatus) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj44 int
	var yyb44 bool
	var yyhl44 bool = l >= 0
	yyj44++
	if yyhl44 {
		yyb44 = yyj44 > l
	} else {
		yyb44 = r.CheckBreak()
	}
	if yyb44 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.PodDisruptionAllowed = false
	} else {
		x.PodDisruptionAllowed = bool(r.DecodeBool())
	}
	yyj44++
	if yyhl44 {
		yyb44 = yyj44 > l
	} else {
		yyb44 = r.CheckBreak()
	}
	if yyb44 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.CurrentHealthy = 0
	} else {
		x.CurrentHealthy = int32(r.DecodeInt(32))
	}
	yyj44++
	if yyhl44 {
		yyb44 = yyj44 > l
	} else {
		yyb44 = r.CheckBreak()
	}
	if yyb44 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.DesiredHealthy = 0
	} else {
		x.DesiredHealthy = int32(r.DecodeInt(32))
	}
	yyj44++
	if yyhl44 {
		yyb44 = yyj44 > l
	} else {
		yyb44 = r.CheckBreak()
	}
	if yyb44 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.ExpectedPods = 0
	} else {
		x.ExpectedPods = int32(r.DecodeInt(32))
	}
	for {
		yyj44++
		if yyhl44 {
			yyb44 = yyj44 > l
		} else {
			yyb44 = r.CheckBreak()
		}
		if yyb44 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj44-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *PodDisruptionBudget) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym49 := z.EncBinary()
		_ = yym49
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep50 := !z.EncBinary()
			yy2arr50 := z.EncBasicHandle().StructToArray
			var yyq50 [5]bool
			_, _, _ = yysep50, yyq50, yy2arr50
			const yyr50 bool = false
			yyq50[0] = x.Kind != ""
			yyq50[1] = x.APIVersion != ""
			yyq50[2] = true
			yyq50[3] = true
			yyq50[4] = true
			var yynn50 int
			if yyr50 || yy2arr50 {
				r.EncodeArrayStart(5)
			} else {
				yynn50 = 0
				for _, b := range yyq50 {
					if b {
						yynn50++
					}
				}
				r.EncodeMapStart(yynn50)
				yynn50 = 0
			}
			if yyr50 || yy2arr50 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq50[0] {
					yym52 := z.EncBinary()
					_ = yym52
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq50[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("kind"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym53 := z.EncBinary()
					_ = yym53
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				}
			}
			if yyr50 || yy2arr50 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq50[1] {
					yym55 := z.EncBinary()
					_ = yym55
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq50[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("apiVersion"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym56 := z.EncBinary()
					_ = yym56
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				}
			}
			if yyr50 || yy2arr50 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq50[2] {
					yy58 := &x.ObjectMeta
					yy58.CodecEncodeSelf(e)
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq50[2] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("metadata"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy59 := &x.ObjectMeta
					yy59.CodecEncodeSelf(e)
				}
			}
			if yyr50 || yy2arr50 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq50[3] {
					yy61 := &x.Spec
					yy61.CodecEncodeSelf(e)
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq50[3] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("spec"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy62 := &x.Spec
					yy62.CodecEncodeSelf(e)
				}
			}
			if yyr50 || yy2arr50 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq50[4] {
					yy64 := &x.Status
					yy64.CodecEncodeSelf(e)
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq50[4] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("status"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy65 := &x.Status
					yy65.CodecEncodeSelf(e)
				}
			}
			if yyr50 || yy2arr50 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *PodDisruptionBudget) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym66 := z.DecBinary()
	_ = yym66
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct67 := r.ContainerType()
		if yyct67 == codecSelferValueTypeMap1234 {
			yyl67 := r.ReadMapStart()
			if yyl67 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl67, d)
			}
		} else if yyct67 == codecSelferValueTypeArray1234 {
			yyl67 := r.ReadArrayStart()
			if yyl67 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl67, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *PodDisruptionBudget) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys68Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys68Slc
	var yyhl68 bool = l >= 0
	for yyj68 := 0; ; yyj68++ {
		if yyhl68 {
			if yyj68 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys68Slc = r.DecodeBytes(yys68Slc, true, true)
		yys68 := string(yys68Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys68 {
		case "kind":
			if r.TryDecodeAsNil() {
				x.Kind = ""
			} else {
				x.Kind = string(r.DecodeString())
			}
		case "apiVersion":
			if r.TryDecodeAsNil() {
				x.APIVersion = ""
			} else {
				x.APIVersion = string(r.DecodeString())
			}
		case "metadata":
			if r.TryDecodeAsNil() {
				x.ObjectMeta = pkg3_api.ObjectMeta{}
			} else {
				yyv71 := &x.ObjectMeta
				yyv71.CodecDecodeSelf(d)
			}
		case "spec":
			if r.TryDecodeAsNil() {
				x.Spec = PodDisruptionBudgetSpec{}
			} else {
				yyv72 := &x.Spec
				yyv72.CodecDecodeSelf(d)
			}
		case "status":
			if r.TryDecodeAsNil() {
				x.Status = PodDisruptionBudgetStatus{}
			} else {
				yyv73 := &x.Status
				yyv73.CodecDecodeSelf(d)
			}
		default:
			z.DecStructFieldNotFound(-1, yys68)
		} // end switch yys68
	} // end for yyj68
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *PodDisruptionBudget) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj74 int
	var yyb74 bool
	var yyhl74 bool = l >= 0
	yyj74++
	if yyhl74 {
		yyb74 = yyj74 > l
	} else {
		yyb74 = r.CheckBreak()
	}
	if yyb74 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} else {
		x.Kind = string(r.DecodeString())
	}
	yyj74++
	if yyhl74 {
		yyb74 = yyj74 > l
	} else {
		yyb74 = r.CheckBreak()
	}
	if yyb74 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} else {
		x.APIVersion = string(r.DecodeString())
	}
	yyj74++
	if yyhl74 {
		yyb74 = yyj74 > l
	} else {
		yyb74 = r.CheckBreak()
	}
	if yyb74 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.ObjectMeta = pkg3_api.ObjectMeta{}
	} else {
		yyv77 := &x.ObjectMeta
		yyv77.CodecDecodeSelf(d)
	}
	yyj74++
	if yyhl74 {
		yyb74 = yyj74 > l
	} else {
		yyb74 = r.CheckBreak()
	}
	if yyb74 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Spec = PodDisruptionBudgetSpec{}
	} else {
		yyv78 := &x.Spec
		yyv78.CodecDecodeSelf(d)
	}
	yyj74++
	if yyhl74 {
		yyb74 = yyj74 > l
	} else {
		yyb74 = r.CheckBreak()
	}
	if yyb74 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Status = PodDisruptionBudgetStatus{}
	} else {
		yyv79 := &x.Status
		yyv79.CodecDecodeSelf(d)
	}
	for {
		yyj74++
		if yyhl74 {
			yyb74 = yyj74 > l
		} else {
			yyb74 = r.CheckBreak()
		}
		if yyb74 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj74-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *PodDisruptionBudgetList) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym80 := z.EncBinary()
		_ = yym80
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep81 := !z.EncBinary()
			yy2arr81 := z.EncBasicHandle().StructToArray
			var yyq81 [4]bool
			_, _, _ = yysep81, yyq81, yy2arr81
			const yyr81 bool = false
			yyq81[0] = x.Kind != ""
			yyq81[1] = x.APIVersion != ""
			yyq81[2] = true
			var yynn81 int
			if yyr81 || yy2arr81 {
				r.EncodeArrayStart(4)
			} else {
				yynn81 = 1
				for _, b := range yyq81 {
					if b {
						yynn81++
					}
				}
				r.EncodeMapStart(yynn81)
				yynn81 = 0
			}
			if yyr81 || yy2arr81 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq81[0] {
					yym83 := z.EncBinary()
					_ = yym83
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq81[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("kind"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym84 := z.EncBinary()
					_ = yym84
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				}
			}
			if yyr81 || yy2arr81 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq81[1] {
					yym86 := z.EncBinary()
					_ = yym86
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq81[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("apiVersion"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym87 := z.EncBinary()
					_ = yym87
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				}
			}
			if yyr81 || yy2arr81 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq81[2] {
					yy89 := &x.ListMeta
					yym90 := z.EncBinary()
					_ = yym90
					if false {
					} else if z.HasExtensions() && z.EncExt(yy89) {
					} else {
						z.EncFallback(yy89)
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq81[2] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("metadata"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy91 := &x.ListMeta
					yym92 := z.EncBinary()
					_ = yym92
					if false {
					} else if z.HasExtensions() && z.EncExt(yy91) {
					} else {
						z.EncFallback(yy91)
					}
				}
			}
			if yyr81 || yy2arr81 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if x.Items == nil {
					r.EncodeNil()
				} else {
					yym94 := z.EncBinary()
					_ = yym94
					if false {
					} else {
						h.encSlicePodDisruptionBudget(([]PodDisruptionBudget)(x.Items), e)
					}
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("items"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				if x.Items == nil {
					r.EncodeNil()
				} else {
					yym95 := z.EncBinary()
					_ = yym95
					if false {
					} else {
						h.encSlicePodDisruptionBudget(([]PodDisruptionBudget)(x.Items), e)
					}
				}
			}
			if yyr81 || yy2arr81 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *PodDisruptionBudgetList) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym96 := z.DecBinary()
	_ = yym96
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct97 := r.ContainerType()
		if yyct97 == codecSelferValueTypeMap1234 {
			yyl97 := r.ReadMapStart()
			if yyl97 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl97, d)
			}
		} else if yyct97 == codecSelferValueTypeArray1234 {
			yyl97 := r.ReadArrayStart()
			if yyl97 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl97, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *PodDisruptionBudgetList) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys98Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys98Slc
	var yyhl98 bool = l >= 0
	for yyj98 := 0; ; yyj98++ {
		if yyhl98 {
			if yyj98 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys98Slc = r.DecodeBytes(yys98Slc, true, true)
		yys98 := string(yys98Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys98 {
		case "kind":
			if r.TryDecodeAsNil() {
				x.Kind = ""
			} else {
				x.Kind = string(r.DecodeString())
			}
		case "apiVersion":
			if r.TryDecodeAsNil() {
				x.APIVersion = ""
			} else {
				x.APIVersion = string(r.DecodeString())
			}
		case "metadata":
			if r.TryDecodeAsNil() {
				x.ListMeta = pkg2_unversioned.ListMeta{}
			} else {
				yyv101 := &x.ListMeta
				yym102 := z.DecBinary()
				_ = yym102
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv101) {
				} else {
					z.DecFallback(yyv101, false)
				}
			}
		case "items":
			if r.TryDecodeAsNil() {
				x.Items = nil
			} else {
				yyv103 := &x.Items
				yym104 := z.DecBinary()
				_ = yym104
				if false {
				} else {
					h.decSlicePodDisruptionBudget((*[]PodDisruptionBudget)(yyv103), d)
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys98)
		} // end switch yys98
	} // end for yyj98
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *PodDisruptionBudgetList) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj105 int
	var yyb105 bool
	var yyhl105 bool = l >= 0
	yyj105++
	if yyhl105 {
		yyb105 = yyj105 > l
	} else {
		yyb105 = r.CheckBreak()
	}
	if yyb105 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} else {
		x.Kind = string(r.DecodeString())
	}
	yyj105++
	if yyhl105 {
		yyb105 = yyj105 > l
	} else {
		yyb105 = r.CheckBreak()
	}
	if yyb105 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} else {
		x.APIVersion = string(r.DecodeString())
	}
	yyj105++
	if yyhl105 {
		yyb105 = yyj105 > l
	} else {
		yyb105 = r.CheckBreak()
	}
	if yyb105 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.ListMeta = pkg2_unversioned.ListMeta{}
	} else {
		yyv108 := &x.ListMeta
		yym109 := z.DecBinary()
		_ = yym109
		if false {
		} else if z.HasExtensions() && z.DecExt(yyv108) {
		} else {
			z.DecFallback(yyv108, false)
		}
	}
	yyj105++
	if yyhl105 {
		yyb105 = yyj105 > l
	} else {
		yyb105 = r.CheckBreak()
	}
	if yyb105 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Items = nil
	} else {
		yyv110 := &x.Items
		yym111 := z.DecBinary()
		_ = yym111
		if false {
		} else {
			h.decSlicePodDisruptionBudget((*[]PodDisruptionBudget)(yyv110), d)
		}
	}
	for {
		yyj105++
		if yyhl105 {
			yyb105 = yyj105 > l
		} else {
			yyb105 = r.CheckBreak()
		}
		if yyb105 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj105-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x codecSelfer1234) encSlicePodDisruptionBudget(v []PodDisruptionBudget, e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.EncodeArrayStart(len(v))
	for _, yyv112 := range v {
		z.EncSendContainerState(codecSelfer_containerArrayElem1234)
		yy113 := &yyv112
		yy113.CodecEncodeSelf(e)
	}
	z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x codecSelfer1234) decSlicePodDisruptionBudget(v *[]PodDisruptionBudget, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv114 := *v
	yyh114, yyl114 := z.DecSliceHelperStart()
	var yyc114 bool
	if yyl114 == 0 {
		if yyv114 == nil {
			yyv114 = []PodDisruptionBudget{}
			yyc114 = true
		} else if len(yyv114) != 0 {
			yyv114 = yyv114[:0]
			yyc114 = true
		}
	} else if yyl114 > 0 {
		var yyrr114, yyrl114 int
		var yyrt114 bool
		if yyl114 > cap(yyv114) {

			yyrg114 := len(yyv114) > 0
			yyv2114 := yyv114
			yyrl114, yyrt114 = z.DecInferLen(yyl114, z.DecBasicHandle().MaxInitLen, 296)
			if yyrt114 {
				if yyrl114 <= cap(yyv114) {
					yyv114 = yyv114[:yyrl114]
				} else {
					yyv114 = make([]PodDisruptionBudget, yyrl114)
				}
			} else {
				yyv114 = make([]PodDisruptionBudget, yyrl114)
			}
			yyc114 = true
			yyrr114 = len(yyv114)
			if yyrg114 {
				copy(yyv114, yyv2114)
			}
		} else if yyl114 != len(yyv114) {
			yyv114 = yyv114[:yyl114]
			yyc114 = true
		}
		yyj114 := 0
		for ; yyj114 < yyrr114; yyj114++ {
			yyh114.ElemContainerState(yyj114)
			if r.TryDecodeAsNil() {
				yyv114[yyj114] = PodDisruptionBudget{}
			} else {
				yyv115 := &yyv114[yyj114]
				yyv115.CodecDecodeSelf(d)
			}

		}
		if yyrt114 {
			for ; yyj114 < yyl114; yyj114++ {
				yyv114 = append(yyv114, PodDisruptionBudget{})
				yyh114.ElemContainerState(yyj114)
				if r.TryDecodeAsNil() {
					yyv114[yyj114] = PodDisruptionBudget{}
				} else {
					yyv116 := &yyv114[yyj114]
					yyv116.CodecDecodeSelf(d)
				}

			}
		}

	} else {
		yyj114 := 0
		for ; !r.CheckBreak(); yyj114++ {

			if yyj114 >= len(yyv114) {
				yyv114 = append(yyv114, PodDisruptionBudget{}) // var yyz114 PodDisruptionBudget
				yyc114 = true
			}
			yyh114.ElemContainerState(yyj114)
			if yyj114 < len(yyv114) {
				if r.TryDecodeAsNil() {
					yyv114[yyj114] = PodDisruptionBudget{}
				} else {
					yyv117 := &yyv114[yyj114]
					yyv117.CodecDecodeSelf(d)
				}

			} else {
				z.DecSwallow()
			}

		}
		if yyj114 < len(yyv114) {
			yyv114 = yyv114[:yyj114]
			yyc114 = true
		} else if yyj114 == 0 && yyv114 == nil {
			yyv114 = []PodDisruptionBudget{}
			yyc114 = true
		}
	}
	yyh114.End()
	if yyc114 {
		*v = yyv114
	}
}
