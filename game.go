package main

import (
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
	z "github.com/gen2brain/raylib-go/raylib"
)

var (
	levNUM              = 1
	bg, levbg, playerIM []IM
	levbgIM             IM
	sepia               = CRGB(255, 238, 185)
	sepiaDRK            = CRGB(255, 221, 117)
	sepiaDRK2           = CRGB(255, 205, 55)
	handNUM             = 5
	cardW               float32
	cardlist            []CARD
	enemies, enemyLIST  []ENEMY
	cardA               uint8 = 255

	batLR, fairyLR, batORfairy bool
	batT, fairyT               int
	pl                         PLAYER
	playerLIST                 []PLAYER

	levmapREC, levbgREC, deckREC, cardREC, drawingREC, batREC, fairyREC, deckViewREC z.Rectangle

	drawingANIM, bat, fairy, shopkeeper, playcardANM, nexthandANM, pldeathANM ANIM

	enmDeathANIMS, enmhitANMS []ANIM

	viewDeck, viewDiscard, playerdeathONOFF, enmTURN bool

	enmTurnCount, noTurnsT, notselectedT, nomanaT int

	//LEVMAP
	levmapstartR         z.Rectangle
	levmapobjRECS        []z.Rectangle
	levobjsLIST, levobjs []LEVOBJ
	bgIMmap              []IM
	pllevmapNUM          int
	moveplmap            bool
	plmapV2              z.Vector2

	//ENEMY MODS
	dexENM int

	//DRAWING
	drawingT  int
	drawingIM IM
)

type LEVOBJ struct {
	r       z.Rectangle
	nm      string
	im      IM
	numtype int
}
type PLAYER struct {
	im                  IM
	nm, desc            string
	hand, deck, discard []CARD

	hp, hpmax, mana, turns, turnsMAX, drawNUM int
}

type ENEMY struct {
	a, a2, deth, hit ANIM
	r, rC            z.Rectangle
	nm, imFACES      string

	hpHIT, played, fly, atk, sel, selatk, off, dead, immPOISON, immBURN, immFREEZE, immSTUN, immELECTRIC bool

	pos, hp, hpmax, xp, poisonAMOUNT, burnAMOUNT, stunLEN, frozen, def int

	poison, burn []int
}
type CARD struct {
	nm, desc string
	im       IM
	cnt      z.Vector2

	played, off, offPLAYED bool

	cost, costMana, atk, def, freez, poison, electric, burn, evade, defTURN, thorns, draw int

	evadeT, rainT, dexENM, manaPLUS, turnPLUS, burnAMOUNT, poisonAMOUNT, stunAMOUNT int

	ground, air, all bool
}

func INITGAME() { //MARK: INIT GAME

	rl.SetExitKey(rl.KeyEnd) //DELETE
	//LEV BG
	w := float32(SCRH) - UNIT*2
	h := w
	levbgREC = RECCNT(CNT, w, h)
	levbg = mIMSheetFiles("im/levbg")
	w = UNIT * 7
	h = qHeightProportional(ETC[2].r.Width, ETC[2].r.Height, w)
	deckREC = z.NewRectangle(levbgREC.X-w+(w/3), levbgREC.Y+levbgREC.Height+UNIT/2-h, w, h)
	siz := UNIT * 2
	deckViewREC = z.NewRectangle(deckREC.X, deckREC.Y+deckREC.Height-siz, siz, siz)
	cardW = levbgREC.Width / 8
	h = qHeightProportional(ETC[0].r.Width, ETC[0].r.Height, cardW)
	cardREC = z.NewRectangle(CNT.X-cardW/2, levbgREC.Y+levbgREC.Height-(h-UNIT/2), cardW, h)
	bat = mAnimXY("im/bat.png", 0, 0, 64, 0, 4, 1, 12)
	siz = UNIT * 4
	batREC = z.NewRectangle(-siz, UNIT*3, siz, siz)
	batT = int(FPS) * RINT(30, 120)
	fairy = mAnimXY("im/fairy.png", 0, 0, 32, 0, 8, 1, 12)
	fairyREC = z.NewRectangle(-siz, UNIT*3, siz/3, siz/3)
	fairyT = int(FPS) * RINT(30, 120)
	batORfairy = FLIPCOIN()
	shopkeeper = mAnimXY("im/shopkeeper.png", 0, 0, 64, 0, 4, 1, 10)
	playerIM = mIMSheetFiles("im/characters")
	playcardANM = mAnimXY("im/fx/002playcard.png", 0, 0, 64, 0, 11, 1, 50)
	enmDeathANIMS = mAnimSheetFiles1LINEH("im/fx2", 64, 64, 18)
	pldeathANM = mAnimXY("im/fx/004playerdeath.png", 0, 0, 64, 0, 15, 1, 24)
	nexthandANM = mAnimXY("im/fx/003nexthand.png", 0, 0, 64, 0, 12, 1, 44)
	enmhitANMS = mAnimSheetFiles1LINEH("im/fx3", 64, 64, 50)

	mANIM()
	mCARDS()
	mLEV()
	mENEMYLIST()
	mENEMIES()
	mPLAYERLIST()
	mPLAYER()
	mLEVOBJS()
	mMAP()
}

