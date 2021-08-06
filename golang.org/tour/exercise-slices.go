package main

import (
    "golang.org/x/tour/pic"
)

// make a interesting image!
func Pic(dx, dy int) [][]uint8 {
    // creating 2-d slices
    img := make([][]uint8, dx)
    for i := range img {
        img[i] = make([]uint8, dy)
    }

    for i := 0; i < dx; i++ {
        for j := 0; j < dy; j++ {
            img[i][j] = uint8(i*i + j*j)
        }
    }
    
    return img
}

func main() {
	pic.Show(Pic)
}

