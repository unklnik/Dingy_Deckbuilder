package main

import (
	"fmt"
	"time"

	z "github.com/gen2brain/raylib-go/raylib"
)

var (
	selatkY, selatkY2 float32
	selatkUD          bool
	selatkYchange     = float32(50)
	nexthandANMonOFF  bool
	MAPON             = false
)

func DRAWCAM() {

	dIMSheetDrawRecColorAlpha(bg)                                              //BG PATTERN
	dREC(z.NewRectangle(0, 0, float32(SCRW), float32(SCRH)), CA(z.Black, 120)) //BG SHADOW OVER REC

	if MAPON {
		dMAP()
	} else {
		if !optionsON && !viewDeck && !viewDiscard {
			dDECK()    //DECK
			dDISCARD() //DISCARD
			//PLAYER IM
			siz := UNIT * 4
			sizTXT := float32(2)
			portraitREC := z.NewRectangle(levbgREC.X-(siz+UNIT/2), levbgREC.Y+qTXTheight(FONT2, sizTXT), siz, siz)
			dIMAcolor(pl.im, portraitREC, 120, sepiaDRK2)
			dRECLINE(portraitREC, 4, sepiaDRK2)
			dTXTcntRecYToffsetFont2(pl.nm, portraitREC, sizTXT, 0, sepiaDRK2)
			//TURNS TXT
			sizTXT = float32(1.2)
			yTXT := portraitREC.Y + portraitREC.Height + UNIT/4
			dTXTcntXFont2("HP: "+fmt.Sprint(pl.hp), portraitREC.X+portraitREC.Width/2, yTXT, sizTXT, z.Maroon)
			yTXT += qTXTheight(FONT2, sizTXT)
			dTXTcntXFont2("Mana: "+fmt.Sprint(pl.mana), portraitREC.X+portraitREC.Width/2, yTXT, sizTXT, z.DarkBlue)
			yTXT += qTXTheight(FONT2, sizTXT)
			dTXTcntXFont2("Turns: "+fmt.Sprint(pl.turns), portraitREC.X+portraitREC.Width/2, yTXT, sizTXT, z.DarkGreen)
		}
		dRECSHADOWONLY(levbgREC, UNIT/3, z.Black, 170) //LEVREC IM SHADOW
		dIMcolor(levbgIM, levbgREC, sepia)             //LEVREC IM
		dRECLINE(levbgREC, 2, sepia)                   //LEVREC BORDER

		if !optionsON && !viewDeck && !viewDiscard {
			dCARDS()
			dENEMIES()

			//DRAWING
			if cMSREC(deckREC) && !cMSREC(levbgREC) && !cMSREC(deckViewREC) {
				dTXTfont2XYSHADOW("Draw", deckREC.X+UNIT-UNIT/7, deckREC.Y-(UNIT+UNIT/5), 2.2, sepiaDRK2, CA(z.Black, 250), UNIT/7)
				if MSL && drawingT == 0 && pl.turns > 0 {
					pl.turns--
					drawingT = int(FPS / 2)
					drawingREC = z.NewRectangle(deckREC.X, deckREC.Y+cardREC.Height/4, cardREC.Width, cardREC.Height)
					drawingIM = ETC[6]
					drawingANIM.off = false
					drawingANIM.timer = time.Time{}
					drawingANIM.dFrame = 0
					DRAWCARD()
				} else if pl.turns == 0 {
					noTurnsT = int(FPS)
				}
			} else if cMSREC(deckREC) && !cMSREC(levbgREC) && cMSREC(deckViewREC) {
				dTXTfont2XYSHADOW("View", deckREC.X+UNIT-UNIT/7, deckREC.Y-(UNIT+UNIT/5), 2.2, sepiaDRK2, CA(z.Black, 250), UNIT/7)
				if MSL {
					viewDeck = true
				}
			}
			if drawingT > 0 {
				dDRAWING()
			}
			if debug {
				dRECLINE(deckREC, 2, z.Magenta)
			}
		} else if !optionsON && viewDeck && !viewDiscard {
			dVIEWDECK()
		} else if !optionsON && !viewDeck && viewDiscard {
			dVIEWDISCARD()
		} else {
			dOPTIONS()
		}

		//BAT FAIRY
		if batT == 0 {
			dBAT()
		}
		if fairyT == 0 {
			dFAIRY()
		}

		//NO TURNS TEXT
		if noTurnsT > 0 {
			x := levbgREC.X + levbgREC.Width + UNIT/2
			y := levbgREC.Y
			siz := float32(2)
			dTXTfont2XY("Not enough turns!", x, y, siz, sepiaDRK2)
		} else if notselectedT > 0 {
			x := levbgREC.X + levbgREC.Width + UNIT/2
			y := levbgREC.Y
			siz := float32(2)
			dTXTfont2XY("No enemy selected!", x, y, siz, sepiaDRK2)
		} else if nomanaT > 0 {
			x := levbgREC.X + levbgREC.Width + UNIT/2
			y := levbgREC.Y
			siz := float32(2)
			dTXTfont2XY("Not enough mana!", x, y, siz, sepiaDRK2)
		}

		//DRAW NEXT HAND
		if nexthandANMonOFF {
			siz := UNIT * 10
			r := RECCNT(CNT, siz, siz)
			r.Y = levbgREC.Y + levbgREC.Height - (siz/20)*14
			nexthandANM = dAnimRecONCE(nexthandANM, r)
			if nexthandANM.off {
				DRAWNEXTHAND()
				nexthandANM = RESETANIM(nexthandANM)
				nexthandANMonOFF = false
				if len(enemies) > 0 {
					for i := range enemies {
						if enemies[i].selatk {
							enemies[i].selatk = false
						}
					}
				}
			}
		}

		//DRAW PLAYER DEATH
		if playerdeathONOFF {
			siz := float32(SCRH) + UNIT*2
			r := RECCNT(CNT, siz, siz)
			pldeathANM = dAnimRecONCE(pldeathANM, r)
			if pldeathANM.off {
				pldeathANM = RESETANIM(pldeathANM)
				playerdeathONOFF = false
			}
		}
	}

}
func DRAWNOCAMSHADER() { //MARK: NOCAM SHADER
	//MOUSE CURSOR
	siz := UNIT + UNIT/2
	cursorREC = z.NewRectangle(MS.X, MS.Y, siz, siz)
	if clickT > 0 {
		dIMcolorSHADOW(ETC[5], cursorREC, sepiaDRK, CA(z.Black, 180), UNIT/5)
	} else {
		dIMcolorSHADOW(ETC[4], cursorREC, sepiaDRK, CA(z.Black, 180), UNIT/5)
	}

}
func DRAWNOCAMNOSHADER() { //MARK: NOCAM NOSHADER
	SCAN(1, 2, z.Fade(z.Black, 0.3))

	if debug {
		DEBUG()

	}
	if scrollONOFF {
		SCROLL()
	}
}
func dMAP() { //MARK: MAP
	dIM(ETC[1], levmapREC) //MAP BG
	if debug {
		dRECLINE(levmapREC, 2, z.Magenta)
	}

	//START
	dIMshadow(ETC[50], levmapstartR, CA(z.DarkGray, RUINT8(100, 120)), UNIT/12)
	//LEV OBJS
	for i := range levobjs {
		if cMSREC(levobjs[i].r) {
			r2 := levobjs[i].r
			siz := UNIT / 2
			r2.X -= siz / 2
			r2.Y -= siz / 2
			r2.Width += siz
			r2.Height += siz
			dIMshadow(levobjs[i].im, r2, CA(z.DarkGray, RUINT8(100, 120)), UNIT/12)
			dTXTcntRecYToffsetFont2(levobjs[i].nm, levmapREC, 4, -UNIT*5, z.Black)
			if plmapV2 == qRecCNT(levmapstartR) {
				if i < 3 {
					if MSL {
						pllevmapNUM = i
						moveplmap = true
					}
				}
			}

		} else {
			dIMshadow(levobjs[i].im, levobjs[i].r, CA(z.DarkGray, RUINT8(100, 120)), UNIT/12)
		}

		if debug {
			dRECLINE(levobjs[i].r, 2, z.Magenta)
		}
	}
	//LINES
	dLineRecV2toRecV2SHADOWxplus(levmapstartR, levobjs[1].r, 2, 4, UNIT/12, UNIT/7, z.Black, CA(z.DarkGray, RUINT8(180, 220)))                       //ROW 1T
	dLineRecV2toRecV2SHADOW(levmapstartR, levobjs[2].r, 3, 1, UNIT/12, UNIT/7, z.Black, CA(z.DarkGray, RUINT8(180, 220)))                            //ROW 1B
	dLineRECsideCNTtoRECsideCNToffsetSHADOW(levmapstartR, levobjs[0].r, 2, 4, UNIT/5, 0, UNIT/12, UNIT/7, z.Black, CA(z.DarkGray, RUINT8(180, 220))) //ROW 1C

	dLineRECsideCNTtoRECsideCNToffsetSHADOW(levobjs[1].r, levobjs[4].r, 2, 4, UNIT/5, 0, UNIT/12, UNIT/7, z.Black, CA(z.DarkGray, RUINT8(180, 220))) //ROW 2T
	dLineRECsideCNTtoRECsideCNToffsetSHADOW(levobjs[2].r, levobjs[5].r, 2, 4, UNIT/5, 0, UNIT/12, UNIT/7, z.Black, CA(z.DarkGray, RUINT8(180, 220))) //ROW 2B
	dLineRECsideCNTtoRECsideCNToffsetSHADOW(levobjs[0].r, levobjs[3].r, 2, 4, UNIT/5, 0, UNIT/12, UNIT/7, z.Black, CA(z.DarkGray, RUINT8(180, 220))) //ROW 2C

	dLineRECsideCNTtoRECsideCNToffsetSHADOW(levobjs[4].r, levobjs[7].r, 2, 4, UNIT/5, 0, UNIT/12, UNIT/7, z.Black, CA(z.DarkGray, RUINT8(180, 220))) //ROW 3T
	dLineRECsideCNTtoRECsideCNToffsetSHADOW(levobjs[5].r, levobjs[8].r, 2, 4, UNIT/5, 0, UNIT/12, UNIT/7, z.Black, CA(z.DarkGray, RUINT8(180, 220))) //ROW 3B
	dLineRECsideCNTtoRECsideCNToffsetSHADOW(levobjs[3].r, levobjs[6].r, 2, 4, UNIT/5, 0, UNIT/12, UNIT/7, z.Black, CA(z.DarkGray, RUINT8(180, 220))) //ROW 3C

	//CENTER EVENT
	dLineRecV2toRecV2SHADOW(levobjs[7].r, levobjs[9].r, 3, 1, UNIT/12, UNIT/7, z.Black, CA(z.DarkGray, RUINT8(180, 220)))
	dLineRecV2toRecV2SHADOWxplus(levobjs[8].r, levobjs[9].r, 2, 4, UNIT/12, UNIT/7, z.Black, CA(z.DarkGray, RUINT8(180, 220)))
	dLineRECsideCNTtoRECsideCNToffsetSHADOW(levobjs[6].r, levobjs[9].r, 2, 4, UNIT/5, 0, UNIT/12, UNIT/7, z.Black, CA(z.DarkGray, RUINT8(180, 220)))

	dLineRecV2toRecV2SHADOWxplus(levobjs[9].r, levobjs[11].r, 2, 4, UNIT/12, UNIT/7, z.Black, CA(z.DarkGray, RUINT8(180, 220)))                       //ROW 4T
	dLineRecV2toRecV2SHADOW(levobjs[9].r, levobjs[12].r, 3, 1, UNIT/12, UNIT/7, z.Black, CA(z.DarkGray, RUINT8(180, 220)))                            //ROW 4B
	dLineRECsideCNTtoRECsideCNToffsetSHADOW(levobjs[9].r, levobjs[10].r, 2, 4, UNIT/5, 0, UNIT/12, UNIT/7, z.Black, CA(z.DarkGray, RUINT8(180, 220))) //ROW 4C

	dLineRECsideCNTtoRECsideCNToffsetSHADOW(levobjs[11].r, levobjs[14].r, 2, 4, UNIT/5, 0, UNIT/12, UNIT/7, z.Black, CA(z.DarkGray, RUINT8(180, 220))) //ROW 3T
	dLineRECsideCNTtoRECsideCNToffsetSHADOW(levobjs[12].r, levobjs[15].r, 2, 4, UNIT/5, 0, UNIT/12, UNIT/7, z.Black, CA(z.DarkGray, RUINT8(180, 220))) //ROW 3B
	dLineRECsideCNTtoRECsideCNToffsetSHADOW(levobjs[10].r, levobjs[13].r, 2, 4, UNIT/5, 0, UNIT/12, UNIT/7, z.Black, CA(z.DarkGray, RUINT8(180, 220))) //ROW 3C

	//BOSS
	dLineRecV2toRecV2SHADOW(levobjs[14].r, levobjs[16].r, 3, 1, UNIT/12, UNIT/7, z.Black, CA(z.DarkGray, RUINT8(180, 220)))
	dLineRecV2toRecV2SHADOWxplus(levobjs[15].r, levobjs[16].r, 2, 4, UNIT/12, UNIT/7, z.Black, CA(z.DarkGray, RUINT8(180, 220)))
	dLineRECsideCNTtoRECsideCNToffsetSHADOW(levobjs[13].r, levobjs[16].r, 2, 4, UNIT/5, 0, UNIT/12, UNIT/7, z.Black, CA(z.DarkGray, RUINT8(180, 220)))

	//PLAYER ICON
	siz := UNIT * 4
	ETC[51] = dIMrotatingColorSHADOW(ETC[51], RECCNT(plmapV2, siz, siz), 2, UNIT/7, z.Magenta, CA(z.Black, 120))
	if plmapV2 != qRecCNT(levobjs[pllevmapNUM].r) && moveplmap {
		x, y := VELXY(plmapV2, qRecCNT(levobjs[pllevmapNUM].r), UNIT/4)
		plmapV2.X += x
		plmapV2.Y += y
	} else if plmapV2 == qRecCNT(levobjs[pllevmapNUM].r) {
		moveplmap = false
	}
}
func dVIEWDECK() { //MARK: VIEW DECK DISCARD
	dREC(z.NewRectangle(0, 0, float32(SCRW), float32(SCRH)), CA(z.Black, 220))
	viewDeck = dCLOSEtopRight(UNIT, UNIT/4, sepiaDRK, viewDeck) //CLOSE
}
func dVIEWDISCARD() {
	dREC(z.NewRectangle(0, 0, float32(SCRW), float32(SCRH)), CA(z.Black, 220))
	viewDiscard = dCLOSEtopRight(UNIT, UNIT/4, sepiaDRK, viewDiscard) //CLOSE
}
func dDECK() {
	if cMSREC(deckREC) && !cMSREC(deckViewREC) {
		dIMroColor(ETC[2], deckREC, -15, sepiaDRK)
	} else {
		dIMcolor(ETC[2], deckREC, sepiaDRK)
	}

	//DECK REMAINDER NUM
	r := qResizeRecProportionalH(ETC[9].r, deckREC.Width/5)
	r.X = levbgREC.X - r.Width - UNIT/4
	r.Y = deckREC.Y + deckREC.Height - r.Height
	dIMcolorSHADOW(ETC[9], r, sepiaDRK, CA(z.Black, 200), UNIT/7)
	txt := fmt.Sprint(len(pl.deck))
	if len(pl.deck) < 10 {
		txt = "0" + txt
	}
	dTXTcntRecFont2(txt, r, 1.4, z.Black)

	//VIEW
	if cMSREC(deckViewREC) {
		dIMroColor(ETC[22], deckViewREC, -30, sepiaDRK)
	} else {
		dIMcolor(ETC[22], deckViewREC, sepiaDRK)
	}

	//END TURN
	clear := false
	siz := UNIT + UNIT/2
	r = z.NewRectangle(deckREC.X-siz, deckREC.Y-siz, siz, siz)
	if cPointRec(MS, r) {
		dIM(ETC[29], r)
		dTXTcntRecYToffsetFont2("End Turn", r, 1.2, 0, sepiaDRK2)
		if MSL && !enmTURN {
			enmTURN = true
			enemies[0].atk = true
		}
	} else {
		dIM(ETC[28], r)
	}
	if clear {
		pl.hand = REMCARDOFF(pl.hand)
	}

}
func dDISCARD() {
	w := UNIT * 3.5
	r := qResizeRecProportionalH(ETC[21].r, w)
	r.X = levbgREC.X + levbgREC.Width + UNIT/2
	r.Y = levbgREC.Y + levbgREC.Height + UNIT/2 - r.Height
	if cMSREC(r) {
		dIMroColor(ETC[21], r, 15, sepiaDRK)
		dTXTcntRecYToffsetFont2("Discard", r, 2.2, UNIT/7, sepiaDRK2)
		if MSL {
			viewDiscard = true
		}
	} else {
		dIMcolor(ETC[21], r, sepiaDRK)
	}
}
func dOPTIONS() { //MARK: OPTIONS
	lineW := float32(2)
	if SCRW > 1920 {
		lineW = float32(4)
	}
	dREC(levbgREC, CA(z.Black, 180))
	dTXTcntRecWfont2("Options", levbgREC, UNIT/4, 3, sepiaDRK)
	x := levbgREC.X + UNIT/2
	y := qTXTheight(FONT2, 3) + UNIT
	siz := float32(2)
	h := qTXTheight(FONT2, siz)
	dTXTfont2XY("Fullscreen", x, y, 2, sepiaDRK2)
	x2 := levbgREC.X + levbgREC.Width/4
	b := mBUTTON(x2, y+h/4, (UNIT/20)*12, lineW, sepiaDRK2, z.DarkGreen, z.Maroon)
	FULLSCREEN = dBUTTON(b, FULLSCREEN)
	y += h
	dTXTfont2XY("Mute FX", x, y, 2, sepiaDRK2)
	b.r.Y = y + h/4
	MUTEFX = dBUTTON(b, MUTEFX)
	y += h
	dTXTfont2XY("Mute Music", x, y, 2, sepiaDRK2)
	b.r.Y = y + h/4
	MUTEMUSIC = dBUTTON(b, MUTEMUSIC)
	y += h

	optionsON = dCLOSEtopRight(UNIT, UNIT/4, sepiaDRK, optionsON) //CLOSE
}
func dBAT() {
	if batLR {
		bat = dAnimRecLoopFlipShadow(bat, batREC, CA(z.Black, 140), UNIT)
		batREC.X -= UNIT / 10
		if batREC.X+batREC.Width < 0 {
			batREC.Y = RF32(UNIT, UNIT*4)
			batLR = false
			batT = int(FPS)
			batORfairy = FLIPCOIN()
		}
	} else {
		bat = dAnimRecLoopShadow(bat, batREC, CA(z.Black, 140), UNIT)
		batREC.X += UNIT / 10
		if batREC.X > float32(SCRW) {
			batREC.Y = RF32(UNIT, UNIT*4)
			batLR = true
			batT = int(FPS)
			batORfairy = FLIPCOIN()
		}
	}
}
func dFAIRY() {
	if fairyLR {
		fairy = dAnimRecLoopFlipShadow(fairy, fairyREC, CA(z.Black, 140), UNIT)
		fairyREC.X -= UNIT / 10
		if fairyREC.X+fairyREC.Width < 0 {
			fairyREC.Y = RF32(UNIT, UNIT*4)
			fairyLR = false
			fairyT = int(FPS) * RINT(30, 120)
			batORfairy = FLIPCOIN()
		}
	} else {
		fairy = dAnimRecLoopShadow(fairy, fairyREC, CA(z.Black, 140), UNIT)
		fairyREC.X += UNIT / 10
		if fairyREC.X > float32(SCRW) {
			fairyREC.Y = RF32(UNIT, UNIT*4)
			fairyLR = true
			fairyT = int(FPS) * RINT(30, 120)
			batORfairy = FLIPCOIN()
		}
	}
}