func mMAP() { //MARK: MAKE MAP
	siz := UNIT / 2
	levmapREC = z.NewRectangle(siz, siz, float32(SCRW)-siz*2, float32(SCRH)-siz*2) //BG IM REC
	//MAP OBJS
	siz = UNIT * 2
	v2 := z.NewVector2(levmapREC.X+UNIT*3+UNIT/2, levmapREC.Y+levmapREC.Height/2)
	plmapV2 = v2
	levmapstartR = RECCNT(v2, siz, siz)
	divW := levmapREC.Width / 9
	divY := levmapREC.Height / 4.5
	v2.X += divW

	siz = UNIT * 3

	//ROW 1
	l := levobjsLIST[0]
	l.r = RECCNT(v2, siz, siz) //CNT
	levobjs = append(levobjs, l)

	v2.Y -= divY

	l = levobjsLIST[0]
	l.r = RECCNT(v2, siz, siz) //TOP
	levobjs = append(levobjs, l)

	v2.Y += divY * 2

	l = levobjsLIST[0]
	l.r = RECCNT(v2, siz, siz) //BOTTOM
	levobjs = append(levobjs, l)

	//ROW 2
	v2.X += divW
	v2.Y -= divY

	l = levobjsLIST[0]
	l.r = RECCNT(v2, siz, siz) //CNT
	levobjs = append(levobjs, l)

	v2.Y -= divY

	l = levobjsLIST[0]
	l.r = RECCNT(v2, siz, siz) //TOP
	levobjs = append(levobjs, l)

	v2.Y += divY * 2

	l = levobjsLIST[0]
	l.r = RECCNT(v2, siz, siz) //BOTTOM
	levobjs = append(levobjs, l)

	//ROW 3
	v2.X += divW
	v2.Y -= divY

	l = levobjsLIST[0]
	l.r = RECCNT(v2, siz, siz) //CNT
	levobjs = append(levobjs, l)

	v2.Y -= divY

	l = levobjsLIST[0]
	l.r = RECCNT(v2, siz, siz) //TOP
	levobjs = append(levobjs, l)

	v2.Y += divY * 2

	l = levobjsLIST[0]
	l.r = RECCNT(v2, siz, siz) //BOTTOM
	levobjs = append(levobjs, l)

	//CENTER EVENT
	v2.X += divW
	v2.Y -= divY
	l = levobjsLIST[1]
	l.r = RECCNT(v2, siz, siz) //CNT
	levobjs = append(levobjs, l)

	//ROW 4
	v2.X += divW

	l = levobjsLIST[0]
	l.r = RECCNT(v2, siz, siz) //CNT
	levobjs = append(levobjs, l)

	v2.Y -= divY

	l = levobjsLIST[0]
	l.r = RECCNT(v2, siz, siz) //TOP
	levobjs = append(levobjs, l)

	v2.Y += divY * 2

	l = levobjsLIST[0]
	l.r = RECCNT(v2, siz, siz) //BOTTOM
	levobjs = append(levobjs, l)

	//ROW 5
	v2.X += divW
	v2.Y -= divY

	l = levobjsLIST[0]
	l.r = RECCNT(v2, siz, siz) //CNT
	levobjs = append(levobjs, l)

	v2.Y -= divY

	l = levobjsLIST[0]
	l.r = RECCNT(v2, siz, siz) //TOP
	levobjs = append(levobjs, l)

	v2.Y += divY * 2

	l = levobjsLIST[0]
	l.r = RECCNT(v2, siz, siz) //BOTTOM
	levobjs = append(levobjs, l)

	//BOSS
	siz = UNIT * 4
	v2.X += divW
	v2.Y -= divY

	l = levobjsLIST[6]
	l.r = RECCNT(v2, siz, siz) //CNT
	levobjs = append(levobjs, l)

	//ADD DIFFERENT LEV OBJS
	var row1, row2, row3 bool
	for {
		count := 0
		if !row1 {
			if FLIPCOIN() {
				num := RPICKINT([]int{0, 1, 2})
				num2 := RPICKINT([]int{4, 5})
				levobjs[num] = copylevobj(levobjs[num], num2)
				row1 = true
				count++
			}
		} else {
			count++
		}
		if !row2 {
			if FLIPCOIN() {
				num := RPICKINT([]int{10, 11, 12})
				num2 := RPICKINT([]int{2, 3})
				levobjs[num] = copylevobj(levobjs[num], num2)
				row2 = true
				count++
			}
		} else {
			count++
		}
		if !row3 {
			if FLIPCOIN() {
				num := RPICKINT([]int{13, 14, 15})
				num2 := RPICKINT([]int{7, 9, 10})
				levobjs[num] = copylevobj(levobjs[num], num2)
				row3 = true
				count++
			}
		} else {
			count++
		}
		if count >= 2 {
			break
		}
	}

	var row4, row5 bool
	for {
		if !row4 {
			if FLIPCOIN() {
				num := RPICKINT([]int{0, 1, 2})
				num2 := RPICKINT([]int{2, 3, 4, 5, 7, 9, 10})
				levobjs[num] = copylevobj(levobjs[num], num2)
				row4 = true
			}
		}
		if !row5 {
			if FLIPCOIN() {
				num := RPICKINT([]int{3, 4, 5})
				num2 := RPICKINT([]int{2, 3, 4, 5, 7, 9, 10})
				levobjs[num] = copylevobj(levobjs[num], num2)
				row5 = true
			}
		}
		if row4 || row5 {
			break
		}
	}

}
func copylevobj(l LEVOBJ, numtype int) LEVOBJ {
	l.nm = levobjsLIST[numtype].nm
	l.im = levobjsLIST[numtype].im
	l.numtype = levobjsLIST[numtype].numtype
	return l
}
func mLEVOBJS() { //MARK: MAKE LEVOBJS
	l := LEVOBJ{}
	l.numtype = 0 //ENCOUNTER
	l.nm = "Encounter"
	l.im = ETC[39]
	levobjsLIST = append(levobjsLIST, l)

	l = LEVOBJ{}
	l.numtype = 1 //GEM SHOP
	l.nm = "Gem Shop"
	l.im = ETC[48]
	levobjsLIST = append(levobjsLIST, l)

	l = LEVOBJ{}
	l.numtype = 2 //WEAPON SHOP
	l.nm = "Weapon Shop"
	l.im = ETC[43]
	levobjsLIST = append(levobjsLIST, l)

	l = LEVOBJ{}
	l.numtype = 3 //CARD SHOP
	l.nm = "Card Shop"
	l.im = ETC[35]
	levobjsLIST = append(levobjsLIST, l)

	l = LEVOBJ{}
	l.numtype = 4 //EVENT
	l.nm = "Event"
	l.im = ETC[36]
	levobjsLIST = append(levobjsLIST, l)

	l = LEVOBJ{}
	l.numtype = 5 //TREASURE
	l.nm = "Treasure"
	l.im = ETC[44]
	levobjsLIST = append(levobjsLIST, l)

	l = LEVOBJ{}
	l.numtype = 6 //BOSS
	l.nm = "Boss Encounter"
	l.im = ETC[41]
	levobjsLIST = append(levobjsLIST, l)

	l = LEVOBJ{}
	l.numtype = 7 //SHRINE
	l.nm = "Shrine"
	l.im = ETC[57]
	levobjsLIST = append(levobjsLIST, l)

	l = LEVOBJ{}
	l.numtype = 8 //UFO
	l.nm = "UFO"
	l.im = ETC[63]
	levobjsLIST = append(levobjsLIST, l)

	l = LEVOBJ{}
	l.numtype = 9 //Potion Shop
	l.nm = "Potion Shop"
	l.im = ETC[64]
	levobjsLIST = append(levobjsLIST, l)

	l = LEVOBJ{}
	l.numtype = 10 //Healer
	l.nm = "Healer"
	l.im = ETC[65]
	levobjsLIST = append(levobjsLIST, l)

}
func cENEMYSEL() bool {
	sel := false
	for i := range enemies {
		if enemies[i].selatk {
			sel = true
			break
		}
	}
	return sel
}

