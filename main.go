package main

import (
	"fmt"
	"log"
	"time"
	"os"
	"crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "crypto/sha256"
    "encoding/base64"
    "io"
	"strings" 
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
	Click int
	CntBossWin int
	PointsEarned int
	PointsSpent int
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
	prevTabState bool
	prevKState bool
	prevLState bool
	prevYState bool
	prevNState bool
	prevConfirm1State bool
	prevConfirm2State bool
	prevConfirm3State bool
	WinBattleOne bool
	WinBattleTwo bool
	WinBattleThree bool
	BattleOne bool
	BattleTwo bool
	BattleThree bool
	ScreenEnd bool
	lastAutoClick time.Time
}

func createHash(key string) []byte {
    hasher := sha256.New()
    hasher.Write([]byte(key))
    return hasher.Sum(nil)
}

func encrypt(data []byte, passphrase string) ([]byte, error) {
    key := createHash(passphrase)
    block, _ := aes.NewCipher(key)
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }
    nonce := make([]byte, gcm.NonceSize())
    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, err
    }
    ciphertext := gcm.Seal(nonce, nonce, data, nil)
    return ciphertext, nil
}

func decrypt(data []byte, passphrase string) ([]byte, error) {
    key := createHash(passphrase)
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }
    
    nonceSize := gcm.NonceSize()
    nonce, ciphertext := data[:nonceSize], data[nonceSize:]
    
    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return nil, err
    }
    
    return plaintext, nil
}

