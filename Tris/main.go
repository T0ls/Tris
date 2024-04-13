package main

import (
	. "fmt"
	"math/rand"
	"os"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
)

func Menù(options []string) int {
	c1 := color.New(color.FgRed)
	selectedOption := 0
	Println(" Choose game modality: ")
	Println()
	// Loop per leggere l'input dell'utente
	for {
		for i, option := range options {
			if i == selectedOption {
				c1.Printf(" >> - %s << \n", option)
			} else {
				Printf("    - %s\n", option)
			}
		}
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		// Esegui il codice in base all'input dell'utente
		if key == keyboard.KeyArrowUp || char == 'w' { // Up || w
			selectedOption--
			if selectedOption < 0 {
				selectedOption = len(options) - 1
			}
		} else if key == keyboard.KeyArrowDown || char == 's' { // Down || s
			selectedOption++
			if selectedOption >= len(options) {
				selectedOption = 0
			}
		} else if key == keyboard.KeyEnter || key == keyboard.KeySpace { // Enter || SpaceBar
			Println()
			Println("You chose:", options[selectedOption])
			return selectedOption
		}
		println()
		os.Stdout.WriteString("\033[2J\033[1;1H") // pulisco la board
		Println(" Choose game modality: ")
		Println()
	}
}

// Minimax Algoritmo

func switchPlayer(player string) string {
	if player == "X" {
		return "O"
	} else {
		return "X"
	}
}

// Definizione della funzione di valutazione
func evaluate(board [][]string) int {
	// Controllo delle righe
	for _, row := range board {
		if row[0] == row[1] && row[1] == row[2] {
			if row[0] == "X" {
				return 100
			} else if row[0] == "O" {
				return -100
			}
		}
	}

	// Controllo delle colonne
	for col := 0; col < 3; col++ {
		if board[0][col] == board[1][col] && board[1][col] == board[2][col] {
			if board[0][col] == "X" {
				return 100
			} else if board[0][col] == "O" {
				return -100
			}
		}
	}

	// Controllo delle diagonali
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		if board[0][0] == "X" {
			return 100
		} else if board[0][0] == "O" {
			return -100
		}
	} else if board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		if board[0][2] == "X" {
			return 100
		} else if board[0][2] == "O" {
			return -100
		}
	}

	// Se non ci sono vittorie, restituisce 0
	return 0
}

// Definizione della funzione Minimax
func minimax(board [][]string, player string, depth int) (int, []int) {
	// Controllo dei nodi terminali
	score := evaluate(board)
	if score == 100 || score == -100 || depth == 0 {
		return score, nil
	}

	// Scelta del giocatore corrente
	var bestScore int
	var bestMove []int
	if player == "X" {
		bestScore = -1000
	} else {
		bestScore = 1000
	}

	// Iterazione sulle mosse possibili
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if board[row][col] == "" {
				// Applica la mossa e valuta il risultato
				board[row][col] = player
				var score int
				if player == "X" {
					score, _ = minimax(board, "O", depth-1)
				} else {
					score, _ = minimax(board, "X", depth-1)
				}
				board[row][col] = ""

				// Aggiorna il miglior risultato e la miglior mossa
				if player == "X" && score > bestScore {
					bestScore = score
					bestMove = []int{row, col}
				} else if player == "O" && score < bestScore {
					bestScore = score
					bestMove = []int{row, col}
				}
			}
		}
	}

	// Restituisce il miglior risultato e la miglior mossa
	return bestScore, bestMove
}

/*func minimax(board [][]string, player string, depth int) (int, []int) {
    // Controllo dei nodi terminali
    score := evaluate(board)
    if score == 100 || score == -100 || depth == 0 {
        return score, nil
    }

    // Scelta del giocatore corrente
    var bestScore int
    var bestMove []int
    if player == "X" {
        bestScore = -1000
    } else {
        bestScore = 1000
    }

    // Iterazione sulle mosse possibili
    for row := 0; row < 3; row++ {
        for col := 0; col < 3; col++ {
            if board[row][col] == "" {
                // Applica la mossa e valuta il risultato
                board[row][col] = player
                var score int
                if player == "X" {
                    score, _ = minimax(board, "O", depth-1)
                } else {
                    score, _ = minimax(board, "X", depth-1)
                }
                board[row][col] = ""

                // Aggiorna il miglior risultato e la miglior mossa
                if player == "X" && score > bestScore {
                    bestScore = score
                    bestMove = []int{row, col}
                } else if player == "O" && score < bestScore {
                    bestScore = score
                    bestMove = []int{row, col}
                }
            }
        }
    }

    // Restituisce il miglior risultato e la miglior mossa
    return bestScore, bestMove
}*/

