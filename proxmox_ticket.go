package main

// type ProxmoxAPITicketResp struct {
// 	Data struct {
// 		CSRFPreventionToken string `json:"CSRFPreventionToken"`
// 		Ticket              string `json:"ticket"`
// 		Username            string `json:"username"`
// 	} `json:"data"`
// }

// type Ticket struct {
// 	ticket string
// 	mux    sync.Mutex
// }

// func (c *Ticket) Get() string {
// 	c.mux.Lock()
// 	defer c.mux.Unlock()
// 	return c.ticket
// }

// func (c *Ticket) Set(ticket string) {
// 	c.mux.Lock()
// 	c.ticket = ticket
// 	c.mux.Unlock()
// }

// func getNewTicket(api ProxmoxAPI) (string, error) {
// 	url := "https://" + api.Host + ":" + api.Port + "/api2/json/access/ticket?username=" + api.User + "&password=" + api.Pass
// 	c := &tls.Config{
// 		InsecureSkipVerify: true,
// 	}
// 	tr := &http.Transport{TLSClientConfig: c}
// 	client := &http.Client{Transport: tr}
// 	//client := &http.Client{}
// 	req, err := http.NewRequest("POST", url, nil)
// 	if err != nil {
// 		return "", err
// 	}
// 	req.Header.Add("Content-type", "application/json")

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close() //Close the resp body when finished

// 	respBody := ProxmoxAPITicketResp{}
// 	err = json.NewDecoder(resp.Body).Decode(&respBody)
// 	if err != nil {
// 		return "", err
// 	}

// 	return respBody.Data.Ticket, nil

// 	// // //Check if success, print errors from api if not.
// 	// // if !respBody.Success {
// 	// // 	// for _, e := range respBody.Errors {
// 	// // 	// 	log.Printf("Error code %d: %s\n", e.Code, e.Message)
// 	// // 	// }
// 	// // 	return fmt.Errorf("Api call failed")
// 	// // }

// 	// if len(respBody.Result) == 0 {
// 	// 	return []CloudflareRecord{}, errors.New("Could not find any TXT records")
// 	// }

// 	// return respBody.Result, nil

// 	// resp, err := http.Get("http://api.ipify.org")
// 	// if err != nil {
// 	// 	return "", err
// 	// }
// 	// publicIPRaw, err := ioutil.ReadAll(resp.Body)
// 	// resp.Body.Close()
// 	// if err != nil {
// 	// 	return "", err
// 	// }

// 	// return string(publicIPRaw), nil
// }

// func refreshTicket(api *ProxmoxAPI) {
// 	ticker := time.NewTicker(5 * time.Second)
// 	for {
// 		newTicket, err := getNewTicket(api)
// 		if err != nil {
// 			fmt.Println("Could not retrieve new ticket. Retry on next check...")
// 		}
// 		currentTicket.Set(newTicket)
// 		fmt.Println("Refreshed ticket")
// 		<-ticker.C
// 	}
// }

// func waitForTicket(currentTicket *Ticket) {
// 	ticker := time.NewTicker(time.Second)
// 	for {
// 		if currentTicket.Get() != "" {
// 			break
// 		} else {
// 			fmt.Println("Waiting to get ticket")
// 		}
// 		<-ticker.C
// 	}
// }
