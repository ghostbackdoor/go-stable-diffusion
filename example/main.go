package main

import (
	stableDiffusion "github.com/ghostbackdoor/go-stable-diffusion"
)

func main() {
	stableDiffusion.GenerateImage(
		256,
		256,
		0,
		15,
		42,
		"floating hair, portrait, ((loli)), ((one girl)), cute face, hidden hands, asymmetrical bangs, beautiful detailed eyes, eye shadow, hair ornament, ribbons, bowties, buttons, pleated skirt, (((masterpiece))), ((best quality)), colorful",
		"((part of the head)), ((((mutated hands and fingers)))), deformed, blurry, bad anatomy, disfigured, poorly drawn face, mutation, mutated, extra limb, ugly, poorly drawn hands, missing limb, blurry, floating limbs, disconnected limbs, malformed hands, blur, out of focus, long neck, long body, Octane renderer, lowres, bad anatomy, bad hands, text", "./test.png",
		"",
		"assets",
	)
}
