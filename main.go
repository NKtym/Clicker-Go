package main

import (
	"fmt"
	"log"
	"time"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	_ "image/jpeg"
	"strconv"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	BattleScore int
	ScoreText string
	prevSpaceState bool
	prevSState bool
	prevFState bool
	WinBattleOne bool
	WinBattleTwo bool
	WinBattleThree bool
	BattleOne bool
	BattleTwo bool
	BattleThree bool
	ScreenEnd bool
	lastAutoClick time.Time
}

func (g *Game) Update() error {
	g.prevSpaceState = ebiten.IsKeyPressed(ebiten.KeySpace)
	currentEscState := inpututil.IsKeyJustPressed(ebiten.KeyEscape)
	currentSState := inpututil.IsKeyJustPressed(ebiten.KeyS)
	currentFState := inpututil.IsKeyJustPressed(ebiten.KeyF)
	currentOneState := inpututil.IsKeyJustPressed(ebiten.Key1)
	currentTwoState := inpututil.IsKeyJustPressed(ebiten.Key2)
	currentThreeState := inpututil.IsKeyJustPressed(ebiten.Key3)
	if g.BattleOne {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace){
			g.BattleScore -= 1 + g.TapLevel
			g.prevSpaceState = true
		}
		if g.BattleScore <= 0 {
			g.BattleOne = false
			g.WinBattleOne = true
			g.Score += 200
			g.ScoreText = strconv.Itoa(g.Score)
		}
		if currentEscState{
			return ebiten.Termination
		}
	} else if g.BattleTwo{
		if inpututil.IsKeyJustPressed(ebiten.KeySpace){
			g.BattleScore -= 1 + g.TapLevel
			g.prevSpaceState = true
		}
		if g.BattleScore <= 0 {
			g.BattleTwo = false
			g.WinBattleTwo = true
			g.Score += 500
			g.ScoreText = strconv.Itoa(g.Score)
		}
		if currentEscState{
			return ebiten.Termination
		}
	} else {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace){
			g.Score = g.Score + 1 + g.TapLevel 
			g.ScoreText = strconv.Itoa(g.Score)
			fmt.Println("Score:", g.Score)
			g.prevSpaceState = true
		}
		if currentEscState{
			return ebiten.Termination
		}
		if currentSState{
			g.prevSState = !g.prevSState
		}
		if currentFState{
			g.prevFState = !g.prevFState
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
		if g.prevFState && currentOneState && !g.WinBattleOne {
			g.BattleOne = true
			g.BattleScore = 200
		}
		if g.prevFState && currentTwoState && !g.WinBattleTwo {
			g.BattleTwo = true
			g.BattleScore = 500
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
	if(g.ScreenEnd){
		time.Sleep(5 * time.Second)
		g.ScreenEnd = false
	}
	if(g.BattleOne){
		screen.Fill(color.Black)
		if(g.prevSpaceState){
			geoM := ebiten.GeoM{}
			geoM.Translate(float64(screenWidth/2-300), float64(screenHeight/2-140))
			geoM.Scale(0.198, 0.198)
			logo, _, err := ebitenutil.NewImageFromFile("images/Boss1_2.jpg")
			if err != nil {
				log.Fatal(err)
			}
			op := &ebiten.DrawImageOptions{GeoM: geoM}
			screen.DrawImage(logo, op)
			g.prevSpaceState = false
		} else {
			geoM := ebiten.GeoM{}
			geoM.Translate(float64(screenWidth/2-300), float64(screenHeight/2-140))
			geoM.Scale(0.2, 0.2)
			logo, _, err := ebitenutil.NewImageFromFile("images/Boss1_1.jpg")
			if err != nil {
				log.Fatal(err)
			}
			op := &ebiten.DrawImageOptions{GeoM: geoM}
			screen.DrawImage(logo, op)
		}
		healthText := "BOSS HP: " + strconv.Itoa(g.BattleScore)
		ebitenutil.DebugPrintAt(screen, healthText, 132, 40)
	} else if g.BattleTwo {
		screen.Fill(color.Black)
		if(g.prevSpaceState){
			geoM := ebiten.GeoM{}
			geoM.Translate(float64(screenWidth/2-35), float64(screenHeight/2-92))
			geoM.Scale(0.198, 0.198)
			logo, _, err := ebitenutil.NewImageFromFile("images/Boss2_2.jpg")
			if err != nil {
				log.Fatal(err)
			}
			op := &ebiten.DrawImageOptions{GeoM: geoM}
			screen.DrawImage(logo, op)
			g.prevSpaceState = false
		} else {
			geoM := ebiten.GeoM{}
			geoM.Translate(float64(screenWidth/2-35), float64(screenHeight/2-92))
			geoM.Scale(0.2, 0.2)
			logo, _, err := ebitenutil.NewImageFromFile("images/Boss2_1.jpg")
			if err != nil {
				log.Fatal(err)
			}
			op := &ebiten.DrawImageOptions{GeoM: geoM}
			screen.DrawImage(logo, op)
		}
		healthText := "BOSS HP: " + strconv.Itoa(g.BattleScore)
		ebitenutil.DebugPrintAt(screen, healthText, 132, 50)
	} else {
		if(g.prevSpaceState){
			geoM := ebiten.GeoM{}
			geoM.Translate(float64(screenWidth/2-50), float64(screenHeight/2-150))
			geoM.Scale(0.198, 0.198)
			logo, _, err := ebitenutil.NewImageFromFile("images/2224.jpg")
			if err != nil {
				log.Fatal(err)
			}
			op := &ebiten.DrawImageOptions{GeoM: geoM}
			screen.DrawImage(logo, op)
			g.prevSpaceState = false
		} else{
			geoM := ebiten.GeoM{}
			geoM.Translate(float64(screenWidth/2-50), float64(screenHeight/2-150))
			geoM.Scale(0.2, 0.2)
			logo, _, err := ebitenutil.NewImageFromFile("images/2224.jpg")
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
		if(g.prevFState){
			screen.Fill(color.Black)
			ebitenutil.DebugPrintAt(screen, "(1)Boss 1", 130, 70)
			ebitenutil.DebugPrintAt(screen, "(2)Boss 2", 130, 110)
			ebitenutil.DebugPrintAt(screen, "(3)Boss 3", 130, 150)
		}
	}
	if(g.WinBattleOne){
		screen.Fill(color.Black)
		ebitenutil.DebugPrintAt(screen, "You win!!!", 135, 40)
		geoM := ebiten.GeoM{}
		geoM.Translate(float64(screenWidth/2-300), float64(screenHeight/2-140))
		geoM.Scale(0.2, 0.2)
		logo, _, err := ebitenutil.NewImageFromFile("images/Boss1_3.jpg")
		if err != nil {
			log.Fatal(err)
		}
		op := &ebiten.DrawImageOptions{GeoM: geoM}
		screen.DrawImage(logo, op)
		g.WinBattleOne = false
		g.ScreenEnd = true
	} else if (g.WinBattleTwo){
		screen.Fill(color.Black)
		ebitenutil.DebugPrintAt(screen, "You win!!!", 135, 50)
		geoM := ebiten.GeoM{}
		geoM.Translate(float64(screenWidth/2-35), float64(screenHeight/2-92))
		geoM.Scale(0.2, 0.2)
		logo, _, err := ebitenutil.NewImageFromFile("images/Boss2_3.jpg")
		if err != nil {
			log.Fatal(err)
		}
		op := &ebiten.DrawImageOptions{GeoM: geoM}
		screen.DrawImage(logo, op)
		g.WinBattleTwo = false
		g.ScreenEnd = true
	}
	ebitenutil.DebugPrint(screen, "Score: " + g.ScoreText)
	ebitenutil.DebugPrintAt(screen, "Shop(S)", 276, 1)
	ebitenutil.DebugPrintAt(screen, "Exit(Esc)", 266, 225)
	ebitenutil.DebugPrintAt(screen, "Battle(F)", 0, 225)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Test screan")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}