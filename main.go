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
	"github.com/hajimehoshi/ebiten/v2/text"
    "golang.org/x/image/font/basicfont"
	"io/ioutil"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font"
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
	prevIState bool
	prevConfirm1State bool
	prevConfirm2State bool
	prevConfirm3State bool
	WinBattleOne bool
	WinBattleTwo bool
	WinBattleThree bool
	WinBattleOneSkin bool
	WinBattleTwoSkin bool
	WinBattleThreeSkin bool
	BattleOne bool
	BattleTwo bool
	BattleThree bool
	SkinsTwo bool
	SkinsThree bool
	SkinsFour bool
	SkinsFive bool
	SkinsSix bool
	InstallSkinsTwo bool
	InstallSkinsThree bool
	InstallSkinsFour bool
	InstallSkinsFive bool
	InstallSkinsSix bool
	Skin int
	SkinPrice int
	ScreenEnd bool
	lastAutoClick time.Time
}

var smallFont font.Face

func init() {
  // читаем файл
  b, err := ioutil.ReadFile("assets/IBM2.ttf")
  if err != nil {
    log.Fatal(err)
  }
  // парсим
  tt, err := opentype.Parse(b)
  if err != nil {
    log.Fatal(err)
  }
  // создаём face размером 12 точек (можно 8, 10, 14 и т.д.)
  smallFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
    Size:    5,      // размер шрифта в pt
    DPI:     144,
    Hinting: font.HintingFull,
  })
  if err != nil {
    log.Fatal(err)
  }

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
	currentIState := inpututil.IsKeyJustPressed(ebiten.KeyI)
	currentOneState := inpututil.IsKeyJustPressed(ebiten.Key1)
	currentTwoState := inpututil.IsKeyJustPressed(ebiten.Key2)
	currentThreeState := inpututil.IsKeyJustPressed(ebiten.Key3)
	currentFourState := inpututil.IsKeyJustPressed(ebiten.Key4)
	currentFiveState := inpututil.IsKeyJustPressed(ebiten.Key5)
	currentSixState := inpututil.IsKeyJustPressed(ebiten.Key6)
	if (g.WinBattleOneSkin && g.WinBattleTwoSkin && g.WinBattleThreeSkin){
		g.SkinsFive = true
	}
	if (g.Click >= 500){
		g.SkinsSix = true
	}
	if (g.TapBotLevel >= 10){
		g.SkinsThree = true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyTab){
		g.prevTabState = !g.prevTabState
		g.prevKState = false
		g.prevSState = false
		g.prevIState = false
		g.prevFState = false
		g.prevLState = false
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
			g.SkinsFour = true
			g.Score += 200
			g.ScoreText = strconv.Itoa(g.Score)
			g.CntBossWin++
			g.PointsEarned += 200
			g.WinBattleOneSkin = true
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
			g.WinBattleTwoSkin = true
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
			g.WinBattleThreeSkin = true
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
			g.prevKState = false
			g.prevLState = false
			g.prevIState = false
			g.prevFState = false
			g.prevTabState = false
		}
		if currentFState{
			g.prevFState = !g.prevFState
			g.prevKState = false
			g.prevSState = false
			g.prevIState = false
			g.prevLState = false
			g.prevTabState = false
		}
		if currentIState{
			g.prevIState = !g.prevIState
			g.prevKState = false
			g.prevSState = false
			g.prevLState = false
			g.prevFState = false
			g.prevTabState = false
		}
		if g.prevIState && currentFourState{
			if g.SkinsFour{
				g.InstallSkinsTwo = false
				g.InstallSkinsThree = false
				g.InstallSkinsFour = true
				g.InstallSkinsFive = false
				g.InstallSkinsSix = false
			}
		} else if g.prevIState && currentOneState{
			g.InstallSkinsTwo = false
			g.InstallSkinsThree = false
			g.InstallSkinsFour = false
			g.InstallSkinsFive = false
			g.InstallSkinsSix = false
		} else if g.prevIState && currentSixState{
			if g.SkinsSix{
				g.InstallSkinsTwo = false
				g.InstallSkinsThree = false
				g.InstallSkinsFour = false
				g.InstallSkinsFive = false
				g.InstallSkinsSix = true
			}
		} else if g.prevIState && currentThreeState{
			if g.SkinsThree{
				g.InstallSkinsTwo = false
				g.InstallSkinsThree = true
				g.InstallSkinsFour = false
				g.InstallSkinsFive = false
				g.InstallSkinsSix = false
			}
		} else if g.prevIState && currentTwoState{
			if g.SkinsTwo{
				g.InstallSkinsTwo = true
				g.InstallSkinsThree = false
				g.InstallSkinsFour = false
				g.InstallSkinsFive = false
				g.InstallSkinsSix = false
			}
		} else if g.prevIState && currentFiveState{
			if g.SkinsTwo{
				g.InstallSkinsTwo = false
				g.InstallSkinsThree = false
				g.InstallSkinsFour = false
				g.InstallSkinsFive = true
				g.InstallSkinsSix = false
			}
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
		} else if g.prevSState && currentTwoState{
			if g.TapBotPrice <= g.Score{
				g.Score -= g.TapBotPrice
				g.ScoreText = strconv.Itoa(g.Score)
				fmt.Println("Score:", g.Score)
				g.TapBotPrice += 100;
				g.TapBotLevel += 1;
				g.PointsSpent += g.TapBotPrice
			}
		} else if g.prevSState && currentFourState{
			if g.SkinPrice <= g.Score{
				g.Score -= g.SkinPrice
				g.ScoreText = strconv.Itoa(g.Score)
				fmt.Println("Score:", g.Score)
				g.SkinPrice += 6000
				g.Skin++
				g.SkinsTwo = true
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
			g.prevLState = false
			g.prevSState = false
			g.prevIState = false
			g.prevFState = false
			g.prevTabState = false
		}
		if g.prevKState && currentOneState{
			g.prevConfirm1State = true
		} else if g.prevKState && currentTwoState{
			g.prevConfirm2State = true
		} else if g.prevKState && currentThreeState{
			g.prevConfirm3State = true
		} else if g.prevKState && currentYState && g.prevConfirm1State{
			g.prevYState = true
			dataStr := fmt.Sprintf("%d\n%d\n%d\n%d\n%d\n%d\n%d\n%d\n%d\n%t\n%t\n%t\n%t\n%d\n%d\n%t", g.Score, g.Click, g.CntBossWin, g.PointsEarned, g.PointsSpent,	g.TapLevel, g.TapBotLevel, g.TapBotPrice, g.TapPrice, g.SkinsFour, g.WinBattleOneSkin, g.WinBattleTwoSkin, g.WinBattleThreeSkin, g.Skin, g.SkinPrice, g.SkinsTwo)
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
			dataStr := fmt.Sprintf("%d\n%d\n%d\n%d\n%d\n%d\n%d\n%d\n%d\n%t\n%t\n%t\n%t\n%d\n%d\n%t", g.Score, g.Click, g.CntBossWin, g.PointsEarned, g.PointsSpent,	g.TapLevel, g.TapBotLevel, g.TapBotPrice, g.TapPrice, g.SkinsFour, g.WinBattleOneSkin, g.WinBattleTwoSkin, g.WinBattleThreeSkin, g.Skin, g.SkinPrice, g.SkinsTwo)
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
			dataStr := fmt.Sprintf("%d\n%d\n%d\n%d\n%d\n%d\n%d\n%d\n%d\n%t\n%t\n%t\n%t\n%d\n%d\n%t", g.Score, g.Click, g.CntBossWin, g.PointsEarned, g.PointsSpent,	g.TapLevel, g.TapBotLevel, g.TapBotPrice, g.TapPrice, g.SkinsFour, g.WinBattleOneSkin, g.WinBattleTwoSkin, g.WinBattleThreeSkin, g.Skin, g.SkinPrice, g.SkinsTwo)
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
			g.prevKState = false
			g.prevSState = false
			g.prevIState = false
			g.prevFState = false
			g.prevTabState = false
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
						if len(data) >= 16 {
							g.Score, _ = strconv.Atoi(data[0])
							g.Click, _ = strconv.Atoi(data[1])
							g.CntBossWin, _ = strconv.Atoi(data[2])
							g.PointsEarned, _ = strconv.Atoi(data[3])
							g.PointsSpent, _ = strconv.Atoi(data[4])
							g.TapLevel, _ = strconv.Atoi(data[5])
							g.TapBotLevel, _ = strconv.Atoi(data[6])
							g.TapBotPrice, _ = strconv.Atoi(data[7])
							g.TapPrice, _ = strconv.Atoi(data[8])
							g.SkinsFour, _ = strconv.ParseBool(data[9])
							g.ScoreText = strconv.Itoa(g.Score)
							g.WinBattleOneSkin, _ = strconv.ParseBool(data[10])
							g.WinBattleTwoSkin, _ = strconv.ParseBool(data[11])
							g.WinBattleThreeSkin, _ = strconv.ParseBool(data[12])
							g.Skin, _ = strconv.Atoi(data[13])
							g.SkinPrice, _ = strconv.Atoi(data[14])
							g.SkinsTwo, _ = strconv.ParseBool(data[15])
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
						if len(data) >= 16 {
							g.Score, _ = strconv.Atoi(data[0])
							g.Click, _ = strconv.Atoi(data[1])
							g.CntBossWin, _ = strconv.Atoi(data[2])
							g.PointsEarned, _ = strconv.Atoi(data[3])
							g.PointsSpent, _ = strconv.Atoi(data[4])
							g.TapLevel, _ = strconv.Atoi(data[5])
							g.TapBotLevel, _ = strconv.Atoi(data[6])
							g.TapBotPrice, _ = strconv.Atoi(data[7])
							g.TapPrice, _ = strconv.Atoi(data[8])
							g.SkinsFour, _ = strconv.ParseBool(data[9])
							g.ScoreText = strconv.Itoa(g.Score)
							g.WinBattleOneSkin, _ = strconv.ParseBool(data[10])
							g.WinBattleTwoSkin, _ = strconv.ParseBool(data[11])
							g.WinBattleThreeSkin, _ = strconv.ParseBool(data[12])
							g.Skin, _ = strconv.Atoi(data[13])
							g.SkinPrice, _ = strconv.Atoi(data[14])
							g.SkinsTwo, _ = strconv.ParseBool(data[15])
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
						if len(data) >= 16 {
							g.Score, _ = strconv.Atoi(data[0])
							g.Click, _ = strconv.Atoi(data[1])
							g.CntBossWin, _ = strconv.Atoi(data[2])
							g.PointsEarned, _ = strconv.Atoi(data[3])
							g.PointsSpent, _ = strconv.Atoi(data[4])
							g.TapLevel, _ = strconv.Atoi(data[5])
							g.TapBotLevel, _ = strconv.Atoi(data[6])
							g.TapBotPrice, _ = strconv.Atoi(data[7])
							g.TapPrice, _ = strconv.Atoi(data[8])
							g.SkinsFour, _ = strconv.ParseBool(data[9])
							g.ScoreText = strconv.Itoa(g.Score)
							g.WinBattleOneSkin, _ = strconv.ParseBool(data[10])
							g.WinBattleTwoSkin, _ = strconv.ParseBool(data[11])
							g.WinBattleThreeSkin, _ = strconv.ParseBool(data[12])
							g.Skin, _ = strconv.Atoi(data[13])
							g.SkinPrice, _ = strconv.Atoi(data[14])
							g.SkinsTwo, _ = strconv.ParseBool(data[15])
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
		if now.Sub(g.lastAutoClick) >= time.Second {
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
			if g.InstallSkinsFour{
				logo, _, err = ebitenutil.NewImageFromFile("images/skins4.jpg")
			} else if g.InstallSkinsSix{
				logo, _, err = ebitenutil.NewImageFromFile("images/skins6.jpg")
			} else if g.InstallSkinsThree{
				logo, _, err = ebitenutil.NewImageFromFile("images/skins3.jpg")
			} else if g.InstallSkinsTwo{
				logo, _, err = ebitenutil.NewImageFromFile("images/skins2.jpg")
			} else if g.InstallSkinsFive{
				logo, _, err = ebitenutil.NewImageFromFile("images/skins5.jpg")
			}
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
			if g.InstallSkinsFour{
				logo, _, err = ebitenutil.NewImageFromFile("images/skins4.jpg")
			} else if g.InstallSkinsSix{
				logo, _, err = ebitenutil.NewImageFromFile("images/skins6.jpg")
			} else if g.InstallSkinsThree{
				logo, _, err = ebitenutil.NewImageFromFile("images/skins3.jpg")
			} else if g.InstallSkinsTwo{
				logo, _, err = ebitenutil.NewImageFromFile("images/skins2.jpg")
			} else if g.InstallSkinsFive{
				logo, _, err = ebitenutil.NewImageFromFile("images/skins5.jpg")
			}
			if err != nil {
				log.Fatal(err)
			}
			op := &ebiten.DrawImageOptions{GeoM: geoM}
			screen.DrawImage(logo, op)
		}
		if(g.prevSState){
			screen.Fill(color.Black)
			ebitenutil.DebugPrintAt(screen, "(1)Tap level: " + strconv.Itoa(g.TapLevel) + "  " + strconv.Itoa(g.TapPrice) + " price", 75, 50)
			ebitenutil.DebugPrintAt(screen, "(2)Tap bot level: " + strconv.Itoa(g.TapBotLevel) + "  " + strconv.Itoa(g.TapBotPrice) + " price", 75, 90)
			ebitenutil.DebugPrintAt(screen, "(3)Present: " + strconv.Itoa(g.Present) + "/2  " + strconv.Itoa(g.PresentPrice) + " price", 75, 130)
			ebitenutil.DebugPrintAt(screen, "(4)Skins: " + strconv.Itoa(g.Skin) + "/2  " + strconv.Itoa(g.SkinPrice) + " price", 75, 170)
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
		if g.prevIState {
			screen.Fill(color.Black)
			gray := color.RGBA{128, 128, 128, 255}
			text.Draw(screen, "(1)Base skins", basicfont.Face7x13, 135, 53, color.White)
			geoM := ebiten.GeoM{}
			geoM.Translate(float64(screenWidth+1150), float64(screenHeight))
			geoM.Scale(0.04, 0.04)
			logo, _, err := ebitenutil.NewImageFromFile("images/2224.jpg")
			if err != nil {
				log.Fatal(err)
			}
			op := &ebiten.DrawImageOptions{GeoM: geoM}
			screen.DrawImage(logo, op)
			geoM.Translate(float64(screenWidth/4-300), float64(screenHeight/4-195))
			geoM.Scale(1, 1)
			logo2, _, err2 := ebitenutil.NewImageFromFile("images/skins2.jpg")
			if err2 != nil {
				log.Fatal(err2)
			}
			op2 := &ebiten.DrawImageOptions{GeoM: geoM}
			if !g.SkinsTwo{
				op2.ColorM.Scale(0.5, 0.5, 0.5, 1.0)
			}
			screen.DrawImage(logo2, op2)
			geoM.Translate(float64(screenWidth/4-300), float64(screenHeight/4-195))
			logo3, _, err3 := ebitenutil.NewImageFromFile("images/skins3.jpg")
			if err3 != nil {
				log.Fatal(err3)
			}
			op3 := &ebiten.DrawImageOptions{GeoM: geoM}
			if !g.SkinsThree{
				op3.ColorM.Scale(0.5, 0.5, 0.5, 1.0)
			}
			screen.DrawImage(logo3, op3)
			geoM.Translate(float64(screenWidth/4-300), float64(screenHeight/4-195))
			logo4, _, err4 := ebitenutil.NewImageFromFile("images/skins4.jpg")
			if err4 != nil {
				log.Fatal(err4)
			}
			op4 := &ebiten.DrawImageOptions{GeoM: geoM}
			if !g.SkinsFour{
				op4.ColorM.Scale(0.5, 0.5, 0.5, 1.0)
			}
			screen.DrawImage(logo4, op4)
			geoM.Translate(float64(screenWidth/4-300), float64(screenHeight/4-195))
			logo5, _, err5 := ebitenutil.NewImageFromFile("images/skins5.jpg")
			if err5 != nil {
				log.Fatal(err5)
			}
			op5 := &ebiten.DrawImageOptions{GeoM: geoM}
			if !g.SkinsFive{
				op5.ColorM.Scale(0.5, 0.5, 0.5, 1.0)
			}
			screen.DrawImage(logo5, op5)
			if g.SkinsTwo{
				text.Draw(screen, "(2)Seated skins", basicfont.Face7x13, 135, 83, color.White)
			} else {
				text.Draw(screen, "(2)Seated skins", basicfont.Face7x13, 135, 78, gray)
				text.Draw(screen, "Buy in shop", smallFont, 135, 88, gray)
			}
			if g.SkinsThree{
				text.Draw(screen, "(3)Yami skins", basicfont.Face7x13, 135, 113, color.White)
			} else {
				text.Draw(screen, "(3)Yami skins", basicfont.Face7x13, 135, 108, gray)
				text.Draw(screen, "Upgrade autoclick to lvl 10", smallFont, 135, 118, gray)
			}
			if g.SkinsFour{
				text.Draw(screen, "(4)Vahui skins", basicfont.Face7x13, 135, 143, color.White)
			} else {
				text.Draw(screen, "(4)Vahui skins", basicfont.Face7x13, 135, 138, gray)
				text.Draw(screen, "To open, kill the first boss", smallFont, 135, 148, gray)
			}
			if g.SkinsFive{
				text.Draw(screen, "(5)Loaf skins", basicfont.Face7x13, 135, 173, color.White)
			} else {
				text.Draw(screen, "(5)Loaf skins", basicfont.Face7x13, 135, 168, gray)
				text.Draw(screen, "To open, kill all bosses", smallFont, 135, 178, gray)
			}
			if g.SkinsSix{
				text.Draw(screen, "(6)Secret skins", basicfont.Face7x13, 135, 203, color.White)
				geoM.Translate(float64(screenWidth/4-308), float64(screenHeight/4-214))
				geoM.Scale(1.1, 1.1)
				logo6, _, err6 := ebitenutil.NewImageFromFile("images/skins6.jpg")
				if err6 != nil {
					log.Fatal(err6)
				}
				op6 := &ebiten.DrawImageOptions{GeoM: geoM}
				screen.DrawImage(logo6, op6)
			} else {
				text.Draw(screen, "(6)Secret skins", basicfont.Face7x13, 135, 203, gray)
				geoM.Translate(float64(screenWidth/4-308), float64(screenHeight/4-214))
				geoM.Scale(1.1, 1.1)
				logo6, _, err6 := ebitenutil.NewImageFromFile("images/skins666.png")
				if err6 != nil {
					log.Fatal(err6)
				}
				op6 := &ebiten.DrawImageOptions{GeoM: geoM}
				screen.DrawImage(logo6, op6)
			}
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
				ebitenutil.DebugPrintAt(screen, "(2)Save 2 - empty", 105, 100)
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
				ebitenutil.DebugPrintAt(screen, "(2)Save 2 - empty", 105, 100)
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
	ebitenutil.DebugPrintAt(screen, "Skins(I)", 0, 100)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	err := os.MkdirAll("save", 0777)
	if err != nil {
		fmt.Println("Ошибка при создании директории:", err)
		return
	}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Test screan")
	if err := ebiten.RunGame(&Game{
		SkinPrice: 1500,
		TapBotPrice: 100,
		TapPrice: 20,
		PresentPrice: 500,
	}); err != nil {
		log.Fatal(err)
	}
}