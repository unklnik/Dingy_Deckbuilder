package main

import (
	"time"
)

func UP() {
	INP()
	UPGAME()
	TIMERS()
}

var (
	FRAMES      int32
	FRAMESTIMER time.Time
	SECONDS     int
)

func TIMERS() {

	//FRAMES
	if FRAMESTIMER.IsZero() {
		FRAMESTIMER = time.Now()
	}
	if time.Since(FRAMESTIMER) >= time.Second/time.Duration(FPS) {
		FRAMES++
		if FRAMES == FPS {
			SECONDS++
			FRAMES = 0
			FRAMESTIMER = time.Time{}
		}
	}
}

func UPGAME() {
	UPENEMIES()
	//MOUSE
	if MSL && clickT == 0 {
		clickT = clickP
	}

	//TIMERS
	if nomanaT > 0 {
		nomanaT--
	}
	if notselectedT > 0 {
		notselectedT--
	}
	if noTurnsT > 0 {
		noTurnsT--
	}
	if batORfairy {
		if batT > 0 {
			batT--
		}
	} else {
		if fairyT > 0 {
			fairyT--
		}
	}

	if clickT > 0 {
		clickT--
	}
	if drawingT > 0 {
		drawingT--
	}
}

func UPENEMIES() {

	if enmTURN {
		if enmTurnCount < len(enemies)-1 {
			if !enemies[enmTurnCount].atk && enemies[enmTurnCount].played {
				enmTurnCount++
				enemies[enmTurnCount].atk = true
			}
		} else {
			for i := range enemies {
				enemies[i].played = false
			}
			enmTurnCount = 0
			enmTURN = false
		}
	}

	for i := range enemies {
		//BURN
		if len(enemies[i].burn) != 0 {
			for j := range enemies[i].burn {
				if enemies[i].burn[j] > 0 {
					enemies[i].hp -= enemies[i].burnAMOUNT
					enemies[i].burn[j]--
				}
			}
		}
		//POISON

		if len(enemies[i].poison) != 0 {
			for j := range enemies[i].poison {
				if enemies[i].poison[j] > 0 {
					enemies[i].hp -= enemies[i].poisonAMOUNT
					enemies[i].poison[j]--
				}
			}
		}
	}

}