func dENEMIES() { //MARK: ENEMIES
	clear := false
	for i := range enemies {
		if !enemies[i].off {
			flip := false
			switch enemies[i].pos {
			case 1, 4:
				if enemies[i].imFACES == "l" {
					flip = true
				}
			case 3, 6:
				if enemies[i].imFACES == "r" {
					flip = true
				}
			}
			if enemies[i].atk {
				if flip {
					enemies[i].a2 = dAnimRecLoopFlipShadowOUTLINEonce(enemies[i].a2, enemies[i].r, CA(z.Black, 140), UNIT/3, UNIT/9)

				} else {
					enemies[i].a2 = dAnimRecLoopShadowOUTLINEonce(enemies[i].a2, enemies[i].r, CA(z.Black, 140), UNIT/3, UNIT/9)
				}
			} else {
				if flip {
					enemies[i].a = dAnimRecLoopFlipShadowOUTLINE(enemies[i].a, enemies[i].r, CA(z.Black, 140), UNIT/3, UNIT/9)
					if enemies[i].sel {
						enemies[i].a = dAnimRecLoopFlipColor(enemies[i].a, enemies[i].r, CA(z.Magenta, RUINT8(100, 150)))
					}
				} else {
					enemies[i].a = dAnimRecLoopShadowOUTLINE(enemies[i].a, enemies[i].r, CA(z.Black, 140), UNIT/3, UNIT/9)
					if enemies[i].sel {
						enemies[i].a = dAnimRecLoopColor(enemies[i].a, enemies[i].r, CA(z.Magenta, RUINT8(100, 150)))
					}
				}
			}

			if enemies[i].hpHIT {
				enemies[i].hit = dAnimRecONCEalpha(enemies[i].hit, qResizeRecOffsetLRGR(enemies[i].rC, UNIT*2), 180)
				if enemies[i].hit.off {
					enemies[i].hit = RESETANIM(enemies[i].hit)
					enemies[i].hpHIT = false
				}
			}

			if enemies[i].a2.off { //RESET ATTACK
				enemies[i].a2 = RESETANIM(enemies[i].a2)
				enemies[i].played = true
				enemies[i].atk = false
				if i == len(enemies)-1 {
					for i := range enemies {
						enemies[i].played = false
					}
					if len(pl.hand) > 0 {
						clear = true
						for i := range pl.hand {
							pl.discard = append(pl.discard, pl.hand[i])
							pl.hand[i].off = true
							pl.hand[i].offPLAYED = true
						}
					}
					pl.turns = pl.turnsMAX
					nexthandANMonOFF = true
				}
			}

			//SELATK ARROW
			if enemies[i].selatk {
				siz := UNIT * 2
				r := z.NewRectangle(enemies[i].rC.X+enemies[i].rC.Width/2-siz/2, selatkY, siz, siz)
				yt := enemies[i].rC.Y - UNIT
				switch enemies[i].nm {
				case "Barbara", "Hogan":
					yt -= UNIT * 2
					selatkY2 = yt - selatkYchange
				case "Spaghetti":
					yt -= UNIT
					selatkY2 = yt - selatkYchange
					if flip {
						r.X += UNIT / 2
					} else {
						r.X -= UNIT / 2
					}
				case "Spike":
					yt -= UNIT
					selatkY2 = yt - selatkYchange
					if flip {
						r.X -= UNIT
					} else {
						r.X += UNIT
					}
				case "The Crow", "Pothead", "Audrey", "Baldrick", "Pet Rock", "Octeye", "Marlin":
					yt -= UNIT
					selatkY2 = yt - selatkYchange
				case "Litchi", "Ogg", "Priscilla", "Blob", "Slime Too":
					yt -= UNIT / 2
					selatkY2 = yt - selatkYchange
				case "Igor", "Jim":
					yt += UNIT / 2
					selatkY2 = yt - selatkYchange
					if flip {
						r.X -= UNIT
					} else {
						r.X += UNIT
					}
				case "Mischeevis":
					yt += UNIT
					selatkY2 = yt - selatkYchange
					if flip {
						r.X += UNIT
					} else {
						r.X -= UNIT
					}
				}
				dIMcolor(ETC[23], r, CA(cRAN(), 150)) //SEL ARROW
				if selatkUD {
					selatkY += 5
					if selatkY > yt {
						selatkY = yt
						selatkUD = false
					}
				} else {
					selatkY -= 5
					if selatkY < selatkY2 {
						selatkY = selatkY2
						selatkUD = true
					}
				}
			}
			//HP DEF EFFECTS BAR
			spc := float32(UNIT / 10)
			var totallen float32
			bgrec := R(enemies[i].rC.X, enemies[i].rC.Y-(UNIT+UNIT/2), (UNIT/20)*22, (UNIT/20)*12)
			bgrecW := bgrec.Width
			if enemies[i].hp > 9 {
				bgrecW += (UNIT / 20) * 4
			}
			totallen += bgrecW
			if enemies[i].poisonAMOUNT > 0 {
				totallen += spc
				if enemies[i].hp > 9 && enemies[i].poisonAMOUNT < 10 {
					bgrecW -= (UNIT / 20) * 4
				} else if enemies[i].hp < 10 && enemies[i].poisonAMOUNT > 9 {
					bgrecW += (UNIT / 20) * 4
				}
				totallen += bgrecW
			}
			if enemies[i].burnAMOUNT > 0 {
				totallen += spc
				if enemies[i].poisonAMOUNT > 9 && enemies[i].burnAMOUNT < 10 {
					bgrecW -= (UNIT / 20) * 4
				} else if enemies[i].poisonAMOUNT < 10 && enemies[i].burnAMOUNT > 9 {
					bgrecW += (UNIT / 20) * 4
				}
				totallen += bgrecW
			}
			if enemies[i].stunLEN > 0 {
				totallen += spc
				if enemies[i].burnAMOUNT > 9 && enemies[i].stunLEN < 10 {
					bgrecW -= (UNIT / 20) * 4
				} else if enemies[i].burnAMOUNT < 10 && enemies[i].stunLEN > 9 {
					bgrecW += (UNIT / 20) * 4
				}
				totallen += bgrecW
			}

			bgrec.X = enemies[i].rC.X + enemies[i].rC.Width/2 - totallen/2 - bgrec.Width/2
			//ALIGN
			switch enemies[i].nm {
			case "Octeye":
				bgrec.X += (UNIT / 20) * 10
			case "Claude":
				bgrec.X += (UNIT / 20) * 15
				bgrec.Y += (UNIT / 20) * 5
			case "Spaghetti":
				if flip {
					bgrec.X += (UNIT / 20) * 18
				} else {
					//bgrec.X += (UNIT / 20) * 7
				}
				bgrec.Y += (UNIT / 20) * 2
			case "Priscilla":
				bgrec.X += (UNIT / 20) * 15
				bgrec.Y += (UNIT / 20) * 10
			case "Jim":
				if flip {
					bgrec.X -= (UNIT / 20) * 14
				} else {
					bgrec.X += (UNIT / 20) * 25
				}
				bgrec.Y += (UNIT / 20) * 2
			case "Audrey":
				if flip {
					bgrec.X += (UNIT / 20) * 20
				} else {
					bgrec.X += (UNIT / 20) * 10
				}
			case "Pothead":
				if flip {
					bgrec.X += (UNIT / 20) * 40
				} else {
					bgrec.X += (UNIT / 20) * 5
				}
			case "Spike":
				if flip {
					bgrec.X -= (UNIT / 20) * 10
				} else {
					bgrec.X += (UNIT / 20) * 25
				}
			case "Barbara":
				bgrec.X += (UNIT / 20) * 10
				bgrec.Y -= (UNIT / 20) * 22
			case "Martin":
				bgrec.X += (UNIT / 20) * 15
				bgrec.Y += (UNIT / 20) * 25
			case "Baldrick":
				bgrec.X += (UNIT / 20) * 10
			case "Pet Rock":
				bgrec.X += (UNIT / 20) * 18
			case "Ogg":
				if flip {
					bgrec.X += (UNIT / 20) * 8
				} else {
					bgrec.X += (UNIT / 20) * 18
				}
				bgrec.Y += (UNIT / 20) * 2
			case "Marlin":
				bgrec.X += (UNIT / 20) * 15
			case "Harkonen":
				bgrec.X += (UNIT / 20) * 15
				bgrec.Y += (UNIT / 20) * 12
			case "Blob":
				bgrec.X += (UNIT / 20) * 12
				bgrec.Y += (UNIT / 20) * 8
			case "Slime":
				bgrec.X += (UNIT / 20) * 12
				bgrec.Y += (UNIT / 20) * 4
			case "Slime Too", "Bernie":
				bgrec.X += (UNIT / 20) * 12
			case "Hogan":
				bgrec.X += (UNIT / 20) * 10
				bgrec.Y -= (UNIT / 20) * 4
			case "The Crow":
				if flip {
					bgrec.X += (UNIT / 20) * 10
				} else {
					bgrec.X += (UNIT / 20) * 7
				}
				bgrec.Y += (UNIT / 20) * 2
			case "Litchi":
				if flip {
					bgrec.X += (UNIT / 20) * 14
				} else {
					bgrec.X += (UNIT / 20) * 7
				}
				bgrec.Y += (UNIT / 20) * 2
			case "Mischeevis":
				if flip {
					bgrec.X += (UNIT / 20) * 30
				} else {
					bgrec.X -= (UNIT / 20) * 7
				}
				bgrec.Y += (UNIT / 20) * 18
			case "Igor":
				if flip {
					bgrec.X -= (UNIT / 20) * 5
				} else {
					bgrec.X += (UNIT / 20) * 25
				}
				bgrec.Y += (UNIT / 20) * 18
			}

			//HP
			if enemies[i].hp > 9 {
				bgrec.Width += (UNIT / 20) * 4
			}
			dREC(bgrec, CA(sepiaDRK2, 100))
			dRECLINE(bgrec, 2, z.Black)
			y := bgrec.Y
			r := R(bgrec.X, y, (UNIT/20)*12, (UNIT/20)*12)
			dIMcolor(ETC[24], r, z.Maroon)
			dTXTfont1XY(fmt.Sprint(enemies[i].hp), bgrec.X+bgrec.Width/2, y, 0.8, z.Black)
			if cPointRec(MS, r) {
				dTXTfont2XYbgrec("HP", r.X, r.Y-qTXTheight(FONT1, 1), 1, sepiaDRK2, CA(z.Black, 180), 2)
			}
			//POISON
			if enemies[i].poisonAMOUNT > 0 {
				bgrec.X += bgrec.Width + spc
				r.X += bgrec.Width + spc
				if enemies[i].hp > 9 && enemies[i].poisonAMOUNT < 10 {
					bgrec.Width -= (UNIT / 20) * 4
				} else if enemies[i].hp < 10 && enemies[i].poisonAMOUNT > 9 {
					bgrec.Width += (UNIT / 20) * 4
				}
				dREC(bgrec, CA(sepiaDRK2, 100))
				dRECLINE(bgrec, 2, z.Black)
				dIMcolor(ETC[25], r, GREENdark1())
				dTXTfont1XY(fmt.Sprint(enemies[i].poisonAMOUNT), bgrec.X+bgrec.Width/2, y, 0.8, z.Black)
				if cPointRec(MS, r) {
					dTXTfont2XYbgrec("Poison", r.X, r.Y-qTXTheight(FONT1, 1), 1, sepiaDRK2, CA(z.Black, 180), 2)
				}
			}
			//BURN
			if enemies[i].burnAMOUNT > 0 {
				bgrec.X += bgrec.Width + spc
				r.X += bgrec.Width + spc
				if enemies[i].poisonAMOUNT > 9 && enemies[i].burnAMOUNT < 10 {
					bgrec.Width -= (UNIT / 20) * 4
				} else if enemies[i].poisonAMOUNT < 10 && enemies[i].burnAMOUNT > 9 {
					bgrec.Width += (UNIT / 20) * 4
				}
				dREC(bgrec, CA(sepiaDRK2, 100))
				dRECLINE(bgrec, 2, z.Black)
				dIMcolor(ETC[26], r, z.Maroon)
				dTXTfont1XY(fmt.Sprint(enemies[i].burnAMOUNT), bgrec.X+bgrec.Width/2, y, 0.8, z.Black)
				if cPointRec(MS, r) {
					dTXTfont2XYbgrec("Burn", r.X, r.Y-qTXTheight(FONT1, 1), 1, sepiaDRK2, CA(z.Black, 180), 2)
				}
			}
			//STUN
			if enemies[i].stunLEN > 0 {
				bgrec.X += bgrec.Width + spc
				r.X += bgrec.Width + spc
				r = qResizeRecOffsetSMLR(r, 2)
				if enemies[i].burnAMOUNT > 9 && enemies[i].stunLEN < 10 {
					bgrec.Width -= (UNIT / 20) * 4
				} else if enemies[i].burnAMOUNT < 10 && enemies[i].stunLEN > 9 {
					bgrec.Width += (UNIT / 20) * 4
				}
				dREC(bgrec, CA(sepiaDRK2, 100))
				dRECLINE(bgrec, 2, z.Black)
				dIMcolor(ETC[27], r, z.Magenta)
				dTXTfont1XY(fmt.Sprint(enemies[i].stunLEN), bgrec.X+bgrec.Width/2, y, 0.8, z.Black)
				if cPointRec(MS, r) {
					dTXTfont2XYbgrec("Stun", r.X, r.Y-qTXTheight(FONT1, 1), 1, sepiaDRK2, CA(z.Black, 180), 2)
				}
			}

			if debug {
				dRECLINE(enemies[i].r, 2, z.Magenta)
				dRECLINE(enemies[i].rC, 4, z.Magenta)
			}

			if cMSREC(enemies[i].rC) && noTurnsT == 0 {
				dENEMYSTATS(enemies[i])
				enemies[i].sel = true
				if MSL {
					if enemies[i].selatk {
						enemies[i].selatk = false
					} else {
						for i := range enemies {
							enemies[i].selatk = false
						}
						enemies[i].selatk = true
						selatkY = enemies[i].rC.Y
						selatkY -= UNIT
						selatkY2 = selatkY - selatkYchange
					}

				}
			} else {
				enemies[i].sel = false
			}

		} else {
			siz := UNIT * 10
			enemies[i].deth = dAnimRecONCE(enemies[i].deth, RECCNT(qRecCNT(enemies[i].rC), siz, siz))
			if enemies[i].deth.off {
				enemies[i].dead = true
				clear = true
			}
		}

	}
	if clear {
		enemies = REMENMDEAD(enemies)
	}
}

