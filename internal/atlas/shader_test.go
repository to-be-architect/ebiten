// Copyright 2021 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package atlas_test

import (
	"image/color"
	"testing"

	"github.com/hajimehoshi/ebiten/v2/internal/affine"
	"github.com/hajimehoshi/ebiten/v2/internal/atlas"
	"github.com/hajimehoshi/ebiten/v2/internal/graphics"
	"github.com/hajimehoshi/ebiten/v2/internal/graphicsdriver"
	etesting "github.com/hajimehoshi/ebiten/v2/internal/testing"
	"github.com/hajimehoshi/ebiten/v2/internal/ui"
)

func TestShaderFillTwice(t *testing.T) {
	const w, h = 1, 1

	dst := atlas.NewImage(w, h)

	vs := quadVertices(w, h, 0, 0, 1)
	is := graphics.QuadIndices()
	dr := graphicsdriver.Region{
		X:      0,
		Y:      0,
		Width:  w,
		Height: h,
	}
	g := ui.GraphicsDriverForTesting()
	s0 := atlas.NewShader(etesting.ShaderProgramFill(0xff, 0xff, 0xff, 0xff))
	dst.DrawTriangles([graphics.ShaderImageNum]*atlas.Image{}, vs, is, affine.ColorMIdentity{}, graphicsdriver.CompositeModeCopy, graphicsdriver.FilterNearest, graphicsdriver.AddressUnsafe, dr, graphicsdriver.Region{}, [graphics.ShaderImageNum - 1][2]float32{}, s0, nil, false)

	// Vertices must be recreated (#1755)
	vs = quadVertices(w, h, 0, 0, 1)
	s1 := atlas.NewShader(etesting.ShaderProgramFill(0x80, 0x80, 0x80, 0xff))
	dst.DrawTriangles([graphics.ShaderImageNum]*atlas.Image{}, vs, is, affine.ColorMIdentity{}, graphicsdriver.CompositeModeCopy, graphicsdriver.FilterNearest, graphicsdriver.AddressUnsafe, dr, graphicsdriver.Region{}, [graphics.ShaderImageNum - 1][2]float32{}, s1, nil, false)

	pix, err := dst.Pixels(g)
	if err != nil {
		t.Error(err)
	}
	if got, want := (color.RGBA{pix[0], pix[1], pix[2], pix[3]}), (color.RGBA{0x80, 0x80, 0x80, 0xff}); got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestImageDrawTwice(t *testing.T) {
	const w, h = 1, 1

	dst := atlas.NewImage(w, h)
	src0 := atlas.NewImage(w, h)
	src0.ReplacePixels([]byte{0xff, 0xff, 0xff, 0xff}, nil)
	src1 := atlas.NewImage(w, h)
	src1.ReplacePixels([]byte{0x80, 0x80, 0x80, 0xff}, nil)

	vs := quadVertices(w, h, 0, 0, 1)
	is := graphics.QuadIndices()
	dr := graphicsdriver.Region{
		X:      0,
		Y:      0,
		Width:  w,
		Height: h,
	}
	dst.DrawTriangles([graphics.ShaderImageNum]*atlas.Image{src0}, vs, is, affine.ColorMIdentity{}, graphicsdriver.CompositeModeCopy, graphicsdriver.FilterNearest, graphicsdriver.AddressUnsafe, dr, graphicsdriver.Region{}, [graphics.ShaderImageNum - 1][2]float32{}, nil, nil, false)

	// Vertices must be recreated (#1755)
	vs = quadVertices(w, h, 0, 0, 1)
	dst.DrawTriangles([graphics.ShaderImageNum]*atlas.Image{src1}, vs, is, affine.ColorMIdentity{}, graphicsdriver.CompositeModeCopy, graphicsdriver.FilterNearest, graphicsdriver.AddressUnsafe, dr, graphicsdriver.Region{}, [graphics.ShaderImageNum - 1][2]float32{}, nil, nil, false)

	pix, err := dst.Pixels(ui.GraphicsDriverForTesting())
	if err != nil {
		t.Error(err)
	}
	if got, want := (color.RGBA{pix[0], pix[1], pix[2], pix[3]}), (color.RGBA{0x80, 0x80, 0x80, 0xff}); got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
