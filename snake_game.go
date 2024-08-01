package main

import (
	"container/list"
	"fmt"
)

type cordinates struct {
	x int
	y int
}

type SnakeGame struct {
	food   []cordinates
	snake  *list.List
	height int
	width  int
	score  int
}

func Constructor(width int, height int, food [][]int) SnakeGame {
	snakeGame := SnakeGame{
		score: 0,
	}
	for _, f := range food {
		snakeGame.food = append(snakeGame.food, cordinates{f[0], f[1]})
	}
	snakeGame.snake = list.New()
	snakeGame.snake.PushBack(cordinates{0, 0})
	snakeGame.height = height
	snakeGame.width = width

	return snakeGame
}

func (this *SnakeGame) Move(direction string) int {
	currentPosition := this.snake.Back().Value.(cordinates)
	switch direction {
	case "U":
		x := currentPosition.x - 1
		y := currentPosition.y
		if this.validate(x, y) == -1 {
			return -1
		}
		currentCord := cordinates{x, y}
		if len(this.food) > 0 && currentCord == this.food[0] {
			this.food = this.food[1:]
		} else {
			this.snake.Remove(this.snake.Front())
		}
		this.snake.PushBack(cordinates{x, y})

	case "D":
		x := currentPosition.x + 1
		y := currentPosition.y
		if this.validate(x, y) == -1 {
			return -1
		}
		currentCord := cordinates{x, y}
		if len(this.food) > 0 && currentCord == this.food[0] {
			this.food = this.food[1:]
		} else {
			this.snake.Remove(this.snake.Front())
		}
		this.snake.PushBack(cordinates{x, y})

	case "R":
		x := currentPosition.x
		y := currentPosition.y + 1
		if this.validate(x, y) == -1 {
			return -1
		}
		currentCord := cordinates{x, y}
		if len(this.food) > 0 && currentCord == this.food[0] {
			this.food = this.food[1:]
		} else {
			this.snake.Remove(this.snake.Front())
		}
		this.snake.PushBack(cordinates{x, y})
		return this.snake.Len() - 1

	case "L":
		x := currentPosition.x
		y := currentPosition.y - 1
		if this.validate(x, y) == -1 {
			return -1
		}
		currentCord := cordinates{x, y}
		if len(this.food) > 0 && currentCord == this.food[0] {
			this.food = this.food[1:]
		} else {
			this.snake.Remove(this.snake.Front())
		}
		this.snake.PushBack(cordinates{x, y})
	}
	return this.snake.Len() - 1
}

func (this *SnakeGame) validate(x, y int) int {
	if x < 0 || x == this.height {
		return -1
	}
	if y < 0 || y == this.width {
		return -1
	}
	return 0
}

func main() {
	snakeGame := Constructor(2, 2, [][]int{{0, 1}})
	fmt.Println(snakeGame.Move("R")) // return 0
	fmt.Println(snakeGame.Move("D")) // return 0
	//fmt.Println(snakeGame.Move("R")) // return 1, snake eats the first piece of food. The second piece of food appears at (0, 1).
	//fmt.Println(snakeGame.Move("U")) // return 1
	//fmt.Println(snakeGame.Move("L")) // return 2, snake eats the second food. No more food appears.
	//fmt.Println(snakeGame.Move("U")) // return -1, game over because snake collides
}