func PLAYCARD(c CARD) { //MARK: PLAY CARD

	if c.atk != 0 { //ATK
		if c.all {
			for i := range enemies {
				if enemies[i].def != 0 {
					if enemies[i].def >= c.atk {
						enemies[i].def -= c.atk
					} else {
						diff := c.atk - enemies[i].def
						enemies[i].def = 0
						enemies[i].hp -= diff
						enemies[i].hpHIT = true
					}
				} else {
					enemies[i].hp -= c.atk
					enemies[i].hpHIT = true
				}
			}
		} else {
			for i := range enemies {
				if enemies[i].selatk {
					if enemies[i].def != 0 {
						if enemies[i].def >= c.atk {
							enemies[i].def -= c.atk
						} else {
							diff := c.atk - enemies[i].def
							enemies[i].def = 0
							enemies[i].hp -= diff
							enemies[i].hpHIT = true
						}
					} else {
						enemies[i].hp -= c.atk
						enemies[i].hpHIT = true

					}
				}
			}
		}
	}
	if c.poison != 0 { //POISON
		if c.all {
			for i := range enemies {
				if !enemies[i].immPOISON {
					enemies[i].poison = append(enemies[i].poison, c.poison)
					enemies[i].poisonAMOUNT += c.poisonAMOUNT
				}
			}
		} else {
			for i := range enemies {
				if enemies[i].selatk {
					if !enemies[i].immPOISON {
						enemies[i].poison = append(enemies[i].poison, c.poison)
						enemies[i].poisonAMOUNT += c.poisonAMOUNT
					}
				}
			}
		}
	}

	if c.burn != 0 { //BURN
		if c.all {
			for i := range enemies {
				if !enemies[i].immBURN {
					enemies[i].burn = append(enemies[i].burn, c.burn)
					enemies[i].burnAMOUNT += c.burnAMOUNT
				}
			}
		} else {
			for i := range enemies {
				if enemies[i].selatk {
					if !enemies[i].immBURN {
						enemies[i].burn = append(enemies[i].burn, c.burn)
						enemies[i].burnAMOUNT += c.burnAMOUNT
					}
				}
			}
		}
	}

	if c.stunAMOUNT != 0 { //STUN
		if c.all {
			for i := range enemies {
				if !enemies[i].immSTUN {
					enemies[i].stunLEN += c.stunAMOUNT
				}
			}
		} else {
			for i := range enemies {
				if enemies[i].selatk {
					if !enemies[i].immSTUN {
						enemies[i].stunLEN += c.stunAMOUNT
					}
				}
			}
		}
	}

	//DESTROY ENEMY HP 0
	for i := range enemies {
		if enemies[i].hp <= 0 {
			enemies[i].off = true
			enemies[i].deth = enmDeathANIMS[RINT(0, len(enmDeathANIMS))]
		}
	}
}
func DRAWCARD() { //MARK: DRAW CARD
	choose := RINT(0, len(pl.deck))
	pl.hand = append(pl.hand, pl.deck[choose])
	pl.deck = REM(pl.deck, choose)
	if len(pl.deck) == 0 {
		DISCARD2DECK()
	}
}
func DISCARD2DECK() {
	pl.deck = append(pl.deck, pl.discard...)
	pl.discard = nil
}

