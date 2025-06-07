package main

import (
	"fmt"
	"log"
	"time"
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
	lastAutoClick time.Time
}

//func (g *Game) init(){
//	g.cnt: 0
//}

func (g *Game) Update() error {
	currentSpaceState := ebiten.IsKeyPressed(ebiten.KeySpace)
	currentSState := ebiten.IsKeyPressed(ebiten.KeyS)
	currentOneState := ebiten.IsKeyPressed(ebiten.Key1)
	currentTwoState := ebiten.IsKeyPressed(ebiten.Key2)
	if currentSpaceState && !g.prevSpaceState {
		g.Score = g.Score + 1 + g.TapLevel 
		g.ScoreText = strconv.Itoa(g.Score)
		fmt.Println("Score:", g.Score)
	}
	//if currentSState && !g.prevSState {
	//	g.Score += 1
	//	g.ScoreText = strconv.Itoa(g.Score)
	//	fmt.Println("Score:", g.Score)
	//}
	g.prevSpaceState = currentSpaceState
	if currentSState{
		g.prevSState = !g.prevSState
	}
	if g.prevSState && currentOneState{
		if g.TapPrice <= g.Score{
			g.Score -= g.TapPrice
			g.ScoreText = strconv.Itoa(g.Score)
			fmt.Println("Score:", g.Score)
			g.TapPrice += 20;
			g.TapLevel += 1;
		}
	}
	if g.prevSState && currentTwoState{
		if g.TapBotPrice <= g.Score{
			g.Score -= g.TapBotPrice
			g.ScoreText = strconv.Itoa(g.Score)
			fmt.Println("Score:", g.Score)
			g.TapBotPrice += 100;
			g.TapBotLevel += 1;
		}
	}

	if g.TapBotLevel > 0 {
		now := time.Now()
		if now.Sub(g.lastAutoClick) >= 2*time.Second {
			g.Score += g.TapBotLevel
			g.ScoreText = strconv.Itoa(g.Score)
			g.lastAutoClick = now
			fmt.Println("Auto-click! Score:", g.Score)
		}
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
		ebitenutil.DebugPrintAt(screen, "(1)Tap level: " + strconv.Itoa(g.TapLevel) + "  " + strconv.Itoa(g.TapPrice) + " price", 90, 60)
		ebitenutil.DebugPrintAt(screen, "(2)Tap bot level: " + strconv.Itoa(g.TapBotLevel) + "  " + strconv.Itoa(g.TapBotPrice) + " price", 90, 100)
		ebitenutil.DebugPrintAt(screen, "(3)Present: " + strconv.Itoa(g.Present) + "/2  " + strconv.Itoa(g.PresentPrice) + " price", 90, 140)
	}
	/*if(g.TapBotLevel != 0){
		time.Sleep(2 * time.Second)
		g.Score += g.TapBotLevel
		g.ScoreText = strconv.Itoa(g.Score)
	}*/
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