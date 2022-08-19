package main

import (
	"fmt"

	"github.com/karim-w/kddd/app/api"
	"github.com/karim-w/kddd/domain/health"
	"github.com/karim-w/kddd/infra/cache"
	"github.com/karim-w/kddd/infra/repositories"
)

func main() {
	fmt.Println(`
	KKKKKKKKK    KKKKKKKDDDDDDDDDDDDD      DDDDDDDDDDDDD      DDDDDDDDDDDDD        
K:::::::K    K:::::KD::::::::::::DDD   D::::::::::::DDD   D::::::::::::DDD     
K:::::::K    K:::::KD:::::::::::::::DD D:::::::::::::::DD D:::::::::::::::DD   
K:::::::K   K::::::KDDD:::::DDDDD:::::DDDD:::::DDDDD:::::DDDD:::::DDDDD:::::D  
KK::::::K  K:::::KKK  D:::::D    D:::::D D:::::D    D:::::D D:::::D    D:::::D 
  K:::::K K:::::K     D:::::D     D:::::DD:::::D     D:::::DD:::::D     D:::::D
  K::::::K:::::K      D:::::D     D:::::DD:::::D     D:::::DD:::::D     D:::::D
  K:::::::::::K       D:::::D     D:::::DD:::::D     D:::::DD:::::D     D:::::D
  K:::::::::::K       D:::::D     D:::::DD:::::D     D:::::DD:::::D     D:::::D
  K::::::K:::::K      D:::::D     D:::::DD:::::D     D:::::DD:::::D     D:::::D
  K:::::K K:::::K     D:::::D     D:::::DD:::::D     D:::::DD:::::D     D:::::D
KK::::::K  K:::::KKK  D:::::D    D:::::D D:::::D    D:::::D D:::::D    D:::::D 
K:::::::K   K::::::KDDD:::::DDDDD:::::DDDD:::::DDDDD:::::DDDD:::::DDDDD:::::D  
K:::::::K    K:::::KD:::::::::::::::DD D:::::::::::::::DD D:::::::::::::::DD   
K:::::::K    K:::::KD::::::::::::DDD   D::::::::::::DDD   D::::::::::::DDD     
KKKKKKKKK    KKKKKKKDDDDDDDDDDDDD      DDDDDDDDDDDDD      DDDDDDDDDDDDD        
                                                                               
	`)
	fmt.Println("Starting Server...")
	cache := cache.NewMemcache()
	hr := repositories.NewHealthRepository(cache)
	hs := health.NewHealthService(hr)
	hh := api.NewHealthCheckHandler(hs)
	api.SetupRoutes(hh)

}