func mPLAYERLIST() { //MARK: MAKE PLAYER LIST
	p := PLAYER{} //0 Horknee
	p.im = playerIM[0]
	p.nm = "Horknee"
	p.desc = "some text here"
	p.hp = 10
	p.hpmax = pl.hp
	p.mana = 3
	p.turns = 3
	p.turnsMAX = p.turns
	p.drawNUM = 5
	playerLIST = append(playerLIST, p)

	p.im = playerIM[1] //1 Sharpee
	p.nm = "Sharpee"
	p.desc = "some text here"
	playerLIST = append(playerLIST, p)

	p.im = playerIM[2] //2 Mr Snuggles
	p.nm = "Mr Snuggles"
	p.desc = "some text here"
	playerLIST = append(playerLIST, p)

	p.im = playerIM[3] //3 Drakola
	p.nm = "Drakola"
	p.desc = "some text here"
	playerLIST = append(playerLIST, p)

	p.im = playerIM[4] //4 Zorg
	p.nm = "Zorg"
	p.desc = "some text here"
	playerLIST = append(playerLIST, p)

	p.im = playerIM[5] //5 Woofles
	p.nm = "Woofles"
	p.desc = "some text here"
	playerLIST = append(playerLIST, p)

	p.im = playerIM[6] //6 Panpan
	p.nm = "Panpan"
	p.desc = "some text here"
	playerLIST = append(playerLIST, p)

	p.im = playerIM[7] //7 Constipo
	p.nm = "Constipo"
	p.desc = "some text here"
	playerLIST = append(playerLIST, p)

	p.im = playerIM[8] //8 Vlambo
	p.nm = "Vlambo"
	p.desc = "some text here"
	playerLIST = append(playerLIST, p)

	p.im = playerIM[9] //9 Marg
	p.nm = "Marg"
	p.desc = "some text here"
	playerLIST = append(playerLIST, p)

	p.im = playerIM[10] //10 Zagzig
	p.nm = "Zagzig"
	p.desc = "some text here"
	playerLIST = append(playerLIST, p)

	p.im = playerIM[11] //11 Eyeskeem
	p.nm = "Eyeskeem"
	p.desc = "some text here"
	playerLIST = append(playerLIST, p)

	p.im = playerIM[12] //12 Scruffles
	p.nm = "Scruffles"
	p.desc = "some text here"
	playerLIST = append(playerLIST, p)

	p.im = playerIM[13] //13 Gadget
	p.nm = "Gadget"
	p.desc = "some text here"
	playerLIST = append(playerLIST, p)

	p.im = playerIM[14] //14 Lizo
	p.nm = "Lizo"
	p.desc = "some text here"
	playerLIST = append(playerLIST, p)

	p.im = playerIM[15] //15 Buggo
	p.nm = "Buggo"
	p.desc = "some text here"
	playerLIST = append(playerLIST, p)

	p.im = playerIM[16] //16 Daisy
	p.nm = "Daisy"
	p.desc = "some text here"
	playerLIST = append(playerLIST, p)

	p.im = playerIM[17] //17 Rex
	p.nm = "Rex"
	p.desc = "some text here"
	playerLIST = append(playerLIST, p)

	p.im = playerIM[18] //18 Izzy
	p.nm = "Izzy"
	p.desc = "some text here"
	playerLIST = append(playerLIST, p)
}
func mPLAYER() { //MARK: MAKE PLAYER
	pl = playerLIST[RINT(0, len(playerLIST))]
	pl.deck = append(pl.deck, cardlist[0], cardlist[0], cardlist[0], cardlist[1], cardlist[1], cardlist[4], cardlist[8], cardlist[11], cardlist[11], cardlist[11], cardlist[12], cardlist[14], cardlist[15], cardlist[24])

	CHOOSECARDSINITIAL()
}
func CHOOSECARDSINITIAL() { //MARK: CHOOSE CARDS INITIAL
	pl.hand = nil
	//CARDS
	for range pl.drawNUM {
		choose := RINT(0, len(pl.deck))
		pl.hand = append(pl.hand, pl.deck[choose])
		pl.deck = REM(pl.deck, choose)
	}
}
func DRAWNEXTHAND() { //MARK: DRAW NEXT HAND
	for range pl.drawNUM {
		if len(pl.deck) > 1 {
			choose := RINT(0, len(pl.deck))
			pl.hand = append(pl.hand, pl.deck[choose])
			pl.deck = REM(pl.deck, choose)
		} else if len(pl.deck) == 1 {
			pl.hand = append(pl.hand, pl.deck[0])
			pl.deck = REM(pl.deck, 0)
		} else {
			DISCARD2DECK()
			choose := RINT(0, len(pl.deck))
			pl.hand = append(pl.hand, pl.deck[choose])
			pl.deck = REM(pl.deck, choose)
		}
	}
}
func mCARDS() { //MARK: MAKE CARDS
	c := CARD{}
	c.nm = "Jab" //0 KNIFE JAB
	c.atk = 1
	c.cost = 1
	c.ground = true
	c.im = ETC[10]
	cardlist = append(cardlist, c)

	c = CARD{}
	c.nm = "Block" //1 BLOCK
	c.def = 1
	c.cost = 1
	c.ground = true
	c.im = ETC[14]
	cardlist = append(cardlist, c)

	c = CARD{}
	c.nm = "Block Above" //2 BLOCK ABOVE
	c.def = 1
	c.cost = 1
	c.air = true
	c.im = ETC[10]
	cardlist = append(cardlist, c)

	c = CARD{}
	c.nm = "Shield Wall" //3 SHIELD WALL
	c.def = 2
	c.cost = 2
	c.air = true
	c.ground = true
	c.im = ETC[10]
	cardlist = append(cardlist, c)

	c = CARD{}
	c.nm = "Duck" //4 DUCK
	c.evade = 2
	c.cost = 2
	c.air = true
	c.ground = true
	c.im = ETC[19]
	cardlist = append(cardlist, c)

	c = CARD{}
	c.nm = "Dodge" //5 DODGE
	c.evade = 1
	c.cost = 1
	c.ground = true
	c.im = ETC[10]
	cardlist = append(cardlist, c)

	c = CARD{}
	c.nm = "Smoke Bomb" //6 SMOKE
	c.evade = 1
	c.evadeT = 2
	c.dexENM = -1
	c.cost = 1
	c.all = true
	c.ground = true
	c.air = true
	c.im = ETC[18]
	cardlist = append(cardlist, c)

	c = CARD{}
	c.nm = "Rain" //7 RAIN
	c.costMana = 1
	c.rainT = 2
	c.dexENM = -2
	c.all = true
	c.ground = true
	c.air = true
	c.im = ETC[10]
	cardlist = append(cardlist, c)

	c = CARD{}
	c.nm = "Groin Kick" //8 GROIN KICK
	c.atk = 1
	c.dexENM = -2
	c.stunAMOUNT = 1
	c.cost = 1
	c.ground = true
	c.im = ETC[17]
	cardlist = append(cardlist, c)

	c = CARD{}
	c.nm = "Marshmallow" //9 MARSHMALLOW
	c.defTURN = 1
	c.cost = 2
	c.air = true
	c.ground = true
	c.im = ETC[20]
	cardlist = append(cardlist, c)

	c = CARD{}
	c.nm = "Pineapple" //10 PINEAPPLE
	c.costMana = 1
	c.thorns = 1
	c.air = true
	c.ground = true
	c.im = ETC[16]
	cardlist = append(cardlist, c)

	c = CARD{}
	c.nm = "Arrow" //11 ARROW
	c.atk = 1
	c.cost = 1
	c.air = true
	c.im = ETC[10]
	cardlist = append(cardlist, c)

	c = CARD{}
	c.nm = "Bolt" //12 BOLT
	c.costMana = 2
	c.electric = 1
	c.all = true
	c.air = true
	c.ground = true
	c.im = ETC[10]
	cardlist = append(cardlist, c)

	c = CARD{}
	c.nm = "Vine" //13 VINE
	c.costMana = 1
	c.stunAMOUNT = 2
	c.ground = true
	c.im = ETC[10]
	cardlist = append(cardlist, c)

	c = CARD{}
	c.nm = "Throw Vial" //14 VIAL
	c.cost = 1
	c.poison = 3
	c.poisonAMOUNT = 1
	c.ground = true
	c.air = true
	c.im = ETC[10]
	cardlist = append(cardlist, c)

	c = CARD{}
	c.nm = "Slash" //15 SLASH
	c.cost = 3
	c.atk = 2
	c.all = true
	c.ground = true
	c.im = ETC[13]
	cardlist = append(cardlist, c)

	c = CARD{}
	c.nm = "Rummage" //16 RUMMAGE
	c.cost = 1
	c.draw = 1
	c.desc = "from discard"
	c.im = ETC[10]
	cardlist = append(cardlist, c)

	c = CARD{}
	c.nm = "Card +" //17 CARD+
	c.cost = 1
	c.draw = 1
	c.desc = "from deck"
	c.im = ETC[10]
	cardlist = append(cardlist, c)

	c = CARD{}
	c.nm = "Card ++" //18 CARD++
	c.cost = 2
	c.draw = 2
	c.desc = "from deck"
	c.im = ETC[10]
	cardlist = append(cardlist, c)

	c = CARD{}
	c.nm = "Mana +" //19 MANA+
	c.cost = 1
	c.manaPLUS = 1
	c.im = ETC[10]
	cardlist = append(cardlist, c)

	c = CARD{}
	c.nm = "Mana ++" //20 MANA++
	c.cost = 2
	c.manaPLUS = 2
	c.im = ETC[10]
	cardlist = append(cardlist, c)

	c = CARD{}
	c.nm = "Turn +" //21 TURN+
	c.cost = 1
	c.turnPLUS = 1
	c.im = ETC[10]
	cardlist = append(cardlist, c)

	c = CARD{}
	c.nm = "Turn ++" //22 TURN++
	c.cost = 2
	c.turnPLUS = 2
	c.im = ETC[10]
	cardlist = append(cardlist, c)

	c = CARD{}
	c.nm = "Fireball" //23 FIREBALL
	c.costMana = 3
	c.burn = 2
	c.burnAMOUNT = 1
	c.all = true
	c.ground = true
	c.air = true
	c.im = ETC[15]
	cardlist = append(cardlist, c)

	c = CARD{}
	c.nm = "Firework" //24 FIREWORK
	c.cost = 1
	c.burn = 1
	c.burnAMOUNT = 1
	c.ground = true
	c.air = true
	c.im = ETC[12]
	cardlist = append(cardlist, c)

	c = CARD{}
	c.nm = "Hot Coals" //25 HOT COALS
	c.cost = 3
	c.burn = 3
	c.burnAMOUNT = 1
	c.ground = true
	c.all = true
	c.im = ETC[10]
	cardlist = append(cardlist, c)

	c = CARD{}
	c.nm = "Rotten Food" //26 ROTTEN FOOD
	c.cost = 1
	c.poison = 1
	c.poisonAMOUNT = 1
	c.ground = true
	c.im = ETC[10]
	cardlist = append(cardlist, c)

	c = CARD{}
	c.nm = "Urchin Spine" //27 URCHIN SPINE
	c.cost = 1
	c.poison = 2
	c.poisonAMOUNT = 2
	c.ground = true
	c.im = ETC[10]
	cardlist = append(cardlist, c)

	c = CARD{}
	c.nm = "Poison Dart" //28 POISON DART
	c.cost = 1
	c.poison = 2
	c.poisonAMOUNT = 1
	c.ground = true
	c.air = true
	c.im = ETC[10]
	cardlist = append(cardlist, c)

	c = CARD{}
	c.nm = "Gas Cloud" //29 GAS CLOUD
	c.cost = 2
	c.poison = 2
	c.poisonAMOUNT = 1
	c.all = true
	c.ground = true
	c.air = true
	c.im = ETC[10]
	cardlist = append(cardlist, c)
}

