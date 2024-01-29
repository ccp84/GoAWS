         coat := "I am a coat"
         receipt := &coat
// receipt points to coat not a copy of
         receipt2 := &coat

         fmt.Println(*receipt) 
		//  says tell me whats at the address of where I point to. Without the * this would give a weird output
         fmt.Println(*receipt2)