func (g *Game) Update() error {
	g.prevSpaceState = ebiten.IsKeyPressed(ebiten.KeySpace)
	currentEscState := inpututil.IsKeyJustPressed(ebiten.KeyEscape)
	currentSState := inpututil.IsKeyJustPressed(ebiten.KeyS)
	currentFState := inpututil.IsKeyJustPressed(ebiten.KeyF)
	currentKState := inpututil.IsKeyJustPressed(ebiten.KeyK)
	currentLState := inpututil.IsKeyJustPressed(ebiten.KeyL)
	currentYState := inpututil.IsKeyJustPressed(ebiten.KeyY)
	currentNState := inpututil.IsKeyJustPressed(ebiten.KeyN)
	currentOneState := inpututil.IsKeyJustPressed(ebiten.Key1)
	currentTwoState := inpututil.IsKeyJustPressed(ebiten.Key2)
	currentThreeState := inpututil.IsKeyJustPressed(ebiten.Key3)
	if inpututil.IsKeyJustPressed(ebiten.KeyTab){
		g.prevTabState = !g.prevTabState
	}
	if g.BattleOne {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace){
			g.BattleScore -= 1 + g.TapLevel
			g.prevSpaceState = true
			g.Click++
		}
		if g.BattleScore <= 0 {
			g.BattleOne = false
			g.WinBattleOne = true
			g.Score += 200
			g.ScoreText = strconv.Itoa(g.Score)
			g.CntBossWin++
			g.PointsEarned += 200
		}
		if currentEscState{
			return ebiten.Termination
		}
	} else if g.BattleTwo{
		if inpututil.IsKeyJustPressed(ebiten.KeySpace){
			g.BattleScore -= 1 + g.TapLevel
			g.prevSpaceState = true
			g.Click++
		}
		if g.BattleScore <= 0 {
			g.BattleTwo = false
			g.WinBattleTwo = true
			g.Score += 500
			g.ScoreText = strconv.Itoa(g.Score)
			g.CntBossWin++
			g.PointsEarned += 500
		}
		if currentEscState{
			return ebiten.Termination
		}
	} else if g.BattleThree{
		if inpututil.IsKeyJustPressed(ebiten.KeySpace){
			g.BattleScore -= 1 + g.TapLevel
			g.prevSpaceState = true
			g.Click++
		}
		if g.BattleScore <= 0 {
			g.BattleThree = false
			g.WinBattleThree = true
			g.Score += 1000
			g.ScoreText = strconv.Itoa(g.Score)
			g.CntBossWin++
			g.PointsEarned += 1000
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
			g.Click++
			g.PointsEarned += 1 + g.TapLevel
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
				g.PointsSpent += g.TapPrice
			}
		}
		if g.prevSState && currentTwoState{
			if g.TapBotPrice <= g.Score{
				g.Score -= g.TapBotPrice
				g.ScoreText = strconv.Itoa(g.Score)
				fmt.Println("Score:", g.Score)
				g.TapBotPrice += 100;
				g.TapBotLevel += 1;
				g.PointsSpent += g.TapBotPrice
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
		if g.prevFState && currentThreeState && !g.WinBattleThree {
			g.BattleThree = true
			g.BattleScore = 1000
		}
		if currentKState{
			g.prevKState = !g.prevKState
		}
		if g.prevKState && currentOneState{
			g.prevConfirm1State = true
		} else if g.prevKState && currentTwoState{
			g.prevConfirm2State = true
		} else if g.prevKState && currentThreeState{
			g.prevConfirm3State = true
		} else if g.prevKState && currentYState && g.prevConfirm1State{
			g.prevYState = true
			dataStr := fmt.Sprintf("%d\n%d\n%d\n%d\n%d\n%d\n%d\n%d\n%d", g.Score, g.Click, g.CntBossWin, g.PointsEarned, g.PointsSpent,	g.TapLevel, g.TapBotLevel, g.TapBotPrice, g.TapPrice)
			encrypted, err := encrypt([]byte(dataStr), "your-secret-password")
			if err != nil {
				log.Println("Ошибка шифрования:", err)
			} else {
				encoded := base64.StdEncoding.EncodeToString(encrypted)
				err = os.WriteFile("save/save_1.enc", []byte(encoded), 0644)
				if err != nil {
					log.Println("Ошибка записи файла:", err)
				}
			}
		} else if g.prevKState && currentYState && g.prevConfirm2State{
			g.prevYState = true
			dataStr := fmt.Sprintf("%d\n%d\n%d\n%d\n%d\n%d\n%d\n%d\n%d", g.Score, g.Click, g.CntBossWin, g.PointsEarned, g.PointsSpent,	g.TapLevel, g.TapBotLevel, g.TapBotPrice, g.TapPrice)
			encrypted, err := encrypt([]byte(dataStr), "your-secret-password")
			if err != nil {
				log.Println("Ошибка шифрования:", err)
			} else {
				encoded := base64.StdEncoding.EncodeToString(encrypted)
				err = os.WriteFile("save/save_2.enc", []byte(encoded), 0644)
				if err != nil {
					log.Println("Ошибка записи файла:", err)
				}
			}
		} else if g.prevKState && currentYState && g.prevConfirm3State{
			g.prevYState = true
			dataStr := fmt.Sprintf("%d\n%d\n%d\n%d\n%d\n%d\n%d\n%d\n%d", g.Score, g.Click, g.CntBossWin, g.PointsEarned, g.PointsSpent,	g.TapLevel, g.TapBotLevel, g.TapBotPrice, g.TapPrice)
			encrypted, err := encrypt([]byte(dataStr), "your-secret-password")
			if err != nil {
				log.Println("Ошибка шифрования:", err)
			} else {
				encoded := base64.StdEncoding.EncodeToString(encrypted)
				err = os.WriteFile("save/save_3.enc", []byte(encoded), 0644)
				if err != nil {
					log.Println("Ошибка записи файла:", err)
				}
			}
		} else if g.prevKState && currentNState && (g.prevConfirm1State || g.prevConfirm2State || g.prevConfirm3State) {
			g.prevNState = true
		}
		if currentLState {
			g.prevLState = !g.prevLState
		}
		if g.prevLState && currentOneState{
			g.prevConfirm1State = true
		} else if g.prevLState && currentTwoState{
			g.prevConfirm2State = true
		} else if g.prevLState && currentThreeState{
			g.prevConfirm3State = true
		} else if g.prevLState && currentYState && g.prevConfirm1State{
			g.prevYState = true
			encrypted, err := os.ReadFile("save/save_1.enc")
    		if err != nil {
    			log.Println("Ошибка чтения файла:", err)
			} else {
				decoded, err := base64.StdEncoding.DecodeString(string(encrypted))
				if err != nil {
					log.Println("Ошибка декодирования base64:", err)
				} else {
					decrypted, err := decrypt(decoded, "your-secret-password")
					if err != nil {
						log.Println("Ошибка дешифрования:", err)
					} else {
						data := strings.Split(string(decrypted), "\n")
						if len(data) >= 9 {
							g.Score, _ = strconv.Atoi(data[0])
							g.Click, _ = strconv.Atoi(data[1])
							g.CntBossWin, _ = strconv.Atoi(data[2])
							g.PointsEarned, _ = strconv.Atoi(data[3])
							g.PointsSpent, _ = strconv.Atoi(data[4])
							g.TapLevel, _ = strconv.Atoi(data[5])
							g.TapBotLevel, _ = strconv.Atoi(data[6])
							g.TapBotPrice, _ = strconv.Atoi(data[7])
							g.TapPrice, _ = strconv.Atoi(data[8])
							g.ScoreText = strconv.Itoa(g.Score)
						}
					}
				}
			}
		} else if g.prevLState && currentYState && g.prevConfirm2State{
			g.prevYState = true
			encrypted, err := os.ReadFile("save/save_2.enc")
    		if err != nil {
    			log.Println("Ошибка чтения файла:", err)
			} else {
				decoded, err := base64.StdEncoding.DecodeString(string(encrypted))
				if err != nil {
					log.Println("Ошибка декодирования base64:", err)
				} else {
					decrypted, err := decrypt(decoded, "your-secret-password")
					if err != nil {
						log.Println("Ошибка дешифрования:", err)
					} else {
						data := strings.Split(string(decrypted), "\n")
						if len(data) >= 9 {
							g.Score, _ = strconv.Atoi(data[0])
							g.Click, _ = strconv.Atoi(data[1])
							g.CntBossWin, _ = strconv.Atoi(data[2])
							g.PointsEarned, _ = strconv.Atoi(data[3])
							g.PointsSpent, _ = strconv.Atoi(data[4])
							g.TapLevel, _ = strconv.Atoi(data[5])
							g.TapBotLevel, _ = strconv.Atoi(data[6])
							g.TapBotPrice, _ = strconv.Atoi(data[7])
							g.TapPrice, _ = strconv.Atoi(data[8])
							g.ScoreText = strconv.Itoa(g.Score)
						}
					}
				}
			}
		} else if g.prevLState && currentYState && g.prevConfirm3State{
			g.prevYState = true
			encrypted, err := os.ReadFile("save/save_3.enc")
    		if err != nil {
    			log.Println("Ошибка чтения файла:", err)
			} else {
				decoded, err := base64.StdEncoding.DecodeString(string(encrypted))
				if err != nil {
					log.Println("Ошибка декодирования base64:", err)
				} else {
					decrypted, err := decrypt(decoded, "your-secret-password")
					if err != nil {
						log.Println("Ошибка дешифрования:", err)
					} else {
						data := strings.Split(string(decrypted), "\n")
						if len(data) >= 9 {
							g.Score, _ = strconv.Atoi(data[0])
							g.Click, _ = strconv.Atoi(data[1])
							g.CntBossWin, _ = strconv.Atoi(data[2])
							g.PointsEarned, _ = strconv.Atoi(data[3])
							g.PointsSpent, _ = strconv.Atoi(data[4])
							g.TapLevel, _ = strconv.Atoi(data[5])
							g.TapBotLevel, _ = strconv.Atoi(data[6])
							g.TapBotPrice, _ = strconv.Atoi(data[7])
							g.TapPrice, _ = strconv.Atoi(data[8])
							g.ScoreText = strconv.Itoa(g.Score)
						}
					}
				}
			}
		} else if g.prevLState && currentNState && (g.prevConfirm1State || g.prevConfirm2State || g.prevConfirm3State) {
			g.prevNState = true
		}
	}
	if g.TapBotLevel > 0 {
		now := time.Now()
		if now.Sub(g.lastAutoClick) >= 2*time.Second {
			g.Score += g.TapBotLevel
			g.ScoreText = strconv.Itoa(g.Score)
			g.lastAutoClick = now
			fmt.Println("Auto-Click! Score:", g.Score)
			g.PointsEarned += g.TapBotLevel
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
	} else if g.BattleThree {
		screen.Fill(color.Black)
		if(g.prevSpaceState){
			geoM := ebiten.GeoM{}
			geoM.Translate(float64(screenWidth/2-155), float64(screenHeight/2-135))
			geoM.Scale(0.198, 0.198)
			logo, _, err := ebitenutil.NewImageFromFile("images/Boss3_2.jpg")
			if err != nil {
				log.Fatal(err)
			}
			op := &ebiten.DrawImageOptions{GeoM: geoM}
			screen.DrawImage(logo, op)
			g.prevSpaceState = false
		} else {
			geoM := ebiten.GeoM{}
			geoM.Translate(float64(screenWidth/2-155), float64(screenHeight/2-135))
			geoM.Scale(0.2, 0.2)
			logo, _, err := ebitenutil.NewImageFromFile("images/Boss3_1.jpg")
			if err != nil {
				log.Fatal(err)
			}
			op := &ebiten.DrawImageOptions{GeoM: geoM}
			screen.DrawImage(logo, op)
		}
		healthText := "BOSS HP: " + strconv.Itoa(g.BattleScore)
		ebitenutil.DebugPrintAt(screen, healthText, 128, 38)
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
		if g.prevTabState {
			screen.Fill(color.Black)
			ebitenutil.DebugPrintAt(screen, "Click: " + strconv.Itoa(g.Click), 80, 55)
			ebitenutil.DebugPrintAt(screen, "Points earned over all time: " + strconv.Itoa(g.PointsEarned), 80, 95)
			ebitenutil.DebugPrintAt(screen, "Points spent over all time: " + strconv.Itoa(g.PointsSpent), 80, 135)
			ebitenutil.DebugPrintAt(screen, "Defeated bosses: " + strconv.Itoa(g.CntBossWin), 80, 175)
		}
		if g.prevKState {
			screen.Fill(color.Black)
			ebitenutil.DebugPrintAt(screen, "Save", 147, 50)
			f1, _ := os.Open("save/save_1.enc")
			f2, _ := os.Open("save/save_2.enc")
			f3, _ := os.Open("save/save_3.enc")
			if f1 != nil {
				ebitenutil.DebugPrintAt(screen, "(1)Save 1 - exists", 105, 80)
			} else {
				ebitenutil.DebugPrintAt(screen, "(1)Save 1 - empty", 105, 80)
			}
			if f2 != nil {
				ebitenutil.DebugPrintAt(screen, "(2)Save 2 - exists", 105, 100)
			} else {
				ebitenutil.DebugPrintAt(screen, "(3)Save 3 - empty", 105, 100)
			}
			if f3 != nil{
				ebitenutil.DebugPrintAt(screen, "(3)Save 3 - exists", 105, 120)
			} else {
				ebitenutil.DebugPrintAt(screen, "(3)Save 3 - empty", 105, 120)
			}
			if g.prevConfirm1State{
				screen.Fill(color.Black)
				ebitenutil.DebugPrintAt(screen, "Are you sure you want to save your", 55, 85)
				ebitenutil.DebugPrintAt(screen, "progress in save 1", 100, 105)
				ebitenutil.DebugPrintAt(screen, "(Y)Yes", 105, 135)
				ebitenutil.DebugPrintAt(screen, "(N)No", 175, 135)
				if (g.prevYState || g.prevNState){
					g.prevNState = false
					g.prevYState = false
					g.prevConfirm1State = false
				}
			} else if g.prevConfirm2State{
				screen.Fill(color.Black)
				ebitenutil.DebugPrintAt(screen, "Are you sure you want to save your", 55, 85)
				ebitenutil.DebugPrintAt(screen, "progress in save 2", 100, 105)
				ebitenutil.DebugPrintAt(screen, "(Y)Yes", 105, 135)
				ebitenutil.DebugPrintAt(screen, "(N)No", 175, 135)
				if (g.prevYState || g.prevNState){
					g.prevNState = false
					g.prevYState = false
					g.prevConfirm2State = false
				}
			} else if g.prevConfirm3State{
				screen.Fill(color.Black)
				ebitenutil.DebugPrintAt(screen, "Are you sure you want to save your", 55, 85)
				ebitenutil.DebugPrintAt(screen, "progress in save 3", 100, 105)
				ebitenutil.DebugPrintAt(screen, "(Y)Yes", 105, 135)
				ebitenutil.DebugPrintAt(screen, "(N)No", 175, 135)
				if (g.prevYState || g.prevNState){
					g.prevNState = false
					g.prevYState = false
					g.prevConfirm3State = false
				}
			}
		}
		if g.prevLState {
			screen.Fill(color.Black)
			ebitenutil.DebugPrintAt(screen, "Load", 147, 50)
			f1, _ := os.Open("save/save_1.enc")
			f2, _ := os.Open("save/save_2.enc")
			f3, _ := os.Open("save/save_3.enc")
			if f1 != nil {
				ebitenutil.DebugPrintAt(screen, "(1)Save 1 - exists", 105, 80)
			} else {
				ebitenutil.DebugPrintAt(screen, "(1)Save 1 - empty", 105, 80)
			}
			if f2 != nil {
				ebitenutil.DebugPrintAt(screen, "(2)Save 2 - exists", 105, 100)
			} else {
				ebitenutil.DebugPrintAt(screen, "(3)Save 3 - empty", 105, 100)
			}
			if f3 != nil{
				ebitenutil.DebugPrintAt(screen, "(3)Save 3 - exists", 105, 120)
			} else {
				ebitenutil.DebugPrintAt(screen, "(3)Save 3 - empty", 105, 120)
			}
			if g.prevConfirm1State{
				screen.Fill(color.Black)
				ebitenutil.DebugPrintAt(screen, "Are you sure you want to load", 75, 80)
				ebitenutil.DebugPrintAt(screen, "progress from save 1", 100, 105)
				ebitenutil.DebugPrintAt(screen, "(Y)Yes", 105, 135)
				ebitenutil.DebugPrintAt(screen, "(N)No", 175, 135)
				if (g.prevYState || g.prevNState){
					g.prevNState = false
					g.prevYState = false
					g.prevConfirm1State = false
				}
			} else if g.prevConfirm2State{
				screen.Fill(color.Black)
				ebitenutil.DebugPrintAt(screen, "Are you sure you want to load", 75, 80)
				ebitenutil.DebugPrintAt(screen, "progress from save 2", 100, 105)
				ebitenutil.DebugPrintAt(screen, "(Y)Yes", 105, 135)
				ebitenutil.DebugPrintAt(screen, "(N)No", 175, 135)
				if (g.prevYState || g.prevNState){
					g.prevNState = false
					g.prevYState = false
					g.prevConfirm2State = false
				}
			} else if g.prevConfirm3State{
				screen.Fill(color.Black)
				ebitenutil.DebugPrintAt(screen, "Are you sure you want to load", 75, 80)
				ebitenutil.DebugPrintAt(screen, "progress from save 3", 100, 105)
				ebitenutil.DebugPrintAt(screen, "(Y)Yes", 105, 135)
				ebitenutil.DebugPrintAt(screen, "(N)No", 175, 135)
				if (g.prevYState || g.prevNState){
					g.prevNState = false
					g.prevYState = false
					g.prevConfirm3State = false
				}
			}
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
	} else if (g.WinBattleThree){
		screen.Fill(color.Black)
		ebitenutil.DebugPrintAt(screen, "You win!!!", 128, 38)
		geoM := ebiten.GeoM{}
		geoM.Translate(float64(screenWidth/2-155), float64(screenHeight/2-135))
		geoM.Scale(0.2, 0.2)
		logo, _, err := ebitenutil.NewImageFromFile("images/Boss3_3.jpg")
		if err != nil {
			log.Fatal(err)
		}
		op := &ebiten.DrawImageOptions{GeoM: geoM}
		screen.DrawImage(logo, op)
		g.WinBattleThree = false
		g.ScreenEnd = true
	}
	ebitenutil.DebugPrint(screen, "Score: " + g.ScoreText)
	ebitenutil.DebugPrintAt(screen, "Shop(S)", 276, 1)
	ebitenutil.DebugPrintAt(screen, "Exit(Esc)", 266, 225)
	ebitenutil.DebugPrintAt(screen, "Battle(F)", 0, 225)
	ebitenutil.DebugPrintAt(screen, "Statistics(Tab)", 120, 1)
	ebitenutil.DebugPrintAt(screen, "Save(K)", 275, 100)
	ebitenutil.DebugPrintAt(screen, "Load(L)", 275, 120)
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