// MARK: MAKE
func mENEMYLIST() { //MARK: MAKE ENEMY LIST

	e := ENEMY{} //0 Slime
	e.nm = "Slime"
	e.imFACES = "c"
	e.hp = 4
	e.hpmax = e.hp
	e.a = mAnimXY("im/enemies/002slimeIDLE.png", 0, 0, 64, 0, 6, 1, 7)
	e.a2 = mAnimXY("im/enemies/002slimeATK.png", 0, 0, 64, 0, 10, 1, 10)
	w := levbgREC.Width / 2
	e.r = z.NewRectangle(0, 0, w, w)
	e.rC = qResizeRecOffsetSMLR(e.r, e.r.Width/3)
	e.immPOISON = true
	enemyLIST = append(enemyLIST, e)

	e = ENEMY{} //1 Slime Too
	e.nm = "Slime Too"
	e.imFACES = "c"
	e.hp = 4
	e.hpmax = e.hp
	e.a = mAnimXY("im/enemies/003slimeIDLE.png", 0, 0, 64, 0, 6, 1, 7)
	e.a2 = mAnimXY("im/enemies/003slimeATK.png", 0, 0, 64, 0, 11, 1, 10)
	w = levbgREC.Width / 2
	e.r = z.NewRectangle(0, 0, w, w)
	e.rC = qResizeRecOffsetSMLR(e.r, e.r.Width/3)
	e.immPOISON = true
	enemyLIST = append(enemyLIST, e)

	e = ENEMY{} //2 Blob
	e.nm = "Blob"
	e.imFACES = "c"
	e.hp = 8
	e.hpmax = e.hp
	e.a = mAnimXY("im/enemies/024blobIDLE.png", 0, 0, 64, 0, 6, 1, 10)
	e.a2 = mAnimXY("im/enemies/024blobATK.png", 0, 0, 64, 0, 8, 1, 12)
	w = levbgREC.Width / 2
	e.r = z.NewRectangle(0, 0, w, w)
	e.rC = qResizeRecOffsetSMLR(e.r, (e.r.Width/40)*15)
	e.rC.Y += UNIT + UNIT/2
	e.immPOISON = true
	enemyLIST = append(enemyLIST, e)

	e = ENEMY{} //3 The Crow
	e.nm = "The Crow"
	e.fly = true
	e.imFACES = "r"
	e.hp = 8
	e.hpmax = e.hp
	e.a = mAnimXY("im/enemies/017crowIDLE.png", 0, 0, 128, 0, 4, 1, 7)
	e.a2 = mAnimXY("im/enemies/017crowATK.png", 0, 0, 128, 0, 6, 1, 10)
	w = (levbgREC.Width / 20) * 18
	e.r = z.NewRectangle(0, 0, w, w)
	e.rC = qResizeRecOffsetSMLR(e.r, (e.r.Width/40)*16)
	e.rC = qResizeRecSHORTERbottom(e.rC, UNIT)
	enemyLIST = append(enemyLIST, e)

	e = ENEMY{} //4 Bernie
	e.nm = "Bernie"
	e.imFACES = "c"
	e.hp = 10
	e.hpmax = e.hp
	e.a = mAnimXY("im/enemies/004slimeIDLE.png", 0, 0, 64, 0, 6, 1, 7)
	e.a2 = mAnimXY("im/enemies/004slimeATK.png", 0, 0, 64, 0, 9, 1, 10)
	w = levbgREC.Width / 2
	e.r = z.NewRectangle(0, 0, w, w)
	e.rC = qResizeRecOffsetSMLR(e.r, e.r.Width/3)
	e.immBURN = true
	e.immFREEZE = true
	e.immPOISON = true
	enemyLIST = append(enemyLIST, e)

	e = ENEMY{} //5 Mischeevis
	e.nm = "Mischeevis"
	e.fly = true
	e.imFACES = "l"
	e.hp = 10
	e.hpmax = e.hp
	e.a = mAnimXYWH("im/enemies/005flyingIDLE.png", 0, 0, 81, 71, 0, 4, 1, 7)
	e.a2 = mAnimXYWH("im/enemies/005flyingATK.png", 0, 0, 81, 71, 0, 8, 1, 10)
	w = levbgREC.Width / 3
	e.r = z.NewRectangle(0, 0, w, qHeightProportional(81, 71, w))
	e.rC = e.r
	e.immBURN = true
	enemyLIST = append(enemyLIST, e)

	e = ENEMY{} //6 Baldrick
	e.nm = "Baldrick"
	e.imFACES = "c"
	e.hp = 14
	e.hpmax = e.hp
	e.a = mAnimXY("im/enemies/012orcIDLE.png", 0, 0, 64, 0, 4, 1, 7)
	e.a2 = mAnimXY("im/enemies/012orcATK.png", 0, 0, 64, 0, 8, 1, 10)
	w = levbgREC.Width / 2.5
	e.r = z.NewRectangle(0, 0, w, w)
	e.rC = qResizeRecOffsetSMLR(e.r, e.r.Width/4)
	enemyLIST = append(enemyLIST, e)

	e = ENEMY{} //7 Hogan
	e.nm = "Hogan"
	e.imFACES = "c"
	e.hp = 18
	e.hpmax = e.hp
	e.a = mAnimXY("im/enemies/013orcIDLE.png", 0, 0, 64, 0, 4, 1, 7)
	e.a2 = mAnimXY("im/enemies/013orcATK.png", 0, 0, 64, 0, 8, 1, 10)
	w = levbgREC.Width / 2.5
	e.r = z.NewRectangle(0, 0, w, w)
	e.rC = qResizeRecOffsetSMLR(e.r, e.r.Width/4)
	enemyLIST = append(enemyLIST, e)

	e = ENEMY{} //8 Barbara
	e.nm = "Barbara"
	e.imFACES = "c"
	e.hp = 16
	e.hpmax = e.hp
	e.a = mAnimXY("im/enemies/014orcIDLE.png", 0, 0, 64, 0, 4, 1, 7)
	e.a2 = mAnimXY("im/enemies/014orcATK.png", 0, 0, 64, 0, 8, 1, 10)
	w = levbgREC.Width / 2.5
	e.r = z.NewRectangle(0, 0, w, w)
	e.rC = qResizeRecOffsetSMLR(e.r, e.r.Width/4)
	enemyLIST = append(enemyLIST, e)

	e = ENEMY{} //9 Pothead
	e.nm = "Pothead"
	e.imFACES = "r"
	e.hp = 12
	e.hpmax = e.hp
	e.a = mAnimXY("im/enemies/007plantIDLE.png", 0, 0, 64, 0, 4, 1, 7)
	e.a2 = mAnimXY("im/enemies/007plantATK.png", 0, 0, 64, 0, 7, 1, 10)
	w = levbgREC.Width / 2.2
	e.r = z.NewRectangle(0, 0, w, w)
	e.rC = qResizeRecOffsetSMLR(e.r, e.r.Width/4)
	e.immPOISON = true
	e.immSTUN = true
	enemyLIST = append(enemyLIST, e)

	e = ENEMY{} //10 Audrey
	e.nm = "Audrey"
	e.imFACES = "l"
	e.hp = 16
	e.hpmax = e.hp
	e.a = mAnimXY("im/enemies/001plantIDLE.png", 0, 0, 64, 0, 4, 1, 7)
	e.a2 = mAnimXY("im/enemies/001plantATK.png", 0, 0, 64, 0, 7, 1, 10)
	w = levbgREC.Width / 2.2
	e.r = z.NewRectangle(0, 0, w, w)
	e.rC = qResizeRecOffsetSMLR(e.r, e.r.Width/4)
	e.immPOISON = true
	e.immSTUN = true
	enemyLIST = append(enemyLIST, e)

	e = ENEMY{} //11 Spike
	e.nm = "Spike"
	e.imFACES = "r"
	e.hp = 16
	e.hpmax = e.hp
	e.a = mAnimXY("im/enemies/008plantIDLE.png", 0, 0, 64, 0, 4, 1, 7)
	e.a2 = mAnimXY("im/enemies/008plantATK.png", 0, 0, 64, 0, 7, 1, 10)
	w = levbgREC.Width / 2.2
	e.r = z.NewRectangle(0, 0, w, w)
	e.rC = qResizeRecOffsetSMLR(e.r, e.r.Width/4)
	e.immPOISON = true
	e.immSTUN = true
	enemyLIST = append(enemyLIST, e)

	e = ENEMY{} //12 Jim
	e.nm = "Jim"
	e.imFACES = "r"
	e.hp = 18
	e.hpmax = e.hp
	e.a = mAnimXY("im/enemies/006wormIDLE.png", 0, 0, 90, 0, 9, 1, 8)
	e.a2 = mAnimXY("im/enemies/006wormATK.png", 0, 0, 90, 0, 16, 1, 18)
	w = levbgREC.Width / 2
	e.r = z.NewRectangle(0, 0, w, w)
	e.rC = qResizeRecOffsetSMLR(e.r, e.r.Width/4)
	e.immBURN = true
	enemyLIST = append(enemyLIST, e)

	e = ENEMY{} //13 Octeye
	e.nm = "Octeye"
	e.imFACES = "c"
	e.hp = 18
	e.hpmax = e.hp
	e.a = mAnimXY("im/enemies/025octiIDLE.png", 0, 0, 64, 0, 8, 1, 5)
	e.a2 = mAnimXY("im/enemies/025octiATK.png", 0, 0, 64, 0, 8, 1, 15)
	w = levbgREC.Width / 4
	e.r = z.NewRectangle(0, 0, w, w)
	e.rC = qResizeRecOffsetSMLR(e.r, UNIT)
	e.rC = qResizeRecSHORTERtop(e.rC, UNIT/4)
	enemyLIST = append(enemyLIST, e)

	e = ENEMY{} //14 Igor
	e.nm = "Igor"
	e.fly = true
	e.imFACES = "r"
	e.hp = 18
	e.hpmax = e.hp
	e.a = mAnimXY("im/enemies/011flyingIDLE.png", 0, 0, 150, 0, 8, 1, 7)
	e.a2 = mAnimXY("im/enemies/011flyingATK.png", 0, 0, 150, 0, 8, 1, 10)
	w = levbgREC.Width
	e.r = z.NewRectangle(0, 0, w, w)
	e.rC = qResizeRecOffsetSMLR(e.r, (e.r.Width/40)*16)
	e.immSTUN = true
	enemyLIST = append(enemyLIST, e)

	e = ENEMY{} //15 Ogg
	e.nm = "Ogg"
	e.imFACES = "r"
	e.hp = 16
	e.hpmax = e.hp
	e.a = mAnimXYWH("im/enemies/015frogIDLE.png", 0, 0, 384, 128, 0, 5, 1, 7)
	e.a2 = mAnimXYWH("im/enemies/015frogATK.png", 0, 0, 384, 128, 0, 8, 1, 10)
	w = (levbgREC.Width / 20) * 17
	e.r = z.NewRectangle(0, 0, w, qHeightProportional(384, 128, w))
	e.rC = qResizeRecNARROWERcnt(e.r, UNIT+e.r.Width/3)
	e.rC = qResizeRecSHORTERtop(e.rC, UNIT+UNIT/2)
	e.immPOISON = true
	enemyLIST = append(enemyLIST, e)

	e = ENEMY{} //16 Pet Rock
	e.nm = "Pet Rock"
	e.imFACES = "r"
	e.hp = 20
	e.hpmax = e.hp
	e.a = mAnimXYWH("im/enemies/016rockIDLE.png", 0, 0, 384, 128, 0, 5, 1, 7)
	e.a2 = mAnimXYWH("im/enemies/016rockATK.png", 0, 0, 384, 128, 0, 19, 1, 10)
	w = levbgREC.Width + levbgREC.Width/4
	e.r = z.NewRectangle(0, 0, w, qHeightProportional(384, 128, w))
	e.rC = qResizeRecNARROWERcnt(e.r, (e.r.Width/40)*16)
	e.rC = qResizeRecSHORTERtop(e.rC, UNIT/2+UNIT*2)
	e.immELECTRIC = true
	e.immSTUN = true
	e.immFREEZE = true
	e.immPOISON = true
	e.immBURN = true
	enemyLIST = append(enemyLIST, e)

	e = ENEMY{} //17 Marlin
	e.nm = "Marlin"
	e.imFACES = "r"
	e.hp = 20
	e.hpmax = e.hp
	e.a = mAnimXY("im/enemies/009wizardIDLE.png", 0, 0, 150, 0, 8, 1, 7)
	e.a2 = mAnimXY("im/enemies/009wizardATK.png", 0, 0, 150, 0, 8, 1, 10)
	w = (levbgREC.Width / 20) * 14
	e.r = z.NewRectangle(0, 0, w, w)
	e.rC = qResizeRecOffsetSMLR(e.r, e.r.Width/3)
	e.rC = qResizeRecNARROWERcnt(e.rC, UNIT)
	e.immELECTRIC = true
	e.immSTUN = true
	e.immPOISON = true
	enemyLIST = append(enemyLIST, e)

	e = ENEMY{} //18 Martin
	e.nm = "Martin"
	e.imFACES = "r"
	e.hp = 20
	e.hpmax = e.hp
	e.a = mAnimXY("im/enemies/010wizardIDLE.png", 0, 0, 250, 0, 8, 1, 7)
	e.a2 = mAnimXY("im/enemies/010wizardATK.png", 0, 0, 250, 0, 8, 1, 10)
	w = (levbgREC.Width / 20) * 17
	e.r = z.NewRectangle(0, 0, w, w)
	e.rC = qResizeRecOffsetSMLR(e.r, (e.r.Width/40)*15)
	e.rC = qResizeRecNARROWERcnt(e.rC, UNIT/2)
	e.immELECTRIC = true
	e.immSTUN = true
	e.immFREEZE = true
	enemyLIST = append(enemyLIST, e)

	e = ENEMY{} //19 Claude
	e.nm = "Claude"
	e.imFACES = "r"
	e.hp = 20
	e.hpmax = e.hp
	e.a = mAnimXYWH("im/enemies/018hand.png", 0, 0, 62, 64, 0, 8, 1, 6)
	e.a2 = e.a
	w = levbgREC.Width / 5
	e.r = z.NewRectangle(0, 0, w, w)
	e.rC = qResizeRecNARROWERcnt(e.r, UNIT/2)
	e.immSTUN = true
	e.immPOISON = true
	enemyLIST = append(enemyLIST, e)

	e = ENEMY{} //20 Priscilla
	e.nm = "Priscilla"
	e.imFACES = "c"
	e.hp = 20
	e.hpmax = e.hp
	e.a = mAnimXY("im/enemies/019prism.png", 0, 0, 64, 0, 7, 1, 10)
	e.a2 = e.a
	w = levbgREC.Width / 4
	e.r = z.NewRectangle(0, 0, w, w)
	e.rC = qResizeRecOffsetSMLR(e.r, (e.r.Width/40)*5)
	e.rC = qResizeRecNARROWERright(e.rC, UNIT/2)
	e.immELECTRIC = true
	e.immSTUN = true
	e.immFREEZE = true
	e.immPOISON = true
	e.immBURN = true
	enemyLIST = append(enemyLIST, e)

	e = ENEMY{} //21 Litchi
	e.nm = "Litchi"
	e.fly = true
	e.imFACES = "l"
	e.hp = 20
	e.hpmax = e.hp
	e.a = mAnimXY("im/enemies/020mouthIDLE.png", 0, 0, 80, 0, 4, 1, 7)
	e.a2 = mAnimXY("im/enemies/020mouthATK.png", 0, 0, 80, 0, 4, 1, 10)
	w = levbgREC.Width / 2
	e.r = z.NewRectangle(0, 0, w, w)
	e.rC = qResizeRecOffsetSMLR(e.r, (e.r.Width/40)*14)
	enemyLIST = append(enemyLIST, e)

	e = ENEMY{} //22 Harkonen
	e.nm = "Harkonen"
	e.imFACES = "l"
	e.hp = 20
	e.hpmax = e.hp
	e.a = mAnimXY("im/enemies/021wormIDLE.png", 0, 0, 64, 0, 4, 1, 7)
	e.a2 = mAnimXY("im/enemies/021wormATK.png", 0, 0, 64, 0, 4, 1, 10)
	w = levbgREC.Width / 3
	e.r = z.NewRectangle(0, 0, w, w)
	e.rC = qResizeRecOffsetSMLR(e.r, (e.r.Width/40)*10)
	e.immPOISON = true
	enemyLIST = append(enemyLIST, e)

	e = ENEMY{} //23 Spaghetti
	e.nm = "Spaghetti"
	e.imFACES = "r"
	e.hp = 20
	e.hpmax = e.hp
	e.a = mAnimXY("im/enemies/022gorgonIDLE.png", 0, 0, 128, 0, 7, 1, 6)
	e.a2 = mAnimXY("im/enemies/022gorgonATK.png", 0, 0, 128, 0, 7, 1, 9)
	w = levbgREC.Width / 2.5
	e.r = z.NewRectangle(0, 0, w, w)
	e.rC = qResizeRecNARROWERcnt(e.r, e.r.Width/3)
	e.rC = qResizeRecSHORTERtop(e.rC, UNIT+UNIT/2)
	e.immSTUN = true
	e.immFREEZE = true
	e.immPOISON = true
	enemyLIST = append(enemyLIST, e)

	e = ENEMY{} //24 Mr Bones
	e.nm = "Mr Bones"
	e.imFACES = "l"
	e.hp = 20
	e.hpmax = e.hp
	e.a = mAnimXYWH("im/enemies/023skeletonIDLE.png", 0, 0, 146, 64, 0, 8, 1, 10)
	e.a2 = mAnimXYWHsubtract("im/enemies/023skeletonATK.png", 0, 0, 146, 64, 0, 5, 5, 2, 10)
	w = levbgREC.Width / 1.5
	e.r = z.NewRectangle(0, 0, w, qHeightProportional(146, 64, w))
	e.rC = qResizeRecNARROWERcnt(e.r, (e.r.Width/40)*15)
	e.rC = qResizeRecSHORTERtop(e.rC, UNIT)
	e.immELECTRIC = true
	e.immSTUN = true
	e.immFREEZE = true
	e.immPOISON = true
	e.immBURN = true
	enemyLIST = append(enemyLIST, e)

	for i := range enemyLIST {
		enemyLIST[i].hit = enmhitANMS[RINT(0, len(enmhitANMS))]
	}
}
func mENEMIES() { //MARK: MAKE ENEMIES

	var num, startnum, endnum int
	switch levNUM {
	case 1, 2:
		num = 2
		endnum = 3
	}
	countbreak := 0
	e := ENEMY{}

	endnum = len(enemyLIST) - 1
	num = RINT(3, 7)

	var enmnums []int

	var pos1, pos2, pos3, pos4, pos5, pos6 bool

	for num > 0 && countbreak < 100 {
		canadd := true
		choose := RINT(startnum, endnum)
		if len(enmnums) > 0 {
			if slices.Contains(enmnums, choose) {
				canadd = false
			} else {
				enmnums = append(enmnums, choose)
				e = enemyLIST[choose]
			}
		} else {
			enmnums = append(enmnums, choose)
			e = enemyLIST[choose]
		}

		if canadd {
			if e.fly {
				if pos4 && pos5 && pos6 {
					canadd = false
				} else {
					found := false
					for !found {
						switch RINT(4, 7) {
						case 4:
							if !pos4 {
								e.pos = 4
								pos4 = true
								found = true
							}
						case 5:
							if !pos5 {
								e.pos = 5
								pos5 = true
								found = true
							}
						case 6:
							if !pos6 {
								e.pos = 6
								pos6 = true
								found = true
							}
						}
					}
				}
			} else {
				if pos1 && pos2 && pos3 {
					canadd = false
				} else {
					found := false
					for !found {
						switch RINT(1, 4) {
						case 1:
							if !pos1 {
								e.pos = 1
								pos1 = true
								found = true
							}
						case 2:
							if !pos2 {
								e.pos = 2
								pos2 = true
								found = true
							}
						case 3:
							if !pos3 {
								e.pos = 3
								pos3 = true
								found = true
							}
						}
					}
				}
			}
		}

		if canadd {
			enemies = append(enemies, e)
			num--
		}

		countbreak++
	}

	mENEMIESPOSADJUST()
}
func mENEMIESPOSADJUST() { //MARK: ENEMY POSITION ADJUST
	/*
		DRAW POSITIONS

		v4	v5	v6	AIR

		v1	v2	v3	GROUND
	*/

	for i := range enemies {
		v1 := z.NewVector2(levbgREC.X+UNIT*2, levbgREC.Y+(levbgREC.Height/20)*12)
		v2 := z.NewVector2(levbgREC.X+levbgREC.Width/2-UNIT*2, levbgREC.Y+(levbgREC.Height/20)*12)
		v3 := z.NewVector2(levbgREC.X+levbgREC.Width-UNIT*5, levbgREC.Y+(levbgREC.Height/20)*12)

		v4 := z.NewVector2(levbgREC.X+UNIT*2, levbgREC.Y+UNIT*2)
		v5 := z.NewVector2(levbgREC.X+levbgREC.Width/2-UNIT*2, levbgREC.Y+UNIT*2)
		v6 := z.NewVector2(levbgREC.X+levbgREC.Width-UNIT*5, levbgREC.Y+UNIT*2)

		switch enemies[i].nm {
		case "Pet Rock":
			v3.X -= UNIT
		case "The Crow":
			v4.Y += UNIT / 2
			v5.Y += UNIT / 2
			v6.Y += UNIT / 2
			v5.X += UNIT / 2
		case "Priscilla":
			v2.X += UNIT / 2
		case "Spike", "Pothead", "Audrey":
			v1.Y -= UNIT / 2
			v2.Y -= UNIT / 2
			v3.Y -= UNIT / 2
		case "Marlin", "Martin":
			v1.Y -= UNIT
			v2.Y -= UNIT
			v3.Y -= UNIT
			v2.X += UNIT / 2
		case "Slime", "Ogg", "Slime Too", "Bernie":
			v1.Y += UNIT
			v2.Y += UNIT
			v3.Y += UNIT
		case "Spaghetti":
			v1.Y -= UNIT + UNIT/2
			v2.Y -= UNIT + UNIT/2
			v3.Y -= UNIT + UNIT/2
			v2.X += UNIT
			v1.X += UNIT / 2
		case "Octeye":
			v1.Y += UNIT + UNIT/2
			v2.Y += UNIT + UNIT/2
			v3.Y += UNIT + UNIT/2
		case "Blob":
			v1.Y += UNIT + UNIT/2
			v2.Y += UNIT + UNIT/2
			v3.Y += UNIT + UNIT/2
			v2.X += UNIT / 2
		case "Mischeevis":
			v4.X -= UNIT + UNIT/2
			v6.X -= UNIT
			v4.Y -= UNIT / 2
			v5.Y -= UNIT / 2
			v6.Y -= UNIT / 2
		case "Claude":
			v2.X += UNIT / 2
		case "Jim":
			v3.X -= UNIT
			v1.X -= UNIT / 2
		case "Hogan", "Barbara", "Baldrick":
			v1.Y += UNIT / 2
			v2.Y += UNIT / 2
			v3.Y += UNIT / 2
		case "Litchi":
			v5.X += UNIT
			v6.X += UNIT / 2
			v4.Y += UNIT / 2
			v5.Y += UNIT / 2
			v6.Y += UNIT / 2
		case "Igor":
			v4.Y -= UNIT / 2
			v5.Y -= UNIT / 2
			v6.Y -= UNIT / 2
		}

		switch enemies[i].pos {
		case 1:
			enemies[i].r, enemies[i].rC = moveRECCOLLIStoV2topleftandREC(enemies[i].r, enemies[i].rC, v1)
		case 2:
			enemies[i].r, enemies[i].rC = moveRECCOLLIStoV2topleftandREC(enemies[i].r, enemies[i].rC, v2)
		case 3:
			enemies[i].r, enemies[i].rC = moveRECCOLLIStoV2topleftandREC(enemies[i].r, enemies[i].rC, v3)
		case 4:
			enemies[i].r, enemies[i].rC = moveRECCOLLIStoV2topleftandREC(enemies[i].r, enemies[i].rC, v4)
		case 5:
			enemies[i].r, enemies[i].rC = moveRECCOLLIStoV2topleftandREC(enemies[i].r, enemies[i].rC, v5)
		case 6:
			enemies[i].r, enemies[i].rC = moveRECCOLLIStoV2topleftandREC(enemies[i].r, enemies[i].rC, v6)
		}
	}

}

