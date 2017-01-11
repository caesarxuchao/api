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

package v1beta1

import (
	"errors"
	"fmt"
	codec1978 "github.com/ugorji/go/codec"
	pkg1_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	pkg3_types "k8s.io/apimachinery/pkg/types"
	pkg2_v1 "k8s.io/kubernetes/pkg/api/v1"
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
		var v0 pkg1_v1.TypeMeta
		var v1 pkg3_types.UID
		var v2 pkg2_v1.ObjectMeta
		var v3 time.Time
		_, _, _, _ = v0, v1, v2, v3
	}
}

func (x *TokenReview) CodecEncodeSelf(e *codec1978.Encoder) {
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
			var yyq2 [5]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			yyq2[0] = x.Kind != ""
			yyq2[1] = x.APIVersion != ""
			yyq2[2] = true
			yyq2[4] = true
			var yynn2 int
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(5)
			} else {
				yynn2 = 1
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
					yym4 := z.EncBinary()
					_ = yym4
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("kind"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym5 := z.EncBinary()
					_ = yym5
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[1] {
					yym7 := z.EncBinary()
					_ = yym7
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("apiVersion"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym8 := z.EncBinary()
					_ = yym8
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[2] {
					yy10 := &x.ObjectMeta
					yy10.CodecEncodeSelf(e)
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[2] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("metadata"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy12 := &x.ObjectMeta
					yy12.CodecEncodeSelf(e)
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				yy15 := &x.Spec
				yy15.CodecEncodeSelf(e)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("spec"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				yy17 := &x.Spec
				yy17.CodecEncodeSelf(e)
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[4] {
					yy20 := &x.Status
					yy20.CodecEncodeSelf(e)
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[4] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("status"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy22 := &x.Status
					yy22.CodecEncodeSelf(e)
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

func (x *TokenReview) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap1234 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray1234 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *TokenReview) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys3Slc = r.DecodeBytes(yys3Slc, true, true)
		yys3 := string(yys3Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys3 {
		case "kind":
			if r.TryDecodeAsNil() {
				x.Kind = ""
			} else {
				yyv4 := &x.Kind
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else {
					*((*string)(yyv4)) = r.DecodeString()
				}
			}
		case "apiVersion":
			if r.TryDecodeAsNil() {
				x.APIVersion = ""
			} else {
				yyv6 := &x.APIVersion
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} else {
					*((*string)(yyv6)) = r.DecodeString()
				}
			}
		case "metadata":
			if r.TryDecodeAsNil() {
				x.ObjectMeta = pkg2_v1.ObjectMeta{}
			} else {
				yyv8 := &x.ObjectMeta
				yyv8.CodecDecodeSelf(d)
			}
		case "spec":
			if r.TryDecodeAsNil() {
				x.Spec = TokenReviewSpec{}
			} else {
				yyv9 := &x.Spec
				yyv9.CodecDecodeSelf(d)
			}
		case "status":
			if r.TryDecodeAsNil() {
				x.Status = TokenReviewStatus{}
			} else {
				yyv10 := &x.Status
				yyv10.CodecDecodeSelf(d)
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *TokenReview) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj11 int
	var yyb11 bool
	var yyhl11 bool = l >= 0
	yyj11++
	if yyhl11 {
		yyb11 = yyj11 > l
	} else {
		yyb11 = r.CheckBreak()
	}
	if yyb11 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} else {
		yyv12 := &x.Kind
		yym13 := z.DecBinary()
		_ = yym13
		if false {
		} else {
			*((*string)(yyv12)) = r.DecodeString()
		}
	}
	yyj11++
	if yyhl11 {
		yyb11 = yyj11 > l
	} else {
		yyb11 = r.CheckBreak()
	}
	if yyb11 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} else {
		yyv14 := &x.APIVersion
		yym15 := z.DecBinary()
		_ = yym15
		if false {
		} else {
			*((*string)(yyv14)) = r.DecodeString()
		}
	}
	yyj11++
	if yyhl11 {
		yyb11 = yyj11 > l
	} else {
		yyb11 = r.CheckBreak()
	}
	if yyb11 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.ObjectMeta = pkg2_v1.ObjectMeta{}
	} else {
		yyv16 := &x.ObjectMeta
		yyv16.CodecDecodeSelf(d)
	}
	yyj11++
	if yyhl11 {
		yyb11 = yyj11 > l
	} else {
		yyb11 = r.CheckBreak()
	}
	if yyb11 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Spec = TokenReviewSpec{}
	} else {
		yyv17 := &x.Spec
		yyv17.CodecDecodeSelf(d)
	}
	yyj11++
	if yyhl11 {
		yyb11 = yyj11 > l
	} else {
		yyb11 = r.CheckBreak()
	}
	if yyb11 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Status = TokenReviewStatus{}
	} else {
		yyv18 := &x.Status
		yyv18.CodecDecodeSelf(d)
	}
	for {
		yyj11++
		if yyhl11 {
			yyb11 = yyj11 > l
		} else {
			yyb11 = r.CheckBreak()
		}
		if yyb11 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj11-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *TokenReviewSpec) CodecEncodeSelf(e *codec1978.Encoder) {
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
			var yyq2 [1]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			yyq2[0] = x.Token != ""
			var yynn2 int
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(1)
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
					yym4 := z.EncBinary()
					_ = yym4
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Token))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("token"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym5 := z.EncBinary()
					_ = yym5
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Token))
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

func (x *TokenReviewSpec) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap1234 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray1234 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *TokenReviewSpec) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys3Slc = r.DecodeBytes(yys3Slc, true, true)
		yys3 := string(yys3Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys3 {
		case "token":
			if r.TryDecodeAsNil() {
				x.Token = ""
			} else {
				yyv4 := &x.Token
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else {
					*((*string)(yyv4)) = r.DecodeString()
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *TokenReviewSpec) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj6 int
	var yyb6 bool
	var yyhl6 bool = l >= 0
	yyj6++
	if yyhl6 {
		yyb6 = yyj6 > l
	} else {
		yyb6 = r.CheckBreak()
	}
	if yyb6 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Token = ""
	} else {
		yyv7 := &x.Token
		yym8 := z.DecBinary()
		_ = yym8
		if false {
		} else {
			*((*string)(yyv7)) = r.DecodeString()
		}
	}
	for {
		yyj6++
		if yyhl6 {
			yyb6 = yyj6 > l
		} else {
			yyb6 = r.CheckBreak()
		}
		if yyb6 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj6-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *TokenReviewStatus) CodecEncodeSelf(e *codec1978.Encoder) {
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
			var yyq2 [3]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			yyq2[0] = x.Authenticated != false
			yyq2[1] = true
			yyq2[2] = x.Error != ""
			var yynn2 int
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(3)
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
					yym4 := z.EncBinary()
					_ = yym4
					if false {
					} else {
						r.EncodeBool(bool(x.Authenticated))
					}
				} else {
					r.EncodeBool(false)
				}
			} else {
				if yyq2[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("authenticated"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym5 := z.EncBinary()
					_ = yym5
					if false {
					} else {
						r.EncodeBool(bool(x.Authenticated))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[1] {
					yy7 := &x.User
					yy7.CodecEncodeSelf(e)
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("user"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy9 := &x.User
					yy9.CodecEncodeSelf(e)
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[2] {
					yym12 := z.EncBinary()
					_ = yym12
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Error))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[2] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("error"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym13 := z.EncBinary()
					_ = yym13
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Error))
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

func (x *TokenReviewStatus) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap1234 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray1234 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *TokenReviewStatus) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys3Slc = r.DecodeBytes(yys3Slc, true, true)
		yys3 := string(yys3Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys3 {
		case "authenticated":
			if r.TryDecodeAsNil() {
				x.Authenticated = false
			} else {
				yyv4 := &x.Authenticated
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else {
					*((*bool)(yyv4)) = r.DecodeBool()
				}
			}
		case "user":
			if r.TryDecodeAsNil() {
				x.User = UserInfo{}
			} else {
				yyv6 := &x.User
				yyv6.CodecDecodeSelf(d)
			}
		case "error":
			if r.TryDecodeAsNil() {
				x.Error = ""
			} else {
				yyv7 := &x.Error
				yym8 := z.DecBinary()
				_ = yym8
				if false {
				} else {
					*((*string)(yyv7)) = r.DecodeString()
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *TokenReviewStatus) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj9 int
	var yyb9 bool
	var yyhl9 bool = l >= 0
	yyj9++
	if yyhl9 {
		yyb9 = yyj9 > l
	} else {
		yyb9 = r.CheckBreak()
	}
	if yyb9 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Authenticated = false
	} else {
		yyv10 := &x.Authenticated
		yym11 := z.DecBinary()
		_ = yym11
		if false {
		} else {
			*((*bool)(yyv10)) = r.DecodeBool()
		}
	}
	yyj9++
	if yyhl9 {
		yyb9 = yyj9 > l
	} else {
		yyb9 = r.CheckBreak()
	}
	if yyb9 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.User = UserInfo{}
	} else {
		yyv12 := &x.User
		yyv12.CodecDecodeSelf(d)
	}
	yyj9++
	if yyhl9 {
		yyb9 = yyj9 > l
	} else {
		yyb9 = r.CheckBreak()
	}
	if yyb9 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Error = ""
	} else {
		yyv13 := &x.Error
		yym14 := z.DecBinary()
		_ = yym14
		if false {
		} else {
			*((*string)(yyv13)) = r.DecodeString()
		}
	}
	for {
		yyj9++
		if yyhl9 {
			yyb9 = yyj9 > l
		} else {
			yyb9 = r.CheckBreak()
		}
		if yyb9 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj9-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *UserInfo) CodecEncodeSelf(e *codec1978.Encoder) {
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
			var yyq2 [4]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			yyq2[0] = x.Username != ""
			yyq2[1] = x.UID != ""
			yyq2[2] = len(x.Groups) != 0
			yyq2[3] = len(x.Extra) != 0
			var yynn2 int
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(4)
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
					yym4 := z.EncBinary()
					_ = yym4
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Username))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("username"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym5 := z.EncBinary()
					_ = yym5
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Username))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[1] {
					yym7 := z.EncBinary()
					_ = yym7
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.UID))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("uid"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym8 := z.EncBinary()
					_ = yym8
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.UID))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[2] {
					if x.Groups == nil {
						r.EncodeNil()
					} else {
						yym10 := z.EncBinary()
						_ = yym10
						if false {
						} else {
							z.F.EncSliceStringV(x.Groups, false, e)
						}
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[2] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("groups"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					if x.Groups == nil {
						r.EncodeNil()
					} else {
						yym11 := z.EncBinary()
						_ = yym11
						if false {
						} else {
							z.F.EncSliceStringV(x.Groups, false, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[3] {
					if x.Extra == nil {
						r.EncodeNil()
					} else {
						yym13 := z.EncBinary()
						_ = yym13
						if false {
						} else {
							h.encMapstringExtraValue((map[string]ExtraValue)(x.Extra), e)
						}
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[3] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("extra"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					if x.Extra == nil {
						r.EncodeNil()
					} else {
						yym14 := z.EncBinary()
						_ = yym14
						if false {
						} else {
							h.encMapstringExtraValue((map[string]ExtraValue)(x.Extra), e)
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

func (x *UserInfo) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap1234 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray1234 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *UserInfo) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys3Slc = r.DecodeBytes(yys3Slc, true, true)
		yys3 := string(yys3Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys3 {
		case "username":
			if r.TryDecodeAsNil() {
				x.Username = ""
			} else {
				yyv4 := &x.Username
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else {
					*((*string)(yyv4)) = r.DecodeString()
				}
			}
		case "uid":
			if r.TryDecodeAsNil() {
				x.UID = ""
			} else {
				yyv6 := &x.UID
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} else {
					*((*string)(yyv6)) = r.DecodeString()
				}
			}
		case "groups":
			if r.TryDecodeAsNil() {
				x.Groups = nil
			} else {
				yyv8 := &x.Groups
				yym9 := z.DecBinary()
				_ = yym9
				if false {
				} else {
					z.F.DecSliceStringX(yyv8, false, d)
				}
			}
		case "extra":
			if r.TryDecodeAsNil() {
				x.Extra = nil
			} else {
				yyv10 := &x.Extra
				yym11 := z.DecBinary()
				_ = yym11
				if false {
				} else {
					h.decMapstringExtraValue((*map[string]ExtraValue)(yyv10), d)
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *UserInfo) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj12 int
	var yyb12 bool
	var yyhl12 bool = l >= 0
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Username = ""
	} else {
		yyv13 := &x.Username
		yym14 := z.DecBinary()
		_ = yym14
		if false {
		} else {
			*((*string)(yyv13)) = r.DecodeString()
		}
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.UID = ""
	} else {
		yyv15 := &x.UID
		yym16 := z.DecBinary()
		_ = yym16
		if false {
		} else {
			*((*string)(yyv15)) = r.DecodeString()
		}
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Groups = nil
	} else {
		yyv17 := &x.Groups
		yym18 := z.DecBinary()
		_ = yym18
		if false {
		} else {
			z.F.DecSliceStringX(yyv17, false, d)
		}
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Extra = nil
	} else {
		yyv19 := &x.Extra
		yym20 := z.DecBinary()
		_ = yym20
		if false {
		} else {
			h.decMapstringExtraValue((*map[string]ExtraValue)(yyv19), d)
		}
	}
	for {
		yyj12++
		if yyhl12 {
			yyb12 = yyj12 > l
		} else {
			yyb12 = r.CheckBreak()
		}
		if yyb12 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj12-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x ExtraValue) CodecEncodeSelf(e *codec1978.Encoder) {
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
			h.encExtraValue((ExtraValue)(x), e)
		}
	}
}

func (x *ExtraValue) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		h.decExtraValue((*ExtraValue)(x), d)
	}
}

func (x codecSelfer1234) encMapstringExtraValue(v map[string]ExtraValue, e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.EncodeMapStart(len(v))
	for yyk1, yyv1 := range v {
		z.EncSendContainerState(codecSelfer_containerMapKey1234)
		yym2 := z.EncBinary()
		_ = yym2
		if false {
		} else {
			r.EncodeString(codecSelferC_UTF81234, string(yyk1))
		}
		z.EncSendContainerState(codecSelfer_containerMapValue1234)
		if yyv1 == nil {
			r.EncodeNil()
		} else {
			yyv1.CodecEncodeSelf(e)
		}
	}
	z.EncSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x codecSelfer1234) decMapstringExtraValue(v *map[string]ExtraValue, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv1 := *v
	yyl1 := r.ReadMapStart()
	yybh1 := z.DecBasicHandle()
	if yyv1 == nil {
		yyrl1, _ := z.DecInferLen(yyl1, yybh1.MaxInitLen, 40)
		yyv1 = make(map[string]ExtraValue, yyrl1)
		*v = yyv1
	}
	var yymk1 string
	var yymv1 ExtraValue
	var yymg1 bool
	if yybh1.MapValueReset {
		yymg1 = true
	}
	if yyl1 > 0 {
		for yyj1 := 0; yyj1 < yyl1; yyj1++ {
			z.DecSendContainerState(codecSelfer_containerMapKey1234)
			if r.TryDecodeAsNil() {
				yymk1 = ""
			} else {
				yyv2 := &yymk1
				yym3 := z.DecBinary()
				_ = yym3
				if false {
				} else {
					*((*string)(yyv2)) = r.DecodeString()
				}
			}

			if yymg1 {
				yymv1 = yyv1[yymk1]
			} else {
				yymv1 = nil
			}
			z.DecSendContainerState(codecSelfer_containerMapValue1234)
			if r.TryDecodeAsNil() {
				yymv1 = nil
			} else {
				yyv4 := &yymv1
				yyv4.CodecDecodeSelf(d)
			}

			if yyv1 != nil {
				yyv1[yymk1] = yymv1
			}
		}
	} else if yyl1 < 0 {
		for yyj1 := 0; !r.CheckBreak(); yyj1++ {
			z.DecSendContainerState(codecSelfer_containerMapKey1234)
			if r.TryDecodeAsNil() {
				yymk1 = ""
			} else {
				yyv5 := &yymk1
				yym6 := z.DecBinary()
				_ = yym6
				if false {
				} else {
					*((*string)(yyv5)) = r.DecodeString()
				}
			}

			if yymg1 {
				yymv1 = yyv1[yymk1]
			} else {
				yymv1 = nil
			}
			z.DecSendContainerState(codecSelfer_containerMapValue1234)
			if r.TryDecodeAsNil() {
				yymv1 = nil
			} else {
				yyv7 := &yymv1
				yyv7.CodecDecodeSelf(d)
			}

			if yyv1 != nil {
				yyv1[yymk1] = yymv1
			}
		}
	} // else len==0: TODO: Should we clear map entries?
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x codecSelfer1234) encExtraValue(v ExtraValue, e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.EncodeArrayStart(len(v))
	for _, yyv1 := range v {
		z.EncSendContainerState(codecSelfer_containerArrayElem1234)
		yym2 := z.EncBinary()
		_ = yym2
		if false {
		} else {
			r.EncodeString(codecSelferC_UTF81234, string(yyv1))
		}
	}
	z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x codecSelfer1234) decExtraValue(v *ExtraValue, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv1 := *v
	yyh1, yyl1 := z.DecSliceHelperStart()
	var yyc1 bool
	_ = yyc1
	if yyl1 == 0 {
		if yyv1 == nil {
			yyv1 = []string{}
			yyc1 = true
		} else if len(yyv1) != 0 {
			yyv1 = yyv1[:0]
			yyc1 = true
		}
	} else if yyl1 > 0 {
		var yyrr1, yyrl1 int
		var yyrt1 bool
		_, _ = yyrl1, yyrt1
		yyrr1 = yyl1 // len(yyv1)
		if yyl1 > cap(yyv1) {

			yyrl1, yyrt1 = z.DecInferLen(yyl1, z.DecBasicHandle().MaxInitLen, 16)
			if yyrt1 {
				if yyrl1 <= cap(yyv1) {
					yyv1 = yyv1[:yyrl1]
				} else {
					yyv1 = make([]string, yyrl1)
				}
			} else {
				yyv1 = make([]string, yyrl1)
			}
			yyc1 = true
			yyrr1 = len(yyv1)
		} else if yyl1 != len(yyv1) {
			yyv1 = yyv1[:yyl1]
			yyc1 = true
		}
		yyj1 := 0
		for ; yyj1 < yyrr1; yyj1++ {
			yyh1.ElemContainerState(yyj1)
			if r.TryDecodeAsNil() {
				yyv1[yyj1] = ""
			} else {
				yyv2 := &yyv1[yyj1]
				yym3 := z.DecBinary()
				_ = yym3
				if false {
				} else {
					*((*string)(yyv2)) = r.DecodeString()
				}
			}

		}
		if yyrt1 {
			for ; yyj1 < yyl1; yyj1++ {
				yyv1 = append(yyv1, "")
				yyh1.ElemContainerState(yyj1)
				if r.TryDecodeAsNil() {
					yyv1[yyj1] = ""
				} else {
					yyv4 := &yyv1[yyj1]
					yym5 := z.DecBinary()
					_ = yym5
					if false {
					} else {
						*((*string)(yyv4)) = r.DecodeString()
					}
				}

			}
		}

	} else {
		yyj1 := 0
		for ; !r.CheckBreak(); yyj1++ {

			if yyj1 >= len(yyv1) {
				yyv1 = append(yyv1, "") // var yyz1 string
				yyc1 = true
			}
			yyh1.ElemContainerState(yyj1)
			if yyj1 < len(yyv1) {
				if r.TryDecodeAsNil() {
					yyv1[yyj1] = ""
				} else {
					yyv6 := &yyv1[yyj1]
					yym7 := z.DecBinary()
					_ = yym7
					if false {
					} else {
						*((*string)(yyv6)) = r.DecodeString()
					}
				}

			} else {
				z.DecSwallow()
			}

		}
		if yyj1 < len(yyv1) {
			yyv1 = yyv1[:yyj1]
			yyc1 = true
		} else if yyj1 == 0 && yyv1 == nil {
			yyv1 = []string{}
			yyc1 = true
		}
	}
	yyh1.End()
	if yyc1 {
		*v = yyv1
	}
}