// Mossa Pc
func MossaPc(Campo [3][3]string, player string, x int) [3][3]string {
	/*var m = [3][3]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	Println(" Round of:", player)
	Println()
	bMove := Minimax(Campo, player)
	Println()
	Println("Mossa Pc:", bMove)
	Println()
	Campo[bMove[0]][bMove[1]] = player
	printCampo(Campo, m)
	Println()*/
	return Campo
}

// Mossa
func Mossa(Campo [3][3]string, player string) [3][3]string {
	var CampoPrev = [3][3]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}}
	occ := false
	skip := false
	var cRT string
	CampoPrev = Campo
	countV := 1
	countH := 1
	cRT = Campo[countV][countH]
	CampoPrev[countV][countH] = player
	Println(" Round of:", player)
	Println()
	printCampoPrev(CampoPrev, countV, countH, player)
	Println()
	// Loop per leggere l'input dell'utente
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		// Eseguo il codice in base all'input dell'utente
		if key == keyboard.KeyArrowUp || char == 'w' { // Up || w
			CampoPrev[countV][countH] = cRT
			countV--
			if countV < 0 {
				countV = 2
			}
			cRT = Campo[countV][countH]
			CampoPrev[countV][countH] = player
			os.Stdout.WriteString("\033[2J\033[1;1H") // pulisco la board
		} else if key == keyboard.KeyArrowDown || char == 's' { // Dow || s
			CampoPrev[countV][countH] = cRT
			countV++
			if countV >= 3 {
				countV = 0
			}
			cRT = Campo[countV][countH]
			CampoPrev[countV][countH] = player
			os.Stdout.WriteString("\033[2J\033[1;1H") // pulisco la board
		} else if key == keyboard.KeyArrowRight || char == 'd' { // Right || d
			CampoPrev[countV][countH] = cRT
			countH++
			if countH >= 3 {
				countH = 0
			}
			cRT = Campo[countV][countH]
			CampoPrev[countV][countH] = player
			os.Stdout.WriteString("\033[2J\033[1;1H") // pulisco la board
		} else if key == keyboard.KeyArrowLeft || char == 'a' { // Left || a
			CampoPrev[countV][countH] = cRT
			countH--
			if countH < 0 {
				countH = 2
			}
			cRT = Campo[countV][countH]
			CampoPrev[countV][countH] = player
			os.Stdout.WriteString("\033[2J\033[1;1H") // pulisco la board
		} else if key == keyboard.KeyEnter || key == keyboard.KeySpace { // Enter || SpaceBar
			if Campo[countV][countH] == " " {
				Campo[countV][countH] = player
				return Campo
			} else {
				occ = true
			}
			os.Stdout.WriteString("\033[2J\033[1;1H") // pulisco la board
		} else {
			skip = true
		}
		// Prompt inserito non valido
		if !skip {
			// Vera stamapa, modificare qui
			Println(" Round of:", player)
			Println()
			printCampoPrev(CampoPrev, countV, countH, player)
			println()
			if occ { // controllo x casella già occupata
				Println("Box already occupied!")
				Println()
				occ = false
			}
		} else {
			skip = false
		}
	}
}

// Stampa campo prev con colori
func printCampoPrev(Campo [3][3]string, coutnV int, countH int, player string) {
	c1 := color.New(color.FgRed)
	c2 := color.New(color.FgBlue)
	c3 := color.New(color.FgYellow)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j < 2 {
				Print(" ")
				if i == coutnV && j == countH {
					c3.Print(Campo[i][j])
				} else {
					if Campo[i][j] == "X" {
						c1.Print(Campo[i][j])
					} else {
						c2.Print(Campo[i][j])
					}
				}
				Print(" ¦")
			} else {
				Print(" ")
				if i == coutnV && j == countH {
					c3.Print(Campo[i][j])
				} else {
					if Campo[i][j] == "X" {
						c1.Print(Campo[i][j])
					} else {
						c2.Print(Campo[i][j])
					}
				}
				Print(" ")
			}
		}
		if i < 2 {
			Println("\n---+---+---")
		} else {
			Println()
		}
	}
}

// Print risultati  highlighted
func printCampo(Campo [3][3]string, m [3][3]int) {
	c1 := color.New(color.FgRed)
	c2 := color.New(color.FgBlue)
	c3 := color.New(color.FgGreen)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j < 2 {
				Print(" ")
				if m[i][j] == 1 {
					c3.Print(Campo[i][j])
				} else {
					if Campo[i][j] == "X" {
						c1.Print(Campo[i][j])
					} else {
						c2.Print(Campo[i][j])
					}
				}
				Print(" ¦")
			} else {
				Print(" ")
				if m[i][j] == 1 {
					c3.Print(Campo[i][j])
				} else {
					if Campo[i][j] == "X" {
						c1.Print(Campo[i][j])
					} else {
						c2.Print(Campo[i][j])
					}
				}
				Print(" ")
			}
		}
		if i < 2 {
			Println("\n---+---+---")
		} else {
			Println()
		}
	}
}

