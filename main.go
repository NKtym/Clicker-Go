package main

import (
	"fmt"
	"log"
	//"time"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	//"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	_ "image/jpeg"
	"strconv"
	//"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	screenWidth  = 1200
	screenHeight = 900 
)

type Game struct{
	Score int
	TapLevel int
	TapPrice int
	TapBotLevel int
	TapBotPrice int
	Present int
	PresentPrice int
	ScoreText string
	prevSpaceState bool
	prevSState bool
}

//func (g *Game) init(){
//	g.cnt: 0
//}

func (g *Game) Update() error {
	currentSpaceState := ebiten.IsKeyPressed(ebiten.KeySpace)
	if currentSpaceState && !g.prevSpaceState {
		g.Score += 1
		g.ScoreText = strconv.Itoa(g.Score)
		fmt.Println("Score:", g.Score)
	}
	currentSState := ebiten.IsKeyPressed(ebiten.KeyS)
	//if currentSState && !g.prevSState {
	//	g.Score += 1
	//	g.ScoreText = strconv.Itoa(g.Score)
	//	fmt.Println("Score:", g.Score)
	//}
	g.prevSpaceState = currentSpaceState
	if currentSState{
		g.prevSState = !g.prevSState
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//screen.Fill(color.RGBA{})
	//screen.DrowImage("222.jpg")
	if(g.prevSpaceState){
		geoM := ebiten.GeoM{}
		geoM.Translate(float64(screenWidth/2-50), float64(screenHeight/2-150))
		geoM.Scale(0.198, 0.198)
		logo, _, err := ebitenutil.NewImageFromFile("2224.jpg")
		if err != nil {
			log.Fatal(err)
		}
		op := &ebiten.DrawImageOptions{GeoM: geoM}
		screen.DrawImage(logo, op)
	} else{
		geoM := ebiten.GeoM{}
		geoM.Translate(float64(screenWidth/2-50), float64(screenHeight/2-150))
		geoM.Scale(0.2, 0.2)
		logo, _, err := ebitenutil.NewImageFromFile("2224.jpg")
		if err != nil {
			log.Fatal(err)
		}
		op := &ebiten.DrawImageOptions{GeoM: geoM}
		screen.DrawImage(logo, op)
	}
	if(g.prevSState){
		screen.Fill(color.Black)
		ebitenutil.DebugPrintAt(screen, "Tap level: " + strconv.Itoa(g.TapLevel) + "  " + strconv.Itoa(g.TapPrice) + " price", 90, 60)
		ebitenutil.DebugPrintAt(screen, "Tap bot level: " + strconv.Itoa(g.TapBotLevel) + "  " + strconv.Itoa(g.TapBotPrice) + " price", 90, 100)
		ebitenutil.DebugPrintAt(screen, "Present: " + strconv.Itoa(g.Present) + "/2  " + strconv.Itoa(g.PresentPrice) + " price", 90, 140)
	}
	ebitenutil.DebugPrint(screen, "Score: " + g.ScoreText)
	ebitenutil.DebugPrintAt(screen, "Shop(s)", 276, 1)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	//init()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Test screan")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}