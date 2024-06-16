package assets

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"image"
	_ "image/png"
	"io/fs"
	"path"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type SpriteSlice struct {
	Vertical   []int `json:"vertical"`
	Horizontal []int `json:"horizontal"`
}

type SpriteData struct {
	Slices *SpriteSlice `json:"slices"`
}

type Assets struct {
	spriteData   map[string]*SpriteData
	sprites      map[string]*ebiten.Image
	spriteSlices map[string]*ebiten.Image
	fonts        map[string]*font.Face
}

func (asset *Assets) GetSprite(dir string) *ebiten.Image {
	return asset.sprites[dir+".png"]
}

func (asset *Assets) GetFont(dir string) *font.Face {
	return asset.fonts[dir+".ttf"]
}

func (asset *Assets) GetSpriteData(dir string) *SpriteData {
	return asset.spriteData[dir+".json"]
}

func (asset *Assets) GetSpriteSlices(dir string, x int, y int) *ebiten.Image {
	sliceDir := fmt.Sprintf("%v[%v,%v]", dir, x, y)
	return asset.spriteSlices[sliceDir]
}

//go:embed embed/*
var embedDir embed.FS

func getAllFilenames(efs *embed.FS) ([]string, error) {
	files := []string{}

	error := fs.WalkDir(efs, ".", func(path string, dir fs.DirEntry, err error) error {
		if dir.IsDir() {
			return nil
		}

		files = append(files, path)
		return nil
	})

	if error != nil {
		return nil, error
	}

	return files, nil
}

func Load() *Assets {
	assets := Assets{
		spriteData:   make(map[string]*SpriteData),
		sprites:      make(map[string]*ebiten.Image),
		spriteSlices: make(map[string]*ebiten.Image),
		fonts:        make(map[string]*font.Face),
	}
	entries, error := getAllFilenames(&embedDir)

	if error != nil {
		panic(fmt.Sprintf("Unexpected error while searching assets: %v", error))
	}

	// load assets

	for _, assetDir := range entries {
		assetBytes, error := embedDir.ReadFile(assetDir)

		if error != nil {
			panic(fmt.Sprintf("Unexpected error while reading asset bytes: %v, dir %v", error, assetDir))
		}

		switch extension := path.Ext(assetDir); extension {
		case ".png":
			im, _, error := image.Decode(bytes.NewReader(assetBytes))
			if error != nil {
				panic(fmt.Sprintf("Unexpected error while decoding an image: %v, dir %v", error, assetDir))
			}

			assets.sprites[assetDir] = ebiten.NewImageFromImage(im)

		case ".ttf":
			tt, error := opentype.Parse(assetBytes)

			if error != nil {
				panic(fmt.Sprintf("Unexpected error while parsing font: %v, dir %v", error, assetDir))
			}

			font, error := opentype.NewFace(tt, &opentype.FaceOptions{
				Size:    16,
				DPI:     72,
				Hinting: font.HintingFull,
			})

			if error != nil {
				panic(fmt.Sprintf("Unexpected error while creating font: %v, dir %v", error, assetDir))
			}

			assets.fonts[assetDir] = &font

		case ".json":
			var spriteData SpriteData
			error := json.Unmarshal(assetBytes, &spriteData)

			if error != nil {
				panic(fmt.Sprintf("Error while parsing json: %v, dir %v", error, assetDir))
			}

			assets.spriteData[assetDir] = &spriteData

		default:
			panic(fmt.Sprintf("Unexpected asset type %v, dir %v", error, assetDir))
		}

		fmt.Printf("Stored asset, %v \n", assetDir)
	}

	// process assets

	// process image slices
	for spriteDataDir, spriteData := range assets.spriteData {
		spriteDir := strings.Replace(spriteDataDir, ".json", "", 1)
		sprite := assets.GetSprite(spriteDir)
		if sprite == nil {
			panic(fmt.Sprintf("Unexpected error, the sprite data %v has not sprite %v", spriteDataDir, spriteDir))
		}

		size := sprite.Bounds().Size()

		sliceHorizontal := make([]int, len(spriteData.Slices.Horizontal)+1)
		copy(sliceHorizontal, spriteData.Slices.Horizontal)
		sliceHorizontal[len(sliceHorizontal)-1] = size.X

		sliceVertical := make([]int, len(spriteData.Slices.Vertical)+1)
		copy(sliceVertical, spriteData.Slices.Vertical)
		sliceVertical[len(sliceVertical)-1] = size.Y

		x1 := 0
		for i1, x2 := range sliceHorizontal {
			y1 := 0
			for i2, y2 := range sliceVertical {
				sliceDir := fmt.Sprintf("%v[%v,%v]", spriteDir, i1, i2)

				println(x1, x2, y1, y2)

				sliceImage := sprite.SubImage(image.Rectangle{
					Min: image.Point{
						X: x1,
						Y: y1,
					},
					Max: image.Point{
						X: x2,
						Y: y2,
					},
				})

				assets.spriteSlices[sliceDir] = ebiten.NewImageFromImage(sliceImage)
				fmt.Printf("Stored asset, slice, %v \n", sliceDir)

				y1 = y2 + 1
			}
			x1 = x2 + 1
		}
	}

	return &assets
}
