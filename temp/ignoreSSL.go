/* Example healthcheck with ignore ssl , slice and for loop */

 package main 

 import ( 

         "crypto/tls" 
         "fmt" 
         "net/http" 
         "os" 
 ) 

 func main() { 

         transCfg := &http.Transport{ 

                 TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore SSL certificates 

         } 
         client := &http.Client{Transport: transCfg} 

         urls := [4]string{"https://host1:7999","https://host2:7999", "https://host3:7999", "https://host4:7999"} 

          for _,url := range urls { 

         response, err := client.Get(url) 

         if err != nil { 
                 fmt.Println(err) 
                 os.Exit(1) 
         }
         defer response.Body.Close() 
         status := response.Status 

         if err != nil { 
                 fmt.Println(err) 
                 os.Exit(1) 
         } 
        fmt.Println( url, ":",  status) 
   } 
 } 