func mANIM() {
	drawingANIM = mAnimXY("im/fx/001drawing.png", 0, 0, 64, 0, 14, 1, 30)
}
func mLEV() { //MARK: MAKE LEV
	//LEV BG
	levbgIM = levbg[RINT(0, len(levbg))]
	//BG
	im := PATTERNS[RINT(0, len(PATTERNS))]
	var x, y, zoom float32
	if im.r.Width < UNIT {
		zoom = float32(RINT(3, 7))
	} else if im.r.Width > UNIT*4 {
		zoom = float32(RINT(1, 4))
	} else {
		zoom = float32(RINT(2, 5))
	}
	im.rD.Width = im.r.Width * zoom
	im.rD.Height = im.r.Height * zoom
	for y < float32(SCRH) {
		im.rD.X = x
		im.rD.Y = y
		im.c = CRGB(255, 238, 185)
		im.a = RUINT8(30, 40)
		bg = append(bg, im)
		x += im.rD.Width
		if x > float32(SCRW) {
			x = 0
			y += im.rD.Height
		}
	}

}
func mCARDrecCNT() z.Vector2 {
	cnt := z.NewVector2(cardREC.X+cardREC.Width/2, cardREC.Y+cardREC.Height/2)
	return cnt
}

// UTILS
func REMENMDEAD(e []ENEMY) []ENEMY {
	var e2 []ENEMY
	for i := range e {
		if !e[i].dead {
			e2 = append(e2, e[i])
		}
	}
	return e2
}
func REMCARDOFF(c []CARD) []CARD {
	var c2 []CARD
	for i := range c {
		if !c[i].off {
			c2 = append(c2, c[i])
		}
	}
	return c2
}