func dDRAWING() {
	drawingANIM = dAnimRecColorONCE(drawingANIM, qResizeRecTHREEQUARTER(deckREC), CA(sepia, 100))
	drawingIM = dIMrotatingColorSHADOW(drawingIM, drawingREC, UNIT/4, UNIT/3, sepiaDRK, CA(z.Black, 180))
	drawingREC.X += UNIT / 2
	drawingREC.Y += UNIT / 10
}

func dENEMYSTATS(e ENEMY) { //MARK: ENEMY STATS
	x := levbgREC.X + levbgREC.Width + UNIT/2
	y := levbgREC.Y
	siz := float32(2)
	dTXTfont2XY(e.nm, x, y, siz, sepiaDRK2)
	y += qTXTheight(FONT2, siz)
	siz = 1.2
	dTXTfont2XY("HP : "+fmt.Sprint(e.hp), x, y, siz, sepiaDRK2)
	y += qTXTheight(FONT2, siz)
	if e.immBURN {
		dTXTfont2XY("Immune Burn", x, y, siz, sepiaDRK2)
		y += qTXTheight(FONT2, siz)
	}
	if e.immSTUN {
		dTXTfont2XY("Immune Stun", x, y, siz, sepiaDRK2)
		y += qTXTheight(FONT2, siz)
	}
	if e.immFREEZE {
		dTXTfont2XY("Immune Freeze", x, y, siz, sepiaDRK2)
		y += qTXTheight(FONT2, siz)
	}
	if e.immELECTRIC {
		dTXTfont2XY("Immune Electric", x, y, siz, sepiaDRK2)
		y += qTXTheight(FONT2, siz)
	}
	if e.immPOISON {
		dTXTfont2XY("Immune Poison", x, y, siz, sepiaDRK2)
		y += qTXTheight(FONT2, siz)
	}

}
func dCARDSTATS(c CARD) {
	x := levbgREC.X + levbgREC.Width + UNIT/2
	y := levbgREC.Y
	siz := float32(2)
	dTXTfont2XY(c.nm, x, y, siz, sepiaDRK2)
	h := qTXTheight(FONT2, siz)
	y += h
	siz = float32(1.4)
	if c.cost > 0 {
		dTXTfont2XY("Cost Turns: "+fmt.Sprint(c.cost), x, y, siz, sepiaDRK2)
	} else if c.costMana > 0 {
		dTXTfont2XY("Cost Mana: "+fmt.Sprint(c.costMana), x, y, siz, sepiaDRK2)
	}
	h = qTXTheight(FONT2, siz)
	y += h
	siz = float32(1.2)
	h = qTXTheight(FONT2, siz)
	if c.atk != 0 {
		dTXTfont2XY("Attack +"+fmt.Sprint(c.atk), x, y, siz, sepiaDRK2)
		y += h
	}
	if c.def != 0 {
		dTXTfont2XY("Defend +"+fmt.Sprint(c.def), x, y, siz, sepiaDRK2)
		y += h
	}
	if c.poison != 0 {
		dTXTfont2XY("Poison +"+fmt.Sprint(c.poison), x, y, siz, sepiaDRK2)
		y += h
	}
	if c.burn != 0 {
		dTXTfont2XY("Burn +"+fmt.Sprint(c.burn), x, y, siz, sepiaDRK2)
		y += h
	}
	if c.electric != 0 {
		dTXTfont2XY("Electric +"+fmt.Sprint(c.electric), x, y, siz, sepiaDRK2)
		y += h
	}
	if c.freez != 0 {
		dTXTfont2XY("Freeze +"+fmt.Sprint(c.freez), x, y, siz, sepiaDRK2)
		y += h
	}
	if c.evade != 0 {
		dTXTfont2XY("Evade +"+fmt.Sprint(c.evade), x, y, siz, sepiaDRK2)
		y += h
	}
	if c.thorns != 0 {
		dTXTfont2XY("Thorns +"+fmt.Sprint(c.thorns), x, y, siz, sepiaDRK2)
		y += h
	}
	if c.defTURN != 0 {
		dTXTfont2XY("Absorb Damage +"+fmt.Sprint(c.defTURN), x, y, siz, sepiaDRK2)
		y += h
	}
	if c.draw != 0 {
		dTXTfont2XY("Draw Card +"+fmt.Sprint(c.draw), x, y, siz, sepiaDRK2)
		y += h
	}
	if c.desc != "" {
		dTXTfont2XY(c.desc, x, y, siz, sepiaDRK2)
		y += h
	}

	if c.ground && c.air {
		dTXTfont2XY("Ground & Air", x, y, siz, sepiaDRK2)
	} else if c.ground && !c.air {
		dTXTfont2XY("Ground Only", x, y, siz, sepiaDRK2)
	} else if !c.ground && c.air {
		dTXTfont2XY("Air Only", x, y, siz, sepiaDRK2)
	}
	y += h
	if c.all {
		dTXTfont2XY("All Enemies", x, y, siz, sepiaDRK2)
	}

}
func dCARDS() { //MARK: CARDS HAND
	spc := UNIT / 5
	w := levbgREC.Width - UNIT*2
	var wCard float32
	if len(pl.hand) >= 7 {
		wCard = w / float32(len(pl.hand))
	} else {
		wCard = w / 7
	}
	wCard -= spc / 2
	w2 := float32(len(pl.hand)) * (wCard + spc)
	x := CNT.X - w2/2
	h := qHeightProportional(ETC[0].r.Width, ETC[0].r.Height, wCard)
	for i := range len(pl.hand) {
		if !pl.hand[i].off {
			//CARDS IN HAND
			cardREC = z.NewRectangle(x, float32(SCRH)-(h+UNIT/4), wCard, h)
			v2 := qRecCNT(cardREC)
			v2.Y += cardREC.Height / 10
			rIM := RECCNT(v2, cardREC.Width-((cardREC.Width/20)*9), cardREC.Width-((cardREC.Width/20)*9))
			//CARD PLAYED
			if pl.hand[i].played { //PLAY CARD ANIM
				if !pl.hand[i].offPLAYED {
					pl.hand[i].offPLAYED = true
					PLAYCARD(pl.hand[i])
				}
				dIMAcolor(ETC[0], cardREC, cardA, sepia)
				r2 := RECCNT(qRecCNT(cardREC), cardREC.Width, cardREC.Width)
				r2 = qResizeRecOffsetLRGR(r2, UNIT*4)
				r2.Y -= r2.Height / 2
				playcardANM = dAnimRecShadowONCE(playcardANM, r2, CA(z.Black, 180), UNIT/5)
				cardA -= 10
				if playcardANM.off {
					pl.hand[i].off = true
					playcardANM = RESETANIM(playcardANM)
					cardA = 255
				}
			} else { //HAND CARDS
				if cMSREC(cardREC) {
					cardREC.Y -= UNIT / 2
					rIM.Y -= UNIT / 2
					if noTurnsT == 0 && notselectedT == 0 && nomanaT == 0 {
						dCARDSTATS(pl.hand[i])
					}
					if MSL {
						if pl.hand[i].costMana > 0 {
							if pl.hand[i].costMana <= pl.mana {
								if pl.hand[i].all {
									pl.mana -= pl.hand[i].costMana
									pl.discard = append(pl.discard, pl.hand[i])
									pl.hand[i].played = true
								} else {
									if cENEMYSEL() {
										pl.mana -= pl.hand[i].costMana
										pl.discard = append(pl.discard, pl.hand[i])
										pl.hand[i].played = true
									} else {
										notselectedT = int(FPS)
									}
								}
							} else {
								nomanaT = int(FPS)
							}
						} else {
							if pl.hand[i].cost <= pl.turns {
								if pl.hand[i].all {
									pl.turns -= pl.hand[i].cost
									pl.discard = append(pl.discard, pl.hand[i])
									pl.hand[i].played = true
								} else {
									if cENEMYSEL() {
										pl.turns -= pl.hand[i].cost
										pl.discard = append(pl.discard, pl.hand[i])
										pl.hand[i].played = true
									} else {
										notselectedT = int(FPS)
									}
								}
							} else {
								noTurnsT = int(FPS)
							}
						}
					}
				}
				pl.hand[i].cnt = mCARDrecCNT()
				dRECSHADOWONLY(cardREC, UNIT/7, z.Black, 220)
				dIMcolor(ETC[0], cardREC, sepia)
				dRECLINE(cardREC, 2, z.Black)
				dTXTcntRecWfont2(pl.hand[i].nm, cardREC, cardREC.Height/5, 1, z.Black)
				dIMcolorSHADOW(pl.hand[i].im, rIM, sepiaDRK2, CA(z.Black, 180), UNIT/8)
				//CARD ICONS
				siz := UNIT
				r := z.NewRectangle(cardREC.X+cardREC.Width-(siz+siz/5), cardREC.Y-(siz/20)*4, siz, siz)
				if pl.hand[i].costMana > 0 { //TURNS MANA
					dIMcolorSHADOW(ETC[34], r, z.DarkBlue, CA(z.Black, 120), UNIT/12)
					dTXTfont1XY(fmt.Sprint(pl.hand[i].costMana), r.X+(siz/20)*9, r.Y+(siz/20)*2, 1, z.White)
				} else {
					dIMcolorSHADOW(ETC[34], r, z.DarkGreen, CA(z.Black, 120), UNIT/12)
					dTXTfont1XY(fmt.Sprint(pl.hand[i].cost), r.X+(siz/20)*9, r.Y+(siz/20)*2, 1, z.White)
				}

				siz = (UNIT / 10) * 6
				r = z.NewRectangle(pl.hand[i].cnt.X-siz/2, cardREC.Y+cardREC.Height-(siz/10)*7, siz, siz)
				spc := UNIT / 10
				if pl.hand[i].def > 0 { //DEF
					dRECLINE(r, 2, z.Black)
					dREC(r, CA(z.Black, 100))
					dIMcolor(ETC[33], r, z.Blue)
					r.X += siz + spc
				}
				if pl.hand[i].burn > 0 { //BURN
					dRECLINE(r, 2, z.Black)
					dREC(r, CA(z.Black, 100))
					dIMcolor(ETC[31], r, z.Orange)
					r.X += siz + spc
				}
				if pl.hand[i].electric > 0 { //ELECTRIC
					dRECLINE(r, 2, z.Black)
					dREC(r, CA(z.Black, 100))
					dIMcolor(ETC[30], r, z.SkyBlue)
					r.X += siz + spc
				}
				if pl.hand[i].poison > 0 { //POISON
					dRECLINE(r, 2, z.Black)
					dREC(r, CA(z.Black, 100))
					dIMcolor(ETC[32], r, z.DarkGreen)
					r.X += siz + spc
				}
			}

			//MOVE X TO NEXT CARD
			x += wCard + spc

		}
	}
	pl.hand = REMCARDOFF(pl.hand)

}
