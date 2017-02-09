package main

import (
	Kiai "gitlab.com/tapir/kiai/client"
	"math/rand"
	"time"
)

var start bool = false //Variable to store the info if the game just started
var path int //Variable to show next character movement direction after reaching each corner. This variable is global because it will be set and used by different functions at different times.
var iterate int = 0 //Variable to set how many rounds bot spends at a corner
var turn *Kiai.TurnStatus //Variable to use turnstatus everywhere. This creates a variable with a TurnStatus type. The outcome unit of WaitForTurn() function is Turnstatus which consists of multiple int and []int data. X,Y,Health,MovesLeft,FiresLeft,PlayersLeft,Radar[]
var maxiteration int = 5

//Below here lies the function to understand where the bot should go at the start of the game

func firstdecision(gconfig Kiai.GameConfig)  {
	
	if turn.X <= gconfig.Width / 2 && turn.Y < gconfig.Height / 2 { //Be careful, this method favours the top left corner due to int variables rounding the results of the division to lower limit.
		path = 1   
	}
	if turn.X >= gconfig.Width / 2 && turn.Y < gconfig.Height / 2 {
		path = 2   
	}
	if turn.X >= gconfig.Width / 2 && turn.Y > gconfig.Height / 2 {
		path = 3   
	}
	if turn.X <= gconfig.Width / 2 && turn.Y > gconfig.Height / 2 {
		path = 4   
	}
	
	start = true
}

//Below here lies the function that will make the Easyfucka Bot fire towards his doomed opponents if a bot is seen on the radar

func firing() { 
	for i := 0; i < 8; i++ {
		if turn.Radar[i] > 1 {
			turn.Fire(i)
			iterate = 0
			if path < 4 {
				path += 1 
			} else {
				path = 1	
			}
		}
	}	
}

//Below lies the tactic after goign towards a corner. It is incomplete.

func maintactic () {
	
	firing()	
	var oddeven int
	
	if turn.MovesLeft % 2 != 0 {
		oddeven = turn.MovesLeft - 1
	} else {
		oddeven = turn.MovesLeft
	}
		if path == 1 {
			for i := 0; i < oddeven; i++ {
				if i <= oddeven/2 - 1 {
					turn.Move(Kiai.SouthEast)
				} else {
					turn.Move(Kiai.NorthWest)
				}
				firing() //If a condition is in both if and else statement, take it out of it and put it as common
			}
		} else if path == 2 {
			for i := 0; i < oddeven; i++ {
				if i <= oddeven/2 - 1  {
					turn.Move(Kiai.SouthWest)
				} else {
					turn.Move(Kiai.NorthEast)
				}
				firing()
			}
		} else if path == 3 {
			for i := 0; i < oddeven; i++ {
				if i <= oddeven/2 - 1 {
					turn.Move(Kiai.NorthWest)
				} else {
					turn.Move(Kiai.SouthEast)
				}
				firing()
			}
		} else if path == 4 {
			for i := 0; i < oddeven; i++ {
				if i <= oddeven/2 - 1 {
					turn.Move(Kiai.NorthEast)
				} else {
					turn.Move(Kiai.SouthWest)
				}
				firing()
			}
		} 	
}

//Below here lies the sacred main code of the Easyfucka Bot who will rape any other bot that he will cum across.

func main() {
	rand.Seed(time.Now().UnixNano())
	gconfig := Kiai.Connect("localhost:4000", "Easyfucka", 1, 0) //Creates a variable with a GameConfig type. The output unit of connect function is Gameconfig which consists of multiple int data. Width,Height,MaxMoves,MaxFires,MaxHealth,PlayerCount,Timeout
	defer Kiai.Disconnect() //Makes sure that at the end of the main function bot disconnects from server.
	
	
	for {
		turn = Kiai.WaitForTurn() //Assigns the starting values of the wait for turn to turn variable.
		
		if start == false {
			firstdecision(gconfig)
		}
		
		if iterate == maxiteration {
			iterate = 0
			if path < 4 {
				path += 1 
			} else {
				path = 1	
			}		
		}
		
		if path == 1 { //This part is one of the four movement tactics that belong to a path. For all the moves left it directs the bot towards the corner and if corner is reached, it starts main offense.
			
			for i := 0; i < turn.MovesLeft; i++ {
				firing()
				
				if turn.X == 0 && turn.Y == 0 {
					maintactic()
					break
				} else if turn.X != 0 && turn.Y != 0 {
					turn.Move(Kiai.NorthWest)
				} else if turn.X == 0 && turn.Y != 0 {
					turn.Move(Kiai.North)
				} else if turn.X != 0 && turn.Y == 0 {
					turn.Move(Kiai.West)
				}
			}
		}
		
		if path == 2 {
			
			for i := 0; i < turn.MovesLeft; i++ {
				firing()
				
				if turn.X == gconfig.Width-1 && turn.Y == 0 {
					maintactic()
					break
				} else if turn.X != gconfig.Width-1 && turn.Y != 0 {
					turn.Move(Kiai.NorthEast)
				} else if turn.X == gconfig.Width-1 && turn.Y != 0 {
					turn.Move(Kiai.North)
				} else if turn.X != gconfig.Width-1 && turn.Y == 0 {
					turn.Move(Kiai.East)
				}
			}
		}
		
		if path == 3 {
			
			for i := 0; i < turn.MovesLeft; i++ {
				firing()
				
				if turn.X == gconfig.Width-1 && turn.Y == gconfig.Height-1 {
					maintactic()
					break
				} else if turn.X != gconfig.Width-1 && turn.Y != gconfig.Height-1 {
					turn.Move(Kiai.SouthEast)
				} else if turn.X == gconfig.Width-1 && turn.Y != gconfig.Height-1 {
					turn.Move(Kiai.South)
				} else if turn.X != gconfig.Width-1 && turn.Y == gconfig.Height-1 {
					turn.Move(Kiai.East)
				}
			}
		}
		
		if path == 4 {
			
			for i := 0; i < turn.MovesLeft; i++ {
				firing()
				
				if turn.X == 0 && turn.Y == gconfig.Height-1 {
					maintactic()
					break
				} else if turn.X != 0 && turn.Y != gconfig.Height-1 {
					turn.Move(Kiai.SouthWest)
				} else if turn.X == 0 && turn.Y != gconfig.Height-1 {
					turn.Move(Kiai.South)
				} else if turn.X != 0 && turn.Y == gconfig.Height-1 {
					turn.Move(Kiai.West)
				}
			}
		}
	
//This section just shoots randomly towards the middle of the board at the last move if someone is not seen around.
	
		if path == 1 {
			if turn.Y == 0 {
				turn.Fire(Kiai.SouthEast)
			} else {
				turn.Fire(Kiai.East)
			}	 
		} 
		
		if path == 2 {
			if turn.X == gconfig.Width - 1 {
				turn.Fire(Kiai.SouthWest)
			} else {
				turn.Fire(Kiai.South)
			}		
		}
		
		if path == 3 {
			if turn.Y == gconfig.Height - 1 {
				turn.Fire(Kiai.NorthWest)
			} else {
				turn.Fire(Kiai.West)
			}
		}
		
		if path == 4 {
			if turn.X == 0 {
				turn.Fire(Kiai.NorthEast)
			} else {
				turn.Fire(Kiai.North)
			}
		}
		
		iterate += 1
		Kiai.EndTurn()
	}
}