// Verifica il vincitore
func checkWin(Campo [3][3]string, player string) (bool, [3][3]int) {
	var m = [3][3]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	// Verifica righe
	for i := 0; i < 3; i++ {
		if Campo[i][0] == player && Campo[i][1] == player && Campo[i][2] == player {
			m[i][0] = 1
			m[i][1] = 1
			m[i][2] = 1
			return true, m
		}
	}

	// Verifica colonne
	for j := 0; j < 3; j++ {
		if Campo[0][j] == player && Campo[1][j] == player && Campo[2][j] == player {
			m[0][j] = 1
			m[1][j] = 1
			m[2][j] = 1
			return true, m
		}
	}

	// Verifica diagonali
	if Campo[0][0] == player && Campo[1][1] == player && Campo[2][2] == player {
		m[0][0] = 1
		m[1][1] = 1
		m[2][2] = 1
		return true, m
	}

	if Campo[0][2] == player && Campo[1][1] == player && Campo[2][0] == player {
		m[0][2] = 1
		m[1][1] = 1
		m[2][0] = 1
		return true, m
	}

	return false, m
}

// Partita 1 vs Pc
func Partita1vsPc(x int, Campo [3][3]string) {
	os.Stdout.WriteString("\033[2J\033[1;1H") // pulisco la board
	var player string
	rand.Seed(time.Now().UnixNano()) // inizializza il generatore di numeri casuali con un seme univoco
	num := rand.Intn(2)
	for {
		if x%2 == 0 {
			if num == 0 {
				player = "0"
				Campo = Mossa(Campo, player)
			} else {
				player = "0"
				Campo = MossaPc(Campo, player, x)
			}
			os.Stdout.WriteString("\033[2J\033[1;1H") // pulisco la board
		} else {
			if num == 1 {
				player = "X"
				Campo = Mossa(Campo, player)
			} else {
				player = "X"
				Campo = MossaPc(Campo, player, x)
			}
			os.Stdout.WriteString("\033[2J\033[1;1H") // pulisco la board
		}
		// Controllo vincita
		bol, m := checkWin(Campo, player)
		if bol {
			os.Stdout.WriteString("\033[2J\033[1;1H") // pulisco la board
			Println("The Winner is:", player)
			Println()
			printCampo(Campo, m)
			Println()
			Println("Game Ended!")
			Println()
			break
		}
		// Controllo pareggio
		if x == 9 {
			os.Stdout.WriteString("\033[2J\033[1;1H") // pulisco la board
			Println("   Draw!")
			Println()
			printCampo(Campo, m)
			Println()
			Println("Game Ended!")
			Println()
			break
		}
		x++
	}
}

// Partita 1 vs 1
func Partita1vs1(x int, Campo [3][3]string) {
	os.Stdout.WriteString("\033[2J\033[1;1H") // pulisco la board
	var player string
	for {
		if x%2 == 0 {
			player = "0"
			Campo = Mossa(Campo, player)
			os.Stdout.WriteString("\033[2J\033[1;1H") // pulisco la board
		} else {
			player = "X"
			Campo = Mossa(Campo, player)
			os.Stdout.WriteString("\033[2J\033[1;1H") // pulisco la board
		}
		// Controllo vincita
		bol, m := checkWin(Campo, player)
		if bol {
			os.Stdout.WriteString("\033[2J\033[1;1H") // pulisco la board
			Println("The Winner is:", player)
			Println()
			printCampo(Campo, m)
			Println()
			Println("Game Ended!")
			Println()
			break
		}
		// Controllo pareggio
		if x == 9 {
			os.Stdout.WriteString("\033[2J\033[1;1H") // pulisco la board
			Println("   Draw!")
			Println()
			printCampo(Campo, m)
			Println()
			Println("Game Ended!")
			Println()
			break
		}
		x++
	}
}

func main() {
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}

	defer keyboard.Close()
	os.Stdout.WriteString("\033[2J\033[1;1H") // pulisco la board
	// Inizializzo le variabili
	x := 1
	var Campo = [3][3]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}}
	// Menù
	options := []string{"1 vs Pc", "1 vs 1", "Quit"}
	mChoice := Menù(options)
	switch mChoice {
	case 0:
		// Partita 1 vs Pc
		Partita1vsPc(x, Campo)
	case 1:
		// Partita 1 vs 1
		Partita1vs1(x, Campo)
	case 2:
		return
	}
}

//TO DO:
/*
- sviluppo pc contro cui giocare
	- implementazione algoritmo minimax
*/
